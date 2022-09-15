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

// NewSubscriptionCreateParams creates a new SubscriptionCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionCreateParams() *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionCreateParamsWithTimeout creates a new SubscriptionCreateParams object
// with the ability to set a timeout on a request.
func NewSubscriptionCreateParamsWithTimeout(timeout time.Duration) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		timeout: timeout,
	}
}

// NewSubscriptionCreateParamsWithContext creates a new SubscriptionCreateParams object
// with the ability to set a context for a request.
func NewSubscriptionCreateParamsWithContext(ctx context.Context) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		Context: ctx,
	}
}

// NewSubscriptionCreateParamsWithHTTPClient creates a new SubscriptionCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionCreateParamsWithHTTPClient(client *http.Client) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		HTTPClient: client,
	}
}

/*
SubscriptionCreateParams contains all the parameters to send to the API endpoint

	for the subscription create operation.

	Typically these are written to a http.Request.
*/
type SubscriptionCreateParams struct {

	// Body.
	Body *models.CreateSubscriptionCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionCreateParams) WithDefaults() *SubscriptionCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription create params
func (o *SubscriptionCreateParams) WithTimeout(timeout time.Duration) *SubscriptionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription create params
func (o *SubscriptionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription create params
func (o *SubscriptionCreateParams) WithContext(ctx context.Context) *SubscriptionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription create params
func (o *SubscriptionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription create params
func (o *SubscriptionCreateParams) WithHTTPClient(client *http.Client) *SubscriptionCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription create params
func (o *SubscriptionCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscription create params
func (o *SubscriptionCreateParams) WithBody(body *models.CreateSubscriptionCommand) *SubscriptionCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscription create params
func (o *SubscriptionCreateParams) SetBody(body *models.CreateSubscriptionCommand) {
	o.Body = body
}

// WithV adds the v to the subscription create params
func (o *SubscriptionCreateParams) WithV(v string) *SubscriptionCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the subscription create params
func (o *SubscriptionCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
