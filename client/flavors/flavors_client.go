// Code generated by go-swagger; DO NOT EDIT.

package flavors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new flavors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for flavors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	FlavorsAwsFlavors(params *FlavorsAwsFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsAwsFlavorsOK, error)

	FlavorsAzureFlavors(params *FlavorsAzureFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsAzureFlavorsOK, error)

	FlavorsBindToProject(params *FlavorsBindToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsBindToProjectOK, error)

	FlavorsDropdownRecordDtos(params *FlavorsDropdownRecordDtosParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsDropdownRecordDtosOK, error)

	FlavorsGetSelectedFlavorsForProject(params *FlavorsGetSelectedFlavorsForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsGetSelectedFlavorsForProjectOK, error)

	FlavorsGoogleFlavors(params *FlavorsGoogleFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsGoogleFlavorsOK, error)

	FlavorsOpenstackFlavors(params *FlavorsOpenstackFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsOpenstackFlavorsOK, error)

	FlavorsUnbindFromProject(params *FlavorsUnbindFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsUnbindFromProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
FlavorsAwsFlavors retrieves aws flavors
*/
func (a *Client) FlavorsAwsFlavors(params *FlavorsAwsFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsAwsFlavorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsAwsFlavorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_AwsFlavors",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/aws/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsAwsFlavorsReader{formats: a.formats},
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
	success, ok := result.(*FlavorsAwsFlavorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_AwsFlavors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsAzureFlavors retrieves azure flavors
*/
func (a *Client) FlavorsAzureFlavors(params *FlavorsAzureFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsAzureFlavorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsAzureFlavorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_AzureFlavors",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/azure/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsAzureFlavorsReader{formats: a.formats},
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
	success, ok := result.(*FlavorsAzureFlavorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_AzureFlavors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsBindToProject binds flavors to project
*/
func (a *Client) FlavorsBindToProject(params *FlavorsBindToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsBindToProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsBindToProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_BindToProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Flavors/bind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsBindToProjectReader{formats: a.formats},
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
	success, ok := result.(*FlavorsBindToProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_BindToProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsDropdownRecordDtos retrieves cloud credentials dropdown list
*/
func (a *Client) FlavorsDropdownRecordDtos(params *FlavorsDropdownRecordDtosParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsDropdownRecordDtosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsDropdownRecordDtosParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_DropdownRecordDtos",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/credentials/dropdown/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsDropdownRecordDtosReader{formats: a.formats},
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
	success, ok := result.(*FlavorsDropdownRecordDtosOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_DropdownRecordDtos: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsGetSelectedFlavorsForProject retrieves selected flavors for projects
*/
func (a *Client) FlavorsGetSelectedFlavorsForProject(params *FlavorsGetSelectedFlavorsForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsGetSelectedFlavorsForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsGetSelectedFlavorsForProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_GetSelectedFlavorsForProject",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/projects/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsGetSelectedFlavorsForProjectReader{formats: a.formats},
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
	success, ok := result.(*FlavorsGetSelectedFlavorsForProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_GetSelectedFlavorsForProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsGoogleFlavors retrieves google flavors
*/
func (a *Client) FlavorsGoogleFlavors(params *FlavorsGoogleFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsGoogleFlavorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsGoogleFlavorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_GoogleFlavors",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/google/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsGoogleFlavorsReader{formats: a.formats},
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
	success, ok := result.(*FlavorsGoogleFlavorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_GoogleFlavors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsOpenstackFlavors retrieves openstack flavors
*/
func (a *Client) FlavorsOpenstackFlavors(params *FlavorsOpenstackFlavorsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsOpenstackFlavorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsOpenstackFlavorsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_OpenstackFlavors",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Flavors/openstack/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsOpenstackFlavorsReader{formats: a.formats},
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
	success, ok := result.(*FlavorsOpenstackFlavorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_OpenstackFlavors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
FlavorsUnbindFromProject unbinds flavors from project
*/
func (a *Client) FlavorsUnbindFromProject(params *FlavorsUnbindFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FlavorsUnbindFromProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFlavorsUnbindFromProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Flavors_UnbindFromProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Flavors/unbind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FlavorsUnbindFromProjectReader{formats: a.formats},
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
	success, ok := result.(*FlavorsUnbindFromProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Flavors_UnbindFromProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
