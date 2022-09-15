// Code generated by go-swagger; DO NOT EDIT.

package security_group

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

// NewSecurityGroupEditParams creates a new SecurityGroupEditParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityGroupEditParams() *SecurityGroupEditParams {
	return &SecurityGroupEditParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityGroupEditParamsWithTimeout creates a new SecurityGroupEditParams object
// with the ability to set a timeout on a request.
func NewSecurityGroupEditParamsWithTimeout(timeout time.Duration) *SecurityGroupEditParams {
	return &SecurityGroupEditParams{
		timeout: timeout,
	}
}

// NewSecurityGroupEditParamsWithContext creates a new SecurityGroupEditParams object
// with the ability to set a context for a request.
func NewSecurityGroupEditParamsWithContext(ctx context.Context) *SecurityGroupEditParams {
	return &SecurityGroupEditParams{
		Context: ctx,
	}
}

// NewSecurityGroupEditParamsWithHTTPClient creates a new SecurityGroupEditParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityGroupEditParamsWithHTTPClient(client *http.Client) *SecurityGroupEditParams {
	return &SecurityGroupEditParams{
		HTTPClient: client,
	}
}

/*
SecurityGroupEditParams contains all the parameters to send to the API endpoint

	for the security group edit operation.

	Typically these are written to a http.Request.
*/
type SecurityGroupEditParams struct {

	// Body.
	Body *models.EditSecurityGroupCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security group edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupEditParams) WithDefaults() *SecurityGroupEditParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security group edit params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupEditParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security group edit params
func (o *SecurityGroupEditParams) WithTimeout(timeout time.Duration) *SecurityGroupEditParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security group edit params
func (o *SecurityGroupEditParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security group edit params
func (o *SecurityGroupEditParams) WithContext(ctx context.Context) *SecurityGroupEditParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security group edit params
func (o *SecurityGroupEditParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security group edit params
func (o *SecurityGroupEditParams) WithHTTPClient(client *http.Client) *SecurityGroupEditParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security group edit params
func (o *SecurityGroupEditParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the security group edit params
func (o *SecurityGroupEditParams) WithBody(body *models.EditSecurityGroupCommand) *SecurityGroupEditParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the security group edit params
func (o *SecurityGroupEditParams) SetBody(body *models.EditSecurityGroupCommand) {
	o.Body = body
}

// WithV adds the v to the security group edit params
func (o *SecurityGroupEditParams) WithV(v string) *SecurityGroupEditParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the security group edit params
func (o *SecurityGroupEditParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityGroupEditParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
