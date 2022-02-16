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

// NewShowbackRulesListParams creates a new ShowbackRulesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackRulesListParams() *ShowbackRulesListParams {
	return &ShowbackRulesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackRulesListParamsWithTimeout creates a new ShowbackRulesListParams object
// with the ability to set a timeout on a request.
func NewShowbackRulesListParamsWithTimeout(timeout time.Duration) *ShowbackRulesListParams {
	return &ShowbackRulesListParams{
		timeout: timeout,
	}
}

// NewShowbackRulesListParamsWithContext creates a new ShowbackRulesListParams object
// with the ability to set a context for a request.
func NewShowbackRulesListParamsWithContext(ctx context.Context) *ShowbackRulesListParams {
	return &ShowbackRulesListParams{
		Context: ctx,
	}
}

// NewShowbackRulesListParamsWithHTTPClient creates a new ShowbackRulesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackRulesListParamsWithHTTPClient(client *http.Client) *ShowbackRulesListParams {
	return &ShowbackRulesListParams{
		HTTPClient: client,
	}
}

/* ShowbackRulesListParams contains all the parameters to send to the API endpoint
   for the showback rules list operation.

   Typically these are written to a http.Request.
*/
type ShowbackRulesListParams struct {

	// ID.
	//
	// Format: int32
	ID *int32

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

	// Search.
	Search *string

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

// WithDefaults hydrates default values in the showback rules list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesListParams) WithDefaults() *ShowbackRulesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback rules list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback rules list params
func (o *ShowbackRulesListParams) WithTimeout(timeout time.Duration) *ShowbackRulesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback rules list params
func (o *ShowbackRulesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback rules list params
func (o *ShowbackRulesListParams) WithContext(ctx context.Context) *ShowbackRulesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback rules list params
func (o *ShowbackRulesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback rules list params
func (o *ShowbackRulesListParams) WithHTTPClient(client *http.Client) *ShowbackRulesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback rules list params
func (o *ShowbackRulesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the showback rules list params
func (o *ShowbackRulesListParams) WithID(id *int32) *ShowbackRulesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the showback rules list params
func (o *ShowbackRulesListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the showback rules list params
func (o *ShowbackRulesListParams) WithLimit(limit *int32) *ShowbackRulesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the showback rules list params
func (o *ShowbackRulesListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the showback rules list params
func (o *ShowbackRulesListParams) WithOffset(offset *int32) *ShowbackRulesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the showback rules list params
func (o *ShowbackRulesListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the showback rules list params
func (o *ShowbackRulesListParams) WithOrganizationID(organizationID *int32) *ShowbackRulesListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback rules list params
func (o *ShowbackRulesListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the showback rules list params
func (o *ShowbackRulesListParams) WithSearch(search *string) *ShowbackRulesListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the showback rules list params
func (o *ShowbackRulesListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the showback rules list params
func (o *ShowbackRulesListParams) WithSortBy(sortBy *string) *ShowbackRulesListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the showback rules list params
func (o *ShowbackRulesListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the showback rules list params
func (o *ShowbackRulesListParams) WithSortDirection(sortDirection *string) *ShowbackRulesListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the showback rules list params
func (o *ShowbackRulesListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the showback rules list params
func (o *ShowbackRulesListParams) WithV(v string) *ShowbackRulesListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback rules list params
func (o *ShowbackRulesListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackRulesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
