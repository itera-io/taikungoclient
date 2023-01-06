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
)

// NewSubscriptionSubscriptionForLandingPageParams creates a new SubscriptionSubscriptionForLandingPageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionSubscriptionForLandingPageParams() *SubscriptionSubscriptionForLandingPageParams {
	return &SubscriptionSubscriptionForLandingPageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionSubscriptionForLandingPageParamsWithTimeout creates a new SubscriptionSubscriptionForLandingPageParams object
// with the ability to set a timeout on a request.
func NewSubscriptionSubscriptionForLandingPageParamsWithTimeout(timeout time.Duration) *SubscriptionSubscriptionForLandingPageParams {
	return &SubscriptionSubscriptionForLandingPageParams{
		timeout: timeout,
	}
}

// NewSubscriptionSubscriptionForLandingPageParamsWithContext creates a new SubscriptionSubscriptionForLandingPageParams object
// with the ability to set a context for a request.
func NewSubscriptionSubscriptionForLandingPageParamsWithContext(ctx context.Context) *SubscriptionSubscriptionForLandingPageParams {
	return &SubscriptionSubscriptionForLandingPageParams{
		Context: ctx,
	}
}

// NewSubscriptionSubscriptionForLandingPageParamsWithHTTPClient creates a new SubscriptionSubscriptionForLandingPageParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionSubscriptionForLandingPageParamsWithHTTPClient(client *http.Client) *SubscriptionSubscriptionForLandingPageParams {
	return &SubscriptionSubscriptionForLandingPageParams{
		HTTPClient: client,
	}
}

/*
SubscriptionSubscriptionForLandingPageParams contains all the parameters to send to the API endpoint

	for the subscription subscription for landing page operation.

	Typically these are written to a http.Request.
*/
type SubscriptionSubscriptionForLandingPageParams struct {

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription subscription for landing page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionSubscriptionForLandingPageParams) WithDefaults() *SubscriptionSubscriptionForLandingPageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription subscription for landing page params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionSubscriptionForLandingPageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) WithTimeout(timeout time.Duration) *SubscriptionSubscriptionForLandingPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) WithContext(ctx context.Context) *SubscriptionSubscriptionForLandingPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) WithHTTPClient(client *http.Client) *SubscriptionSubscriptionForLandingPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV adds the v to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) WithV(v string) *SubscriptionSubscriptionForLandingPageParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the subscription subscription for landing page params
func (o *SubscriptionSubscriptionForLandingPageParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionSubscriptionForLandingPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}