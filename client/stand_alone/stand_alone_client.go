// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stand alone API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stand alone API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	StandAloneCommit(params *StandAloneCommitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneCommitOK, error)

	StandAloneCreate(params *StandAloneCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneCreateOK, error)

	StandAloneDelete(params *StandAloneDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneDeleteOK, error)

	StandAloneIPManagement(params *StandAloneIPManagementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneIPManagementOK, error)

	StandAloneProjectDetails(params *StandAloneProjectDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneProjectDetailsOK, error)

	StandAlonePurge(params *StandAlonePurgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAlonePurgeOK, error)

	StandAloneReset(params *StandAloneResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneResetOK, error)

	StandAloneUpdate(params *StandAloneUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StandAloneCommit commits stand alone vm
*/
func (a *Client) StandAloneCommit(params *StandAloneCommitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneCommitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneCommitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Commit",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/commit",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneCommitReader{formats: a.formats},
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
	success, ok := result.(*StandAloneCommitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Commit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneCreate creates stand alone vm
*/
func (a *Client) StandAloneCreate(params *StandAloneCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneCreateReader{formats: a.formats},
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
	success, ok := result.(*StandAloneCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneDelete deletes purge stand alone vm
*/
func (a *Client) StandAloneDelete(params *StandAloneDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneDeleteReader{formats: a.formats},
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
	success, ok := result.(*StandAloneDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneIPManagement enables disable stand alone public ip
*/
func (a *Client) StandAloneIPManagement(params *StandAloneIPManagementParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneIPManagementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneIPManagementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_IpManagement",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/ip/management",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneIPManagementReader{formats: a.formats},
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
	success, ok := result.(*StandAloneIPManagementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_IpManagement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneProjectDetails retrieves details of the project by Id
*/
func (a *Client) StandAloneProjectDetails(params *StandAloneProjectDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneProjectDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneProjectDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_ProjectDetails",
		Method:             "GET",
		PathPattern:        "/api/v{v}/StandAlone/project/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneProjectDetailsReader{formats: a.formats},
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
	success, ok := result.(*StandAloneProjectDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_ProjectDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAlonePurge purges stand alone based project
*/
func (a *Client) StandAlonePurge(params *StandAlonePurgeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAlonePurgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAlonePurgeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Purge",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/purge",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAlonePurgeReader{formats: a.formats},
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
	success, ok := result.(*StandAlonePurgeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Purge: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneReset resets stand alone vm status
*/
func (a *Client) StandAloneReset(params *StandAloneResetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneResetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Reset",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/reset",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneResetReader{formats: a.formats},
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
	success, ok := result.(*StandAloneResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Reset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StandAloneUpdate updates stand alone vm
*/
func (a *Client) StandAloneUpdate(params *StandAloneUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StandAloneUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStandAloneUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StandAlone_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/StandAlone/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StandAloneUpdateReader{formats: a.formats},
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
	success, ok := result.(*StandAloneUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StandAlone_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
