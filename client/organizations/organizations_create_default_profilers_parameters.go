// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewOrganizationsCreateDefaultProfilersParams creates a new OrganizationsCreateDefaultProfilersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsCreateDefaultProfilersParams() *OrganizationsCreateDefaultProfilersParams {
	return &OrganizationsCreateDefaultProfilersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsCreateDefaultProfilersParamsWithTimeout creates a new OrganizationsCreateDefaultProfilersParams object
// with the ability to set a timeout on a request.
func NewOrganizationsCreateDefaultProfilersParamsWithTimeout(timeout time.Duration) *OrganizationsCreateDefaultProfilersParams {
	return &OrganizationsCreateDefaultProfilersParams{
		timeout: timeout,
	}
}

// NewOrganizationsCreateDefaultProfilersParamsWithContext creates a new OrganizationsCreateDefaultProfilersParams object
// with the ability to set a context for a request.
func NewOrganizationsCreateDefaultProfilersParamsWithContext(ctx context.Context) *OrganizationsCreateDefaultProfilersParams {
	return &OrganizationsCreateDefaultProfilersParams{
		Context: ctx,
	}
}

// NewOrganizationsCreateDefaultProfilersParamsWithHTTPClient creates a new OrganizationsCreateDefaultProfilersParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsCreateDefaultProfilersParamsWithHTTPClient(client *http.Client) *OrganizationsCreateDefaultProfilersParams {
	return &OrganizationsCreateDefaultProfilersParams{
		HTTPClient: client,
	}
}

/* OrganizationsCreateDefaultProfilersParams contains all the parameters to send to the API endpoint
   for the organizations create default profilers operation.

   Typically these are written to a http.Request.
*/
type OrganizationsCreateDefaultProfilersParams struct {

	// Body.
	Body *models.DefaultProfilesCreateCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations create default profilers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsCreateDefaultProfilersParams) WithDefaults() *OrganizationsCreateDefaultProfilersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations create default profilers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsCreateDefaultProfilersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) WithTimeout(timeout time.Duration) *OrganizationsCreateDefaultProfilersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) WithContext(ctx context.Context) *OrganizationsCreateDefaultProfilersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) WithHTTPClient(client *http.Client) *OrganizationsCreateDefaultProfilersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) WithBody(body *models.DefaultProfilesCreateCommand) *OrganizationsCreateDefaultProfilersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) SetBody(body *models.DefaultProfilesCreateCommand) {
	o.Body = body
}

// WithV adds the v to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) WithV(v string) *OrganizationsCreateDefaultProfilersParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the organizations create default profilers params
func (o *OrganizationsCreateDefaultProfilersParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsCreateDefaultProfilersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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