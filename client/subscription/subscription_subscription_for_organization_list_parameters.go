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

// NewSubscriptionSubscriptionForOrganizationListParams creates a new SubscriptionSubscriptionForOrganizationListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionSubscriptionForOrganizationListParams() *SubscriptionSubscriptionForOrganizationListParams {
	return &SubscriptionSubscriptionForOrganizationListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionSubscriptionForOrganizationListParamsWithTimeout creates a new SubscriptionSubscriptionForOrganizationListParams object
// with the ability to set a timeout on a request.
func NewSubscriptionSubscriptionForOrganizationListParamsWithTimeout(timeout time.Duration) *SubscriptionSubscriptionForOrganizationListParams {
	return &SubscriptionSubscriptionForOrganizationListParams{
		timeout: timeout,
	}
}

// NewSubscriptionSubscriptionForOrganizationListParamsWithContext creates a new SubscriptionSubscriptionForOrganizationListParams object
// with the ability to set a context for a request.
func NewSubscriptionSubscriptionForOrganizationListParamsWithContext(ctx context.Context) *SubscriptionSubscriptionForOrganizationListParams {
	return &SubscriptionSubscriptionForOrganizationListParams{
		Context: ctx,
	}
}

// NewSubscriptionSubscriptionForOrganizationListParamsWithHTTPClient creates a new SubscriptionSubscriptionForOrganizationListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionSubscriptionForOrganizationListParamsWithHTTPClient(client *http.Client) *SubscriptionSubscriptionForOrganizationListParams {
	return &SubscriptionSubscriptionForOrganizationListParams{
		HTTPClient: client,
	}
}

/* SubscriptionSubscriptionForOrganizationListParams contains all the parameters to send to the API endpoint
   for the subscription subscription for organization list operation.

   Typically these are written to a http.Request.
*/
type SubscriptionSubscriptionForOrganizationListParams struct {

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription subscription for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionSubscriptionForOrganizationListParams) WithDefaults() *SubscriptionSubscriptionForOrganizationListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription subscription for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionSubscriptionForOrganizationListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) WithTimeout(timeout time.Duration) *SubscriptionSubscriptionForOrganizationListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) WithContext(ctx context.Context) *SubscriptionSubscriptionForOrganizationListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) WithHTTPClient(client *http.Client) *SubscriptionSubscriptionForOrganizationListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV adds the v to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) WithV(v string) *SubscriptionSubscriptionForOrganizationListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the subscription subscription for organization list params
func (o *SubscriptionSubscriptionForOrganizationListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionSubscriptionForOrganizationListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
