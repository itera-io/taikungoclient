// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

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

// NewAlertingProfilesDeleteParams creates a new AlertingProfilesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAlertingProfilesDeleteParams() *AlertingProfilesDeleteParams {
	return &AlertingProfilesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAlertingProfilesDeleteParamsWithTimeout creates a new AlertingProfilesDeleteParams object
// with the ability to set a timeout on a request.
func NewAlertingProfilesDeleteParamsWithTimeout(timeout time.Duration) *AlertingProfilesDeleteParams {
	return &AlertingProfilesDeleteParams{
		timeout: timeout,
	}
}

// NewAlertingProfilesDeleteParamsWithContext creates a new AlertingProfilesDeleteParams object
// with the ability to set a context for a request.
func NewAlertingProfilesDeleteParamsWithContext(ctx context.Context) *AlertingProfilesDeleteParams {
	return &AlertingProfilesDeleteParams{
		Context: ctx,
	}
}

// NewAlertingProfilesDeleteParamsWithHTTPClient creates a new AlertingProfilesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAlertingProfilesDeleteParamsWithHTTPClient(client *http.Client) *AlertingProfilesDeleteParams {
	return &AlertingProfilesDeleteParams{
		HTTPClient: client,
	}
}

/* AlertingProfilesDeleteParams contains all the parameters to send to the API endpoint
   for the alerting profiles delete operation.

   Typically these are written to a http.Request.
*/
type AlertingProfilesDeleteParams struct {

	// Body.
	Body *models.DeleteAlertingProfilesCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the alerting profiles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesDeleteParams) WithDefaults() *AlertingProfilesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the alerting profiles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) WithTimeout(timeout time.Duration) *AlertingProfilesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) WithContext(ctx context.Context) *AlertingProfilesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) WithHTTPClient(client *http.Client) *AlertingProfilesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) WithBody(body *models.DeleteAlertingProfilesCommand) *AlertingProfilesDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) SetBody(body *models.DeleteAlertingProfilesCommand) {
	o.Body = body
}

// WithV adds the v to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) WithV(v string) *AlertingProfilesDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the alerting profiles delete params
func (o *AlertingProfilesDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AlertingProfilesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
