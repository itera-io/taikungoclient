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

// NewKubernetesDescribeJobParams creates a new KubernetesDescribeJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeJobParams() *KubernetesDescribeJobParams {
	return &KubernetesDescribeJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeJobParamsWithTimeout creates a new KubernetesDescribeJobParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeJobParamsWithTimeout(timeout time.Duration) *KubernetesDescribeJobParams {
	return &KubernetesDescribeJobParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeJobParamsWithContext creates a new KubernetesDescribeJobParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeJobParamsWithContext(ctx context.Context) *KubernetesDescribeJobParams {
	return &KubernetesDescribeJobParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeJobParamsWithHTTPClient creates a new KubernetesDescribeJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeJobParamsWithHTTPClient(client *http.Client) *KubernetesDescribeJobParams {
	return &KubernetesDescribeJobParams{
		HTTPClient: client,
	}
}

/* KubernetesDescribeJobParams contains all the parameters to send to the API endpoint
   for the kubernetes describe job operation.

   Typically these are written to a http.Request.
*/
type KubernetesDescribeJobParams struct {

	// Body.
	Body *models.DescribeJobCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeJobParams) WithDefaults() *KubernetesDescribeJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) WithTimeout(timeout time.Duration) *KubernetesDescribeJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) WithContext(ctx context.Context) *KubernetesDescribeJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) WithHTTPClient(client *http.Client) *KubernetesDescribeJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) WithBody(body *models.DescribeJobCommand) *KubernetesDescribeJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) SetBody(body *models.DescribeJobCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) WithV(v string) *KubernetesDescribeJobParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe job params
func (o *KubernetesDescribeJobParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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