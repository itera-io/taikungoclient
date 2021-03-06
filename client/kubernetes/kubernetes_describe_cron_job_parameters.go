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

// NewKubernetesDescribeCronJobParams creates a new KubernetesDescribeCronJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDescribeCronJobParams() *KubernetesDescribeCronJobParams {
	return &KubernetesDescribeCronJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDescribeCronJobParamsWithTimeout creates a new KubernetesDescribeCronJobParams object
// with the ability to set a timeout on a request.
func NewKubernetesDescribeCronJobParamsWithTimeout(timeout time.Duration) *KubernetesDescribeCronJobParams {
	return &KubernetesDescribeCronJobParams{
		timeout: timeout,
	}
}

// NewKubernetesDescribeCronJobParamsWithContext creates a new KubernetesDescribeCronJobParams object
// with the ability to set a context for a request.
func NewKubernetesDescribeCronJobParamsWithContext(ctx context.Context) *KubernetesDescribeCronJobParams {
	return &KubernetesDescribeCronJobParams{
		Context: ctx,
	}
}

// NewKubernetesDescribeCronJobParamsWithHTTPClient creates a new KubernetesDescribeCronJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDescribeCronJobParamsWithHTTPClient(client *http.Client) *KubernetesDescribeCronJobParams {
	return &KubernetesDescribeCronJobParams{
		HTTPClient: client,
	}
}

/* KubernetesDescribeCronJobParams contains all the parameters to send to the API endpoint
   for the kubernetes describe cron job operation.

   Typically these are written to a http.Request.
*/
type KubernetesDescribeCronJobParams struct {

	// Body.
	Body *models.DescribeCronJobCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes describe cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeCronJobParams) WithDefaults() *KubernetesDescribeCronJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes describe cron job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDescribeCronJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) WithTimeout(timeout time.Duration) *KubernetesDescribeCronJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) WithContext(ctx context.Context) *KubernetesDescribeCronJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) WithHTTPClient(client *http.Client) *KubernetesDescribeCronJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) WithBody(body *models.DescribeCronJobCommand) *KubernetesDescribeCronJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) SetBody(body *models.DescribeCronJobCommand) {
	o.Body = body
}

// WithV adds the v to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) WithV(v string) *KubernetesDescribeCronJobParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes describe cron job params
func (o *KubernetesDescribeCronJobParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDescribeCronJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
