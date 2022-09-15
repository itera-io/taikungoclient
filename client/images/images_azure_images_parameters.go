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

// NewImagesAzureImagesParams creates a new ImagesAzureImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImagesAzureImagesParams() *ImagesAzureImagesParams {
	return &ImagesAzureImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImagesAzureImagesParamsWithTimeout creates a new ImagesAzureImagesParams object
// with the ability to set a timeout on a request.
func NewImagesAzureImagesParamsWithTimeout(timeout time.Duration) *ImagesAzureImagesParams {
	return &ImagesAzureImagesParams{
		timeout: timeout,
	}
}

// NewImagesAzureImagesParamsWithContext creates a new ImagesAzureImagesParams object
// with the ability to set a context for a request.
func NewImagesAzureImagesParamsWithContext(ctx context.Context) *ImagesAzureImagesParams {
	return &ImagesAzureImagesParams{
		Context: ctx,
	}
}

// NewImagesAzureImagesParamsWithHTTPClient creates a new ImagesAzureImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImagesAzureImagesParamsWithHTTPClient(client *http.Client) *ImagesAzureImagesParams {
	return &ImagesAzureImagesParams{
		HTTPClient: client,
	}
}

/*
ImagesAzureImagesParams contains all the parameters to send to the API endpoint

	for the images azure images operation.

	Typically these are written to a http.Request.
*/
type ImagesAzureImagesParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	// Latest.
	Latest *bool

	/* Limit.

	   Limits size (by default 50)

	   Format: int32
	*/
	Limit *int32

	// Offer.
	Offer string

	/* Offset.

	   Skip elements

	   Format: int32
	*/
	Offset *int32

	// PublisherName.
	PublisherName string

	// Search.
	Search *string

	// Sku.
	Sku string

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

// WithDefaults hydrates default values in the images azure images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesAzureImagesParams) WithDefaults() *ImagesAzureImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the images azure images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesAzureImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the images azure images params
func (o *ImagesAzureImagesParams) WithTimeout(timeout time.Duration) *ImagesAzureImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images azure images params
func (o *ImagesAzureImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images azure images params
func (o *ImagesAzureImagesParams) WithContext(ctx context.Context) *ImagesAzureImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images azure images params
func (o *ImagesAzureImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the images azure images params
func (o *ImagesAzureImagesParams) WithHTTPClient(client *http.Client) *ImagesAzureImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the images azure images params
func (o *ImagesAzureImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the images azure images params
func (o *ImagesAzureImagesParams) WithCloudID(cloudID int32) *ImagesAzureImagesParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the images azure images params
func (o *ImagesAzureImagesParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithLatest adds the latest to the images azure images params
func (o *ImagesAzureImagesParams) WithLatest(latest *bool) *ImagesAzureImagesParams {
	o.SetLatest(latest)
	return o
}

// SetLatest adds the latest to the images azure images params
func (o *ImagesAzureImagesParams) SetLatest(latest *bool) {
	o.Latest = latest
}

// WithLimit adds the limit to the images azure images params
func (o *ImagesAzureImagesParams) WithLimit(limit *int32) *ImagesAzureImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the images azure images params
func (o *ImagesAzureImagesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffer adds the offer to the images azure images params
func (o *ImagesAzureImagesParams) WithOffer(offer string) *ImagesAzureImagesParams {
	o.SetOffer(offer)
	return o
}

// SetOffer adds the offer to the images azure images params
func (o *ImagesAzureImagesParams) SetOffer(offer string) {
	o.Offer = offer
}

// WithOffset adds the offset to the images azure images params
func (o *ImagesAzureImagesParams) WithOffset(offset *int32) *ImagesAzureImagesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the images azure images params
func (o *ImagesAzureImagesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPublisherName adds the publisherName to the images azure images params
func (o *ImagesAzureImagesParams) WithPublisherName(publisherName string) *ImagesAzureImagesParams {
	o.SetPublisherName(publisherName)
	return o
}

// SetPublisherName adds the publisherName to the images azure images params
func (o *ImagesAzureImagesParams) SetPublisherName(publisherName string) {
	o.PublisherName = publisherName
}

// WithSearch adds the search to the images azure images params
func (o *ImagesAzureImagesParams) WithSearch(search *string) *ImagesAzureImagesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the images azure images params
func (o *ImagesAzureImagesParams) SetSearch(search *string) {
	o.Search = search
}

// WithSku adds the sku to the images azure images params
func (o *ImagesAzureImagesParams) WithSku(sku string) *ImagesAzureImagesParams {
	o.SetSku(sku)
	return o
}

// SetSku adds the sku to the images azure images params
func (o *ImagesAzureImagesParams) SetSku(sku string) {
	o.Sku = sku
}

// WithSortBy adds the sortBy to the images azure images params
func (o *ImagesAzureImagesParams) WithSortBy(sortBy *string) *ImagesAzureImagesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the images azure images params
func (o *ImagesAzureImagesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the images azure images params
func (o *ImagesAzureImagesParams) WithSortDirection(sortDirection *string) *ImagesAzureImagesParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the images azure images params
func (o *ImagesAzureImagesParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the images azure images params
func (o *ImagesAzureImagesParams) WithV(v string) *ImagesAzureImagesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the images azure images params
func (o *ImagesAzureImagesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesAzureImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudId
	if err := r.SetPathParam("cloudId", swag.FormatInt32(o.CloudID)); err != nil {
		return err
	}

	if o.Latest != nil {

		// query param latest
		var qrLatest bool

		if o.Latest != nil {
			qrLatest = *o.Latest
		}
		qLatest := swag.FormatBool(qrLatest)
		if qLatest != "" {

			if err := r.SetQueryParam("latest", qLatest); err != nil {
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

	// path param offer
	if err := r.SetPathParam("offer", o.Offer); err != nil {
		return err
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

	// path param publisherName
	if err := r.SetPathParam("publisherName", o.PublisherName); err != nil {
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

	// path param sku
	if err := r.SetPathParam("sku", o.Sku); err != nil {
		return err
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
