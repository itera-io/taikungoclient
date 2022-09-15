// Code generated by go-swagger; DO NOT EDIT.

package ticket

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

// NewTicketTransferParams creates a new TicketTransferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTicketTransferParams() *TicketTransferParams {
	return &TicketTransferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTicketTransferParamsWithTimeout creates a new TicketTransferParams object
// with the ability to set a timeout on a request.
func NewTicketTransferParamsWithTimeout(timeout time.Duration) *TicketTransferParams {
	return &TicketTransferParams{
		timeout: timeout,
	}
}

// NewTicketTransferParamsWithContext creates a new TicketTransferParams object
// with the ability to set a context for a request.
func NewTicketTransferParamsWithContext(ctx context.Context) *TicketTransferParams {
	return &TicketTransferParams{
		Context: ctx,
	}
}

// NewTicketTransferParamsWithHTTPClient creates a new TicketTransferParams object
// with the ability to set a custom HTTPClient for a request.
func NewTicketTransferParamsWithHTTPClient(client *http.Client) *TicketTransferParams {
	return &TicketTransferParams{
		HTTPClient: client,
	}
}

/*
TicketTransferParams contains all the parameters to send to the API endpoint

	for the ticket transfer operation.

	Typically these are written to a http.Request.
*/
type TicketTransferParams struct {

	// Body.
	Body *models.TransferTicketCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ticket transfer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TicketTransferParams) WithDefaults() *TicketTransferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ticket transfer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TicketTransferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ticket transfer params
func (o *TicketTransferParams) WithTimeout(timeout time.Duration) *TicketTransferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket transfer params
func (o *TicketTransferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket transfer params
func (o *TicketTransferParams) WithContext(ctx context.Context) *TicketTransferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket transfer params
func (o *TicketTransferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ticket transfer params
func (o *TicketTransferParams) WithHTTPClient(client *http.Client) *TicketTransferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket transfer params
func (o *TicketTransferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ticket transfer params
func (o *TicketTransferParams) WithBody(body *models.TransferTicketCommand) *TicketTransferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ticket transfer params
func (o *TicketTransferParams) SetBody(body *models.TransferTicketCommand) {
	o.Body = body
}

// WithV adds the v to the ticket transfer params
func (o *TicketTransferParams) WithV(v string) *TicketTransferParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ticket transfer params
func (o *TicketTransferParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *TicketTransferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
