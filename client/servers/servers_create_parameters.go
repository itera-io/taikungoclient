// Code generated by go-swagger; DO NOT EDIT.

package servers

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

// NewServersCreateParams creates a new ServersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServersCreateParams() *ServersCreateParams {
	return &ServersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServersCreateParamsWithTimeout creates a new ServersCreateParams object
// with the ability to set a timeout on a request.
func NewServersCreateParamsWithTimeout(timeout time.Duration) *ServersCreateParams {
	return &ServersCreateParams{
		timeout: timeout,
	}
}

// NewServersCreateParamsWithContext creates a new ServersCreateParams object
// with the ability to set a context for a request.
func NewServersCreateParamsWithContext(ctx context.Context) *ServersCreateParams {
	return &ServersCreateParams{
		Context: ctx,
	}
}

// NewServersCreateParamsWithHTTPClient creates a new ServersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewServersCreateParamsWithHTTPClient(client *http.Client) *ServersCreateParams {
	return &ServersCreateParams{
		HTTPClient: client,
	}
}

/*
ServersCreateParams contains all the parameters to send to the API endpoint

	for the servers create operation.

	Typically these are written to a http.Request.
*/
type ServersCreateParams struct {

	// Body.
	Body ServersCreateBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersCreateParams) WithDefaults() *ServersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the servers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the servers create params
func (o *ServersCreateParams) WithTimeout(timeout time.Duration) *ServersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers create params
func (o *ServersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers create params
func (o *ServersCreateParams) WithContext(ctx context.Context) *ServersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers create params
func (o *ServersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers create params
func (o *ServersCreateParams) WithHTTPClient(client *http.Client) *ServersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers create params
func (o *ServersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the servers create params
func (o *ServersCreateParams) WithBody(body ServersCreateBody) *ServersCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the servers create params
func (o *ServersCreateParams) SetBody(body ServersCreateBody) {
	o.Body = body
}

// WithV adds the v to the servers create params
func (o *ServersCreateParams) WithV(v string) *ServersCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the servers create params
func (o *ServersCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ServersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
