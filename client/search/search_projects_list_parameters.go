// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewSearchProjectsListParams creates a new SearchProjectsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchProjectsListParams() *SearchProjectsListParams {
	return &SearchProjectsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchProjectsListParamsWithTimeout creates a new SearchProjectsListParams object
// with the ability to set a timeout on a request.
func NewSearchProjectsListParamsWithTimeout(timeout time.Duration) *SearchProjectsListParams {
	return &SearchProjectsListParams{
		timeout: timeout,
	}
}

// NewSearchProjectsListParamsWithContext creates a new SearchProjectsListParams object
// with the ability to set a context for a request.
func NewSearchProjectsListParamsWithContext(ctx context.Context) *SearchProjectsListParams {
	return &SearchProjectsListParams{
		Context: ctx,
	}
}

// NewSearchProjectsListParamsWithHTTPClient creates a new SearchProjectsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchProjectsListParamsWithHTTPClient(client *http.Client) *SearchProjectsListParams {
	return &SearchProjectsListParams{
		HTTPClient: client,
	}
}

/*
SearchProjectsListParams contains all the parameters to send to the API endpoint

	for the search projects list operation.

	Typically these are written to a http.Request.
*/
type SearchProjectsListParams struct {

	// Body.
	Body *models.ProjectsSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search projects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchProjectsListParams) WithDefaults() *SearchProjectsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search projects list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchProjectsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search projects list params
func (o *SearchProjectsListParams) WithTimeout(timeout time.Duration) *SearchProjectsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search projects list params
func (o *SearchProjectsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search projects list params
func (o *SearchProjectsListParams) WithContext(ctx context.Context) *SearchProjectsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search projects list params
func (o *SearchProjectsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search projects list params
func (o *SearchProjectsListParams) WithHTTPClient(client *http.Client) *SearchProjectsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search projects list params
func (o *SearchProjectsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search projects list params
func (o *SearchProjectsListParams) WithBody(body *models.ProjectsSearchCommand) *SearchProjectsListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search projects list params
func (o *SearchProjectsListParams) SetBody(body *models.ProjectsSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search projects list params
func (o *SearchProjectsListParams) WithV(v string) *SearchProjectsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search projects list params
func (o *SearchProjectsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchProjectsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
