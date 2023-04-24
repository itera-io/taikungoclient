// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for payment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PaymentClearPayment(params *PaymentClearPaymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentClearPaymentOK, error)

	PaymentCreateCustomer(params *PaymentCreateCustomerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentCreateCustomerOK, error)

	PaymentGetBillingInfo(params *PaymentGetBillingInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetBillingInfoOK, error)

	PaymentGetCardInfo(params *PaymentGetCardInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetCardInfoOK, error)

	PaymentGetFinalPrice(params *PaymentGetFinalPriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetFinalPriceOK, error)

	PaymentGetStripeInvoices(params *PaymentGetStripeInvoicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetStripeInvoicesOK, error)

	PaymentPayInvoice(params *PaymentPayInvoiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentPayInvoiceOK, error)

	PaymentUpdateCard(params *PaymentUpdateCardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentUpdateCardOK, error)

	PaymentWebhook(params *PaymentWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentWebhookOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PaymentClearPayment clears payment
*/
func (a *Client) PaymentClearPayment(params *PaymentClearPaymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentClearPaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentClearPaymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_ClearPayment",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/clear",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentClearPaymentReader{formats: a.formats},
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
	success, ok := result.(*PaymentClearPaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_ClearPayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentCreateCustomer creates customer
*/
func (a *Client) PaymentCreateCustomer(params *PaymentCreateCustomerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentCreateCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentCreateCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_CreateCustomer",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/createcustomer",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentCreateCustomerReader{formats: a.formats},
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
	success, ok := result.(*PaymentCreateCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_CreateCustomer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentGetBillingInfo gets billing info for organization
*/
func (a *Client) PaymentGetBillingInfo(params *PaymentGetBillingInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetBillingInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentGetBillingInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_GetBillingInfo",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Payment/billing-info",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentGetBillingInfoReader{formats: a.formats},
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
	success, ok := result.(*PaymentGetBillingInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_GetBillingInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentGetCardInfo gets card information
*/
func (a *Client) PaymentGetCardInfo(params *PaymentGetCardInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetCardInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentGetCardInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_GetCardInfo",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Payment/cardinfo",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentGetCardInfoReader{formats: a.formats},
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
	success, ok := result.(*PaymentGetCardInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_GetCardInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentGetFinalPrice updates your card
*/
func (a *Client) PaymentGetFinalPrice(params *PaymentGetFinalPriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetFinalPriceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentGetFinalPriceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_GetFinalPrice",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/finalprice",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentGetFinalPriceReader{formats: a.formats},
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
	success, ok := result.(*PaymentGetFinalPriceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_GetFinalPrice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentGetStripeInvoices gets required stripe invoices by stripe subscription id
*/
func (a *Client) PaymentGetStripeInvoices(params *PaymentGetStripeInvoicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentGetStripeInvoicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentGetStripeInvoicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_GetStripeInvoices",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Payment/stripeinvoices/{subscriptionId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentGetStripeInvoicesReader{formats: a.formats},
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
	success, ok := result.(*PaymentGetStripeInvoicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_GetStripeInvoices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentPayInvoice pays invoice
*/
func (a *Client) PaymentPayInvoice(params *PaymentPayInvoiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentPayInvoiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentPayInvoiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_PayInvoice",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/pay",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentPayInvoiceReader{formats: a.formats},
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
	success, ok := result.(*PaymentPayInvoiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_PayInvoice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentUpdateCard updates your card
*/
func (a *Client) PaymentUpdateCard(params *PaymentUpdateCardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentUpdateCardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentUpdateCardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_UpdateCard",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/updatecard",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentUpdateCardReader{formats: a.formats},
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
	success, ok := result.(*PaymentUpdateCardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_UpdateCard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PaymentWebhook listens to payment webhook
*/
func (a *Client) PaymentWebhook(params *PaymentWebhookParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PaymentWebhookOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPaymentWebhookParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Payment_Webhook",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Payment/webhook",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PaymentWebhookReader{formats: a.formats},
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
	success, ok := result.(*PaymentWebhookOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Payment_Webhook: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
