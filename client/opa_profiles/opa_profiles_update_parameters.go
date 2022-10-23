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

// NewOpaProfilesUpdateParams creates a new OpaProfilesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpaProfilesUpdateParams() *OpaProfilesUpdateParams {
	return &OpaProfilesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpaProfilesUpdateParamsWithTimeout creates a new OpaProfilesUpdateParams object
// with the ability to set a timeout on a request.
func NewOpaProfilesUpdateParamsWithTimeout(timeout time.Duration) *OpaProfilesUpdateParams {
	return &OpaProfilesUpdateParams{
		timeout: timeout,
	}
}

// NewOpaProfilesUpdateParamsWithContext creates a new OpaProfilesUpdateParams object
// with the ability to set a context for a request.
func NewOpaProfilesUpdateParamsWithContext(ctx context.Context) *OpaProfilesUpdateParams {
	return &OpaProfilesUpdateParams{
		Context: ctx,
	}
}

// NewOpaProfilesUpdateParamsWithHTTPClient creates a new OpaProfilesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpaProfilesUpdateParamsWithHTTPClient(client *http.Client) *OpaProfilesUpdateParams {
	return &OpaProfilesUpdateParams{
		HTTPClient: client,
	}
}

/*
OpaProfilesUpdateParams contains all the parameters to send to the API endpoint

	for the opa profiles update operation.

	Typically these are written to a http.Request.
*/
type OpaProfilesUpdateParams struct {

	// Body.
	Body *models.OpaProfileUpdateCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the opa profiles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesUpdateParams) WithDefaults() *OpaProfilesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the opa profiles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the opa profiles update params
func (o *OpaProfilesUpdateParams) WithTimeout(timeout time.Duration) *OpaProfilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the opa profiles update params
func (o *OpaProfilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the opa profiles update params
func (o *OpaProfilesUpdateParams) WithContext(ctx context.Context) *OpaProfilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the opa profiles update params
func (o *OpaProfilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the opa profiles update params
func (o *OpaProfilesUpdateParams) WithHTTPClient(client *http.Client) *OpaProfilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the opa profiles update params
func (o *OpaProfilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the opa profiles update params
func (o *OpaProfilesUpdateParams) WithBody(body *models.OpaProfileUpdateCommand) *OpaProfilesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the opa profiles update params
func (o *OpaProfilesUpdateParams) SetBody(body *models.OpaProfileUpdateCommand) {
	o.Body = body
}

// WithV adds the v to the opa profiles update params
func (o *OpaProfilesUpdateParams) WithV(v string) *OpaProfilesUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the opa profiles update params
func (o *OpaProfilesUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpaProfilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
