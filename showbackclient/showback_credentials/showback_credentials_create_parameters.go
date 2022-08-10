// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

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

// NewShowbackCredentialsCreateParams creates a new ShowbackCredentialsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackCredentialsCreateParams() *ShowbackCredentialsCreateParams {
	return &ShowbackCredentialsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackCredentialsCreateParamsWithTimeout creates a new ShowbackCredentialsCreateParams object
// with the ability to set a timeout on a request.
func NewShowbackCredentialsCreateParamsWithTimeout(timeout time.Duration) *ShowbackCredentialsCreateParams {
	return &ShowbackCredentialsCreateParams{
		timeout: timeout,
	}
}

// NewShowbackCredentialsCreateParamsWithContext creates a new ShowbackCredentialsCreateParams object
// with the ability to set a context for a request.
func NewShowbackCredentialsCreateParamsWithContext(ctx context.Context) *ShowbackCredentialsCreateParams {
	return &ShowbackCredentialsCreateParams{
		Context: ctx,
	}
}

// NewShowbackCredentialsCreateParamsWithHTTPClient creates a new ShowbackCredentialsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackCredentialsCreateParamsWithHTTPClient(client *http.Client) *ShowbackCredentialsCreateParams {
	return &ShowbackCredentialsCreateParams{
		HTTPClient: client,
	}
}

/* ShowbackCredentialsCreateParams contains all the parameters to send to the API endpoint
   for the showback credentials create operation.

   Typically these are written to a http.Request.
*/
type ShowbackCredentialsCreateParams struct {

	// Body.
	Body *models.CreateShowbackCredentialCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsCreateParams) WithDefaults() *ShowbackCredentialsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback credentials create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) WithTimeout(timeout time.Duration) *ShowbackCredentialsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) WithContext(ctx context.Context) *ShowbackCredentialsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) WithHTTPClient(client *http.Client) *ShowbackCredentialsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) WithBody(body *models.CreateShowbackCredentialCommand) *ShowbackCredentialsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) SetBody(body *models.CreateShowbackCredentialCommand) {
	o.Body = body
}

// WithV adds the v to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) WithV(v string) *ShowbackCredentialsCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback credentials create params
func (o *ShowbackCredentialsCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackCredentialsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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