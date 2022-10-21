// Code generated by go-swagger; DO NOT EDIT.

package project_infracosts

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

// NewProjectInfracostsDetailsParams creates a new ProjectInfracostsDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectInfracostsDetailsParams() *ProjectInfracostsDetailsParams {
	return &ProjectInfracostsDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectInfracostsDetailsParamsWithTimeout creates a new ProjectInfracostsDetailsParams object
// with the ability to set a timeout on a request.
func NewProjectInfracostsDetailsParamsWithTimeout(timeout time.Duration) *ProjectInfracostsDetailsParams {
	return &ProjectInfracostsDetailsParams{
		timeout: timeout,
	}
}

// NewProjectInfracostsDetailsParamsWithContext creates a new ProjectInfracostsDetailsParams object
// with the ability to set a context for a request.
func NewProjectInfracostsDetailsParamsWithContext(ctx context.Context) *ProjectInfracostsDetailsParams {
	return &ProjectInfracostsDetailsParams{
		Context: ctx,
	}
}

// NewProjectInfracostsDetailsParamsWithHTTPClient creates a new ProjectInfracostsDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectInfracostsDetailsParamsWithHTTPClient(client *http.Client) *ProjectInfracostsDetailsParams {
	return &ProjectInfracostsDetailsParams{
		HTTPClient: client,
	}
}

/*
ProjectInfracostsDetailsParams contains all the parameters to send to the API endpoint

	for the project infracosts details operation.

	Typically these are written to a http.Request.
*/
type ProjectInfracostsDetailsParams struct {

	// ID.
	ID string

	/* ProjectID.

	   Project Id

	   Format: int32
	*/
	ProjectID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project infracosts details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectInfracostsDetailsParams) WithDefaults() *ProjectInfracostsDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project infracosts details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectInfracostsDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithTimeout(timeout time.Duration) *ProjectInfracostsDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithContext(ctx context.Context) *ProjectInfracostsDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithHTTPClient(client *http.Client) *ProjectInfracostsDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithID(id string) *ProjectInfracostsDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithProjectID(projectID *int32) *ProjectInfracostsDetailsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) WithV(v string) *ProjectInfracostsDetailsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project infracosts details params
func (o *ProjectInfracostsDetailsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectInfracostsDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
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
