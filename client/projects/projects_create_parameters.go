// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsCreateParams creates a new ProjectsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsCreateParams() *ProjectsCreateParams {
	return &ProjectsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsCreateParamsWithTimeout creates a new ProjectsCreateParams object
// with the ability to set a timeout on a request.
func NewProjectsCreateParamsWithTimeout(timeout time.Duration) *ProjectsCreateParams {
	return &ProjectsCreateParams{
		timeout: timeout,
	}
}

// NewProjectsCreateParamsWithContext creates a new ProjectsCreateParams object
// with the ability to set a context for a request.
func NewProjectsCreateParamsWithContext(ctx context.Context) *ProjectsCreateParams {
	return &ProjectsCreateParams{
		Context: ctx,
	}
}

// NewProjectsCreateParamsWithHTTPClient creates a new ProjectsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsCreateParamsWithHTTPClient(client *http.Client) *ProjectsCreateParams {
	return &ProjectsCreateParams{
		HTTPClient: client,
	}
}

/*
ProjectsCreateParams contains all the parameters to send to the API endpoint

	for the projects create operation.

	Typically these are written to a http.Request.
*/
type ProjectsCreateParams struct {

	/* Body.

	   Create command
	*/
	Body ProjectsCreateBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsCreateParams) WithDefaults() *ProjectsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects create params
func (o *ProjectsCreateParams) WithTimeout(timeout time.Duration) *ProjectsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects create params
func (o *ProjectsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects create params
func (o *ProjectsCreateParams) WithContext(ctx context.Context) *ProjectsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects create params
func (o *ProjectsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects create params
func (o *ProjectsCreateParams) WithHTTPClient(client *http.Client) *ProjectsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects create params
func (o *ProjectsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects create params
func (o *ProjectsCreateParams) WithBody(body ProjectsCreateBody) *ProjectsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects create params
func (o *ProjectsCreateParams) SetBody(body ProjectsCreateBody) {
	o.Body = body
}

// WithV adds the v to the projects create params
func (o *ProjectsCreateParams) WithV(v string) *ProjectsCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects create params
func (o *ProjectsCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
