// Code generated by go-swagger; DO NOT EDIT.

package user_groups

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

// NewUserGroupsListParams creates a new UserGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupsListParams() *UserGroupsListParams {
	return &UserGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupsListParamsWithTimeout creates a new UserGroupsListParams object
// with the ability to set a timeout on a request.
func NewUserGroupsListParamsWithTimeout(timeout time.Duration) *UserGroupsListParams {
	return &UserGroupsListParams{
		timeout: timeout,
	}
}

// NewUserGroupsListParamsWithContext creates a new UserGroupsListParams object
// with the ability to set a context for a request.
func NewUserGroupsListParamsWithContext(ctx context.Context) *UserGroupsListParams {
	return &UserGroupsListParams{
		Context: ctx,
	}
}

// NewUserGroupsListParamsWithHTTPClient creates a new UserGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupsListParamsWithHTTPClient(client *http.Client) *UserGroupsListParams {
	return &UserGroupsListParams{
		HTTPClient: client,
	}
}

/* UserGroupsListParams contains all the parameters to send to the API endpoint
   for the user groups list operation.

   Typically these are written to a http.Request.
*/
type UserGroupsListParams struct {

	// ID.
	//
	// Format: int32
	ID *int32

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// Search.
	Search *string

	// SearchID.
	SearchID *string

	// SortBy.
	SortBy *string

	// SortDirection.
	SortDirection *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsListParams) WithDefaults() *UserGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user groups list params
func (o *UserGroupsListParams) WithTimeout(timeout time.Duration) *UserGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user groups list params
func (o *UserGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user groups list params
func (o *UserGroupsListParams) WithContext(ctx context.Context) *UserGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user groups list params
func (o *UserGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user groups list params
func (o *UserGroupsListParams) WithHTTPClient(client *http.Client) *UserGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user groups list params
func (o *UserGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user groups list params
func (o *UserGroupsListParams) WithID(id *int32) *UserGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the user groups list params
func (o *UserGroupsListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the user groups list params
func (o *UserGroupsListParams) WithLimit(limit *int32) *UserGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the user groups list params
func (o *UserGroupsListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the user groups list params
func (o *UserGroupsListParams) WithOffset(offset *int32) *UserGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the user groups list params
func (o *UserGroupsListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the user groups list params
func (o *UserGroupsListParams) WithOrganizationID(organizationID *int32) *UserGroupsListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the user groups list params
func (o *UserGroupsListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the user groups list params
func (o *UserGroupsListParams) WithSearch(search *string) *UserGroupsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the user groups list params
func (o *UserGroupsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the user groups list params
func (o *UserGroupsListParams) WithSearchID(searchID *string) *UserGroupsListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the user groups list params
func (o *UserGroupsListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithSortBy adds the sortBy to the user groups list params
func (o *UserGroupsListParams) WithSortBy(sortBy *string) *UserGroupsListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the user groups list params
func (o *UserGroupsListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the user groups list params
func (o *UserGroupsListParams) WithSortDirection(sortDirection *string) *UserGroupsListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the user groups list params
func (o *UserGroupsListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the user groups list params
func (o *UserGroupsListParams) WithV(v string) *UserGroupsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the user groups list params
func (o *UserGroupsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID int32

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt32(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

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

	if o.SearchID != nil {

		// query param searchId
		var qrSearchID string

		if o.SearchID != nil {
			qrSearchID = *o.SearchID
		}
		qSearchID := qrSearchID
		if qSearchID != "" {

			if err := r.SetQueryParam("searchId", qSearchID); err != nil {
				return err
			}
		}
	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortDirection != nil {

		// query param sortDirection
		var qrSortDirection string

		if o.SortDirection != nil {
			qrSortDirection = *o.SortDirection
		}
		qSortDirection := qrSortDirection
		if qSortDirection != "" {

			if err := r.SetQueryParam("sortDirection", qSortDirection); err != nil {
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
