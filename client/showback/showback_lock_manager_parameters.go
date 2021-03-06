// Code generated by go-swagger; DO NOT EDIT.

package showback

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

// NewShowbackLockManagerParams creates a new ShowbackLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackLockManagerParams() *ShowbackLockManagerParams {
	return &ShowbackLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackLockManagerParamsWithTimeout creates a new ShowbackLockManagerParams object
// with the ability to set a timeout on a request.
func NewShowbackLockManagerParamsWithTimeout(timeout time.Duration) *ShowbackLockManagerParams {
	return &ShowbackLockManagerParams{
		timeout: timeout,
	}
}

// NewShowbackLockManagerParamsWithContext creates a new ShowbackLockManagerParams object
// with the ability to set a context for a request.
func NewShowbackLockManagerParamsWithContext(ctx context.Context) *ShowbackLockManagerParams {
	return &ShowbackLockManagerParams{
		Context: ctx,
	}
}

// NewShowbackLockManagerParamsWithHTTPClient creates a new ShowbackLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackLockManagerParamsWithHTTPClient(client *http.Client) *ShowbackLockManagerParams {
	return &ShowbackLockManagerParams{
		HTTPClient: client,
	}
}

/* ShowbackLockManagerParams contains all the parameters to send to the API endpoint
   for the showback lock manager operation.

   Typically these are written to a http.Request.
*/
type ShowbackLockManagerParams struct {

	// Body.
	Body *models.ShowbackCredentialLockCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackLockManagerParams) WithDefaults() *ShowbackLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback lock manager params
func (o *ShowbackLockManagerParams) WithTimeout(timeout time.Duration) *ShowbackLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback lock manager params
func (o *ShowbackLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback lock manager params
func (o *ShowbackLockManagerParams) WithContext(ctx context.Context) *ShowbackLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback lock manager params
func (o *ShowbackLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback lock manager params
func (o *ShowbackLockManagerParams) WithHTTPClient(client *http.Client) *ShowbackLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback lock manager params
func (o *ShowbackLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback lock manager params
func (o *ShowbackLockManagerParams) WithBody(body *models.ShowbackCredentialLockCommand) *ShowbackLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback lock manager params
func (o *ShowbackLockManagerParams) SetBody(body *models.ShowbackCredentialLockCommand) {
	o.Body = body
}

// WithV adds the v to the showback lock manager params
func (o *ShowbackLockManagerParams) WithV(v string) *ShowbackLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback lock manager params
func (o *ShowbackLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
