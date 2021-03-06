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

// NewKubernetesGetSupportedListParams creates a new KubernetesGetSupportedListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesGetSupportedListParams() *KubernetesGetSupportedListParams {
	return &KubernetesGetSupportedListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesGetSupportedListParamsWithTimeout creates a new KubernetesGetSupportedListParams object
// with the ability to set a timeout on a request.
func NewKubernetesGetSupportedListParamsWithTimeout(timeout time.Duration) *KubernetesGetSupportedListParams {
	return &KubernetesGetSupportedListParams{
		timeout: timeout,
	}
}

// NewKubernetesGetSupportedListParamsWithContext creates a new KubernetesGetSupportedListParams object
// with the ability to set a context for a request.
func NewKubernetesGetSupportedListParamsWithContext(ctx context.Context) *KubernetesGetSupportedListParams {
	return &KubernetesGetSupportedListParams{
		Context: ctx,
	}
}

// NewKubernetesGetSupportedListParamsWithHTTPClient creates a new KubernetesGetSupportedListParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesGetSupportedListParamsWithHTTPClient(client *http.Client) *KubernetesGetSupportedListParams {
	return &KubernetesGetSupportedListParams{
		HTTPClient: client,
	}
}

/* KubernetesGetSupportedListParams contains all the parameters to send to the API endpoint
   for the kubernetes get supported list operation.

   Typically these are written to a http.Request.
*/
type KubernetesGetSupportedListParams struct {

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes get supported list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetSupportedListParams) WithDefaults() *KubernetesGetSupportedListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes get supported list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetSupportedListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) WithTimeout(timeout time.Duration) *KubernetesGetSupportedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) WithContext(ctx context.Context) *KubernetesGetSupportedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) WithHTTPClient(client *http.Client) *KubernetesGetSupportedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV adds the v to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) WithV(v string) *KubernetesGetSupportedListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes get supported list params
func (o *KubernetesGetSupportedListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesGetSupportedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
