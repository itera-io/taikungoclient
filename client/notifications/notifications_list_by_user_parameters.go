// Code generated by go-swagger; DO NOT EDIT.

package notifications

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

// NewNotificationsListByUserParams creates a new NotificationsListByUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNotificationsListByUserParams() *NotificationsListByUserParams {
	return &NotificationsListByUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNotificationsListByUserParamsWithTimeout creates a new NotificationsListByUserParams object
// with the ability to set a timeout on a request.
func NewNotificationsListByUserParamsWithTimeout(timeout time.Duration) *NotificationsListByUserParams {
	return &NotificationsListByUserParams{
		timeout: timeout,
	}
}

// NewNotificationsListByUserParamsWithContext creates a new NotificationsListByUserParams object
// with the ability to set a context for a request.
func NewNotificationsListByUserParamsWithContext(ctx context.Context) *NotificationsListByUserParams {
	return &NotificationsListByUserParams{
		Context: ctx,
	}
}

// NewNotificationsListByUserParamsWithHTTPClient creates a new NotificationsListByUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewNotificationsListByUserParamsWithHTTPClient(client *http.Client) *NotificationsListByUserParams {
	return &NotificationsListByUserParams{
		HTTPClient: client,
	}
}

/* NotificationsListByUserParams contains all the parameters to send to the API endpoint
   for the notifications list by user operation.

   Typically these are written to a http.Request.
*/
type NotificationsListByUserParams struct {

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

	   Page number

	   Format: int32
	*/
	Offset *int32

	// Search.
	Search *string

	// StartDate.
	//
	// Format: date-time
	StartDate *strfmt.DateTime

	// Type.
	Type *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the notifications list by user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsListByUserParams) WithDefaults() *NotificationsListByUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the notifications list by user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NotificationsListByUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the notifications list by user params
func (o *NotificationsListByUserParams) WithTimeout(timeout time.Duration) *NotificationsListByUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the notifications list by user params
func (o *NotificationsListByUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the notifications list by user params
func (o *NotificationsListByUserParams) WithContext(ctx context.Context) *NotificationsListByUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the notifications list by user params
func (o *NotificationsListByUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the notifications list by user params
func (o *NotificationsListByUserParams) WithHTTPClient(client *http.Client) *NotificationsListByUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the notifications list by user params
func (o *NotificationsListByUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the notifications list by user params
func (o *NotificationsListByUserParams) WithEndDate(endDate *strfmt.DateTime) *NotificationsListByUserParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the notifications list by user params
func (o *NotificationsListByUserParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithLimit adds the limit to the notifications list by user params
func (o *NotificationsListByUserParams) WithLimit(limit *int32) *NotificationsListByUserParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the notifications list by user params
func (o *NotificationsListByUserParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the notifications list by user params
func (o *NotificationsListByUserParams) WithOffset(offset *int32) *NotificationsListByUserParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the notifications list by user params
func (o *NotificationsListByUserParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearch adds the search to the notifications list by user params
func (o *NotificationsListByUserParams) WithSearch(search *string) *NotificationsListByUserParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the notifications list by user params
func (o *NotificationsListByUserParams) SetSearch(search *string) {
	o.Search = search
}

// WithStartDate adds the startDate to the notifications list by user params
func (o *NotificationsListByUserParams) WithStartDate(startDate *strfmt.DateTime) *NotificationsListByUserParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the notifications list by user params
func (o *NotificationsListByUserParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithType adds the typeVar to the notifications list by user params
func (o *NotificationsListByUserParams) WithType(typeVar *string) *NotificationsListByUserParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the notifications list by user params
func (o *NotificationsListByUserParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithV adds the v to the notifications list by user params
func (o *NotificationsListByUserParams) WithV(v string) *NotificationsListByUserParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the notifications list by user params
func (o *NotificationsListByUserParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *NotificationsListByUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
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
