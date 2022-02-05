# taikungoclient

Go client library for the Taikun API generated using [go-swagger](https://goswagger.io/).

This package is used by the
[taikun-cli](https://github.com/itera-io/taikun-cli) and
[terraform-provider-taikun](https://github.com/itera-io/terraform-provider-taikun)
projects.

## Example use
This example Go program uses the `taikungoclient` package to output your Taikun username.
```go
package main

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/itera-io/taikungoclient/client"
	"github.com/itera-io/taikungoclient/client/auth"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
)

// Taikun API version
const taikunApiVersion = "1"

// Wrapper around the Taikungoclient struct to include authentication data
type taikunClient struct {
	// Taikun API client
	Client *client.Taikungoclient

	email    string
	password string

	token string
}

func main() {
	// Initialize Taikun API client with Taikun credentials
	taikunClient := &taikunClient{
		Client:   client.NewHTTPClient(nil), // default client
		email:    "jdoe@example.com",
		password: "taikunPassword123",
	}

	// User information parameters
	params := users.NewUsersDetailsParams().WithV(taikunApiVersion)

	// Send user info request
	response, err := taikunClient.Client.Users.UsersDetails(params, taikunClient)
	if err != nil {
		log.Fatal(err)
	}

	// Print username
	fmt.Println(response.Payload.Data.Username)
}

// Taikun authentication function used by go-openapi
func (c *taikunClient) AuthenticateRequest(request runtime.ClientRequest, _ strfmt.Registry) error {
	if len(c.token) == 0 {
		// Authentication request body
		authBody := models.LoginCommand{
			Email:    c.email,
			Password: c.password,
		}

		// Authentication request parameters
		authParams := auth.NewAuthLoginParams().WithV(taikunApiVersion)
		authParams = authParams.WithBody(&authBody)

		// Send authentication request
		authResult, err := c.Client.Auth.AuthLogin(authParams, nil)
		if err != nil {
			return err
		}

		// Save authentication token
		c.token = authResult.Payload.Token
	}

	// Add authorization header to request
	return request.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", c.token))
}
```
