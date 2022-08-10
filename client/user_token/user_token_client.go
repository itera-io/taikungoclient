// Code generated by go-swagger; DO NOT EDIT.

package user_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserTokenAvailableEndpointList(params *UserTokenAvailableEndpointListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenAvailableEndpointListOK, error)

	UserTokenBindUnbind(params *UserTokenBindUnbindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenBindUnbindOK, error)

	UserTokenCreate(params *UserTokenCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenCreateOK, error)

	UserTokenDelete(params *UserTokenDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenDeleteOK, error)

	UserTokenList(params *UserTokenListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserTokenAvailableEndpointList gets available endpoints list
*/
func (a *Client) UserTokenAvailableEndpointList(params *UserTokenAvailableEndpointListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenAvailableEndpointListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTokenAvailableEndpointListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserToken_AvailableEndpointList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/UserToken/available-endpoints",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserTokenAvailableEndpointListReader{formats: a.formats},
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
	success, ok := result.(*UserTokenAvailableEndpointListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserToken_AvailableEndpointList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserTokenBindUnbind binds and unbind endpoints
*/
func (a *Client) UserTokenBindUnbind(params *UserTokenBindUnbindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenBindUnbindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTokenBindUnbindParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserToken_BindUnbind",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserToken/bind-unbind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserTokenBindUnbindReader{formats: a.formats},
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
	success, ok := result.(*UserTokenBindUnbindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserToken_BindUnbind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserTokenCreate creates a new user token
*/
func (a *Client) UserTokenCreate(params *UserTokenCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTokenCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserToken_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserToken/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserTokenCreateReader{formats: a.formats},
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
	success, ok := result.(*UserTokenCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserToken_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserTokenDelete deletes token
*/
func (a *Client) UserTokenDelete(params *UserTokenDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTokenDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserToken_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/UserToken/delete/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserTokenDeleteReader{formats: a.formats},
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
	success, ok := result.(*UserTokenDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserToken_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserTokenList gets user token list
*/
func (a *Client) UserTokenList(params *UserTokenListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserTokenListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserTokenListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserToken_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/UserToken/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserTokenListReader{formats: a.formats},
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
	success, ok := result.(*UserTokenListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserToken_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}