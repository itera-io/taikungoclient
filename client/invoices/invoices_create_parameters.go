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

	"github.com/itera-io/taikungoclient/models"
)

// NewInvoicesCreateParams creates a new InvoicesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvoicesCreateParams() *InvoicesCreateParams {
	return &InvoicesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvoicesCreateParamsWithTimeout creates a new InvoicesCreateParams object
// with the ability to set a timeout on a request.
func NewInvoicesCreateParamsWithTimeout(timeout time.Duration) *InvoicesCreateParams {
	return &InvoicesCreateParams{
		timeout: timeout,
	}
}

// NewInvoicesCreateParamsWithContext creates a new InvoicesCreateParams object
// with the ability to set a context for a request.
func NewInvoicesCreateParamsWithContext(ctx context.Context) *InvoicesCreateParams {
	return &InvoicesCreateParams{
		Context: ctx,
	}
}

// NewInvoicesCreateParamsWithHTTPClient creates a new InvoicesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvoicesCreateParamsWithHTTPClient(client *http.Client) *InvoicesCreateParams {
	return &InvoicesCreateParams{
		HTTPClient: client,
	}
}

/*
InvoicesCreateParams contains all the parameters to send to the API endpoint

	for the invoices create operation.

	Typically these are written to a http.Request.
*/
type InvoicesCreateParams struct {

	// Body.
	Body *models.CreateInvoiceCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invoices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesCreateParams) WithDefaults() *InvoicesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invoices create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invoices create params
func (o *InvoicesCreateParams) WithTimeout(timeout time.Duration) *InvoicesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoices create params
func (o *InvoicesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoices create params
func (o *InvoicesCreateParams) WithContext(ctx context.Context) *InvoicesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoices create params
func (o *InvoicesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invoices create params
func (o *InvoicesCreateParams) WithHTTPClient(client *http.Client) *InvoicesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invoices create params
func (o *InvoicesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invoices create params
func (o *InvoicesCreateParams) WithBody(body *models.CreateInvoiceCommand) *InvoicesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invoices create params
func (o *InvoicesCreateParams) SetBody(body *models.CreateInvoiceCommand) {
	o.Body = body
}

// WithV adds the v to the invoices create params
func (o *InvoicesCreateParams) WithV(v string) *InvoicesCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the invoices create params
func (o *InvoicesCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *InvoicesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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