// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

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

// NewStandAloneDeleteParams creates a new StandAloneDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneDeleteParams() *StandAloneDeleteParams {
	return &StandAloneDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneDeleteParamsWithTimeout creates a new StandAloneDeleteParams object
// with the ability to set a timeout on a request.
func NewStandAloneDeleteParamsWithTimeout(timeout time.Duration) *StandAloneDeleteParams {
	return &StandAloneDeleteParams{
		timeout: timeout,
	}
}

// NewStandAloneDeleteParamsWithContext creates a new StandAloneDeleteParams object
// with the ability to set a context for a request.
func NewStandAloneDeleteParamsWithContext(ctx context.Context) *StandAloneDeleteParams {
	return &StandAloneDeleteParams{
		Context: ctx,
	}
}

// NewStandAloneDeleteParamsWithHTTPClient creates a new StandAloneDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneDeleteParamsWithHTTPClient(client *http.Client) *StandAloneDeleteParams {
	return &StandAloneDeleteParams{
		HTTPClient: client,
	}
}

/*
StandAloneDeleteParams contains all the parameters to send to the API endpoint

	for the stand alone delete operation.

	Typically these are written to a http.Request.
*/
type StandAloneDeleteParams struct {

	// Body.
	Body *models.DeleteStandAloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneDeleteParams) WithDefaults() *StandAloneDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone delete params
func (o *StandAloneDeleteParams) WithTimeout(timeout time.Duration) *StandAloneDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone delete params
func (o *StandAloneDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone delete params
func (o *StandAloneDeleteParams) WithContext(ctx context.Context) *StandAloneDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone delete params
func (o *StandAloneDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone delete params
func (o *StandAloneDeleteParams) WithHTTPClient(client *http.Client) *StandAloneDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone delete params
func (o *StandAloneDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone delete params
func (o *StandAloneDeleteParams) WithBody(body *models.DeleteStandAloneVMCommand) *StandAloneDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone delete params
func (o *StandAloneDeleteParams) SetBody(body *models.DeleteStandAloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone delete params
func (o *StandAloneDeleteParams) WithV(v string) *StandAloneDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone delete params
func (o *StandAloneDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
