// Code generated by go-swagger; DO NOT EDIT.

package pre_defined_queries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pre defined queries API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pre defined queries API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PreDefinedQueriesCreatePrometheusDashboard(params *PreDefinedQueriesCreatePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesCreatePrometheusDashboardOK, error)

	PreDefinedQueriesDeletePrometheusDashboard(params *PreDefinedQueriesDeletePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesDeletePrometheusDashboardOK, error)

	PreDefinedQueriesGetPrometheusCommonDashboardList(params *PreDefinedQueriesGetPrometheusCommonDashboardListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesGetPrometheusCommonDashboardListOK, error)

	PreDefinedQueriesGetPrometheusDashboardList(params *PreDefinedQueriesGetPrometheusDashboardListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesGetPrometheusDashboardListOK, error)

	PreDefinedQueriesUpdatePrometheusDashboard(params *PreDefinedQueriesUpdatePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesUpdatePrometheusDashboardOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PreDefinedQueriesCreatePrometheusDashboard creates prometheus dashboard pre defined query
*/
func (a *Client) PreDefinedQueriesCreatePrometheusDashboard(params *PreDefinedQueriesCreatePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesCreatePrometheusDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreDefinedQueriesCreatePrometheusDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreDefinedQueries_CreatePrometheusDashboard",
		Method:             "POST",
		PathPattern:        "/api/v{v}/PreDefinedQueries/prometheus/dashboard/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreDefinedQueriesCreatePrometheusDashboardReader{formats: a.formats},
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
	success, ok := result.(*PreDefinedQueriesCreatePrometheusDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PreDefinedQueries_CreatePrometheusDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PreDefinedQueriesDeletePrometheusDashboard deletes prometheus dashboard pre defined query
*/
func (a *Client) PreDefinedQueriesDeletePrometheusDashboard(params *PreDefinedQueriesDeletePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesDeletePrometheusDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreDefinedQueriesDeletePrometheusDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreDefinedQueries_DeletePrometheusDashboard",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/PreDefinedQueries/prometheus/dashboard/delete/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreDefinedQueriesDeletePrometheusDashboardReader{formats: a.formats},
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
	success, ok := result.(*PreDefinedQueriesDeletePrometheusDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PreDefinedQueries_DeletePrometheusDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PreDefinedQueriesGetPrometheusCommonDashboardList gets list of pre defined common prometheus dashboard elements
*/
func (a *Client) PreDefinedQueriesGetPrometheusCommonDashboardList(params *PreDefinedQueriesGetPrometheusCommonDashboardListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesGetPrometheusCommonDashboardListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreDefinedQueriesGetPrometheusCommonDashboardListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreDefinedQueries_GetPrometheusCommonDashboardList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/PreDefinedQueries/prometheus/dashboard/common",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreDefinedQueriesGetPrometheusCommonDashboardListReader{formats: a.formats},
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
	success, ok := result.(*PreDefinedQueriesGetPrometheusCommonDashboardListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PreDefinedQueries_GetPrometheusCommonDashboardList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PreDefinedQueriesGetPrometheusDashboardList gets list of pre defined organization prometheus dashboard elements
*/
func (a *Client) PreDefinedQueriesGetPrometheusDashboardList(params *PreDefinedQueriesGetPrometheusDashboardListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesGetPrometheusDashboardListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreDefinedQueriesGetPrometheusDashboardListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreDefinedQueries_GetPrometheusDashboardList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/PreDefinedQueries/prometheus/dashboard/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreDefinedQueriesGetPrometheusDashboardListReader{formats: a.formats},
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
	success, ok := result.(*PreDefinedQueriesGetPrometheusDashboardListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PreDefinedQueries_GetPrometheusDashboardList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PreDefinedQueriesUpdatePrometheusDashboard updates prometheus dashboard pre defined query
*/
func (a *Client) PreDefinedQueriesUpdatePrometheusDashboard(params *PreDefinedQueriesUpdatePrometheusDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PreDefinedQueriesUpdatePrometheusDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPreDefinedQueriesUpdatePrometheusDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PreDefinedQueries_UpdatePrometheusDashboard",
		Method:             "POST",
		PathPattern:        "/api/v{v}/PreDefinedQueries/prometheus/dashboard/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PreDefinedQueriesUpdatePrometheusDashboardReader{formats: a.formats},
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
	success, ok := result.(*PreDefinedQueriesUpdatePrometheusDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PreDefinedQueries_UpdatePrometheusDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
