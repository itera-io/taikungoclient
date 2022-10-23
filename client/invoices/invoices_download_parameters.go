// Code generated by go-swagger; DO NOT EDIT.

package invoices

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

// NewInvoicesDownloadParams creates a new InvoicesDownloadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvoicesDownloadParams() *InvoicesDownloadParams {
	return &InvoicesDownloadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvoicesDownloadParamsWithTimeout creates a new InvoicesDownloadParams object
// with the ability to set a timeout on a request.
func NewInvoicesDownloadParamsWithTimeout(timeout time.Duration) *InvoicesDownloadParams {
	return &InvoicesDownloadParams{
		timeout: timeout,
	}
}

// NewInvoicesDownloadParamsWithContext creates a new InvoicesDownloadParams object
// with the ability to set a context for a request.
func NewInvoicesDownloadParamsWithContext(ctx context.Context) *InvoicesDownloadParams {
	return &InvoicesDownloadParams{
		Context: ctx,
	}
}

// NewInvoicesDownloadParamsWithHTTPClient creates a new InvoicesDownloadParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvoicesDownloadParamsWithHTTPClient(client *http.Client) *InvoicesDownloadParams {
	return &InvoicesDownloadParams{
		HTTPClient: client,
	}
}

/*
InvoicesDownloadParams contains all the parameters to send to the API endpoint

	for the invoices download operation.

	Typically these are written to a http.Request.
*/
type InvoicesDownloadParams struct {

	// Body.
	Body InvoicesDownloadBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invoices download params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesDownloadParams) WithDefaults() *InvoicesDownloadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invoices download params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesDownloadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invoices download params
func (o *InvoicesDownloadParams) WithTimeout(timeout time.Duration) *InvoicesDownloadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoices download params
func (o *InvoicesDownloadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoices download params
func (o *InvoicesDownloadParams) WithContext(ctx context.Context) *InvoicesDownloadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoices download params
func (o *InvoicesDownloadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invoices download params
func (o *InvoicesDownloadParams) WithHTTPClient(client *http.Client) *InvoicesDownloadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invoices download params
func (o *InvoicesDownloadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invoices download params
func (o *InvoicesDownloadParams) WithBody(body InvoicesDownloadBody) *InvoicesDownloadParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invoices download params
func (o *InvoicesDownloadParams) SetBody(body InvoicesDownloadBody) {
	o.Body = body
}

// WithV adds the v to the invoices download params
func (o *InvoicesDownloadParams) WithV(v string) *InvoicesDownloadParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the invoices download params
func (o *InvoicesDownloadParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *InvoicesDownloadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
