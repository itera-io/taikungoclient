// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

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

// NewAccessProfilesDeleteParams creates a new AccessProfilesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessProfilesDeleteParams() *AccessProfilesDeleteParams {
	return &AccessProfilesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessProfilesDeleteParamsWithTimeout creates a new AccessProfilesDeleteParams object
// with the ability to set a timeout on a request.
func NewAccessProfilesDeleteParamsWithTimeout(timeout time.Duration) *AccessProfilesDeleteParams {
	return &AccessProfilesDeleteParams{
		timeout: timeout,
	}
}

// NewAccessProfilesDeleteParamsWithContext creates a new AccessProfilesDeleteParams object
// with the ability to set a context for a request.
func NewAccessProfilesDeleteParamsWithContext(ctx context.Context) *AccessProfilesDeleteParams {
	return &AccessProfilesDeleteParams{
		Context: ctx,
	}
}

// NewAccessProfilesDeleteParamsWithHTTPClient creates a new AccessProfilesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessProfilesDeleteParamsWithHTTPClient(client *http.Client) *AccessProfilesDeleteParams {
	return &AccessProfilesDeleteParams{
		HTTPClient: client,
	}
}

/* AccessProfilesDeleteParams contains all the parameters to send to the API endpoint
   for the access profiles delete operation.

   Typically these are written to a http.Request.
*/
type AccessProfilesDeleteParams struct {

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

// WithDefaults hydrates default values in the access profiles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesDeleteParams) WithDefaults() *AccessProfilesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access profiles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access profiles delete params
func (o *AccessProfilesDeleteParams) WithTimeout(timeout time.Duration) *AccessProfilesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access profiles delete params
func (o *AccessProfilesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access profiles delete params
func (o *AccessProfilesDeleteParams) WithContext(ctx context.Context) *AccessProfilesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access profiles delete params
func (o *AccessProfilesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access profiles delete params
func (o *AccessProfilesDeleteParams) WithHTTPClient(client *http.Client) *AccessProfilesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access profiles delete params
func (o *AccessProfilesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the access profiles delete params
func (o *AccessProfilesDeleteParams) WithID(id int32) *AccessProfilesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the access profiles delete params
func (o *AccessProfilesDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the access profiles delete params
func (o *AccessProfilesDeleteParams) WithV(v string) *AccessProfilesDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the access profiles delete params
func (o *AccessProfilesDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AccessProfilesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
