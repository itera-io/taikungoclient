// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersToggleNotificationModeParams creates a new UsersToggleNotificationModeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersToggleNotificationModeParams() *UsersToggleNotificationModeParams {
	return &UsersToggleNotificationModeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersToggleNotificationModeParamsWithTimeout creates a new UsersToggleNotificationModeParams object
// with the ability to set a timeout on a request.
func NewUsersToggleNotificationModeParamsWithTimeout(timeout time.Duration) *UsersToggleNotificationModeParams {
	return &UsersToggleNotificationModeParams{
		timeout: timeout,
	}
}

// NewUsersToggleNotificationModeParamsWithContext creates a new UsersToggleNotificationModeParams object
// with the ability to set a context for a request.
func NewUsersToggleNotificationModeParamsWithContext(ctx context.Context) *UsersToggleNotificationModeParams {
	return &UsersToggleNotificationModeParams{
		Context: ctx,
	}
}

// NewUsersToggleNotificationModeParamsWithHTTPClient creates a new UsersToggleNotificationModeParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersToggleNotificationModeParamsWithHTTPClient(client *http.Client) *UsersToggleNotificationModeParams {
	return &UsersToggleNotificationModeParams{
		HTTPClient: client,
	}
}

/* UsersToggleNotificationModeParams contains all the parameters to send to the API endpoint
   for the users toggle notification mode operation.

   Typically these are written to a http.Request.
*/
type UsersToggleNotificationModeParams struct {

	// Mode.
	Mode *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users toggle notification mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersToggleNotificationModeParams) WithDefaults() *UsersToggleNotificationModeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users toggle notification mode params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersToggleNotificationModeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) WithTimeout(timeout time.Duration) *UsersToggleNotificationModeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) WithContext(ctx context.Context) *UsersToggleNotificationModeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) WithHTTPClient(client *http.Client) *UsersToggleNotificationModeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMode adds the mode to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) WithMode(mode *string) *UsersToggleNotificationModeParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithV adds the v to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) WithV(v string) *UsersToggleNotificationModeParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the users toggle notification mode params
func (o *UsersToggleNotificationModeParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UsersToggleNotificationModeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
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