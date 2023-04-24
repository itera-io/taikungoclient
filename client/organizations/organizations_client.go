// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	OrganizationsAcceptOffer(params *OrganizationsAcceptOfferParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsAcceptOfferOK, error)

	OrganizationsCreate(params *OrganizationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsCreateOK, error)

	OrganizationsDelete(params *OrganizationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsDeleteOK, *OrganizationsDeleteNoContent, error)

	OrganizationsDetails(params *OrganizationsDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsDetailsOK, error)

	OrganizationsExportCsv(params *OrganizationsExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsExportCsvOK, error)

	OrganizationsLeaveTaikun(params *OrganizationsLeaveTaikunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsLeaveTaikunOK, error)

	OrganizationsList(params *OrganizationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsListOK, error)

	OrganizationsOrganizationList(params *OrganizationsOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsOrganizationListOK, error)

	OrganizationsToggleKeycloak(params *OrganizationsToggleKeycloakParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsToggleKeycloakOK, error)

	OrganizationsUpdate(params *OrganizationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdateOK, error)

	OrganizationsUpdatePaymentMethod(params *OrganizationsUpdatePaymentMethodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdatePaymentMethodOK, error)

	OrganizationsUpdateSubscription(params *OrganizationsUpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdateSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
OrganizationsAcceptOffer accepts discount offer
*/
func (a *Client) OrganizationsAcceptOffer(params *OrganizationsAcceptOfferParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsAcceptOfferOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsAcceptOfferParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_AcceptOffer",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/accept-offer",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsAcceptOfferReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsAcceptOfferOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_AcceptOffer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsCreate creates a new organization only available for admins
*/
func (a *Client) OrganizationsCreate(params *OrganizationsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsCreateReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsDelete deletes the specified organization only available for admins
*/
func (a *Client) OrganizationsDelete(params *OrganizationsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsDeleteOK, *OrganizationsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/delete/{organizationId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsDeleteReader{formats: a.formats},
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
	case *OrganizationsDeleteOK:
		return value, nil, nil
	case *OrganizationsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for organizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsDetails retrieves all data about current organization by Id
*/
func (a *Client) OrganizationsDetails(params *OrganizationsDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_Details",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Organizations/details",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsDetailsReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_Details: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsExportCsv exports csv file
*/
func (a *Client) OrganizationsExportCsv(params *OrganizationsExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsExportCsvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsExportCsvParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_ExportCsv",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Organizations/export",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsExportCsvReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsExportCsvOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_ExportCsv: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsLeaveTaikun leaves taikun
*/
func (a *Client) OrganizationsLeaveTaikun(params *OrganizationsLeaveTaikunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsLeaveTaikunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsLeaveTaikunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_LeaveTaikun",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/leave",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsLeaveTaikunReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsLeaveTaikunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_LeaveTaikun: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsList retrieves all organizations
*/
func (a *Client) OrganizationsList(params *OrganizationsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Organizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsListReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsOrganizationList retrieves organizations
*/
func (a *Client) OrganizationsOrganizationList(params *OrganizationsOrganizationListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsOrganizationListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsOrganizationListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_OrganizationList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Organizations/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsOrganizationListReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsOrganizationListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_OrganizationList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsToggleKeycloak toggles keycloak login option
*/
func (a *Client) OrganizationsToggleKeycloak(params *OrganizationsToggleKeycloakParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsToggleKeycloakOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsToggleKeycloakParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_ToggleKeycloak",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/toggle/keycloak",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsToggleKeycloakReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsToggleKeycloakOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_ToggleKeycloak: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsUpdate updates organization by Id
*/
func (a *Client) OrganizationsUpdate(params *OrganizationsUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsUpdateReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsUpdatePaymentMethod updates organization payment Id
*/
func (a *Client) OrganizationsUpdatePaymentMethod(params *OrganizationsUpdatePaymentMethodParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdatePaymentMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsUpdatePaymentMethodParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_UpdatePaymentMethod",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/updatepaymentmethod",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsUpdatePaymentMethodReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsUpdatePaymentMethodOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_UpdatePaymentMethod: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrganizationsUpdateSubscription updates subscription
*/
func (a *Client) OrganizationsUpdateSubscription(params *OrganizationsUpdateSubscriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrganizationsUpdateSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrganizationsUpdateSubscriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Organizations_UpdateSubscription",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Organizations/updatesubscription",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrganizationsUpdateSubscriptionReader{formats: a.formats},
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
	success, ok := result.(*OrganizationsUpdateSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Organizations_UpdateSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
