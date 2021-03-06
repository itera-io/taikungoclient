// Code generated by go-swagger; DO NOT EDIT.

package common

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

// NewCommonGetSortingElementsParams creates a new CommonGetSortingElementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommonGetSortingElementsParams() *CommonGetSortingElementsParams {
	return &CommonGetSortingElementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommonGetSortingElementsParamsWithTimeout creates a new CommonGetSortingElementsParams object
// with the ability to set a timeout on a request.
func NewCommonGetSortingElementsParamsWithTimeout(timeout time.Duration) *CommonGetSortingElementsParams {
	return &CommonGetSortingElementsParams{
		timeout: timeout,
	}
}

// NewCommonGetSortingElementsParamsWithContext creates a new CommonGetSortingElementsParams object
// with the ability to set a context for a request.
func NewCommonGetSortingElementsParamsWithContext(ctx context.Context) *CommonGetSortingElementsParams {
	return &CommonGetSortingElementsParams{
		Context: ctx,
	}
}

// NewCommonGetSortingElementsParamsWithHTTPClient creates a new CommonGetSortingElementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommonGetSortingElementsParamsWithHTTPClient(client *http.Client) *CommonGetSortingElementsParams {
	return &CommonGetSortingElementsParams{
		HTTPClient: client,
	}
}

/* CommonGetSortingElementsParams contains all the parameters to send to the API endpoint
   for the common get sorting elements operation.

   Typically these are written to a http.Request.
*/
type CommonGetSortingElementsParams struct {

	// Type.
	Type string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the common get sorting elements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommonGetSortingElementsParams) WithDefaults() *CommonGetSortingElementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the common get sorting elements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommonGetSortingElementsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the common get sorting elements params
func (o *CommonGetSortingElementsParams) WithTimeout(timeout time.Duration) *CommonGetSortingElementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the common get sorting elements params
func (o *CommonGetSortingElementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the common get sorting elements params
func (o *CommonGetSortingElementsParams) WithContext(ctx context.Context) *CommonGetSortingElementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the common get sorting elements params
func (o *CommonGetSortingElementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the common get sorting elements params
func (o *CommonGetSortingElementsParams) WithHTTPClient(client *http.Client) *CommonGetSortingElementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the common get sorting elements params
func (o *CommonGetSortingElementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the common get sorting elements params
func (o *CommonGetSortingElementsParams) WithType(typeVar string) *CommonGetSortingElementsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the common get sorting elements params
func (o *CommonGetSortingElementsParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WithV adds the v to the common get sorting elements params
func (o *CommonGetSortingElementsParams) WithV(v string) *CommonGetSortingElementsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the common get sorting elements params
func (o *CommonGetSortingElementsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CommonGetSortingElementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
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
