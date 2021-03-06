// Code generated by go-swagger; DO NOT EDIT.

package backup

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

// NewBackupListAllSchedulesParams creates a new BackupListAllSchedulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupListAllSchedulesParams() *BackupListAllSchedulesParams {
	return &BackupListAllSchedulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupListAllSchedulesParamsWithTimeout creates a new BackupListAllSchedulesParams object
// with the ability to set a timeout on a request.
func NewBackupListAllSchedulesParamsWithTimeout(timeout time.Duration) *BackupListAllSchedulesParams {
	return &BackupListAllSchedulesParams{
		timeout: timeout,
	}
}

// NewBackupListAllSchedulesParamsWithContext creates a new BackupListAllSchedulesParams object
// with the ability to set a context for a request.
func NewBackupListAllSchedulesParamsWithContext(ctx context.Context) *BackupListAllSchedulesParams {
	return &BackupListAllSchedulesParams{
		Context: ctx,
	}
}

// NewBackupListAllSchedulesParamsWithHTTPClient creates a new BackupListAllSchedulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupListAllSchedulesParamsWithHTTPClient(client *http.Client) *BackupListAllSchedulesParams {
	return &BackupListAllSchedulesParams{
		HTTPClient: client,
	}
}

/* BackupListAllSchedulesParams contains all the parameters to send to the API endpoint
   for the backup list all schedules operation.

   Typically these are written to a http.Request.
*/
type BackupListAllSchedulesParams struct {

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// ProjectID.
	//
	// Format: int32
	ProjectID int32

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

// WithDefaults hydrates default values in the backup list all schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupListAllSchedulesParams) WithDefaults() *BackupListAllSchedulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup list all schedules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupListAllSchedulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithTimeout(timeout time.Duration) *BackupListAllSchedulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithContext(ctx context.Context) *BackupListAllSchedulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithHTTPClient(client *http.Client) *BackupListAllSchedulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithLimit(limit *int32) *BackupListAllSchedulesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithOffset(offset *int32) *BackupListAllSchedulesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithProjectID adds the projectID to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithProjectID(projectID int32) *BackupListAllSchedulesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithSearch(search *string) *BackupListAllSchedulesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithSortBy(sortBy *string) *BackupListAllSchedulesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithSortDirection(sortDirection *string) *BackupListAllSchedulesParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the backup list all schedules params
func (o *BackupListAllSchedulesParams) WithV(v string) *BackupListAllSchedulesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the backup list all schedules params
func (o *BackupListAllSchedulesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BackupListAllSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param projectId
	if err := r.SetPathParam("projectId", swag.FormatInt32(o.ProjectID)); err != nil {
		return err
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
