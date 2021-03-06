// Code generated by go-swagger; DO NOT EDIT.

package opa_profiles

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

// NewOpaProfilesLockManagerParams creates a new OpaProfilesLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpaProfilesLockManagerParams() *OpaProfilesLockManagerParams {
	return &OpaProfilesLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpaProfilesLockManagerParamsWithTimeout creates a new OpaProfilesLockManagerParams object
// with the ability to set a timeout on a request.
func NewOpaProfilesLockManagerParamsWithTimeout(timeout time.Duration) *OpaProfilesLockManagerParams {
	return &OpaProfilesLockManagerParams{
		timeout: timeout,
	}
}

// NewOpaProfilesLockManagerParamsWithContext creates a new OpaProfilesLockManagerParams object
// with the ability to set a context for a request.
func NewOpaProfilesLockManagerParamsWithContext(ctx context.Context) *OpaProfilesLockManagerParams {
	return &OpaProfilesLockManagerParams{
		Context: ctx,
	}
}

// NewOpaProfilesLockManagerParamsWithHTTPClient creates a new OpaProfilesLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpaProfilesLockManagerParamsWithHTTPClient(client *http.Client) *OpaProfilesLockManagerParams {
	return &OpaProfilesLockManagerParams{
		HTTPClient: client,
	}
}

/* OpaProfilesLockManagerParams contains all the parameters to send to the API endpoint
   for the opa profiles lock manager operation.

   Typically these are written to a http.Request.
*/
type OpaProfilesLockManagerParams struct {

	// Body.
	Body *models.OpaProfileLockManagerCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the opa profiles lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesLockManagerParams) WithDefaults() *OpaProfilesLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the opa profiles lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) WithTimeout(timeout time.Duration) *OpaProfilesLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) WithContext(ctx context.Context) *OpaProfilesLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) WithHTTPClient(client *http.Client) *OpaProfilesLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) WithBody(body *models.OpaProfileLockManagerCommand) *OpaProfilesLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) SetBody(body *models.OpaProfileLockManagerCommand) {
	o.Body = body
}

// WithV adds the v to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) WithV(v string) *OpaProfilesLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the opa profiles lock manager params
func (o *OpaProfilesLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpaProfilesLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
