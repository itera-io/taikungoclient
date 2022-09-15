// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_actions

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

// NewStandAloneActionsShelveParams creates a new StandAloneActionsShelveParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneActionsShelveParams() *StandAloneActionsShelveParams {
	return &StandAloneActionsShelveParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneActionsShelveParamsWithTimeout creates a new StandAloneActionsShelveParams object
// with the ability to set a timeout on a request.
func NewStandAloneActionsShelveParamsWithTimeout(timeout time.Duration) *StandAloneActionsShelveParams {
	return &StandAloneActionsShelveParams{
		timeout: timeout,
	}
}

// NewStandAloneActionsShelveParamsWithContext creates a new StandAloneActionsShelveParams object
// with the ability to set a context for a request.
func NewStandAloneActionsShelveParamsWithContext(ctx context.Context) *StandAloneActionsShelveParams {
	return &StandAloneActionsShelveParams{
		Context: ctx,
	}
}

// NewStandAloneActionsShelveParamsWithHTTPClient creates a new StandAloneActionsShelveParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneActionsShelveParamsWithHTTPClient(client *http.Client) *StandAloneActionsShelveParams {
	return &StandAloneActionsShelveParams{
		HTTPClient: client,
	}
}

/*
StandAloneActionsShelveParams contains all the parameters to send to the API endpoint

	for the stand alone actions shelve operation.

	Typically these are written to a http.Request.
*/
type StandAloneActionsShelveParams struct {

	// Body.
	Body *models.ShelveStandAloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone actions shelve params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsShelveParams) WithDefaults() *StandAloneActionsShelveParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone actions shelve params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsShelveParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) WithTimeout(timeout time.Duration) *StandAloneActionsShelveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) WithContext(ctx context.Context) *StandAloneActionsShelveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) WithHTTPClient(client *http.Client) *StandAloneActionsShelveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) WithBody(body *models.ShelveStandAloneVMCommand) *StandAloneActionsShelveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) SetBody(body *models.ShelveStandAloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) WithV(v string) *StandAloneActionsShelveParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone actions shelve params
func (o *StandAloneActionsShelveParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneActionsShelveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
