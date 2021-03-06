// Code generated by go-swagger; DO NOT EDIT.

package showback

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

// NewShowbackShowbackCredentialsDropdownParams creates a new ShowbackShowbackCredentialsDropdownParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackShowbackCredentialsDropdownParams() *ShowbackShowbackCredentialsDropdownParams {
	return &ShowbackShowbackCredentialsDropdownParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackShowbackCredentialsDropdownParamsWithTimeout creates a new ShowbackShowbackCredentialsDropdownParams object
// with the ability to set a timeout on a request.
func NewShowbackShowbackCredentialsDropdownParamsWithTimeout(timeout time.Duration) *ShowbackShowbackCredentialsDropdownParams {
	return &ShowbackShowbackCredentialsDropdownParams{
		timeout: timeout,
	}
}

// NewShowbackShowbackCredentialsDropdownParamsWithContext creates a new ShowbackShowbackCredentialsDropdownParams object
// with the ability to set a context for a request.
func NewShowbackShowbackCredentialsDropdownParamsWithContext(ctx context.Context) *ShowbackShowbackCredentialsDropdownParams {
	return &ShowbackShowbackCredentialsDropdownParams{
		Context: ctx,
	}
}

// NewShowbackShowbackCredentialsDropdownParamsWithHTTPClient creates a new ShowbackShowbackCredentialsDropdownParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackShowbackCredentialsDropdownParamsWithHTTPClient(client *http.Client) *ShowbackShowbackCredentialsDropdownParams {
	return &ShowbackShowbackCredentialsDropdownParams{
		HTTPClient: client,
	}
}

/* ShowbackShowbackCredentialsDropdownParams contains all the parameters to send to the API endpoint
   for the showback showback credentials dropdown operation.

   Typically these are written to a http.Request.
*/
type ShowbackShowbackCredentialsDropdownParams struct {

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

// WithDefaults hydrates default values in the showback showback credentials dropdown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackShowbackCredentialsDropdownParams) WithDefaults() *ShowbackShowbackCredentialsDropdownParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback showback credentials dropdown params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackShowbackCredentialsDropdownParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) WithTimeout(timeout time.Duration) *ShowbackShowbackCredentialsDropdownParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) WithContext(ctx context.Context) *ShowbackShowbackCredentialsDropdownParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) WithHTTPClient(client *http.Client) *ShowbackShowbackCredentialsDropdownParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) WithOrganizationID(organizationID *int32) *ShowbackShowbackCredentialsDropdownParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) WithV(v string) *ShowbackShowbackCredentialsDropdownParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback showback credentials dropdown params
func (o *ShowbackShowbackCredentialsDropdownParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackShowbackCredentialsDropdownParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
