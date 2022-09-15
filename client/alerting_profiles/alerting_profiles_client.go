// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerting profiles API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerting profiles API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AlertingProfilesAlertingProfilesForOrganizationList(params *AlertingProfilesAlertingProfilesForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAlertingProfilesForOrganizationListOK, error)

	AlertingProfilesAssignEmails(params *AlertingProfilesAssignEmailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAssignEmailsOK, error)

	AlertingProfilesAssignWebhooks(params *AlertingProfilesAssignWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAssignWebhooksOK, error)

	AlertingProfilesAttach(params *AlertingProfilesAttachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAttachOK, error)

	AlertingProfilesCreate(params *AlertingProfilesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesCreateOK, error)

	AlertingProfilesDelete(params *AlertingProfilesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesDeleteOK, *AlertingProfilesDeleteNoContent, error)

	AlertingProfilesDetach(params *AlertingProfilesDetachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesDetachOK, error)

	AlertingProfilesEdit(params *AlertingProfilesEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesEditOK, error)

	AlertingProfilesList(params *AlertingProfilesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesListOK, error)

	AlertingProfilesLockManager(params *AlertingProfilesLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesLockManagerOK, error)

	AlertingProfilesVerifyWebhook(params *AlertingProfilesVerifyWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesVerifyWebhookOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AlertingProfilesAlertingProfilesForOrganizationList retrieves all alerting profiles for organization
*/
func (a *Client) AlertingProfilesAlertingProfilesForOrganizationList(params *AlertingProfilesAlertingProfilesForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAlertingProfilesForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesAlertingProfilesForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_AlertingProfilesForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/AlertingProfiles/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesAlertingProfilesForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesAlertingProfilesForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_AlertingProfilesForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesAssignEmails assigns alerting emails
*/
func (a *Client) AlertingProfilesAssignEmails(params *AlertingProfilesAssignEmailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAssignEmailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesAssignEmailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_AssignEmails",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/AlertingProfiles/assignemails/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesAssignEmailsReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesAssignEmailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_AssignEmails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesAssignWebhooks assigns alerting webhooks
*/
func (a *Client) AlertingProfilesAssignWebhooks(params *AlertingProfilesAssignWebhooksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAssignWebhooksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesAssignWebhooksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_AssignWebhooks",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/AlertingProfiles/assignwebhooks/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesAssignWebhooksReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesAssignWebhooksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_AssignWebhooks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesAttach attaches alerting profile to project
*/
func (a *Client) AlertingProfilesAttach(params *AlertingProfilesAttachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesAttachOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesAttachParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_Attach",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/attach",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesAttachReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesAttachOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_Attach: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesCreate adds alerting profiles
*/
func (a *Client) AlertingProfilesCreate(params *AlertingProfilesCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesCreateReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesDelete removes alerting profiles by Id
*/
func (a *Client) AlertingProfilesDelete(params *AlertingProfilesDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesDeleteOK, *AlertingProfilesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesDeleteReader{formats: a.formats},
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
	case *AlertingProfilesDeleteOK:
		return value, nil, nil
	case *AlertingProfilesDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for alerting_profiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesDetach detaches alerting profile from project
*/
func (a *Client) AlertingProfilesDetach(params *AlertingProfilesDetachParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesDetachOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesDetachParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_Detach",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/detach",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesDetachReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesDetachOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_Detach: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesEdit updates alerting profiles
*/
func (a *Client) AlertingProfilesEdit(params *AlertingProfilesEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesEditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_Edit",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/edit",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesEditReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesEditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_Edit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesList retrieves all alerting profiles
*/
func (a *Client) AlertingProfilesList(params *AlertingProfilesListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/AlertingProfiles",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesListReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesLockManager locks unlock alerting profiles
*/
func (a *Client) AlertingProfilesLockManager(params *AlertingProfilesLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_LockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesLockManagerReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_LockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
AlertingProfilesVerifyWebhook verifies webhook endpoint
*/
func (a *Client) AlertingProfilesVerifyWebhook(params *AlertingProfilesVerifyWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingProfilesVerifyWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingProfilesVerifyWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingProfiles_VerifyWebhook",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingProfiles/verifywebhook",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingProfilesVerifyWebhookReader{formats: a.formats},
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
	success, ok := result.(*AlertingProfilesVerifyWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingProfiles_VerifyWebhook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
