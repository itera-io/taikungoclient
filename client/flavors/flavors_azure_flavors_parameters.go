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

// NewFlavorsAzureFlavorsParams creates a new FlavorsAzureFlavorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlavorsAzureFlavorsParams() *FlavorsAzureFlavorsParams {
	return &FlavorsAzureFlavorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlavorsAzureFlavorsParamsWithTimeout creates a new FlavorsAzureFlavorsParams object
// with the ability to set a timeout on a request.
func NewFlavorsAzureFlavorsParamsWithTimeout(timeout time.Duration) *FlavorsAzureFlavorsParams {
	return &FlavorsAzureFlavorsParams{
		timeout: timeout,
	}
}

// NewFlavorsAzureFlavorsParamsWithContext creates a new FlavorsAzureFlavorsParams object
// with the ability to set a context for a request.
func NewFlavorsAzureFlavorsParamsWithContext(ctx context.Context) *FlavorsAzureFlavorsParams {
	return &FlavorsAzureFlavorsParams{
		Context: ctx,
	}
}

// NewFlavorsAzureFlavorsParamsWithHTTPClient creates a new FlavorsAzureFlavorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlavorsAzureFlavorsParamsWithHTTPClient(client *http.Client) *FlavorsAzureFlavorsParams {
	return &FlavorsAzureFlavorsParams{
		HTTPClient: client,
	}
}

/* FlavorsAzureFlavorsParams contains all the parameters to send to the API endpoint
   for the flavors azure flavors operation.

   Typically these are written to a http.Request.
*/
type FlavorsAzureFlavorsParams struct {

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

// WithDefaults hydrates default values in the flavors azure flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsAzureFlavorsParams) WithDefaults() *FlavorsAzureFlavorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flavors azure flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlavorsAzureFlavorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithTimeout(timeout time.Duration) *FlavorsAzureFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithContext(ctx context.Context) *FlavorsAzureFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithHTTPClient(client *http.Client) *FlavorsAzureFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithCloudID(cloudID int32) *FlavorsAzureFlavorsParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithEndCPU adds the endCPU to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithEndCPU(endCPU *int32) *FlavorsAzureFlavorsParams {
	o.SetEndCPU(endCPU)
	return o
}

// SetEndCPU adds the endCpu to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetEndCPU(endCPU *int32) {
	o.EndCPU = endCPU
}

// WithEndRAM adds the endRAM to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithEndRAM(endRAM *int32) *FlavorsAzureFlavorsParams {
	o.SetEndRAM(endRAM)
	return o
}

// SetEndRAM adds the endRam to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetEndRAM(endRAM *int32) {
	o.EndRAM = endRAM
}

// WithLimit adds the limit to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithLimit(limit *int32) *FlavorsAzureFlavorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithOffset(offset *int32) *FlavorsAzureFlavorsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSearch adds the search to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithSearch(search *string) *FlavorsAzureFlavorsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetSearch(search *string) {
	o.Search = search
}

// WithSortBy adds the sortBy to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithSortBy(sortBy *string) *FlavorsAzureFlavorsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortDirection adds the sortDirection to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithSortDirection(sortDirection *string) *FlavorsAzureFlavorsParams {
	o.SetSortDirection(sortDirection)
	return o
}

// SetSortDirection adds the sortDirection to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetSortDirection(sortDirection *string) {
	o.SortDirection = sortDirection
}

// WithStartCPU adds the startCPU to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithStartCPU(startCPU *int32) *FlavorsAzureFlavorsParams {
	o.SetStartCPU(startCPU)
	return o
}

// SetStartCPU adds the startCpu to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetStartCPU(startCPU *int32) {
	o.StartCPU = startCPU
}

// WithStartRAM adds the startRAM to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithStartRAM(startRAM *int32) *FlavorsAzureFlavorsParams {
	o.SetStartRAM(startRAM)
	return o
}

// SetStartRAM adds the startRam to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetStartRAM(startRAM *int32) {
	o.StartRAM = startRAM
}

// WithV adds the v to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) WithV(v string) *FlavorsAzureFlavorsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the flavors azure flavors params
func (o *FlavorsAzureFlavorsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *FlavorsAzureFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
