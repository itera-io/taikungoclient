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

// NewProxmoxStorageListParams creates a new ProxmoxStorageListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProxmoxStorageListParams() *ProxmoxStorageListParams {
	return &ProxmoxStorageListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProxmoxStorageListParamsWithTimeout creates a new ProxmoxStorageListParams object
// with the ability to set a timeout on a request.
func NewProxmoxStorageListParamsWithTimeout(timeout time.Duration) *ProxmoxStorageListParams {
	return &ProxmoxStorageListParams{
		timeout: timeout,
	}
}

// NewProxmoxStorageListParamsWithContext creates a new ProxmoxStorageListParams object
// with the ability to set a context for a request.
func NewProxmoxStorageListParamsWithContext(ctx context.Context) *ProxmoxStorageListParams {
	return &ProxmoxStorageListParams{
		Context: ctx,
	}
}

// NewProxmoxStorageListParamsWithHTTPClient creates a new ProxmoxStorageListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProxmoxStorageListParamsWithHTTPClient(client *http.Client) *ProxmoxStorageListParams {
	return &ProxmoxStorageListParams{
		HTTPClient: client,
	}
}

/*
ProxmoxStorageListParams contains all the parameters to send to the API endpoint

	for the proxmox storage list operation.

	Typically these are written to a http.Request.
*/
type ProxmoxStorageListParams struct {

	// Body.
	Body *models.StorageListCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the proxmox storage list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxStorageListParams) WithDefaults() *ProxmoxStorageListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the proxmox storage list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxStorageListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the proxmox storage list params
func (o *ProxmoxStorageListParams) WithTimeout(timeout time.Duration) *ProxmoxStorageListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the proxmox storage list params
func (o *ProxmoxStorageListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the proxmox storage list params
func (o *ProxmoxStorageListParams) WithContext(ctx context.Context) *ProxmoxStorageListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the proxmox storage list params
func (o *ProxmoxStorageListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the proxmox storage list params
func (o *ProxmoxStorageListParams) WithHTTPClient(client *http.Client) *ProxmoxStorageListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the proxmox storage list params
func (o *ProxmoxStorageListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the proxmox storage list params
func (o *ProxmoxStorageListParams) WithBody(body *models.StorageListCommand) *ProxmoxStorageListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the proxmox storage list params
func (o *ProxmoxStorageListParams) SetBody(body *models.StorageListCommand) {
	o.Body = body
}

// WithV adds the v to the proxmox storage list params
func (o *ProxmoxStorageListParams) WithV(v string) *ProxmoxStorageListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the proxmox storage list params
func (o *ProxmoxStorageListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProxmoxStorageListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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