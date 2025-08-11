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

// Enviroment variables which can be read from the command line evn.
const TaikunEmailEnvVar = "TAIKUN_EMAIL"
const TaikunPasswordEnvVar = "TAIKUN_PASSWORD"
const TaikunAuthMode = "TAIKUN_AUTH_MODE"
const TaikunAccessKey = "TAIKUN_ACCESS_KEY"
const TaikunSecretKey = "TAIKUN_SECRET_KEY"
const TaikunApiHostEnvVar = "TAIKUN_API_HOST"

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
	// This is not a /auth/login or /auth/refresh request
	if strings.Contains(req.URL.Path, "/api/v1/taikun-lb") {
		req.Header.Set("Authorization", "Bearer "+c.Client.token)
	} else if req.URL.Path != "/api/v1/auth/login" && req.URL.Path != "/api/v1/auth/refresh" {
		if c.Client.token == "" { // We do not have a token, get a lock
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.Client.token == "" {

				var loginCmd *taikuncore.LoginCommand
				if c.Client.authMode != "" {
					// I have an authMode specified (eg. "default" or "keycloak")
					// Use for modes: keycloak, default, token, autoscaling
					loginCmd = &taikuncore.LoginCommand{
						SecretKey: *taikuncore.NewNullableString(&c.Client.secretKey),
						AccessKey: *taikuncore.NewNullableString(&c.Client.accessKey),
						Mode:      *taikuncore.NewNullableString(&c.Client.authMode),
					}
				} else {
					// I do not have an authMode specified
					// Use authentication with email + password
					loginCmd = &taikuncore.LoginCommand{
						Email:    *taikuncore.NewNullableString(&c.Client.email),
						Password: *taikuncore.NewNullableString(&c.Client.password),
					}
				}
				result, body, err := c.Client.Client.AuthManagementAPI.AuthLogin(req.Context()).LoginCommand(*loginCmd).Execute()
				if err != nil {
					return nil, CreateError(body, err)
				}
				c.Client.token = result.Token
				c.Client.refreshToken = *result.RefreshToken.Get()
			}
		} else { // We do have a token
			if c.Client.token != "" && c.hasTokenExpired() {
				result, body, err := c.Client.Client.AuthManagementAPI.AuthRefresh(req.Context()).RefreshTokenCommand(taikuncore.RefreshTokenCommand{
					RefreshToken: *taikuncore.NewNullableString(&c.Client.refreshToken),
					Token:        *taikuncore.NewNullableString(&c.Client.token),
				}).Execute()
				if err != nil {
					return nil, CreateError(body, err)
				}
				c.Client.token = result.Token
				c.Client.refreshToken = *result.RefreshToken.Get()
			}
		}
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

// NewClientFromToken is a helper function not intended to be used by the user.
// It returns a client based on the token provided.
func NewClientFromToken(token string, apiHost string) *Client {

	// Create a configuration object for the Taikun Web API
	cfg := taikuncore.NewConfiguration()
	cfg.Host = apiHost
	cfg.Scheme = "https"

	// Create a configuration object for the Taikun Showback service
	cfg2 := taikunshowback.NewConfiguration()
	cfg2.Host = apiHost
	cfg2.Scheme = "https"

	// Actually create the new client
	client := &Client{
		token: token,
	}

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

// NewClientFromCredentials is a helper function not intended to be used by the user.
// It returns a client based on the authMode and provided credentials
func NewClientFromCredentials(email string, password string, accessKey string, secretKey string, authMode string, apiHost string) *Client {

	// Create a configuration object for the Taikun Web API
	cfg := taikuncore.NewConfiguration()
	cfg.Host = apiHost
	cfg.Scheme = "https"

	// Create a configuration object for the Taikun Showback service
	cfg2 := taikunshowback.NewConfiguration()
	cfg2.Host = apiHost
	cfg2.Scheme = "https"

	// Actually create the new client
	client := &Client{
		email:     email,
		password:  password,
		accessKey: accessKey,
		secretKey: secretKey,
		authMode:  authMode,
	}

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

// NewClient returns a client authenticated based on your environment variables defined above.
// It is intended to be called by other modules or packages.
func NewClient() *Client {

	// What endpoint are we connecting to?
	// If not defined, default is production - api.taikun.cloud
	apiHost := os.Getenv(TaikunApiHostEnvVar)
	if apiHost == "" {
		apiHost = "api.taikun.cloud"
	}

	// What Authorization mode was chosen?
	// Default, empty or not set = Try to use Username and Password
	taikunAuthMode, customMode := os.LookupEnv(TaikunAuthMode)
	if !customMode || taikunAuthMode == "" || taikunAuthMode == "default" {
		email := os.Getenv(TaikunEmailEnvVar)
		password := os.Getenv(TaikunPasswordEnvVar)
		if email == "" || password == "" {
			fmt.Println("Please set your Taikun credentials. Password or Email was empty.")
			os.Exit(1)
		}
		return NewClientFromCredentials(email, password, "", "", "", apiHost) // Create and return the client
	}

	// Any other mode was chosen. Try to use AccessKey + Secret key.
	accessKey := os.Getenv(TaikunAccessKey)
	secretKey := os.Getenv(TaikunSecretKey)
	if accessKey == "" || secretKey == "" {
		fmt.Println("Please set your Taikun credentials. AccessKey or SecretKey was empty.")
		os.Exit(1)
	}
	return NewClientFromCredentials("", "", accessKey, secretKey, taikunAuthMode, apiHost) // Create and return the client
}
