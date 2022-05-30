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

// NewStandAloneActionsStartParams creates a new StandAloneActionsStartParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneActionsStartParams() *StandAloneActionsStartParams {
	return &StandAloneActionsStartParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneActionsStartParamsWithTimeout creates a new StandAloneActionsStartParams object
// with the ability to set a timeout on a request.
func NewStandAloneActionsStartParamsWithTimeout(timeout time.Duration) *StandAloneActionsStartParams {
	return &StandAloneActionsStartParams{
		timeout: timeout,
	}
}

// NewStandAloneActionsStartParamsWithContext creates a new StandAloneActionsStartParams object
// with the ability to set a context for a request.
func NewStandAloneActionsStartParamsWithContext(ctx context.Context) *StandAloneActionsStartParams {
	return &StandAloneActionsStartParams{
		Context: ctx,
	}
}

// NewStandAloneActionsStartParamsWithHTTPClient creates a new StandAloneActionsStartParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneActionsStartParamsWithHTTPClient(client *http.Client) *StandAloneActionsStartParams {
	return &StandAloneActionsStartParams{
		HTTPClient: client,
	}
}

/* StandAloneActionsStartParams contains all the parameters to send to the API endpoint
   for the stand alone actions start operation.

   Typically these are written to a http.Request.
*/
type StandAloneActionsStartParams struct {

	// Body.
	Body *models.StartStandaloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone actions start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsStartParams) WithDefaults() *StandAloneActionsStartParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone actions start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsStartParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone actions start params
func (o *StandAloneActionsStartParams) WithTimeout(timeout time.Duration) *StandAloneActionsStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone actions start params
func (o *StandAloneActionsStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone actions start params
func (o *StandAloneActionsStartParams) WithContext(ctx context.Context) *StandAloneActionsStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone actions start params
func (o *StandAloneActionsStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone actions start params
func (o *StandAloneActionsStartParams) WithHTTPClient(client *http.Client) *StandAloneActionsStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone actions start params
func (o *StandAloneActionsStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone actions start params
func (o *StandAloneActionsStartParams) WithBody(body *models.StartStandaloneVMCommand) *StandAloneActionsStartParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone actions start params
func (o *StandAloneActionsStartParams) SetBody(body *models.StartStandaloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone actions start params
func (o *StandAloneActionsStartParams) WithV(v string) *StandAloneActionsStartParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone actions start params
func (o *StandAloneActionsStartParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneActionsStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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