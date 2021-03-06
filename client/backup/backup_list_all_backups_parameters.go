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

// NewBackupListAllBackupsParams creates a new BackupListAllBackupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupListAllBackupsParams() *BackupListAllBackupsParams {
	return &BackupListAllBackupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupListAllBackupsParamsWithTimeout creates a new BackupListAllBackupsParams object
// with the ability to set a timeout on a request.
func NewBackupListAllBackupsParamsWithTimeout(timeout time.Duration) *BackupListAllBackupsParams {
	return &BackupListAllBackupsParams{
		timeout: timeout,
	}
}

// NewBackupListAllBackupsParamsWithContext creates a new BackupListAllBackupsParams object
// with the ability to set a context for a request.
func NewBackupListAllBackupsParamsWithContext(ctx context.Context) *BackupListAllBackupsParams {
	return &BackupListAllBackupsParams{
		Context: ctx,
	}
}

// NewBackupListAllBackupsParamsWithHTTPClient creates a new BackupListAllBackupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupListAllBackupsParamsWithHTTPClient(client *http.Client) *BackupListAllBackupsParams {
	return &BackupListAllBackupsParams{
		HTTPClient: client,
	}
}

/* BackupListAllBackupsParams contains all the parameters to send to the API endpoint
   for the backup list all backups operation.

   Typically these are written to a http.Request.
*/
type BackupListAllBackupsParams struct {

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

// WithDefaults hydrates default values in the backup list all backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupListAllBackupsParams) WithDefaults() *BackupListAllBackupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup list all backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupListAllBackupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup list all backups params
func (o *BackupListAllBackupsParams) WithTimeout(timeout time.Duration) *BackupListAllBackupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup list all backups params
func (o *BackupListAllBackupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup list all backups params
func (o *BackupListAllBackupsParams) WithContext(ctx context.Context) *BackupListAllBackupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup list all backups params
func (o *BackupListAllBackupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup list all backups params
func (o *BackupListAllBackupsParams) WithHTTPClient(client *http.Client) *BackupListAllBackupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup list all backups params
func (o *BackupListAllBackupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the backup list all backups params
func (o *BackupListAllBackupsParams) WithLimit(limit *int32) *BackupListAllBackupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the backup list all backups params
func (o *BackupListAllBackupsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the backup list all backups params
func (o *BackupListAllBackupsParams) WithOffset(offset *int32) *BackupListAllBackupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the backup list all backups params
func (o *BackupListAllBackupsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithProjectID adds the projectID to the backup list all backups params
func (o *BackupListAllBackupsParams) WithProjectID(projectID int32) *BackupListAllBackupsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the backup list all backups params
func (o *BackupListAllBackupsParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the backup list all backups params
func (o *BackupListAllBackupsParams) WithSearch(search *string) *BackupListAllBackupsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the backup list all backups params
func (o *BackupListAllBackupsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the backup list all backups params
func (o *BackupListAllBackupsParams) WithSortBy(sortBy *string) *BackupListAllBackupsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the backup list all backups params
func (o *BackupListAllBackupsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the backup list all backups params
func (o *BackupListAllBackupsParams) WithSortDirection(sortDirection *string) *BackupListAllBackupsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the backup list all backups params
func (o *BackupListAllBackupsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the backup list all backups params
func (o *BackupListAllBackupsParams) WithV(v string) *BackupListAllBackupsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the backup list all backups params
func (o *BackupListAllBackupsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BackupListAllBackupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
