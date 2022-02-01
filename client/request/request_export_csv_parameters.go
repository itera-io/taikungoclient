// Code generated by go-swagger; DO NOT EDIT.

package request

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

// NewRequestExportCsvParams creates a new RequestExportCsvParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRequestExportCsvParams() *RequestExportCsvParams {
	return &RequestExportCsvParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRequestExportCsvParamsWithTimeout creates a new RequestExportCsvParams object
// with the ability to set a timeout on a request.
func NewRequestExportCsvParamsWithTimeout(timeout time.Duration) *RequestExportCsvParams {
	return &RequestExportCsvParams{
		timeout: timeout,
	}
}

// NewRequestExportCsvParamsWithContext creates a new RequestExportCsvParams object
// with the ability to set a context for a request.
func NewRequestExportCsvParamsWithContext(ctx context.Context) *RequestExportCsvParams {
	return &RequestExportCsvParams{
		Context: ctx,
	}
}

// NewRequestExportCsvParamsWithHTTPClient creates a new RequestExportCsvParams object
// with the ability to set a custom HTTPClient for a request.
func NewRequestExportCsvParamsWithHTTPClient(client *http.Client) *RequestExportCsvParams {
	return &RequestExportCsvParams{
		HTTPClient: client,
	}
}

/* RequestExportCsvParams contains all the parameters to send to the API endpoint
   for the request export csv operation.

   Typically these are written to a http.Request.
*/
type RequestExportCsvParams struct {

	// EndDate.
	//
	// Format: date-time
	EndDate *strfmt.DateTime

	// FilterBy.
	FilterBy *string

	// IsEmailEnabled.
	IsEmailEnabled *bool

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// Search.
	Search *string

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

// WithDefaults hydrates default values in the request export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestExportCsvParams) WithDefaults() *RequestExportCsvParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the request export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RequestExportCsvParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the request export csv params
func (o *RequestExportCsvParams) WithTimeout(timeout time.Duration) *RequestExportCsvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request export csv params
func (o *RequestExportCsvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request export csv params
func (o *RequestExportCsvParams) WithContext(ctx context.Context) *RequestExportCsvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request export csv params
func (o *RequestExportCsvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request export csv params
func (o *RequestExportCsvParams) WithHTTPClient(client *http.Client) *RequestExportCsvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request export csv params
func (o *RequestExportCsvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the request export csv params
func (o *RequestExportCsvParams) WithEndDate(endDate *strfmt.DateTime) *RequestExportCsvParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the request export csv params
func (o *RequestExportCsvParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithFilterBy adds the filterBy to the request export csv params
func (o *RequestExportCsvParams) WithFilterBy(filterBy *string) *RequestExportCsvParams {
	o.SetFilterBy(filterBy)
	return o
}

// SetFilterBy adds the filterBy to the request export csv params
func (o *RequestExportCsvParams) SetFilterBy(filterBy *string) {
	o.FilterBy = filterBy
}

// WithIsEmailEnabled adds the isEmailEnabled to the request export csv params
func (o *RequestExportCsvParams) WithIsEmailEnabled(isEmailEnabled *bool) *RequestExportCsvParams {
	o.SetIsEmailEnabled(isEmailEnabled)
	return o
}

// SetIsEmailEnabled adds the isEmailEnabled to the request export csv params
func (o *RequestExportCsvParams) SetIsEmailEnabled(isEmailEnabled *bool) {
	o.IsEmailEnabled = isEmailEnabled
}

// WithOrganizationID adds the organizationID to the request export csv params
func (o *RequestExportCsvParams) WithOrganizationID(organizationID *int32) *RequestExportCsvParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the request export csv params
func (o *RequestExportCsvParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the request export csv params
func (o *RequestExportCsvParams) WithSearch(search *string) *RequestExportCsvParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the request export csv params
func (o *RequestExportCsvParams) SetSearch(search *string) {
	o.Search = search
}

// WithStartDate adds the startDate to the request export csv params
func (o *RequestExportCsvParams) WithStartDate(startDate *strfmt.DateTime) *RequestExportCsvParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the request export csv params
func (o *RequestExportCsvParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithV adds the v to the request export csv params
func (o *RequestExportCsvParams) WithV(v string) *RequestExportCsvParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the request export csv params
func (o *RequestExportCsvParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *RequestExportCsvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsEmailEnabled != nil {

		// query param isEmailEnabled
		var qrIsEmailEnabled bool

		if o.IsEmailEnabled != nil {
			qrIsEmailEnabled = *o.IsEmailEnabled
		}
		qIsEmailEnabled := swag.FormatBool(qrIsEmailEnabled)
		if qIsEmailEnabled != "" {

			if err := r.SetQueryParam("isEmailEnabled", qIsEmailEnabled); err != nil {
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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
