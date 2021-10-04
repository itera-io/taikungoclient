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

// NewSearchDaemonSetListParams creates a new SearchDaemonSetListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchDaemonSetListParams() *SearchDaemonSetListParams {
	return &SearchDaemonSetListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchDaemonSetListParamsWithTimeout creates a new SearchDaemonSetListParams object
// with the ability to set a timeout on a request.
func NewSearchDaemonSetListParamsWithTimeout(timeout time.Duration) *SearchDaemonSetListParams {
	return &SearchDaemonSetListParams{
		timeout: timeout,
	}
}

// NewSearchDaemonSetListParamsWithContext creates a new SearchDaemonSetListParams object
// with the ability to set a context for a request.
func NewSearchDaemonSetListParamsWithContext(ctx context.Context) *SearchDaemonSetListParams {
	return &SearchDaemonSetListParams{
		Context: ctx,
	}
}

// NewSearchDaemonSetListParamsWithHTTPClient creates a new SearchDaemonSetListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchDaemonSetListParamsWithHTTPClient(client *http.Client) *SearchDaemonSetListParams {
	return &SearchDaemonSetListParams{
		HTTPClient: client,
	}
}

/* SearchDaemonSetListParams contains all the parameters to send to the API endpoint
   for the search daemon set list operation.

   Typically these are written to a http.Request.
*/
type SearchDaemonSetListParams struct {

	// Body.
	Body *models.DaemonSetSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search daemon set list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDaemonSetListParams) WithDefaults() *SearchDaemonSetListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search daemon set list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchDaemonSetListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search daemon set list params
func (o *SearchDaemonSetListParams) WithTimeout(timeout time.Duration) *SearchDaemonSetListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search daemon set list params
func (o *SearchDaemonSetListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search daemon set list params
func (o *SearchDaemonSetListParams) WithContext(ctx context.Context) *SearchDaemonSetListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search daemon set list params
func (o *SearchDaemonSetListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search daemon set list params
func (o *SearchDaemonSetListParams) WithHTTPClient(client *http.Client) *SearchDaemonSetListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search daemon set list params
func (o *SearchDaemonSetListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search daemon set list params
func (o *SearchDaemonSetListParams) WithBody(body *models.DaemonSetSearchCommand) *SearchDaemonSetListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search daemon set list params
func (o *SearchDaemonSetListParams) SetBody(body *models.DaemonSetSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search daemon set list params
func (o *SearchDaemonSetListParams) WithV(v string) *SearchDaemonSetListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search daemon set list params
func (o *SearchDaemonSetListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchDaemonSetListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
