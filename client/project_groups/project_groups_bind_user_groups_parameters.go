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

// NewProjectGroupsBindUserGroupsParams creates a new ProjectGroupsBindUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGroupsBindUserGroupsParams() *ProjectGroupsBindUserGroupsParams {
	return &ProjectGroupsBindUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGroupsBindUserGroupsParamsWithTimeout creates a new ProjectGroupsBindUserGroupsParams object
// with the ability to set a timeout on a request.
func NewProjectGroupsBindUserGroupsParamsWithTimeout(timeout time.Duration) *ProjectGroupsBindUserGroupsParams {
	return &ProjectGroupsBindUserGroupsParams{
		timeout: timeout,
	}
}

// NewProjectGroupsBindUserGroupsParamsWithContext creates a new ProjectGroupsBindUserGroupsParams object
// with the ability to set a context for a request.
func NewProjectGroupsBindUserGroupsParamsWithContext(ctx context.Context) *ProjectGroupsBindUserGroupsParams {
	return &ProjectGroupsBindUserGroupsParams{
		Context: ctx,
	}
}

// NewProjectGroupsBindUserGroupsParamsWithHTTPClient creates a new ProjectGroupsBindUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGroupsBindUserGroupsParamsWithHTTPClient(client *http.Client) *ProjectGroupsBindUserGroupsParams {
	return &ProjectGroupsBindUserGroupsParams{
		HTTPClient: client,
	}
}

/*
ProjectGroupsBindUserGroupsParams contains all the parameters to send to the API endpoint

	for the project groups bind user groups operation.

	Typically these are written to a http.Request.
*/
type ProjectGroupsBindUserGroupsParams struct {

	// Body.
	Body *models.BindUserGroupsToProjectGroupCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project groups bind user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsBindUserGroupsParams) WithDefaults() *ProjectGroupsBindUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project groups bind user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsBindUserGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) WithTimeout(timeout time.Duration) *ProjectGroupsBindUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) WithContext(ctx context.Context) *ProjectGroupsBindUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) WithHTTPClient(client *http.Client) *ProjectGroupsBindUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) WithBody(body *models.BindUserGroupsToProjectGroupCommand) *ProjectGroupsBindUserGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) SetBody(body *models.BindUserGroupsToProjectGroupCommand) {
	o.Body = body
}

// WithV adds the v to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) WithV(v string) *ProjectGroupsBindUserGroupsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project groups bind user groups params
func (o *ProjectGroupsBindUserGroupsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGroupsBindUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
