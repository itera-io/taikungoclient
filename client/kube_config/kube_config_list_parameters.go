// Code generated by go-swagger; DO NOT EDIT.

package kube_config

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

// NewKubeConfigListParams creates a new KubeConfigListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubeConfigListParams() *KubeConfigListParams {
	return &KubeConfigListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubeConfigListParamsWithTimeout creates a new KubeConfigListParams object
// with the ability to set a timeout on a request.
func NewKubeConfigListParamsWithTimeout(timeout time.Duration) *KubeConfigListParams {
	return &KubeConfigListParams{
		timeout: timeout,
	}
}

// NewKubeConfigListParamsWithContext creates a new KubeConfigListParams object
// with the ability to set a context for a request.
func NewKubeConfigListParamsWithContext(ctx context.Context) *KubeConfigListParams {
	return &KubeConfigListParams{
		Context: ctx,
	}
}

// NewKubeConfigListParamsWithHTTPClient creates a new KubeConfigListParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubeConfigListParamsWithHTTPClient(client *http.Client) *KubeConfigListParams {
	return &KubeConfigListParams{
		HTTPClient: client,
	}
}

/* KubeConfigListParams contains all the parameters to send to the API endpoint
   for the kube config list operation.

   Typically these are written to a http.Request.
*/
type KubeConfigListParams struct {

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

	/* OrganizationID.

	   Only for admin sort by org id

	   Format: int32
	*/
	OrganizationID *int32

	/* ProjectID.

	   Get kube configs by projectId

	   Format: int32
	*/
	ProjectID *int32

	/* Search.

	   Keyword for searching
	*/
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

// WithDefaults hydrates default values in the kube config list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubeConfigListParams) WithDefaults() *KubeConfigListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kube config list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubeConfigListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kube config list params
func (o *KubeConfigListParams) WithTimeout(timeout time.Duration) *KubeConfigListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kube config list params
func (o *KubeConfigListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kube config list params
func (o *KubeConfigListParams) WithContext(ctx context.Context) *KubeConfigListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kube config list params
func (o *KubeConfigListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kube config list params
func (o *KubeConfigListParams) WithHTTPClient(client *http.Client) *KubeConfigListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kube config list params
func (o *KubeConfigListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the kube config list params
func (o *KubeConfigListParams) WithLimit(limit *int32) *KubeConfigListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the kube config list params
func (o *KubeConfigListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the kube config list params
func (o *KubeConfigListParams) WithOffset(offset *int32) *KubeConfigListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the kube config list params
func (o *KubeConfigListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the kube config list params
func (o *KubeConfigListParams) WithOrganizationID(organizationID *int32) *KubeConfigListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the kube config list params
func (o *KubeConfigListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the kube config list params
func (o *KubeConfigListParams) WithProjectID(projectID *int32) *KubeConfigListParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kube config list params
func (o *KubeConfigListParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the kube config list params
func (o *KubeConfigListParams) WithSearch(search *string) *KubeConfigListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the kube config list params
func (o *KubeConfigListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the kube config list params
func (o *KubeConfigListParams) WithSortBy(sortBy *string) *KubeConfigListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the kube config list params
func (o *KubeConfigListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the kube config list params
func (o *KubeConfigListParams) WithSortDirection(sortDirection *string) *KubeConfigListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the kube config list params
func (o *KubeConfigListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the kube config list params
func (o *KubeConfigListParams) WithV(v string) *KubeConfigListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kube config list params
func (o *KubeConfigListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubeConfigListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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
