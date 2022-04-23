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

// NewFlavorsGoogleFlavorsParams creates a new FlavorsGoogleFlavorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlavorsGoogleFlavorsParams() *FlavorsGoogleFlavorsParams {
	return &FlavorsGoogleFlavorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlavorsGoogleFlavorsParamsWithTimeout creates a new FlavorsGoogleFlavorsParams object
// with the ability to set a timeout on a request.
func NewFlavorsGoogleFlavorsParamsWithTimeout(timeout time.Duration) *FlavorsGoogleFlavorsParams {
	return &FlavorsGoogleFlavorsParams{
		timeout: timeout,
	}
}

// NewFlavorsGoogleFlavorsParamsWithContext creates a new FlavorsGoogleFlavorsParams object
// with the ability to set a context for a request.
func NewFlavorsGoogleFlavorsParamsWithContext(ctx context.Context) *FlavorsGoogleFlavorsParams {
	return &FlavorsGoogleFlavorsParams{
		Context: ctx,
	}
}

// NewFlavorsGoogleFlavorsParamsWithHTTPClient creates a new FlavorsGoogleFlavorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlavorsGoogleFlavorsParamsWithHTTPClient(client *http.Client) *FlavorsGoogleFlavorsParams {
	return &FlavorsGoogleFlavorsParams{
		HTTPClient: client,
	}
}

/* FlavorsGoogleFlavorsParams contains all the parameters to send to the API endpoint
   for the flavors google flavors operation.

   Typically these are written to a http.Request.
*/
type FlavorsGoogleFlavorsParams struct {

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
	// Format: int64
	EndRAM *int64

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
	// Format: int64
	StartRAM *int64

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the flavors google flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsGoogleFlavorsParams) WithDefaults() *FlavorsGoogleFlavorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flavors google flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsGoogleFlavorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithTimeout(timeout time.Duration) *FlavorsGoogleFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithContext(ctx context.Context) *FlavorsGoogleFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithHTTPClient(client *http.Client) *FlavorsGoogleFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithCloudID(cloudID int32) *FlavorsGoogleFlavorsParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithEndCPU adds the endCPU to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithEndCPU(endCPU *int32) *FlavorsGoogleFlavorsParams {
	o.SetEndCPU(endCPU)
	return o
}

// SetEndCPU adds the endCpu to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetEndCPU(endCPU *int32) {
	o.EndCPU = endCPU
}

// WithEndRAM adds the endRAM to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithEndRAM(endRAM *int64) *FlavorsGoogleFlavorsParams {
	o.SetEndRAM(endRAM)
	return o
}

// SetEndRAM adds the endRam to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetEndRAM(endRAM *int64) {
	o.EndRAM = endRAM
}

// WithLimit adds the limit to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithLimit(limit *int32) *FlavorsGoogleFlavorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithOffset(offset *int32) *FlavorsGoogleFlavorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearch adds the search to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithSearch(search *string) *FlavorsGoogleFlavorsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithSortBy(sortBy *string) *FlavorsGoogleFlavorsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithSortDirection(sortDirection *string) *FlavorsGoogleFlavorsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartCPU adds the startCPU to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithStartCPU(startCPU *int32) *FlavorsGoogleFlavorsParams {
	o.SetStartCPU(startCPU)
	return o
}

// SetStartCPU adds the startCpu to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetStartCPU(startCPU *int32) {
	o.StartCPU = startCPU
}

// WithStartRAM adds the startRAM to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithStartRAM(startRAM *int64) *FlavorsGoogleFlavorsParams {
	o.SetStartRAM(startRAM)
	return o
}

// SetStartRAM adds the startRam to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetStartRAM(startRAM *int64) {
	o.StartRAM = startRAM
}

// WithV adds the v to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) WithV(v string) *FlavorsGoogleFlavorsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the flavors google flavors params
func (o *FlavorsGoogleFlavorsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *FlavorsGoogleFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
