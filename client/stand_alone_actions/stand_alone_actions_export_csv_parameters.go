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
	"github.com/go-openapi/swag"
)

// NewStandAloneActionsExportCsvParams creates a new StandAloneActionsExportCsvParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneActionsExportCsvParams() *StandAloneActionsExportCsvParams {
	return &StandAloneActionsExportCsvParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneActionsExportCsvParamsWithTimeout creates a new StandAloneActionsExportCsvParams object
// with the ability to set a timeout on a request.
func NewStandAloneActionsExportCsvParamsWithTimeout(timeout time.Duration) *StandAloneActionsExportCsvParams {
	return &StandAloneActionsExportCsvParams{
		timeout: timeout,
	}
}

// NewStandAloneActionsExportCsvParamsWithContext creates a new StandAloneActionsExportCsvParams object
// with the ability to set a context for a request.
func NewStandAloneActionsExportCsvParamsWithContext(ctx context.Context) *StandAloneActionsExportCsvParams {
	return &StandAloneActionsExportCsvParams{
		Context: ctx,
	}
}

// NewStandAloneActionsExportCsvParamsWithHTTPClient creates a new StandAloneActionsExportCsvParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneActionsExportCsvParamsWithHTTPClient(client *http.Client) *StandAloneActionsExportCsvParams {
	return &StandAloneActionsExportCsvParams{
		HTTPClient: client,
	}
}

/* StandAloneActionsExportCsvParams contains all the parameters to send to the API endpoint
   for the stand alone actions export csv operation.

   Typically these are written to a http.Request.
*/
type StandAloneActionsExportCsvParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone actions export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsExportCsvParams) WithDefaults() *StandAloneActionsExportCsvParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone actions export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsExportCsvParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) WithTimeout(timeout time.Duration) *StandAloneActionsExportCsvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) WithContext(ctx context.Context) *StandAloneActionsExportCsvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) WithHTTPClient(client *http.Client) *StandAloneActionsExportCsvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) WithID(id int32) *StandAloneActionsExportCsvParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) WithV(v string) *StandAloneActionsExportCsvParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone actions export csv params
func (o *StandAloneActionsExportCsvParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneActionsExportCsvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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