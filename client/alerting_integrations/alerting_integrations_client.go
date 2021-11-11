// Code generated by go-swagger; DO NOT EDIT.

package alerting_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerting integrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerting integrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AlertingIntegrationsCreate(params *AlertingIntegrationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsCreateOK, error)

	AlertingIntegrationsDelete(params *AlertingIntegrationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsDeleteOK, *AlertingIntegrationsDeleteNoContent, error)

	AlertingIntegrationsEdit(params *AlertingIntegrationsEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsEditOK, error)

	AlertingIntegrationsList(params *AlertingIntegrationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AlertingIntegrationsCreate creates alerting profile alerting integration
*/
func (a *Client) AlertingIntegrationsCreate(params *AlertingIntegrationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingIntegrationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingIntegrations_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/AlertingIntegrations/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingIntegrationsCreateReader{formats: a.formats},
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
	success, ok := result.(*AlertingIntegrationsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingIntegrations_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AlertingIntegrationsDelete deletes alerting profile alerting integration
*/
func (a *Client) AlertingIntegrationsDelete(params *AlertingIntegrationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsDeleteOK, *AlertingIntegrationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingIntegrationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingIntegrations_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/AlertingIntegrations/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingIntegrationsDeleteReader{formats: a.formats},
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
	case *AlertingIntegrationsDeleteOK:
		return value, nil, nil
	case *AlertingIntegrationsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for alerting_integrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AlertingIntegrationsEdit edits alerting profile alerting integration
*/
func (a *Client) AlertingIntegrationsEdit(params *AlertingIntegrationsEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingIntegrationsEditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingIntegrations_Edit",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/AlertingIntegrations/edit",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingIntegrationsEditReader{formats: a.formats},
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
	success, ok := result.(*AlertingIntegrationsEditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingIntegrations_Edit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AlertingIntegrationsList lists alerting integrations by profile id
*/
func (a *Client) AlertingIntegrationsList(params *AlertingIntegrationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AlertingIntegrationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAlertingIntegrationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AlertingIntegrations_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/AlertingIntegrations/{alertingProfileId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AlertingIntegrationsListReader{formats: a.formats},
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
	success, ok := result.(*AlertingIntegrationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AlertingIntegrations_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
