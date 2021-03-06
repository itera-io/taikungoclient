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

// NewProjectsMonitoringOperationsParams creates a new ProjectsMonitoringOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsMonitoringOperationsParams() *ProjectsMonitoringOperationsParams {
	return &ProjectsMonitoringOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsMonitoringOperationsParamsWithTimeout creates a new ProjectsMonitoringOperationsParams object
// with the ability to set a timeout on a request.
func NewProjectsMonitoringOperationsParamsWithTimeout(timeout time.Duration) *ProjectsMonitoringOperationsParams {
	return &ProjectsMonitoringOperationsParams{
		timeout: timeout,
	}
}

// NewProjectsMonitoringOperationsParamsWithContext creates a new ProjectsMonitoringOperationsParams object
// with the ability to set a context for a request.
func NewProjectsMonitoringOperationsParamsWithContext(ctx context.Context) *ProjectsMonitoringOperationsParams {
	return &ProjectsMonitoringOperationsParams{
		Context: ctx,
	}
}

// NewProjectsMonitoringOperationsParamsWithHTTPClient creates a new ProjectsMonitoringOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsMonitoringOperationsParamsWithHTTPClient(client *http.Client) *ProjectsMonitoringOperationsParams {
	return &ProjectsMonitoringOperationsParams{
		HTTPClient: client,
	}
}

/* ProjectsMonitoringOperationsParams contains all the parameters to send to the API endpoint
   for the projects monitoring operations operation.

   Typically these are written to a http.Request.
*/
type ProjectsMonitoringOperationsParams struct {

	// Body.
	Body *models.MonitoringOperationsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects monitoring operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsMonitoringOperationsParams) WithDefaults() *ProjectsMonitoringOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects monitoring operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsMonitoringOperationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) WithTimeout(timeout time.Duration) *ProjectsMonitoringOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) WithContext(ctx context.Context) *ProjectsMonitoringOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) WithHTTPClient(client *http.Client) *ProjectsMonitoringOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) WithBody(body *models.MonitoringOperationsCommand) *ProjectsMonitoringOperationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) SetBody(body *models.MonitoringOperationsCommand) {
	o.Body = body
}

// WithV adds the v to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) WithV(v string) *ProjectsMonitoringOperationsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects monitoring operations params
func (o *ProjectsMonitoringOperationsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsMonitoringOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
