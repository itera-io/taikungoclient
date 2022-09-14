// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new common API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for common API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CommonGetCountryList(params *CommonGetCountryListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetCountryListOK, error)

	CommonGetEnumValues(params *CommonGetEnumValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetEnumValuesOK, error)

	CommonGetSortingElements(params *CommonGetSortingElementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetSortingElementsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CommonGetCountryList retrieves country list
*/
func (a *Client) CommonGetCountryList(params *CommonGetCountryListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetCountryListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommonGetCountryListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Common_GetCountryList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Common/countries",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommonGetCountryListReader{formats: a.formats},
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
	success, ok := result.(*CommonGetCountryListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Common_GetCountryList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CommonGetEnumValues retrieves enum values
*/
func (a *Client) CommonGetEnumValues(params *CommonGetEnumValuesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetEnumValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommonGetEnumValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Common_GetEnumValues",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Common/enumvalues",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommonGetEnumValuesReader{formats: a.formats},
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
	success, ok := result.(*CommonGetEnumValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Common_GetEnumValues: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CommonGetSortingElements retrieves country list
*/
func (a *Client) CommonGetSortingElements(params *CommonGetSortingElementsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CommonGetSortingElementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommonGetSortingElementsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Common_GetSortingElements",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Common/sorting-elements/{type}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommonGetSortingElementsReader{formats: a.formats},
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
	success, ok := result.(*CommonGetSortingElementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Common_GetSortingElements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
