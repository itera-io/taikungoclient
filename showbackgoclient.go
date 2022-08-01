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
	"github.com/itera-io/taikungoclient/client/keycloak"
	"github.com/itera-io/taikungoclient/models"
	"github.com/itera-io/taikungoclient/showbackclient"
)

// Wrapper around the generated Showbackgoclient to include authentication
type ShowbackClient struct {
	Client *showbackclient.Showbackgoclient

	email               string
	password            string
	useKeycloakEndpoint bool

	token        string
	refreshToken string
}

// Create a new authenticated Showbackgoclient, Taikun or Keycloak credentials
// environment variables must be set
func NewShowbackClient() (*ShowbackClient, error) {
	email, keycloakEnabled := os.LookupEnv(TaikunKeycloakEmailEnvVar)
	password := os.Getenv(TaikunKeycloakPasswordEnvVar)

	if !keycloakEnabled {
		email = os.Getenv(TaikunEmailEnvVar)
		password = os.Getenv(TaikunPasswordEnvVar)
	}

	if email == "" || password == "" {
		return nil, fmt.Errorf(
			`Please set your Taikun credentials.

To authenticate with your Taikun account, set the following environment variables:
%s
%s

To authenticate with Keycloak, set the following environment variables:
%s
%s

To override the default API host, set the following environment variable:
%s (default value is: %s)`,
			TaikunEmailEnvVar,
			TaikunPasswordEnvVar,
			TaikunKeycloakEmailEnvVar,
			TaikunKeycloakPasswordEnvVar,
			TaikunApiHostEnvVar,
			client.DefaultHost,
		)
	}

	transportConfig := showbackclient.DefaultTransportConfig()
	if apiHost, apiHostIsSet := os.LookupEnv(TaikunApiHostEnvVar); apiHostIsSet {
		transportConfig = transportConfig.WithHost(apiHost).WithBasePath("/showback")
	}

	return &ShowbackClient{
		Client:              showbackclient.NewHTTPClientWithConfig(nil, transportConfig),
		email:               email,
		password:            password,
		useKeycloakEndpoint: keycloakEnabled,
	}, nil
}

// Authenticate a ClientRequest, obtain a new token if the current one is
// expired
func (showbackClient *ShowbackClient) AuthenticateRequest(request runtime.ClientRequest, _ strfmt.Registry) error {
	if len(showbackClient.token) == 0 {
		if !showbackClient.useKeycloakEndpoint {
			apiClient, err := NewClient()
			if err != nil {
				return err
			}
			loginResult, err := apiClient.Client.Auth.AuthLogin(
				auth.NewAuthLoginParams().WithV(Version).WithBody(
					&models.LoginCommand{Email: showbackClient.email, Password: showbackClient.password},
				), nil,
			)
			if err != nil {
				return err
			}

			showbackClient.token = loginResult.Payload.Token
			showbackClient.refreshToken = loginResult.Payload.RefreshToken
		} else {
			apiClient, err := NewClient()
			if err != nil {
				return err
			}
			loginResult, err := apiClient.Client.Keycloak.KeycloakLogin(
				keycloak.NewKeycloakLoginParams().WithV(Version).WithBody(
					&models.LoginWithKeycloakCommand{Email: showbackClient.email, Password: showbackClient.password},
				), nil,
			)
			if err != nil {
				return err
			}

			showbackClient.token = loginResult.Payload.Token
			showbackClient.refreshToken = loginResult.Payload.RefreshToken
		}
	}

	if showbackClient.hasTokenExpired() {
		apiClient, err := NewClient()
		if err != nil {
			return err
		}
		refreshResult, err := apiClient.Client.Auth.AuthRefreshToken(
			auth.NewAuthRefreshTokenParams().WithV(Version).WithBody(
				&models.RefreshTokenCommand{
					RefreshToken: showbackClient.refreshToken,
					Token:        showbackClient.token,
				}), nil,
		)
		if err != nil {
			return err
		}

		showbackClient.token = refreshResult.Payload.Token
		showbackClient.refreshToken = refreshResult.Payload.RefreshToken
	}

	err := request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", showbackClient.token))
	if err != nil {
		return err
	}

	return nil
}

func (showbackClient *ShowbackClient) hasTokenExpired() bool {
	jwtSplit := strings.Split(showbackClient.token, ".")
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
