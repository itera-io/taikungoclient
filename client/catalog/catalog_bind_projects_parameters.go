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

	"github.com/itera-io/taikungoclient/models"
)

// NewCatalogBindProjectsParams creates a new CatalogBindProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogBindProjectsParams() *CatalogBindProjectsParams {
	return &CatalogBindProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogBindProjectsParamsWithTimeout creates a new CatalogBindProjectsParams object
// with the ability to set a timeout on a request.
func NewCatalogBindProjectsParamsWithTimeout(timeout time.Duration) *CatalogBindProjectsParams {
	return &CatalogBindProjectsParams{
		timeout: timeout,
	}
}

// NewCatalogBindProjectsParamsWithContext creates a new CatalogBindProjectsParams object
// with the ability to set a context for a request.
func NewCatalogBindProjectsParamsWithContext(ctx context.Context) *CatalogBindProjectsParams {
	return &CatalogBindProjectsParams{
		Context: ctx,
	}
}

// NewCatalogBindProjectsParamsWithHTTPClient creates a new CatalogBindProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogBindProjectsParamsWithHTTPClient(client *http.Client) *CatalogBindProjectsParams {
	return &CatalogBindProjectsParams{
		HTTPClient: client,
	}
}

/* CatalogBindProjectsParams contains all the parameters to send to the API endpoint
   for the catalog bind projects operation.

   Typically these are written to a http.Request.
*/
type CatalogBindProjectsParams struct {

	// Body.
	Body *models.BindProjectsToCatalogCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog bind projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogBindProjectsParams) WithDefaults() *CatalogBindProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog bind projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogBindProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog bind projects params
func (o *CatalogBindProjectsParams) WithTimeout(timeout time.Duration) *CatalogBindProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog bind projects params
func (o *CatalogBindProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog bind projects params
func (o *CatalogBindProjectsParams) WithContext(ctx context.Context) *CatalogBindProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog bind projects params
func (o *CatalogBindProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog bind projects params
func (o *CatalogBindProjectsParams) WithHTTPClient(client *http.Client) *CatalogBindProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog bind projects params
func (o *CatalogBindProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog bind projects params
func (o *CatalogBindProjectsParams) WithBody(body *models.BindProjectsToCatalogCommand) *CatalogBindProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog bind projects params
func (o *CatalogBindProjectsParams) SetBody(body *models.BindProjectsToCatalogCommand) {
	o.Body = body
}

// WithV adds the v to the catalog bind projects params
func (o *CatalogBindProjectsParams) WithV(v string) *CatalogBindProjectsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog bind projects params
func (o *CatalogBindProjectsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogBindProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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