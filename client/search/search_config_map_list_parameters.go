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

// NewSearchConfigMapListParams creates a new SearchConfigMapListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchConfigMapListParams() *SearchConfigMapListParams {
	return &SearchConfigMapListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchConfigMapListParamsWithTimeout creates a new SearchConfigMapListParams object
// with the ability to set a timeout on a request.
func NewSearchConfigMapListParamsWithTimeout(timeout time.Duration) *SearchConfigMapListParams {
	return &SearchConfigMapListParams{
		timeout: timeout,
	}
}

// NewSearchConfigMapListParamsWithContext creates a new SearchConfigMapListParams object
// with the ability to set a context for a request.
func NewSearchConfigMapListParamsWithContext(ctx context.Context) *SearchConfigMapListParams {
	return &SearchConfigMapListParams{
		Context: ctx,
	}
}

// NewSearchConfigMapListParamsWithHTTPClient creates a new SearchConfigMapListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchConfigMapListParamsWithHTTPClient(client *http.Client) *SearchConfigMapListParams {
	return &SearchConfigMapListParams{
		HTTPClient: client,
	}
}

/* SearchConfigMapListParams contains all the parameters to send to the API endpoint
   for the search config map list operation.

   Typically these are written to a http.Request.
*/
type SearchConfigMapListParams struct {

	// Body.
	Body *models.ConfigMapSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search config map list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchConfigMapListParams) WithDefaults() *SearchConfigMapListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search config map list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchConfigMapListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search config map list params
func (o *SearchConfigMapListParams) WithTimeout(timeout time.Duration) *SearchConfigMapListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search config map list params
func (o *SearchConfigMapListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search config map list params
func (o *SearchConfigMapListParams) WithContext(ctx context.Context) *SearchConfigMapListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search config map list params
func (o *SearchConfigMapListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search config map list params
func (o *SearchConfigMapListParams) WithHTTPClient(client *http.Client) *SearchConfigMapListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search config map list params
func (o *SearchConfigMapListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search config map list params
func (o *SearchConfigMapListParams) WithBody(body *models.ConfigMapSearchCommand) *SearchConfigMapListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search config map list params
func (o *SearchConfigMapListParams) SetBody(body *models.ConfigMapSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search config map list params
func (o *SearchConfigMapListParams) WithV(v string) *SearchConfigMapListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search config map list params
func (o *SearchConfigMapListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchConfigMapListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
