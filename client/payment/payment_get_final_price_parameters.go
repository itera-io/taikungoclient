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

// NewPaymentGetFinalPriceParams creates a new PaymentGetFinalPriceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPaymentGetFinalPriceParams() *PaymentGetFinalPriceParams {
	return &PaymentGetFinalPriceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPaymentGetFinalPriceParamsWithTimeout creates a new PaymentGetFinalPriceParams object
// with the ability to set a timeout on a request.
func NewPaymentGetFinalPriceParamsWithTimeout(timeout time.Duration) *PaymentGetFinalPriceParams {
	return &PaymentGetFinalPriceParams{
		timeout: timeout,
	}
}

// NewPaymentGetFinalPriceParamsWithContext creates a new PaymentGetFinalPriceParams object
// with the ability to set a context for a request.
func NewPaymentGetFinalPriceParamsWithContext(ctx context.Context) *PaymentGetFinalPriceParams {
	return &PaymentGetFinalPriceParams{
		Context: ctx,
	}
}

// NewPaymentGetFinalPriceParamsWithHTTPClient creates a new PaymentGetFinalPriceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPaymentGetFinalPriceParamsWithHTTPClient(client *http.Client) *PaymentGetFinalPriceParams {
	return &PaymentGetFinalPriceParams{
		HTTPClient: client,
	}
}

/* PaymentGetFinalPriceParams contains all the parameters to send to the API endpoint
   for the payment get final price operation.

   Typically these are written to a http.Request.
*/
type PaymentGetFinalPriceParams struct {

	// Body.
	Body *models.FinalPriceCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the payment get final price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGetFinalPriceParams) WithDefaults() *PaymentGetFinalPriceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the payment get final price params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PaymentGetFinalPriceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the payment get final price params
func (o *PaymentGetFinalPriceParams) WithTimeout(timeout time.Duration) *PaymentGetFinalPriceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the payment get final price params
func (o *PaymentGetFinalPriceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the payment get final price params
func (o *PaymentGetFinalPriceParams) WithContext(ctx context.Context) *PaymentGetFinalPriceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the payment get final price params
func (o *PaymentGetFinalPriceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the payment get final price params
func (o *PaymentGetFinalPriceParams) WithHTTPClient(client *http.Client) *PaymentGetFinalPriceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the payment get final price params
func (o *PaymentGetFinalPriceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the payment get final price params
func (o *PaymentGetFinalPriceParams) WithBody(body *models.FinalPriceCommand) *PaymentGetFinalPriceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the payment get final price params
func (o *PaymentGetFinalPriceParams) SetBody(body *models.FinalPriceCommand) {
	o.Body = body
}

// WithV adds the v to the payment get final price params
func (o *PaymentGetFinalPriceParams) WithV(v string) *PaymentGetFinalPriceParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the payment get final price params
func (o *PaymentGetFinalPriceParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PaymentGetFinalPriceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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