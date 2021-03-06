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

	"github.com/itera-io/taikungoclient/models"
)

// NewKubernetesProfilesCreateParams creates a new KubernetesProfilesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesProfilesCreateParams() *KubernetesProfilesCreateParams {
	return &KubernetesProfilesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesProfilesCreateParamsWithTimeout creates a new KubernetesProfilesCreateParams object
// with the ability to set a timeout on a request.
func NewKubernetesProfilesCreateParamsWithTimeout(timeout time.Duration) *KubernetesProfilesCreateParams {
	return &KubernetesProfilesCreateParams{
		timeout: timeout,
	}
}

// NewKubernetesProfilesCreateParamsWithContext creates a new KubernetesProfilesCreateParams object
// with the ability to set a context for a request.
func NewKubernetesProfilesCreateParamsWithContext(ctx context.Context) *KubernetesProfilesCreateParams {
	return &KubernetesProfilesCreateParams{
		Context: ctx,
	}
}

// NewKubernetesProfilesCreateParamsWithHTTPClient creates a new KubernetesProfilesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesProfilesCreateParamsWithHTTPClient(client *http.Client) *KubernetesProfilesCreateParams {
	return &KubernetesProfilesCreateParams{
		HTTPClient: client,
	}
}

/* KubernetesProfilesCreateParams contains all the parameters to send to the API endpoint
   for the kubernetes profiles create operation.

   Typically these are written to a http.Request.
*/
type KubernetesProfilesCreateParams struct {

	// Body.
	Body *models.CreateKubernetesProfileCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes profiles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesCreateParams) WithDefaults() *KubernetesProfilesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes profiles create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) WithTimeout(timeout time.Duration) *KubernetesProfilesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) WithContext(ctx context.Context) *KubernetesProfilesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) WithHTTPClient(client *http.Client) *KubernetesProfilesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) WithBody(body *models.CreateKubernetesProfileCommand) *KubernetesProfilesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) SetBody(body *models.CreateKubernetesProfileCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) WithV(v string) *KubernetesProfilesCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes profiles create params
func (o *KubernetesProfilesCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesProfilesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
