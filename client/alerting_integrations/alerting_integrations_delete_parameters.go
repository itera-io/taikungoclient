// Code generated by go-swagger; DO NOT EDIT.

package alerting_integrations

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
	"github.com/go-openapi/swag"
)

// NewAlertingIntegrationsDeleteParams creates a new AlertingIntegrationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAlertingIntegrationsDeleteParams() *AlertingIntegrationsDeleteParams {
	return &AlertingIntegrationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAlertingIntegrationsDeleteParamsWithTimeout creates a new AlertingIntegrationsDeleteParams object
// with the ability to set a timeout on a request.
func NewAlertingIntegrationsDeleteParamsWithTimeout(timeout time.Duration) *AlertingIntegrationsDeleteParams {
	return &AlertingIntegrationsDeleteParams{
		timeout: timeout,
	}
}

// NewAlertingIntegrationsDeleteParamsWithContext creates a new AlertingIntegrationsDeleteParams object
// with the ability to set a context for a request.
func NewAlertingIntegrationsDeleteParamsWithContext(ctx context.Context) *AlertingIntegrationsDeleteParams {
	return &AlertingIntegrationsDeleteParams{
		Context: ctx,
	}
}

// NewAlertingIntegrationsDeleteParamsWithHTTPClient creates a new AlertingIntegrationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAlertingIntegrationsDeleteParamsWithHTTPClient(client *http.Client) *AlertingIntegrationsDeleteParams {
	return &AlertingIntegrationsDeleteParams{
		HTTPClient: client,
	}
}

/*
AlertingIntegrationsDeleteParams contains all the parameters to send to the API endpoint

	for the alerting integrations delete operation.

	Typically these are written to a http.Request.
*/
type AlertingIntegrationsDeleteParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the alerting integrations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingIntegrationsDeleteParams) WithDefaults() *AlertingIntegrationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the alerting integrations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingIntegrationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) WithTimeout(timeout time.Duration) *AlertingIntegrationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) WithContext(ctx context.Context) *AlertingIntegrationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) WithHTTPClient(client *http.Client) *AlertingIntegrationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) WithID(id int32) *AlertingIntegrationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) WithV(v string) *AlertingIntegrationsDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the alerting integrations delete params
func (o *AlertingIntegrationsDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AlertingIntegrationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
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
