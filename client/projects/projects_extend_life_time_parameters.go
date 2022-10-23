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

// NewProjectsExtendLifeTimeParams creates a new ProjectsExtendLifeTimeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsExtendLifeTimeParams() *ProjectsExtendLifeTimeParams {
	return &ProjectsExtendLifeTimeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsExtendLifeTimeParamsWithTimeout creates a new ProjectsExtendLifeTimeParams object
// with the ability to set a timeout on a request.
func NewProjectsExtendLifeTimeParamsWithTimeout(timeout time.Duration) *ProjectsExtendLifeTimeParams {
	return &ProjectsExtendLifeTimeParams{
		timeout: timeout,
	}
}

// NewProjectsExtendLifeTimeParamsWithContext creates a new ProjectsExtendLifeTimeParams object
// with the ability to set a context for a request.
func NewProjectsExtendLifeTimeParamsWithContext(ctx context.Context) *ProjectsExtendLifeTimeParams {
	return &ProjectsExtendLifeTimeParams{
		Context: ctx,
	}
}

// NewProjectsExtendLifeTimeParamsWithHTTPClient creates a new ProjectsExtendLifeTimeParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsExtendLifeTimeParamsWithHTTPClient(client *http.Client) *ProjectsExtendLifeTimeParams {
	return &ProjectsExtendLifeTimeParams{
		HTTPClient: client,
	}
}

/*
ProjectsExtendLifeTimeParams contains all the parameters to send to the API endpoint

	for the projects extend life time operation.

	Typically these are written to a http.Request.
*/
type ProjectsExtendLifeTimeParams struct {

	/* Body.

	   Command
	*/
	Body ProjectsExtendLifeTimeBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects extend life time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsExtendLifeTimeParams) WithDefaults() *ProjectsExtendLifeTimeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects extend life time params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsExtendLifeTimeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) WithTimeout(timeout time.Duration) *ProjectsExtendLifeTimeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) WithContext(ctx context.Context) *ProjectsExtendLifeTimeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) WithHTTPClient(client *http.Client) *ProjectsExtendLifeTimeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) WithBody(body ProjectsExtendLifeTimeBody) *ProjectsExtendLifeTimeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) SetBody(body ProjectsExtendLifeTimeBody) {
	o.Body = body
}

// WithV adds the v to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) WithV(v string) *ProjectsExtendLifeTimeParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects extend life time params
func (o *ProjectsExtendLifeTimeParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsExtendLifeTimeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
