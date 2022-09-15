// Code generated by go-swagger; DO NOT EDIT.

package google_cloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new google cloud API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for google cloud API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GoogleCloudBillingAccountList(params *GoogleCloudBillingAccountListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudBillingAccountListOK, error)

	GoogleCloudCreate(params *GoogleCloudCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudCreateOK, error)

	GoogleCloudList(params *GoogleCloudListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudListOK, error)

	GoogleCloudRegionList(params *GoogleCloudRegionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudRegionListOK, error)

	GoogleCloudZoneList(params *GoogleCloudZoneListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudZoneListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GoogleCloudBillingAccountList retrieves google billing accounts list
*/
func (a *Client) GoogleCloudBillingAccountList(params *GoogleCloudBillingAccountListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudBillingAccountListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoogleCloudBillingAccountListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GoogleCloud_BillingAccountList",
		Method:             "POST",
		PathPattern:        "/api/v{v}/GoogleCloud/billing-accounts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoogleCloudBillingAccountListReader{formats: a.formats},
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
	success, ok := result.(*GoogleCloudBillingAccountListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GoogleCloud_BillingAccountList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GoogleCloudCreate creates google cloud credential
*/
func (a *Client) GoogleCloudCreate(params *GoogleCloudCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoogleCloudCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GoogleCloud_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/GoogleCloud/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoogleCloudCreateReader{formats: a.formats},
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
	success, ok := result.(*GoogleCloudCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GoogleCloud_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GoogleCloudList retrieves list of google cloud credentials
*/
func (a *Client) GoogleCloudList(params *GoogleCloudListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoogleCloudListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GoogleCloud_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/GoogleCloud/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoogleCloudListReader{formats: a.formats},
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
	success, ok := result.(*GoogleCloudListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GoogleCloud_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GoogleCloudRegionList retrieves google region list
*/
func (a *Client) GoogleCloudRegionList(params *GoogleCloudRegionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudRegionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoogleCloudRegionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GoogleCloud_RegionList",
		Method:             "POST",
		PathPattern:        "/api/v{v}/GoogleCloud/regions",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoogleCloudRegionListReader{formats: a.formats},
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
	success, ok := result.(*GoogleCloudRegionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GoogleCloud_RegionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GoogleCloudZoneList retrieves google zone list
*/
func (a *Client) GoogleCloudZoneList(params *GoogleCloudZoneListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GoogleCloudZoneListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGoogleCloudZoneListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GoogleCloud_ZoneList",
		Method:             "POST",
		PathPattern:        "/api/v{v}/GoogleCloud/zones",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GoogleCloudZoneListReader{formats: a.formats},
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
	success, ok := result.(*GoogleCloudZoneListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GoogleCloud_ZoneList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
