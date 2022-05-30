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

// NewStandAloneCommitParams creates a new StandAloneCommitParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneCommitParams() *StandAloneCommitParams {
	return &StandAloneCommitParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneCommitParamsWithTimeout creates a new StandAloneCommitParams object
// with the ability to set a timeout on a request.
func NewStandAloneCommitParamsWithTimeout(timeout time.Duration) *StandAloneCommitParams {
	return &StandAloneCommitParams{
		timeout: timeout,
	}
}

// NewStandAloneCommitParamsWithContext creates a new StandAloneCommitParams object
// with the ability to set a context for a request.
func NewStandAloneCommitParamsWithContext(ctx context.Context) *StandAloneCommitParams {
	return &StandAloneCommitParams{
		Context: ctx,
	}
}

// NewStandAloneCommitParamsWithHTTPClient creates a new StandAloneCommitParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneCommitParamsWithHTTPClient(client *http.Client) *StandAloneCommitParams {
	return &StandAloneCommitParams{
		HTTPClient: client,
	}
}

/* StandAloneCommitParams contains all the parameters to send to the API endpoint
   for the stand alone commit operation.

   Typically these are written to a http.Request.
*/
type StandAloneCommitParams struct {

	// Body.
	Body *models.CommitStandAloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneCommitParams) WithDefaults() *StandAloneCommitParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone commit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneCommitParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone commit params
func (o *StandAloneCommitParams) WithTimeout(timeout time.Duration) *StandAloneCommitParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone commit params
func (o *StandAloneCommitParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone commit params
func (o *StandAloneCommitParams) WithContext(ctx context.Context) *StandAloneCommitParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone commit params
func (o *StandAloneCommitParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone commit params
func (o *StandAloneCommitParams) WithHTTPClient(client *http.Client) *StandAloneCommitParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone commit params
func (o *StandAloneCommitParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone commit params
func (o *StandAloneCommitParams) WithBody(body *models.CommitStandAloneVMCommand) *StandAloneCommitParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone commit params
func (o *StandAloneCommitParams) SetBody(body *models.CommitStandAloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone commit params
func (o *StandAloneCommitParams) WithV(v string) *StandAloneCommitParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone commit params
func (o *StandAloneCommitParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneCommitParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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