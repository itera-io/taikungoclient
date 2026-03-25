package taikungoclient

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	taikuncore "github.com/itera-io/taikungoclient/client"             // Main Taikun Web API
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient" // API for Taikun showback service
)

// Environment variables for configuring authentication.
const TaikunEmailEnvVar = "TAIKUN_EMAIL"
const TaikunPasswordEnvVar = "TAIKUN_PASSWORD"
const TaikunAuthMode = "TAIKUN_AUTH_MODE"
const TaikunAccessKey = "TAIKUN_ACCESS_KEY"
const TaikunSecretKey = "TAIKUN_SECRET_KEY"
const TaikunApiHostEnvVar = "TAIKUN_API_HOST"
const TaikunAccountName = "TAIKUN_ACCOUNT_NAME"

type jwtData struct {
	Nameid     string `json:"nameid"`
	Email      string `json:"email"`
	UniqueName string `json:"unique_name"`
	Role       string `json:"role"`
	Nbf        int    `json:"nbf"`
	Exp        int    `json:"exp"`
	Iat        int    `json:"iat"`
}

type taikunError struct {
	HTTPStatusCode int
	Message        string
	GoError        error // Optional underlying Go error (e.g., network, decoding)

	Method    string // HTTP method of the failed request
	URL       string // URL of the failed request (with sensitive params redacted)
	RequestID string // Request ID header from the server, if available
}

type customTransport struct {
	Transport http.RoundTripper
	Client    *Client
	mu        sync.Mutex
}

// Client is a struct representing an authenticated connection to simultaneously both Taikun Web API and Showback client
type Client struct {
	// Set by the program
	Client          *taikuncore.APIClient
	ShowbackClient  *taikunshowback.APIClient
	CustomTransport *customTransport
	token           string
	refreshToken    string

	// Set by user
	email     string
	password  string
	accessKey string
	secretKey string
	authMode  string
}

// Getter for token (Used in CLI usertoken get-bearer)
func (c *Client) GetToken() string {
	return c.token
}

// Transport wrapper
func (c *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.Client.accessKey != "" && c.Client.secretKey != "" {
		// Robot users: HTTP Basic Auth on every request, no JWT
		req.SetBasicAuth(c.Client.accessKey, c.Client.secretKey)
	} else if strings.Contains(req.URL.Path, "/api/v1/taikun-lb") {
		req.Header.Set("Authorization", "Bearer "+c.Client.token)
	} else if req.URL.Path != "/api/v1/auth/login" && req.URL.Path != "/api/v1/auth/refresh" {
		c.mu.Lock()
		if c.Client.token == "" {
			// No token yet: login with email + password
			loginCmd := taikuncore.LoginCommand{
				Email:    *taikuncore.NewNullableString(&c.Client.email),
				Password: *taikuncore.NewNullableString(&c.Client.password),
			}
			if c.Client.authMode != "" {
				loginCmd.Mode = *taikuncore.NewNullableString(&c.Client.authMode)
			}
			result, body, err := c.Client.Client.AuthManagementAPI.AuthLogin(req.Context()).LoginCommand(loginCmd).Execute()
			if err != nil {
				c.mu.Unlock()
				return nil, CreateError(body, err)
			}
			c.Client.token = result.GetToken()
			if result.RefreshToken.Get() != nil {
				c.Client.refreshToken = *result.RefreshToken.Get()
			}
		} else if c.hasTokenExpired() {
			// Token expired: refresh
			result, body, err := c.Client.Client.AuthManagementAPI.AuthRefresh(req.Context()).RefreshTokenCommand(taikuncore.RefreshTokenCommand{
				RefreshToken: *taikuncore.NewNullableString(&c.Client.refreshToken),
				Token:        *taikuncore.NewNullableString(&c.Client.token),
			}).Execute()
			if err != nil {
				c.mu.Unlock()
				return nil, CreateError(body, err)
			}
			c.Client.token = result.GetToken()
			if result.RefreshToken.Get() != nil {
				c.Client.refreshToken = *result.RefreshToken.Get()
			}
		}
		c.mu.Unlock()
		req.Header.Set("Authorization", "Bearer "+c.Client.token)
	}

	return c.Transport.RoundTrip(req)
}

