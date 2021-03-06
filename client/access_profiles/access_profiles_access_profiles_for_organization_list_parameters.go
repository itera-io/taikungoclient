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

// NewAccessProfilesAccessProfilesForOrganizationListParams creates a new AccessProfilesAccessProfilesForOrganizationListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessProfilesAccessProfilesForOrganizationListParams() *AccessProfilesAccessProfilesForOrganizationListParams {
	return &AccessProfilesAccessProfilesForOrganizationListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessProfilesAccessProfilesForOrganizationListParamsWithTimeout creates a new AccessProfilesAccessProfilesForOrganizationListParams object
// with the ability to set a timeout on a request.
func NewAccessProfilesAccessProfilesForOrganizationListParamsWithTimeout(timeout time.Duration) *AccessProfilesAccessProfilesForOrganizationListParams {
	return &AccessProfilesAccessProfilesForOrganizationListParams{
		timeout: timeout,
	}
}

// NewAccessProfilesAccessProfilesForOrganizationListParamsWithContext creates a new AccessProfilesAccessProfilesForOrganizationListParams object
// with the ability to set a context for a request.
func NewAccessProfilesAccessProfilesForOrganizationListParamsWithContext(ctx context.Context) *AccessProfilesAccessProfilesForOrganizationListParams {
	return &AccessProfilesAccessProfilesForOrganizationListParams{
		Context: ctx,
	}
}

// NewAccessProfilesAccessProfilesForOrganizationListParamsWithHTTPClient creates a new AccessProfilesAccessProfilesForOrganizationListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessProfilesAccessProfilesForOrganizationListParamsWithHTTPClient(client *http.Client) *AccessProfilesAccessProfilesForOrganizationListParams {
	return &AccessProfilesAccessProfilesForOrganizationListParams{
		HTTPClient: client,
	}
}

/* AccessProfilesAccessProfilesForOrganizationListParams contains all the parameters to send to the API endpoint
   for the access profiles access profiles for organization list operation.

   Typically these are written to a http.Request.
*/
type AccessProfilesAccessProfilesForOrganizationListParams struct {

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

// WithDefaults hydrates default values in the access profiles access profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithDefaults() *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access profiles access profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithTimeout(timeout time.Duration) *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithContext(ctx context.Context) *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithHTTPClient(client *http.Client) *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithOrganizationID(organizationID *int32) *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WithV(v string) *AccessProfilesAccessProfilesForOrganizationListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the access profiles access profiles for organization list params
func (o *AccessProfilesAccessProfilesForOrganizationListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AccessProfilesAccessProfilesForOrganizationListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
