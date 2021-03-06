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

// NewStandAloneIPManagementParams creates a new StandAloneIPManagementParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneIPManagementParams() *StandAloneIPManagementParams {
	return &StandAloneIPManagementParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneIPManagementParamsWithTimeout creates a new StandAloneIPManagementParams object
// with the ability to set a timeout on a request.
func NewStandAloneIPManagementParamsWithTimeout(timeout time.Duration) *StandAloneIPManagementParams {
	return &StandAloneIPManagementParams{
		timeout: timeout,
	}
}

// NewStandAloneIPManagementParamsWithContext creates a new StandAloneIPManagementParams object
// with the ability to set a context for a request.
func NewStandAloneIPManagementParamsWithContext(ctx context.Context) *StandAloneIPManagementParams {
	return &StandAloneIPManagementParams{
		Context: ctx,
	}
}

// NewStandAloneIPManagementParamsWithHTTPClient creates a new StandAloneIPManagementParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneIPManagementParamsWithHTTPClient(client *http.Client) *StandAloneIPManagementParams {
	return &StandAloneIPManagementParams{
		HTTPClient: client,
	}
}

/* StandAloneIPManagementParams contains all the parameters to send to the API endpoint
   for the stand alone Ip management operation.

   Typically these are written to a http.Request.
*/
type StandAloneIPManagementParams struct {

	// Body.
	Body *models.StandAloneVMIPManagementCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone Ip management params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneIPManagementParams) WithDefaults() *StandAloneIPManagementParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone Ip management params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneIPManagementParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone Ip management params
func (o *StandAloneIPManagementParams) WithTimeout(timeout time.Duration) *StandAloneIPManagementParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone Ip management params
func (o *StandAloneIPManagementParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone Ip management params
func (o *StandAloneIPManagementParams) WithContext(ctx context.Context) *StandAloneIPManagementParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone Ip management params
func (o *StandAloneIPManagementParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone Ip management params
func (o *StandAloneIPManagementParams) WithHTTPClient(client *http.Client) *StandAloneIPManagementParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone Ip management params
func (o *StandAloneIPManagementParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone Ip management params
func (o *StandAloneIPManagementParams) WithBody(body *models.StandAloneVMIPManagementCommand) *StandAloneIPManagementParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone Ip management params
func (o *StandAloneIPManagementParams) SetBody(body *models.StandAloneVMIPManagementCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone Ip management params
func (o *StandAloneIPManagementParams) WithV(v string) *StandAloneIPManagementParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone Ip management params
func (o *StandAloneIPManagementParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneIPManagementParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
