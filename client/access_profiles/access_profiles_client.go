// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new access profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for access profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AccessProfilesAccessProfilesForOrganizationList(params *AccessProfilesAccessProfilesForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesAccessProfilesForOrganizationListOK, error)

	AccessProfilesCreate(params *AccessProfilesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesCreateOK, error)

	AccessProfilesDelete(params *AccessProfilesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesDeleteOK, *AccessProfilesDeleteNoContent, error)

	AccessProfilesList(params *AccessProfilesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesListOK, error)

	AccessProfilesLockManager(params *AccessProfilesLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesLockManagerOK, error)

	AccessProfilesUpdate(params *AccessProfilesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AccessProfilesAccessProfilesForOrganizationList retrieves access profiles by organization Id
*/
func (a *Client) AccessProfilesAccessProfilesForOrganizationList(params *AccessProfilesAccessProfilesForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesAccessProfilesForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesAccessProfilesForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_AccessProfilesForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/AccessProfiles/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesAccessProfilesForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*AccessProfilesAccessProfilesForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AccessProfiles_AccessProfilesForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AccessProfilesCreate creates access profile
*/
func (a *Client) AccessProfilesCreate(params *AccessProfilesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AccessProfiles/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesCreateReader{formats: a.formats},
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
	success, ok := result.(*AccessProfilesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AccessProfiles_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AccessProfilesDelete deletes access profile by Id
*/
func (a *Client) AccessProfilesDelete(params *AccessProfilesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesDeleteOK, *AccessProfilesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/AccessProfiles/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesDeleteReader{formats: a.formats},
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
	case *AccessProfilesDeleteOK:
		return value, nil, nil
	case *AccessProfilesDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for access_profiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AccessProfilesList retrieves all access profiles
*/
func (a *Client) AccessProfilesList(params *AccessProfilesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/AccessProfiles",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesListReader{formats: a.formats},
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
	success, ok := result.(*AccessProfilesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AccessProfiles_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AccessProfilesLockManager locks unlock access profiles
*/
func (a *Client) AccessProfilesLockManager(params *AccessProfilesLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_LockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AccessProfiles/lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesLockManagerReader{formats: a.formats},
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
	success, ok := result.(*AccessProfilesLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AccessProfiles_LockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AccessProfilesUpdate updates access profile
*/
func (a *Client) AccessProfilesUpdate(params *AccessProfilesUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AccessProfilesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAccessProfilesUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AccessProfiles_Update",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/AccessProfiles/update/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AccessProfilesUpdateReader{formats: a.formats},
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
	success, ok := result.(*AccessProfilesUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AccessProfiles_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
