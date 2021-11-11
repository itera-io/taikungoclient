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

// NewSearchAccessProfilesListParams creates a new SearchAccessProfilesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchAccessProfilesListParams() *SearchAccessProfilesListParams {
	return &SearchAccessProfilesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAccessProfilesListParamsWithTimeout creates a new SearchAccessProfilesListParams object
// with the ability to set a timeout on a request.
func NewSearchAccessProfilesListParamsWithTimeout(timeout time.Duration) *SearchAccessProfilesListParams {
	return &SearchAccessProfilesListParams{
		timeout: timeout,
	}
}

// NewSearchAccessProfilesListParamsWithContext creates a new SearchAccessProfilesListParams object
// with the ability to set a context for a request.
func NewSearchAccessProfilesListParamsWithContext(ctx context.Context) *SearchAccessProfilesListParams {
	return &SearchAccessProfilesListParams{
		Context: ctx,
	}
}

// NewSearchAccessProfilesListParamsWithHTTPClient creates a new SearchAccessProfilesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchAccessProfilesListParamsWithHTTPClient(client *http.Client) *SearchAccessProfilesListParams {
	return &SearchAccessProfilesListParams{
		HTTPClient: client,
	}
}

/* SearchAccessProfilesListParams contains all the parameters to send to the API endpoint
   for the search access profiles list operation.

   Typically these are written to a http.Request.
*/
type SearchAccessProfilesListParams struct {

	// Body.
	Body *models.AccessProfilesSearchCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search access profiles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAccessProfilesListParams) WithDefaults() *SearchAccessProfilesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search access profiles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchAccessProfilesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search access profiles list params
func (o *SearchAccessProfilesListParams) WithTimeout(timeout time.Duration) *SearchAccessProfilesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search access profiles list params
func (o *SearchAccessProfilesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search access profiles list params
func (o *SearchAccessProfilesListParams) WithContext(ctx context.Context) *SearchAccessProfilesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search access profiles list params
func (o *SearchAccessProfilesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search access profiles list params
func (o *SearchAccessProfilesListParams) WithHTTPClient(client *http.Client) *SearchAccessProfilesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search access profiles list params
func (o *SearchAccessProfilesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search access profiles list params
func (o *SearchAccessProfilesListParams) WithBody(body *models.AccessProfilesSearchCommand) *SearchAccessProfilesListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search access profiles list params
func (o *SearchAccessProfilesListParams) SetBody(body *models.AccessProfilesSearchCommand) {
	o.Body = body
}

// WithV adds the v to the search access profiles list params
func (o *SearchAccessProfilesListParams) WithV(v string) *SearchAccessProfilesListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the search access profiles list params
func (o *SearchAccessProfilesListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAccessProfilesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
