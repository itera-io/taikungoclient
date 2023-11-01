## Contributing

The client's code should not be modified manually, it is automatically
generated using [go-swagger](https://goswagger.io/). Thus any manual changes
would be overwritten by the next update.

A Github Action is set up to fetch the `swagger.json` file from the
`$SWAGGER_URL` variable daily.  A checksum of the swagger file is then computed
and [stored](swagger_sha1sum.txt).  If the new checksum doesn't match
the Go client is regenerated by a [bash script](generate_client.sh).

This script makes some type modifications to the JSON swagger file before using
go-swagger to generate the client. These modifications are necessary to avoid
Go type errors when developing the
[taikun-cli](https://github.com/itera-io/taikun-cli) and
[terraform-provider-taikun](https://github.com/itera-io/terraform-provider-taikun).
The script may need to be updated in the future as the API evolves.