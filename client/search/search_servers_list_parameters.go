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

// NewSearchServersListParams creates a new SearchServersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchServersListParams() *SearchServersListParams {
	return &SearchServersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchServersListParamsWithTimeout creates a new SearchServersListParams object
// with the ability to set a timeout on a request.
func NewSearchServersListParamsWithTimeout(timeout time.Duration) *SearchServersListParams {
	return &SearchServersListParams{
		timeout: timeout,
	}
}

// NewSearchServersListParamsWithContext creates a new SearchServersListParams object
// with the ability to set a context for a request.
func NewSearchServersListParamsWithContext(ctx context.Context) *SearchServersListParams {
	return &SearchServersListParams{
		Context: ctx,
	}
}

// NewSearchServersListParamsWithHTTPClient creates a new SearchServersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchServersListParamsWithHTTPClient(client *http.Client) *SearchServersListParams {
	return &SearchServersListParams{
		HTTPClient: client,
	}
}

/* SearchServersListParams contains all the parameters to send to the API endpoint
   for the search servers list operation.

   Typically these are written to a http.Request.
*/
type SearchServersListParams struct {

	// Body.
	Body *models.ServersSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search servers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchServersListParams) WithDefaults() *SearchServersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search servers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchServersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search servers list params
func (o *SearchServersListParams) WithTimeout(timeout time.Duration) *SearchServersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search servers list params
func (o *SearchServersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search servers list params
func (o *SearchServersListParams) WithContext(ctx context.Context) *SearchServersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search servers list params
func (o *SearchServersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search servers list params
func (o *SearchServersListParams) WithHTTPClient(client *http.Client) *SearchServersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search servers list params
func (o *SearchServersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search servers list params
func (o *SearchServersListParams) WithBody(body *models.ServersSearchCommand) *SearchServersListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search servers list params
func (o *SearchServersListParams) SetBody(body *models.ServersSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search servers list params
func (o *SearchServersListParams) WithV(v string) *SearchServersListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search servers list params
func (o *SearchServersListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchServersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
