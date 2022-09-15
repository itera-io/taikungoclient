// Code generated by go-swagger; DO NOT EDIT.

package prometheus

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

// NewPrometheusBillingListParams creates a new PrometheusBillingListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrometheusBillingListParams() *PrometheusBillingListParams {
	return &PrometheusBillingListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrometheusBillingListParamsWithTimeout creates a new PrometheusBillingListParams object
// with the ability to set a timeout on a request.
func NewPrometheusBillingListParamsWithTimeout(timeout time.Duration) *PrometheusBillingListParams {
	return &PrometheusBillingListParams{
		timeout: timeout,
	}
}

// NewPrometheusBillingListParamsWithContext creates a new PrometheusBillingListParams object
// with the ability to set a context for a request.
func NewPrometheusBillingListParamsWithContext(ctx context.Context) *PrometheusBillingListParams {
	return &PrometheusBillingListParams{
		Context: ctx,
	}
}

// NewPrometheusBillingListParamsWithHTTPClient creates a new PrometheusBillingListParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrometheusBillingListParamsWithHTTPClient(client *http.Client) *PrometheusBillingListParams {
	return &PrometheusBillingListParams{
		HTTPClient: client,
	}
}

/*
PrometheusBillingListParams contains all the parameters to send to the API endpoint

	for the prometheus billing list operation.

	Typically these are written to a http.Request.
*/
type PrometheusBillingListParams struct {

	// EndDate.
	//
	// Format: date-time
	EndDate *strfmt.DateTime

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

	// SortBy.
	SortBy *string

	// SortDirection.
	SortDirection *string

	// StartDate.
	//
	// Format: date-time
	StartDate *strfmt.DateTime

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the prometheus billing list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBillingListParams) WithDefaults() *PrometheusBillingListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the prometheus billing list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusBillingListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the prometheus billing list params
func (o *PrometheusBillingListParams) WithTimeout(timeout time.Duration) *PrometheusBillingListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prometheus billing list params
func (o *PrometheusBillingListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prometheus billing list params
func (o *PrometheusBillingListParams) WithContext(ctx context.Context) *PrometheusBillingListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prometheus billing list params
func (o *PrometheusBillingListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prometheus billing list params
func (o *PrometheusBillingListParams) WithHTTPClient(client *http.Client) *PrometheusBillingListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prometheus billing list params
func (o *PrometheusBillingListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the prometheus billing list params
func (o *PrometheusBillingListParams) WithEndDate(endDate *strfmt.DateTime) *PrometheusBillingListParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the prometheus billing list params
func (o *PrometheusBillingListParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithLimit adds the limit to the prometheus billing list params
func (o *PrometheusBillingListParams) WithLimit(limit *int32) *PrometheusBillingListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the prometheus billing list params
func (o *PrometheusBillingListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the prometheus billing list params
func (o *PrometheusBillingListParams) WithOffset(offset *int32) *PrometheusBillingListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the prometheus billing list params
func (o *PrometheusBillingListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the prometheus billing list params
func (o *PrometheusBillingListParams) WithOrganizationID(organizationID *int32) *PrometheusBillingListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the prometheus billing list params
func (o *PrometheusBillingListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSortBy adds the sortBy to the prometheus billing list params
func (o *PrometheusBillingListParams) WithSortBy(sortBy *string) *PrometheusBillingListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the prometheus billing list params
func (o *PrometheusBillingListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the prometheus billing list params
func (o *PrometheusBillingListParams) WithSortDirection(sortDirection *string) *PrometheusBillingListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the prometheus billing list params
func (o *PrometheusBillingListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartDate adds the startDate to the prometheus billing list params
func (o *PrometheusBillingListParams) WithStartDate(startDate *strfmt.DateTime) *PrometheusBillingListParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the prometheus billing list params
func (o *PrometheusBillingListParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithV adds the v to the prometheus billing list params
func (o *PrometheusBillingListParams) WithV(v string) *PrometheusBillingListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the prometheus billing list params
func (o *PrometheusBillingListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PrometheusBillingListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.DateTime

		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {

			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
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

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.DateTime

		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {

			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
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
