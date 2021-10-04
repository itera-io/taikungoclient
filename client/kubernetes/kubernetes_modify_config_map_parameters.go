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

// NewKubernetesModifyConfigMapParams creates a new KubernetesModifyConfigMapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesModifyConfigMapParams() *KubernetesModifyConfigMapParams {
	return &KubernetesModifyConfigMapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesModifyConfigMapParamsWithTimeout creates a new KubernetesModifyConfigMapParams object
// with the ability to set a timeout on a request.
func NewKubernetesModifyConfigMapParamsWithTimeout(timeout time.Duration) *KubernetesModifyConfigMapParams {
	return &KubernetesModifyConfigMapParams{
		timeout: timeout,
	}
}

// NewKubernetesModifyConfigMapParamsWithContext creates a new KubernetesModifyConfigMapParams object
// with the ability to set a context for a request.
func NewKubernetesModifyConfigMapParamsWithContext(ctx context.Context) *KubernetesModifyConfigMapParams {
	return &KubernetesModifyConfigMapParams{
		Context: ctx,
	}
}

// NewKubernetesModifyConfigMapParamsWithHTTPClient creates a new KubernetesModifyConfigMapParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesModifyConfigMapParamsWithHTTPClient(client *http.Client) *KubernetesModifyConfigMapParams {
	return &KubernetesModifyConfigMapParams{
		HTTPClient: client,
	}
}

/* KubernetesModifyConfigMapParams contains all the parameters to send to the API endpoint
   for the kubernetes modify config map operation.

   Typically these are written to a http.Request.
*/
type KubernetesModifyConfigMapParams struct {

	// Body.
	Body *models.ModifyConfigMapCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes modify config map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesModifyConfigMapParams) WithDefaults() *KubernetesModifyConfigMapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes modify config map params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesModifyConfigMapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) WithTimeout(timeout time.Duration) *KubernetesModifyConfigMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) WithContext(ctx context.Context) *KubernetesModifyConfigMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) WithHTTPClient(client *http.Client) *KubernetesModifyConfigMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) WithBody(body *models.ModifyConfigMapCommand) *KubernetesModifyConfigMapParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) SetBody(body *models.ModifyConfigMapCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) WithV(v string) *KubernetesModifyConfigMapParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes modify config map params
func (o *KubernetesModifyConfigMapParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesModifyConfigMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
