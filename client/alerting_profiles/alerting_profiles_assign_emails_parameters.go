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
	"github.com/go-openapi/swag"
)

// NewAlertingProfilesAssignEmailsParams creates a new AlertingProfilesAssignEmailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAlertingProfilesAssignEmailsParams() *AlertingProfilesAssignEmailsParams {
	return &AlertingProfilesAssignEmailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAlertingProfilesAssignEmailsParamsWithTimeout creates a new AlertingProfilesAssignEmailsParams object
// with the ability to set a timeout on a request.
func NewAlertingProfilesAssignEmailsParamsWithTimeout(timeout time.Duration) *AlertingProfilesAssignEmailsParams {
	return &AlertingProfilesAssignEmailsParams{
		timeout: timeout,
	}
}

// NewAlertingProfilesAssignEmailsParamsWithContext creates a new AlertingProfilesAssignEmailsParams object
// with the ability to set a context for a request.
func NewAlertingProfilesAssignEmailsParamsWithContext(ctx context.Context) *AlertingProfilesAssignEmailsParams {
	return &AlertingProfilesAssignEmailsParams{
		Context: ctx,
	}
}

// NewAlertingProfilesAssignEmailsParamsWithHTTPClient creates a new AlertingProfilesAssignEmailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAlertingProfilesAssignEmailsParamsWithHTTPClient(client *http.Client) *AlertingProfilesAssignEmailsParams {
	return &AlertingProfilesAssignEmailsParams{
		HTTPClient: client,
	}
}

/*
AlertingProfilesAssignEmailsParams contains all the parameters to send to the API endpoint

	for the alerting profiles assign emails operation.

	Typically these are written to a http.Request.
*/
type AlertingProfilesAssignEmailsParams struct {

	// Body.
	Body []*AlertingProfilesAssignEmailsParamsBodyItems0

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

// WithDefaults hydrates default values in the alerting profiles assign emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesAssignEmailsParams) WithDefaults() *AlertingProfilesAssignEmailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the alerting profiles assign emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesAssignEmailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithTimeout(timeout time.Duration) *AlertingProfilesAssignEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithContext(ctx context.Context) *AlertingProfilesAssignEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithHTTPClient(client *http.Client) *AlertingProfilesAssignEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithBody(body []*AlertingProfilesAssignEmailsParamsBodyItems0) *AlertingProfilesAssignEmailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetBody(body []*AlertingProfilesAssignEmailsParamsBodyItems0) {
	o.Body = body
}

// WithID adds the id to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithID(id int32) *AlertingProfilesAssignEmailsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) WithV(v string) *AlertingProfilesAssignEmailsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the alerting profiles assign emails params
func (o *AlertingProfilesAssignEmailsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AlertingProfilesAssignEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
