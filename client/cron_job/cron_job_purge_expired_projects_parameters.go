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

// NewCronJobPurgeExpiredProjectsParams creates a new CronJobPurgeExpiredProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobPurgeExpiredProjectsParams() *CronJobPurgeExpiredProjectsParams {
	return &CronJobPurgeExpiredProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobPurgeExpiredProjectsParamsWithTimeout creates a new CronJobPurgeExpiredProjectsParams object
// with the ability to set a timeout on a request.
func NewCronJobPurgeExpiredProjectsParamsWithTimeout(timeout time.Duration) *CronJobPurgeExpiredProjectsParams {
	return &CronJobPurgeExpiredProjectsParams{
		timeout: timeout,
	}
}

// NewCronJobPurgeExpiredProjectsParamsWithContext creates a new CronJobPurgeExpiredProjectsParams object
// with the ability to set a context for a request.
func NewCronJobPurgeExpiredProjectsParamsWithContext(ctx context.Context) *CronJobPurgeExpiredProjectsParams {
	return &CronJobPurgeExpiredProjectsParams{
		Context: ctx,
	}
}

// NewCronJobPurgeExpiredProjectsParamsWithHTTPClient creates a new CronJobPurgeExpiredProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobPurgeExpiredProjectsParamsWithHTTPClient(client *http.Client) *CronJobPurgeExpiredProjectsParams {
	return &CronJobPurgeExpiredProjectsParams{
		HTTPClient: client,
	}
}

/*
CronJobPurgeExpiredProjectsParams contains all the parameters to send to the API endpoint

	for the cron job purge expired projects operation.

	Typically these are written to a http.Request.
*/
type CronJobPurgeExpiredProjectsParams struct {

	// Body.
	Body models.PurgeExpiredProjectsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job purge expired projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobPurgeExpiredProjectsParams) WithDefaults() *CronJobPurgeExpiredProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job purge expired projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobPurgeExpiredProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) WithTimeout(timeout time.Duration) *CronJobPurgeExpiredProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) WithContext(ctx context.Context) *CronJobPurgeExpiredProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) WithHTTPClient(client *http.Client) *CronJobPurgeExpiredProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) WithBody(body models.PurgeExpiredProjectsCommand) *CronJobPurgeExpiredProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) SetBody(body models.PurgeExpiredProjectsCommand) {
	o.Body = body
}

// WithV adds the v to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) WithV(v string) *CronJobPurgeExpiredProjectsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job purge expired projects params
func (o *CronJobPurgeExpiredProjectsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobPurgeExpiredProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
