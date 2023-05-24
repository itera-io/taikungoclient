// Code generated by go-swagger; DO NOT EDIT.

package tanzu

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

// NewTanzuStorageListParams creates a new TanzuStorageListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTanzuStorageListParams() *TanzuStorageListParams {
	return &TanzuStorageListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTanzuStorageListParamsWithTimeout creates a new TanzuStorageListParams object
// with the ability to set a timeout on a request.
func NewTanzuStorageListParamsWithTimeout(timeout time.Duration) *TanzuStorageListParams {
	return &TanzuStorageListParams{
		timeout: timeout,
	}
}

// NewTanzuStorageListParamsWithContext creates a new TanzuStorageListParams object
// with the ability to set a context for a request.
func NewTanzuStorageListParamsWithContext(ctx context.Context) *TanzuStorageListParams {
	return &TanzuStorageListParams{
		Context: ctx,
	}
}

// NewTanzuStorageListParamsWithHTTPClient creates a new TanzuStorageListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTanzuStorageListParamsWithHTTPClient(client *http.Client) *TanzuStorageListParams {
	return &TanzuStorageListParams{
		HTTPClient: client,
	}
}

/*
TanzuStorageListParams contains all the parameters to send to the API endpoint

	for the tanzu storage list operation.

	Typically these are written to a http.Request.
*/
type TanzuStorageListParams struct {

	// Body.
	Body *models.TanzuStorageListCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tanzu storage list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TanzuStorageListParams) WithDefaults() *TanzuStorageListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tanzu storage list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TanzuStorageListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tanzu storage list params
func (o *TanzuStorageListParams) WithTimeout(timeout time.Duration) *TanzuStorageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tanzu storage list params
func (o *TanzuStorageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tanzu storage list params
func (o *TanzuStorageListParams) WithContext(ctx context.Context) *TanzuStorageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tanzu storage list params
func (o *TanzuStorageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tanzu storage list params
func (o *TanzuStorageListParams) WithHTTPClient(client *http.Client) *TanzuStorageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tanzu storage list params
func (o *TanzuStorageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tanzu storage list params
func (o *TanzuStorageListParams) WithBody(body *models.TanzuStorageListCommand) *TanzuStorageListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tanzu storage list params
func (o *TanzuStorageListParams) SetBody(body *models.TanzuStorageListCommand) {
	o.Body = body
}

// WithV adds the v to the tanzu storage list params
func (o *TanzuStorageListParams) WithV(v string) *TanzuStorageListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the tanzu storage list params
func (o *TanzuStorageListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *TanzuStorageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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