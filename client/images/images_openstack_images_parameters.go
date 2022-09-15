// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewImagesOpenstackImagesParams creates a new ImagesOpenstackImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImagesOpenstackImagesParams() *ImagesOpenstackImagesParams {
	return &ImagesOpenstackImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImagesOpenstackImagesParamsWithTimeout creates a new ImagesOpenstackImagesParams object
// with the ability to set a timeout on a request.
func NewImagesOpenstackImagesParamsWithTimeout(timeout time.Duration) *ImagesOpenstackImagesParams {
	return &ImagesOpenstackImagesParams{
		timeout: timeout,
	}
}

// NewImagesOpenstackImagesParamsWithContext creates a new ImagesOpenstackImagesParams object
// with the ability to set a context for a request.
func NewImagesOpenstackImagesParamsWithContext(ctx context.Context) *ImagesOpenstackImagesParams {
	return &ImagesOpenstackImagesParams{
		Context: ctx,
	}
}

// NewImagesOpenstackImagesParamsWithHTTPClient creates a new ImagesOpenstackImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImagesOpenstackImagesParamsWithHTTPClient(client *http.Client) *ImagesOpenstackImagesParams {
	return &ImagesOpenstackImagesParams{
		HTTPClient: client,
	}
}

/*
ImagesOpenstackImagesParams contains all the parameters to send to the API endpoint

	for the images openstack images operation.

	Typically these are written to a http.Request.
*/
type ImagesOpenstackImagesParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	/* Limit.

	   Limits size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Skip elements

	   Format: int32
	*/
	Offset *int32

	// Personal.
	Personal *bool

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

// WithDefaults hydrates default values in the images openstack images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesOpenstackImagesParams) WithDefaults() *ImagesOpenstackImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the images openstack images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesOpenstackImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithTimeout(timeout time.Duration) *ImagesOpenstackImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithContext(ctx context.Context) *ImagesOpenstackImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithHTTPClient(client *http.Client) *ImagesOpenstackImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithCloudID(cloudID int32) *ImagesOpenstackImagesParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithLimit adds the limit to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithLimit(limit *int32) *ImagesOpenstackImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithOffset(offset *int32) *ImagesOpenstackImagesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPersonal adds the personal to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithPersonal(personal *bool) *ImagesOpenstackImagesParams {
	o.SetPersonal(personal)
	return o
}

// SetPersonal adds the personal to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetPersonal(personal *bool) {
	o.Personal = personal
}

// WithSearch adds the search to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithSearch(search *string) *ImagesOpenstackImagesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithSortBy(sortBy *string) *ImagesOpenstackImagesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithSortDirection(sortDirection *string) *ImagesOpenstackImagesParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the images openstack images params
func (o *ImagesOpenstackImagesParams) WithV(v string) *ImagesOpenstackImagesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the images openstack images params
func (o *ImagesOpenstackImagesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesOpenstackImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudId
	if err := r.SetPathParam("cloudId", swag.FormatInt32(o.CloudID)); err != nil {
		return err
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

	if o.Personal != nil {

		// query param personal
		var qrPersonal bool

		if o.Personal != nil {
			qrPersonal = *o.Personal
		}
		qPersonal := swag.FormatBool(qrPersonal)
		if qPersonal != "" {

			if err := r.SetQueryParam("personal", qPersonal); err != nil {
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
