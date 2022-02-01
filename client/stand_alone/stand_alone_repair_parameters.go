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

// NewStandAloneRepairParams creates a new StandAloneRepairParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneRepairParams() *StandAloneRepairParams {
	return &StandAloneRepairParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneRepairParamsWithTimeout creates a new StandAloneRepairParams object
// with the ability to set a timeout on a request.
func NewStandAloneRepairParamsWithTimeout(timeout time.Duration) *StandAloneRepairParams {
	return &StandAloneRepairParams{
		timeout: timeout,
	}
}

// NewStandAloneRepairParamsWithContext creates a new StandAloneRepairParams object
// with the ability to set a context for a request.
func NewStandAloneRepairParamsWithContext(ctx context.Context) *StandAloneRepairParams {
	return &StandAloneRepairParams{
		Context: ctx,
	}
}

// NewStandAloneRepairParamsWithHTTPClient creates a new StandAloneRepairParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneRepairParamsWithHTTPClient(client *http.Client) *StandAloneRepairParams {
	return &StandAloneRepairParams{
		HTTPClient: client,
	}
}

/* StandAloneRepairParams contains all the parameters to send to the API endpoint
   for the stand alone repair operation.

   Typically these are written to a http.Request.
*/
type StandAloneRepairParams struct {

	// Body.
	Body *models.RepairStandAloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone repair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneRepairParams) WithDefaults() *StandAloneRepairParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone repair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneRepairParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone repair params
func (o *StandAloneRepairParams) WithTimeout(timeout time.Duration) *StandAloneRepairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone repair params
func (o *StandAloneRepairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone repair params
func (o *StandAloneRepairParams) WithContext(ctx context.Context) *StandAloneRepairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone repair params
func (o *StandAloneRepairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone repair params
func (o *StandAloneRepairParams) WithHTTPClient(client *http.Client) *StandAloneRepairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone repair params
func (o *StandAloneRepairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone repair params
func (o *StandAloneRepairParams) WithBody(body *models.RepairStandAloneVMCommand) *StandAloneRepairParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone repair params
func (o *StandAloneRepairParams) SetBody(body *models.RepairStandAloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone repair params
func (o *StandAloneRepairParams) WithV(v string) *StandAloneRepairParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone repair params
func (o *StandAloneRepairParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneRepairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
