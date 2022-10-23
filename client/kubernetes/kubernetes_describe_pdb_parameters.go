// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewKubernetesDescribePdbParams creates a new KubernetesDescribePdbParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribePdbParams() *KubernetesDescribePdbParams {
	return &KubernetesDescribePdbParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribePdbParamsWithTimeout creates a new KubernetesDescribePdbParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribePdbParamsWithTimeout(timeout time.Duration) *KubernetesDescribePdbParams {
	return &KubernetesDescribePdbParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribePdbParamsWithContext creates a new KubernetesDescribePdbParams object
// with the ability to set a context for a request.
func NewKubernetesDescribePdbParamsWithContext(ctx context.Context) *KubernetesDescribePdbParams {
	return &KubernetesDescribePdbParams{
		Context: ctx,
	}
}

// NewKubernetesDescribePdbParamsWithHTTPClient creates a new KubernetesDescribePdbParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribePdbParamsWithHTTPClient(client *http.Client) *KubernetesDescribePdbParams {
	return &KubernetesDescribePdbParams{
		HTTPClient: client,
	}
}

/*
KubernetesDescribePdbParams contains all the parameters to send to the API endpoint

	for the kubernetes describe pdb operation.

	Typically these are written to a http.Request.
*/
type KubernetesDescribePdbParams struct {

	// Body.
	Body KubernetesDescribePdbBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe pdb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribePdbParams) WithDefaults() *KubernetesDescribePdbParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe pdb params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribePdbParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) WithTimeout(timeout time.Duration) *KubernetesDescribePdbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) WithContext(ctx context.Context) *KubernetesDescribePdbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) WithHTTPClient(client *http.Client) *KubernetesDescribePdbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) WithBody(body KubernetesDescribePdbBody) *KubernetesDescribePdbParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) SetBody(body KubernetesDescribePdbBody) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) WithV(v string) *KubernetesDescribePdbParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe pdb params
func (o *KubernetesDescribePdbParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribePdbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
