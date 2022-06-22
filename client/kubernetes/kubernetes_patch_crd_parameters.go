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

// NewKubernetesPatchCrdParams creates a new KubernetesPatchCrdParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesPatchCrdParams() *KubernetesPatchCrdParams {
	return &KubernetesPatchCrdParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesPatchCrdParamsWithTimeout creates a new KubernetesPatchCrdParams object
// with the ability to set a timeout on a request.
func NewKubernetesPatchCrdParamsWithTimeout(timeout time.Duration) *KubernetesPatchCrdParams {
	return &KubernetesPatchCrdParams{
		timeout: timeout,
	}
}

// NewKubernetesPatchCrdParamsWithContext creates a new KubernetesPatchCrdParams object
// with the ability to set a context for a request.
func NewKubernetesPatchCrdParamsWithContext(ctx context.Context) *KubernetesPatchCrdParams {
	return &KubernetesPatchCrdParams{
		Context: ctx,
	}
}

// NewKubernetesPatchCrdParamsWithHTTPClient creates a new KubernetesPatchCrdParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesPatchCrdParamsWithHTTPClient(client *http.Client) *KubernetesPatchCrdParams {
	return &KubernetesPatchCrdParams{
		HTTPClient: client,
	}
}

/* KubernetesPatchCrdParams contains all the parameters to send to the API endpoint
   for the kubernetes patch crd operation.

   Typically these are written to a http.Request.
*/
type KubernetesPatchCrdParams struct {

	// Body.
	Body *models.PatchCrdCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes patch crd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchCrdParams) WithDefaults() *KubernetesPatchCrdParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes patch crd params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesPatchCrdParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) WithTimeout(timeout time.Duration) *KubernetesPatchCrdParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) WithContext(ctx context.Context) *KubernetesPatchCrdParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) WithHTTPClient(client *http.Client) *KubernetesPatchCrdParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) WithBody(body *models.PatchCrdCommand) *KubernetesPatchCrdParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) SetBody(body *models.PatchCrdCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) WithV(v string) *KubernetesPatchCrdParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes patch crd params
func (o *KubernetesPatchCrdParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesPatchCrdParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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