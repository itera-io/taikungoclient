# taikungoclient

Go client library for the [Taikun API](https://api.taikun.cloud/) generated
using [go-swagger](https://goswagger.io/).

This package is used by the
[taikun-cli](https://github.com/itera-io/taikun-cli) and
[terraform-provider-taikun](https://github.com/itera-io/terraform-provider-taikun)
projects.

## API
The `taikungoclient` package provides the following constants, structs and
functions.

- `Version`: The API version, must be passed as an argument to client calls
    using `.WithV()` (see below for an example).
- `Client`: Wrapper struct around the generated Taikungoclient to include
    authentication.
- `NewClient()`: Create a new authenticated Taikungoclient. Taikun or Keycloak
    credentials environment variables must be set.
- `Client.AuthenticateRequest(request, ...)`: Authenticate a ClientRequest,
    obtain a new token if the current one is expired. You probably won't need
    to use this one as it is called by the client structs located in
    [./client](./client).

## Example use
The following snippet is from the
[Taikun CLI](https://github.com/itera-io/taikun-cli). It implements the
`taikun whoami` command which prints the name of the CLI's authenticated user.

```go
package whoami

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

func NewCmdWhoAmI() *cobra.Command {
	cmd := cobra.Command{
		Use:   "whoami",
		Short: "Print username",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return whoAmIRun()
		},
	}

	return &cmd
}

func whoAmIRun() (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(taikungoclient.Version)

	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		out.Println(response.Payload.Data.Username)
	}

	return
}
```
