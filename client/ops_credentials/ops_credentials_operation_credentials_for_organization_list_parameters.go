// Code generated by go-swagger; DO NOT EDIT.

package ops_credentials

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

// NewOpsCredentialsOperationCredentialsForOrganizationListParams creates a new OpsCredentialsOperationCredentialsForOrganizationListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpsCredentialsOperationCredentialsForOrganizationListParams() *OpsCredentialsOperationCredentialsForOrganizationListParams {
	return &OpsCredentialsOperationCredentialsForOrganizationListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithTimeout creates a new OpsCredentialsOperationCredentialsForOrganizationListParams object
// with the ability to set a timeout on a request.
func NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithTimeout(timeout time.Duration) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	return &OpsCredentialsOperationCredentialsForOrganizationListParams{
		timeout: timeout,
	}
}

// NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithContext creates a new OpsCredentialsOperationCredentialsForOrganizationListParams object
// with the ability to set a context for a request.
func NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithContext(ctx context.Context) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	return &OpsCredentialsOperationCredentialsForOrganizationListParams{
		Context: ctx,
	}
}

// NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithHTTPClient creates a new OpsCredentialsOperationCredentialsForOrganizationListParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpsCredentialsOperationCredentialsForOrganizationListParamsWithHTTPClient(client *http.Client) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	return &OpsCredentialsOperationCredentialsForOrganizationListParams{
		HTTPClient: client,
	}
}

/* OpsCredentialsOperationCredentialsForOrganizationListParams contains all the parameters to send to the API endpoint
   for the ops credentials operation credentials for organization list operation.

   Typically these are written to a http.Request.
*/
type OpsCredentialsOperationCredentialsForOrganizationListParams struct {

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

// WithDefaults hydrates default values in the ops credentials operation credentials for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithDefaults() *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ops credentials operation credentials for organization list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithTimeout(timeout time.Duration) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithContext(ctx context.Context) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithHTTPClient(client *http.Client) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithOrganizationID(organizationID *int32) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithSearch(search *string) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WithV(v string) *OpsCredentialsOperationCredentialsForOrganizationListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ops credentials operation credentials for organization list params
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpsCredentialsOperationCredentialsForOrganizationListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
