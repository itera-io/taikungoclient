// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_profile

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

// NewStandAloneProfileDropdownListParams creates a new StandAloneProfileDropdownListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneProfileDropdownListParams() *StandAloneProfileDropdownListParams {
	return &StandAloneProfileDropdownListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneProfileDropdownListParamsWithTimeout creates a new StandAloneProfileDropdownListParams object
// with the ability to set a timeout on a request.
func NewStandAloneProfileDropdownListParamsWithTimeout(timeout time.Duration) *StandAloneProfileDropdownListParams {
	return &StandAloneProfileDropdownListParams{
		timeout: timeout,
	}
}

// NewStandAloneProfileDropdownListParamsWithContext creates a new StandAloneProfileDropdownListParams object
// with the ability to set a context for a request.
func NewStandAloneProfileDropdownListParamsWithContext(ctx context.Context) *StandAloneProfileDropdownListParams {
	return &StandAloneProfileDropdownListParams{
		Context: ctx,
	}
}

// NewStandAloneProfileDropdownListParamsWithHTTPClient creates a new StandAloneProfileDropdownListParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneProfileDropdownListParamsWithHTTPClient(client *http.Client) *StandAloneProfileDropdownListParams {
	return &StandAloneProfileDropdownListParams{
		HTTPClient: client,
	}
}

/* StandAloneProfileDropdownListParams contains all the parameters to send to the API endpoint
   for the stand alone profile dropdown list operation.

   Typically these are written to a http.Request.
*/
type StandAloneProfileDropdownListParams struct {

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone profile dropdown list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneProfileDropdownListParams) WithDefaults() *StandAloneProfileDropdownListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone profile dropdown list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneProfileDropdownListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) WithTimeout(timeout time.Duration) *StandAloneProfileDropdownListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) WithContext(ctx context.Context) *StandAloneProfileDropdownListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) WithHTTPClient(client *http.Client) *StandAloneProfileDropdownListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV adds the v to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) WithV(v string) *StandAloneProfileDropdownListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone profile dropdown list params
func (o *StandAloneProfileDropdownListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneProfileDropdownListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
