// Code generated by go-swagger; DO NOT EDIT.

package ticket

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

// NewTicketTicketListParams creates a new TicketTicketListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTicketTicketListParams() *TicketTicketListParams {
	return &TicketTicketListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTicketTicketListParamsWithTimeout creates a new TicketTicketListParams object
// with the ability to set a timeout on a request.
func NewTicketTicketListParamsWithTimeout(timeout time.Duration) *TicketTicketListParams {
	return &TicketTicketListParams{
		timeout: timeout,
	}
}

// NewTicketTicketListParamsWithContext creates a new TicketTicketListParams object
// with the ability to set a context for a request.
func NewTicketTicketListParamsWithContext(ctx context.Context) *TicketTicketListParams {
	return &TicketTicketListParams{
		Context: ctx,
	}
}

// NewTicketTicketListParamsWithHTTPClient creates a new TicketTicketListParams object
// with the ability to set a custom HTTPClient for a request.
func NewTicketTicketListParamsWithHTTPClient(client *http.Client) *TicketTicketListParams {
	return &TicketTicketListParams{
		HTTPClient: client,
	}
}

/* TicketTicketListParams contains all the parameters to send to the API endpoint
   for the ticket ticket list operation.

   Typically these are written to a http.Request.
*/
type TicketTicketListParams struct {

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

	/* Search.

	   Keyword for searching
	*/
	Search *string

	// StartDate.
	//
	// Format: date-time
	StartDate *strfmt.DateTime

	// TicketID.
	TicketID *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ticket ticket list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TicketTicketListParams) WithDefaults() *TicketTicketListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ticket ticket list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TicketTicketListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ticket ticket list params
func (o *TicketTicketListParams) WithTimeout(timeout time.Duration) *TicketTicketListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket ticket list params
func (o *TicketTicketListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket ticket list params
func (o *TicketTicketListParams) WithContext(ctx context.Context) *TicketTicketListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket ticket list params
func (o *TicketTicketListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ticket ticket list params
func (o *TicketTicketListParams) WithHTTPClient(client *http.Client) *TicketTicketListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket ticket list params
func (o *TicketTicketListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the ticket ticket list params
func (o *TicketTicketListParams) WithEndDate(endDate *strfmt.DateTime) *TicketTicketListParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the ticket ticket list params
func (o *TicketTicketListParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithLimit adds the limit to the ticket ticket list params
func (o *TicketTicketListParams) WithLimit(limit *int32) *TicketTicketListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ticket ticket list params
func (o *TicketTicketListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the ticket ticket list params
func (o *TicketTicketListParams) WithOffset(offset *int32) *TicketTicketListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ticket ticket list params
func (o *TicketTicketListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the ticket ticket list params
func (o *TicketTicketListParams) WithOrganizationID(organizationID *int32) *TicketTicketListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the ticket ticket list params
func (o *TicketTicketListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the ticket ticket list params
func (o *TicketTicketListParams) WithSearch(search *string) *TicketTicketListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the ticket ticket list params
func (o *TicketTicketListParams) SetSearch(search *string) {
	o.Search = search
}

// WithStartDate adds the startDate to the ticket ticket list params
func (o *TicketTicketListParams) WithStartDate(startDate *strfmt.DateTime) *TicketTicketListParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the ticket ticket list params
func (o *TicketTicketListParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithTicketID adds the ticketID to the ticket ticket list params
func (o *TicketTicketListParams) WithTicketID(ticketID *string) *TicketTicketListParams {
	o.SetTicketID(ticketID)
	return o
}

// SetTicketID adds the ticketId to the ticket ticket list params
func (o *TicketTicketListParams) SetTicketID(ticketID *string) {
	o.TicketID = ticketID
}

// WithV adds the v to the ticket ticket list params
func (o *TicketTicketListParams) WithV(v string) *TicketTicketListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ticket ticket list params
func (o *TicketTicketListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *TicketTicketListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TicketID != nil {

		// query param ticketId
		var qrTicketID string

		if o.TicketID != nil {
			qrTicketID = *o.TicketID
		}
		qTicketID := qrTicketID
		if qTicketID != "" {

			if err := r.SetQueryParam("ticketId", qTicketID); err != nil {
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
