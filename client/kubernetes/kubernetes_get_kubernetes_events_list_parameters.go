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

// NewKubernetesGetKubernetesEventsListParams creates a new KubernetesGetKubernetesEventsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesGetKubernetesEventsListParams() *KubernetesGetKubernetesEventsListParams {
	return &KubernetesGetKubernetesEventsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesGetKubernetesEventsListParamsWithTimeout creates a new KubernetesGetKubernetesEventsListParams object
// with the ability to set a timeout on a request.
func NewKubernetesGetKubernetesEventsListParamsWithTimeout(timeout time.Duration) *KubernetesGetKubernetesEventsListParams {
	return &KubernetesGetKubernetesEventsListParams{
		timeout: timeout,
	}
}

// NewKubernetesGetKubernetesEventsListParamsWithContext creates a new KubernetesGetKubernetesEventsListParams object
// with the ability to set a context for a request.
func NewKubernetesGetKubernetesEventsListParamsWithContext(ctx context.Context) *KubernetesGetKubernetesEventsListParams {
	return &KubernetesGetKubernetesEventsListParams{
		Context: ctx,
	}
}

// NewKubernetesGetKubernetesEventsListParamsWithHTTPClient creates a new KubernetesGetKubernetesEventsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesGetKubernetesEventsListParamsWithHTTPClient(client *http.Client) *KubernetesGetKubernetesEventsListParams {
	return &KubernetesGetKubernetesEventsListParams{
		HTTPClient: client,
	}
}

/* KubernetesGetKubernetesEventsListParams contains all the parameters to send to the API endpoint
   for the kubernetes get kubernetes events list operation.

   Typically these are written to a http.Request.
*/
type KubernetesGetKubernetesEventsListParams struct {

	// FirstTimeStamp.
	//
	// Format: date-time
	FirstTimeStamp *strfmt.DateTime

	// LastTimeStamp.
	//
	// Format: date-time
	LastTimeStamp *strfmt.DateTime

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

// WithDefaults hydrates default values in the kubernetes get kubernetes events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetKubernetesEventsListParams) WithDefaults() *KubernetesGetKubernetesEventsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes get kubernetes events list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesGetKubernetesEventsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithTimeout(timeout time.Duration) *KubernetesGetKubernetesEventsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithContext(ctx context.Context) *KubernetesGetKubernetesEventsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithHTTPClient(client *http.Client) *KubernetesGetKubernetesEventsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFirstTimeStamp adds the firstTimeStamp to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithFirstTimeStamp(firstTimeStamp *strfmt.DateTime) *KubernetesGetKubernetesEventsListParams {
	o.SetFirstTimeStamp(firstTimeStamp)
	return o
}

// SetFirstTimeStamp adds the firstTimeStamp to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetFirstTimeStamp(firstTimeStamp *strfmt.DateTime) {
	o.FirstTimeStamp = firstTimeStamp
}

// WithLastTimeStamp adds the lastTimeStamp to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithLastTimeStamp(lastTimeStamp *strfmt.DateTime) *KubernetesGetKubernetesEventsListParams {
	o.SetLastTimeStamp(lastTimeStamp)
	return o
}

// SetLastTimeStamp adds the lastTimeStamp to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetLastTimeStamp(lastTimeStamp *strfmt.DateTime) {
	o.LastTimeStamp = lastTimeStamp
}

// WithLimit adds the limit to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithLimit(limit *int32) *KubernetesGetKubernetesEventsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithOffset(offset *int32) *KubernetesGetKubernetesEventsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithProjectID adds the projectID to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithProjectID(projectID int32) *KubernetesGetKubernetesEventsListParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithSearch(search *string) *KubernetesGetKubernetesEventsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithSortBy(sortBy *string) *KubernetesGetKubernetesEventsListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithSortDirection(sortDirection *string) *KubernetesGetKubernetesEventsListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) WithV(v string) *KubernetesGetKubernetesEventsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes get kubernetes events list params
func (o *KubernetesGetKubernetesEventsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesGetKubernetesEventsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FirstTimeStamp != nil {

		// query param firstTimeStamp
		var qrFirstTimeStamp strfmt.DateTime

		if o.FirstTimeStamp != nil {
			qrFirstTimeStamp = *o.FirstTimeStamp
		}
		qFirstTimeStamp := qrFirstTimeStamp.String()
		if qFirstTimeStamp != "" {

			if err := r.SetQueryParam("firstTimeStamp", qFirstTimeStamp); err != nil {
				return err
			}
		}
	}

	if o.LastTimeStamp != nil {

		// query param lastTimeStamp
		var qrLastTimeStamp strfmt.DateTime

		if o.LastTimeStamp != nil {
			qrLastTimeStamp = *o.LastTimeStamp
		}
		qLastTimeStamp := qrLastTimeStamp.String()
		if qLastTimeStamp != "" {

			if err := r.SetQueryParam("lastTimeStamp", qLastTimeStamp); err != nil {
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
