// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersChangePasswordParams creates a new UsersChangePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersChangePasswordParams() *UsersChangePasswordParams {
	return &UsersChangePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersChangePasswordParamsWithTimeout creates a new UsersChangePasswordParams object
// with the ability to set a timeout on a request.
func NewUsersChangePasswordParamsWithTimeout(timeout time.Duration) *UsersChangePasswordParams {
	return &UsersChangePasswordParams{
		timeout: timeout,
	}
}

// NewUsersChangePasswordParamsWithContext creates a new UsersChangePasswordParams object
// with the ability to set a context for a request.
func NewUsersChangePasswordParamsWithContext(ctx context.Context) *UsersChangePasswordParams {
	return &UsersChangePasswordParams{
		Context: ctx,
	}
}

// NewUsersChangePasswordParamsWithHTTPClient creates a new UsersChangePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersChangePasswordParamsWithHTTPClient(client *http.Client) *UsersChangePasswordParams {
	return &UsersChangePasswordParams{
		HTTPClient: client,
	}
}

/*
UsersChangePasswordParams contains all the parameters to send to the API endpoint

	for the users change password operation.

	Typically these are written to a http.Request.
*/
type UsersChangePasswordParams struct {

	// Body.
	Body UsersChangePasswordBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersChangePasswordParams) WithDefaults() *UsersChangePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users change password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersChangePasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users change password params
func (o *UsersChangePasswordParams) WithTimeout(timeout time.Duration) *UsersChangePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users change password params
func (o *UsersChangePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users change password params
func (o *UsersChangePasswordParams) WithContext(ctx context.Context) *UsersChangePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users change password params
func (o *UsersChangePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users change password params
func (o *UsersChangePasswordParams) WithHTTPClient(client *http.Client) *UsersChangePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users change password params
func (o *UsersChangePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users change password params
func (o *UsersChangePasswordParams) WithBody(body UsersChangePasswordBody) *UsersChangePasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users change password params
func (o *UsersChangePasswordParams) SetBody(body UsersChangePasswordBody) {
	o.Body = body
}

// WithV adds the v to the users change password params
func (o *UsersChangePasswordParams) WithV(v string) *UsersChangePasswordParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the users change password params
func (o *UsersChangePasswordParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UsersChangePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
