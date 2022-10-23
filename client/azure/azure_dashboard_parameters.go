// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAzureDashboardParams creates a new AzureDashboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureDashboardParams() *AzureDashboardParams {
	return &AzureDashboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureDashboardParamsWithTimeout creates a new AzureDashboardParams object
// with the ability to set a timeout on a request.
func NewAzureDashboardParamsWithTimeout(timeout time.Duration) *AzureDashboardParams {
	return &AzureDashboardParams{
		timeout: timeout,
	}
}

// NewAzureDashboardParamsWithContext creates a new AzureDashboardParams object
// with the ability to set a context for a request.
func NewAzureDashboardParamsWithContext(ctx context.Context) *AzureDashboardParams {
	return &AzureDashboardParams{
		Context: ctx,
	}
}

// NewAzureDashboardParamsWithHTTPClient creates a new AzureDashboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureDashboardParamsWithHTTPClient(client *http.Client) *AzureDashboardParams {
	return &AzureDashboardParams{
		HTTPClient: client,
	}
}

/*
AzureDashboardParams contains all the parameters to send to the API endpoint

	for the azure dashboard operation.

	Typically these are written to a http.Request.
*/
type AzureDashboardParams struct {

	// Body.
	Body AzureDashboardBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureDashboardParams) WithDefaults() *AzureDashboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure dashboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureDashboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the azure dashboard params
func (o *AzureDashboardParams) WithTimeout(timeout time.Duration) *AzureDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure dashboard params
func (o *AzureDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure dashboard params
func (o *AzureDashboardParams) WithContext(ctx context.Context) *AzureDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure dashboard params
func (o *AzureDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure dashboard params
func (o *AzureDashboardParams) WithHTTPClient(client *http.Client) *AzureDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure dashboard params
func (o *AzureDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the azure dashboard params
func (o *AzureDashboardParams) WithBody(body AzureDashboardBody) *AzureDashboardParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the azure dashboard params
func (o *AzureDashboardParams) SetBody(body AzureDashboardBody) {
	o.Body = body
}

// WithV adds the v to the azure dashboard params
func (o *AzureDashboardParams) WithV(v string) *AzureDashboardParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the azure dashboard params
func (o *AzureDashboardParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AzureDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
