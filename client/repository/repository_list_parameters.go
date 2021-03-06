// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// NewRepositoryListParams creates a new RepositoryListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepositoryListParams() *RepositoryListParams {
	return &RepositoryListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepositoryListParamsWithTimeout creates a new RepositoryListParams object
// with the ability to set a timeout on a request.
func NewRepositoryListParamsWithTimeout(timeout time.Duration) *RepositoryListParams {
	return &RepositoryListParams{
		timeout: timeout,
	}
}

// NewRepositoryListParamsWithContext creates a new RepositoryListParams object
// with the ability to set a context for a request.
func NewRepositoryListParamsWithContext(ctx context.Context) *RepositoryListParams {
	return &RepositoryListParams{
		Context: ctx,
	}
}

// NewRepositoryListParamsWithHTTPClient creates a new RepositoryListParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepositoryListParamsWithHTTPClient(client *http.Client) *RepositoryListParams {
	return &RepositoryListParams{
		HTTPClient: client,
	}
}

/* RepositoryListParams contains all the parameters to send to the API endpoint
   for the repository list operation.

   Typically these are written to a http.Request.
*/
type RepositoryListParams struct {

	// ID.
	ID *string

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

// WithDefaults hydrates default values in the repository list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryListParams) WithDefaults() *RepositoryListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repository list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepositoryListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repository list params
func (o *RepositoryListParams) WithTimeout(timeout time.Duration) *RepositoryListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repository list params
func (o *RepositoryListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repository list params
func (o *RepositoryListParams) WithContext(ctx context.Context) *RepositoryListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repository list params
func (o *RepositoryListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repository list params
func (o *RepositoryListParams) WithHTTPClient(client *http.Client) *RepositoryListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repository list params
func (o *RepositoryListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the repository list params
func (o *RepositoryListParams) WithID(id *string) *RepositoryListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the repository list params
func (o *RepositoryListParams) SetID(id *string) {
	o.ID = id
}

// WithLimit adds the limit to the repository list params
func (o *RepositoryListParams) WithLimit(limit *int32) *RepositoryListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the repository list params
func (o *RepositoryListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the repository list params
func (o *RepositoryListParams) WithOffset(offset *int32) *RepositoryListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the repository list params
func (o *RepositoryListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearch adds the search to the repository list params
func (o *RepositoryListParams) WithSearch(search *string) *RepositoryListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the repository list params
func (o *RepositoryListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the repository list params
func (o *RepositoryListParams) WithSortBy(sortBy *string) *RepositoryListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the repository list params
func (o *RepositoryListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the repository list params
func (o *RepositoryListParams) WithSortDirection(sortDirection *string) *RepositoryListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the repository list params
func (o *RepositoryListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the repository list params
func (o *RepositoryListParams) WithV(v string) *RepositoryListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the repository list params
func (o *RepositoryListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *RepositoryListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
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
