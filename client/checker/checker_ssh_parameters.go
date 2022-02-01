// Code generated by go-swagger; DO NOT EDIT.

package checker

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

// NewCheckerSSHParams creates a new CheckerSSHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckerSSHParams() *CheckerSSHParams {
	return &CheckerSSHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckerSSHParamsWithTimeout creates a new CheckerSSHParams object
// with the ability to set a timeout on a request.
func NewCheckerSSHParamsWithTimeout(timeout time.Duration) *CheckerSSHParams {
	return &CheckerSSHParams{
		timeout: timeout,
	}
}

// NewCheckerSSHParamsWithContext creates a new CheckerSSHParams object
// with the ability to set a context for a request.
func NewCheckerSSHParamsWithContext(ctx context.Context) *CheckerSSHParams {
	return &CheckerSSHParams{
		Context: ctx,
	}
}

// NewCheckerSSHParamsWithHTTPClient creates a new CheckerSSHParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckerSSHParamsWithHTTPClient(client *http.Client) *CheckerSSHParams {
	return &CheckerSSHParams{
		HTTPClient: client,
	}
}

/* CheckerSSHParams contains all the parameters to send to the API endpoint
   for the checker Ssh operation.

   Typically these are written to a http.Request.
*/
type CheckerSSHParams struct {

	// Body.
	Body *models.SSHKeyCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checker Ssh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerSSHParams) WithDefaults() *CheckerSSHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checker Ssh params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerSSHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checker Ssh params
func (o *CheckerSSHParams) WithTimeout(timeout time.Duration) *CheckerSSHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checker Ssh params
func (o *CheckerSSHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checker Ssh params
func (o *CheckerSSHParams) WithContext(ctx context.Context) *CheckerSSHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checker Ssh params
func (o *CheckerSSHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checker Ssh params
func (o *CheckerSSHParams) WithHTTPClient(client *http.Client) *CheckerSSHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checker Ssh params
func (o *CheckerSSHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the checker Ssh params
func (o *CheckerSSHParams) WithBody(body *models.SSHKeyCommand) *CheckerSSHParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the checker Ssh params
func (o *CheckerSSHParams) SetBody(body *models.SSHKeyCommand) {
	o.Body = body
}

// WithV adds the v to the checker Ssh params
func (o *CheckerSSHParams) WithV(v string) *CheckerSSHParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the checker Ssh params
func (o *CheckerSSHParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CheckerSSHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
