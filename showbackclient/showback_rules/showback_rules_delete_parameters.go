// Code generated by go-swagger; DO NOT EDIT.

package showback_rules

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

// NewShowbackRulesDeleteParams creates a new ShowbackRulesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackRulesDeleteParams() *ShowbackRulesDeleteParams {
	return &ShowbackRulesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackRulesDeleteParamsWithTimeout creates a new ShowbackRulesDeleteParams object
// with the ability to set a timeout on a request.
func NewShowbackRulesDeleteParamsWithTimeout(timeout time.Duration) *ShowbackRulesDeleteParams {
	return &ShowbackRulesDeleteParams{
		timeout: timeout,
	}
}

// NewShowbackRulesDeleteParamsWithContext creates a new ShowbackRulesDeleteParams object
// with the ability to set a context for a request.
func NewShowbackRulesDeleteParamsWithContext(ctx context.Context) *ShowbackRulesDeleteParams {
	return &ShowbackRulesDeleteParams{
		Context: ctx,
	}
}

// NewShowbackRulesDeleteParamsWithHTTPClient creates a new ShowbackRulesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackRulesDeleteParamsWithHTTPClient(client *http.Client) *ShowbackRulesDeleteParams {
	return &ShowbackRulesDeleteParams{
		HTTPClient: client,
	}
}

/* ShowbackRulesDeleteParams contains all the parameters to send to the API endpoint
   for the showback rules delete operation.

   Typically these are written to a http.Request.
*/
type ShowbackRulesDeleteParams struct {

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

// WithDefaults hydrates default values in the showback rules delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesDeleteParams) WithDefaults() *ShowbackRulesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback rules delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback rules delete params
func (o *ShowbackRulesDeleteParams) WithTimeout(timeout time.Duration) *ShowbackRulesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback rules delete params
func (o *ShowbackRulesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback rules delete params
func (o *ShowbackRulesDeleteParams) WithContext(ctx context.Context) *ShowbackRulesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback rules delete params
func (o *ShowbackRulesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback rules delete params
func (o *ShowbackRulesDeleteParams) WithHTTPClient(client *http.Client) *ShowbackRulesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback rules delete params
func (o *ShowbackRulesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the showback rules delete params
func (o *ShowbackRulesDeleteParams) WithID(id int32) *ShowbackRulesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the showback rules delete params
func (o *ShowbackRulesDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the showback rules delete params
func (o *ShowbackRulesDeleteParams) WithV(v string) *ShowbackRulesDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback rules delete params
func (o *ShowbackRulesDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackRulesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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