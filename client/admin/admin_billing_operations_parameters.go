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

// NewAdminBillingOperationsParams creates a new AdminBillingOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminBillingOperationsParams() *AdminBillingOperationsParams {
	return &AdminBillingOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminBillingOperationsParamsWithTimeout creates a new AdminBillingOperationsParams object
// with the ability to set a timeout on a request.
func NewAdminBillingOperationsParamsWithTimeout(timeout time.Duration) *AdminBillingOperationsParams {
	return &AdminBillingOperationsParams{
		timeout: timeout,
	}
}

// NewAdminBillingOperationsParamsWithContext creates a new AdminBillingOperationsParams object
// with the ability to set a context for a request.
func NewAdminBillingOperationsParamsWithContext(ctx context.Context) *AdminBillingOperationsParams {
	return &AdminBillingOperationsParams{
		Context: ctx,
	}
}

// NewAdminBillingOperationsParamsWithHTTPClient creates a new AdminBillingOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminBillingOperationsParamsWithHTTPClient(client *http.Client) *AdminBillingOperationsParams {
	return &AdminBillingOperationsParams{
		HTTPClient: client,
	}
}

/* AdminBillingOperationsParams contains all the parameters to send to the API endpoint
   for the admin billing operations operation.

   Typically these are written to a http.Request.
*/
type AdminBillingOperationsParams struct {

	// Body.
	Body *models.AdminBillingOperationCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin billing operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminBillingOperationsParams) WithDefaults() *AdminBillingOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin billing operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminBillingOperationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin billing operations params
func (o *AdminBillingOperationsParams) WithTimeout(timeout time.Duration) *AdminBillingOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin billing operations params
func (o *AdminBillingOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin billing operations params
func (o *AdminBillingOperationsParams) WithContext(ctx context.Context) *AdminBillingOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin billing operations params
func (o *AdminBillingOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin billing operations params
func (o *AdminBillingOperationsParams) WithHTTPClient(client *http.Client) *AdminBillingOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin billing operations params
func (o *AdminBillingOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin billing operations params
func (o *AdminBillingOperationsParams) WithBody(body *models.AdminBillingOperationCommand) *AdminBillingOperationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin billing operations params
func (o *AdminBillingOperationsParams) SetBody(body *models.AdminBillingOperationCommand) {
	o.Body = body
}

// WithV adds the v to the admin billing operations params
func (o *AdminBillingOperationsParams) WithV(v string) *AdminBillingOperationsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the admin billing operations params
func (o *AdminBillingOperationsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AdminBillingOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
