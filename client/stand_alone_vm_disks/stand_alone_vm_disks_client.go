// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_vm_disks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stand alone vm disks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stand alone vm disks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StandAloneVMDisksCreate(params *StandAloneVMDisksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksCreateOK, error)

	StandAloneVMDisksDelete(params *StandAloneVMDisksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksDeleteOK, error)

	StandAloneVMDisksReset(params *StandAloneVMDisksResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksResetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StandAloneVMDisksCreate adds disk for stand alone vm
*/
func (a *Client) StandAloneVMDisksCreate(params *StandAloneVMDisksCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneVMDisksCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAloneVmDisks_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAloneVmDisks/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneVMDisksCreateReader{formats: a.formats},
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
	success, ok := result.(*StandAloneVMDisksCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAloneVmDisks_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneVMDisksDelete deletes disk by id
*/
func (a *Client) StandAloneVMDisksDelete(params *StandAloneVMDisksDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneVMDisksDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAloneVmDisks_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAloneVmDisks/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneVMDisksDeleteReader{formats: a.formats},
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
	success, ok := result.(*StandAloneVMDisksDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAloneVmDisks_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneVMDisksReset updates statuses of disks by vm Id
*/
func (a *Client) StandAloneVMDisksReset(params *StandAloneVMDisksResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneVMDisksResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneVMDisksResetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAloneVmDisks_Reset",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAloneVmDisks/reset",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneVMDisksResetReader{formats: a.formats},
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
	success, ok := result.(*StandAloneVMDisksResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAloneVmDisks_Reset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
