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

// NewCatalogCreateCatalogAppParams creates a new CatalogCreateCatalogAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogCreateCatalogAppParams() *CatalogCreateCatalogAppParams {
	return &CatalogCreateCatalogAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCreateCatalogAppParamsWithTimeout creates a new CatalogCreateCatalogAppParams object
// with the ability to set a timeout on a request.
func NewCatalogCreateCatalogAppParamsWithTimeout(timeout time.Duration) *CatalogCreateCatalogAppParams {
	return &CatalogCreateCatalogAppParams{
		timeout: timeout,
	}
}

// NewCatalogCreateCatalogAppParamsWithContext creates a new CatalogCreateCatalogAppParams object
// with the ability to set a context for a request.
func NewCatalogCreateCatalogAppParamsWithContext(ctx context.Context) *CatalogCreateCatalogAppParams {
	return &CatalogCreateCatalogAppParams{
		Context: ctx,
	}
}

// NewCatalogCreateCatalogAppParamsWithHTTPClient creates a new CatalogCreateCatalogAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogCreateCatalogAppParamsWithHTTPClient(client *http.Client) *CatalogCreateCatalogAppParams {
	return &CatalogCreateCatalogAppParams{
		HTTPClient: client,
	}
}

/*
CatalogCreateCatalogAppParams contains all the parameters to send to the API endpoint

	for the catalog create catalog app operation.

	Typically these are written to a http.Request.
*/
type CatalogCreateCatalogAppParams struct {

	// Body.
	Body CatalogCreateCatalogAppBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog create catalog app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCreateCatalogAppParams) WithDefaults() *CatalogCreateCatalogAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog create catalog app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCreateCatalogAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) WithTimeout(timeout time.Duration) *CatalogCreateCatalogAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) WithContext(ctx context.Context) *CatalogCreateCatalogAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) WithHTTPClient(client *http.Client) *CatalogCreateCatalogAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) WithBody(body CatalogCreateCatalogAppBody) *CatalogCreateCatalogAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) SetBody(body CatalogCreateCatalogAppBody) {
	o.Body = body
}

// WithV adds the v to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) WithV(v string) *CatalogCreateCatalogAppParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog create catalog app params
func (o *CatalogCreateCatalogAppParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCreateCatalogAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