func (transport *customTransport) hasTokenExpired() bool {
	jwtSplit := strings.Split(transport.Client.token, ".")
	if len(jwtSplit) != 3 {
		return true
	}

	data, err := base64.RawURLEncoding.DecodeString(jwtSplit[1])
	if err != nil {
		return true
	}

	jwtData := jwtData{}

	err = json.Unmarshal(data, &jwtData)
	if err != nil {
		return true
	}

	tm := time.Unix(int64(jwtData.Exp), 0)

	return tm.Before(time.Now())
}

// Error returns a readable string representation of a taikunError.
func (e *taikunError) Error() string {
	var base string
	if e.GoError != nil {
		base = fmt.Sprintf("Taikun Error: %s (HTTP %d) (GO_ERROR %v)", e.Message, e.HTTPStatusCode, e.GoError)
	} else {
		base = fmt.Sprintf("Taikun Error: %s (HTTP %d)", e.Message, e.HTTPStatusCode)
	}
	// Append HTTP method, URL if available
	if e.Method != "" || e.URL != "" {
		base += fmt.Sprintf(" (%s %s)", e.Method, e.URL)
	}
	return base
}

// CreateError returns a taikunError describing the HTTP/Go-level error.
// Returns nil for successful (2xx) responses with no Go error.
// Function is exported because it is used by the Terraform taikun provider and Taikun CLI to show errors.
func CreateError(resp *http.Response, err error) error {
	// Capture request metadata for debugging
	var method, urlStr string
	if resp != nil && resp.Request != nil {
		if resp.Request.Method != "" {
			method = resp.Request.Method
		}
		if resp.Request.URL != nil {
			urlStr = redactURL(resp.Request.URL)
		}
	}

	// Case 1: No response at all — wrap the Go-level error into taikunError
	if resp == nil {
		if err != nil {
			return &taikunError{
				HTTPStatusCode: 0, // no HTTP status available
				Message:        "no HTTP response",
				GoError:        err,
				Method:         method,
				URL:            urlStr,
			}
		}
		return &taikunError{
			HTTPStatusCode: 0,
			Message:        "unknown error: both response and error are nil",
			GoError:        nil,
			Method:         method,
			URL:            urlStr,
		}
	}

	// Case 2: Successful HTTP response (2xx) and no Go-level error → return nil
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	var bodyBytes []byte
	var readErr error

	// Read the response body (if present)
	if resp.Body != nil {
		defer func() { _ = resp.Body.Close() }()
		bodyBytes, readErr = io.ReadAll(resp.Body)
	}

	bodyStr := strings.TrimSpace(string(bodyBytes))

	// Remove duplicate Go-level errors that simply restate HTTP status
	goError := err
	if err != nil {
		statusText := http.StatusText(resp.StatusCode)
		if strings.Contains(err.Error(), strconv.Itoa(resp.StatusCode)) &&
			strings.Contains(err.Error(), statusText) {
			goError = nil
		}
	}

	// Case 3: Response body unreadable
	if readErr != nil {
		return &taikunError{
			HTTPStatusCode: resp.StatusCode,
			Message:        fmt.Sprintf("failed to read error response body: %v", readErr),
			GoError:        goError,
			Method:         method,
			URL:            urlStr,
		}
	}

	// Case 4: Empty response body
	if bodyStr == "" {
		return &taikunError{
			HTTPStatusCode: resp.StatusCode,
			Message:        "empty response body",
			GoError:        goError,
			Method:         method,
			URL:            urlStr,
		}
	}

	// Case 5: JSON error body with known fields ("title", "detail", "message", "error")
	var parsed map[string]interface{}
	if json.Unmarshal(bodyBytes, &parsed) == nil {
		var title, message string

		if val, ok := parsed["title"]; ok {
			title = fmt.Sprintf("%v", val)
		}
		for _, key := range []string{"detail", "message", "error"} {
			if val, ok := parsed[key]; ok {
				message = fmt.Sprintf("%v", val)
				break
			}
		}

		// Sanitize newlines
		title = strings.ReplaceAll(title, "\n", " ")
		message = strings.ReplaceAll(message, "\n", " ")

		// Construct combined message
		var combinedMsg string
		switch {
		case title != "" && message != "":
			combinedMsg = fmt.Sprintf("(TITLE %s) (DETAIL %s)", title, message)
		case message != "":
			combinedMsg = message
		case title != "":
			combinedMsg = title
		default:
			combinedMsg = fmt.Sprintf("(RESPONSE %s)", strings.ReplaceAll(string(bodyBytes), "\n", " "))
		}

		return &taikunError{
			HTTPStatusCode: resp.StatusCode,
			Message:        combinedMsg,
			GoError:        goError,
			Method:         method,
			URL:            urlStr,
		}
	}

	// Case 6: Fallback — plain text body, sanitized
	bodyStr = strings.ReplaceAll(bodyStr, "\n", " ")

	return &taikunError{
		HTTPStatusCode: resp.StatusCode,
		Message:        bodyStr,
		GoError:        goError,
		Method:         method,
		URL:            urlStr,
	}
}

