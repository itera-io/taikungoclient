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

	"github.com/itera-io/taikungoclient/models"
)

// NewShowbackDeleteRuleParams creates a new ShowbackDeleteRuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackDeleteRuleParams() *ShowbackDeleteRuleParams {
	return &ShowbackDeleteRuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackDeleteRuleParamsWithTimeout creates a new ShowbackDeleteRuleParams object
// with the ability to set a timeout on a request.
func NewShowbackDeleteRuleParamsWithTimeout(timeout time.Duration) *ShowbackDeleteRuleParams {
	return &ShowbackDeleteRuleParams{
		timeout: timeout,
	}
}

// NewShowbackDeleteRuleParamsWithContext creates a new ShowbackDeleteRuleParams object
// with the ability to set a context for a request.
func NewShowbackDeleteRuleParamsWithContext(ctx context.Context) *ShowbackDeleteRuleParams {
	return &ShowbackDeleteRuleParams{
		Context: ctx,
	}
}

// NewShowbackDeleteRuleParamsWithHTTPClient creates a new ShowbackDeleteRuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackDeleteRuleParamsWithHTTPClient(client *http.Client) *ShowbackDeleteRuleParams {
	return &ShowbackDeleteRuleParams{
		HTTPClient: client,
	}
}

/* ShowbackDeleteRuleParams contains all the parameters to send to the API endpoint
   for the showback delete rule operation.

   Typically these are written to a http.Request.
*/
type ShowbackDeleteRuleParams struct {

	// Body.
	Body *models.DeleteShowbackRuleCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback delete rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackDeleteRuleParams) WithDefaults() *ShowbackDeleteRuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback delete rule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackDeleteRuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback delete rule params
func (o *ShowbackDeleteRuleParams) WithTimeout(timeout time.Duration) *ShowbackDeleteRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback delete rule params
func (o *ShowbackDeleteRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback delete rule params
func (o *ShowbackDeleteRuleParams) WithContext(ctx context.Context) *ShowbackDeleteRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback delete rule params
func (o *ShowbackDeleteRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback delete rule params
func (o *ShowbackDeleteRuleParams) WithHTTPClient(client *http.Client) *ShowbackDeleteRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback delete rule params
func (o *ShowbackDeleteRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback delete rule params
func (o *ShowbackDeleteRuleParams) WithBody(body *models.DeleteShowbackRuleCommand) *ShowbackDeleteRuleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback delete rule params
func (o *ShowbackDeleteRuleParams) SetBody(body *models.DeleteShowbackRuleCommand) {
	o.Body = body
}

// WithV adds the v to the showback delete rule params
func (o *ShowbackDeleteRuleParams) WithV(v string) *ShowbackDeleteRuleParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback delete rule params
func (o *ShowbackDeleteRuleParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackDeleteRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
