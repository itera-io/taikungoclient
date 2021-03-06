// Code generated by go-swagger; DO NOT EDIT.

package billing

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

// NewBillingGroupedListParams creates a new BillingGroupedListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingGroupedListParams() *BillingGroupedListParams {
	return &BillingGroupedListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingGroupedListParamsWithTimeout creates a new BillingGroupedListParams object
// with the ability to set a timeout on a request.
func NewBillingGroupedListParamsWithTimeout(timeout time.Duration) *BillingGroupedListParams {
	return &BillingGroupedListParams{
		timeout: timeout,
	}
}

// NewBillingGroupedListParamsWithContext creates a new BillingGroupedListParams object
// with the ability to set a context for a request.
func NewBillingGroupedListParamsWithContext(ctx context.Context) *BillingGroupedListParams {
	return &BillingGroupedListParams{
		Context: ctx,
	}
}

// NewBillingGroupedListParamsWithHTTPClient creates a new BillingGroupedListParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingGroupedListParamsWithHTTPClient(client *http.Client) *BillingGroupedListParams {
	return &BillingGroupedListParams{
		HTTPClient: client,
	}
}

/* BillingGroupedListParams contains all the parameters to send to the API endpoint
   for the billing grouped list operation.

   Typically these are written to a http.Request.
*/
type BillingGroupedListParams struct {

	// IsDeleted.
	IsDeleted *bool

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// PeriodDuration.
	PeriodDuration *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing grouped list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingGroupedListParams) WithDefaults() *BillingGroupedListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing grouped list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingGroupedListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing grouped list params
func (o *BillingGroupedListParams) WithTimeout(timeout time.Duration) *BillingGroupedListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing grouped list params
func (o *BillingGroupedListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing grouped list params
func (o *BillingGroupedListParams) WithContext(ctx context.Context) *BillingGroupedListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing grouped list params
func (o *BillingGroupedListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing grouped list params
func (o *BillingGroupedListParams) WithHTTPClient(client *http.Client) *BillingGroupedListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing grouped list params
func (o *BillingGroupedListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsDeleted adds the isDeleted to the billing grouped list params
func (o *BillingGroupedListParams) WithIsDeleted(isDeleted *bool) *BillingGroupedListParams {
	o.SetIsDeleted(isDeleted)
	return o
}

// SetIsDeleted adds the isDeleted to the billing grouped list params
func (o *BillingGroupedListParams) SetIsDeleted(isDeleted *bool) {
	o.IsDeleted = isDeleted
}

// WithOrganizationID adds the organizationID to the billing grouped list params
func (o *BillingGroupedListParams) WithOrganizationID(organizationID *int32) *BillingGroupedListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the billing grouped list params
func (o *BillingGroupedListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithPeriodDuration adds the periodDuration to the billing grouped list params
func (o *BillingGroupedListParams) WithPeriodDuration(periodDuration *string) *BillingGroupedListParams {
	o.SetPeriodDuration(periodDuration)
	return o
}

// SetPeriodDuration adds the periodDuration to the billing grouped list params
func (o *BillingGroupedListParams) SetPeriodDuration(periodDuration *string) {
	o.PeriodDuration = periodDuration
}

// WithV adds the v to the billing grouped list params
func (o *BillingGroupedListParams) WithV(v string) *BillingGroupedListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the billing grouped list params
func (o *BillingGroupedListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BillingGroupedListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
