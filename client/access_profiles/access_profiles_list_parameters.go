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

// NewAccessProfilesListParams creates a new AccessProfilesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessProfilesListParams() *AccessProfilesListParams {
	return &AccessProfilesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessProfilesListParamsWithTimeout creates a new AccessProfilesListParams object
// with the ability to set a timeout on a request.
func NewAccessProfilesListParamsWithTimeout(timeout time.Duration) *AccessProfilesListParams {
	return &AccessProfilesListParams{
		timeout: timeout,
	}
}

// NewAccessProfilesListParamsWithContext creates a new AccessProfilesListParams object
// with the ability to set a context for a request.
func NewAccessProfilesListParamsWithContext(ctx context.Context) *AccessProfilesListParams {
	return &AccessProfilesListParams{
		Context: ctx,
	}
}

// NewAccessProfilesListParamsWithHTTPClient creates a new AccessProfilesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessProfilesListParamsWithHTTPClient(client *http.Client) *AccessProfilesListParams {
	return &AccessProfilesListParams{
		HTTPClient: client,
	}
}

/* AccessProfilesListParams contains all the parameters to send to the API endpoint
   for the access profiles list operation.

   Typically these are written to a http.Request.
*/
type AccessProfilesListParams struct {

	/* Limit.

	   Limits user size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Page number

	   Format: int32
	*/
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

// WithDefaults hydrates default values in the access profiles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesListParams) WithDefaults() *AccessProfilesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access profiles list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access profiles list params
func (o *AccessProfilesListParams) WithTimeout(timeout time.Duration) *AccessProfilesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access profiles list params
func (o *AccessProfilesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access profiles list params
func (o *AccessProfilesListParams) WithContext(ctx context.Context) *AccessProfilesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access profiles list params
func (o *AccessProfilesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access profiles list params
func (o *AccessProfilesListParams) WithHTTPClient(client *http.Client) *AccessProfilesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access profiles list params
func (o *AccessProfilesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the access profiles list params
func (o *AccessProfilesListParams) WithLimit(limit *int32) *AccessProfilesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the access profiles list params
func (o *AccessProfilesListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the access profiles list params
func (o *AccessProfilesListParams) WithOffset(offset *int32) *AccessProfilesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the access profiles list params
func (o *AccessProfilesListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the access profiles list params
func (o *AccessProfilesListParams) WithOrganizationID(organizationID *int32) *AccessProfilesListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the access profiles list params
func (o *AccessProfilesListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the access profiles list params
func (o *AccessProfilesListParams) WithSearch(search *string) *AccessProfilesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the access profiles list params
func (o *AccessProfilesListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the access profiles list params
func (o *AccessProfilesListParams) WithSearchID(searchID *string) *AccessProfilesListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the access profiles list params
func (o *AccessProfilesListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithSortBy adds the sortBy to the access profiles list params
func (o *AccessProfilesListParams) WithSortBy(sortBy *string) *AccessProfilesListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the access profiles list params
func (o *AccessProfilesListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the access profiles list params
func (o *AccessProfilesListParams) WithSortDirection(sortDirection *string) *AccessProfilesListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the access profiles list params
func (o *AccessProfilesListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the access profiles list params
func (o *AccessProfilesListParams) WithV(v string) *AccessProfilesListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the access profiles list params
func (o *AccessProfilesListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AccessProfilesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
