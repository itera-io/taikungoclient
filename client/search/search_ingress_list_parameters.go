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

// NewSearchIngressListParams creates a new SearchIngressListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchIngressListParams() *SearchIngressListParams {
	return &SearchIngressListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchIngressListParamsWithTimeout creates a new SearchIngressListParams object
// with the ability to set a timeout on a request.
func NewSearchIngressListParamsWithTimeout(timeout time.Duration) *SearchIngressListParams {
	return &SearchIngressListParams{
		timeout: timeout,
	}
}

// NewSearchIngressListParamsWithContext creates a new SearchIngressListParams object
// with the ability to set a context for a request.
func NewSearchIngressListParamsWithContext(ctx context.Context) *SearchIngressListParams {
	return &SearchIngressListParams{
		Context: ctx,
	}
}

// NewSearchIngressListParamsWithHTTPClient creates a new SearchIngressListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchIngressListParamsWithHTTPClient(client *http.Client) *SearchIngressListParams {
	return &SearchIngressListParams{
		HTTPClient: client,
	}
}

/* SearchIngressListParams contains all the parameters to send to the API endpoint
   for the search ingress list operation.

   Typically these are written to a http.Request.
*/
type SearchIngressListParams struct {

	// Body.
	Body *models.IngressSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search ingress list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchIngressListParams) WithDefaults() *SearchIngressListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search ingress list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchIngressListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search ingress list params
func (o *SearchIngressListParams) WithTimeout(timeout time.Duration) *SearchIngressListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search ingress list params
func (o *SearchIngressListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search ingress list params
func (o *SearchIngressListParams) WithContext(ctx context.Context) *SearchIngressListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search ingress list params
func (o *SearchIngressListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search ingress list params
func (o *SearchIngressListParams) WithHTTPClient(client *http.Client) *SearchIngressListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search ingress list params
func (o *SearchIngressListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search ingress list params
func (o *SearchIngressListParams) WithBody(body *models.IngressSearchCommand) *SearchIngressListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search ingress list params
func (o *SearchIngressListParams) SetBody(body *models.IngressSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search ingress list params
func (o *SearchIngressListParams) WithV(v string) *SearchIngressListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search ingress list params
func (o *SearchIngressListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchIngressListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
