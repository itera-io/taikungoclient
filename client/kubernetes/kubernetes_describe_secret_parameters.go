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

// NewKubernetesDescribeSecretParams creates a new KubernetesDescribeSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeSecretParams() *KubernetesDescribeSecretParams {
	return &KubernetesDescribeSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeSecretParamsWithTimeout creates a new KubernetesDescribeSecretParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeSecretParamsWithTimeout(timeout time.Duration) *KubernetesDescribeSecretParams {
	return &KubernetesDescribeSecretParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeSecretParamsWithContext creates a new KubernetesDescribeSecretParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeSecretParamsWithContext(ctx context.Context) *KubernetesDescribeSecretParams {
	return &KubernetesDescribeSecretParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeSecretParamsWithHTTPClient creates a new KubernetesDescribeSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeSecretParamsWithHTTPClient(client *http.Client) *KubernetesDescribeSecretParams {
	return &KubernetesDescribeSecretParams{
		HTTPClient: client,
	}
}

/*
KubernetesDescribeSecretParams contains all the parameters to send to the API endpoint

	for the kubernetes describe secret operation.

	Typically these are written to a http.Request.
*/
type KubernetesDescribeSecretParams struct {

	// Body.
	Body KubernetesDescribeSecretBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeSecretParams) WithDefaults() *KubernetesDescribeSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) WithTimeout(timeout time.Duration) *KubernetesDescribeSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) WithContext(ctx context.Context) *KubernetesDescribeSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) WithHTTPClient(client *http.Client) *KubernetesDescribeSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) WithBody(body KubernetesDescribeSecretBody) *KubernetesDescribeSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) SetBody(body KubernetesDescribeSecretBody) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) WithV(v string) *KubernetesDescribeSecretParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe secret params
func (o *KubernetesDescribeSecretParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
