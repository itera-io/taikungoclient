// Code generated by go-swagger; DO NOT EDIT.

package allowed_host

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

// NewAllowedHostCreateParams creates a new AllowedHostCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllowedHostCreateParams() *AllowedHostCreateParams {
	return &AllowedHostCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllowedHostCreateParamsWithTimeout creates a new AllowedHostCreateParams object
// with the ability to set a timeout on a request.
func NewAllowedHostCreateParamsWithTimeout(timeout time.Duration) *AllowedHostCreateParams {
	return &AllowedHostCreateParams{
		timeout: timeout,
	}
}

// NewAllowedHostCreateParamsWithContext creates a new AllowedHostCreateParams object
// with the ability to set a context for a request.
func NewAllowedHostCreateParamsWithContext(ctx context.Context) *AllowedHostCreateParams {
	return &AllowedHostCreateParams{
		Context: ctx,
	}
}

// NewAllowedHostCreateParamsWithHTTPClient creates a new AllowedHostCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllowedHostCreateParamsWithHTTPClient(client *http.Client) *AllowedHostCreateParams {
	return &AllowedHostCreateParams{
		HTTPClient: client,
	}
}

/* AllowedHostCreateParams contains all the parameters to send to the API endpoint
   for the allowed host create operation.

   Typically these are written to a http.Request.
*/
type AllowedHostCreateParams struct {

	// Body.
	Body *models.CreateAllowedHostCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allowed host create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllowedHostCreateParams) WithDefaults() *AllowedHostCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allowed host create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllowedHostCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allowed host create params
func (o *AllowedHostCreateParams) WithTimeout(timeout time.Duration) *AllowedHostCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allowed host create params
func (o *AllowedHostCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allowed host create params
func (o *AllowedHostCreateParams) WithContext(ctx context.Context) *AllowedHostCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allowed host create params
func (o *AllowedHostCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allowed host create params
func (o *AllowedHostCreateParams) WithHTTPClient(client *http.Client) *AllowedHostCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allowed host create params
func (o *AllowedHostCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the allowed host create params
func (o *AllowedHostCreateParams) WithBody(body *models.CreateAllowedHostCommand) *AllowedHostCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the allowed host create params
func (o *AllowedHostCreateParams) SetBody(body *models.CreateAllowedHostCommand) {
	o.Body = body
}

// WithV adds the v to the allowed host create params
func (o *AllowedHostCreateParams) WithV(v string) *AllowedHostCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the allowed host create params
func (o *AllowedHostCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AllowedHostCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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