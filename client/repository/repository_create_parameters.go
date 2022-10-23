// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepositoryCreateParams creates a new RepositoryCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryCreateParams() *RepositoryCreateParams {
	return &RepositoryCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryCreateParamsWithTimeout creates a new RepositoryCreateParams object
// with the ability to set a timeout on a request.
func NewRepositoryCreateParamsWithTimeout(timeout time.Duration) *RepositoryCreateParams {
	return &RepositoryCreateParams{
		timeout: timeout,
	}
}

// NewRepositoryCreateParamsWithContext creates a new RepositoryCreateParams object
// with the ability to set a context for a request.
func NewRepositoryCreateParamsWithContext(ctx context.Context) *RepositoryCreateParams {
	return &RepositoryCreateParams{
		Context: ctx,
	}
}

// NewRepositoryCreateParamsWithHTTPClient creates a new RepositoryCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryCreateParamsWithHTTPClient(client *http.Client) *RepositoryCreateParams {
	return &RepositoryCreateParams{
		HTTPClient: client,
	}
}

/*
RepositoryCreateParams contains all the parameters to send to the API endpoint

	for the repository create operation.

	Typically these are written to a http.Request.
*/
type RepositoryCreateParams struct {

	// Body.
	Body RepositoryCreateBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryCreateParams) WithDefaults() *RepositoryCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository create params
func (o *RepositoryCreateParams) WithTimeout(timeout time.Duration) *RepositoryCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository create params
func (o *RepositoryCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository create params
func (o *RepositoryCreateParams) WithContext(ctx context.Context) *RepositoryCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository create params
func (o *RepositoryCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository create params
func (o *RepositoryCreateParams) WithHTTPClient(client *http.Client) *RepositoryCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository create params
func (o *RepositoryCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository create params
func (o *RepositoryCreateParams) WithBody(body RepositoryCreateBody) *RepositoryCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository create params
func (o *RepositoryCreateParams) SetBody(body RepositoryCreateBody) {
	o.Body = body
}

// WithV adds the v to the repository create params
func (o *RepositoryCreateParams) WithV(v string) *RepositoryCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the repository create params
func (o *RepositoryCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
