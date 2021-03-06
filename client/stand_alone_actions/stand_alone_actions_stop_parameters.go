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

// NewStandAloneActionsStopParams creates a new StandAloneActionsStopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneActionsStopParams() *StandAloneActionsStopParams {
	return &StandAloneActionsStopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneActionsStopParamsWithTimeout creates a new StandAloneActionsStopParams object
// with the ability to set a timeout on a request.
func NewStandAloneActionsStopParamsWithTimeout(timeout time.Duration) *StandAloneActionsStopParams {
	return &StandAloneActionsStopParams{
		timeout: timeout,
	}
}

// NewStandAloneActionsStopParamsWithContext creates a new StandAloneActionsStopParams object
// with the ability to set a context for a request.
func NewStandAloneActionsStopParamsWithContext(ctx context.Context) *StandAloneActionsStopParams {
	return &StandAloneActionsStopParams{
		Context: ctx,
	}
}

// NewStandAloneActionsStopParamsWithHTTPClient creates a new StandAloneActionsStopParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneActionsStopParamsWithHTTPClient(client *http.Client) *StandAloneActionsStopParams {
	return &StandAloneActionsStopParams{
		HTTPClient: client,
	}
}

/* StandAloneActionsStopParams contains all the parameters to send to the API endpoint
   for the stand alone actions stop operation.

   Typically these are written to a http.Request.
*/
type StandAloneActionsStopParams struct {

	// Body.
	Body *models.StopStandaloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone actions stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsStopParams) WithDefaults() *StandAloneActionsStopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone actions stop params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsStopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone actions stop params
func (o *StandAloneActionsStopParams) WithTimeout(timeout time.Duration) *StandAloneActionsStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone actions stop params
func (o *StandAloneActionsStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone actions stop params
func (o *StandAloneActionsStopParams) WithContext(ctx context.Context) *StandAloneActionsStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone actions stop params
func (o *StandAloneActionsStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone actions stop params
func (o *StandAloneActionsStopParams) WithHTTPClient(client *http.Client) *StandAloneActionsStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone actions stop params
func (o *StandAloneActionsStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone actions stop params
func (o *StandAloneActionsStopParams) WithBody(body *models.StopStandaloneVMCommand) *StandAloneActionsStopParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone actions stop params
func (o *StandAloneActionsStopParams) SetBody(body *models.StopStandaloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone actions stop params
func (o *StandAloneActionsStopParams) WithV(v string) *StandAloneActionsStopParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone actions stop params
func (o *StandAloneActionsStopParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneActionsStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
