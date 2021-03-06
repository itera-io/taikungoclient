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

	SlackDelete(params *SlackDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackDeleteOK, *SlackDeleteNoContent, error)

	SlackList(params *SlackListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SlackCreate updates or add slack configuration
*/
func (a *Client) SlackCreate(params *SlackCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack",
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
  SlackDelete deletes slack config file
*/
func (a *Client) SlackDelete(params *SlackDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SlackDeleteOK, *SlackDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSlackDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Slack_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Slack/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SlackDeleteReader{formats: a.formats},
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
	case *SlackDeleteOK:
		return value, nil, nil
	case *SlackDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for slack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
