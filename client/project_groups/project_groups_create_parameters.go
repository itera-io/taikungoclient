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

	"github.com/itera-io/taikungoclient/models"
)

// NewProjectGroupsCreateParams creates a new ProjectGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGroupsCreateParams() *ProjectGroupsCreateParams {
	return &ProjectGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGroupsCreateParamsWithTimeout creates a new ProjectGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewProjectGroupsCreateParamsWithTimeout(timeout time.Duration) *ProjectGroupsCreateParams {
	return &ProjectGroupsCreateParams{
		timeout: timeout,
	}
}

// NewProjectGroupsCreateParamsWithContext creates a new ProjectGroupsCreateParams object
// with the ability to set a context for a request.
func NewProjectGroupsCreateParamsWithContext(ctx context.Context) *ProjectGroupsCreateParams {
	return &ProjectGroupsCreateParams{
		Context: ctx,
	}
}

// NewProjectGroupsCreateParamsWithHTTPClient creates a new ProjectGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGroupsCreateParamsWithHTTPClient(client *http.Client) *ProjectGroupsCreateParams {
	return &ProjectGroupsCreateParams{
		HTTPClient: client,
	}
}

/* ProjectGroupsCreateParams contains all the parameters to send to the API endpoint
   for the project groups create operation.

   Typically these are written to a http.Request.
*/
type ProjectGroupsCreateParams struct {

	// Body.
	Body *models.CreateProjectGroupCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsCreateParams) WithDefaults() *ProjectGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project groups create params
func (o *ProjectGroupsCreateParams) WithTimeout(timeout time.Duration) *ProjectGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project groups create params
func (o *ProjectGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project groups create params
func (o *ProjectGroupsCreateParams) WithContext(ctx context.Context) *ProjectGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project groups create params
func (o *ProjectGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project groups create params
func (o *ProjectGroupsCreateParams) WithHTTPClient(client *http.Client) *ProjectGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project groups create params
func (o *ProjectGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project groups create params
func (o *ProjectGroupsCreateParams) WithBody(body *models.CreateProjectGroupCommand) *ProjectGroupsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project groups create params
func (o *ProjectGroupsCreateParams) SetBody(body *models.CreateProjectGroupCommand) {
	o.Body = body
}

// WithV adds the v to the project groups create params
func (o *ProjectGroupsCreateParams) WithV(v string) *ProjectGroupsCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project groups create params
func (o *ProjectGroupsCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
