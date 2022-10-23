// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewAdminUpdateUserKubeConfigParams creates a new AdminUpdateUserKubeConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminUpdateUserKubeConfigParams() *AdminUpdateUserKubeConfigParams {
	return &AdminUpdateUserKubeConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateUserKubeConfigParamsWithTimeout creates a new AdminUpdateUserKubeConfigParams object
// with the ability to set a timeout on a request.
func NewAdminUpdateUserKubeConfigParamsWithTimeout(timeout time.Duration) *AdminUpdateUserKubeConfigParams {
	return &AdminUpdateUserKubeConfigParams{
		timeout: timeout,
	}
}

// NewAdminUpdateUserKubeConfigParamsWithContext creates a new AdminUpdateUserKubeConfigParams object
// with the ability to set a context for a request.
func NewAdminUpdateUserKubeConfigParamsWithContext(ctx context.Context) *AdminUpdateUserKubeConfigParams {
	return &AdminUpdateUserKubeConfigParams{
		Context: ctx,
	}
}

// NewAdminUpdateUserKubeConfigParamsWithHTTPClient creates a new AdminUpdateUserKubeConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminUpdateUserKubeConfigParamsWithHTTPClient(client *http.Client) *AdminUpdateUserKubeConfigParams {
	return &AdminUpdateUserKubeConfigParams{
		HTTPClient: client,
	}
}

/*
AdminUpdateUserKubeConfigParams contains all the parameters to send to the API endpoint

	for the admin update user kube config operation.

	Typically these are written to a http.Request.
*/
type AdminUpdateUserKubeConfigParams struct {

	// Body.
	Body *models.AdminUpdateUserKubeConfigCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin update user kube config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserKubeConfigParams) WithDefaults() *AdminUpdateUserKubeConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin update user kube config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateUserKubeConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) WithTimeout(timeout time.Duration) *AdminUpdateUserKubeConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) WithContext(ctx context.Context) *AdminUpdateUserKubeConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) WithHTTPClient(client *http.Client) *AdminUpdateUserKubeConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) WithBody(body *models.AdminUpdateUserKubeConfigCommand) *AdminUpdateUserKubeConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) SetBody(body *models.AdminUpdateUserKubeConfigCommand) {
	o.Body = body
}

// WithV adds the v to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) WithV(v string) *AdminUpdateUserKubeConfigParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the admin update user kube config params
func (o *AdminUpdateUserKubeConfigParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateUserKubeConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
