// Code generated by go-swagger; DO NOT EDIT.

package prometheus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new prometheus API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for prometheus API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PrometheusBillingList(params *PrometheusBillingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBillingListOK, error)

	PrometheusBindOrganizations(params *PrometheusBindOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBindOrganizationsOK, error)

	PrometheusBindRules(params *PrometheusBindRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBindRulesOK, error)

	PrometheusCreate(params *PrometheusCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusCreateOK, error)

	PrometheusCreateBilling(params *PrometheusCreateBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusCreateBillingOK, error)

	PrometheusDelete(params *PrometheusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusDeleteOK, error)

	PrometheusDetails(params *PrometheusDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusDetailsOK, error)

	PrometheusExportCsv(params *PrometheusExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusExportCsvOK, error)

	PrometheusGroupedList(params *PrometheusGroupedListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusGroupedListOK, error)

	PrometheusListOfRules(params *PrometheusListOfRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusListOfRulesOK, error)

	PrometheusMetricName(params *PrometheusMetricNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusMetricNameOK, error)

	PrometheusUpdate(params *PrometheusUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PrometheusBillingList retrieves all prometheus billing
*/
func (a *Client) PrometheusBillingList(params *PrometheusBillingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBillingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusBillingListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_BillingList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Prometheus/billing",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusBillingListReader{formats: a.formats},
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
	success, ok := result.(*PrometheusBillingListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_BillingList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusBindOrganizations binds organizations to prometheus rule
*/
func (a *Client) PrometheusBindOrganizations(params *PrometheusBindOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBindOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusBindOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_BindOrganizations",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus/bindorganizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusBindOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*PrometheusBindOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_BindOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusBindRules binds rules to organizations
*/
func (a *Client) PrometheusBindRules(params *PrometheusBindRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusBindRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusBindRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_BindRules",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus/bindrules",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusBindRulesReader{formats: a.formats},
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
	success, ok := result.(*PrometheusBindRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_BindRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusCreate adds prometheus rule
*/
func (a *Client) PrometheusCreate(params *PrometheusCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusCreateReader{formats: a.formats},
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
	success, ok := result.(*PrometheusCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusCreateBilling adds prometheus billing
*/
func (a *Client) PrometheusCreateBilling(params *PrometheusCreateBillingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusCreateBillingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusCreateBillingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_CreateBilling",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus/billing",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusCreateBillingReader{formats: a.formats},
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
	success, ok := result.(*PrometheusCreateBillingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_CreateBilling: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusDelete removes prometheus rule
*/
func (a *Client) PrometheusDelete(params *PrometheusDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/Prometheus/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusDeleteReader{formats: a.formats},
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
	success, ok := result.(*PrometheusDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusDetails retrieves all prometheus rules with detailed info
*/
func (a *Client) PrometheusDetails(params *PrometheusDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_Details",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Prometheus/details",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusDetailsReader{formats: a.formats},
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
	success, ok := result.(*PrometheusDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_Details: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusExportCsv exports csv file
*/
func (a *Client) PrometheusExportCsv(params *PrometheusExportCsvParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusExportCsvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusExportCsvParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_ExportCsv",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Prometheus/export",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusExportCsvReader{formats: a.formats},
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
	success, ok := result.(*PrometheusExportCsvOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_ExportCsv: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusGroupedList retrieves a list of grouped prometheus billing
*/
func (a *Client) PrometheusGroupedList(params *PrometheusGroupedListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusGroupedListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusGroupedListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_GroupedList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Prometheus/grouped",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusGroupedListReader{formats: a.formats},
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
	success, ok := result.(*PrometheusGroupedListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_GroupedList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusListOfRules retrieves a list of prometheus rules
*/
func (a *Client) PrometheusListOfRules(params *PrometheusListOfRulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusListOfRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusListOfRulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_ListOfRules",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Prometheus",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusListOfRulesReader{formats: a.formats},
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
	success, ok := result.(*PrometheusListOfRulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_ListOfRules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusMetricName fetches prometheus metric names
*/
func (a *Client) PrometheusMetricName(params *PrometheusMetricNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusMetricNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusMetricNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_MetricName",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus/metricname",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusMetricNameReader{formats: a.formats},
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
	success, ok := result.(*PrometheusMetricNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_MetricName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PrometheusUpdate edits prometheus rule
*/
func (a *Client) PrometheusUpdate(params *PrometheusUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PrometheusUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPrometheusUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Prometheus_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Prometheus/update/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PrometheusUpdateReader{formats: a.formats},
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
	success, ok := result.(*PrometheusUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Prometheus_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
