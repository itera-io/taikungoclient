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
	taikunGoClient := Client{
		authMode: strings.ToLower(os.Getenv(TaikunAuthModeEnvVar)),
	}

	switch taikunGoClient.authMode {
	case "taikun":
		taikunGoClient.email = os.Getenv(TaikunEmailEnvVar)
		taikunGoClient.password = os.Getenv(TaikunPasswordEnvVar)
	case "keycloak":
		taikunGoClient.email = os.Getenv(TaikunKeycloakEmailEnvVar)
		taikunGoClient.password = os.Getenv(TaikunKeycloakPasswordEnvVar)
	case "autoscaling", "token":
		taikunGoClient.accessKey = os.Getenv(TaikunAccessKeyEnvVar)
		taikunGoClient.secretKey = os.Getenv(TaikunSecretKeyEnvVar)
	default:
		return nil, fmt.Errorf(
			`Authentication mode must be one of normal, keycloak, autoscaling or token.
Set your authentication with the following environment variable:
%s

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
	}

	apiHost := os.Getenv(TaikunApiHostEnvVar)

	return InitializeClient(&taikunGoClient, apiHost), nil
}

func InitializeClient(taikunGoClient *Client, apiHost string) *Client {
	transportConfig := client.DefaultTransportConfig()
	showbackTransportConfig := showbackclient.DefaultTransportConfig()
	if apiHost != "" {
		transportConfig = transportConfig.WithHost(apiHost)
		showbackTransportConfig = showbackTransportConfig.WithHost(apiHost)
	}

	taikunGoClient.Client = client.NewHTTPClientWithConfig(nil, transportConfig)
	taikunGoClient.ShowbackClient = showbackclient.NewHTTPClientWithConfig(nil, showbackTransportConfig)

	return taikunGoClient
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

		authMode := os.Getenv(TaikunAuthModeEnvVar)
		accessKey := os.Getenv(TaikunAccessKeyEnvVar)
		secretKey := os.Getenv(TaikunSecretKeyEnvVar)

		switch authMode {

		case "keycloak":
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(Version).WithBody(
					&models.LoginCommand{Email: apiClient.email, Password: apiClient.password, Mode: "keycloak"},
				), nil,
			)
			if err != nil {
				return err
			}

			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken

		case "autoscaling":
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(Version).WithBody(
					&models.LoginCommand{AccessKey: accessKey, SecretKey: secretKey, Mode: "autoscaling"},
				), nil,
			)
			if err != nil {
				return err
			}

			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken

		case "token":
			content := models.LoginCommand{AccessKey: accessKey, SecretKey: secretKey, Mode: "token"}
			fmt.Println(content)
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(Version).WithBody(
					&models.LoginCommand{AccessKey: accessKey, SecretKey: secretKey, Mode: "token"},
				), nil,
			)
			if err != nil {
				return err
			}

			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken

		default:
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(Version).WithBody(
					&models.LoginCommand{Email: apiClient.email, Password: apiClient.password},
				), nil,
			)
			if err != nil {
				return err
			}

			apiClient.token = loginResult.Payload.Token
			apiClient.refreshToken = loginResult.Payload.RefreshToken

		}
	}

	if apiClient.hasTokenExpired() {
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
	}

	err := request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", apiClient.token))
	if err != nil {
		return err
	}

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
