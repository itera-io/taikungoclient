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

// NewKubernetesRestartStsParams creates a new KubernetesRestartStsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesRestartStsParams() *KubernetesRestartStsParams {
	return &KubernetesRestartStsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesRestartStsParamsWithTimeout creates a new KubernetesRestartStsParams object
// with the ability to set a timeout on a request.
func NewKubernetesRestartStsParamsWithTimeout(timeout time.Duration) *KubernetesRestartStsParams {
	return &KubernetesRestartStsParams{
		timeout: timeout,
	}
}

// NewKubernetesRestartStsParamsWithContext creates a new KubernetesRestartStsParams object
// with the ability to set a context for a request.
func NewKubernetesRestartStsParamsWithContext(ctx context.Context) *KubernetesRestartStsParams {
	return &KubernetesRestartStsParams{
		Context: ctx,
	}
}

// NewKubernetesRestartStsParamsWithHTTPClient creates a new KubernetesRestartStsParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesRestartStsParamsWithHTTPClient(client *http.Client) *KubernetesRestartStsParams {
	return &KubernetesRestartStsParams{
		HTTPClient: client,
	}
}

/*
KubernetesRestartStsParams contains all the parameters to send to the API endpoint

	for the kubernetes restart sts operation.

	Typically these are written to a http.Request.
*/
type KubernetesRestartStsParams struct {

	// Body.
	Body KubernetesRestartStsBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes restart sts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesRestartStsParams) WithDefaults() *KubernetesRestartStsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes restart sts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesRestartStsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) WithTimeout(timeout time.Duration) *KubernetesRestartStsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) WithContext(ctx context.Context) *KubernetesRestartStsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) WithHTTPClient(client *http.Client) *KubernetesRestartStsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) WithBody(body KubernetesRestartStsBody) *KubernetesRestartStsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) SetBody(body KubernetesRestartStsBody) {
	o.Body = body
}

// WithV adds the v to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) WithV(v string) *KubernetesRestartStsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes restart sts params
func (o *KubernetesRestartStsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesRestartStsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
