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

// NewCronJobDeleteAwsSpotInstancesParams creates a new CronJobDeleteAwsSpotInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobDeleteAwsSpotInstancesParams() *CronJobDeleteAwsSpotInstancesParams {
	return &CronJobDeleteAwsSpotInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobDeleteAwsSpotInstancesParamsWithTimeout creates a new CronJobDeleteAwsSpotInstancesParams object
// with the ability to set a timeout on a request.
func NewCronJobDeleteAwsSpotInstancesParamsWithTimeout(timeout time.Duration) *CronJobDeleteAwsSpotInstancesParams {
	return &CronJobDeleteAwsSpotInstancesParams{
		timeout: timeout,
	}
}

// NewCronJobDeleteAwsSpotInstancesParamsWithContext creates a new CronJobDeleteAwsSpotInstancesParams object
// with the ability to set a context for a request.
func NewCronJobDeleteAwsSpotInstancesParamsWithContext(ctx context.Context) *CronJobDeleteAwsSpotInstancesParams {
	return &CronJobDeleteAwsSpotInstancesParams{
		Context: ctx,
	}
}

// NewCronJobDeleteAwsSpotInstancesParamsWithHTTPClient creates a new CronJobDeleteAwsSpotInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobDeleteAwsSpotInstancesParamsWithHTTPClient(client *http.Client) *CronJobDeleteAwsSpotInstancesParams {
	return &CronJobDeleteAwsSpotInstancesParams{
		HTTPClient: client,
	}
}

/*
CronJobDeleteAwsSpotInstancesParams contains all the parameters to send to the API endpoint

	for the cron job delete aws spot instances operation.

	Typically these are written to a http.Request.
*/
type CronJobDeleteAwsSpotInstancesParams struct {

	// Body.
	Body models.DeleteRemovedSpotInstancesCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job delete aws spot instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteAwsSpotInstancesParams) WithDefaults() *CronJobDeleteAwsSpotInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job delete aws spot instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteAwsSpotInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) WithTimeout(timeout time.Duration) *CronJobDeleteAwsSpotInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) WithContext(ctx context.Context) *CronJobDeleteAwsSpotInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) WithHTTPClient(client *http.Client) *CronJobDeleteAwsSpotInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) WithBody(body models.DeleteRemovedSpotInstancesCommand) *CronJobDeleteAwsSpotInstancesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) SetBody(body models.DeleteRemovedSpotInstancesCommand) {
	o.Body = body
}

// WithV adds the v to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) WithV(v string) *CronJobDeleteAwsSpotInstancesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job delete aws spot instances params
func (o *CronJobDeleteAwsSpotInstancesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobDeleteAwsSpotInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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