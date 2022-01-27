// Code generated by go-swagger; DO NOT EDIT.

package project_revisions

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

	"github.com/itera-io/taikungoclient/models"
)

// NewProjectRevisionsEditParams creates a new ProjectRevisionsEditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectRevisionsEditParams() *ProjectRevisionsEditParams {
	return &ProjectRevisionsEditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectRevisionsEditParamsWithTimeout creates a new ProjectRevisionsEditParams object
// with the ability to set a timeout on a request.
func NewProjectRevisionsEditParamsWithTimeout(timeout time.Duration) *ProjectRevisionsEditParams {
	return &ProjectRevisionsEditParams{
		timeout: timeout,
	}
}

// NewProjectRevisionsEditParamsWithContext creates a new ProjectRevisionsEditParams object
// with the ability to set a context for a request.
func NewProjectRevisionsEditParamsWithContext(ctx context.Context) *ProjectRevisionsEditParams {
	return &ProjectRevisionsEditParams{
		Context: ctx,
	}
}

// NewProjectRevisionsEditParamsWithHTTPClient creates a new ProjectRevisionsEditParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectRevisionsEditParamsWithHTTPClient(client *http.Client) *ProjectRevisionsEditParams {
	return &ProjectRevisionsEditParams{
		HTTPClient: client,
	}
}

/* ProjectRevisionsEditParams contains all the parameters to send to the API endpoint
   for the project revisions edit operation.

   Typically these are written to a http.Request.
*/
type ProjectRevisionsEditParams struct {

	// Body.
	Body *models.ProjectRevisionUpdateDto

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

// WithDefaults hydrates default values in the project revisions edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectRevisionsEditParams) WithDefaults() *ProjectRevisionsEditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project revisions edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectRevisionsEditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithTimeout(timeout time.Duration) *ProjectRevisionsEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithContext(ctx context.Context) *ProjectRevisionsEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithHTTPClient(client *http.Client) *ProjectRevisionsEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithBody(body *models.ProjectRevisionUpdateDto) *ProjectRevisionsEditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetBody(body *models.ProjectRevisionUpdateDto) {
	o.Body = body
}

// WithProjectID adds the projectID to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithProjectID(projectID int32) *ProjectRevisionsEditParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the project revisions edit params
func (o *ProjectRevisionsEditParams) WithV(v string) *ProjectRevisionsEditParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project revisions edit params
func (o *ProjectRevisionsEditParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectRevisionsEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
