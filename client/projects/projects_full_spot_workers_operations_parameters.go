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

// NewProjectsFullSpotWorkersOperationsParams creates a new ProjectsFullSpotWorkersOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsFullSpotWorkersOperationsParams() *ProjectsFullSpotWorkersOperationsParams {
	return &ProjectsFullSpotWorkersOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsFullSpotWorkersOperationsParamsWithTimeout creates a new ProjectsFullSpotWorkersOperationsParams object
// with the ability to set a timeout on a request.
func NewProjectsFullSpotWorkersOperationsParamsWithTimeout(timeout time.Duration) *ProjectsFullSpotWorkersOperationsParams {
	return &ProjectsFullSpotWorkersOperationsParams{
		timeout: timeout,
	}
}

// NewProjectsFullSpotWorkersOperationsParamsWithContext creates a new ProjectsFullSpotWorkersOperationsParams object
// with the ability to set a context for a request.
func NewProjectsFullSpotWorkersOperationsParamsWithContext(ctx context.Context) *ProjectsFullSpotWorkersOperationsParams {
	return &ProjectsFullSpotWorkersOperationsParams{
		Context: ctx,
	}
}

// NewProjectsFullSpotWorkersOperationsParamsWithHTTPClient creates a new ProjectsFullSpotWorkersOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsFullSpotWorkersOperationsParamsWithHTTPClient(client *http.Client) *ProjectsFullSpotWorkersOperationsParams {
	return &ProjectsFullSpotWorkersOperationsParams{
		HTTPClient: client,
	}
}

/* ProjectsFullSpotWorkersOperationsParams contains all the parameters to send to the API endpoint
   for the projects full spot workers operations operation.

   Typically these are written to a http.Request.
*/
type ProjectsFullSpotWorkersOperationsParams struct {

	// Body.
	Body *models.FullSpotOperationCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects full spot workers operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsFullSpotWorkersOperationsParams) WithDefaults() *ProjectsFullSpotWorkersOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects full spot workers operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsFullSpotWorkersOperationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) WithTimeout(timeout time.Duration) *ProjectsFullSpotWorkersOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) WithContext(ctx context.Context) *ProjectsFullSpotWorkersOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) WithHTTPClient(client *http.Client) *ProjectsFullSpotWorkersOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) WithBody(body *models.FullSpotOperationCommand) *ProjectsFullSpotWorkersOperationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) SetBody(body *models.FullSpotOperationCommand) {
	o.Body = body
}

// WithV adds the v to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) WithV(v string) *ProjectsFullSpotWorkersOperationsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects full spot workers operations params
func (o *ProjectsFullSpotWorkersOperationsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsFullSpotWorkersOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
