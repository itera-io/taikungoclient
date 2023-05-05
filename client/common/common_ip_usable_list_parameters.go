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

	"github.com/itera-io/taikungoclient/models"
)

// NewCommonIPUsableListParams creates a new CommonIPUsableListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommonIPUsableListParams() *CommonIPUsableListParams {
	return &CommonIPUsableListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommonIPUsableListParamsWithTimeout creates a new CommonIPUsableListParams object
// with the ability to set a timeout on a request.
func NewCommonIPUsableListParamsWithTimeout(timeout time.Duration) *CommonIPUsableListParams {
	return &CommonIPUsableListParams{
		timeout: timeout,
	}
}

// NewCommonIPUsableListParamsWithContext creates a new CommonIPUsableListParams object
// with the ability to set a context for a request.
func NewCommonIPUsableListParamsWithContext(ctx context.Context) *CommonIPUsableListParams {
	return &CommonIPUsableListParams{
		Context: ctx,
	}
}

// NewCommonIPUsableListParamsWithHTTPClient creates a new CommonIPUsableListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommonIPUsableListParamsWithHTTPClient(client *http.Client) *CommonIPUsableListParams {
	return &CommonIPUsableListParams{
		HTTPClient: client,
	}
}

/*
CommonIPUsableListParams contains all the parameters to send to the API endpoint

	for the common Ip usable list operation.

	Typically these are written to a http.Request.
*/
type CommonIPUsableListParams struct {

	// Body.
	Body *models.IPListCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the common Ip usable list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommonIPUsableListParams) WithDefaults() *CommonIPUsableListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the common Ip usable list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommonIPUsableListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the common Ip usable list params
func (o *CommonIPUsableListParams) WithTimeout(timeout time.Duration) *CommonIPUsableListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the common Ip usable list params
func (o *CommonIPUsableListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the common Ip usable list params
func (o *CommonIPUsableListParams) WithContext(ctx context.Context) *CommonIPUsableListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the common Ip usable list params
func (o *CommonIPUsableListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the common Ip usable list params
func (o *CommonIPUsableListParams) WithHTTPClient(client *http.Client) *CommonIPUsableListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the common Ip usable list params
func (o *CommonIPUsableListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the common Ip usable list params
func (o *CommonIPUsableListParams) WithBody(body *models.IPListCommand) *CommonIPUsableListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the common Ip usable list params
func (o *CommonIPUsableListParams) SetBody(body *models.IPListCommand) {
	o.Body = body
}

// WithV adds the v to the common Ip usable list params
func (o *CommonIPUsableListParams) WithV(v string) *CommonIPUsableListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the common Ip usable list params
func (o *CommonIPUsableListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CommonIPUsableListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
