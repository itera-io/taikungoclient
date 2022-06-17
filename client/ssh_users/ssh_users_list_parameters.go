// Code generated by go-swagger; DO NOT EDIT.

package ssh_users

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
	"github.com/go-openapi/swag"
)

// NewSSHUsersListParams creates a new SSHUsersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSHUsersListParams() *SSHUsersListParams {
	return &SSHUsersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSHUsersListParamsWithTimeout creates a new SSHUsersListParams object
// with the ability to set a timeout on a request.
func NewSSHUsersListParamsWithTimeout(timeout time.Duration) *SSHUsersListParams {
	return &SSHUsersListParams{
		timeout: timeout,
	}
}

// NewSSHUsersListParamsWithContext creates a new SSHUsersListParams object
// with the ability to set a context for a request.
func NewSSHUsersListParamsWithContext(ctx context.Context) *SSHUsersListParams {
	return &SSHUsersListParams{
		Context: ctx,
	}
}

// NewSSHUsersListParamsWithHTTPClient creates a new SSHUsersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSHUsersListParamsWithHTTPClient(client *http.Client) *SSHUsersListParams {
	return &SSHUsersListParams{
		HTTPClient: client,
	}
}

/* SSHUsersListParams contains all the parameters to send to the API endpoint
   for the Ssh users list operation.

   Typically these are written to a http.Request.
*/
type SSHUsersListParams struct {

	// AccessProfileID.
	//
	// Format: int32
	AccessProfileID int32

	// Search.
	Search *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the Ssh users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHUsersListParams) WithDefaults() *SSHUsersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the Ssh users list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSHUsersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the Ssh users list params
func (o *SSHUsersListParams) WithTimeout(timeout time.Duration) *SSHUsersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Ssh users list params
func (o *SSHUsersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Ssh users list params
func (o *SSHUsersListParams) WithContext(ctx context.Context) *SSHUsersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Ssh users list params
func (o *SSHUsersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Ssh users list params
func (o *SSHUsersListParams) WithHTTPClient(client *http.Client) *SSHUsersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Ssh users list params
func (o *SSHUsersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessProfileID adds the accessProfileID to the Ssh users list params
func (o *SSHUsersListParams) WithAccessProfileID(accessProfileID int32) *SSHUsersListParams {
	o.SetAccessProfileID(accessProfileID)
	return o
}

// SetAccessProfileID adds the accessProfileId to the Ssh users list params
func (o *SSHUsersListParams) SetAccessProfileID(accessProfileID int32) {
	o.AccessProfileID = accessProfileID
}

// WithSearch adds the search to the Ssh users list params
func (o *SSHUsersListParams) WithSearch(search *string) *SSHUsersListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the Ssh users list params
func (o *SSHUsersListParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the Ssh users list params
func (o *SSHUsersListParams) WithV(v string) *SSHUsersListParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the Ssh users list params
func (o *SSHUsersListParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *SSHUsersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accessProfileId
	if err := r.SetPathParam("accessProfileId", swag.FormatInt32(o.AccessProfileID)); err != nil {
		return err
	}

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
