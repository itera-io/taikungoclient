// Code generated by go-swagger; DO NOT EDIT.

package showback

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

// NewShowbackGroupedShowbackSummaryListParams creates a new ShowbackGroupedShowbackSummaryListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackGroupedShowbackSummaryListParams() *ShowbackGroupedShowbackSummaryListParams {
	return &ShowbackGroupedShowbackSummaryListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackGroupedShowbackSummaryListParamsWithTimeout creates a new ShowbackGroupedShowbackSummaryListParams object
// with the ability to set a timeout on a request.
func NewShowbackGroupedShowbackSummaryListParamsWithTimeout(timeout time.Duration) *ShowbackGroupedShowbackSummaryListParams {
	return &ShowbackGroupedShowbackSummaryListParams{
		timeout: timeout,
	}
}

// NewShowbackGroupedShowbackSummaryListParamsWithContext creates a new ShowbackGroupedShowbackSummaryListParams object
// with the ability to set a context for a request.
func NewShowbackGroupedShowbackSummaryListParamsWithContext(ctx context.Context) *ShowbackGroupedShowbackSummaryListParams {
	return &ShowbackGroupedShowbackSummaryListParams{
		Context: ctx,
	}
}

// NewShowbackGroupedShowbackSummaryListParamsWithHTTPClient creates a new ShowbackGroupedShowbackSummaryListParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackGroupedShowbackSummaryListParamsWithHTTPClient(client *http.Client) *ShowbackGroupedShowbackSummaryListParams {
	return &ShowbackGroupedShowbackSummaryListParams{
		HTTPClient: client,
	}
}

/* ShowbackGroupedShowbackSummaryListParams contains all the parameters to send to the API endpoint
   for the showback grouped showback summary list operation.

   Typically these are written to a http.Request.
*/
type ShowbackGroupedShowbackSummaryListParams struct {

	// OrganizationID.
	//
	// Format: int32
	OrganizationID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback grouped showback summary list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackGroupedShowbackSummaryListParams) WithDefaults() *ShowbackGroupedShowbackSummaryListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback grouped showback summary list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackGroupedShowbackSummaryListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) WithTimeout(timeout time.Duration) *ShowbackGroupedShowbackSummaryListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) WithContext(ctx context.Context) *ShowbackGroupedShowbackSummaryListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) WithHTTPClient(client *http.Client) *ShowbackGroupedShowbackSummaryListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) WithOrganizationID(organizationID *int32) *ShowbackGroupedShowbackSummaryListParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) SetOrganizationID(organizationID *int32) {
	o.OrganizationID = organizationID
}

// WithV adds the v to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) WithV(v string) *ShowbackGroupedShowbackSummaryListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback grouped showback summary list params
func (o *ShowbackGroupedShowbackSummaryListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackGroupedShowbackSummaryListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param v
	if err := r.SetPathParam("v", o.V); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
