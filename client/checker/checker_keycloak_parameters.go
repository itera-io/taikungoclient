// Code generated by go-swagger; DO NOT EDIT.

package checker

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

// NewCheckerKeycloakParams creates a new CheckerKeycloakParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckerKeycloakParams() *CheckerKeycloakParams {
	return &CheckerKeycloakParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckerKeycloakParamsWithTimeout creates a new CheckerKeycloakParams object
// with the ability to set a timeout on a request.
func NewCheckerKeycloakParamsWithTimeout(timeout time.Duration) *CheckerKeycloakParams {
	return &CheckerKeycloakParams{
		timeout: timeout,
	}
}

// NewCheckerKeycloakParamsWithContext creates a new CheckerKeycloakParams object
// with the ability to set a context for a request.
func NewCheckerKeycloakParamsWithContext(ctx context.Context) *CheckerKeycloakParams {
	return &CheckerKeycloakParams{
		Context: ctx,
	}
}

// NewCheckerKeycloakParamsWithHTTPClient creates a new CheckerKeycloakParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckerKeycloakParamsWithHTTPClient(client *http.Client) *CheckerKeycloakParams {
	return &CheckerKeycloakParams{
		HTTPClient: client,
	}
}

/*
CheckerKeycloakParams contains all the parameters to send to the API endpoint

	for the checker keycloak operation.

	Typically these are written to a http.Request.
*/
type CheckerKeycloakParams struct {

	// Body.
	Body *models.KeycloakCheckerCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checker keycloak params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerKeycloakParams) WithDefaults() *CheckerKeycloakParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checker keycloak params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckerKeycloakParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checker keycloak params
func (o *CheckerKeycloakParams) WithTimeout(timeout time.Duration) *CheckerKeycloakParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checker keycloak params
func (o *CheckerKeycloakParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checker keycloak params
func (o *CheckerKeycloakParams) WithContext(ctx context.Context) *CheckerKeycloakParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checker keycloak params
func (o *CheckerKeycloakParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checker keycloak params
func (o *CheckerKeycloakParams) WithHTTPClient(client *http.Client) *CheckerKeycloakParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checker keycloak params
func (o *CheckerKeycloakParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the checker keycloak params
func (o *CheckerKeycloakParams) WithBody(body *models.KeycloakCheckerCommand) *CheckerKeycloakParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the checker keycloak params
func (o *CheckerKeycloakParams) SetBody(body *models.KeycloakCheckerCommand) {
	o.Body = body
}

// WithV adds the v to the checker keycloak params
func (o *CheckerKeycloakParams) WithV(v string) *CheckerKeycloakParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the checker keycloak params
func (o *CheckerKeycloakParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *CheckerKeycloakParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
