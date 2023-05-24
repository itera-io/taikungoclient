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

// NewCronJobTriggerTemplatesParams creates a new CronJobTriggerTemplatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobTriggerTemplatesParams() *CronJobTriggerTemplatesParams {
	return &CronJobTriggerTemplatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobTriggerTemplatesParamsWithTimeout creates a new CronJobTriggerTemplatesParams object
// with the ability to set a timeout on a request.
func NewCronJobTriggerTemplatesParamsWithTimeout(timeout time.Duration) *CronJobTriggerTemplatesParams {
	return &CronJobTriggerTemplatesParams{
		timeout: timeout,
	}
}

// NewCronJobTriggerTemplatesParamsWithContext creates a new CronJobTriggerTemplatesParams object
// with the ability to set a context for a request.
func NewCronJobTriggerTemplatesParamsWithContext(ctx context.Context) *CronJobTriggerTemplatesParams {
	return &CronJobTriggerTemplatesParams{
		Context: ctx,
	}
}

// NewCronJobTriggerTemplatesParamsWithHTTPClient creates a new CronJobTriggerTemplatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobTriggerTemplatesParamsWithHTTPClient(client *http.Client) *CronJobTriggerTemplatesParams {
	return &CronJobTriggerTemplatesParams{
		HTTPClient: client,
	}
}

/*
CronJobTriggerTemplatesParams contains all the parameters to send to the API endpoint

	for the cron job trigger templates operation.

	Typically these are written to a http.Request.
*/
type CronJobTriggerTemplatesParams struct {

	// Body.
	Body models.TriggerTemplateCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job trigger templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobTriggerTemplatesParams) WithDefaults() *CronJobTriggerTemplatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job trigger templates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobTriggerTemplatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) WithTimeout(timeout time.Duration) *CronJobTriggerTemplatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) WithContext(ctx context.Context) *CronJobTriggerTemplatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) WithHTTPClient(client *http.Client) *CronJobTriggerTemplatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) WithBody(body models.TriggerTemplateCommand) *CronJobTriggerTemplatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) SetBody(body models.TriggerTemplateCommand) {
	o.Body = body
}

// WithV adds the v to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) WithV(v string) *CronJobTriggerTemplatesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job trigger templates params
func (o *CronJobTriggerTemplatesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobTriggerTemplatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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