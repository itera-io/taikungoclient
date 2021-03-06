// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BillingExportCsv(params *BillingExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingExportCsvOK, error)

	BillingGroupedList(params *BillingGroupedListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingGroupedListOK, error)

	BillingList(params *BillingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BillingExportCsv exports csv
*/
func (a *Client) BillingExportCsv(params *BillingExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingExportCsvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingExportCsvParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Billing_ExportCsv",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Billing/export",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BillingExportCsvReader{formats: a.formats},
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
	success, ok := result.(*BillingExportCsvOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Billing_ExportCsv: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BillingGroupedList retrieves a grouped list of billing summaries
*/
func (a *Client) BillingGroupedList(params *BillingGroupedListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingGroupedListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingGroupedListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Billing_GroupedList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Billing/grouped",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BillingGroupedListReader{formats: a.formats},
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
	success, ok := result.(*BillingGroupedListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Billing_GroupedList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  BillingList retrieves billing info
*/
func (a *Client) BillingList(params *BillingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BillingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBillingListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Billing_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Billing",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BillingListReader{formats: a.formats},
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
	success, ok := result.(*BillingListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Billing_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
