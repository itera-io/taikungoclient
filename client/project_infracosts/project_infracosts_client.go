// Code generated by go-swagger; DO NOT EDIT.

package project_infracosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new project infracosts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project infracosts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectInfracostsDelete(params *ProjectInfracostsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectInfracostsDeleteOK, *ProjectInfracostsDeleteNoContent, error)

	ProjectInfracostsEdit(params *ProjectInfracostsEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectInfracostsEditOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ProjectInfracostsDelete deletes the project infracost
*/
func (a *Client) ProjectInfracostsDelete(params *ProjectInfracostsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectInfracostsDeleteOK, *ProjectInfracostsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectInfracostsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectInfracosts_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/ProjectInfracosts/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectInfracostsDeleteReader{formats: a.formats},
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
	case *ProjectInfracostsDeleteOK:
		return value, nil, nil
	case *ProjectInfracostsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for project_infracosts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ProjectInfracostsEdit upserts project infracost by project Id
*/
func (a *Client) ProjectInfracostsEdit(params *ProjectInfracostsEditParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectInfracostsEditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectInfracostsEditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectInfracosts_Edit",
		Method:             "POST",
		PathPattern:        "/api/v{v}/ProjectInfracosts/upsert/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectInfracostsEditReader{formats: a.formats},
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
	success, ok := result.(*ProjectInfracostsEditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectInfracosts_Edit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
