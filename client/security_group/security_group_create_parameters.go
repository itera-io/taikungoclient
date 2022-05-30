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

// NewSecurityGroupCreateParams creates a new SecurityGroupCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityGroupCreateParams() *SecurityGroupCreateParams {
	return &SecurityGroupCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityGroupCreateParamsWithTimeout creates a new SecurityGroupCreateParams object
// with the ability to set a timeout on a request.
func NewSecurityGroupCreateParamsWithTimeout(timeout time.Duration) *SecurityGroupCreateParams {
	return &SecurityGroupCreateParams{
		timeout: timeout,
	}
}

// NewSecurityGroupCreateParamsWithContext creates a new SecurityGroupCreateParams object
// with the ability to set a context for a request.
func NewSecurityGroupCreateParamsWithContext(ctx context.Context) *SecurityGroupCreateParams {
	return &SecurityGroupCreateParams{
		Context: ctx,
	}
}

// NewSecurityGroupCreateParamsWithHTTPClient creates a new SecurityGroupCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityGroupCreateParamsWithHTTPClient(client *http.Client) *SecurityGroupCreateParams {
	return &SecurityGroupCreateParams{
		HTTPClient: client,
	}
}

/* SecurityGroupCreateParams contains all the parameters to send to the API endpoint
   for the security group create operation.

   Typically these are written to a http.Request.
*/
type SecurityGroupCreateParams struct {

	// Body.
	Body *models.CreateSecurityGroupCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupCreateParams) WithDefaults() *SecurityGroupCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security group create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityGroupCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security group create params
func (o *SecurityGroupCreateParams) WithTimeout(timeout time.Duration) *SecurityGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security group create params
func (o *SecurityGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security group create params
func (o *SecurityGroupCreateParams) WithContext(ctx context.Context) *SecurityGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security group create params
func (o *SecurityGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security group create params
func (o *SecurityGroupCreateParams) WithHTTPClient(client *http.Client) *SecurityGroupCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security group create params
func (o *SecurityGroupCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the security group create params
func (o *SecurityGroupCreateParams) WithBody(body *models.CreateSecurityGroupCommand) *SecurityGroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the security group create params
func (o *SecurityGroupCreateParams) SetBody(body *models.CreateSecurityGroupCommand) {
	o.Body = body
}

// WithV adds the v to the security group create params
func (o *SecurityGroupCreateParams) WithV(v string) *SecurityGroupCreateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the security group create params
func (o *SecurityGroupCreateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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