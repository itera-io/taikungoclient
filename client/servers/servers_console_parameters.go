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

// NewServersConsoleParams creates a new ServersConsoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServersConsoleParams() *ServersConsoleParams {
	return &ServersConsoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServersConsoleParamsWithTimeout creates a new ServersConsoleParams object
// with the ability to set a timeout on a request.
func NewServersConsoleParamsWithTimeout(timeout time.Duration) *ServersConsoleParams {
	return &ServersConsoleParams{
		timeout: timeout,
	}
}

// NewServersConsoleParamsWithContext creates a new ServersConsoleParams object
// with the ability to set a context for a request.
func NewServersConsoleParamsWithContext(ctx context.Context) *ServersConsoleParams {
	return &ServersConsoleParams{
		Context: ctx,
	}
}

// NewServersConsoleParamsWithHTTPClient creates a new ServersConsoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewServersConsoleParamsWithHTTPClient(client *http.Client) *ServersConsoleParams {
	return &ServersConsoleParams{
		HTTPClient: client,
	}
}

/* ServersConsoleParams contains all the parameters to send to the API endpoint
   for the servers console operation.

   Typically these are written to a http.Request.
*/
type ServersConsoleParams struct {

	// Body.
	Body *models.ConsoleScreenshotCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the servers console params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersConsoleParams) WithDefaults() *ServersConsoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the servers console params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServersConsoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the servers console params
func (o *ServersConsoleParams) WithTimeout(timeout time.Duration) *ServersConsoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers console params
func (o *ServersConsoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers console params
func (o *ServersConsoleParams) WithContext(ctx context.Context) *ServersConsoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers console params
func (o *ServersConsoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers console params
func (o *ServersConsoleParams) WithHTTPClient(client *http.Client) *ServersConsoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers console params
func (o *ServersConsoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the servers console params
func (o *ServersConsoleParams) WithBody(body *models.ConsoleScreenshotCommand) *ServersConsoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the servers console params
func (o *ServersConsoleParams) SetBody(body *models.ConsoleScreenshotCommand) {
	o.Body = body
}

// WithV adds the v to the servers console params
func (o *ServersConsoleParams) WithV(v string) *ServersConsoleParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the servers console params
func (o *ServersConsoleParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ServersConsoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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