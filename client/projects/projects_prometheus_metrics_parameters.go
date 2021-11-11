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

// NewProjectsPrometheusMetricsParams creates a new ProjectsPrometheusMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsPrometheusMetricsParams() *ProjectsPrometheusMetricsParams {
	return &ProjectsPrometheusMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsPrometheusMetricsParamsWithTimeout creates a new ProjectsPrometheusMetricsParams object
// with the ability to set a timeout on a request.
func NewProjectsPrometheusMetricsParamsWithTimeout(timeout time.Duration) *ProjectsPrometheusMetricsParams {
	return &ProjectsPrometheusMetricsParams{
		timeout: timeout,
	}
}

// NewProjectsPrometheusMetricsParamsWithContext creates a new ProjectsPrometheusMetricsParams object
// with the ability to set a context for a request.
func NewProjectsPrometheusMetricsParamsWithContext(ctx context.Context) *ProjectsPrometheusMetricsParams {
	return &ProjectsPrometheusMetricsParams{
		Context: ctx,
	}
}

// NewProjectsPrometheusMetricsParamsWithHTTPClient creates a new ProjectsPrometheusMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsPrometheusMetricsParamsWithHTTPClient(client *http.Client) *ProjectsPrometheusMetricsParams {
	return &ProjectsPrometheusMetricsParams{
		HTTPClient: client,
	}
}

/* ProjectsPrometheusMetricsParams contains all the parameters to send to the API endpoint
   for the projects prometheus metrics operation.

   Typically these are written to a http.Request.
*/
type ProjectsPrometheusMetricsParams struct {

	// Body.
	Body *models.PrometheusMetricsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects prometheus metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsPrometheusMetricsParams) WithDefaults() *ProjectsPrometheusMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects prometheus metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsPrometheusMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) WithTimeout(timeout time.Duration) *ProjectsPrometheusMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) WithContext(ctx context.Context) *ProjectsPrometheusMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) WithHTTPClient(client *http.Client) *ProjectsPrometheusMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) WithBody(body *models.PrometheusMetricsCommand) *ProjectsPrometheusMetricsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) SetBody(body *models.PrometheusMetricsCommand) {
	o.Body = body
}

// WithV adds the v to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) WithV(v string) *ProjectsPrometheusMetricsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects prometheus metrics params
func (o *ProjectsPrometheusMetricsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsPrometheusMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
