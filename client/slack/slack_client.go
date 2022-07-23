// Code generated by go-swagger; DO NOT EDIT.

package slack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new slack API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for slack API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SlackCreate(params *SlackCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackCreateOK, error)

	SlackDeleteMultiple(params *SlackDeleteMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackDeleteMultipleOK, error)

	SlackList(params *SlackListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackListOK, error)

	SlackSlackConfigurationForOrganizationList(params *SlackSlackConfigurationForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackSlackConfigurationForOrganizationListOK, error)

	SlackUpdate(params *SlackUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackUpdateOK, error)

	SlackUpsert(params *SlackUpsertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackUpsertOK, error)

	SlackVerifySlackCredentials(params *SlackVerifySlackCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackVerifySlackCredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SlackCreate creates slack configuration
*/
func (a *Client) SlackCreate(params *SlackCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackCreateReader{formats: a.formats},
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
	success, ok := result.(*SlackCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackDeleteMultiple deletes slack configurations s
*/
func (a *Client) SlackDeleteMultiple(params *SlackDeleteMultipleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackDeleteMultipleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackDeleteMultipleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_DeleteMultiple",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack/delete-multiple",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackDeleteMultipleReader{formats: a.formats},
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
	success, ok := result.(*SlackDeleteMultipleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_DeleteMultiple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackList retrieves a list of slack configs
*/
func (a *Client) SlackList(params *SlackListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Slack",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackListReader{formats: a.formats},
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
	success, ok := result.(*SlackListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackSlackConfigurationForOrganizationList retrieves all slack configs for organization
*/
func (a *Client) SlackSlackConfigurationForOrganizationList(params *SlackSlackConfigurationForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackSlackConfigurationForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackSlackConfigurationForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_SlackConfigurationForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Slack/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackSlackConfigurationForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*SlackSlackConfigurationForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_SlackConfigurationForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackUpdate updates slack configuration
*/
func (a *Client) SlackUpdate(params *SlackUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_Update",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Slack/update/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackUpdateReader{formats: a.formats},
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
	success, ok := result.(*SlackUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackUpsert updates or add slack configuration
*/
func (a *Client) SlackUpsert(params *SlackUpsertParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackUpsertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackUpsertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_Upsert",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackUpsertReader{formats: a.formats},
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
	success, ok := result.(*SlackUpsertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_Upsert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SlackVerifySlackCredentials verifies slack credentials
*/
func (a *Client) SlackVerifySlackCredentials(params *SlackVerifySlackCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackVerifySlackCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackVerifySlackCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_VerifySlackCredentials",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack/verify",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackVerifySlackCredentialsReader{formats: a.formats},
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
	success, ok := result.(*SlackVerifySlackCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Slack_VerifySlackCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
