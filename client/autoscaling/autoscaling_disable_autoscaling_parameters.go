// Code generated by go-swagger; DO NOT EDIT.

package autoscaling

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

// NewAutoscalingDisableAutoscalingParams creates a new AutoscalingDisableAutoscalingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutoscalingDisableAutoscalingParams() *AutoscalingDisableAutoscalingParams {
	return &AutoscalingDisableAutoscalingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutoscalingDisableAutoscalingParamsWithTimeout creates a new AutoscalingDisableAutoscalingParams object
// with the ability to set a timeout on a request.
func NewAutoscalingDisableAutoscalingParamsWithTimeout(timeout time.Duration) *AutoscalingDisableAutoscalingParams {
	return &AutoscalingDisableAutoscalingParams{
		timeout: timeout,
	}
}

// NewAutoscalingDisableAutoscalingParamsWithContext creates a new AutoscalingDisableAutoscalingParams object
// with the ability to set a context for a request.
func NewAutoscalingDisableAutoscalingParamsWithContext(ctx context.Context) *AutoscalingDisableAutoscalingParams {
	return &AutoscalingDisableAutoscalingParams{
		Context: ctx,
	}
}

// NewAutoscalingDisableAutoscalingParamsWithHTTPClient creates a new AutoscalingDisableAutoscalingParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutoscalingDisableAutoscalingParamsWithHTTPClient(client *http.Client) *AutoscalingDisableAutoscalingParams {
	return &AutoscalingDisableAutoscalingParams{
		HTTPClient: client,
	}
}

/*
AutoscalingDisableAutoscalingParams contains all the parameters to send to the API endpoint

	for the autoscaling disable autoscaling operation.

	Typically these are written to a http.Request.
*/
type AutoscalingDisableAutoscalingParams struct {

	// Body.
	Body AutoscalingDisableAutoscalingBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the autoscaling disable autoscaling params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoscalingDisableAutoscalingParams) WithDefaults() *AutoscalingDisableAutoscalingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the autoscaling disable autoscaling params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoscalingDisableAutoscalingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) WithTimeout(timeout time.Duration) *AutoscalingDisableAutoscalingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) WithContext(ctx context.Context) *AutoscalingDisableAutoscalingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) WithHTTPClient(client *http.Client) *AutoscalingDisableAutoscalingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) WithBody(body AutoscalingDisableAutoscalingBody) *AutoscalingDisableAutoscalingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) SetBody(body AutoscalingDisableAutoscalingBody) {
	o.Body = body
}

// WithV adds the v to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) WithV(v string) *AutoscalingDisableAutoscalingParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the autoscaling disable autoscaling params
func (o *AutoscalingDisableAutoscalingParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AutoscalingDisableAutoscalingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
