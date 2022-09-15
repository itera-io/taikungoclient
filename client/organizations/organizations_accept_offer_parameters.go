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

// NewOrganizationsAcceptOfferParams creates a new OrganizationsAcceptOfferParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationsAcceptOfferParams() *OrganizationsAcceptOfferParams {
	return &OrganizationsAcceptOfferParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationsAcceptOfferParamsWithTimeout creates a new OrganizationsAcceptOfferParams object
// with the ability to set a timeout on a request.
func NewOrganizationsAcceptOfferParamsWithTimeout(timeout time.Duration) *OrganizationsAcceptOfferParams {
	return &OrganizationsAcceptOfferParams{
		timeout: timeout,
	}
}

// NewOrganizationsAcceptOfferParamsWithContext creates a new OrganizationsAcceptOfferParams object
// with the ability to set a context for a request.
func NewOrganizationsAcceptOfferParamsWithContext(ctx context.Context) *OrganizationsAcceptOfferParams {
	return &OrganizationsAcceptOfferParams{
		Context: ctx,
	}
}

// NewOrganizationsAcceptOfferParamsWithHTTPClient creates a new OrganizationsAcceptOfferParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationsAcceptOfferParamsWithHTTPClient(client *http.Client) *OrganizationsAcceptOfferParams {
	return &OrganizationsAcceptOfferParams{
		HTTPClient: client,
	}
}

/*
OrganizationsAcceptOfferParams contains all the parameters to send to the API endpoint

	for the organizations accept offer operation.

	Typically these are written to a http.Request.
*/
type OrganizationsAcceptOfferParams struct {

	// Body.
	Body models.AcceptOfferCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organizations accept offer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsAcceptOfferParams) WithDefaults() *OrganizationsAcceptOfferParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organizations accept offer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationsAcceptOfferParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) WithTimeout(timeout time.Duration) *OrganizationsAcceptOfferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) WithContext(ctx context.Context) *OrganizationsAcceptOfferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) WithHTTPClient(client *http.Client) *OrganizationsAcceptOfferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) WithBody(body models.AcceptOfferCommand) *OrganizationsAcceptOfferParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) SetBody(body models.AcceptOfferCommand) {
	o.Body = body
}

// WithV adds the v to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) WithV(v string) *OrganizationsAcceptOfferParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the organizations accept offer params
func (o *OrganizationsAcceptOfferParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationsAcceptOfferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
