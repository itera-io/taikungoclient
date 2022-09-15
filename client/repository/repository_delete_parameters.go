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

	"github.com/itera-io/taikungoclient/models"
)

// NewRepositoryDeleteParams creates a new RepositoryDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryDeleteParams() *RepositoryDeleteParams {
	return &RepositoryDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryDeleteParamsWithTimeout creates a new RepositoryDeleteParams object
// with the ability to set a timeout on a request.
func NewRepositoryDeleteParamsWithTimeout(timeout time.Duration) *RepositoryDeleteParams {
	return &RepositoryDeleteParams{
		timeout: timeout,
	}
}

// NewRepositoryDeleteParamsWithContext creates a new RepositoryDeleteParams object
// with the ability to set a context for a request.
func NewRepositoryDeleteParamsWithContext(ctx context.Context) *RepositoryDeleteParams {
	return &RepositoryDeleteParams{
		Context: ctx,
	}
}

// NewRepositoryDeleteParamsWithHTTPClient creates a new RepositoryDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryDeleteParamsWithHTTPClient(client *http.Client) *RepositoryDeleteParams {
	return &RepositoryDeleteParams{
		HTTPClient: client,
	}
}

/*
RepositoryDeleteParams contains all the parameters to send to the API endpoint

	for the repository delete operation.

	Typically these are written to a http.Request.
*/
type RepositoryDeleteParams struct {

	// Body.
	Body *models.UnbindAppRepositoryCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repository delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryDeleteParams) WithDefaults() *RepositoryDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository delete params
func (o *RepositoryDeleteParams) WithTimeout(timeout time.Duration) *RepositoryDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository delete params
func (o *RepositoryDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository delete params
func (o *RepositoryDeleteParams) WithContext(ctx context.Context) *RepositoryDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository delete params
func (o *RepositoryDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository delete params
func (o *RepositoryDeleteParams) WithHTTPClient(client *http.Client) *RepositoryDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository delete params
func (o *RepositoryDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the repository delete params
func (o *RepositoryDeleteParams) WithBody(body *models.UnbindAppRepositoryCommand) *RepositoryDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the repository delete params
func (o *RepositoryDeleteParams) SetBody(body *models.UnbindAppRepositoryCommand) {
	o.Body = body
}

// WithV adds the v to the repository delete params
func (o *RepositoryDeleteParams) WithV(v string) *RepositoryDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the repository delete params
func (o *RepositoryDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
