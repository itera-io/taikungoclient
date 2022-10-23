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
)

// NewCronJobSyncProjectsParams creates a new CronJobSyncProjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobSyncProjectsParams() *CronJobSyncProjectsParams {
	return &CronJobSyncProjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobSyncProjectsParamsWithTimeout creates a new CronJobSyncProjectsParams object
// with the ability to set a timeout on a request.
func NewCronJobSyncProjectsParamsWithTimeout(timeout time.Duration) *CronJobSyncProjectsParams {
	return &CronJobSyncProjectsParams{
		timeout: timeout,
	}
}

// NewCronJobSyncProjectsParamsWithContext creates a new CronJobSyncProjectsParams object
// with the ability to set a context for a request.
func NewCronJobSyncProjectsParamsWithContext(ctx context.Context) *CronJobSyncProjectsParams {
	return &CronJobSyncProjectsParams{
		Context: ctx,
	}
}

// NewCronJobSyncProjectsParamsWithHTTPClient creates a new CronJobSyncProjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobSyncProjectsParamsWithHTTPClient(client *http.Client) *CronJobSyncProjectsParams {
	return &CronJobSyncProjectsParams{
		HTTPClient: client,
	}
}

/*
CronJobSyncProjectsParams contains all the parameters to send to the API endpoint

	for the cron job sync projects operation.

	Typically these are written to a http.Request.
*/
type CronJobSyncProjectsParams struct {

	// Body.
	Body interface{}

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job sync projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobSyncProjectsParams) WithDefaults() *CronJobSyncProjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job sync projects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobSyncProjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job sync projects params
func (o *CronJobSyncProjectsParams) WithTimeout(timeout time.Duration) *CronJobSyncProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job sync projects params
func (o *CronJobSyncProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job sync projects params
func (o *CronJobSyncProjectsParams) WithContext(ctx context.Context) *CronJobSyncProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job sync projects params
func (o *CronJobSyncProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job sync projects params
func (o *CronJobSyncProjectsParams) WithHTTPClient(client *http.Client) *CronJobSyncProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job sync projects params
func (o *CronJobSyncProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job sync projects params
func (o *CronJobSyncProjectsParams) WithBody(body interface{}) *CronJobSyncProjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job sync projects params
func (o *CronJobSyncProjectsParams) SetBody(body interface{}) {
	o.Body = body
}

// WithV adds the v to the cron job sync projects params
func (o *CronJobSyncProjectsParams) WithV(v string) *CronJobSyncProjectsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job sync projects params
func (o *CronJobSyncProjectsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobSyncProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
