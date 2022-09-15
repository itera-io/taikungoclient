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
	"github.com/go-openapi/swag"
)

// NewCatalogCatalogAppDetailsParams creates a new CatalogCatalogAppDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogCatalogAppDetailsParams() *CatalogCatalogAppDetailsParams {
	return &CatalogCatalogAppDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCatalogAppDetailsParamsWithTimeout creates a new CatalogCatalogAppDetailsParams object
// with the ability to set a timeout on a request.
func NewCatalogCatalogAppDetailsParamsWithTimeout(timeout time.Duration) *CatalogCatalogAppDetailsParams {
	return &CatalogCatalogAppDetailsParams{
		timeout: timeout,
	}
}

// NewCatalogCatalogAppDetailsParamsWithContext creates a new CatalogCatalogAppDetailsParams object
// with the ability to set a context for a request.
func NewCatalogCatalogAppDetailsParamsWithContext(ctx context.Context) *CatalogCatalogAppDetailsParams {
	return &CatalogCatalogAppDetailsParams{
		Context: ctx,
	}
}

// NewCatalogCatalogAppDetailsParamsWithHTTPClient creates a new CatalogCatalogAppDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogCatalogAppDetailsParamsWithHTTPClient(client *http.Client) *CatalogCatalogAppDetailsParams {
	return &CatalogCatalogAppDetailsParams{
		HTTPClient: client,
	}
}

/*
CatalogCatalogAppDetailsParams contains all the parameters to send to the API endpoint

	for the catalog catalog app details operation.

	Typically these are written to a http.Request.
*/
type CatalogCatalogAppDetailsParams struct {

	/* ID.

	   catalog app id

	   Format: int32
	*/
	ID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog catalog app details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppDetailsParams) WithDefaults() *CatalogCatalogAppDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog catalog app details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) WithTimeout(timeout time.Duration) *CatalogCatalogAppDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) WithContext(ctx context.Context) *CatalogCatalogAppDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) WithHTTPClient(client *http.Client) *CatalogCatalogAppDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) WithID(id int32) *CatalogCatalogAppDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) WithV(v string) *CatalogCatalogAppDetailsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog catalog app details params
func (o *CatalogCatalogAppDetailsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCatalogAppDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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
