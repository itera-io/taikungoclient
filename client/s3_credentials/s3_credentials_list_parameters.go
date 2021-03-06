// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

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

// NewS3CredentialsListParams creates a new S3CredentialsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3CredentialsListParams() *S3CredentialsListParams {
	return &S3CredentialsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3CredentialsListParamsWithTimeout creates a new S3CredentialsListParams object
// with the ability to set a timeout on a request.
func NewS3CredentialsListParamsWithTimeout(timeout time.Duration) *S3CredentialsListParams {
	return &S3CredentialsListParams{
		timeout: timeout,
	}
}

// NewS3CredentialsListParamsWithContext creates a new S3CredentialsListParams object
// with the ability to set a context for a request.
func NewS3CredentialsListParamsWithContext(ctx context.Context) *S3CredentialsListParams {
	return &S3CredentialsListParams{
		Context: ctx,
	}
}

// NewS3CredentialsListParamsWithHTTPClient creates a new S3CredentialsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3CredentialsListParamsWithHTTPClient(client *http.Client) *S3CredentialsListParams {
	return &S3CredentialsListParams{
		HTTPClient: client,
	}
}

/* S3CredentialsListParams contains all the parameters to send to the API endpoint
   for the s3 credentials list operation.

   Typically these are written to a http.Request.
*/
type S3CredentialsListParams struct {

	// ID.
	//
	// Format: int32
	ID *int32

	/* Limit.

	   Limits user size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Skip elements

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

// WithDefaults hydrates default values in the s3 credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsListParams) WithDefaults() *S3CredentialsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 credentials list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 credentials list params
func (o *S3CredentialsListParams) WithTimeout(timeout time.Duration) *S3CredentialsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 credentials list params
func (o *S3CredentialsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 credentials list params
func (o *S3CredentialsListParams) WithContext(ctx context.Context) *S3CredentialsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 credentials list params
func (o *S3CredentialsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 credentials list params
func (o *S3CredentialsListParams) WithHTTPClient(client *http.Client) *S3CredentialsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 credentials list params
func (o *S3CredentialsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the s3 credentials list params
func (o *S3CredentialsListParams) WithID(id *int32) *S3CredentialsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the s3 credentials list params
func (o *S3CredentialsListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the s3 credentials list params
func (o *S3CredentialsListParams) WithLimit(limit *int32) *S3CredentialsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the s3 credentials list params
func (o *S3CredentialsListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the s3 credentials list params
func (o *S3CredentialsListParams) WithOffset(offset *int32) *S3CredentialsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the s3 credentials list params
func (o *S3CredentialsListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the s3 credentials list params
func (o *S3CredentialsListParams) WithOrganizationID(organizationID *int32) *S3CredentialsListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the s3 credentials list params
func (o *S3CredentialsListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the s3 credentials list params
func (o *S3CredentialsListParams) WithSearch(search *string) *S3CredentialsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the s3 credentials list params
func (o *S3CredentialsListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSearchID adds the searchID to the s3 credentials list params
func (o *S3CredentialsListParams) WithSearchID(searchID *string) *S3CredentialsListParams {
	o.SetSearchID(searchID)
	return o
}

// SetSearchID adds the searchId to the s3 credentials list params
func (o *S3CredentialsListParams) SetSearchID(searchID *string) {
	o.SearchID = searchID
}

// WithV adds the v to the s3 credentials list params
func (o *S3CredentialsListParams) WithV(v string) *S3CredentialsListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the s3 credentials list params
func (o *S3CredentialsListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *S3CredentialsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID int32

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt32(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

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
