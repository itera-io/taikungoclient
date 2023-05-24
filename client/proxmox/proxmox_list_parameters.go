// Code generated by go-swagger; DO NOT EDIT.

package proxmox

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

// NewProxmoxListParams creates a new ProxmoxListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProxmoxListParams() *ProxmoxListParams {
	return &ProxmoxListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProxmoxListParamsWithTimeout creates a new ProxmoxListParams object
// with the ability to set a timeout on a request.
func NewProxmoxListParamsWithTimeout(timeout time.Duration) *ProxmoxListParams {
	return &ProxmoxListParams{
		timeout: timeout,
	}
}

// NewProxmoxListParamsWithContext creates a new ProxmoxListParams object
// with the ability to set a context for a request.
func NewProxmoxListParamsWithContext(ctx context.Context) *ProxmoxListParams {
	return &ProxmoxListParams{
		Context: ctx,
	}
}

// NewProxmoxListParamsWithHTTPClient creates a new ProxmoxListParams object
// with the ability to set a custom HTTPClient for a request.
func NewProxmoxListParamsWithHTTPClient(client *http.Client) *ProxmoxListParams {
	return &ProxmoxListParams{
		HTTPClient: client,
	}
}

/*
ProxmoxListParams contains all the parameters to send to the API endpoint

	for the proxmox list operation.

	Typically these are written to a http.Request.
*/
type ProxmoxListParams struct {

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

// WithDefaults hydrates default values in the proxmox list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxListParams) WithDefaults() *ProxmoxListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the proxmox list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProxmoxListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the proxmox list params
func (o *ProxmoxListParams) WithTimeout(timeout time.Duration) *ProxmoxListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the proxmox list params
func (o *ProxmoxListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the proxmox list params
func (o *ProxmoxListParams) WithContext(ctx context.Context) *ProxmoxListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the proxmox list params
func (o *ProxmoxListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the proxmox list params
func (o *ProxmoxListParams) WithHTTPClient(client *http.Client) *ProxmoxListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the proxmox list params
func (o *ProxmoxListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the proxmox list params
func (o *ProxmoxListParams) WithID(id *int32) *ProxmoxListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the proxmox list params
func (o *ProxmoxListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the proxmox list params
func (o *ProxmoxListParams) WithLimit(limit *int32) *ProxmoxListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the proxmox list params
func (o *ProxmoxListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the proxmox list params
func (o *ProxmoxListParams) WithOffset(offset *int32) *ProxmoxListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the proxmox list params
func (o *ProxmoxListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the proxmox list params
func (o *ProxmoxListParams) WithOrganizationID(organizationID *int32) *ProxmoxListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the proxmox list params
func (o *ProxmoxListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the proxmox list params
func (o *ProxmoxListParams) WithSearch(search *string) *ProxmoxListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the proxmox list params
func (o *ProxmoxListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the proxmox list params
func (o *ProxmoxListParams) WithSearchID(searchID *string) *ProxmoxListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the proxmox list params
func (o *ProxmoxListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithSortBy adds the sortBy to the proxmox list params
func (o *ProxmoxListParams) WithSortBy(sortBy *string) *ProxmoxListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the proxmox list params
func (o *ProxmoxListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the proxmox list params
func (o *ProxmoxListParams) WithSortDirection(sortDirection *string) *ProxmoxListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the proxmox list params
func (o *ProxmoxListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the proxmox list params
func (o *ProxmoxListParams) WithV(v string) *ProxmoxListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the proxmox list params
func (o *ProxmoxListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProxmoxListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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