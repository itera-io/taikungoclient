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

// NewProjectsRepairParams creates a new ProjectsRepairParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsRepairParams() *ProjectsRepairParams {
	return &ProjectsRepairParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsRepairParamsWithTimeout creates a new ProjectsRepairParams object
// with the ability to set a timeout on a request.
func NewProjectsRepairParamsWithTimeout(timeout time.Duration) *ProjectsRepairParams {
	return &ProjectsRepairParams{
		timeout: timeout,
	}
}

// NewProjectsRepairParamsWithContext creates a new ProjectsRepairParams object
// with the ability to set a context for a request.
func NewProjectsRepairParamsWithContext(ctx context.Context) *ProjectsRepairParams {
	return &ProjectsRepairParams{
		Context: ctx,
	}
}

// NewProjectsRepairParamsWithHTTPClient creates a new ProjectsRepairParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsRepairParamsWithHTTPClient(client *http.Client) *ProjectsRepairParams {
	return &ProjectsRepairParams{
		HTTPClient: client,
	}
}

/* ProjectsRepairParams contains all the parameters to send to the API endpoint
   for the projects repair operation.

   Typically these are written to a http.Request.
*/
type ProjectsRepairParams struct {

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

// WithDefaults hydrates default values in the projects repair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsRepairParams) WithDefaults() *ProjectsRepairParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects repair params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsRepairParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects repair params
func (o *ProjectsRepairParams) WithTimeout(timeout time.Duration) *ProjectsRepairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects repair params
func (o *ProjectsRepairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects repair params
func (o *ProjectsRepairParams) WithContext(ctx context.Context) *ProjectsRepairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects repair params
func (o *ProjectsRepairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects repair params
func (o *ProjectsRepairParams) WithHTTPClient(client *http.Client) *ProjectsRepairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects repair params
func (o *ProjectsRepairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the projects repair params
func (o *ProjectsRepairParams) WithProjectID(projectID int32) *ProjectsRepairParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the projects repair params
func (o *ProjectsRepairParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the projects repair params
func (o *ProjectsRepairParams) WithV(v string) *ProjectsRepairParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects repair params
func (o *ProjectsRepairParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsRepairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
