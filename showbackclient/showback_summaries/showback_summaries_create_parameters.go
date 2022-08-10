// Code generated by go-swagger; DO NOT EDIT.

package showback_summaries

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

	"github.com/itera-io/taikungoclient/models"
)

// NewShowbackSummariesCreateParams creates a new ShowbackSummariesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackSummariesCreateParams() *ShowbackSummariesCreateParams {
	return &ShowbackSummariesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackSummariesCreateParamsWithTimeout creates a new ShowbackSummariesCreateParams object
// with the ability to set a timeout on a request.
func NewShowbackSummariesCreateParamsWithTimeout(timeout time.Duration) *ShowbackSummariesCreateParams {
	return &ShowbackSummariesCreateParams{
		timeout: timeout,
	}
}

// NewShowbackSummariesCreateParamsWithContext creates a new ShowbackSummariesCreateParams object
// with the ability to set a context for a request.
func NewShowbackSummariesCreateParamsWithContext(ctx context.Context) *ShowbackSummariesCreateParams {
	return &ShowbackSummariesCreateParams{
		Context: ctx,
	}
}

// NewShowbackSummariesCreateParamsWithHTTPClient creates a new ShowbackSummariesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackSummariesCreateParamsWithHTTPClient(client *http.Client) *ShowbackSummariesCreateParams {
	return &ShowbackSummariesCreateParams{
		HTTPClient: client,
	}
}

/* ShowbackSummariesCreateParams contains all the parameters to send to the API endpoint
   for the showback summaries create operation.

   Typically these are written to a http.Request.
*/
type ShowbackSummariesCreateParams struct {

	// Body.
	Body *models.CreateShowbackSummaryCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback summaries create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesCreateParams) WithDefaults() *ShowbackSummariesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback summaries create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackSummariesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback summaries create params
func (o *ShowbackSummariesCreateParams) WithTimeout(timeout time.Duration) *ShowbackSummariesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback summaries create params
func (o *ShowbackSummariesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback summaries create params
func (o *ShowbackSummariesCreateParams) WithContext(ctx context.Context) *ShowbackSummariesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback summaries create params
func (o *ShowbackSummariesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback summaries create params
func (o *ShowbackSummariesCreateParams) WithHTTPClient(client *http.Client) *ShowbackSummariesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback summaries create params
func (o *ShowbackSummariesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback summaries create params
func (o *ShowbackSummariesCreateParams) WithBody(body *models.CreateShowbackSummaryCommand) *ShowbackSummariesCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback summaries create params
func (o *ShowbackSummariesCreateParams) SetBody(body *models.CreateShowbackSummaryCommand) {
	o.Body = body
}

// WithV adds the v to the showback summaries create params
func (o *ShowbackSummariesCreateParams) WithV(v string) *ShowbackSummariesCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback summaries create params
func (o *ShowbackSummariesCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackSummariesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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