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

	"github.com/itera-io/taikungoclient/models"
)

// NewKubernetesDescribeStorageClassParams creates a new KubernetesDescribeStorageClassParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeStorageClassParams() *KubernetesDescribeStorageClassParams {
	return &KubernetesDescribeStorageClassParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeStorageClassParamsWithTimeout creates a new KubernetesDescribeStorageClassParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeStorageClassParamsWithTimeout(timeout time.Duration) *KubernetesDescribeStorageClassParams {
	return &KubernetesDescribeStorageClassParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeStorageClassParamsWithContext creates a new KubernetesDescribeStorageClassParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeStorageClassParamsWithContext(ctx context.Context) *KubernetesDescribeStorageClassParams {
	return &KubernetesDescribeStorageClassParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeStorageClassParamsWithHTTPClient creates a new KubernetesDescribeStorageClassParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeStorageClassParamsWithHTTPClient(client *http.Client) *KubernetesDescribeStorageClassParams {
	return &KubernetesDescribeStorageClassParams{
		HTTPClient: client,
	}
}

/*
KubernetesDescribeStorageClassParams contains all the parameters to send to the API endpoint

	for the kubernetes describe storage class operation.

	Typically these are written to a http.Request.
*/
type KubernetesDescribeStorageClassParams struct {

	// Body.
	Body *models.DescribeStorageClassCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe storage class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeStorageClassParams) WithDefaults() *KubernetesDescribeStorageClassParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe storage class params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeStorageClassParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) WithTimeout(timeout time.Duration) *KubernetesDescribeStorageClassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) WithContext(ctx context.Context) *KubernetesDescribeStorageClassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) WithHTTPClient(client *http.Client) *KubernetesDescribeStorageClassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) WithBody(body *models.DescribeStorageClassCommand) *KubernetesDescribeStorageClassParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) SetBody(body *models.DescribeStorageClassCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) WithV(v string) *KubernetesDescribeStorageClassParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe storage class params
func (o *KubernetesDescribeStorageClassParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeStorageClassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
