// Code generated by go-swagger; DO NOT EDIT.

package proxmox

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

// NewProxmoxUpdateHypervisorsParams creates a new ProxmoxUpdateHypervisorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProxmoxUpdateHypervisorsParams() *ProxmoxUpdateHypervisorsParams {
	return &ProxmoxUpdateHypervisorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProxmoxUpdateHypervisorsParamsWithTimeout creates a new ProxmoxUpdateHypervisorsParams object
// with the ability to set a timeout on a request.
func NewProxmoxUpdateHypervisorsParamsWithTimeout(timeout time.Duration) *ProxmoxUpdateHypervisorsParams {
	return &ProxmoxUpdateHypervisorsParams{
		timeout: timeout,
	}
}

// NewProxmoxUpdateHypervisorsParamsWithContext creates a new ProxmoxUpdateHypervisorsParams object
// with the ability to set a context for a request.
func NewProxmoxUpdateHypervisorsParamsWithContext(ctx context.Context) *ProxmoxUpdateHypervisorsParams {
	return &ProxmoxUpdateHypervisorsParams{
		Context: ctx,
	}
}

// NewProxmoxUpdateHypervisorsParamsWithHTTPClient creates a new ProxmoxUpdateHypervisorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewProxmoxUpdateHypervisorsParamsWithHTTPClient(client *http.Client) *ProxmoxUpdateHypervisorsParams {
	return &ProxmoxUpdateHypervisorsParams{
		HTTPClient: client,
	}
}

/*
ProxmoxUpdateHypervisorsParams contains all the parameters to send to the API endpoint

	for the proxmox update hypervisors operation.

	Typically these are written to a http.Request.
*/
type ProxmoxUpdateHypervisorsParams struct {

	// Body.
	Body *models.UpdateHypervisorsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the proxmox update hypervisors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxUpdateHypervisorsParams) WithDefaults() *ProxmoxUpdateHypervisorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the proxmox update hypervisors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxUpdateHypervisorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) WithTimeout(timeout time.Duration) *ProxmoxUpdateHypervisorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) WithContext(ctx context.Context) *ProxmoxUpdateHypervisorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) WithHTTPClient(client *http.Client) *ProxmoxUpdateHypervisorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) WithBody(body *models.UpdateHypervisorsCommand) *ProxmoxUpdateHypervisorsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) SetBody(body *models.UpdateHypervisorsCommand) {
	o.Body = body
}

// WithV adds the v to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) WithV(v string) *ProxmoxUpdateHypervisorsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the proxmox update hypervisors params
func (o *ProxmoxUpdateHypervisorsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProxmoxUpdateHypervisorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
