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

	"github.com/itera-io/taikungoclient/models"
)

// NewNotificationsUpdateParams creates a new NotificationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationsUpdateParams() *NotificationsUpdateParams {
	return &NotificationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsUpdateParamsWithTimeout creates a new NotificationsUpdateParams object
// with the ability to set a timeout on a request.
func NewNotificationsUpdateParamsWithTimeout(timeout time.Duration) *NotificationsUpdateParams {
	return &NotificationsUpdateParams{
		timeout: timeout,
	}
}

// NewNotificationsUpdateParamsWithContext creates a new NotificationsUpdateParams object
// with the ability to set a context for a request.
func NewNotificationsUpdateParamsWithContext(ctx context.Context) *NotificationsUpdateParams {
	return &NotificationsUpdateParams{
		Context: ctx,
	}
}

// NewNotificationsUpdateParamsWithHTTPClient creates a new NotificationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationsUpdateParamsWithHTTPClient(client *http.Client) *NotificationsUpdateParams {
	return &NotificationsUpdateParams{
		HTTPClient: client,
	}
}

/*
NotificationsUpdateParams contains all the parameters to send to the API endpoint

	for the notifications update operation.

	Typically these are written to a http.Request.
*/
type NotificationsUpdateParams struct {

	// Body.
	Body *models.UpdateNotificationsCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsUpdateParams) WithDefaults() *NotificationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifications update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifications update params
func (o *NotificationsUpdateParams) WithTimeout(timeout time.Duration) *NotificationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications update params
func (o *NotificationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications update params
func (o *NotificationsUpdateParams) WithContext(ctx context.Context) *NotificationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications update params
func (o *NotificationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications update params
func (o *NotificationsUpdateParams) WithHTTPClient(client *http.Client) *NotificationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications update params
func (o *NotificationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the notifications update params
func (o *NotificationsUpdateParams) WithBody(body *models.UpdateNotificationsCommand) *NotificationsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the notifications update params
func (o *NotificationsUpdateParams) SetBody(body *models.UpdateNotificationsCommand) {
	o.Body = body
}

// WithV adds the v to the notifications update params
func (o *NotificationsUpdateParams) WithV(v string) *NotificationsUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the notifications update params
func (o *NotificationsUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
