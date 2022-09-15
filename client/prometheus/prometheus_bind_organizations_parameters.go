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

// NewPrometheusBindOrganizationsParams creates a new PrometheusBindOrganizationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrometheusBindOrganizationsParams() *PrometheusBindOrganizationsParams {
	return &PrometheusBindOrganizationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrometheusBindOrganizationsParamsWithTimeout creates a new PrometheusBindOrganizationsParams object
// with the ability to set a timeout on a request.
func NewPrometheusBindOrganizationsParamsWithTimeout(timeout time.Duration) *PrometheusBindOrganizationsParams {
	return &PrometheusBindOrganizationsParams{
		timeout: timeout,
	}
}

// NewPrometheusBindOrganizationsParamsWithContext creates a new PrometheusBindOrganizationsParams object
// with the ability to set a context for a request.
func NewPrometheusBindOrganizationsParamsWithContext(ctx context.Context) *PrometheusBindOrganizationsParams {
	return &PrometheusBindOrganizationsParams{
		Context: ctx,
	}
}

// NewPrometheusBindOrganizationsParamsWithHTTPClient creates a new PrometheusBindOrganizationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrometheusBindOrganizationsParamsWithHTTPClient(client *http.Client) *PrometheusBindOrganizationsParams {
	return &PrometheusBindOrganizationsParams{
		HTTPClient: client,
	}
}

/*
PrometheusBindOrganizationsParams contains all the parameters to send to the API endpoint

	for the prometheus bind organizations operation.

	Typically these are written to a http.Request.
*/
type PrometheusBindOrganizationsParams struct {

	// Body.
	Body *models.BindPrometheusOrganizationsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the prometheus bind organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBindOrganizationsParams) WithDefaults() *PrometheusBindOrganizationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the prometheus bind organizations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBindOrganizationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) WithTimeout(timeout time.Duration) *PrometheusBindOrganizationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) WithContext(ctx context.Context) *PrometheusBindOrganizationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) WithHTTPClient(client *http.Client) *PrometheusBindOrganizationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) WithBody(body *models.BindPrometheusOrganizationsCommand) *PrometheusBindOrganizationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) SetBody(body *models.BindPrometheusOrganizationsCommand) {
	o.Body = body
}

// WithV adds the v to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) WithV(v string) *PrometheusBindOrganizationsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the prometheus bind organizations params
func (o *PrometheusBindOrganizationsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PrometheusBindOrganizationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
