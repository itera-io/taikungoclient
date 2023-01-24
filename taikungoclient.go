package taikungoclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/itera-io/taikungoclient/client"
	"github.com/itera-io/taikungoclient/client/auth"
	"github.com/itera-io/taikungoclient/models"
	"github.com/itera-io/taikungoclient/showbackclient"
)

// API version
const Version = "1"

// Help message displayed when user is not authenticated
var helpMessage = fmt.Errorf(
	`Authentication mode must be one of taikun, keycloak, autoscaling or token.
Set your authentication with the following environment variable:
%s (default value is taikun if not defined)

For normal mode, please set your Taikun credentials with the variables:
%s
%s

To authenticate with Keycloak, set the following variables:
%s
%s

To authenticate in autoscaling or token mode, set the access and secret keys with the following variables:
%s
%s

To override the default API host, set the following environment variable:
%s (default value is: %s)`,
	TaikunAuthModeEnvVar,
	TaikunEmailEnvVar,
	TaikunPasswordEnvVar,
	TaikunKeycloakEmailEnvVar,
	TaikunKeycloakPasswordEnvVar,
	TaikunAccessKeyEnvVar,
	TaikunSecretKeyEnvVar,
	TaikunApiHostEnvVar,
	client.DefaultHost,
)

// Credential environment variables
const TaikunEmailEnvVar = "TAIKUN_EMAIL"
const TaikunPasswordEnvVar = "TAIKUN_PASSWORD"
const TaikunKeycloakEmailEnvVar = "TAIKUN_KEYCLOAK_EMAIL"
const TaikunKeycloakPasswordEnvVar = "TAIKUN_KEYCLOAK_PASSWORD"
const TaikunApiHostEnvVar = "TAIKUN_API_HOST"
const TaikunAuthModeEnvVar = "TAIKUN_AUTH_MODE"
const TaikunAccessKeyEnvVar = "TAIKUN_ACCESS_KEY"
const TaikunSecretKeyEnvVar = "TAIKUN_SECRET_KEY"

// Wrapper around the generated Taikungoclient to include authentication
type Client struct {
	// Client structures
	Client         *client.Taikungoclient
	ShowbackClient *showbackclient.Showbackgoclient

	// Authentication mode, one of : taikun, keycloak, autoscaling, token
	authMode string

	// Taikun or Keycloak credentials
	email    string
	password string

	// Access and secret keys for autoscaling or token authentication mode
	accessKey string
	secretKey string

	// Taikun JWT tokens
	token        string
	refreshToken string
}

// Create a new authenticated Taikungoclient from environment variables.
// Taikun or Keycloak credentials environment variables must be set
func NewClient() (*Client, error) {
	apiClient := Client{
		authMode: strings.ToLower(os.Getenv(TaikunAuthModeEnvVar)),
	}

	// Set default auth mode to taikun
	if apiClient.authMode == "" {
		apiClient.authMode = "taikun"
	}

	switch apiClient.authMode {
	case "taikun":
		apiClient.email = os.Getenv(TaikunEmailEnvVar)
		apiClient.password = os.Getenv(TaikunPasswordEnvVar)
		if apiClient.email == "" || apiClient.password == "" {
			return nil, helpMessage
		}
	case "keycloak":
		apiClient.email = os.Getenv(TaikunKeycloakEmailEnvVar)
		apiClient.password = os.Getenv(TaikunKeycloakPasswordEnvVar)
		if apiClient.email == "" || apiClient.password == "" {
			return nil, helpMessage
		}
	case "autoscaling", "token":
		apiClient.accessKey = os.Getenv(TaikunAccessKeyEnvVar)
		apiClient.secretKey = os.Getenv(TaikunSecretKeyEnvVar)
		if apiClient.accessKey == "" || apiClient.secretKey == "" {
			return nil, helpMessage
		}
	default:
		return nil, helpMessage
	}

	apiHost := os.Getenv(TaikunApiHostEnvVar)

	return InitializeClient(&apiClient, apiHost), nil
}

func NewClientFromCredentials(email string, password string, keycloakEnabled bool, apiHost string) (*Client, error) {
	apiClient := Client{
		email:    email,
		password: password,
	}

	if keycloakEnabled {
		apiClient.authMode = "keycloak"
	} else {
		apiClient.authMode = "taikun"
	}

	return InitializeClient(&apiClient, apiHost), nil
}

func InitializeClient(apiClient *Client, apiHost string) *Client {
	transportConfig := client.DefaultTransportConfig()
	showbackTransportConfig := showbackclient.DefaultTransportConfig()
	if apiHost != "" {
		transportConfig = transportConfig.WithHost(apiHost)
		showbackTransportConfig = showbackTransportConfig.WithHost(apiHost)
	}

	apiClient.Client = client.NewHTTPClientWithConfig(nil, transportConfig)
	apiClient.ShowbackClient = showbackclient.NewHTTPClientWithConfig(nil, showbackTransportConfig)

	return apiClient
}

type jwtData struct {
	Nameid     string `json:"nameid"`
	Email      string `json:"email"`
	UniqueName string `json:"unique_name"`
	Role       string `json:"role"`
	Nbf        int    `json:"nbf"`
	Exp        int    `json:"exp"`
	Iat        int    `json:"iat"`
}

// Authenticate a ClientRequest, obtain a new token if the current one is
// expired. Only Client is authenticated since ShowbackClient uses Client as
// its runtime.ClientAuthInfoWriter.
func (apiClient *Client) AuthenticateRequest(request runtime.ClientRequest, _ strfmt.Registry) error {
	if len(apiClient.token) == 0 {

		var loginCommand models.LoginCommand

		switch apiClient.authMode {
		case "taikun":
			// Mode is null when authenticating with Taikun credentials
			loginCommand = models.LoginCommand{Email: apiClient.email, Password: apiClient.password}
		case "keycloak":
			loginCommand = models.LoginCommand{Email: apiClient.email, Password: apiClient.password, Mode: apiClient.authMode}
		default: // autoscaling or token
			loginCommand = models.LoginCommand{AccessKey: apiClient.accessKey, SecretKey: apiClient.secretKey, Mode: apiClient.authMode}
		}

		loginParams := auth.NewAuthLoginParams().WithV(Version).WithBody(&loginCommand)
		loginResult, err := apiClient.Client.Auth.AuthLogin(loginParams, nil)
		if err != nil {
			return err
		}

		apiClient.token = loginResult.Payload.Token
		apiClient.refreshToken = loginResult.Payload.RefreshToken
	}

	if apiClient.hasTokenExpired() {
		if err := apiClient.refresh(); err != nil {
			return err
		}
	}

	err := request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", apiClient.token))
	if err != nil {
		return err
	}

	return nil
}

func (apiClient *Client) refresh() error {
	refreshResult, err := apiClient.Client.Auth.AuthRefreshToken(
		auth.NewAuthRefreshTokenParams().WithV(Version).WithBody(
			&models.RefreshTokenCommand{
				RefreshToken: apiClient.refreshToken,
				Token:        apiClient.token,
			}), nil,
	)
	if err != nil {
		return err
	}

	apiClient.token = refreshResult.Payload.Token
	apiClient.refreshToken = refreshResult.Payload.RefreshToken

	return nil
}

func (apiClient *Client) hasTokenExpired() bool {
	jwtSplit := strings.Split(apiClient.token, ".")
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
