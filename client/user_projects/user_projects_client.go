// Code generated by go-swagger; DO NOT EDIT.

package user_projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user projects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user projects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserProjectsBindProjects(params *UserProjectsBindProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsBindProjectsOK, error)

	UserProjectsBindUsers(params *UserProjectsBindUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsBindUsersOK, error)

	UserProjectsUsersListByProject(params *UserProjectsUsersListByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsUsersListByProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserProjectsBindProjects binds projects to users
*/
func (a *Client) UserProjectsBindProjects(params *UserProjectsBindProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsBindProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserProjectsBindProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserProjects_BindProjects",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserProjects/bindprojects",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserProjectsBindProjectsReader{formats: a.formats},
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
	success, ok := result.(*UserProjectsBindProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserProjects_BindProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserProjectsBindUsers binds users to projects
*/
func (a *Client) UserProjectsBindUsers(params *UserProjectsBindUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsBindUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserProjectsBindUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserProjects_BindUsers",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserProjects/bindusers",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserProjectsBindUsersReader{formats: a.formats},
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
	success, ok := result.(*UserProjectsBindUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserProjects_BindUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserProjectsUsersListByProject users list by project id
*/
func (a *Client) UserProjectsUsersListByProject(params *UserProjectsUsersListByProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserProjectsUsersListByProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserProjectsUsersListByProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserProjects_UsersListByProject",
		Method:             "GET",
		PathPattern:        "/api/v{v}/UserProjects/users/list/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserProjectsUsersListByProjectReader{formats: a.formats},
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
	success, ok := result.(*UserProjectsUsersListByProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserProjects_UsersListByProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
