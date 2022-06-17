// Code generated by go-swagger; DO NOT EDIT.

package dns_servers

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
)

// NewDNSServersDeleteParams creates a new DNSServersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDNSServersDeleteParams() *DNSServersDeleteParams {
	return &DNSServersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDNSServersDeleteParamsWithTimeout creates a new DNSServersDeleteParams object
// with the ability to set a timeout on a request.
func NewDNSServersDeleteParamsWithTimeout(timeout time.Duration) *DNSServersDeleteParams {
	return &DNSServersDeleteParams{
		timeout: timeout,
	}
}

// NewDNSServersDeleteParamsWithContext creates a new DNSServersDeleteParams object
// with the ability to set a context for a request.
func NewDNSServersDeleteParamsWithContext(ctx context.Context) *DNSServersDeleteParams {
	return &DNSServersDeleteParams{
		Context: ctx,
	}
}

// NewDNSServersDeleteParamsWithHTTPClient creates a new DNSServersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDNSServersDeleteParamsWithHTTPClient(client *http.Client) *DNSServersDeleteParams {
	return &DNSServersDeleteParams{
		HTTPClient: client,
	}
}

/* DNSServersDeleteParams contains all the parameters to send to the API endpoint
   for the Dns servers delete operation.

   Typically these are written to a http.Request.
*/
type DNSServersDeleteParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Dns servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSServersDeleteParams) WithDefaults() *DNSServersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Dns servers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DNSServersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Dns servers delete params
func (o *DNSServersDeleteParams) WithTimeout(timeout time.Duration) *DNSServersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Dns servers delete params
func (o *DNSServersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Dns servers delete params
func (o *DNSServersDeleteParams) WithContext(ctx context.Context) *DNSServersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Dns servers delete params
func (o *DNSServersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Dns servers delete params
func (o *DNSServersDeleteParams) WithHTTPClient(client *http.Client) *DNSServersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Dns servers delete params
func (o *DNSServersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the Dns servers delete params
func (o *DNSServersDeleteParams) WithID(id int32) *DNSServersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the Dns servers delete params
func (o *DNSServersDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the Dns servers delete params
func (o *DNSServersDeleteParams) WithV(v string) *DNSServersDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the Dns servers delete params
func (o *DNSServersDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *DNSServersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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
