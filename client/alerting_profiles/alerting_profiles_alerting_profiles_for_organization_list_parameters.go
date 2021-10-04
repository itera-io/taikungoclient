// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

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

// NewAlertingProfilesAlertingProfilesForOrganizationListParams creates a new AlertingProfilesAlertingProfilesForOrganizationListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAlertingProfilesAlertingProfilesForOrganizationListParams() *AlertingProfilesAlertingProfilesForOrganizationListParams {
	return &AlertingProfilesAlertingProfilesForOrganizationListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithTimeout creates a new AlertingProfilesAlertingProfilesForOrganizationListParams object
// with the ability to set a timeout on a request.
func NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithTimeout(timeout time.Duration) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	return &AlertingProfilesAlertingProfilesForOrganizationListParams{
		timeout: timeout,
	}
}

// NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithContext creates a new AlertingProfilesAlertingProfilesForOrganizationListParams object
// with the ability to set a context for a request.
func NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithContext(ctx context.Context) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	return &AlertingProfilesAlertingProfilesForOrganizationListParams{
		Context: ctx,
	}
}

// NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithHTTPClient creates a new AlertingProfilesAlertingProfilesForOrganizationListParams object
// with the ability to set a custom HTTPClient for a request.
func NewAlertingProfilesAlertingProfilesForOrganizationListParamsWithHTTPClient(client *http.Client) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	return &AlertingProfilesAlertingProfilesForOrganizationListParams{
		HTTPClient: client,
	}
}

/* AlertingProfilesAlertingProfilesForOrganizationListParams contains all the parameters to send to the API endpoint
   for the alerting profiles alerting profiles for organization list operation.

   Typically these are written to a http.Request.
*/
type AlertingProfilesAlertingProfilesForOrganizationListParams struct {

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

// WithDefaults hydrates default values in the alerting profiles alerting profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithDefaults() *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the alerting profiles alerting profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithTimeout(timeout time.Duration) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithContext(ctx context.Context) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithHTTPClient(client *http.Client) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithOrganizationID(organizationID *int32) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WithV(v string) *AlertingProfilesAlertingProfilesForOrganizationListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the alerting profiles alerting profiles for organization list params
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AlertingProfilesAlertingProfilesForOrganizationListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
