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

// NewCatalogLockManagerParams creates a new CatalogLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogLockManagerParams() *CatalogLockManagerParams {
	return &CatalogLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogLockManagerParamsWithTimeout creates a new CatalogLockManagerParams object
// with the ability to set a timeout on a request.
func NewCatalogLockManagerParamsWithTimeout(timeout time.Duration) *CatalogLockManagerParams {
	return &CatalogLockManagerParams{
		timeout: timeout,
	}
}

// NewCatalogLockManagerParamsWithContext creates a new CatalogLockManagerParams object
// with the ability to set a context for a request.
func NewCatalogLockManagerParamsWithContext(ctx context.Context) *CatalogLockManagerParams {
	return &CatalogLockManagerParams{
		Context: ctx,
	}
}

// NewCatalogLockManagerParamsWithHTTPClient creates a new CatalogLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogLockManagerParamsWithHTTPClient(client *http.Client) *CatalogLockManagerParams {
	return &CatalogLockManagerParams{
		HTTPClient: client,
	}
}

/* CatalogLockManagerParams contains all the parameters to send to the API endpoint
   for the catalog lock manager operation.

   Typically these are written to a http.Request.
*/
type CatalogLockManagerParams struct {

	// Body.
	Body *models.CatalogLockManagementCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogLockManagerParams) WithDefaults() *CatalogLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog lock manager params
func (o *CatalogLockManagerParams) WithTimeout(timeout time.Duration) *CatalogLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog lock manager params
func (o *CatalogLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog lock manager params
func (o *CatalogLockManagerParams) WithContext(ctx context.Context) *CatalogLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog lock manager params
func (o *CatalogLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog lock manager params
func (o *CatalogLockManagerParams) WithHTTPClient(client *http.Client) *CatalogLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog lock manager params
func (o *CatalogLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog lock manager params
func (o *CatalogLockManagerParams) WithBody(body *models.CatalogLockManagementCommand) *CatalogLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog lock manager params
func (o *CatalogLockManagerParams) SetBody(body *models.CatalogLockManagementCommand) {
	o.Body = body
}

// WithV adds the v to the catalog lock manager params
func (o *CatalogLockManagerParams) WithV(v string) *CatalogLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog lock manager params
func (o *CatalogLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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