// Code generated by go-swagger; DO NOT EDIT.

package prometheus

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

	"github.com/itera-io/taikungoclient/models"
)

// NewPrometheusBindRulesParams creates a new PrometheusBindRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrometheusBindRulesParams() *PrometheusBindRulesParams {
	return &PrometheusBindRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrometheusBindRulesParamsWithTimeout creates a new PrometheusBindRulesParams object
// with the ability to set a timeout on a request.
func NewPrometheusBindRulesParamsWithTimeout(timeout time.Duration) *PrometheusBindRulesParams {
	return &PrometheusBindRulesParams{
		timeout: timeout,
	}
}

// NewPrometheusBindRulesParamsWithContext creates a new PrometheusBindRulesParams object
// with the ability to set a context for a request.
func NewPrometheusBindRulesParamsWithContext(ctx context.Context) *PrometheusBindRulesParams {
	return &PrometheusBindRulesParams{
		Context: ctx,
	}
}

// NewPrometheusBindRulesParamsWithHTTPClient creates a new PrometheusBindRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrometheusBindRulesParamsWithHTTPClient(client *http.Client) *PrometheusBindRulesParams {
	return &PrometheusBindRulesParams{
		HTTPClient: client,
	}
}

/*
PrometheusBindRulesParams contains all the parameters to send to the API endpoint

	for the prometheus bind rules operation.

	Typically these are written to a http.Request.
*/
type PrometheusBindRulesParams struct {

	// Body.
	Body *models.BindRulesCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the prometheus bind rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBindRulesParams) WithDefaults() *PrometheusBindRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the prometheus bind rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBindRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the prometheus bind rules params
func (o *PrometheusBindRulesParams) WithTimeout(timeout time.Duration) *PrometheusBindRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prometheus bind rules params
func (o *PrometheusBindRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prometheus bind rules params
func (o *PrometheusBindRulesParams) WithContext(ctx context.Context) *PrometheusBindRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prometheus bind rules params
func (o *PrometheusBindRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prometheus bind rules params
func (o *PrometheusBindRulesParams) WithHTTPClient(client *http.Client) *PrometheusBindRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prometheus bind rules params
func (o *PrometheusBindRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the prometheus bind rules params
func (o *PrometheusBindRulesParams) WithBody(body *models.BindRulesCommand) *PrometheusBindRulesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the prometheus bind rules params
func (o *PrometheusBindRulesParams) SetBody(body *models.BindRulesCommand) {
	o.Body = body
}

// WithV adds the v to the prometheus bind rules params
func (o *PrometheusBindRulesParams) WithV(v string) *PrometheusBindRulesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the prometheus bind rules params
func (o *PrometheusBindRulesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PrometheusBindRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
