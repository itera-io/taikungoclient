// Code generated by go-swagger; DO NOT EDIT.

package showback_credentials

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

// NewShowbackCredentialsDropdownParams creates a new ShowbackCredentialsDropdownParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackCredentialsDropdownParams() *ShowbackCredentialsDropdownParams {
	return &ShowbackCredentialsDropdownParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackCredentialsDropdownParamsWithTimeout creates a new ShowbackCredentialsDropdownParams object
// with the ability to set a timeout on a request.
func NewShowbackCredentialsDropdownParamsWithTimeout(timeout time.Duration) *ShowbackCredentialsDropdownParams {
	return &ShowbackCredentialsDropdownParams{
		timeout: timeout,
	}
}

// NewShowbackCredentialsDropdownParamsWithContext creates a new ShowbackCredentialsDropdownParams object
// with the ability to set a context for a request.
func NewShowbackCredentialsDropdownParamsWithContext(ctx context.Context) *ShowbackCredentialsDropdownParams {
	return &ShowbackCredentialsDropdownParams{
		Context: ctx,
	}
}

// NewShowbackCredentialsDropdownParamsWithHTTPClient creates a new ShowbackCredentialsDropdownParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackCredentialsDropdownParamsWithHTTPClient(client *http.Client) *ShowbackCredentialsDropdownParams {
	return &ShowbackCredentialsDropdownParams{
		HTTPClient: client,
	}
}

/*
ShowbackCredentialsDropdownParams contains all the parameters to send to the API endpoint

	for the showback credentials dropdown operation.

	Typically these are written to a http.Request.
*/
type ShowbackCredentialsDropdownParams struct {

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback credentials dropdown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsDropdownParams) WithDefaults() *ShowbackCredentialsDropdownParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback credentials dropdown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackCredentialsDropdownParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) WithTimeout(timeout time.Duration) *ShowbackCredentialsDropdownParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) WithContext(ctx context.Context) *ShowbackCredentialsDropdownParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) WithHTTPClient(client *http.Client) *ShowbackCredentialsDropdownParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) WithOrganizationID(organizationID *int32) *ShowbackCredentialsDropdownParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) WithV(v string) *ShowbackCredentialsDropdownParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback credentials dropdown params
func (o *ShowbackCredentialsDropdownParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackCredentialsDropdownParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OrganizationID != nil {

		// query param organizationId
		var qrOrganizationID int32

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := swag.FormatInt32(qrOrganizationID)
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationId", qOrganizationID); err != nil {
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
