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

	"github.com/itera-io/taikungoclient/models"
)

// NewProjectsProjectMonitoringAlertsParams creates a new ProjectsProjectMonitoringAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsProjectMonitoringAlertsParams() *ProjectsProjectMonitoringAlertsParams {
	return &ProjectsProjectMonitoringAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsProjectMonitoringAlertsParamsWithTimeout creates a new ProjectsProjectMonitoringAlertsParams object
// with the ability to set a timeout on a request.
func NewProjectsProjectMonitoringAlertsParamsWithTimeout(timeout time.Duration) *ProjectsProjectMonitoringAlertsParams {
	return &ProjectsProjectMonitoringAlertsParams{
		timeout: timeout,
	}
}

// NewProjectsProjectMonitoringAlertsParamsWithContext creates a new ProjectsProjectMonitoringAlertsParams object
// with the ability to set a context for a request.
func NewProjectsProjectMonitoringAlertsParamsWithContext(ctx context.Context) *ProjectsProjectMonitoringAlertsParams {
	return &ProjectsProjectMonitoringAlertsParams{
		Context: ctx,
	}
}

// NewProjectsProjectMonitoringAlertsParamsWithHTTPClient creates a new ProjectsProjectMonitoringAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsProjectMonitoringAlertsParamsWithHTTPClient(client *http.Client) *ProjectsProjectMonitoringAlertsParams {
	return &ProjectsProjectMonitoringAlertsParams{
		HTTPClient: client,
	}
}

/*
ProjectsProjectMonitoringAlertsParams contains all the parameters to send to the API endpoint

	for the projects project monitoring alerts operation.

	Typically these are written to a http.Request.
*/
type ProjectsProjectMonitoringAlertsParams struct {

	// Body.
	Body *models.ProjectsMonitoringAlertsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects project monitoring alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectMonitoringAlertsParams) WithDefaults() *ProjectsProjectMonitoringAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects project monitoring alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsProjectMonitoringAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) WithTimeout(timeout time.Duration) *ProjectsProjectMonitoringAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) WithContext(ctx context.Context) *ProjectsProjectMonitoringAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) WithHTTPClient(client *http.Client) *ProjectsProjectMonitoringAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) WithBody(body *models.ProjectsMonitoringAlertsCommand) *ProjectsProjectMonitoringAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) SetBody(body *models.ProjectsMonitoringAlertsCommand) {
	o.Body = body
}

// WithV adds the v to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) WithV(v string) *ProjectsProjectMonitoringAlertsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects project monitoring alerts params
func (o *ProjectsProjectMonitoringAlertsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsProjectMonitoringAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
