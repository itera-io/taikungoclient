// Code generated by go-swagger; DO NOT EDIT.

package keycloak

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
)

// NewKeycloakLoginParams creates a new KeycloakLoginParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeycloakLoginParams() *KeycloakLoginParams {
	return &KeycloakLoginParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeycloakLoginParamsWithTimeout creates a new KeycloakLoginParams object
// with the ability to set a timeout on a request.
func NewKeycloakLoginParamsWithTimeout(timeout time.Duration) *KeycloakLoginParams {
	return &KeycloakLoginParams{
		timeout: timeout,
	}
}

// NewKeycloakLoginParamsWithContext creates a new KeycloakLoginParams object
// with the ability to set a context for a request.
func NewKeycloakLoginParamsWithContext(ctx context.Context) *KeycloakLoginParams {
	return &KeycloakLoginParams{
		Context: ctx,
	}
}

// NewKeycloakLoginParamsWithHTTPClient creates a new KeycloakLoginParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeycloakLoginParamsWithHTTPClient(client *http.Client) *KeycloakLoginParams {
	return &KeycloakLoginParams{
		HTTPClient: client,
	}
}

/*
KeycloakLoginParams contains all the parameters to send to the API endpoint

	for the keycloak login operation.

	Typically these are written to a http.Request.
*/
type KeycloakLoginParams struct {

	// Body.
	Body KeycloakLoginBody

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the keycloak login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeycloakLoginParams) WithDefaults() *KeycloakLoginParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the keycloak login params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeycloakLoginParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the keycloak login params
func (o *KeycloakLoginParams) WithTimeout(timeout time.Duration) *KeycloakLoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the keycloak login params
func (o *KeycloakLoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the keycloak login params
func (o *KeycloakLoginParams) WithContext(ctx context.Context) *KeycloakLoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the keycloak login params
func (o *KeycloakLoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the keycloak login params
func (o *KeycloakLoginParams) WithHTTPClient(client *http.Client) *KeycloakLoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the keycloak login params
func (o *KeycloakLoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the keycloak login params
func (o *KeycloakLoginParams) WithBody(body KeycloakLoginBody) *KeycloakLoginParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the keycloak login params
func (o *KeycloakLoginParams) SetBody(body KeycloakLoginBody) {
	o.Body = body
}

// WithV adds the v to the keycloak login params
func (o *KeycloakLoginParams) WithV(v string) *KeycloakLoginParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the keycloak login params
func (o *KeycloakLoginParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KeycloakLoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
