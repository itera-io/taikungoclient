// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

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

// NewS3CredentialsMakeDefaultParams creates a new S3CredentialsMakeDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3CredentialsMakeDefaultParams() *S3CredentialsMakeDefaultParams {
	return &S3CredentialsMakeDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3CredentialsMakeDefaultParamsWithTimeout creates a new S3CredentialsMakeDefaultParams object
// with the ability to set a timeout on a request.
func NewS3CredentialsMakeDefaultParamsWithTimeout(timeout time.Duration) *S3CredentialsMakeDefaultParams {
	return &S3CredentialsMakeDefaultParams{
		timeout: timeout,
	}
}

// NewS3CredentialsMakeDefaultParamsWithContext creates a new S3CredentialsMakeDefaultParams object
// with the ability to set a context for a request.
func NewS3CredentialsMakeDefaultParamsWithContext(ctx context.Context) *S3CredentialsMakeDefaultParams {
	return &S3CredentialsMakeDefaultParams{
		Context: ctx,
	}
}

// NewS3CredentialsMakeDefaultParamsWithHTTPClient creates a new S3CredentialsMakeDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3CredentialsMakeDefaultParamsWithHTTPClient(client *http.Client) *S3CredentialsMakeDefaultParams {
	return &S3CredentialsMakeDefaultParams{
		HTTPClient: client,
	}
}

/*
S3CredentialsMakeDefaultParams contains all the parameters to send to the API endpoint

	for the s3 credentials make default operation.

	Typically these are written to a http.Request.
*/
type S3CredentialsMakeDefaultParams struct {

	// Body.
	Body S3CredentialsMakeDefaultBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 credentials make default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsMakeDefaultParams) WithDefaults() *S3CredentialsMakeDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 credentials make default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsMakeDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) WithTimeout(timeout time.Duration) *S3CredentialsMakeDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) WithContext(ctx context.Context) *S3CredentialsMakeDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) WithHTTPClient(client *http.Client) *S3CredentialsMakeDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) WithBody(body S3CredentialsMakeDefaultBody) *S3CredentialsMakeDefaultParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) SetBody(body S3CredentialsMakeDefaultBody) {
	o.Body = body
}

// WithV adds the v to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) WithV(v string) *S3CredentialsMakeDefaultParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the s3 credentials make default params
func (o *S3CredentialsMakeDefaultParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *S3CredentialsMakeDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
