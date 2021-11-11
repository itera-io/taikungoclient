// Code generated by go-swagger; DO NOT EDIT.

package keycloak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new keycloak API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for keycloak API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	KeycloakCreate(params *KeycloakCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakCreateOK, error)

	KeycloakDelete(params *KeycloakDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakDeleteOK, error)

	KeycloakEdit(params *KeycloakEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakEditOK, error)

	KeycloakGet(params *KeycloakGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakGetOK, error)

	KeycloakLogin(params *KeycloakLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakLoginOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  KeycloakCreate creates keycloak configuration for organization
*/
func (a *Client) KeycloakCreate(params *KeycloakCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeycloakCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Keycloak_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Keycloak/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KeycloakCreateReader{formats: a.formats},
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
	success, ok := result.(*KeycloakCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Keycloak_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeycloakDelete deletes keycloak configuration
*/
func (a *Client) KeycloakDelete(params *KeycloakDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeycloakDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Keycloak_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Keycloak/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KeycloakDeleteReader{formats: a.formats},
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
	success, ok := result.(*KeycloakDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Keycloak_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeycloakEdit edits keycloak configuration for organization
*/
func (a *Client) KeycloakEdit(params *KeycloakEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeycloakEditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Keycloak_Edit",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Keycloak/edit",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KeycloakEditReader{formats: a.formats},
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
	success, ok := result.(*KeycloakEditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Keycloak_Edit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeycloakGet gets keycloak configuration
*/
func (a *Client) KeycloakGet(params *KeycloakGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeycloakGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Keycloak_Get",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Keycloak",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KeycloakGetReader{formats: a.formats},
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
	success, ok := result.(*KeycloakGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Keycloak_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KeycloakLogin logins to API by using keycloak email and password
*/
func (a *Client) KeycloakLogin(params *KeycloakLoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KeycloakLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKeycloakLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Keycloak_Login",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Keycloak/login",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KeycloakLoginReader{formats: a.formats},
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
	success, ok := result.(*KeycloakLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Keycloak_Login: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
