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
	"github.com/go-openapi/swag"
)

// NewProjectAppUninstallParams creates a new ProjectAppUninstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectAppUninstallParams() *ProjectAppUninstallParams {
	return &ProjectAppUninstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectAppUninstallParamsWithTimeout creates a new ProjectAppUninstallParams object
// with the ability to set a timeout on a request.
func NewProjectAppUninstallParamsWithTimeout(timeout time.Duration) *ProjectAppUninstallParams {
	return &ProjectAppUninstallParams{
		timeout: timeout,
	}
}

// NewProjectAppUninstallParamsWithContext creates a new ProjectAppUninstallParams object
// with the ability to set a context for a request.
func NewProjectAppUninstallParamsWithContext(ctx context.Context) *ProjectAppUninstallParams {
	return &ProjectAppUninstallParams{
		Context: ctx,
	}
}

// NewProjectAppUninstallParamsWithHTTPClient creates a new ProjectAppUninstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectAppUninstallParamsWithHTTPClient(client *http.Client) *ProjectAppUninstallParams {
	return &ProjectAppUninstallParams{
		HTTPClient: client,
	}
}

/* ProjectAppUninstallParams contains all the parameters to send to the API endpoint
   for the project app uninstall operation.

   Typically these are written to a http.Request.
*/
type ProjectAppUninstallParams struct {

	// ProjectAppID.
	//
	// Format: int32
	ProjectAppID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project app uninstall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectAppUninstallParams) WithDefaults() *ProjectAppUninstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project app uninstall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectAppUninstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project app uninstall params
func (o *ProjectAppUninstallParams) WithTimeout(timeout time.Duration) *ProjectAppUninstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project app uninstall params
func (o *ProjectAppUninstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project app uninstall params
func (o *ProjectAppUninstallParams) WithContext(ctx context.Context) *ProjectAppUninstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project app uninstall params
func (o *ProjectAppUninstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project app uninstall params
func (o *ProjectAppUninstallParams) WithHTTPClient(client *http.Client) *ProjectAppUninstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project app uninstall params
func (o *ProjectAppUninstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectAppID adds the projectAppID to the project app uninstall params
func (o *ProjectAppUninstallParams) WithProjectAppID(projectAppID int32) *ProjectAppUninstallParams {
	o.SetProjectAppID(projectAppID)
	return o
}

// SetProjectAppID adds the projectAppId to the project app uninstall params
func (o *ProjectAppUninstallParams) SetProjectAppID(projectAppID int32) {
	o.ProjectAppID = projectAppID
}

// WithV adds the v to the project app uninstall params
func (o *ProjectAppUninstallParams) WithV(v string) *ProjectAppUninstallParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project app uninstall params
func (o *ProjectAppUninstallParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectAppUninstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectAppId
	if err := r.SetPathParam("projectAppId", swag.FormatInt32(o.ProjectAppID)); err != nil {
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
