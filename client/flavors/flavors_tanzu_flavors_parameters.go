// Code generated by go-swagger; DO NOT EDIT.

package flavors

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

// NewFlavorsTanzuFlavorsParams creates a new FlavorsTanzuFlavorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlavorsTanzuFlavorsParams() *FlavorsTanzuFlavorsParams {
	return &FlavorsTanzuFlavorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlavorsTanzuFlavorsParamsWithTimeout creates a new FlavorsTanzuFlavorsParams object
// with the ability to set a timeout on a request.
func NewFlavorsTanzuFlavorsParamsWithTimeout(timeout time.Duration) *FlavorsTanzuFlavorsParams {
	return &FlavorsTanzuFlavorsParams{
		timeout: timeout,
	}
}

// NewFlavorsTanzuFlavorsParamsWithContext creates a new FlavorsTanzuFlavorsParams object
// with the ability to set a context for a request.
func NewFlavorsTanzuFlavorsParamsWithContext(ctx context.Context) *FlavorsTanzuFlavorsParams {
	return &FlavorsTanzuFlavorsParams{
		Context: ctx,
	}
}

// NewFlavorsTanzuFlavorsParamsWithHTTPClient creates a new FlavorsTanzuFlavorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlavorsTanzuFlavorsParamsWithHTTPClient(client *http.Client) *FlavorsTanzuFlavorsParams {
	return &FlavorsTanzuFlavorsParams{
		HTTPClient: client,
	}
}

/*
FlavorsTanzuFlavorsParams contains all the parameters to send to the API endpoint

	for the flavors tanzu flavors operation.

	Typically these are written to a http.Request.
*/
type FlavorsTanzuFlavorsParams struct {

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
	// Format: int32
	EndRAM *int32

	/* Limit.

	   Limits size (by default 50)

	   Format: int32
	*/
	Limit *int32

	/* Offset.

	   Skip elements

	   Format: int32
	*/
	Offset *int32

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
	// Format: int32
	StartRAM *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the flavors tanzu flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsTanzuFlavorsParams) WithDefaults() *FlavorsTanzuFlavorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flavors tanzu flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsTanzuFlavorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithTimeout(timeout time.Duration) *FlavorsTanzuFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithContext(ctx context.Context) *FlavorsTanzuFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithHTTPClient(client *http.Client) *FlavorsTanzuFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithCloudID(cloudID int32) *FlavorsTanzuFlavorsParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithEndCPU adds the endCPU to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithEndCPU(endCPU *int32) *FlavorsTanzuFlavorsParams {
	o.SetEndCPU(endCPU)
	return o
}

// SetEndCPU adds the endCpu to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetEndCPU(endCPU *int32) {
	o.EndCPU = endCPU
}

// WithEndRAM adds the endRAM to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithEndRAM(endRAM *int32) *FlavorsTanzuFlavorsParams {
	o.SetEndRAM(endRAM)
	return o
}

// SetEndRAM adds the endRam to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetEndRAM(endRAM *int32) {
	o.EndRAM = endRAM
}

// WithLimit adds the limit to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithLimit(limit *int32) *FlavorsTanzuFlavorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithOffset(offset *int32) *FlavorsTanzuFlavorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearch adds the search to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithSearch(search *string) *FlavorsTanzuFlavorsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithSortBy(sortBy *string) *FlavorsTanzuFlavorsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithSortDirection(sortDirection *string) *FlavorsTanzuFlavorsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartCPU adds the startCPU to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithStartCPU(startCPU *int32) *FlavorsTanzuFlavorsParams {
	o.SetStartCPU(startCPU)
	return o
}

// SetStartCPU adds the startCpu to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetStartCPU(startCPU *int32) {
	o.StartCPU = startCPU
}

// WithStartRAM adds the startRAM to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithStartRAM(startRAM *int32) *FlavorsTanzuFlavorsParams {
	o.SetStartRAM(startRAM)
	return o
}

// SetStartRAM adds the startRam to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetStartRAM(startRAM *int32) {
	o.StartRAM = startRAM
}

// WithV adds the v to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) WithV(v string) *FlavorsTanzuFlavorsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the flavors tanzu flavors params
func (o *FlavorsTanzuFlavorsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *FlavorsTanzuFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrEndRAM int32

		if o.EndRAM != nil {
			qrEndRAM = *o.EndRAM
		}
		qEndRAM := swag.FormatInt32(qrEndRAM)
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
		var qrStartRAM int32

		if o.StartRAM != nil {
			qrStartRAM = *o.StartRAM
		}
		qStartRAM := swag.FormatInt32(qrStartRAM)
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