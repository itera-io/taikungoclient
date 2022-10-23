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

	"github.com/itera-io/taikungoclient/models"
)

// NewAccessProfilesUpdateParams creates a new AccessProfilesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessProfilesUpdateParams() *AccessProfilesUpdateParams {
	return &AccessProfilesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessProfilesUpdateParamsWithTimeout creates a new AccessProfilesUpdateParams object
// with the ability to set a timeout on a request.
func NewAccessProfilesUpdateParamsWithTimeout(timeout time.Duration) *AccessProfilesUpdateParams {
	return &AccessProfilesUpdateParams{
		timeout: timeout,
	}
}

// NewAccessProfilesUpdateParamsWithContext creates a new AccessProfilesUpdateParams object
// with the ability to set a context for a request.
func NewAccessProfilesUpdateParamsWithContext(ctx context.Context) *AccessProfilesUpdateParams {
	return &AccessProfilesUpdateParams{
		Context: ctx,
	}
}

// NewAccessProfilesUpdateParamsWithHTTPClient creates a new AccessProfilesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessProfilesUpdateParamsWithHTTPClient(client *http.Client) *AccessProfilesUpdateParams {
	return &AccessProfilesUpdateParams{
		HTTPClient: client,
	}
}

/*
AccessProfilesUpdateParams contains all the parameters to send to the API endpoint

	for the access profiles update operation.

	Typically these are written to a http.Request.
*/
type AccessProfilesUpdateParams struct {

	// Body.
	Body *models.UpdateAccessProfileDto

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

// WithDefaults hydrates default values in the access profiles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesUpdateParams) WithDefaults() *AccessProfilesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access profiles update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access profiles update params
func (o *AccessProfilesUpdateParams) WithTimeout(timeout time.Duration) *AccessProfilesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access profiles update params
func (o *AccessProfilesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access profiles update params
func (o *AccessProfilesUpdateParams) WithContext(ctx context.Context) *AccessProfilesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access profiles update params
func (o *AccessProfilesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access profiles update params
func (o *AccessProfilesUpdateParams) WithHTTPClient(client *http.Client) *AccessProfilesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access profiles update params
func (o *AccessProfilesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the access profiles update params
func (o *AccessProfilesUpdateParams) WithBody(body *models.UpdateAccessProfileDto) *AccessProfilesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the access profiles update params
func (o *AccessProfilesUpdateParams) SetBody(body *models.UpdateAccessProfileDto) {
	o.Body = body
}

// WithID adds the id to the access profiles update params
func (o *AccessProfilesUpdateParams) WithID(id int32) *AccessProfilesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the access profiles update params
func (o *AccessProfilesUpdateParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the access profiles update params
func (o *AccessProfilesUpdateParams) WithV(v string) *AccessProfilesUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the access profiles update params
func (o *AccessProfilesUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AccessProfilesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
