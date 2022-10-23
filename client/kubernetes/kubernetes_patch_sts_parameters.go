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

// NewKubernetesPatchStsParams creates a new KubernetesPatchStsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesPatchStsParams() *KubernetesPatchStsParams {
	return &KubernetesPatchStsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesPatchStsParamsWithTimeout creates a new KubernetesPatchStsParams object
// with the ability to set a timeout on a request.
func NewKubernetesPatchStsParamsWithTimeout(timeout time.Duration) *KubernetesPatchStsParams {
	return &KubernetesPatchStsParams{
		timeout: timeout,
	}
}

// NewKubernetesPatchStsParamsWithContext creates a new KubernetesPatchStsParams object
// with the ability to set a context for a request.
func NewKubernetesPatchStsParamsWithContext(ctx context.Context) *KubernetesPatchStsParams {
	return &KubernetesPatchStsParams{
		Context: ctx,
	}
}

// NewKubernetesPatchStsParamsWithHTTPClient creates a new KubernetesPatchStsParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesPatchStsParamsWithHTTPClient(client *http.Client) *KubernetesPatchStsParams {
	return &KubernetesPatchStsParams{
		HTTPClient: client,
	}
}

/*
KubernetesPatchStsParams contains all the parameters to send to the API endpoint

	for the kubernetes patch sts operation.

	Typically these are written to a http.Request.
*/
type KubernetesPatchStsParams struct {

	// Body.
	Body KubernetesPatchStsBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes patch sts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchStsParams) WithDefaults() *KubernetesPatchStsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes patch sts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchStsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) WithTimeout(timeout time.Duration) *KubernetesPatchStsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) WithContext(ctx context.Context) *KubernetesPatchStsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) WithHTTPClient(client *http.Client) *KubernetesPatchStsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) WithBody(body KubernetesPatchStsBody) *KubernetesPatchStsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) SetBody(body KubernetesPatchStsBody) {
	o.Body = body
}

// WithV adds the v to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) WithV(v string) *KubernetesPatchStsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes patch sts params
func (o *KubernetesPatchStsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesPatchStsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
