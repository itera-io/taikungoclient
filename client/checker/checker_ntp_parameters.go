// Code generated by go-swagger; DO NOT EDIT.

package checker

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

// NewCheckerNtpParams creates a new CheckerNtpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckerNtpParams() *CheckerNtpParams {
	return &CheckerNtpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckerNtpParamsWithTimeout creates a new CheckerNtpParams object
// with the ability to set a timeout on a request.
func NewCheckerNtpParamsWithTimeout(timeout time.Duration) *CheckerNtpParams {
	return &CheckerNtpParams{
		timeout: timeout,
	}
}

// NewCheckerNtpParamsWithContext creates a new CheckerNtpParams object
// with the ability to set a context for a request.
func NewCheckerNtpParamsWithContext(ctx context.Context) *CheckerNtpParams {
	return &CheckerNtpParams{
		Context: ctx,
	}
}

// NewCheckerNtpParamsWithHTTPClient creates a new CheckerNtpParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckerNtpParamsWithHTTPClient(client *http.Client) *CheckerNtpParams {
	return &CheckerNtpParams{
		HTTPClient: client,
	}
}

/*
CheckerNtpParams contains all the parameters to send to the API endpoint

	for the checker ntp operation.

	Typically these are written to a http.Request.
*/
type CheckerNtpParams struct {

	// Body.
	Body *models.NtpCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checker ntp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerNtpParams) WithDefaults() *CheckerNtpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checker ntp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerNtpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checker ntp params
func (o *CheckerNtpParams) WithTimeout(timeout time.Duration) *CheckerNtpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checker ntp params
func (o *CheckerNtpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checker ntp params
func (o *CheckerNtpParams) WithContext(ctx context.Context) *CheckerNtpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checker ntp params
func (o *CheckerNtpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checker ntp params
func (o *CheckerNtpParams) WithHTTPClient(client *http.Client) *CheckerNtpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checker ntp params
func (o *CheckerNtpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the checker ntp params
func (o *CheckerNtpParams) WithBody(body *models.NtpCommand) *CheckerNtpParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the checker ntp params
func (o *CheckerNtpParams) SetBody(body *models.NtpCommand) {
	o.Body = body
}

// WithV adds the v to the checker ntp params
func (o *CheckerNtpParams) WithV(v string) *CheckerNtpParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the checker ntp params
func (o *CheckerNtpParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CheckerNtpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
