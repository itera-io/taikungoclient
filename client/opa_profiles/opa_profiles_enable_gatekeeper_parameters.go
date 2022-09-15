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

// NewOpaProfilesEnableGatekeeperParams creates a new OpaProfilesEnableGatekeeperParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpaProfilesEnableGatekeeperParams() *OpaProfilesEnableGatekeeperParams {
	return &OpaProfilesEnableGatekeeperParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpaProfilesEnableGatekeeperParamsWithTimeout creates a new OpaProfilesEnableGatekeeperParams object
// with the ability to set a timeout on a request.
func NewOpaProfilesEnableGatekeeperParamsWithTimeout(timeout time.Duration) *OpaProfilesEnableGatekeeperParams {
	return &OpaProfilesEnableGatekeeperParams{
		timeout: timeout,
	}
}

// NewOpaProfilesEnableGatekeeperParamsWithContext creates a new OpaProfilesEnableGatekeeperParams object
// with the ability to set a context for a request.
func NewOpaProfilesEnableGatekeeperParamsWithContext(ctx context.Context) *OpaProfilesEnableGatekeeperParams {
	return &OpaProfilesEnableGatekeeperParams{
		Context: ctx,
	}
}

// NewOpaProfilesEnableGatekeeperParamsWithHTTPClient creates a new OpaProfilesEnableGatekeeperParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpaProfilesEnableGatekeeperParamsWithHTTPClient(client *http.Client) *OpaProfilesEnableGatekeeperParams {
	return &OpaProfilesEnableGatekeeperParams{
		HTTPClient: client,
	}
}

/*
OpaProfilesEnableGatekeeperParams contains all the parameters to send to the API endpoint

	for the opa profiles enable gatekeeper operation.

	Typically these are written to a http.Request.
*/
type OpaProfilesEnableGatekeeperParams struct {

	// Body.
	Body *models.EnableGatekeeperCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the opa profiles enable gatekeeper params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesEnableGatekeeperParams) WithDefaults() *OpaProfilesEnableGatekeeperParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the opa profiles enable gatekeeper params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpaProfilesEnableGatekeeperParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) WithTimeout(timeout time.Duration) *OpaProfilesEnableGatekeeperParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) WithContext(ctx context.Context) *OpaProfilesEnableGatekeeperParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) WithHTTPClient(client *http.Client) *OpaProfilesEnableGatekeeperParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) WithBody(body *models.EnableGatekeeperCommand) *OpaProfilesEnableGatekeeperParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) SetBody(body *models.EnableGatekeeperCommand) {
	o.Body = body
}

// WithV adds the v to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) WithV(v string) *OpaProfilesEnableGatekeeperParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the opa profiles enable gatekeeper params
func (o *OpaProfilesEnableGatekeeperParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpaProfilesEnableGatekeeperParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
