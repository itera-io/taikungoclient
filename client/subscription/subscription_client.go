// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SubscriptionBind(params *SubscriptionBindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionBindOK, error)

	SubscriptionCreate(params *SubscriptionCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionCreateOK, error)

	SubscriptionDelete(params *SubscriptionDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionDeleteOK, *SubscriptionDeleteNoContent, error)

	SubscriptionList(params *SubscriptionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionListOK, error)

	SubscriptionSubscriptionForLandingPage(params *SubscriptionSubscriptionForLandingPageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionSubscriptionForLandingPageOK, error)

	SubscriptionSubscriptionForOrganizationList(params *SubscriptionSubscriptionForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionSubscriptionForOrganizationListOK, error)

	SubscriptionUpdate(params *SubscriptionUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
SubscriptionBind binds subscription
*/
func (a *Client) SubscriptionBind(params *SubscriptionBindParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionBindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionBindParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_Bind",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Subscription/bind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionBindReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionBindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_Bind: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionCreate adds new subscription package
*/
func (a *Client) SubscriptionCreate(params *SubscriptionCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Subscription/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionCreateReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionDelete deletes subscription package
*/
func (a *Client) SubscriptionDelete(params *SubscriptionDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionDeleteOK, *SubscriptionDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Subscription/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionDeleteReader{formats: a.formats},
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
	case *SubscriptionDeleteOK:
		return value, nil, nil
	case *SubscriptionDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionList retrieves private subscription list for partners
*/
func (a *Client) SubscriptionList(params *SubscriptionListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Subscription",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionListReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionSubscriptionForLandingPage retrieves public subscription list for landing page
*/
func (a *Client) SubscriptionSubscriptionForLandingPage(params *SubscriptionSubscriptionForLandingPageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionSubscriptionForLandingPageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionSubscriptionForLandingPageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_SubscriptionForLandingPage",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Subscription/public",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionSubscriptionForLandingPageReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionSubscriptionForLandingPageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_SubscriptionForLandingPage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionSubscriptionForOrganizationList retrieves subscription for organization bound
*/
func (a *Client) SubscriptionSubscriptionForOrganizationList(params *SubscriptionSubscriptionForOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionSubscriptionForOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionSubscriptionForOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_SubscriptionForOrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Subscription/boundlist",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionSubscriptionForOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionSubscriptionForOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_SubscriptionForOrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscriptionUpdate adds new subscription package
*/
func (a *Client) SubscriptionUpdate(params *SubscriptionUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SubscriptionUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscription_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Subscription/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionUpdateReader{formats: a.formats},
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
	success, ok := result.(*SubscriptionUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscription_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
