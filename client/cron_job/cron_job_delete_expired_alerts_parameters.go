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

// NewCronJobDeleteExpiredAlertsParams creates a new CronJobDeleteExpiredAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobDeleteExpiredAlertsParams() *CronJobDeleteExpiredAlertsParams {
	return &CronJobDeleteExpiredAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobDeleteExpiredAlertsParamsWithTimeout creates a new CronJobDeleteExpiredAlertsParams object
// with the ability to set a timeout on a request.
func NewCronJobDeleteExpiredAlertsParamsWithTimeout(timeout time.Duration) *CronJobDeleteExpiredAlertsParams {
	return &CronJobDeleteExpiredAlertsParams{
		timeout: timeout,
	}
}

// NewCronJobDeleteExpiredAlertsParamsWithContext creates a new CronJobDeleteExpiredAlertsParams object
// with the ability to set a context for a request.
func NewCronJobDeleteExpiredAlertsParamsWithContext(ctx context.Context) *CronJobDeleteExpiredAlertsParams {
	return &CronJobDeleteExpiredAlertsParams{
		Context: ctx,
	}
}

// NewCronJobDeleteExpiredAlertsParamsWithHTTPClient creates a new CronJobDeleteExpiredAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobDeleteExpiredAlertsParamsWithHTTPClient(client *http.Client) *CronJobDeleteExpiredAlertsParams {
	return &CronJobDeleteExpiredAlertsParams{
		HTTPClient: client,
	}
}

/*
CronJobDeleteExpiredAlertsParams contains all the parameters to send to the API endpoint

	for the cron job delete expired alerts operation.

	Typically these are written to a http.Request.
*/
type CronJobDeleteExpiredAlertsParams struct {

	// Body.
	Body models.DeleteExpiredAlertsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job delete expired alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteExpiredAlertsParams) WithDefaults() *CronJobDeleteExpiredAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job delete expired alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteExpiredAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) WithTimeout(timeout time.Duration) *CronJobDeleteExpiredAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) WithContext(ctx context.Context) *CronJobDeleteExpiredAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) WithHTTPClient(client *http.Client) *CronJobDeleteExpiredAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) WithBody(body models.DeleteExpiredAlertsCommand) *CronJobDeleteExpiredAlertsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) SetBody(body models.DeleteExpiredAlertsCommand) {
	o.Body = body
}

// WithV adds the v to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) WithV(v string) *CronJobDeleteExpiredAlertsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job delete expired alerts params
func (o *CronJobDeleteExpiredAlertsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobDeleteExpiredAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
