// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewKubernetesGetPersistentVolumeClaimParams creates a new KubernetesGetPersistentVolumeClaimParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesGetPersistentVolumeClaimParams() *KubernetesGetPersistentVolumeClaimParams {
	return &KubernetesGetPersistentVolumeClaimParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesGetPersistentVolumeClaimParamsWithTimeout creates a new KubernetesGetPersistentVolumeClaimParams object
// with the ability to set a timeout on a request.
func NewKubernetesGetPersistentVolumeClaimParamsWithTimeout(timeout time.Duration) *KubernetesGetPersistentVolumeClaimParams {
	return &KubernetesGetPersistentVolumeClaimParams{
		timeout: timeout,
	}
}

// NewKubernetesGetPersistentVolumeClaimParamsWithContext creates a new KubernetesGetPersistentVolumeClaimParams object
// with the ability to set a context for a request.
func NewKubernetesGetPersistentVolumeClaimParamsWithContext(ctx context.Context) *KubernetesGetPersistentVolumeClaimParams {
	return &KubernetesGetPersistentVolumeClaimParams{
		Context: ctx,
	}
}

// NewKubernetesGetPersistentVolumeClaimParamsWithHTTPClient creates a new KubernetesGetPersistentVolumeClaimParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesGetPersistentVolumeClaimParamsWithHTTPClient(client *http.Client) *KubernetesGetPersistentVolumeClaimParams {
	return &KubernetesGetPersistentVolumeClaimParams{
		HTTPClient: client,
	}
}

/* KubernetesGetPersistentVolumeClaimParams contains all the parameters to send to the API endpoint
   for the kubernetes get persistent volume claim operation.

   Typically these are written to a http.Request.
*/
type KubernetesGetPersistentVolumeClaimParams struct {

	// FilterBy.
	FilterBy *string

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

	// ProjectID.
	//
	// Format: int32
	ProjectID int32

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

// WithDefaults hydrates default values in the kubernetes get persistent volume claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetPersistentVolumeClaimParams) WithDefaults() *KubernetesGetPersistentVolumeClaimParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes get persistent volume claim params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetPersistentVolumeClaimParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithTimeout(timeout time.Duration) *KubernetesGetPersistentVolumeClaimParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithContext(ctx context.Context) *KubernetesGetPersistentVolumeClaimParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithHTTPClient(client *http.Client) *KubernetesGetPersistentVolumeClaimParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilterBy adds the filterBy to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithFilterBy(filterBy *string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetFilterBy(filterBy)
	return o
}

// SetFilterBy adds the filterBy to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetFilterBy(filterBy *string) {
	o.FilterBy = filterBy
}

// WithLimit adds the limit to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithLimit(limit *int32) *KubernetesGetPersistentVolumeClaimParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithOffset(offset *int32) *KubernetesGetPersistentVolumeClaimParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithProjectID adds the projectID to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithProjectID(projectID int32) *KubernetesGetPersistentVolumeClaimParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithSearch(search *string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithSearchID(searchID *string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithSortBy adds the sortBy to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithSortBy(sortBy *string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithSortDirection(sortDirection *string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) WithV(v string) *KubernetesGetPersistentVolumeClaimParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes get persistent volume claim params
func (o *KubernetesGetPersistentVolumeClaimParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesGetPersistentVolumeClaimParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterBy != nil {

		// query param filterBy
		var qrFilterBy string

		if o.FilterBy != nil {
			qrFilterBy = *o.FilterBy
		}
		qFilterBy := qrFilterBy
		if qFilterBy != "" {

			if err := r.SetQueryParam("filterBy", qFilterBy); err != nil {
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
