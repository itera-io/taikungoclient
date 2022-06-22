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

// NewKubernetesPatchSecretParams creates a new KubernetesPatchSecretParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesPatchSecretParams() *KubernetesPatchSecretParams {
	return &KubernetesPatchSecretParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesPatchSecretParamsWithTimeout creates a new KubernetesPatchSecretParams object
// with the ability to set a timeout on a request.
func NewKubernetesPatchSecretParamsWithTimeout(timeout time.Duration) *KubernetesPatchSecretParams {
	return &KubernetesPatchSecretParams{
		timeout: timeout,
	}
}

// NewKubernetesPatchSecretParamsWithContext creates a new KubernetesPatchSecretParams object
// with the ability to set a context for a request.
func NewKubernetesPatchSecretParamsWithContext(ctx context.Context) *KubernetesPatchSecretParams {
	return &KubernetesPatchSecretParams{
		Context: ctx,
	}
}

// NewKubernetesPatchSecretParamsWithHTTPClient creates a new KubernetesPatchSecretParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesPatchSecretParamsWithHTTPClient(client *http.Client) *KubernetesPatchSecretParams {
	return &KubernetesPatchSecretParams{
		HTTPClient: client,
	}
}

/* KubernetesPatchSecretParams contains all the parameters to send to the API endpoint
   for the kubernetes patch secret operation.

   Typically these are written to a http.Request.
*/
type KubernetesPatchSecretParams struct {

	// Body.
	Body *models.PatchSecretCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes patch secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchSecretParams) WithDefaults() *KubernetesPatchSecretParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes patch secret params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchSecretParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) WithTimeout(timeout time.Duration) *KubernetesPatchSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) WithContext(ctx context.Context) *KubernetesPatchSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) WithHTTPClient(client *http.Client) *KubernetesPatchSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) WithBody(body *models.PatchSecretCommand) *KubernetesPatchSecretParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) SetBody(body *models.PatchSecretCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) WithV(v string) *KubernetesPatchSecretParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes patch secret params
func (o *KubernetesPatchSecretParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesPatchSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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