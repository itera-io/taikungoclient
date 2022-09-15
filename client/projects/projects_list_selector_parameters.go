// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsListSelectorParams creates a new ProjectsListSelectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsListSelectorParams() *ProjectsListSelectorParams {
	return &ProjectsListSelectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsListSelectorParamsWithTimeout creates a new ProjectsListSelectorParams object
// with the ability to set a timeout on a request.
func NewProjectsListSelectorParamsWithTimeout(timeout time.Duration) *ProjectsListSelectorParams {
	return &ProjectsListSelectorParams{
		timeout: timeout,
	}
}

// NewProjectsListSelectorParamsWithContext creates a new ProjectsListSelectorParams object
// with the ability to set a context for a request.
func NewProjectsListSelectorParamsWithContext(ctx context.Context) *ProjectsListSelectorParams {
	return &ProjectsListSelectorParams{
		Context: ctx,
	}
}

// NewProjectsListSelectorParamsWithHTTPClient creates a new ProjectsListSelectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsListSelectorParamsWithHTTPClient(client *http.Client) *ProjectsListSelectorParams {
	return &ProjectsListSelectorParams{
		HTTPClient: client,
	}
}

/*
ProjectsListSelectorParams contains all the parameters to send to the API endpoint

	for the projects list selector operation.

	Typically these are written to a http.Request.
*/
type ProjectsListSelectorParams struct {

	// CatalogID.
	//
	// Format: int32
	CatalogID *int32

	// Healthy.
	Healthy *bool

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

// WithDefaults hydrates default values in the projects list selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsListSelectorParams) WithDefaults() *ProjectsListSelectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects list selector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsListSelectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects list selector params
func (o *ProjectsListSelectorParams) WithTimeout(timeout time.Duration) *ProjectsListSelectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects list selector params
func (o *ProjectsListSelectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects list selector params
func (o *ProjectsListSelectorParams) WithContext(ctx context.Context) *ProjectsListSelectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects list selector params
func (o *ProjectsListSelectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects list selector params
func (o *ProjectsListSelectorParams) WithHTTPClient(client *http.Client) *ProjectsListSelectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects list selector params
func (o *ProjectsListSelectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCatalogID adds the catalogID to the projects list selector params
func (o *ProjectsListSelectorParams) WithCatalogID(catalogID *int32) *ProjectsListSelectorParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the projects list selector params
func (o *ProjectsListSelectorParams) SetCatalogID(catalogID *int32) {
	o.CatalogID = catalogID
}

// WithHealthy adds the healthy to the projects list selector params
func (o *ProjectsListSelectorParams) WithHealthy(healthy *bool) *ProjectsListSelectorParams {
	o.SetHealthy(healthy)
	return o
}

// SetHealthy adds the healthy to the projects list selector params
func (o *ProjectsListSelectorParams) SetHealthy(healthy *bool) {
	o.Healthy = healthy
}

// WithOrganizationID adds the organizationID to the projects list selector params
func (o *ProjectsListSelectorParams) WithOrganizationID(organizationID *int32) *ProjectsListSelectorParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the projects list selector params
func (o *ProjectsListSelectorParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the projects list selector params
func (o *ProjectsListSelectorParams) WithSearch(search *string) *ProjectsListSelectorParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the projects list selector params
func (o *ProjectsListSelectorParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the projects list selector params
func (o *ProjectsListSelectorParams) WithV(v string) *ProjectsListSelectorParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects list selector params
func (o *ProjectsListSelectorParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsListSelectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CatalogID != nil {

		// query param catalogId
		var qrCatalogID int32

		if o.CatalogID != nil {
			qrCatalogID = *o.CatalogID
		}
		qCatalogID := swag.FormatInt32(qrCatalogID)
		if qCatalogID != "" {

			if err := r.SetQueryParam("catalogId", qCatalogID); err != nil {
				return err
			}
		}
	}

	if o.Healthy != nil {

		// query param healthy
		var qrHealthy bool

		if o.Healthy != nil {
			qrHealthy = *o.Healthy
		}
		qHealthy := swag.FormatBool(qrHealthy)
		if qHealthy != "" {

			if err := r.SetQueryParam("healthy", qHealthy); err != nil {
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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
