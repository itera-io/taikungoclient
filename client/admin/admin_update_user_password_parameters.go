// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminUpdateUserPasswordParams creates a new AdminUpdateUserPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminUpdateUserPasswordParams() *AdminUpdateUserPasswordParams {
	return &AdminUpdateUserPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateUserPasswordParamsWithTimeout creates a new AdminUpdateUserPasswordParams object
// with the ability to set a timeout on a request.
func NewAdminUpdateUserPasswordParamsWithTimeout(timeout time.Duration) *AdminUpdateUserPasswordParams {
	return &AdminUpdateUserPasswordParams{
		timeout: timeout,
	}
}

// NewAdminUpdateUserPasswordParamsWithContext creates a new AdminUpdateUserPasswordParams object
// with the ability to set a context for a request.
func NewAdminUpdateUserPasswordParamsWithContext(ctx context.Context) *AdminUpdateUserPasswordParams {
	return &AdminUpdateUserPasswordParams{
		Context: ctx,
	}
}

// NewAdminUpdateUserPasswordParamsWithHTTPClient creates a new AdminUpdateUserPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminUpdateUserPasswordParamsWithHTTPClient(client *http.Client) *AdminUpdateUserPasswordParams {
	return &AdminUpdateUserPasswordParams{
		HTTPClient: client,
	}
}

/*
AdminUpdateUserPasswordParams contains all the parameters to send to the API endpoint

	for the admin update user password operation.

	Typically these are written to a http.Request.
*/
type AdminUpdateUserPasswordParams struct {

	// Body.
	Body *models.AdminUsersUpdatePasswordCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin update user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserPasswordParams) WithDefaults() *AdminUpdateUserPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin update user password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin update user password params
func (o *AdminUpdateUserPasswordParams) WithTimeout(timeout time.Duration) *AdminUpdateUserPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update user password params
func (o *AdminUpdateUserPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update user password params
func (o *AdminUpdateUserPasswordParams) WithContext(ctx context.Context) *AdminUpdateUserPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update user password params
func (o *AdminUpdateUserPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin update user password params
func (o *AdminUpdateUserPasswordParams) WithHTTPClient(client *http.Client) *AdminUpdateUserPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update user password params
func (o *AdminUpdateUserPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin update user password params
func (o *AdminUpdateUserPasswordParams) WithBody(body *models.AdminUsersUpdatePasswordCommand) *AdminUpdateUserPasswordParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update user password params
func (o *AdminUpdateUserPasswordParams) SetBody(body *models.AdminUsersUpdatePasswordCommand) {
	o.Body = body
}

// WithV adds the v to the admin update user password params
func (o *AdminUpdateUserPasswordParams) WithV(v string) *AdminUpdateUserPasswordParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the admin update user password params
func (o *AdminUpdateUserPasswordParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateUserPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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