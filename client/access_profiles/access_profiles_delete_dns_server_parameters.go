// Code generated by go-swagger; DO NOT EDIT.

package access_profiles

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

// NewAccessProfilesDeleteDNSServerParams creates a new AccessProfilesDeleteDNSServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccessProfilesDeleteDNSServerParams() *AccessProfilesDeleteDNSServerParams {
	return &AccessProfilesDeleteDNSServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccessProfilesDeleteDNSServerParamsWithTimeout creates a new AccessProfilesDeleteDNSServerParams object
// with the ability to set a timeout on a request.
func NewAccessProfilesDeleteDNSServerParamsWithTimeout(timeout time.Duration) *AccessProfilesDeleteDNSServerParams {
	return &AccessProfilesDeleteDNSServerParams{
		timeout: timeout,
	}
}

// NewAccessProfilesDeleteDNSServerParamsWithContext creates a new AccessProfilesDeleteDNSServerParams object
// with the ability to set a context for a request.
func NewAccessProfilesDeleteDNSServerParamsWithContext(ctx context.Context) *AccessProfilesDeleteDNSServerParams {
	return &AccessProfilesDeleteDNSServerParams{
		Context: ctx,
	}
}

// NewAccessProfilesDeleteDNSServerParamsWithHTTPClient creates a new AccessProfilesDeleteDNSServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccessProfilesDeleteDNSServerParamsWithHTTPClient(client *http.Client) *AccessProfilesDeleteDNSServerParams {
	return &AccessProfilesDeleteDNSServerParams{
		HTTPClient: client,
	}
}

/* AccessProfilesDeleteDNSServerParams contains all the parameters to send to the API endpoint
   for the access profiles delete Dns server operation.

   Typically these are written to a http.Request.
*/
type AccessProfilesDeleteDNSServerParams struct {

	// Body.
	Body *models.DNSServerDeleteCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the access profiles delete Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesDeleteDNSServerParams) WithDefaults() *AccessProfilesDeleteDNSServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the access profiles delete Dns server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccessProfilesDeleteDNSServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) WithTimeout(timeout time.Duration) *AccessProfilesDeleteDNSServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) WithContext(ctx context.Context) *AccessProfilesDeleteDNSServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) WithHTTPClient(client *http.Client) *AccessProfilesDeleteDNSServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) WithBody(body *models.DNSServerDeleteCommand) *AccessProfilesDeleteDNSServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) SetBody(body *models.DNSServerDeleteCommand) {
	o.Body = body
}

// WithV adds the v to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) WithV(v string) *AccessProfilesDeleteDNSServerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the access profiles delete Dns server params
func (o *AccessProfilesDeleteDNSServerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AccessProfilesDeleteDNSServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
