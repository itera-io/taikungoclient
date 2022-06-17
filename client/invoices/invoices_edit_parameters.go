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
	"github.com/go-openapi/swag"

	"github.com/itera-io/taikungoclient/models"
)

// NewInvoicesEditParams creates a new InvoicesEditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvoicesEditParams() *InvoicesEditParams {
	return &InvoicesEditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvoicesEditParamsWithTimeout creates a new InvoicesEditParams object
// with the ability to set a timeout on a request.
func NewInvoicesEditParamsWithTimeout(timeout time.Duration) *InvoicesEditParams {
	return &InvoicesEditParams{
		timeout: timeout,
	}
}

// NewInvoicesEditParamsWithContext creates a new InvoicesEditParams object
// with the ability to set a context for a request.
func NewInvoicesEditParamsWithContext(ctx context.Context) *InvoicesEditParams {
	return &InvoicesEditParams{
		Context: ctx,
	}
}

// NewInvoicesEditParamsWithHTTPClient creates a new InvoicesEditParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvoicesEditParamsWithHTTPClient(client *http.Client) *InvoicesEditParams {
	return &InvoicesEditParams{
		HTTPClient: client,
	}
}

/* InvoicesEditParams contains all the parameters to send to the API endpoint
   for the invoices edit operation.

   Typically these are written to a http.Request.
*/
type InvoicesEditParams struct {

	// Body.
	Body *models.UpdateInvoiceDto

	// InvoiceID.
	//
	// Format: int32
	InvoiceID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invoices edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesEditParams) WithDefaults() *InvoicesEditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invoices edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvoicesEditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invoices edit params
func (o *InvoicesEditParams) WithTimeout(timeout time.Duration) *InvoicesEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invoices edit params
func (o *InvoicesEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invoices edit params
func (o *InvoicesEditParams) WithContext(ctx context.Context) *InvoicesEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invoices edit params
func (o *InvoicesEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invoices edit params
func (o *InvoicesEditParams) WithHTTPClient(client *http.Client) *InvoicesEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invoices edit params
func (o *InvoicesEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invoices edit params
func (o *InvoicesEditParams) WithBody(body *models.UpdateInvoiceDto) *InvoicesEditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invoices edit params
func (o *InvoicesEditParams) SetBody(body *models.UpdateInvoiceDto) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the invoices edit params
func (o *InvoicesEditParams) WithInvoiceID(invoiceID int32) *InvoicesEditParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the invoices edit params
func (o *InvoicesEditParams) SetInvoiceID(invoiceID int32) {
	o.InvoiceID = invoiceID
}

// WithV adds the v to the invoices edit params
func (o *InvoicesEditParams) WithV(v string) *InvoicesEditParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the invoices edit params
func (o *InvoicesEditParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *InvoicesEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", swag.FormatInt32(o.InvoiceID)); err != nil {
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
