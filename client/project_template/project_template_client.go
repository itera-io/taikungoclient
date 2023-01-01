// Code generated by go-swagger; DO NOT EDIT.

package project_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new project template API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project template API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectTemplateCreate(params *ProjectTemplateCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateCreateOK, error)

	ProjectTemplateDelete(params *ProjectTemplateDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateDeleteOK, *ProjectTemplateDeleteNoContent, error)

	ProjectTemplateList(params *ProjectTemplateListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateListOK, error)

	ProjectTemplateProjectTemplateForOrganizationList(params *ProjectTemplateProjectTemplateForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateProjectTemplateForOrganizationListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ProjectTemplateCreate creates project from template
*/
func (a *Client) ProjectTemplateCreate(params *ProjectTemplateCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectTemplateCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectTemplate_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/ProjectTemplate/create-project",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectTemplateCreateReader{formats: a.formats},
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
	success, ok := result.(*ProjectTemplateCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectTemplate_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ProjectTemplateDelete deletes project template by Id
*/
func (a *Client) ProjectTemplateDelete(params *ProjectTemplateDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateDeleteOK, *ProjectTemplateDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectTemplateDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectTemplate_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/ProjectTemplate/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectTemplateDeleteReader{formats: a.formats},
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
	case *ProjectTemplateDeleteOK:
		return value, nil, nil
	case *ProjectTemplateDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for project_template: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ProjectTemplateList retrieves all project templates
*/
func (a *Client) ProjectTemplateList(params *ProjectTemplateListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectTemplateListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectTemplate_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/ProjectTemplate",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectTemplateListReader{formats: a.formats},
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
	success, ok := result.(*ProjectTemplateListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectTemplate_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ProjectTemplateProjectTemplateForOrganizationList retrieves project template by organization Id
*/
func (a *Client) ProjectTemplateProjectTemplateForOrganizationList(params *ProjectTemplateProjectTemplateForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectTemplateProjectTemplateForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectTemplateProjectTemplateForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ProjectTemplate_ProjectTemplateForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/ProjectTemplate/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectTemplateProjectTemplateForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*ProjectTemplateProjectTemplateForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ProjectTemplate_ProjectTemplateForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
