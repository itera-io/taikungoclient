// Code generated by go-swagger; DO NOT EDIT.

package kube_config_role

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

// NewKubeConfigRoleListParams creates a new KubeConfigRoleListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubeConfigRoleListParams() *KubeConfigRoleListParams {
	return &KubeConfigRoleListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubeConfigRoleListParamsWithTimeout creates a new KubeConfigRoleListParams object
// with the ability to set a timeout on a request.
func NewKubeConfigRoleListParamsWithTimeout(timeout time.Duration) *KubeConfigRoleListParams {
	return &KubeConfigRoleListParams{
		timeout: timeout,
	}
}

// NewKubeConfigRoleListParamsWithContext creates a new KubeConfigRoleListParams object
// with the ability to set a context for a request.
func NewKubeConfigRoleListParamsWithContext(ctx context.Context) *KubeConfigRoleListParams {
	return &KubeConfigRoleListParams{
		Context: ctx,
	}
}

// NewKubeConfigRoleListParamsWithHTTPClient creates a new KubeConfigRoleListParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubeConfigRoleListParamsWithHTTPClient(client *http.Client) *KubeConfigRoleListParams {
	return &KubeConfigRoleListParams{
		HTTPClient: client,
	}
}

/* KubeConfigRoleListParams contains all the parameters to send to the API endpoint
   for the kube config role list operation.

   Typically these are written to a http.Request.
*/
type KubeConfigRoleListParams struct {

	// Search.
	Search *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kube config role list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubeConfigRoleListParams) WithDefaults() *KubeConfigRoleListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kube config role list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubeConfigRoleListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kube config role list params
func (o *KubeConfigRoleListParams) WithTimeout(timeout time.Duration) *KubeConfigRoleListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kube config role list params
func (o *KubeConfigRoleListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kube config role list params
func (o *KubeConfigRoleListParams) WithContext(ctx context.Context) *KubeConfigRoleListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kube config role list params
func (o *KubeConfigRoleListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kube config role list params
func (o *KubeConfigRoleListParams) WithHTTPClient(client *http.Client) *KubeConfigRoleListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kube config role list params
func (o *KubeConfigRoleListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSearch adds the search to the kube config role list params
func (o *KubeConfigRoleListParams) WithSearch(search *string) *KubeConfigRoleListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the kube config role list params
func (o *KubeConfigRoleListParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the kube config role list params
func (o *KubeConfigRoleListParams) WithV(v string) *KubeConfigRoleListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kube config role list params
func (o *KubeConfigRoleListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubeConfigRoleListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
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
