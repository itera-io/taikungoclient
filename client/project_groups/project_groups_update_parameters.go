// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewProjectGroupsUpdateParams creates a new ProjectGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGroupsUpdateParams() *ProjectGroupsUpdateParams {
	return &ProjectGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGroupsUpdateParamsWithTimeout creates a new ProjectGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewProjectGroupsUpdateParamsWithTimeout(timeout time.Duration) *ProjectGroupsUpdateParams {
	return &ProjectGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewProjectGroupsUpdateParamsWithContext creates a new ProjectGroupsUpdateParams object
// with the ability to set a context for a request.
func NewProjectGroupsUpdateParamsWithContext(ctx context.Context) *ProjectGroupsUpdateParams {
	return &ProjectGroupsUpdateParams{
		Context: ctx,
	}
}

// NewProjectGroupsUpdateParamsWithHTTPClient creates a new ProjectGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGroupsUpdateParamsWithHTTPClient(client *http.Client) *ProjectGroupsUpdateParams {
	return &ProjectGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* ProjectGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the project groups update operation.

   Typically these are written to a http.Request.
*/
type ProjectGroupsUpdateParams struct {

	// Body.
	Body *models.UpdateProjectGroupDto

	// ProjectGroupID.
	//
	// Format: int32
	ProjectGroupID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsUpdateParams) WithDefaults() *ProjectGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project groups update params
func (o *ProjectGroupsUpdateParams) WithTimeout(timeout time.Duration) *ProjectGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project groups update params
func (o *ProjectGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project groups update params
func (o *ProjectGroupsUpdateParams) WithContext(ctx context.Context) *ProjectGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project groups update params
func (o *ProjectGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project groups update params
func (o *ProjectGroupsUpdateParams) WithHTTPClient(client *http.Client) *ProjectGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project groups update params
func (o *ProjectGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project groups update params
func (o *ProjectGroupsUpdateParams) WithBody(body *models.UpdateProjectGroupDto) *ProjectGroupsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project groups update params
func (o *ProjectGroupsUpdateParams) SetBody(body *models.UpdateProjectGroupDto) {
	o.Body = body
}

// WithProjectGroupID adds the projectGroupID to the project groups update params
func (o *ProjectGroupsUpdateParams) WithProjectGroupID(projectGroupID *int32) *ProjectGroupsUpdateParams {
	o.SetProjectGroupID(projectGroupID)
	return o
}

// SetProjectGroupID adds the projectGroupId to the project groups update params
func (o *ProjectGroupsUpdateParams) SetProjectGroupID(projectGroupID *int32) {
	o.ProjectGroupID = projectGroupID
}

// WithV adds the v to the project groups update params
func (o *ProjectGroupsUpdateParams) WithV(v string) *ProjectGroupsUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project groups update params
func (o *ProjectGroupsUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ProjectGroupID != nil {

		// query param projectGroupId
		var qrProjectGroupID int32

		if o.ProjectGroupID != nil {
			qrProjectGroupID = *o.ProjectGroupID
		}
		qProjectGroupID := swag.FormatInt32(qrProjectGroupID)
		if qProjectGroupID != "" {

			if err := r.SetQueryParam("projectGroupId", qProjectGroupID); err != nil {
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
