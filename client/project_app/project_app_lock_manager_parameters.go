// Code generated by go-swagger; DO NOT EDIT.

package project_app

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

// NewProjectAppLockManagerParams creates a new ProjectAppLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectAppLockManagerParams() *ProjectAppLockManagerParams {
	return &ProjectAppLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectAppLockManagerParamsWithTimeout creates a new ProjectAppLockManagerParams object
// with the ability to set a timeout on a request.
func NewProjectAppLockManagerParamsWithTimeout(timeout time.Duration) *ProjectAppLockManagerParams {
	return &ProjectAppLockManagerParams{
		timeout: timeout,
	}
}

// NewProjectAppLockManagerParamsWithContext creates a new ProjectAppLockManagerParams object
// with the ability to set a context for a request.
func NewProjectAppLockManagerParamsWithContext(ctx context.Context) *ProjectAppLockManagerParams {
	return &ProjectAppLockManagerParams{
		Context: ctx,
	}
}

// NewProjectAppLockManagerParamsWithHTTPClient creates a new ProjectAppLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectAppLockManagerParamsWithHTTPClient(client *http.Client) *ProjectAppLockManagerParams {
	return &ProjectAppLockManagerParams{
		HTTPClient: client,
	}
}

/*
ProjectAppLockManagerParams contains all the parameters to send to the API endpoint

	for the project app lock manager operation.

	Typically these are written to a http.Request.
*/
type ProjectAppLockManagerParams struct {

	// Body.
	Body ProjectAppLockManagerBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project app lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectAppLockManagerParams) WithDefaults() *ProjectAppLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project app lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectAppLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project app lock manager params
func (o *ProjectAppLockManagerParams) WithTimeout(timeout time.Duration) *ProjectAppLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project app lock manager params
func (o *ProjectAppLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project app lock manager params
func (o *ProjectAppLockManagerParams) WithContext(ctx context.Context) *ProjectAppLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project app lock manager params
func (o *ProjectAppLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project app lock manager params
func (o *ProjectAppLockManagerParams) WithHTTPClient(client *http.Client) *ProjectAppLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project app lock manager params
func (o *ProjectAppLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project app lock manager params
func (o *ProjectAppLockManagerParams) WithBody(body ProjectAppLockManagerBody) *ProjectAppLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project app lock manager params
func (o *ProjectAppLockManagerParams) SetBody(body ProjectAppLockManagerBody) {
	o.Body = body
}

// WithV adds the v to the project app lock manager params
func (o *ProjectAppLockManagerParams) WithV(v string) *ProjectAppLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project app lock manager params
func (o *ProjectAppLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectAppLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
