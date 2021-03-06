// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new azure API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for azure API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AzureCreate(params *AzureCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureCreateOK, error)

	AzureDashboard(params *AzureDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureDashboardOK, error)

	AzureLocations(params *AzureLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureLocationsOK, error)

	AzureOffers(params *AzureOffersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureOffersOK, error)

	AzurePublishers(params *AzurePublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzurePublishersOK, error)

	AzureSkus(params *AzureSkusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureSkusOK, error)

	AzureSubscriptions(params *AzureSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureSubscriptionsOK, error)

	AzureUpdate(params *AzureUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureUpdateOK, error)

	AzureZones(params *AzureZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureZonesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AzureCreate adds azure credentials
*/
func (a *Client) AzureCreate(params *AzureCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureCreateReader{formats: a.formats},
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
	success, ok := result.(*AzureCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureDashboard fetches azure quota list
*/
func (a *Client) AzureDashboard(params *AzureDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Dashboard",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/quota/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureDashboardReader{formats: a.formats},
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
	success, ok := result.(*AzureDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Dashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureLocations fetches azure location list
*/
func (a *Client) AzureLocations(params *AzureLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Locations",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/locations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureLocationsReader{formats: a.formats},
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
	success, ok := result.(*AzureLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Locations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureOffers lists azure offer list by publisher
*/
func (a *Client) AzureOffers(params *AzureOffersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureOffersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureOffersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Offers",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Azure/offers/{cloudId}/{publisher}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureOffersReader{formats: a.formats},
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
	success, ok := result.(*AzureOffersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Offers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzurePublishers lists azure publishers list
*/
func (a *Client) AzurePublishers(params *AzurePublishersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzurePublishersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzurePublishersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Publishers",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Azure/publishers/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzurePublishersReader{formats: a.formats},
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
	success, ok := result.(*AzurePublishersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Publishers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureSkus lists azure skus list by publisher and offer
*/
func (a *Client) AzureSkus(params *AzureSkusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureSkusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureSkusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Skus",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Azure/skus/{cloudId}/{publisher}/{offer}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureSkusReader{formats: a.formats},
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
	success, ok := result.(*AzureSkusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Skus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureSubscriptions azures subscriptions list
*/
func (a *Client) AzureSubscriptions(params *AzureSubscriptionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureSubscriptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Subscriptions",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/subscriptions",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureSubscriptionsReader{formats: a.formats},
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
	success, ok := result.(*AzureSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Subscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureUpdate updates azure credentials
*/
func (a *Client) AzureUpdate(params *AzureUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureUpdateReader{formats: a.formats},
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
	success, ok := result.(*AzureUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AzureZones fetches azure zone list
*/
func (a *Client) AzureZones(params *AzureZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AzureZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAzureZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Azure_Zones",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Azure/zones",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AzureZonesReader{formats: a.formats},
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
	success, ok := result.(*AzureZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Azure_Zones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
