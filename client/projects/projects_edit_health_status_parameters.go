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
	"github.com/go-openapi/swag"
)

// NewProjectsEditHealthStatusParams creates a new ProjectsEditHealthStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsEditHealthStatusParams() *ProjectsEditHealthStatusParams {
	return &ProjectsEditHealthStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsEditHealthStatusParamsWithTimeout creates a new ProjectsEditHealthStatusParams object
// with the ability to set a timeout on a request.
func NewProjectsEditHealthStatusParamsWithTimeout(timeout time.Duration) *ProjectsEditHealthStatusParams {
	return &ProjectsEditHealthStatusParams{
		timeout: timeout,
	}
}

// NewProjectsEditHealthStatusParamsWithContext creates a new ProjectsEditHealthStatusParams object
// with the ability to set a context for a request.
func NewProjectsEditHealthStatusParamsWithContext(ctx context.Context) *ProjectsEditHealthStatusParams {
	return &ProjectsEditHealthStatusParams{
		Context: ctx,
	}
}

// NewProjectsEditHealthStatusParamsWithHTTPClient creates a new ProjectsEditHealthStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsEditHealthStatusParamsWithHTTPClient(client *http.Client) *ProjectsEditHealthStatusParams {
	return &ProjectsEditHealthStatusParams{
		HTTPClient: client,
	}
}

/*
ProjectsEditHealthStatusParams contains all the parameters to send to the API endpoint

	for the projects edit health status operation.

	Typically these are written to a http.Request.
*/
type ProjectsEditHealthStatusParams struct {

	// Body.
	//
	// Format: int32
	Body int32

	// ProjectID.
	//
	// Format: int32
	ProjectID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects edit health status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsEditHealthStatusParams) WithDefaults() *ProjectsEditHealthStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects edit health status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsEditHealthStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithTimeout(timeout time.Duration) *ProjectsEditHealthStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithContext(ctx context.Context) *ProjectsEditHealthStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithHTTPClient(client *http.Client) *ProjectsEditHealthStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithBody(body int32) *ProjectsEditHealthStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetBody(body int32) {
	o.Body = body
}

// WithProjectID adds the projectID to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithProjectID(projectID int32) *ProjectsEditHealthStatusParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) WithV(v string) *ProjectsEditHealthStatusParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects edit health status params
func (o *ProjectsEditHealthStatusParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsEditHealthStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", swag.FormatInt32(o.ProjectID)); err != nil {
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
