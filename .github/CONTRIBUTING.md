## Contributing

The client's code should not be modified manually, it is automatically
generated using [openapi-generate](https://openapi-generator.tech/docs/installation). Thus any manual changes
would be overwritten by the next update.

A Github Action is set up to fetch the `swagger.json` openapi v3 API specification
file from the `$SWAGGER_URL` and `$SHOWBACK_SWAGGER_URL` variable daily at midnight.
A checksum of the swagger file is then computed and [stored](swagger_sha1sum.txt).
If the new checksum doesn't match the Go client is regenerated
by a [bash script](generate_client.sh).

This script makes no modifications to the JSON swagger file before using
openapi-generate to generate the client. In the past, modifications were necessary.
When the API was at openapi version 2, the clilent generation was done by
[go-swagger](https://goswagger.io/). 

This Go client is used when developing the
[taikun-cli](https://github.com/itera-io/taikun-cli),
[terraform-provider-taikun](https://github.com/itera-io/terraform-provider-taikun) and
taikun sweeper.

The script may need to be updated in the future as the API evolves.
