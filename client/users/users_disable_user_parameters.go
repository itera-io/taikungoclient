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

	"github.com/itera-io/taikungoclient/models"
)

// NewUsersDisableUserParams creates a new UsersDisableUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersDisableUserParams() *UsersDisableUserParams {
	return &UsersDisableUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersDisableUserParamsWithTimeout creates a new UsersDisableUserParams object
// with the ability to set a timeout on a request.
func NewUsersDisableUserParamsWithTimeout(timeout time.Duration) *UsersDisableUserParams {
	return &UsersDisableUserParams{
		timeout: timeout,
	}
}

// NewUsersDisableUserParamsWithContext creates a new UsersDisableUserParams object
// with the ability to set a context for a request.
func NewUsersDisableUserParamsWithContext(ctx context.Context) *UsersDisableUserParams {
	return &UsersDisableUserParams{
		Context: ctx,
	}
}

// NewUsersDisableUserParamsWithHTTPClient creates a new UsersDisableUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersDisableUserParamsWithHTTPClient(client *http.Client) *UsersDisableUserParams {
	return &UsersDisableUserParams{
		HTTPClient: client,
	}
}

/*
UsersDisableUserParams contains all the parameters to send to the API endpoint

	for the users disable user operation.

	Typically these are written to a http.Request.
*/
type UsersDisableUserParams struct {

	/* Body.

	   Command
	*/
	Body *models.DisableUserCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users disable user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersDisableUserParams) WithDefaults() *UsersDisableUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users disable user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersDisableUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users disable user params
func (o *UsersDisableUserParams) WithTimeout(timeout time.Duration) *UsersDisableUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users disable user params
func (o *UsersDisableUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users disable user params
func (o *UsersDisableUserParams) WithContext(ctx context.Context) *UsersDisableUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users disable user params
func (o *UsersDisableUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users disable user params
func (o *UsersDisableUserParams) WithHTTPClient(client *http.Client) *UsersDisableUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users disable user params
func (o *UsersDisableUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the users disable user params
func (o *UsersDisableUserParams) WithBody(body *models.DisableUserCommand) *UsersDisableUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the users disable user params
func (o *UsersDisableUserParams) SetBody(body *models.DisableUserCommand) {
	o.Body = body
}

// WithV adds the v to the users disable user params
func (o *UsersDisableUserParams) WithV(v string) *UsersDisableUserParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the users disable user params
func (o *UsersDisableUserParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UsersDisableUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
