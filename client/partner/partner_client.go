// Code generated by go-swagger; DO NOT EDIT.

package partner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new partner API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for partner API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PartnerAddWhiteListDomain(params *PartnerAddWhiteListDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerAddWhiteListDomainOK, error)

	PartnerBecomePartner(params *PartnerBecomePartnerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerBecomePartnerOK, error)

	PartnerBindOrganizations(params *PartnerBindOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerBindOrganizationsOK, error)

	PartnerCreate(params *PartnerCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerCreateOK, error)

	PartnerDeleteWhiteListDomain(params *PartnerDeleteWhiteListDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerDeleteWhiteListDomainOK, error)

	PartnerDetails(params *PartnerDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerDetailsOK, error)

	PartnerList(params *PartnerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerListOK, error)

	PartnerPartnerInfoRegistration(params *PartnerPartnerInfoRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerPartnerInfoRegistrationOK, error)

	PartnerPartnerList(params *PartnerPartnerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerPartnerListOK, error)

	PartnerUpdate(params *PartnerUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PartnerAddWhiteListDomain partner add white list domain API
*/
func (a *Client) PartnerAddWhiteListDomain(params *PartnerAddWhiteListDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerAddWhiteListDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerAddWhiteListDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_AddWhiteListDomain",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Partner/add/whitelist/domain",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerAddWhiteListDomainReader{formats: a.formats},
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
	success, ok := result.(*PartnerAddWhiteListDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_AddWhiteListDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerBecomePartner partner become partner API
*/
func (a *Client) PartnerBecomePartner(params *PartnerBecomePartnerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerBecomePartnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerBecomePartnerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_BecomePartner",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Partner/become-a-partner",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerBecomePartnerReader{formats: a.formats},
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
	success, ok := result.(*PartnerBecomePartnerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_BecomePartner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerBindOrganizations partner bind organizations API
*/
func (a *Client) PartnerBindOrganizations(params *PartnerBindOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerBindOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerBindOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_BindOrganizations",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Partner/bindorganizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerBindOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*PartnerBindOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_BindOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerCreate adds partner
*/
func (a *Client) PartnerCreate(params *PartnerCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Partner/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerCreateReader{formats: a.formats},
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
	success, ok := result.(*PartnerCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerDeleteWhiteListDomain partner delete white list domain API
*/
func (a *Client) PartnerDeleteWhiteListDomain(params *PartnerDeleteWhiteListDomainParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerDeleteWhiteListDomainOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerDeleteWhiteListDomainParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_DeleteWhiteListDomain",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Partner/delete/whitelist/domain",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerDeleteWhiteListDomainReader{formats: a.formats},
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
	success, ok := result.(*PartnerDeleteWhiteListDomainOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_DeleteWhiteListDomain: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerDetails details of partners
*/
func (a *Client) PartnerDetails(params *PartnerDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_Details",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Partner/details",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerDetailsReader{formats: a.formats},
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
	success, ok := result.(*PartnerDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_Details: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerList gets partners
*/
func (a *Client) PartnerList(params *PartnerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Partner",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerListReader{formats: a.formats},
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
	success, ok := result.(*PartnerListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerPartnerInfoRegistration partner partner info registration API
*/
func (a *Client) PartnerPartnerInfoRegistration(params *PartnerPartnerInfoRegistrationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerPartnerInfoRegistrationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerPartnerInfoRegistrationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_PartnerInfoRegistration",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Partner/info",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerPartnerInfoRegistrationReader{formats: a.formats},
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
	success, ok := result.(*PartnerPartnerInfoRegistrationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_PartnerInfoRegistration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerPartnerList gets partners dropdown
*/
func (a *Client) PartnerPartnerList(params *PartnerPartnerListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerPartnerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerPartnerListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_PartnerList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Partner/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerPartnerListReader{formats: a.formats},
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
	success, ok := result.(*PartnerPartnerListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_PartnerList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartnerUpdate edits partner data by Id
*/
func (a *Client) PartnerUpdate(params *PartnerUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PartnerUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartnerUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Partner_Update",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Partner/update/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PartnerUpdateReader{formats: a.formats},
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
	success, ok := result.(*PartnerUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Partner_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
