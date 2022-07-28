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

// NewShowbackSummariesExportCsvParams creates a new ShowbackSummariesExportCsvParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackSummariesExportCsvParams() *ShowbackSummariesExportCsvParams {
	return &ShowbackSummariesExportCsvParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackSummariesExportCsvParamsWithTimeout creates a new ShowbackSummariesExportCsvParams object
// with the ability to set a timeout on a request.
func NewShowbackSummariesExportCsvParamsWithTimeout(timeout time.Duration) *ShowbackSummariesExportCsvParams {
	return &ShowbackSummariesExportCsvParams{
		timeout: timeout,
	}
}

// NewShowbackSummariesExportCsvParamsWithContext creates a new ShowbackSummariesExportCsvParams object
// with the ability to set a context for a request.
func NewShowbackSummariesExportCsvParamsWithContext(ctx context.Context) *ShowbackSummariesExportCsvParams {
	return &ShowbackSummariesExportCsvParams{
		Context: ctx,
	}
}

// NewShowbackSummariesExportCsvParamsWithHTTPClient creates a new ShowbackSummariesExportCsvParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackSummariesExportCsvParamsWithHTTPClient(client *http.Client) *ShowbackSummariesExportCsvParams {
	return &ShowbackSummariesExportCsvParams{
		HTTPClient: client,
	}
}

/* ShowbackSummariesExportCsvParams contains all the parameters to send to the API endpoint
   for the showback summaries export csv operation.

   Typically these are written to a http.Request.
*/
type ShowbackSummariesExportCsvParams struct {

	// IsEmailEnabled.
	IsEmailEnabled *bool

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback summaries export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesExportCsvParams) WithDefaults() *ShowbackSummariesExportCsvParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback summaries export csv params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesExportCsvParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithTimeout(timeout time.Duration) *ShowbackSummariesExportCsvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithContext(ctx context.Context) *ShowbackSummariesExportCsvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithHTTPClient(client *http.Client) *ShowbackSummariesExportCsvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsEmailEnabled adds the isEmailEnabled to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithIsEmailEnabled(isEmailEnabled *bool) *ShowbackSummariesExportCsvParams {
	o.SetIsEmailEnabled(isEmailEnabled)
	return o
}

// SetIsEmailEnabled adds the isEmailEnabled to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetIsEmailEnabled(isEmailEnabled *bool) {
	o.IsEmailEnabled = isEmailEnabled
}

// WithOrganizationID adds the organizationID to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithOrganizationID(organizationID *int32) *ShowbackSummariesExportCsvParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) WithV(v string) *ShowbackSummariesExportCsvParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback summaries export csv params
func (o *ShowbackSummariesExportCsvParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackSummariesExportCsvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
