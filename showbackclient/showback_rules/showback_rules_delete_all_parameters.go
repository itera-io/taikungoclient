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
)

// NewShowbackRulesDeleteAllParams creates a new ShowbackRulesDeleteAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackRulesDeleteAllParams() *ShowbackRulesDeleteAllParams {
	return &ShowbackRulesDeleteAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackRulesDeleteAllParamsWithTimeout creates a new ShowbackRulesDeleteAllParams object
// with the ability to set a timeout on a request.
func NewShowbackRulesDeleteAllParamsWithTimeout(timeout time.Duration) *ShowbackRulesDeleteAllParams {
	return &ShowbackRulesDeleteAllParams{
		timeout: timeout,
	}
}

// NewShowbackRulesDeleteAllParamsWithContext creates a new ShowbackRulesDeleteAllParams object
// with the ability to set a context for a request.
func NewShowbackRulesDeleteAllParamsWithContext(ctx context.Context) *ShowbackRulesDeleteAllParams {
	return &ShowbackRulesDeleteAllParams{
		Context: ctx,
	}
}

// NewShowbackRulesDeleteAllParamsWithHTTPClient creates a new ShowbackRulesDeleteAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackRulesDeleteAllParamsWithHTTPClient(client *http.Client) *ShowbackRulesDeleteAllParams {
	return &ShowbackRulesDeleteAllParams{
		HTTPClient: client,
	}
}

/*
ShowbackRulesDeleteAllParams contains all the parameters to send to the API endpoint

	for the showback rules delete all operation.

	Typically these are written to a http.Request.
*/
type ShowbackRulesDeleteAllParams struct {

	// Body.
	Body ShowbackRulesDeleteAllBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback rules delete all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesDeleteAllParams) WithDefaults() *ShowbackRulesDeleteAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback rules delete all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesDeleteAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) WithTimeout(timeout time.Duration) *ShowbackRulesDeleteAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) WithContext(ctx context.Context) *ShowbackRulesDeleteAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) WithHTTPClient(client *http.Client) *ShowbackRulesDeleteAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) WithBody(body ShowbackRulesDeleteAllBody) *ShowbackRulesDeleteAllParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) SetBody(body ShowbackRulesDeleteAllBody) {
	o.Body = body
}

// WithV adds the v to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) WithV(v string) *ShowbackRulesDeleteAllParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback rules delete all params
func (o *ShowbackRulesDeleteAllParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackRulesDeleteAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
