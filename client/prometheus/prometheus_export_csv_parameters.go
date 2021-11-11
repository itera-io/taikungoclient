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

// NewPrometheusExportCsvParams creates a new PrometheusExportCsvParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPrometheusExportCsvParams() *PrometheusExportCsvParams {
	return &PrometheusExportCsvParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPrometheusExportCsvParamsWithTimeout creates a new PrometheusExportCsvParams object
// with the ability to set a timeout on a request.
func NewPrometheusExportCsvParamsWithTimeout(timeout time.Duration) *PrometheusExportCsvParams {
	return &PrometheusExportCsvParams{
		timeout: timeout,
	}
}

// NewPrometheusExportCsvParamsWithContext creates a new PrometheusExportCsvParams object
// with the ability to set a context for a request.
func NewPrometheusExportCsvParamsWithContext(ctx context.Context) *PrometheusExportCsvParams {
	return &PrometheusExportCsvParams{
		Context: ctx,
	}
}

// NewPrometheusExportCsvParamsWithHTTPClient creates a new PrometheusExportCsvParams object
// with the ability to set a custom HTTPClient for a request.
func NewPrometheusExportCsvParamsWithHTTPClient(client *http.Client) *PrometheusExportCsvParams {
	return &PrometheusExportCsvParams{
		HTTPClient: client,
	}
}

/* PrometheusExportCsvParams contains all the parameters to send to the API endpoint
   for the prometheus export csv operation.

   Typically these are written to a http.Request.
*/
type PrometheusExportCsvParams struct {

	// EndDate.
	//
	// Format: date-time
	EndDate *strfmt.DateTime

	// IsEmailEnabled.
	IsEmailEnabled *bool

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

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

// WithDefaults hydrates default values in the prometheus export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusExportCsvParams) WithDefaults() *PrometheusExportCsvParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the prometheus export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PrometheusExportCsvParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithTimeout(timeout time.Duration) *PrometheusExportCsvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithContext(ctx context.Context) *PrometheusExportCsvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithHTTPClient(client *http.Client) *PrometheusExportCsvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithEndDate(endDate *strfmt.DateTime) *PrometheusExportCsvParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetEndDate(endDate *strfmt.DateTime) {
	o.EndDate = endDate
}

// WithIsEmailEnabled adds the isEmailEnabled to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithIsEmailEnabled(isEmailEnabled *bool) *PrometheusExportCsvParams {
	o.SetIsEmailEnabled(isEmailEnabled)
	return o
}

// SetIsEmailEnabled adds the isEmailEnabled to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetIsEmailEnabled(isEmailEnabled *bool) {
	o.IsEmailEnabled = isEmailEnabled
}

// WithOrganizationID adds the organizationID to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithOrganizationID(organizationID *int32) *PrometheusExportCsvParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithStartDate adds the startDate to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithStartDate(startDate *strfmt.DateTime) *PrometheusExportCsvParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetStartDate(startDate *strfmt.DateTime) {
	o.StartDate = startDate
}

// WithV adds the v to the prometheus export csv params
func (o *PrometheusExportCsvParams) WithV(v string) *PrometheusExportCsvParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the prometheus export csv params
func (o *PrometheusExportCsvParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *PrometheusExportCsvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
