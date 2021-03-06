// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewAzureOffersParams creates a new AzureOffersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureOffersParams() *AzureOffersParams {
	return &AzureOffersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureOffersParamsWithTimeout creates a new AzureOffersParams object
// with the ability to set a timeout on a request.
func NewAzureOffersParamsWithTimeout(timeout time.Duration) *AzureOffersParams {
	return &AzureOffersParams{
		timeout: timeout,
	}
}

// NewAzureOffersParamsWithContext creates a new AzureOffersParams object
// with the ability to set a context for a request.
func NewAzureOffersParamsWithContext(ctx context.Context) *AzureOffersParams {
	return &AzureOffersParams{
		Context: ctx,
	}
}

// NewAzureOffersParamsWithHTTPClient creates a new AzureOffersParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureOffersParamsWithHTTPClient(client *http.Client) *AzureOffersParams {
	return &AzureOffersParams{
		HTTPClient: client,
	}
}

/* AzureOffersParams contains all the parameters to send to the API endpoint
   for the azure offers operation.

   Typically these are written to a http.Request.
*/
type AzureOffersParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	// Limit.
	//
	// Format: int32
	Limit *int32

	// Offset.
	//
	// Format: int32
	Offset *int32

	// Publisher.
	Publisher string

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

// WithDefaults hydrates default values in the azure offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureOffersParams) WithDefaults() *AzureOffersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure offers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureOffersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the azure offers params
func (o *AzureOffersParams) WithTimeout(timeout time.Duration) *AzureOffersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure offers params
func (o *AzureOffersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure offers params
func (o *AzureOffersParams) WithContext(ctx context.Context) *AzureOffersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure offers params
func (o *AzureOffersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure offers params
func (o *AzureOffersParams) WithHTTPClient(client *http.Client) *AzureOffersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure offers params
func (o *AzureOffersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the azure offers params
func (o *AzureOffersParams) WithCloudID(cloudID int32) *AzureOffersParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the azure offers params
func (o *AzureOffersParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithLimit adds the limit to the azure offers params
func (o *AzureOffersParams) WithLimit(limit *int32) *AzureOffersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the azure offers params
func (o *AzureOffersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the azure offers params
func (o *AzureOffersParams) WithOffset(offset *int32) *AzureOffersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the azure offers params
func (o *AzureOffersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPublisher adds the publisher to the azure offers params
func (o *AzureOffersParams) WithPublisher(publisher string) *AzureOffersParams {
	o.SetPublisher(publisher)
	return o
}

// SetPublisher adds the publisher to the azure offers params
func (o *AzureOffersParams) SetPublisher(publisher string) {
	o.Publisher = publisher
}

// WithSearch adds the search to the azure offers params
func (o *AzureOffersParams) WithSearch(search *string) *AzureOffersParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the azure offers params
func (o *AzureOffersParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the azure offers params
func (o *AzureOffersParams) WithSortBy(sortBy *string) *AzureOffersParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the azure offers params
func (o *AzureOffersParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the azure offers params
func (o *AzureOffersParams) WithSortDirection(sortDirection *string) *AzureOffersParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the azure offers params
func (o *AzureOffersParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithV adds the v to the azure offers params
func (o *AzureOffersParams) WithV(v string) *AzureOffersParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the azure offers params
func (o *AzureOffersParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AzureOffersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param publisher
	if err := r.SetPathParam("publisher", o.Publisher); err != nil {
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
