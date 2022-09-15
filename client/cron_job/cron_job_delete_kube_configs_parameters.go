// Code generated by go-swagger; DO NOT EDIT.

package cron_job

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

// NewCronJobDeleteKubeConfigsParams creates a new CronJobDeleteKubeConfigsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobDeleteKubeConfigsParams() *CronJobDeleteKubeConfigsParams {
	return &CronJobDeleteKubeConfigsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobDeleteKubeConfigsParamsWithTimeout creates a new CronJobDeleteKubeConfigsParams object
// with the ability to set a timeout on a request.
func NewCronJobDeleteKubeConfigsParamsWithTimeout(timeout time.Duration) *CronJobDeleteKubeConfigsParams {
	return &CronJobDeleteKubeConfigsParams{
		timeout: timeout,
	}
}

// NewCronJobDeleteKubeConfigsParamsWithContext creates a new CronJobDeleteKubeConfigsParams object
// with the ability to set a context for a request.
func NewCronJobDeleteKubeConfigsParamsWithContext(ctx context.Context) *CronJobDeleteKubeConfigsParams {
	return &CronJobDeleteKubeConfigsParams{
		Context: ctx,
	}
}

// NewCronJobDeleteKubeConfigsParamsWithHTTPClient creates a new CronJobDeleteKubeConfigsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobDeleteKubeConfigsParamsWithHTTPClient(client *http.Client) *CronJobDeleteKubeConfigsParams {
	return &CronJobDeleteKubeConfigsParams{
		HTTPClient: client,
	}
}

/*
CronJobDeleteKubeConfigsParams contains all the parameters to send to the API endpoint

	for the cron job delete kube configs operation.

	Typically these are written to a http.Request.
*/
type CronJobDeleteKubeConfigsParams struct {

	// Body.
	Body models.DeleteRemovedUsersKubeConfigCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job delete kube configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteKubeConfigsParams) WithDefaults() *CronJobDeleteKubeConfigsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job delete kube configs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteKubeConfigsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) WithTimeout(timeout time.Duration) *CronJobDeleteKubeConfigsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) WithContext(ctx context.Context) *CronJobDeleteKubeConfigsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) WithHTTPClient(client *http.Client) *CronJobDeleteKubeConfigsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) WithBody(body models.DeleteRemovedUsersKubeConfigCommand) *CronJobDeleteKubeConfigsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) SetBody(body models.DeleteRemovedUsersKubeConfigCommand) {
	o.Body = body
}

// WithV adds the v to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) WithV(v string) *CronJobDeleteKubeConfigsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job delete kube configs params
func (o *CronJobDeleteKubeConfigsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobDeleteKubeConfigsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
