// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

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

// NewShowbackCredentialsDeleteParams creates a new ShowbackCredentialsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackCredentialsDeleteParams() *ShowbackCredentialsDeleteParams {
	return &ShowbackCredentialsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackCredentialsDeleteParamsWithTimeout creates a new ShowbackCredentialsDeleteParams object
// with the ability to set a timeout on a request.
func NewShowbackCredentialsDeleteParamsWithTimeout(timeout time.Duration) *ShowbackCredentialsDeleteParams {
	return &ShowbackCredentialsDeleteParams{
		timeout: timeout,
	}
}

// NewShowbackCredentialsDeleteParamsWithContext creates a new ShowbackCredentialsDeleteParams object
// with the ability to set a context for a request.
func NewShowbackCredentialsDeleteParamsWithContext(ctx context.Context) *ShowbackCredentialsDeleteParams {
	return &ShowbackCredentialsDeleteParams{
		Context: ctx,
	}
}

// NewShowbackCredentialsDeleteParamsWithHTTPClient creates a new ShowbackCredentialsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackCredentialsDeleteParamsWithHTTPClient(client *http.Client) *ShowbackCredentialsDeleteParams {
	return &ShowbackCredentialsDeleteParams{
		HTTPClient: client,
	}
}

/*
ShowbackCredentialsDeleteParams contains all the parameters to send to the API endpoint

	for the showback credentials delete operation.

	Typically these are written to a http.Request.
*/
type ShowbackCredentialsDeleteParams struct {

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

// WithDefaults hydrates default values in the showback credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsDeleteParams) WithDefaults() *ShowbackCredentialsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) WithTimeout(timeout time.Duration) *ShowbackCredentialsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) WithContext(ctx context.Context) *ShowbackCredentialsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) WithHTTPClient(client *http.Client) *ShowbackCredentialsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) WithID(id int32) *ShowbackCredentialsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) WithV(v string) *ShowbackCredentialsDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback credentials delete params
func (o *ShowbackCredentialsDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackCredentialsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
