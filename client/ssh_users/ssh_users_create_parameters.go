// Code generated by go-swagger; DO NOT EDIT.

package ssh_users

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

// NewSSHUsersCreateParams creates a new SSHUsersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSHUsersCreateParams() *SSHUsersCreateParams {
	return &SSHUsersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSHUsersCreateParamsWithTimeout creates a new SSHUsersCreateParams object
// with the ability to set a timeout on a request.
func NewSSHUsersCreateParamsWithTimeout(timeout time.Duration) *SSHUsersCreateParams {
	return &SSHUsersCreateParams{
		timeout: timeout,
	}
}

// NewSSHUsersCreateParamsWithContext creates a new SSHUsersCreateParams object
// with the ability to set a context for a request.
func NewSSHUsersCreateParamsWithContext(ctx context.Context) *SSHUsersCreateParams {
	return &SSHUsersCreateParams{
		Context: ctx,
	}
}

// NewSSHUsersCreateParamsWithHTTPClient creates a new SSHUsersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSHUsersCreateParamsWithHTTPClient(client *http.Client) *SSHUsersCreateParams {
	return &SSHUsersCreateParams{
		HTTPClient: client,
	}
}

/* SSHUsersCreateParams contains all the parameters to send to the API endpoint
   for the Ssh users create operation.

   Typically these are written to a http.Request.
*/
type SSHUsersCreateParams struct {

	// Body.
	Body *models.CreateSSHUserCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Ssh users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHUsersCreateParams) WithDefaults() *SSHUsersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Ssh users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHUsersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Ssh users create params
func (o *SSHUsersCreateParams) WithTimeout(timeout time.Duration) *SSHUsersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Ssh users create params
func (o *SSHUsersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Ssh users create params
func (o *SSHUsersCreateParams) WithContext(ctx context.Context) *SSHUsersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Ssh users create params
func (o *SSHUsersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Ssh users create params
func (o *SSHUsersCreateParams) WithHTTPClient(client *http.Client) *SSHUsersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Ssh users create params
func (o *SSHUsersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the Ssh users create params
func (o *SSHUsersCreateParams) WithBody(body *models.CreateSSHUserCommand) *SSHUsersCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the Ssh users create params
func (o *SSHUsersCreateParams) SetBody(body *models.CreateSSHUserCommand) {
	o.Body = body
}

// WithV adds the v to the Ssh users create params
func (o *SSHUsersCreateParams) WithV(v string) *SSHUsersCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the Ssh users create params
func (o *SSHUsersCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SSHUsersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
