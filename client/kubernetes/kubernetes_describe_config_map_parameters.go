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

// NewKubernetesDescribeConfigMapParams creates a new KubernetesDescribeConfigMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeConfigMapParams() *KubernetesDescribeConfigMapParams {
	return &KubernetesDescribeConfigMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeConfigMapParamsWithTimeout creates a new KubernetesDescribeConfigMapParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeConfigMapParamsWithTimeout(timeout time.Duration) *KubernetesDescribeConfigMapParams {
	return &KubernetesDescribeConfigMapParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeConfigMapParamsWithContext creates a new KubernetesDescribeConfigMapParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeConfigMapParamsWithContext(ctx context.Context) *KubernetesDescribeConfigMapParams {
	return &KubernetesDescribeConfigMapParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeConfigMapParamsWithHTTPClient creates a new KubernetesDescribeConfigMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeConfigMapParamsWithHTTPClient(client *http.Client) *KubernetesDescribeConfigMapParams {
	return &KubernetesDescribeConfigMapParams{
		HTTPClient: client,
	}
}

/* KubernetesDescribeConfigMapParams contains all the parameters to send to the API endpoint
   for the kubernetes describe config map operation.

   Typically these are written to a http.Request.
*/
type KubernetesDescribeConfigMapParams struct {

	// Body.
	Body *models.DescribeConfigMapCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe config map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeConfigMapParams) WithDefaults() *KubernetesDescribeConfigMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe config map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeConfigMapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) WithTimeout(timeout time.Duration) *KubernetesDescribeConfigMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) WithContext(ctx context.Context) *KubernetesDescribeConfigMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) WithHTTPClient(client *http.Client) *KubernetesDescribeConfigMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) WithBody(body *models.DescribeConfigMapCommand) *KubernetesDescribeConfigMapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) SetBody(body *models.DescribeConfigMapCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) WithV(v string) *KubernetesDescribeConfigMapParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe config map params
func (o *KubernetesDescribeConfigMapParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeConfigMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
