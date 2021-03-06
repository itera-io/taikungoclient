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

// NewStandAloneCreateParams creates a new StandAloneCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneCreateParams() *StandAloneCreateParams {
	return &StandAloneCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneCreateParamsWithTimeout creates a new StandAloneCreateParams object
// with the ability to set a timeout on a request.
func NewStandAloneCreateParamsWithTimeout(timeout time.Duration) *StandAloneCreateParams {
	return &StandAloneCreateParams{
		timeout: timeout,
	}
}

// NewStandAloneCreateParamsWithContext creates a new StandAloneCreateParams object
// with the ability to set a context for a request.
func NewStandAloneCreateParamsWithContext(ctx context.Context) *StandAloneCreateParams {
	return &StandAloneCreateParams{
		Context: ctx,
	}
}

// NewStandAloneCreateParamsWithHTTPClient creates a new StandAloneCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneCreateParamsWithHTTPClient(client *http.Client) *StandAloneCreateParams {
	return &StandAloneCreateParams{
		HTTPClient: client,
	}
}

/* StandAloneCreateParams contains all the parameters to send to the API endpoint
   for the stand alone create operation.

   Typically these are written to a http.Request.
*/
type StandAloneCreateParams struct {

	// Body.
	Body *models.CreateStandAloneVMCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneCreateParams) WithDefaults() *StandAloneCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone create params
func (o *StandAloneCreateParams) WithTimeout(timeout time.Duration) *StandAloneCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone create params
func (o *StandAloneCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone create params
func (o *StandAloneCreateParams) WithContext(ctx context.Context) *StandAloneCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone create params
func (o *StandAloneCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone create params
func (o *StandAloneCreateParams) WithHTTPClient(client *http.Client) *StandAloneCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone create params
func (o *StandAloneCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stand alone create params
func (o *StandAloneCreateParams) WithBody(body *models.CreateStandAloneVMCommand) *StandAloneCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stand alone create params
func (o *StandAloneCreateParams) SetBody(body *models.CreateStandAloneVMCommand) {
	o.Body = body
}

// WithV adds the v to the stand alone create params
func (o *StandAloneCreateParams) WithV(v string) *StandAloneCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone create params
func (o *StandAloneCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
