// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_profiles

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

// NewKubernetesProfilesKubernetesProfilesForOrganizationListParams creates a new KubernetesProfilesKubernetesProfilesForOrganizationListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesProfilesKubernetesProfilesForOrganizationListParams() *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	return &KubernetesProfilesKubernetesProfilesForOrganizationListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithTimeout creates a new KubernetesProfilesKubernetesProfilesForOrganizationListParams object
// with the ability to set a timeout on a request.
func NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithTimeout(timeout time.Duration) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	return &KubernetesProfilesKubernetesProfilesForOrganizationListParams{
		timeout: timeout,
	}
}

// NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithContext creates a new KubernetesProfilesKubernetesProfilesForOrganizationListParams object
// with the ability to set a context for a request.
func NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithContext(ctx context.Context) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	return &KubernetesProfilesKubernetesProfilesForOrganizationListParams{
		Context: ctx,
	}
}

// NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithHTTPClient creates a new KubernetesProfilesKubernetesProfilesForOrganizationListParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesProfilesKubernetesProfilesForOrganizationListParamsWithHTTPClient(client *http.Client) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	return &KubernetesProfilesKubernetesProfilesForOrganizationListParams{
		HTTPClient: client,
	}
}

/*
KubernetesProfilesKubernetesProfilesForOrganizationListParams contains all the parameters to send to the API endpoint

	for the kubernetes profiles kubernetes profiles for organization list operation.

	Typically these are written to a http.Request.
*/
type KubernetesProfilesKubernetesProfilesForOrganizationListParams struct {

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// Search.
	Search *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes profiles kubernetes profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithDefaults() *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes profiles kubernetes profiles for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithTimeout(timeout time.Duration) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithContext(ctx context.Context) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithHTTPClient(client *http.Client) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithOrganizationID(organizationID *int32) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithSearch(search *string) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WithV(v string) *KubernetesProfilesKubernetesProfilesForOrganizationListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes profiles kubernetes profiles for organization list params
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesProfilesKubernetesProfilesForOrganizationListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
