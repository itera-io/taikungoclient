// Code generated by go-swagger; DO NOT EDIT.

package user_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	UserGroupsBindProjectGroups(params *UserGroupsBindProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsBindProjectGroupsOK, error)

	UserGroupsCreate(params *UserGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsCreateOK, error)

	UserGroupsDelete(params *UserGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsDeleteOK, *UserGroupsDeleteNoContent, error)

	UserGroupsList(params *UserGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsListOK, error)

	UserGroupsListByProjectGroupID(params *UserGroupsListByProjectGroupIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsListByProjectGroupIDOK, error)

	UserGroupsUpdate(params *UserGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
UserGroupsBindProjectGroups binds project groups
*/
func (a *Client) UserGroupsBindProjectGroups(params *UserGroupsBindProjectGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsBindProjectGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsBindProjectGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_BindProjectGroups",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserGroups/bind-project-groups",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsBindProjectGroupsReader{formats: a.formats},
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
	success, ok := result.(*UserGroupsBindProjectGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserGroups_BindProjectGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserGroupsCreate adds user groups
*/
func (a *Client) UserGroupsCreate(params *UserGroupsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/UserGroups/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsCreateReader{formats: a.formats},
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
	success, ok := result.(*UserGroupsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserGroups_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserGroupsDelete removes user group
*/
func (a *Client) UserGroupsDelete(params *UserGroupsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsDeleteOK, *UserGroupsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/UserGroups/{userGroupId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsDeleteReader{formats: a.formats},
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
	case *UserGroupsDeleteOK:
		return value, nil, nil
	case *UserGroupsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_groups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserGroupsList retrieves list of user groups
*/
func (a *Client) UserGroupsList(params *UserGroupsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/UserGroups/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsListReader{formats: a.formats},
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
	success, ok := result.(*UserGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserGroups_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserGroupsListByProjectGroupID retrieves list of user groups by project group id for dropdown
*/
func (a *Client) UserGroupsListByProjectGroupID(params *UserGroupsListByProjectGroupIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsListByProjectGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsListByProjectGroupIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_ListByProjectGroupId",
		Method:             "GET",
		PathPattern:        "/api/v{v}/UserGroups/list-by-project-group-id",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsListByProjectGroupIDReader{formats: a.formats},
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
	success, ok := result.(*UserGroupsListByProjectGroupIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserGroups_ListByProjectGroupId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UserGroupsUpdate updates user groups
*/
func (a *Client) UserGroupsUpdate(params *UserGroupsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserGroupsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserGroups_Update",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/UserGroups/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsUpdateReader{formats: a.formats},
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
	success, ok := result.(*UserGroupsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UserGroups_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
