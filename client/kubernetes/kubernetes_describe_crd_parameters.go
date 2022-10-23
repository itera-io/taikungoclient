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

// NewKubernetesDescribeCrdParams creates a new KubernetesDescribeCrdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeCrdParams() *KubernetesDescribeCrdParams {
	return &KubernetesDescribeCrdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeCrdParamsWithTimeout creates a new KubernetesDescribeCrdParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeCrdParamsWithTimeout(timeout time.Duration) *KubernetesDescribeCrdParams {
	return &KubernetesDescribeCrdParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeCrdParamsWithContext creates a new KubernetesDescribeCrdParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeCrdParamsWithContext(ctx context.Context) *KubernetesDescribeCrdParams {
	return &KubernetesDescribeCrdParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeCrdParamsWithHTTPClient creates a new KubernetesDescribeCrdParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeCrdParamsWithHTTPClient(client *http.Client) *KubernetesDescribeCrdParams {
	return &KubernetesDescribeCrdParams{
		HTTPClient: client,
	}
}

/*
KubernetesDescribeCrdParams contains all the parameters to send to the API endpoint

	for the kubernetes describe crd operation.

	Typically these are written to a http.Request.
*/
type KubernetesDescribeCrdParams struct {

	// Body.
	Body *models.DescribeCrdCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe crd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeCrdParams) WithDefaults() *KubernetesDescribeCrdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe crd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeCrdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) WithTimeout(timeout time.Duration) *KubernetesDescribeCrdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) WithContext(ctx context.Context) *KubernetesDescribeCrdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) WithHTTPClient(client *http.Client) *KubernetesDescribeCrdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) WithBody(body *models.DescribeCrdCommand) *KubernetesDescribeCrdParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) SetBody(body *models.DescribeCrdCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) WithV(v string) *KubernetesDescribeCrdParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe crd params
func (o *KubernetesDescribeCrdParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeCrdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