// redactURL removes sensitive query values but keeps path visible
func redactURL(u *urlpkg.URL) string {
	if u == nil {
		return ""
	}
	cp := *u
	q := cp.Query()
	for _, k := range []string{"token", "access_token", "auth", "authorization", "api_key", "apikey"} {
		if _, ok := q[k]; ok {
			q.Set(k, "REDACTED")
		}
	}
	cp.RawQuery = q.Encode()
	return cp.String()
}

// NewClientFromToken creates a client using an existing JWT token.
func NewClientFromToken(token string, apiHost string) *Client {
	return newClient(apiHost, func(c *Client) {
		c.token = token
	})
}

// NewClientFromCredentials creates a client using email + password authentication.
// authMode is optional — pass an empty string for the default mode, or a custom
// mode such as "autoscaler".
func NewClientFromCredentials(email string, password string, authMode string, apiHost string) *Client {
	return newClient(apiHost, func(c *Client) {
		c.email = email
		c.password = password
		c.authMode = strings.TrimSpace(authMode)
	})
}

// NewClientFromAccessKey creates a client using robot user credentials (access key + secret key).
// Robot users authenticate via HTTP Basic Auth on every request — no JWT or refresh tokens.
func NewClientFromAccessKey(accessKey string, secretKey string, apiHost string) *Client {
	return newClient(apiHost, func(c *Client) {
		c.accessKey = accessKey
		c.secretKey = secretKey
	})
}

// newClient is the internal constructor shared by all public constructors.
func newClient(apiHost string, configure func(*Client)) *Client {
	cfg := taikuncore.NewConfiguration()
	cfg.Host = apiHost
	cfg.Scheme = "https"

	cfg2 := taikunshowback.NewConfiguration()
	cfg2.Host = apiHost
	cfg2.Scheme = "https"

	client := &Client{}
	configure(client)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	clientTransport := &customTransport{
		Transport: tr,
		Client:    client,
	}
	cfg.HTTPClient = &http.Client{Transport: clientTransport}
	cfg2.HTTPClient = &http.Client{Transport: clientTransport}
	client.Client = taikuncore.NewAPIClient(cfg)
	client.ShowbackClient = taikunshowback.NewAPIClient(cfg2)

	return client
}

// NewClient returns a client authenticated based on environment variables.
//
// Authentication is auto-detected:
//   - If TAIKUN_ACCESS_KEY and TAIKUN_SECRET_KEY are set, robot user auth (HTTP Basic) is used.
//   - Otherwise, TAIKUN_EMAIL and TAIKUN_PASSWORD are used for user auth (JWT).
//   - TAIKUN_API_HOST defaults to "api.taikun.cloud" if not set.
func NewClient() *Client {
	// check if API endpoint is specified, otherwise use standard one
	apiHost := os.Getenv(TaikunApiHostEnvVar)
	if apiHost == "" {
		apiHost = "api.taikun.cloud"
	}

	// check for account name
	an := os.Getenv(TaikunAccountName)
	if an == "" {
		fmt.Printf("Taikun account name must be set in '%s'\n", TaikunAccountName)
		os.Exit(1)
		return nil
	}

	// Check for robot user credentials (access key + secret key) first
	accessKey := os.Getenv(TaikunAccessKey)
	secretKey := os.Getenv(TaikunSecretKey)
	if accessKey != "" && secretKey != "" {
		return NewClientFromAccessKey(accessKey, secretKey, apiHost)
	}

	// Fall back to email + password authentication
	email := os.Getenv(TaikunEmailEnvVar)
	password := os.Getenv(TaikunPasswordEnvVar)
	if email != "" && password != "" {
		authMode := os.Getenv(TaikunAuthMode)
		return NewClientFromCredentials(email, password, authMode, apiHost)
	}

	fmt.Printf("Please set your Taikun credentials. Either '%s' + '%s' or '%s' + '%s' must be set.\n",
		TaikunAccessKey, TaikunSecretKey, TaikunEmailEnvVar, TaikunPasswordEnvVar)
	os.Exit(1)
	return nil
}
