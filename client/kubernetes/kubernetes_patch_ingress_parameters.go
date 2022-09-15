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

// NewKubernetesPatchIngressParams creates a new KubernetesPatchIngressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesPatchIngressParams() *KubernetesPatchIngressParams {
	return &KubernetesPatchIngressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesPatchIngressParamsWithTimeout creates a new KubernetesPatchIngressParams object
// with the ability to set a timeout on a request.
func NewKubernetesPatchIngressParamsWithTimeout(timeout time.Duration) *KubernetesPatchIngressParams {
	return &KubernetesPatchIngressParams{
		timeout: timeout,
	}
}

// NewKubernetesPatchIngressParamsWithContext creates a new KubernetesPatchIngressParams object
// with the ability to set a context for a request.
func NewKubernetesPatchIngressParamsWithContext(ctx context.Context) *KubernetesPatchIngressParams {
	return &KubernetesPatchIngressParams{
		Context: ctx,
	}
}

// NewKubernetesPatchIngressParamsWithHTTPClient creates a new KubernetesPatchIngressParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesPatchIngressParamsWithHTTPClient(client *http.Client) *KubernetesPatchIngressParams {
	return &KubernetesPatchIngressParams{
		HTTPClient: client,
	}
}

/*
KubernetesPatchIngressParams contains all the parameters to send to the API endpoint

	for the kubernetes patch ingress operation.

	Typically these are written to a http.Request.
*/
type KubernetesPatchIngressParams struct {

	// Body.
	Body *models.PatchIngressCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes patch ingress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchIngressParams) WithDefaults() *KubernetesPatchIngressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes patch ingress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchIngressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) WithTimeout(timeout time.Duration) *KubernetesPatchIngressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) WithContext(ctx context.Context) *KubernetesPatchIngressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) WithHTTPClient(client *http.Client) *KubernetesPatchIngressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) WithBody(body *models.PatchIngressCommand) *KubernetesPatchIngressParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) SetBody(body *models.PatchIngressCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) WithV(v string) *KubernetesPatchIngressParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes patch ingress params
func (o *KubernetesPatchIngressParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesPatchIngressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
