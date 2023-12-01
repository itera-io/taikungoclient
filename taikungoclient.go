package taikungoclient

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
					// I do not have and authMode specified
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
				c.Client.token = *result.Token.Get()
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
				c.Client.token = *result.Token.Get()
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

// Error is a method of struct taikunError used to pretty print the recieved HTTP error.
func (e *taikunError) Error() string {
	return fmt.Sprintf("Taikun Error: %s (HTTP %d)", e.Message, e.HTTPStatusCode)
}

// CreateError is a helper function to convert an unsuccessful HTTP response to a taikunError struct.
// Function is exported because it is used by the Terraform taikun provider and Taikun CLI to show errors.
func CreateError(resp *http.Response, err error) error {
	if err == nil || resp == nil {
		return err
	}

	// Decode into map for simple pretty printing (readMap["detail"]) - disabled because of background compatibility
	//var readMap map[string]interface{}
	//err2 := json.NewDecoder(resp.Body).Decode(&readMap)
	// Read into byte array
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		// Reading/decoding failed. Give the user at least the deformed response.
		return fmt.Errorf("Reply in unexpected format. Pasting the raw error: %v %v", resp.Body, err2)
	}
	return &taikunError{
		HTTPStatusCode: resp.StatusCode,
		Message:        fmt.Sprintf("%s", string(body)),
	}
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
			//panic(fmt.Errorf("Please set your Taikun credentials. Password or Email was empty."))
		}
		return NewClientFromCredentials(email, password, "", "", "", apiHost) // Create and return the client
	}

	// Any other mode was chosen. Try to use AccessKey + Secret key.
	accessKey := os.Getenv(TaikunAccessKey)
	secretKey := os.Getenv(TaikunSecretKey)
	if accessKey == "" || secretKey == "" {
		// panic(fmt.Errorf("Please set your Taikun credentials. AccessKey or SecretKey was empty."))
		fmt.Println("Please set your Taikun credentials. AccessKey or SecretKey was empty.")
		os.Exit(1)
	}
	return NewClientFromCredentials("", "", accessKey, secretKey, taikunAuthMode, apiHost) // Create and return the client
}
