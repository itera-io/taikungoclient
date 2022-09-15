// Code generated by go-swagger; DO NOT EDIT.

package ntp_servers

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

// NewNtpServersDeleteParams creates a new NtpServersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNtpServersDeleteParams() *NtpServersDeleteParams {
	return &NtpServersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNtpServersDeleteParamsWithTimeout creates a new NtpServersDeleteParams object
// with the ability to set a timeout on a request.
func NewNtpServersDeleteParamsWithTimeout(timeout time.Duration) *NtpServersDeleteParams {
	return &NtpServersDeleteParams{
		timeout: timeout,
	}
}

// NewNtpServersDeleteParamsWithContext creates a new NtpServersDeleteParams object
// with the ability to set a context for a request.
func NewNtpServersDeleteParamsWithContext(ctx context.Context) *NtpServersDeleteParams {
	return &NtpServersDeleteParams{
		Context: ctx,
	}
}

// NewNtpServersDeleteParamsWithHTTPClient creates a new NtpServersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNtpServersDeleteParamsWithHTTPClient(client *http.Client) *NtpServersDeleteParams {
	return &NtpServersDeleteParams{
		HTTPClient: client,
	}
}

/*
NtpServersDeleteParams contains all the parameters to send to the API endpoint

	for the ntp servers delete operation.

	Typically these are written to a http.Request.
*/
type NtpServersDeleteParams struct {

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

// WithDefaults hydrates default values in the ntp servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NtpServersDeleteParams) WithDefaults() *NtpServersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ntp servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NtpServersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ntp servers delete params
func (o *NtpServersDeleteParams) WithTimeout(timeout time.Duration) *NtpServersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ntp servers delete params
func (o *NtpServersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ntp servers delete params
func (o *NtpServersDeleteParams) WithContext(ctx context.Context) *NtpServersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ntp servers delete params
func (o *NtpServersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ntp servers delete params
func (o *NtpServersDeleteParams) WithHTTPClient(client *http.Client) *NtpServersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ntp servers delete params
func (o *NtpServersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ntp servers delete params
func (o *NtpServersDeleteParams) WithID(id int32) *NtpServersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ntp servers delete params
func (o *NtpServersDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the ntp servers delete params
func (o *NtpServersDeleteParams) WithV(v string) *NtpServersDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ntp servers delete params
func (o *NtpServersDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *NtpServersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
