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

	"github.com/itera-io/taikungoclient/models"
)

// NewServersDeleteParams creates a new ServersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServersDeleteParams() *ServersDeleteParams {
	return &ServersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServersDeleteParamsWithTimeout creates a new ServersDeleteParams object
// with the ability to set a timeout on a request.
func NewServersDeleteParamsWithTimeout(timeout time.Duration) *ServersDeleteParams {
	return &ServersDeleteParams{
		timeout: timeout,
	}
}

// NewServersDeleteParamsWithContext creates a new ServersDeleteParams object
// with the ability to set a context for a request.
func NewServersDeleteParamsWithContext(ctx context.Context) *ServersDeleteParams {
	return &ServersDeleteParams{
		Context: ctx,
	}
}

// NewServersDeleteParamsWithHTTPClient creates a new ServersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewServersDeleteParamsWithHTTPClient(client *http.Client) *ServersDeleteParams {
	return &ServersDeleteParams{
		HTTPClient: client,
	}
}

/*
ServersDeleteParams contains all the parameters to send to the API endpoint

	for the servers delete operation.

	Typically these are written to a http.Request.
*/
type ServersDeleteParams struct {

	// Body.
	Body *models.DeleteServerCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersDeleteParams) WithDefaults() *ServersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the servers delete params
func (o *ServersDeleteParams) WithTimeout(timeout time.Duration) *ServersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers delete params
func (o *ServersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers delete params
func (o *ServersDeleteParams) WithContext(ctx context.Context) *ServersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers delete params
func (o *ServersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers delete params
func (o *ServersDeleteParams) WithHTTPClient(client *http.Client) *ServersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers delete params
func (o *ServersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the servers delete params
func (o *ServersDeleteParams) WithBody(body *models.DeleteServerCommand) *ServersDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the servers delete params
func (o *ServersDeleteParams) SetBody(body *models.DeleteServerCommand) {
	o.Body = body
}

// WithV adds the v to the servers delete params
func (o *ServersDeleteParams) WithV(v string) *ServersDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the servers delete params
func (o *ServersDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ServersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
