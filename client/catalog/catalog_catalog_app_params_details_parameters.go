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

// NewCatalogCatalogAppParamsDetailsParams creates a new CatalogCatalogAppParamsDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogCatalogAppParamsDetailsParams() *CatalogCatalogAppParamsDetailsParams {
	return &CatalogCatalogAppParamsDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogCatalogAppParamsDetailsParamsWithTimeout creates a new CatalogCatalogAppParamsDetailsParams object
// with the ability to set a timeout on a request.
func NewCatalogCatalogAppParamsDetailsParamsWithTimeout(timeout time.Duration) *CatalogCatalogAppParamsDetailsParams {
	return &CatalogCatalogAppParamsDetailsParams{
		timeout: timeout,
	}
}

// NewCatalogCatalogAppParamsDetailsParamsWithContext creates a new CatalogCatalogAppParamsDetailsParams object
// with the ability to set a context for a request.
func NewCatalogCatalogAppParamsDetailsParamsWithContext(ctx context.Context) *CatalogCatalogAppParamsDetailsParams {
	return &CatalogCatalogAppParamsDetailsParams{
		Context: ctx,
	}
}

// NewCatalogCatalogAppParamsDetailsParamsWithHTTPClient creates a new CatalogCatalogAppParamsDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogCatalogAppParamsDetailsParamsWithHTTPClient(client *http.Client) *CatalogCatalogAppParamsDetailsParams {
	return &CatalogCatalogAppParamsDetailsParams{
		HTTPClient: client,
	}
}

/* CatalogCatalogAppParamsDetailsParams contains all the parameters to send to the API endpoint
   for the catalog catalog app params details operation.

   Typically these are written to a http.Request.
*/
type CatalogCatalogAppParamsDetailsParams struct {

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

// WithDefaults hydrates default values in the catalog catalog app params details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppParamsDetailsParams) WithDefaults() *CatalogCatalogAppParamsDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog catalog app params details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogCatalogAppParamsDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) WithTimeout(timeout time.Duration) *CatalogCatalogAppParamsDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) WithContext(ctx context.Context) *CatalogCatalogAppParamsDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) WithHTTPClient(client *http.Client) *CatalogCatalogAppParamsDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) WithID(id int32) *CatalogCatalogAppParamsDetailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) WithV(v string) *CatalogCatalogAppParamsDetailsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog catalog app params details params
func (o *CatalogCatalogAppParamsDetailsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogCatalogAppParamsDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
