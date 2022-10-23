// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewNotificationsNotifyOwnerParams creates a new NotificationsNotifyOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationsNotifyOwnerParams() *NotificationsNotifyOwnerParams {
	return &NotificationsNotifyOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsNotifyOwnerParamsWithTimeout creates a new NotificationsNotifyOwnerParams object
// with the ability to set a timeout on a request.
func NewNotificationsNotifyOwnerParamsWithTimeout(timeout time.Duration) *NotificationsNotifyOwnerParams {
	return &NotificationsNotifyOwnerParams{
		timeout: timeout,
	}
}

// NewNotificationsNotifyOwnerParamsWithContext creates a new NotificationsNotifyOwnerParams object
// with the ability to set a context for a request.
func NewNotificationsNotifyOwnerParamsWithContext(ctx context.Context) *NotificationsNotifyOwnerParams {
	return &NotificationsNotifyOwnerParams{
		Context: ctx,
	}
}

// NewNotificationsNotifyOwnerParamsWithHTTPClient creates a new NotificationsNotifyOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationsNotifyOwnerParamsWithHTTPClient(client *http.Client) *NotificationsNotifyOwnerParams {
	return &NotificationsNotifyOwnerParams{
		HTTPClient: client,
	}
}

/*
NotificationsNotifyOwnerParams contains all the parameters to send to the API endpoint

	for the notifications notify owner operation.

	Typically these are written to a http.Request.
*/
type NotificationsNotifyOwnerParams struct {

	// Body.
	Body NotificationsNotifyOwnerBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notifications notify owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsNotifyOwnerParams) WithDefaults() *NotificationsNotifyOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifications notify owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsNotifyOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) WithTimeout(timeout time.Duration) *NotificationsNotifyOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) WithContext(ctx context.Context) *NotificationsNotifyOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) WithHTTPClient(client *http.Client) *NotificationsNotifyOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) WithBody(body NotificationsNotifyOwnerBody) *NotificationsNotifyOwnerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) SetBody(body NotificationsNotifyOwnerBody) {
	o.Body = body
}

// WithV adds the v to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) WithV(v string) *NotificationsNotifyOwnerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the notifications notify owner params
func (o *NotificationsNotifyOwnerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsNotifyOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
