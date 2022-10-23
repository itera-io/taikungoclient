// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_profiles

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

// NewKubernetesProfilesLockManagerParams creates a new KubernetesProfilesLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesProfilesLockManagerParams() *KubernetesProfilesLockManagerParams {
	return &KubernetesProfilesLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesProfilesLockManagerParamsWithTimeout creates a new KubernetesProfilesLockManagerParams object
// with the ability to set a timeout on a request.
func NewKubernetesProfilesLockManagerParamsWithTimeout(timeout time.Duration) *KubernetesProfilesLockManagerParams {
	return &KubernetesProfilesLockManagerParams{
		timeout: timeout,
	}
}

// NewKubernetesProfilesLockManagerParamsWithContext creates a new KubernetesProfilesLockManagerParams object
// with the ability to set a context for a request.
func NewKubernetesProfilesLockManagerParamsWithContext(ctx context.Context) *KubernetesProfilesLockManagerParams {
	return &KubernetesProfilesLockManagerParams{
		Context: ctx,
	}
}

// NewKubernetesProfilesLockManagerParamsWithHTTPClient creates a new KubernetesProfilesLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesProfilesLockManagerParamsWithHTTPClient(client *http.Client) *KubernetesProfilesLockManagerParams {
	return &KubernetesProfilesLockManagerParams{
		HTTPClient: client,
	}
}

/*
KubernetesProfilesLockManagerParams contains all the parameters to send to the API endpoint

	for the kubernetes profiles lock manager operation.

	Typically these are written to a http.Request.
*/
type KubernetesProfilesLockManagerParams struct {

	// Body.
	Body KubernetesProfilesLockManagerBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes profiles lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesLockManagerParams) WithDefaults() *KubernetesProfilesLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes profiles lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) WithTimeout(timeout time.Duration) *KubernetesProfilesLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) WithContext(ctx context.Context) *KubernetesProfilesLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) WithHTTPClient(client *http.Client) *KubernetesProfilesLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) WithBody(body KubernetesProfilesLockManagerBody) *KubernetesProfilesLockManagerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) SetBody(body KubernetesProfilesLockManagerBody) {
	o.Body = body
}

// WithV adds the v to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) WithV(v string) *KubernetesProfilesLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes profiles lock manager params
func (o *KubernetesProfilesLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesProfilesLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
