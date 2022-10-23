// Code generated by go-swagger; DO NOT EDIT.

package catalog

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
)

// NewCatalogCatalogAppValueAutocompleteParams creates a new CatalogCatalogAppValueAutocompleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogCatalogAppValueAutocompleteParams() *CatalogCatalogAppValueAutocompleteParams {
	return &CatalogCatalogAppValueAutocompleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCatalogAppValueAutocompleteParamsWithTimeout creates a new CatalogCatalogAppValueAutocompleteParams object
// with the ability to set a timeout on a request.
func NewCatalogCatalogAppValueAutocompleteParamsWithTimeout(timeout time.Duration) *CatalogCatalogAppValueAutocompleteParams {
	return &CatalogCatalogAppValueAutocompleteParams{
		timeout: timeout,
	}
}

// NewCatalogCatalogAppValueAutocompleteParamsWithContext creates a new CatalogCatalogAppValueAutocompleteParams object
// with the ability to set a context for a request.
func NewCatalogCatalogAppValueAutocompleteParamsWithContext(ctx context.Context) *CatalogCatalogAppValueAutocompleteParams {
	return &CatalogCatalogAppValueAutocompleteParams{
		Context: ctx,
	}
}

// NewCatalogCatalogAppValueAutocompleteParamsWithHTTPClient creates a new CatalogCatalogAppValueAutocompleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogCatalogAppValueAutocompleteParamsWithHTTPClient(client *http.Client) *CatalogCatalogAppValueAutocompleteParams {
	return &CatalogCatalogAppValueAutocompleteParams{
		HTTPClient: client,
	}
}

/*
CatalogCatalogAppValueAutocompleteParams contains all the parameters to send to the API endpoint

	for the catalog catalog app value autocomplete operation.

	Typically these are written to a http.Request.
*/
type CatalogCatalogAppValueAutocompleteParams struct {

	// Body.
	Body CatalogCatalogAppValueAutocompleteBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog catalog app value autocomplete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppValueAutocompleteParams) WithDefaults() *CatalogCatalogAppValueAutocompleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog catalog app value autocomplete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppValueAutocompleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) WithTimeout(timeout time.Duration) *CatalogCatalogAppValueAutocompleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) WithContext(ctx context.Context) *CatalogCatalogAppValueAutocompleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) WithHTTPClient(client *http.Client) *CatalogCatalogAppValueAutocompleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) WithBody(body CatalogCatalogAppValueAutocompleteBody) *CatalogCatalogAppValueAutocompleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) SetBody(body CatalogCatalogAppValueAutocompleteBody) {
	o.Body = body
}

// WithV adds the v to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) WithV(v string) *CatalogCatalogAppValueAutocompleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog catalog app value autocomplete params
func (o *CatalogCatalogAppValueAutocompleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCatalogAppValueAutocompleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
