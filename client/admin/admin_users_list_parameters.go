// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminUsersListParams creates a new AdminUsersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminUsersListParams() *AdminUsersListParams {
	return &AdminUsersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUsersListParamsWithTimeout creates a new AdminUsersListParams object
// with the ability to set a timeout on a request.
func NewAdminUsersListParamsWithTimeout(timeout time.Duration) *AdminUsersListParams {
	return &AdminUsersListParams{
		timeout: timeout,
	}
}

// NewAdminUsersListParamsWithContext creates a new AdminUsersListParams object
// with the ability to set a context for a request.
func NewAdminUsersListParamsWithContext(ctx context.Context) *AdminUsersListParams {
	return &AdminUsersListParams{
		Context: ctx,
	}
}

// NewAdminUsersListParamsWithHTTPClient creates a new AdminUsersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminUsersListParamsWithHTTPClient(client *http.Client) *AdminUsersListParams {
	return &AdminUsersListParams{
		HTTPClient: client,
	}
}

/* AdminUsersListParams contains all the parameters to send to the API endpoint
   for the admin users list operation.

   Typically these are written to a http.Request.
*/
type AdminUsersListParams struct {

	/* Limit.

	   Limits user size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Skip elements

	   Format: int32
	*/
	Offset *int32

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	/* Search.

	   Keyword for searching
	*/
	Search *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUsersListParams) WithDefaults() *AdminUsersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUsersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin users list params
func (o *AdminUsersListParams) WithTimeout(timeout time.Duration) *AdminUsersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin users list params
func (o *AdminUsersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin users list params
func (o *AdminUsersListParams) WithContext(ctx context.Context) *AdminUsersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin users list params
func (o *AdminUsersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin users list params
func (o *AdminUsersListParams) WithHTTPClient(client *http.Client) *AdminUsersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin users list params
func (o *AdminUsersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the admin users list params
func (o *AdminUsersListParams) WithLimit(limit *int32) *AdminUsersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin users list params
func (o *AdminUsersListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the admin users list params
func (o *AdminUsersListParams) WithOffset(offset *int32) *AdminUsersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the admin users list params
func (o *AdminUsersListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the admin users list params
func (o *AdminUsersListParams) WithOrganizationID(organizationID *int32) *AdminUsersListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the admin users list params
func (o *AdminUsersListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the admin users list params
func (o *AdminUsersListParams) WithSearch(search *string) *AdminUsersListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the admin users list params
func (o *AdminUsersListParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the admin users list params
func (o *AdminUsersListParams) WithV(v string) *AdminUsersListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the admin users list params
func (o *AdminUsersListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUsersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

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

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
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