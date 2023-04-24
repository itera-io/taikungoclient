// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewSubscriptionBindParams creates a new SubscriptionBindParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionBindParams() *SubscriptionBindParams {
	return &SubscriptionBindParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionBindParamsWithTimeout creates a new SubscriptionBindParams object
// with the ability to set a timeout on a request.
func NewSubscriptionBindParamsWithTimeout(timeout time.Duration) *SubscriptionBindParams {
	return &SubscriptionBindParams{
		timeout: timeout,
	}
}

// NewSubscriptionBindParamsWithContext creates a new SubscriptionBindParams object
// with the ability to set a context for a request.
func NewSubscriptionBindParamsWithContext(ctx context.Context) *SubscriptionBindParams {
	return &SubscriptionBindParams{
		Context: ctx,
	}
}

// NewSubscriptionBindParamsWithHTTPClient creates a new SubscriptionBindParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionBindParamsWithHTTPClient(client *http.Client) *SubscriptionBindParams {
	return &SubscriptionBindParams{
		HTTPClient: client,
	}
}

/*
SubscriptionBindParams contains all the parameters to send to the API endpoint

	for the subscription bind operation.

	Typically these are written to a http.Request.
*/
type SubscriptionBindParams struct {

	// Body.
	Body *models.BindSubscriptionCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBindParams) WithDefaults() *SubscriptionBindParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription bind params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionBindParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription bind params
func (o *SubscriptionBindParams) WithTimeout(timeout time.Duration) *SubscriptionBindParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription bind params
func (o *SubscriptionBindParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription bind params
func (o *SubscriptionBindParams) WithContext(ctx context.Context) *SubscriptionBindParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription bind params
func (o *SubscriptionBindParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription bind params
func (o *SubscriptionBindParams) WithHTTPClient(client *http.Client) *SubscriptionBindParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription bind params
func (o *SubscriptionBindParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription bind params
func (o *SubscriptionBindParams) WithBody(body *models.BindSubscriptionCommand) *SubscriptionBindParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription bind params
func (o *SubscriptionBindParams) SetBody(body *models.BindSubscriptionCommand) {
	o.Body = body
}

// WithV adds the v to the subscription bind params
func (o *SubscriptionBindParams) WithV(v string) *SubscriptionBindParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the subscription bind params
func (o *SubscriptionBindParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionBindParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
