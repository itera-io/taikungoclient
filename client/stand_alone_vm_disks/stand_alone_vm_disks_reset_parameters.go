// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_vm_disks

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

// NewStandAloneVMDisksResetParams creates a new StandAloneVMDisksResetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneVMDisksResetParams() *StandAloneVMDisksResetParams {
	return &StandAloneVMDisksResetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneVMDisksResetParamsWithTimeout creates a new StandAloneVMDisksResetParams object
// with the ability to set a timeout on a request.
func NewStandAloneVMDisksResetParamsWithTimeout(timeout time.Duration) *StandAloneVMDisksResetParams {
	return &StandAloneVMDisksResetParams{
		timeout: timeout,
	}
}

// NewStandAloneVMDisksResetParamsWithContext creates a new StandAloneVMDisksResetParams object
// with the ability to set a context for a request.
func NewStandAloneVMDisksResetParamsWithContext(ctx context.Context) *StandAloneVMDisksResetParams {
	return &StandAloneVMDisksResetParams{
		Context: ctx,
	}
}

// NewStandAloneVMDisksResetParamsWithHTTPClient creates a new StandAloneVMDisksResetParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneVMDisksResetParamsWithHTTPClient(client *http.Client) *StandAloneVMDisksResetParams {
	return &StandAloneVMDisksResetParams{
		HTTPClient: client,
	}
}

/* StandAloneVMDisksResetParams contains all the parameters to send to the API endpoint
   for the stand alone Vm disks reset operation.

   Typically these are written to a http.Request.
*/
type StandAloneVMDisksResetParams struct {

	// Body.
	Body *models.ResetStandAloneVMDiskStatusCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone Vm disks reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneVMDisksResetParams) WithDefaults() *StandAloneVMDisksResetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone Vm disks reset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneVMDisksResetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) WithTimeout(timeout time.Duration) *StandAloneVMDisksResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) WithContext(ctx context.Context) *StandAloneVMDisksResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) WithHTTPClient(client *http.Client) *StandAloneVMDisksResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) WithBody(body *models.ResetStandAloneVMDiskStatusCommand) *StandAloneVMDisksResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) SetBody(body *models.ResetStandAloneVMDiskStatusCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) WithV(v string) *StandAloneVMDisksResetParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone Vm disks reset params
func (o *StandAloneVMDisksResetParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneVMDisksResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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