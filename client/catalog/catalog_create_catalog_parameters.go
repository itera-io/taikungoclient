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

// NewCatalogCreateCatalogParams creates a new CatalogCreateCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogCreateCatalogParams() *CatalogCreateCatalogParams {
	return &CatalogCreateCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCreateCatalogParamsWithTimeout creates a new CatalogCreateCatalogParams object
// with the ability to set a timeout on a request.
func NewCatalogCreateCatalogParamsWithTimeout(timeout time.Duration) *CatalogCreateCatalogParams {
	return &CatalogCreateCatalogParams{
		timeout: timeout,
	}
}

// NewCatalogCreateCatalogParamsWithContext creates a new CatalogCreateCatalogParams object
// with the ability to set a context for a request.
func NewCatalogCreateCatalogParamsWithContext(ctx context.Context) *CatalogCreateCatalogParams {
	return &CatalogCreateCatalogParams{
		Context: ctx,
	}
}

// NewCatalogCreateCatalogParamsWithHTTPClient creates a new CatalogCreateCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogCreateCatalogParamsWithHTTPClient(client *http.Client) *CatalogCreateCatalogParams {
	return &CatalogCreateCatalogParams{
		HTTPClient: client,
	}
}

/*
CatalogCreateCatalogParams contains all the parameters to send to the API endpoint

	for the catalog create catalog operation.

	Typically these are written to a http.Request.
*/
type CatalogCreateCatalogParams struct {

	// Body.
	Body CatalogCreateCatalogBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog create catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCreateCatalogParams) WithDefaults() *CatalogCreateCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog create catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCreateCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog create catalog params
func (o *CatalogCreateCatalogParams) WithTimeout(timeout time.Duration) *CatalogCreateCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog create catalog params
func (o *CatalogCreateCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog create catalog params
func (o *CatalogCreateCatalogParams) WithContext(ctx context.Context) *CatalogCreateCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog create catalog params
func (o *CatalogCreateCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog create catalog params
func (o *CatalogCreateCatalogParams) WithHTTPClient(client *http.Client) *CatalogCreateCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog create catalog params
func (o *CatalogCreateCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog create catalog params
func (o *CatalogCreateCatalogParams) WithBody(body CatalogCreateCatalogBody) *CatalogCreateCatalogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog create catalog params
func (o *CatalogCreateCatalogParams) SetBody(body CatalogCreateCatalogBody) {
	o.Body = body
}

// WithV adds the v to the catalog create catalog params
func (o *CatalogCreateCatalogParams) WithV(v string) *CatalogCreateCatalogParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog create catalog params
func (o *CatalogCreateCatalogParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCreateCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
