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
)

// NewStandAlonePurgeParams creates a new StandAlonePurgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAlonePurgeParams() *StandAlonePurgeParams {
	return &StandAlonePurgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAlonePurgeParamsWithTimeout creates a new StandAlonePurgeParams object
// with the ability to set a timeout on a request.
func NewStandAlonePurgeParamsWithTimeout(timeout time.Duration) *StandAlonePurgeParams {
	return &StandAlonePurgeParams{
		timeout: timeout,
	}
}

// NewStandAlonePurgeParamsWithContext creates a new StandAlonePurgeParams object
// with the ability to set a context for a request.
func NewStandAlonePurgeParamsWithContext(ctx context.Context) *StandAlonePurgeParams {
	return &StandAlonePurgeParams{
		Context: ctx,
	}
}

// NewStandAlonePurgeParamsWithHTTPClient creates a new StandAlonePurgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAlonePurgeParamsWithHTTPClient(client *http.Client) *StandAlonePurgeParams {
	return &StandAlonePurgeParams{
		HTTPClient: client,
	}
}

/*
StandAlonePurgeParams contains all the parameters to send to the API endpoint

	for the stand alone purge operation.

	Typically these are written to a http.Request.
*/
type StandAlonePurgeParams struct {

	// Body.
	Body StandAlonePurgeBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone purge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAlonePurgeParams) WithDefaults() *StandAlonePurgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone purge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAlonePurgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone purge params
func (o *StandAlonePurgeParams) WithTimeout(timeout time.Duration) *StandAlonePurgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone purge params
func (o *StandAlonePurgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone purge params
func (o *StandAlonePurgeParams) WithContext(ctx context.Context) *StandAlonePurgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone purge params
func (o *StandAlonePurgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone purge params
func (o *StandAlonePurgeParams) WithHTTPClient(client *http.Client) *StandAlonePurgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone purge params
func (o *StandAlonePurgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone purge params
func (o *StandAlonePurgeParams) WithBody(body StandAlonePurgeBody) *StandAlonePurgeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone purge params
func (o *StandAlonePurgeParams) SetBody(body StandAlonePurgeBody) {
	o.Body = body
}

// WithV adds the v to the stand alone purge params
func (o *StandAlonePurgeParams) WithV(v string) *StandAlonePurgeParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone purge params
func (o *StandAlonePurgeParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAlonePurgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
