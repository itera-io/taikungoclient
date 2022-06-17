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

// NewCatalogEditCatalogAppVersionParams creates a new CatalogEditCatalogAppVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogEditCatalogAppVersionParams() *CatalogEditCatalogAppVersionParams {
	return &CatalogEditCatalogAppVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogEditCatalogAppVersionParamsWithTimeout creates a new CatalogEditCatalogAppVersionParams object
// with the ability to set a timeout on a request.
func NewCatalogEditCatalogAppVersionParamsWithTimeout(timeout time.Duration) *CatalogEditCatalogAppVersionParams {
	return &CatalogEditCatalogAppVersionParams{
		timeout: timeout,
	}
}

// NewCatalogEditCatalogAppVersionParamsWithContext creates a new CatalogEditCatalogAppVersionParams object
// with the ability to set a context for a request.
func NewCatalogEditCatalogAppVersionParamsWithContext(ctx context.Context) *CatalogEditCatalogAppVersionParams {
	return &CatalogEditCatalogAppVersionParams{
		Context: ctx,
	}
}

// NewCatalogEditCatalogAppVersionParamsWithHTTPClient creates a new CatalogEditCatalogAppVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogEditCatalogAppVersionParamsWithHTTPClient(client *http.Client) *CatalogEditCatalogAppVersionParams {
	return &CatalogEditCatalogAppVersionParams{
		HTTPClient: client,
	}
}

/* CatalogEditCatalogAppVersionParams contains all the parameters to send to the API endpoint
   for the catalog edit catalog app version operation.

   Typically these are written to a http.Request.
*/
type CatalogEditCatalogAppVersionParams struct {

	// Body.
	Body *models.EditCatalogAppVersionCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog edit catalog app version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogEditCatalogAppVersionParams) WithDefaults() *CatalogEditCatalogAppVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog edit catalog app version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogEditCatalogAppVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) WithTimeout(timeout time.Duration) *CatalogEditCatalogAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) WithContext(ctx context.Context) *CatalogEditCatalogAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) WithHTTPClient(client *http.Client) *CatalogEditCatalogAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) WithBody(body *models.EditCatalogAppVersionCommand) *CatalogEditCatalogAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) SetBody(body *models.EditCatalogAppVersionCommand) {
	o.Body = body
}

// WithV adds the v to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) WithV(v string) *CatalogEditCatalogAppVersionParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog edit catalog app version params
func (o *CatalogEditCatalogAppVersionParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogEditCatalogAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
