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

// NewCatalogAvailableVersionsParams creates a new CatalogAvailableVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogAvailableVersionsParams() *CatalogAvailableVersionsParams {
	return &CatalogAvailableVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogAvailableVersionsParamsWithTimeout creates a new CatalogAvailableVersionsParams object
// with the ability to set a timeout on a request.
func NewCatalogAvailableVersionsParamsWithTimeout(timeout time.Duration) *CatalogAvailableVersionsParams {
	return &CatalogAvailableVersionsParams{
		timeout: timeout,
	}
}

// NewCatalogAvailableVersionsParamsWithContext creates a new CatalogAvailableVersionsParams object
// with the ability to set a context for a request.
func NewCatalogAvailableVersionsParamsWithContext(ctx context.Context) *CatalogAvailableVersionsParams {
	return &CatalogAvailableVersionsParams{
		Context: ctx,
	}
}

// NewCatalogAvailableVersionsParamsWithHTTPClient creates a new CatalogAvailableVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogAvailableVersionsParamsWithHTTPClient(client *http.Client) *CatalogAvailableVersionsParams {
	return &CatalogAvailableVersionsParams{
		HTTPClient: client,
	}
}

/* CatalogAvailableVersionsParams contains all the parameters to send to the API endpoint
   for the catalog available versions operation.

   Typically these are written to a http.Request.
*/
type CatalogAvailableVersionsParams struct {

	// Body.
	Body *models.ListCatalogAppAvailableVersionsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog available versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogAvailableVersionsParams) WithDefaults() *CatalogAvailableVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog available versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogAvailableVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog available versions params
func (o *CatalogAvailableVersionsParams) WithTimeout(timeout time.Duration) *CatalogAvailableVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog available versions params
func (o *CatalogAvailableVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog available versions params
func (o *CatalogAvailableVersionsParams) WithContext(ctx context.Context) *CatalogAvailableVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog available versions params
func (o *CatalogAvailableVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog available versions params
func (o *CatalogAvailableVersionsParams) WithHTTPClient(client *http.Client) *CatalogAvailableVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog available versions params
func (o *CatalogAvailableVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog available versions params
func (o *CatalogAvailableVersionsParams) WithBody(body *models.ListCatalogAppAvailableVersionsCommand) *CatalogAvailableVersionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog available versions params
func (o *CatalogAvailableVersionsParams) SetBody(body *models.ListCatalogAppAvailableVersionsCommand) {
	o.Body = body
}

// WithV adds the v to the catalog available versions params
func (o *CatalogAvailableVersionsParams) WithV(v string) *CatalogAvailableVersionsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog available versions params
func (o *CatalogAvailableVersionsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogAvailableVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
