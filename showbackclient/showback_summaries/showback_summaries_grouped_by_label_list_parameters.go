// Code generated by go-swagger; DO NOT EDIT.

package showback_summaries

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

// NewShowbackSummariesGroupedByLabelListParams creates a new ShowbackSummariesGroupedByLabelListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackSummariesGroupedByLabelListParams() *ShowbackSummariesGroupedByLabelListParams {
	return &ShowbackSummariesGroupedByLabelListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackSummariesGroupedByLabelListParamsWithTimeout creates a new ShowbackSummariesGroupedByLabelListParams object
// with the ability to set a timeout on a request.
func NewShowbackSummariesGroupedByLabelListParamsWithTimeout(timeout time.Duration) *ShowbackSummariesGroupedByLabelListParams {
	return &ShowbackSummariesGroupedByLabelListParams{
		timeout: timeout,
	}
}

// NewShowbackSummariesGroupedByLabelListParamsWithContext creates a new ShowbackSummariesGroupedByLabelListParams object
// with the ability to set a context for a request.
func NewShowbackSummariesGroupedByLabelListParamsWithContext(ctx context.Context) *ShowbackSummariesGroupedByLabelListParams {
	return &ShowbackSummariesGroupedByLabelListParams{
		Context: ctx,
	}
}

// NewShowbackSummariesGroupedByLabelListParamsWithHTTPClient creates a new ShowbackSummariesGroupedByLabelListParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackSummariesGroupedByLabelListParamsWithHTTPClient(client *http.Client) *ShowbackSummariesGroupedByLabelListParams {
	return &ShowbackSummariesGroupedByLabelListParams{
		HTTPClient: client,
	}
}

/*
ShowbackSummariesGroupedByLabelListParams contains all the parameters to send to the API endpoint

	for the showback summaries grouped by label list operation.

	Typically these are written to a http.Request.
*/
type ShowbackSummariesGroupedByLabelListParams struct {

	// FromDate.
	//
	// Format: date-time
	FromDate *strfmt.DateTime

	// IsDeleted.
	IsDeleted *bool

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// PeriodDuration.
	PeriodDuration *string

	// ToDate.
	//
	// Format: date-time
	ToDate *strfmt.DateTime

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback summaries grouped by label list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesGroupedByLabelListParams) WithDefaults() *ShowbackSummariesGroupedByLabelListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback summaries grouped by label list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesGroupedByLabelListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithTimeout(timeout time.Duration) *ShowbackSummariesGroupedByLabelListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithContext(ctx context.Context) *ShowbackSummariesGroupedByLabelListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithHTTPClient(client *http.Client) *ShowbackSummariesGroupedByLabelListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFromDate adds the fromDate to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithFromDate(fromDate *strfmt.DateTime) *ShowbackSummariesGroupedByLabelListParams {
	o.SetFromDate(fromDate)
	return o
}

// SetFromDate adds the fromDate to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetFromDate(fromDate *strfmt.DateTime) {
	o.FromDate = fromDate
}

// WithIsDeleted adds the isDeleted to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithIsDeleted(isDeleted *bool) *ShowbackSummariesGroupedByLabelListParams {
	o.SetIsDeleted(isDeleted)
	return o
}

// SetIsDeleted adds the isDeleted to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetIsDeleted(isDeleted *bool) {
	o.IsDeleted = isDeleted
}

// WithOrganizationID adds the organizationID to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithOrganizationID(organizationID *int32) *ShowbackSummariesGroupedByLabelListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithPeriodDuration adds the periodDuration to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithPeriodDuration(periodDuration *string) *ShowbackSummariesGroupedByLabelListParams {
	o.SetPeriodDuration(periodDuration)
	return o
}

// SetPeriodDuration adds the periodDuration to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetPeriodDuration(periodDuration *string) {
	o.PeriodDuration = periodDuration
}

// WithToDate adds the toDate to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithToDate(toDate *strfmt.DateTime) *ShowbackSummariesGroupedByLabelListParams {
	o.SetToDate(toDate)
	return o
}

// SetToDate adds the toDate to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetToDate(toDate *strfmt.DateTime) {
	o.ToDate = toDate
}

// WithV adds the v to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) WithV(v string) *ShowbackSummariesGroupedByLabelListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback summaries grouped by label list params
func (o *ShowbackSummariesGroupedByLabelListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackSummariesGroupedByLabelListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FromDate != nil {

		// query param fromDate
		var qrFromDate strfmt.DateTime

		if o.FromDate != nil {
			qrFromDate = *o.FromDate
		}
		qFromDate := qrFromDate.String()
		if qFromDate != "" {

			if err := r.SetQueryParam("fromDate", qFromDate); err != nil {
				return err
			}
		}
	}

	if o.IsDeleted != nil {

		// query param isDeleted
		var qrIsDeleted bool

		if o.IsDeleted != nil {
			qrIsDeleted = *o.IsDeleted
		}
		qIsDeleted := swag.FormatBool(qrIsDeleted)
		if qIsDeleted != "" {

			if err := r.SetQueryParam("isDeleted", qIsDeleted); err != nil {
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

	if o.PeriodDuration != nil {

		// query param periodDuration
		var qrPeriodDuration string

		if o.PeriodDuration != nil {
			qrPeriodDuration = *o.PeriodDuration
		}
		qPeriodDuration := qrPeriodDuration
		if qPeriodDuration != "" {

			if err := r.SetQueryParam("periodDuration", qPeriodDuration); err != nil {
				return err
			}
		}
	}

	if o.ToDate != nil {

		// query param toDate
		var qrToDate strfmt.DateTime

		if o.ToDate != nil {
			qrToDate = *o.ToDate
		}
		qToDate := qrToDate.String()
		if qToDate != "" {

			if err := r.SetQueryParam("toDate", qToDate); err != nil {
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
