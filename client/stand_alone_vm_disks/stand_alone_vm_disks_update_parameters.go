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

// NewStandAloneVMDisksUpdateParams creates a new StandAloneVMDisksUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneVMDisksUpdateParams() *StandAloneVMDisksUpdateParams {
	return &StandAloneVMDisksUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneVMDisksUpdateParamsWithTimeout creates a new StandAloneVMDisksUpdateParams object
// with the ability to set a timeout on a request.
func NewStandAloneVMDisksUpdateParamsWithTimeout(timeout time.Duration) *StandAloneVMDisksUpdateParams {
	return &StandAloneVMDisksUpdateParams{
		timeout: timeout,
	}
}

// NewStandAloneVMDisksUpdateParamsWithContext creates a new StandAloneVMDisksUpdateParams object
// with the ability to set a context for a request.
func NewStandAloneVMDisksUpdateParamsWithContext(ctx context.Context) *StandAloneVMDisksUpdateParams {
	return &StandAloneVMDisksUpdateParams{
		Context: ctx,
	}
}

// NewStandAloneVMDisksUpdateParamsWithHTTPClient creates a new StandAloneVMDisksUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneVMDisksUpdateParamsWithHTTPClient(client *http.Client) *StandAloneVMDisksUpdateParams {
	return &StandAloneVMDisksUpdateParams{
		HTTPClient: client,
	}
}

/*
StandAloneVMDisksUpdateParams contains all the parameters to send to the API endpoint

	for the stand alone Vm disks update operation.

	Typically these are written to a http.Request.
*/
type StandAloneVMDisksUpdateParams struct {

	// Body.
	Body *models.UpdateStandaloneVMDiskCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone Vm disks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneVMDisksUpdateParams) WithDefaults() *StandAloneVMDisksUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone Vm disks update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneVMDisksUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) WithTimeout(timeout time.Duration) *StandAloneVMDisksUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) WithContext(ctx context.Context) *StandAloneVMDisksUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) WithHTTPClient(client *http.Client) *StandAloneVMDisksUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) WithBody(body *models.UpdateStandaloneVMDiskCommand) *StandAloneVMDisksUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) SetBody(body *models.UpdateStandaloneVMDiskCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) WithV(v string) *StandAloneVMDisksUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone Vm disks update params
func (o *StandAloneVMDisksUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneVMDisksUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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