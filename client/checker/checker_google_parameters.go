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
)

// NewCheckerGoogleParams creates a new CheckerGoogleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckerGoogleParams() *CheckerGoogleParams {
	return &CheckerGoogleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckerGoogleParamsWithTimeout creates a new CheckerGoogleParams object
// with the ability to set a timeout on a request.
func NewCheckerGoogleParamsWithTimeout(timeout time.Duration) *CheckerGoogleParams {
	return &CheckerGoogleParams{
		timeout: timeout,
	}
}

// NewCheckerGoogleParamsWithContext creates a new CheckerGoogleParams object
// with the ability to set a context for a request.
func NewCheckerGoogleParamsWithContext(ctx context.Context) *CheckerGoogleParams {
	return &CheckerGoogleParams{
		Context: ctx,
	}
}

// NewCheckerGoogleParamsWithHTTPClient creates a new CheckerGoogleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckerGoogleParamsWithHTTPClient(client *http.Client) *CheckerGoogleParams {
	return &CheckerGoogleParams{
		HTTPClient: client,
	}
}

/* CheckerGoogleParams contains all the parameters to send to the API endpoint
   for the checker google operation.

   Typically these are written to a http.Request.
*/
type CheckerGoogleParams struct {

	// Config.
	Config runtime.NamedReadCloser

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checker google params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerGoogleParams) WithDefaults() *CheckerGoogleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checker google params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerGoogleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checker google params
func (o *CheckerGoogleParams) WithTimeout(timeout time.Duration) *CheckerGoogleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checker google params
func (o *CheckerGoogleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checker google params
func (o *CheckerGoogleParams) WithContext(ctx context.Context) *CheckerGoogleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checker google params
func (o *CheckerGoogleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checker google params
func (o *CheckerGoogleParams) WithHTTPClient(client *http.Client) *CheckerGoogleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checker google params
func (o *CheckerGoogleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the checker google params
func (o *CheckerGoogleParams) WithConfig(config runtime.NamedReadCloser) *CheckerGoogleParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the checker google params
func (o *CheckerGoogleParams) SetConfig(config runtime.NamedReadCloser) {
	o.Config = config
}

// WithV adds the v to the checker google params
func (o *CheckerGoogleParams) WithV(v string) *CheckerGoogleParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the checker google params
func (o *CheckerGoogleParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CheckerGoogleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {

		if o.Config != nil {
			// form file param Config
			if err := r.SetFileParam("Config", o.Config); err != nil {
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