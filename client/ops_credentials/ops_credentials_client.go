// Code generated by go-swagger; DO NOT EDIT.

package ops_credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ops credentials API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ops credentials API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	OpsCredentialsBackupCredentialsForOrganizationList(params *OpsCredentialsBackupCredentialsForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsBackupCredentialsForOrganizationListOK, error)

	OpsCredentialsCreate(params *OpsCredentialsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsCreateOK, error)

	OpsCredentialsDelete(params *OpsCredentialsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsDeleteOK, *OpsCredentialsDeleteNoContent, error)

	OpsCredentialsList(params *OpsCredentialsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsListOK, error)

	OpsCredentialsLockManager(params *OpsCredentialsLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsLockManagerOK, error)

	OpsCredentialsMakeDefault(params *OpsCredentialsMakeDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsMakeDefaultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  OpsCredentialsBackupCredentialsForOrganizationList retrieves backup credentials by organization Id
*/
func (a *Client) OpsCredentialsBackupCredentialsForOrganizationList(params *OpsCredentialsBackupCredentialsForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsBackupCredentialsForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsBackupCredentialsForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_BackupCredentialsForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/OpsCredentials",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsBackupCredentialsForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*OpsCredentialsBackupCredentialsForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpsCredentials_BackupCredentialsForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OpsCredentialsCreate adds operation credential
*/
func (a *Client) OpsCredentialsCreate(params *OpsCredentialsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/OpsCredentials",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsCreateReader{formats: a.formats},
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
	success, ok := result.(*OpsCredentialsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpsCredentials_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OpsCredentialsDelete removes operation credential
*/
func (a *Client) OpsCredentialsDelete(params *OpsCredentialsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsDeleteOK, *OpsCredentialsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/OpsCredentials/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsDeleteReader{formats: a.formats},
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
	case *OpsCredentialsDeleteOK:
		return value, nil, nil
	case *OpsCredentialsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ops_credentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OpsCredentialsList retrieves all operation credentials
*/
func (a *Client) OpsCredentialsList(params *OpsCredentialsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/OpsCredentials/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsListReader{formats: a.formats},
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
	success, ok := result.(*OpsCredentialsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpsCredentials_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OpsCredentialsLockManager locks unlock operation credential
*/
func (a *Client) OpsCredentialsLockManager(params *OpsCredentialsLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_LockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/OpsCredentials/lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsLockManagerReader{formats: a.formats},
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
	success, ok := result.(*OpsCredentialsLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpsCredentials_LockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OpsCredentialsMakeDefault makes ops credentials default
*/
func (a *Client) OpsCredentialsMakeDefault(params *OpsCredentialsMakeDefaultParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OpsCredentialsMakeDefaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpsCredentialsMakeDefaultParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "OpsCredentials_MakeDefault",
		Method:             "POST",
		PathPattern:        "/api/v{v}/OpsCredentials/makedefault",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OpsCredentialsMakeDefaultReader{formats: a.formats},
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
	success, ok := result.(*OpsCredentialsMakeDefaultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for OpsCredentials_MakeDefault: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
