// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

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

// NewCloudCredentialsAllFlavorsParams creates a new CloudCredentialsAllFlavorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudCredentialsAllFlavorsParams() *CloudCredentialsAllFlavorsParams {
	return &CloudCredentialsAllFlavorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudCredentialsAllFlavorsParamsWithTimeout creates a new CloudCredentialsAllFlavorsParams object
// with the ability to set a timeout on a request.
func NewCloudCredentialsAllFlavorsParamsWithTimeout(timeout time.Duration) *CloudCredentialsAllFlavorsParams {
	return &CloudCredentialsAllFlavorsParams{
		timeout: timeout,
	}
}

// NewCloudCredentialsAllFlavorsParamsWithContext creates a new CloudCredentialsAllFlavorsParams object
// with the ability to set a context for a request.
func NewCloudCredentialsAllFlavorsParamsWithContext(ctx context.Context) *CloudCredentialsAllFlavorsParams {
	return &CloudCredentialsAllFlavorsParams{
		Context: ctx,
	}
}

// NewCloudCredentialsAllFlavorsParamsWithHTTPClient creates a new CloudCredentialsAllFlavorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudCredentialsAllFlavorsParamsWithHTTPClient(client *http.Client) *CloudCredentialsAllFlavorsParams {
	return &CloudCredentialsAllFlavorsParams{
		HTTPClient: client,
	}
}

/* CloudCredentialsAllFlavorsParams contains all the parameters to send to the API endpoint
   for the cloud credentials all flavors operation.

   Typically these are written to a http.Request.
*/
type CloudCredentialsAllFlavorsParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	// EndCPU.
	//
	// Format: int32
	EndCPU *int32

	// EndRAM.
	//
	// Format: double
	EndRAM *float64

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

	// SortBy.
	SortBy *string

	// SortDirection.
	SortDirection *string

	// StartCPU.
	//
	// Format: int32
	StartCPU *int32

	// StartRAM.
	//
	// Format: double
	StartRAM *float64

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud credentials all flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudCredentialsAllFlavorsParams) WithDefaults() *CloudCredentialsAllFlavorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud credentials all flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudCredentialsAllFlavorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithTimeout(timeout time.Duration) *CloudCredentialsAllFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithContext(ctx context.Context) *CloudCredentialsAllFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithHTTPClient(client *http.Client) *CloudCredentialsAllFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithCloudID(cloudID int32) *CloudCredentialsAllFlavorsParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithEndCPU adds the endCPU to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithEndCPU(endCPU *int32) *CloudCredentialsAllFlavorsParams {
	o.SetEndCPU(endCPU)
	return o
}

// SetEndCPU adds the endCpu to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetEndCPU(endCPU *int32) {
	o.EndCPU = endCPU
}

// WithEndRAM adds the endRAM to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithEndRAM(endRAM *float64) *CloudCredentialsAllFlavorsParams {
	o.SetEndRAM(endRAM)
	return o
}

// SetEndRAM adds the endRam to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetEndRAM(endRAM *float64) {
	o.EndRAM = endRAM
}

// WithLimit adds the limit to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithLimit(limit *int32) *CloudCredentialsAllFlavorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithOffset(offset *int32) *CloudCredentialsAllFlavorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithOrganizationID(organizationID *int32) *CloudCredentialsAllFlavorsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithSearch adds the search to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithSearch(search *string) *CloudCredentialsAllFlavorsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithSortBy(sortBy *string) *CloudCredentialsAllFlavorsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithSortDirection(sortDirection *string) *CloudCredentialsAllFlavorsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartCPU adds the startCPU to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithStartCPU(startCPU *int32) *CloudCredentialsAllFlavorsParams {
	o.SetStartCPU(startCPU)
	return o
}

// SetStartCPU adds the startCpu to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetStartCPU(startCPU *int32) {
	o.StartCPU = startCPU
}

// WithStartRAM adds the startRAM to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithStartRAM(startRAM *float64) *CloudCredentialsAllFlavorsParams {
	o.SetStartRAM(startRAM)
	return o
}

// SetStartRAM adds the startRam to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetStartRAM(startRAM *float64) {
	o.StartRAM = startRAM
}

// WithV adds the v to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) WithV(v string) *CloudCredentialsAllFlavorsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cloud credentials all flavors params
func (o *CloudCredentialsAllFlavorsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CloudCredentialsAllFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudId
	if err := r.SetPathParam("cloudId", swag.FormatInt32(o.CloudID)); err != nil {
		return err
	}

	if o.EndCPU != nil {

		// query param endCpu
		var qrEndCPU int32

		if o.EndCPU != nil {
			qrEndCPU = *o.EndCPU
		}
		qEndCPU := swag.FormatInt32(qrEndCPU)
		if qEndCPU != "" {

			if err := r.SetQueryParam("endCpu", qEndCPU); err != nil {
				return err
			}
		}
	}

	if o.EndRAM != nil {

		// query param endRam
		var qrEndRAM float64

		if o.EndRAM != nil {
			qrEndRAM = *o.EndRAM
		}
		qEndRAM := swag.FormatFloat64(qrEndRAM)
		if qEndRAM != "" {

			if err := r.SetQueryParam("endRam", qEndRAM); err != nil {
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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string

		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {

			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}
	}

	if o.SortDirection != nil {

		// query param sortDirection
		var qrSortDirection string

		if o.SortDirection != nil {
			qrSortDirection = *o.SortDirection
		}
		qSortDirection := qrSortDirection
		if qSortDirection != "" {

			if err := r.SetQueryParam("sortDirection", qSortDirection); err != nil {
				return err
			}
		}
	}

	if o.StartCPU != nil {

		// query param startCpu
		var qrStartCPU int32

		if o.StartCPU != nil {
			qrStartCPU = *o.StartCPU
		}
		qStartCPU := swag.FormatInt32(qrStartCPU)
		if qStartCPU != "" {

			if err := r.SetQueryParam("startCpu", qStartCPU); err != nil {
				return err
			}
		}
	}

	if o.StartRAM != nil {

		// query param startRam
		var qrStartRAM float64

		if o.StartRAM != nil {
			qrStartRAM = *o.StartRAM
		}
		qStartRAM := swag.FormatFloat64(qrStartRAM)
		if qStartRAM != "" {

			if err := r.SetQueryParam("startRam", qStartRAM); err != nil {
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
