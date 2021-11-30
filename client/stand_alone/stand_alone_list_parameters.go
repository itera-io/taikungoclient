// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

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

// NewStandAloneListParams creates a new StandAloneListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneListParams() *StandAloneListParams {
	return &StandAloneListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneListParamsWithTimeout creates a new StandAloneListParams object
// with the ability to set a timeout on a request.
func NewStandAloneListParamsWithTimeout(timeout time.Duration) *StandAloneListParams {
	return &StandAloneListParams{
		timeout: timeout,
	}
}

// NewStandAloneListParamsWithContext creates a new StandAloneListParams object
// with the ability to set a context for a request.
func NewStandAloneListParamsWithContext(ctx context.Context) *StandAloneListParams {
	return &StandAloneListParams{
		Context: ctx,
	}
}

// NewStandAloneListParamsWithHTTPClient creates a new StandAloneListParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneListParamsWithHTTPClient(client *http.Client) *StandAloneListParams {
	return &StandAloneListParams{
		HTTPClient: client,
	}
}

/* StandAloneListParams contains all the parameters to send to the API endpoint
   for the stand alone list operation.

   Typically these are written to a http.Request.
*/
type StandAloneListParams struct {

	// EndCPU.
	//
	// Format: int32
	EndCPU *int32

	// EndDiskSize.
	//
	// Format: int64
	EndDiskSize *int64

	// EndRAM.
	//
	// Format: int64
	EndRAM *int64

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

	   Page number

	   Format: int32
	*/
	Offset *int32

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// ProjectID.
	//
	// Format: int32
	ProjectID *int32

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

	// StartDiskSize.
	//
	// Format: int64
	StartDiskSize *int64

	// StartRAM.
	//
	// Format: int64
	StartRAM *int64

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneListParams) WithDefaults() *StandAloneListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone list params
func (o *StandAloneListParams) WithTimeout(timeout time.Duration) *StandAloneListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone list params
func (o *StandAloneListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone list params
func (o *StandAloneListParams) WithContext(ctx context.Context) *StandAloneListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone list params
func (o *StandAloneListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone list params
func (o *StandAloneListParams) WithHTTPClient(client *http.Client) *StandAloneListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone list params
func (o *StandAloneListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndCPU adds the endCPU to the stand alone list params
func (o *StandAloneListParams) WithEndCPU(endCPU *int32) *StandAloneListParams {
	o.SetEndCPU(endCPU)
	return o
}

// SetEndCPU adds the endCpu to the stand alone list params
func (o *StandAloneListParams) SetEndCPU(endCPU *int32) {
	o.EndCPU = endCPU
}

// WithEndDiskSize adds the endDiskSize to the stand alone list params
func (o *StandAloneListParams) WithEndDiskSize(endDiskSize *int64) *StandAloneListParams {
	o.SetEndDiskSize(endDiskSize)
	return o
}

// SetEndDiskSize adds the endDiskSize to the stand alone list params
func (o *StandAloneListParams) SetEndDiskSize(endDiskSize *int64) {
	o.EndDiskSize = endDiskSize
}

// WithEndRAM adds the endRAM to the stand alone list params
func (o *StandAloneListParams) WithEndRAM(endRAM *int64) *StandAloneListParams {
	o.SetEndRAM(endRAM)
	return o
}

// SetEndRAM adds the endRam to the stand alone list params
func (o *StandAloneListParams) SetEndRAM(endRAM *int64) {
	o.EndRAM = endRAM
}

// WithID adds the id to the stand alone list params
func (o *StandAloneListParams) WithID(id *int32) *StandAloneListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stand alone list params
func (o *StandAloneListParams) SetID(id *int32) {
	o.ID = id
}

// WithLimit adds the limit to the stand alone list params
func (o *StandAloneListParams) WithLimit(limit *int32) *StandAloneListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the stand alone list params
func (o *StandAloneListParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the stand alone list params
func (o *StandAloneListParams) WithOffset(offset *int32) *StandAloneListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the stand alone list params
func (o *StandAloneListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithOrganizationID adds the organizationID to the stand alone list params
func (o *StandAloneListParams) WithOrganizationID(organizationID *int32) *StandAloneListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the stand alone list params
func (o *StandAloneListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the stand alone list params
func (o *StandAloneListParams) WithProjectID(projectID *int32) *StandAloneListParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the stand alone list params
func (o *StandAloneListParams) SetProjectID(projectID *int32) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the stand alone list params
func (o *StandAloneListParams) WithSearch(search *string) *StandAloneListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the stand alone list params
func (o *StandAloneListParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the stand alone list params
func (o *StandAloneListParams) WithSortBy(sortBy *string) *StandAloneListParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the stand alone list params
func (o *StandAloneListParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the stand alone list params
func (o *StandAloneListParams) WithSortDirection(sortDirection *string) *StandAloneListParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the stand alone list params
func (o *StandAloneListParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartCPU adds the startCPU to the stand alone list params
func (o *StandAloneListParams) WithStartCPU(startCPU *int32) *StandAloneListParams {
	o.SetStartCPU(startCPU)
	return o
}

// SetStartCPU adds the startCpu to the stand alone list params
func (o *StandAloneListParams) SetStartCPU(startCPU *int32) {
	o.StartCPU = startCPU
}

// WithStartDiskSize adds the startDiskSize to the stand alone list params
func (o *StandAloneListParams) WithStartDiskSize(startDiskSize *int64) *StandAloneListParams {
	o.SetStartDiskSize(startDiskSize)
	return o
}

// SetStartDiskSize adds the startDiskSize to the stand alone list params
func (o *StandAloneListParams) SetStartDiskSize(startDiskSize *int64) {
	o.StartDiskSize = startDiskSize
}

// WithStartRAM adds the startRAM to the stand alone list params
func (o *StandAloneListParams) WithStartRAM(startRAM *int64) *StandAloneListParams {
	o.SetStartRAM(startRAM)
	return o
}

// SetStartRAM adds the startRam to the stand alone list params
func (o *StandAloneListParams) SetStartRAM(startRAM *int64) {
	o.StartRAM = startRAM
}

// WithV adds the v to the stand alone list params
func (o *StandAloneListParams) WithV(v string) *StandAloneListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone list params
func (o *StandAloneListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.EndDiskSize != nil {

		// query param endDiskSize
		var qrEndDiskSize int64

		if o.EndDiskSize != nil {
			qrEndDiskSize = *o.EndDiskSize
		}
		qEndDiskSize := swag.FormatInt64(qrEndDiskSize)
		if qEndDiskSize != "" {

			if err := r.SetQueryParam("endDiskSize", qEndDiskSize); err != nil {
				return err
			}
		}
	}

	if o.EndRAM != nil {

		// query param endRam
		var qrEndRAM int64

		if o.EndRAM != nil {
			qrEndRAM = *o.EndRAM
		}
		qEndRAM := swag.FormatInt64(qrEndRAM)
		if qEndRAM != "" {

			if err := r.SetQueryParam("endRam", qEndRAM); err != nil {
				return err
			}
		}
	}

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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID int32

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := swag.FormatInt32(qrProjectID)
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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

	if o.StartDiskSize != nil {

		// query param startDiskSize
		var qrStartDiskSize int64

		if o.StartDiskSize != nil {
			qrStartDiskSize = *o.StartDiskSize
		}
		qStartDiskSize := swag.FormatInt64(qrStartDiskSize)
		if qStartDiskSize != "" {

			if err := r.SetQueryParam("startDiskSize", qStartDiskSize); err != nil {
				return err
			}
		}
	}

	if o.StartRAM != nil {

		// query param startRam
		var qrStartRAM int64

		if o.StartRAM != nil {
			qrStartRAM = *o.StartRAM
		}
		qStartRAM := swag.FormatInt64(qrStartRAM)
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