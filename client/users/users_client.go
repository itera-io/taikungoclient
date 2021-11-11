// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UsersChangePassword(params *UsersChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersChangePasswordOK, error)

	UsersCreate(params *UsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersCreateOK, error)

	UsersDelete(params *UsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersDeleteOK, *UsersDeleteNoContent, error)

	UsersList(params *UsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersListOK, error)

	UsersUpdateUser(params *UsersUpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUpdateUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UsersChangePassword changes user password
*/
func (a *Client) UsersChangePassword(params *UsersChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersChangePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersChangePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Users_ChangePassword",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Users/changepassword",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersChangePasswordReader{formats: a.formats},
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
	success, ok := result.(*UsersChangePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Users_ChangePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersCreate creates a new user only valid for admin and moderators moderators don t need to fill the organization Id parameter for admin also by default his organization Id
*/
func (a *Client) UsersCreate(params *UsersCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Users_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Users",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersCreateReader{formats: a.formats},
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
	success, ok := result.(*UsersCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Users_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersDelete deletes user only valid for admin moderators and partner reminder moderators can delete users from their organization only
*/
func (a *Client) UsersDelete(params *UsersDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersDeleteOK, *UsersDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Users_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/Users/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UsersDeleteOK:
		return value, nil, nil
	case *UsersDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for users: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersList retrieves all users only valid for admin
*/
func (a *Client) UsersList(params *UsersListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Users_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Users",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersListReader{formats: a.formats},
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
	success, ok := result.(*UsersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Users_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UsersUpdateUser updates user credential
*/
func (a *Client) UsersUpdateUser(params *UsersUpdateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UsersUpdateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUsersUpdateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Users_UpdateUser",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Users/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UsersUpdateUserReader{formats: a.formats},
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
	success, ok := result.(*UsersUpdateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Users_UpdateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
