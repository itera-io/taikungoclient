// Code generated by go-swagger; DO NOT EDIT.

package showback_rules

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

// NewShowbackRulesUpdateParams creates a new ShowbackRulesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShowbackRulesUpdateParams() *ShowbackRulesUpdateParams {
	return &ShowbackRulesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShowbackRulesUpdateParamsWithTimeout creates a new ShowbackRulesUpdateParams object
// with the ability to set a timeout on a request.
func NewShowbackRulesUpdateParamsWithTimeout(timeout time.Duration) *ShowbackRulesUpdateParams {
	return &ShowbackRulesUpdateParams{
		timeout: timeout,
	}
}

// NewShowbackRulesUpdateParamsWithContext creates a new ShowbackRulesUpdateParams object
// with the ability to set a context for a request.
func NewShowbackRulesUpdateParamsWithContext(ctx context.Context) *ShowbackRulesUpdateParams {
	return &ShowbackRulesUpdateParams{
		Context: ctx,
	}
}

// NewShowbackRulesUpdateParamsWithHTTPClient creates a new ShowbackRulesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewShowbackRulesUpdateParamsWithHTTPClient(client *http.Client) *ShowbackRulesUpdateParams {
	return &ShowbackRulesUpdateParams{
		HTTPClient: client,
	}
}

/*
ShowbackRulesUpdateParams contains all the parameters to send to the API endpoint

	for the showback rules update operation.

	Typically these are written to a http.Request.
*/
type ShowbackRulesUpdateParams struct {

	// Body.
	Body *models.UpdateShowbackRuleCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the showback rules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesUpdateParams) WithDefaults() *ShowbackRulesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the showback rules update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShowbackRulesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the showback rules update params
func (o *ShowbackRulesUpdateParams) WithTimeout(timeout time.Duration) *ShowbackRulesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the showback rules update params
func (o *ShowbackRulesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the showback rules update params
func (o *ShowbackRulesUpdateParams) WithContext(ctx context.Context) *ShowbackRulesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the showback rules update params
func (o *ShowbackRulesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the showback rules update params
func (o *ShowbackRulesUpdateParams) WithHTTPClient(client *http.Client) *ShowbackRulesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the showback rules update params
func (o *ShowbackRulesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the showback rules update params
func (o *ShowbackRulesUpdateParams) WithBody(body *models.UpdateShowbackRuleCommand) *ShowbackRulesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the showback rules update params
func (o *ShowbackRulesUpdateParams) SetBody(body *models.UpdateShowbackRuleCommand) {
	o.Body = body
}

// WithV adds the v to the showback rules update params
func (o *ShowbackRulesUpdateParams) WithV(v string) *ShowbackRulesUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the showback rules update params
func (o *ShowbackRulesUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ShowbackRulesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
