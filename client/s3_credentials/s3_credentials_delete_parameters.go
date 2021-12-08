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
	"github.com/go-openapi/swag"
)

// NewS3CredentialsDeleteParams creates a new S3CredentialsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3CredentialsDeleteParams() *S3CredentialsDeleteParams {
	return &S3CredentialsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3CredentialsDeleteParamsWithTimeout creates a new S3CredentialsDeleteParams object
// with the ability to set a timeout on a request.
func NewS3CredentialsDeleteParamsWithTimeout(timeout time.Duration) *S3CredentialsDeleteParams {
	return &S3CredentialsDeleteParams{
		timeout: timeout,
	}
}

// NewS3CredentialsDeleteParamsWithContext creates a new S3CredentialsDeleteParams object
// with the ability to set a context for a request.
func NewS3CredentialsDeleteParamsWithContext(ctx context.Context) *S3CredentialsDeleteParams {
	return &S3CredentialsDeleteParams{
		Context: ctx,
	}
}

// NewS3CredentialsDeleteParamsWithHTTPClient creates a new S3CredentialsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3CredentialsDeleteParamsWithHTTPClient(client *http.Client) *S3CredentialsDeleteParams {
	return &S3CredentialsDeleteParams{
		HTTPClient: client,
	}
}

/* S3CredentialsDeleteParams contains all the parameters to send to the API endpoint
   for the s3 credentials delete operation.

   Typically these are written to a http.Request.
*/
type S3CredentialsDeleteParams struct {

	// ID.
	//
	// Format: int32
	ID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsDeleteParams) WithDefaults() *S3CredentialsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 credentials delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3CredentialsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) WithTimeout(timeout time.Duration) *S3CredentialsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) WithContext(ctx context.Context) *S3CredentialsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) WithHTTPClient(client *http.Client) *S3CredentialsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) WithID(id int32) *S3CredentialsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) SetID(id int32) {
	o.ID = id
}

// WithV adds the v to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) WithV(v string) *S3CredentialsDeleteParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the s3 credentials delete params
func (o *S3CredentialsDeleteParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *S3CredentialsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
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
