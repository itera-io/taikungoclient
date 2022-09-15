// Code generated by go-swagger; DO NOT EDIT.

package cloud_credentials

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

// NewCloudCredentialsMakeDefaultParams creates a new CloudCredentialsMakeDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudCredentialsMakeDefaultParams() *CloudCredentialsMakeDefaultParams {
	return &CloudCredentialsMakeDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudCredentialsMakeDefaultParamsWithTimeout creates a new CloudCredentialsMakeDefaultParams object
// with the ability to set a timeout on a request.
func NewCloudCredentialsMakeDefaultParamsWithTimeout(timeout time.Duration) *CloudCredentialsMakeDefaultParams {
	return &CloudCredentialsMakeDefaultParams{
		timeout: timeout,
	}
}

// NewCloudCredentialsMakeDefaultParamsWithContext creates a new CloudCredentialsMakeDefaultParams object
// with the ability to set a context for a request.
func NewCloudCredentialsMakeDefaultParamsWithContext(ctx context.Context) *CloudCredentialsMakeDefaultParams {
	return &CloudCredentialsMakeDefaultParams{
		Context: ctx,
	}
}

// NewCloudCredentialsMakeDefaultParamsWithHTTPClient creates a new CloudCredentialsMakeDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudCredentialsMakeDefaultParamsWithHTTPClient(client *http.Client) *CloudCredentialsMakeDefaultParams {
	return &CloudCredentialsMakeDefaultParams{
		HTTPClient: client,
	}
}

/*
CloudCredentialsMakeDefaultParams contains all the parameters to send to the API endpoint

	for the cloud credentials make default operation.

	Typically these are written to a http.Request.
*/
type CloudCredentialsMakeDefaultParams struct {

	// Body.
	Body *models.CredentialMakeDefaultCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud credentials make default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudCredentialsMakeDefaultParams) WithDefaults() *CloudCredentialsMakeDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud credentials make default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudCredentialsMakeDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) WithTimeout(timeout time.Duration) *CloudCredentialsMakeDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) WithContext(ctx context.Context) *CloudCredentialsMakeDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) WithHTTPClient(client *http.Client) *CloudCredentialsMakeDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) WithBody(body *models.CredentialMakeDefaultCommand) *CloudCredentialsMakeDefaultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) SetBody(body *models.CredentialMakeDefaultCommand) {
	o.Body = body
}

// WithV adds the v to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) WithV(v string) *CloudCredentialsMakeDefaultParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the cloud credentials make default params
func (o *CloudCredentialsMakeDefaultParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CloudCredentialsMakeDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
