// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewPaymentPayInvoiceParams creates a new PaymentPayInvoiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentPayInvoiceParams() *PaymentPayInvoiceParams {
	return &PaymentPayInvoiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentPayInvoiceParamsWithTimeout creates a new PaymentPayInvoiceParams object
// with the ability to set a timeout on a request.
func NewPaymentPayInvoiceParamsWithTimeout(timeout time.Duration) *PaymentPayInvoiceParams {
	return &PaymentPayInvoiceParams{
		timeout: timeout,
	}
}

// NewPaymentPayInvoiceParamsWithContext creates a new PaymentPayInvoiceParams object
// with the ability to set a context for a request.
func NewPaymentPayInvoiceParamsWithContext(ctx context.Context) *PaymentPayInvoiceParams {
	return &PaymentPayInvoiceParams{
		Context: ctx,
	}
}

// NewPaymentPayInvoiceParamsWithHTTPClient creates a new PaymentPayInvoiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentPayInvoiceParamsWithHTTPClient(client *http.Client) *PaymentPayInvoiceParams {
	return &PaymentPayInvoiceParams{
		HTTPClient: client,
	}
}

/* PaymentPayInvoiceParams contains all the parameters to send to the API endpoint
   for the payment pay invoice operation.

   Typically these are written to a http.Request.
*/
type PaymentPayInvoiceParams struct {

	// Body.
	Body *models.PayInvoiceCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment pay invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentPayInvoiceParams) WithDefaults() *PaymentPayInvoiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment pay invoice params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentPayInvoiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment pay invoice params
func (o *PaymentPayInvoiceParams) WithTimeout(timeout time.Duration) *PaymentPayInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment pay invoice params
func (o *PaymentPayInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment pay invoice params
func (o *PaymentPayInvoiceParams) WithContext(ctx context.Context) *PaymentPayInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment pay invoice params
func (o *PaymentPayInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment pay invoice params
func (o *PaymentPayInvoiceParams) WithHTTPClient(client *http.Client) *PaymentPayInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment pay invoice params
func (o *PaymentPayInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the payment pay invoice params
func (o *PaymentPayInvoiceParams) WithBody(body *models.PayInvoiceCommand) *PaymentPayInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the payment pay invoice params
func (o *PaymentPayInvoiceParams) SetBody(body *models.PayInvoiceCommand) {
	o.Body = body
}

// WithV adds the v to the payment pay invoice params
func (o *PaymentPayInvoiceParams) WithV(v string) *PaymentPayInvoiceParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the payment pay invoice params
func (o *PaymentPayInvoiceParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentPayInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
