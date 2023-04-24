// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aws API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aws API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AwsAwsOwners(params *AwsAwsOwnersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsAwsOwnersOK, error)

	AwsAwsZoneList(params *AwsAwsZoneListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsAwsZoneListOK, error)

	AwsCreate(params *AwsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsCreateOK, error)

	AwsDeviceNames(params *AwsDeviceNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsDeviceNamesOK, error)

	AwsList(params *AwsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsListOK, error)

	AwsRegionList(params *AwsRegionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsRegionListOK, error)

	AwsUpdate(params *AwsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AwsAwsOwners retrieves aws verified owner list
*/
func (a *Client) AwsAwsOwners(params *AwsAwsOwnersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsAwsOwnersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsAwsOwnersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_AwsOwners",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Aws/owners",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsAwsOwnersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsAwsOwnersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_AwsOwners: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsAwsZoneList fetches aws zones
*/
func (a *Client) AwsAwsZoneList(params *AwsAwsZoneListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsAwsZoneListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsAwsZoneListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_AwsZoneList",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Aws/zones",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsAwsZoneListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsAwsZoneListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_AwsZoneList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsCreate adds aws credentials
*/
func (a *Client) AwsCreate(params *AwsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Aws/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsDeviceNames aws device name list
*/
func (a *Client) AwsDeviceNames(params *AwsDeviceNamesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsDeviceNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsDeviceNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_DeviceNames",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Aws/device-names",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsDeviceNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsDeviceNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_DeviceNames: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsList retrieves list of aws cloud credentials
*/
func (a *Client) AwsList(params *AwsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Aws/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsRegionList retrieves aws regions list
*/
func (a *Client) AwsRegionList(params *AwsRegionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsRegionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsRegionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_RegionList",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Aws/regions",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsRegionListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsRegionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_RegionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AwsUpdate updates a w s credentials
*/
func (a *Client) AwsUpdate(params *AwsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AwsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAwsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Aws_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Aws/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AwsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AwsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Aws_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
