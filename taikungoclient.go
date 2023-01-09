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
const TaikunAuthMethodEnvVar = "TAIKUN_AUTH_METHOD"
const TaikunAccessKeyEnvVar = "TAIKUN_ACCESS_KEY"
const TaikunSecretKeyEnvVar = "TAIKUN_SECRET_KEY"

// Wrapper around the generated Taikungoclient to include authentication
type Client struct {
	Client         *client.Taikungoclient
	ShowbackClient *showbackclient.Showbackgoclient

	email               string
	password            string
	useKeycloakEndpoint bool

	token        string
	refreshToken string
}

// Create a new authenticated Taikungoclient from environment variables.
// Taikun or Keycloak credentials environment variables must be set
func NewClient() (*Client, error) {
	email, keycloakEnabled := os.LookupEnv(TaikunKeycloakEmailEnvVar)
	password := os.Getenv(TaikunKeycloakPasswordEnvVar)

	if !keycloakEnabled {
		email = os.Getenv(TaikunEmailEnvVar)
		password = os.Getenv(TaikunPasswordEnvVar)
	}

	if email == "" || password == "" {
		return nil, fmt.Errorf(
			`Please set your authentication method (normal, keycloak, autoscaling or token) with the following environment variable:
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
			TaikunAuthMethodEnvVar,
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

	return NewClientFromCredentials(email, password, keycloakEnabled, apiHost)
}

func NewClientFromCredentials(email string, password string, keycloakEnabled bool, apiHost string) (*Client, error) {
	transportConfig := client.DefaultTransportConfig()
	showbackTransportConfig := showbackclient.DefaultTransportConfig()
	if apiHost != "" {
		transportConfig = transportConfig.WithHost(apiHost)
		showbackTransportConfig = showbackTransportConfig.WithHost(apiHost)
	}

	return &Client{
		Client:              client.NewHTTPClientWithConfig(nil, transportConfig),
		ShowbackClient:      showbackclient.NewHTTPClientWithConfig(nil, showbackTransportConfig),
		email:               email,
		password:            password,
		useKeycloakEndpoint: keycloakEnabled,
	}, nil
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

            authMethod := os.Getenv(TaikunAuthMethodEnvVar)
            accessKey := os.Getenv(TaikunAccessKeyEnvVar)
            secretKey := os.Getenv(TaikunSecretKeyEnvVar)

            switch authMethod {

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
