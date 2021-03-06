// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewProjectGroupsListParams creates a new ProjectGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGroupsListParams() *ProjectGroupsListParams {
	return &ProjectGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGroupsListParamsWithTimeout creates a new ProjectGroupsListParams object
// with the ability to set a timeout on a request.
func NewProjectGroupsListParamsWithTimeout(timeout time.Duration) *ProjectGroupsListParams {
	return &ProjectGroupsListParams{
		timeout: timeout,
	}
}

// NewProjectGroupsListParamsWithContext creates a new ProjectGroupsListParams object
// with the ability to set a context for a request.
func NewProjectGroupsListParamsWithContext(ctx context.Context) *ProjectGroupsListParams {
	return &ProjectGroupsListParams{
		Context: ctx,
	}
}

// NewProjectGroupsListParamsWithHTTPClient creates a new ProjectGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGroupsListParamsWithHTTPClient(client *http.Client) *ProjectGroupsListParams {
	return &ProjectGroupsListParams{
		HTTPClient: client,
	}
}

/* ProjectGroupsListParams contains all the parameters to send to the API endpoint
   for the project groups list operation.

   Typically these are written to a http.Request.
*/
type ProjectGroupsListParams struct {

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

// WithDefaults hydrates default values in the project groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsListParams) WithDefaults() *ProjectGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project groups list params
func (o *ProjectGroupsListParams) WithTimeout(timeout time.Duration) *ProjectGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project groups list params
func (o *ProjectGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project groups list params
func (o *ProjectGroupsListParams) WithContext(ctx context.Context) *ProjectGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project groups list params
func (o *ProjectGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project groups list params
func (o *ProjectGroupsListParams) WithHTTPClient(client *http.Client) *ProjectGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project groups list params
func (o *ProjectGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the project groups list params
func (o *ProjectGroupsListParams) WithID(id *int32) *ProjectGroupsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the project groups list params
func (o *ProjectGroupsListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the project groups list params
func (o *ProjectGroupsListParams) WithLimit(limit *int32) *ProjectGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the project groups list params
func (o *ProjectGroupsListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the project groups list params
func (o *ProjectGroupsListParams) WithOffset(offset *int32) *ProjectGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the project groups list params
func (o *ProjectGroupsListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the project groups list params
func (o *ProjectGroupsListParams) WithOrganizationID(organizationID *int32) *ProjectGroupsListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the project groups list params
func (o *ProjectGroupsListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the project groups list params
func (o *ProjectGroupsListParams) WithSearch(search *string) *ProjectGroupsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the project groups list params
func (o *ProjectGroupsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the project groups list params
func (o *ProjectGroupsListParams) WithSearchID(searchID *string) *ProjectGroupsListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the project groups list params
func (o *ProjectGroupsListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithSortBy adds the sortBy to the project groups list params
func (o *ProjectGroupsListParams) WithSortBy(sortBy *string) *ProjectGroupsListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the project groups list params
func (o *ProjectGroupsListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the project groups list params
func (o *ProjectGroupsListParams) WithSortDirection(sortDirection *string) *ProjectGroupsListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the project groups list params
func (o *ProjectGroupsListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the project groups list params
func (o *ProjectGroupsListParams) WithV(v string) *ProjectGroupsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project groups list params
func (o *ProjectGroupsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
