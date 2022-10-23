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

// NewCatalogEditCatalogParams creates a new CatalogEditCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogEditCatalogParams() *CatalogEditCatalogParams {
	return &CatalogEditCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogEditCatalogParamsWithTimeout creates a new CatalogEditCatalogParams object
// with the ability to set a timeout on a request.
func NewCatalogEditCatalogParamsWithTimeout(timeout time.Duration) *CatalogEditCatalogParams {
	return &CatalogEditCatalogParams{
		timeout: timeout,
	}
}

// NewCatalogEditCatalogParamsWithContext creates a new CatalogEditCatalogParams object
// with the ability to set a context for a request.
func NewCatalogEditCatalogParamsWithContext(ctx context.Context) *CatalogEditCatalogParams {
	return &CatalogEditCatalogParams{
		Context: ctx,
	}
}

// NewCatalogEditCatalogParamsWithHTTPClient creates a new CatalogEditCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogEditCatalogParamsWithHTTPClient(client *http.Client) *CatalogEditCatalogParams {
	return &CatalogEditCatalogParams{
		HTTPClient: client,
	}
}

/*
CatalogEditCatalogParams contains all the parameters to send to the API endpoint

	for the catalog edit catalog operation.

	Typically these are written to a http.Request.
*/
type CatalogEditCatalogParams struct {

	// Body.
	Body *models.EditCatalogCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog edit catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogEditCatalogParams) WithDefaults() *CatalogEditCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog edit catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogEditCatalogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog edit catalog params
func (o *CatalogEditCatalogParams) WithTimeout(timeout time.Duration) *CatalogEditCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog edit catalog params
func (o *CatalogEditCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog edit catalog params
func (o *CatalogEditCatalogParams) WithContext(ctx context.Context) *CatalogEditCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog edit catalog params
func (o *CatalogEditCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog edit catalog params
func (o *CatalogEditCatalogParams) WithHTTPClient(client *http.Client) *CatalogEditCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog edit catalog params
func (o *CatalogEditCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog edit catalog params
func (o *CatalogEditCatalogParams) WithBody(body *models.EditCatalogCommand) *CatalogEditCatalogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog edit catalog params
func (o *CatalogEditCatalogParams) SetBody(body *models.EditCatalogCommand) {
	o.Body = body
}

// WithV adds the v to the catalog edit catalog params
func (o *CatalogEditCatalogParams) WithV(v string) *CatalogEditCatalogParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the catalog edit catalog params
func (o *CatalogEditCatalogParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogEditCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
