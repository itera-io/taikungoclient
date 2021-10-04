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

// NewOpsCredentialsListParams creates a new OpsCredentialsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpsCredentialsListParams() *OpsCredentialsListParams {
	return &OpsCredentialsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpsCredentialsListParamsWithTimeout creates a new OpsCredentialsListParams object
// with the ability to set a timeout on a request.
func NewOpsCredentialsListParamsWithTimeout(timeout time.Duration) *OpsCredentialsListParams {
	return &OpsCredentialsListParams{
		timeout: timeout,
	}
}

// NewOpsCredentialsListParamsWithContext creates a new OpsCredentialsListParams object
// with the ability to set a context for a request.
func NewOpsCredentialsListParamsWithContext(ctx context.Context) *OpsCredentialsListParams {
	return &OpsCredentialsListParams{
		Context: ctx,
	}
}

// NewOpsCredentialsListParamsWithHTTPClient creates a new OpsCredentialsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpsCredentialsListParamsWithHTTPClient(client *http.Client) *OpsCredentialsListParams {
	return &OpsCredentialsListParams{
		HTTPClient: client,
	}
}

/* OpsCredentialsListParams contains all the parameters to send to the API endpoint
   for the ops credentials list operation.

   Typically these are written to a http.Request.
*/
type OpsCredentialsListParams struct {

	/* Limit.

	   Limits user size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Page number

	   Format: int32
	*/
	Offset *int32

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// Search.
	Search *string

	// SearchID.
	SearchID *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ops credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsListParams) WithDefaults() *OpsCredentialsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ops credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpsCredentialsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ops credentials list params
func (o *OpsCredentialsListParams) WithTimeout(timeout time.Duration) *OpsCredentialsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ops credentials list params
func (o *OpsCredentialsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ops credentials list params
func (o *OpsCredentialsListParams) WithContext(ctx context.Context) *OpsCredentialsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ops credentials list params
func (o *OpsCredentialsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ops credentials list params
func (o *OpsCredentialsListParams) WithHTTPClient(client *http.Client) *OpsCredentialsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ops credentials list params
func (o *OpsCredentialsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the ops credentials list params
func (o *OpsCredentialsListParams) WithLimit(limit *int32) *OpsCredentialsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ops credentials list params
func (o *OpsCredentialsListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the ops credentials list params
func (o *OpsCredentialsListParams) WithOffset(offset *int32) *OpsCredentialsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ops credentials list params
func (o *OpsCredentialsListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the ops credentials list params
func (o *OpsCredentialsListParams) WithOrganizationID(organizationID *int32) *OpsCredentialsListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the ops credentials list params
func (o *OpsCredentialsListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the ops credentials list params
func (o *OpsCredentialsListParams) WithSearch(search *string) *OpsCredentialsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the ops credentials list params
func (o *OpsCredentialsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the ops credentials list params
func (o *OpsCredentialsListParams) WithSearchID(searchID *string) *OpsCredentialsListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the ops credentials list params
func (o *OpsCredentialsListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithV adds the v to the ops credentials list params
func (o *OpsCredentialsListParams) WithV(v string) *OpsCredentialsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the ops credentials list params
func (o *OpsCredentialsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OpsCredentialsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.SearchID != nil {

		// query param searchId
		var qrSearchID string

		if o.SearchID != nil {
			qrSearchID = *o.SearchID
		}
		qSearchID := qrSearchID
		if qSearchID != "" {

			if err := r.SetQueryParam("searchId", qSearchID); err != nil {
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
