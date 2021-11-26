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

// NewCronJobDeleteExpiredRequestsParams creates a new CronJobDeleteExpiredRequestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCronJobDeleteExpiredRequestsParams() *CronJobDeleteExpiredRequestsParams {
	return &CronJobDeleteExpiredRequestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCronJobDeleteExpiredRequestsParamsWithTimeout creates a new CronJobDeleteExpiredRequestsParams object
// with the ability to set a timeout on a request.
func NewCronJobDeleteExpiredRequestsParamsWithTimeout(timeout time.Duration) *CronJobDeleteExpiredRequestsParams {
	return &CronJobDeleteExpiredRequestsParams{
		timeout: timeout,
	}
}

// NewCronJobDeleteExpiredRequestsParamsWithContext creates a new CronJobDeleteExpiredRequestsParams object
// with the ability to set a context for a request.
func NewCronJobDeleteExpiredRequestsParamsWithContext(ctx context.Context) *CronJobDeleteExpiredRequestsParams {
	return &CronJobDeleteExpiredRequestsParams{
		Context: ctx,
	}
}

// NewCronJobDeleteExpiredRequestsParamsWithHTTPClient creates a new CronJobDeleteExpiredRequestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCronJobDeleteExpiredRequestsParamsWithHTTPClient(client *http.Client) *CronJobDeleteExpiredRequestsParams {
	return &CronJobDeleteExpiredRequestsParams{
		HTTPClient: client,
	}
}

/* CronJobDeleteExpiredRequestsParams contains all the parameters to send to the API endpoint
   for the cron job delete expired requests operation.

   Typically these are written to a http.Request.
*/
type CronJobDeleteExpiredRequestsParams struct {

	// Body.
	Body models.DeleteExpiredRequestCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cron job delete expired requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteExpiredRequestsParams) WithDefaults() *CronJobDeleteExpiredRequestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cron job delete expired requests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CronJobDeleteExpiredRequestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) WithTimeout(timeout time.Duration) *CronJobDeleteExpiredRequestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) WithContext(ctx context.Context) *CronJobDeleteExpiredRequestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) WithHTTPClient(client *http.Client) *CronJobDeleteExpiredRequestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) WithBody(body models.DeleteExpiredRequestCommand) *CronJobDeleteExpiredRequestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) SetBody(body models.DeleteExpiredRequestCommand) {
	o.Body = body
}

// WithV adds the v to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) WithV(v string) *CronJobDeleteExpiredRequestsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cron job delete expired requests params
func (o *CronJobDeleteExpiredRequestsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CronJobDeleteExpiredRequestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
