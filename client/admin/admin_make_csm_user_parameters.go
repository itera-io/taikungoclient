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

// NewAdminMakeCsmUserParams creates a new AdminMakeCsmUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminMakeCsmUserParams() *AdminMakeCsmUserParams {
	return &AdminMakeCsmUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminMakeCsmUserParamsWithTimeout creates a new AdminMakeCsmUserParams object
// with the ability to set a timeout on a request.
func NewAdminMakeCsmUserParamsWithTimeout(timeout time.Duration) *AdminMakeCsmUserParams {
	return &AdminMakeCsmUserParams{
		timeout: timeout,
	}
}

// NewAdminMakeCsmUserParamsWithContext creates a new AdminMakeCsmUserParams object
// with the ability to set a context for a request.
func NewAdminMakeCsmUserParamsWithContext(ctx context.Context) *AdminMakeCsmUserParams {
	return &AdminMakeCsmUserParams{
		Context: ctx,
	}
}

// NewAdminMakeCsmUserParamsWithHTTPClient creates a new AdminMakeCsmUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminMakeCsmUserParamsWithHTTPClient(client *http.Client) *AdminMakeCsmUserParams {
	return &AdminMakeCsmUserParams{
		HTTPClient: client,
	}
}

/* AdminMakeCsmUserParams contains all the parameters to send to the API endpoint
   for the admin make csm user operation.

   Typically these are written to a http.Request.
*/
type AdminMakeCsmUserParams struct {

	// Body.
	Body *models.MakeCsmCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin make csm user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminMakeCsmUserParams) WithDefaults() *AdminMakeCsmUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin make csm user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminMakeCsmUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin make csm user params
func (o *AdminMakeCsmUserParams) WithTimeout(timeout time.Duration) *AdminMakeCsmUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin make csm user params
func (o *AdminMakeCsmUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin make csm user params
func (o *AdminMakeCsmUserParams) WithContext(ctx context.Context) *AdminMakeCsmUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin make csm user params
func (o *AdminMakeCsmUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin make csm user params
func (o *AdminMakeCsmUserParams) WithHTTPClient(client *http.Client) *AdminMakeCsmUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin make csm user params
func (o *AdminMakeCsmUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin make csm user params
func (o *AdminMakeCsmUserParams) WithBody(body *models.MakeCsmCommand) *AdminMakeCsmUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin make csm user params
func (o *AdminMakeCsmUserParams) SetBody(body *models.MakeCsmCommand) {
	o.Body = body
}

// WithV adds the v to the admin make csm user params
func (o *AdminMakeCsmUserParams) WithV(v string) *AdminMakeCsmUserParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the admin make csm user params
func (o *AdminMakeCsmUserParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AdminMakeCsmUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
