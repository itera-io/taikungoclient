// Code generated by go-swagger; DO NOT EDIT.

package ops_credentials

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

// NewOpsCredentialsLockManagerParams creates a new OpsCredentialsLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpsCredentialsLockManagerParams() *OpsCredentialsLockManagerParams {
	return &OpsCredentialsLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpsCredentialsLockManagerParamsWithTimeout creates a new OpsCredentialsLockManagerParams object
// with the ability to set a timeout on a request.
func NewOpsCredentialsLockManagerParamsWithTimeout(timeout time.Duration) *OpsCredentialsLockManagerParams {
	return &OpsCredentialsLockManagerParams{
		timeout: timeout,
	}
}

// NewOpsCredentialsLockManagerParamsWithContext creates a new OpsCredentialsLockManagerParams object
// with the ability to set a context for a request.
func NewOpsCredentialsLockManagerParamsWithContext(ctx context.Context) *OpsCredentialsLockManagerParams {
	return &OpsCredentialsLockManagerParams{
		Context: ctx,
	}
}

// NewOpsCredentialsLockManagerParamsWithHTTPClient creates a new OpsCredentialsLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpsCredentialsLockManagerParamsWithHTTPClient(client *http.Client) *OpsCredentialsLockManagerParams {
	return &OpsCredentialsLockManagerParams{
		HTTPClient: client,
	}
}

/* OpsCredentialsLockManagerParams contains all the parameters to send to the API endpoint
   for the ops credentials lock manager operation.

   Typically these are written to a http.Request.
*/
type OpsCredentialsLockManagerParams struct {

	// Body.
	Body *models.OperationCredentialLockManagerCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ops credentials lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsLockManagerParams) WithDefaults() *OpsCredentialsLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ops credentials lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) WithTimeout(timeout time.Duration) *OpsCredentialsLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) WithContext(ctx context.Context) *OpsCredentialsLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) WithHTTPClient(client *http.Client) *OpsCredentialsLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) WithBody(body *models.OperationCredentialLockManagerCommand) *OpsCredentialsLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) SetBody(body *models.OperationCredentialLockManagerCommand) {
	o.Body = body
}

// WithV adds the v to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) WithV(v string) *OpsCredentialsLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ops credentials lock manager params
func (o *OpsCredentialsLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpsCredentialsLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
