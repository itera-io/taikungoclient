// Code generated by go-swagger; DO NOT EDIT.

package user_groups

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

// NewUserGroupsBindProjectGroupsParams creates a new UserGroupsBindProjectGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupsBindProjectGroupsParams() *UserGroupsBindProjectGroupsParams {
	return &UserGroupsBindProjectGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupsBindProjectGroupsParamsWithTimeout creates a new UserGroupsBindProjectGroupsParams object
// with the ability to set a timeout on a request.
func NewUserGroupsBindProjectGroupsParamsWithTimeout(timeout time.Duration) *UserGroupsBindProjectGroupsParams {
	return &UserGroupsBindProjectGroupsParams{
		timeout: timeout,
	}
}

// NewUserGroupsBindProjectGroupsParamsWithContext creates a new UserGroupsBindProjectGroupsParams object
// with the ability to set a context for a request.
func NewUserGroupsBindProjectGroupsParamsWithContext(ctx context.Context) *UserGroupsBindProjectGroupsParams {
	return &UserGroupsBindProjectGroupsParams{
		Context: ctx,
	}
}

// NewUserGroupsBindProjectGroupsParamsWithHTTPClient creates a new UserGroupsBindProjectGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupsBindProjectGroupsParamsWithHTTPClient(client *http.Client) *UserGroupsBindProjectGroupsParams {
	return &UserGroupsBindProjectGroupsParams{
		HTTPClient: client,
	}
}

/* UserGroupsBindProjectGroupsParams contains all the parameters to send to the API endpoint
   for the user groups bind project groups operation.

   Typically these are written to a http.Request.
*/
type UserGroupsBindProjectGroupsParams struct {

	// Body.
	Body *models.BindProjectGroupsToUserGroupCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user groups bind project groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsBindProjectGroupsParams) WithDefaults() *UserGroupsBindProjectGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user groups bind project groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsBindProjectGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) WithTimeout(timeout time.Duration) *UserGroupsBindProjectGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) WithContext(ctx context.Context) *UserGroupsBindProjectGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) WithHTTPClient(client *http.Client) *UserGroupsBindProjectGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) WithBody(body *models.BindProjectGroupsToUserGroupCommand) *UserGroupsBindProjectGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) SetBody(body *models.BindProjectGroupsToUserGroupCommand) {
	o.Body = body
}

// WithV adds the v to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) WithV(v string) *UserGroupsBindProjectGroupsParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the user groups bind project groups params
func (o *UserGroupsBindProjectGroupsParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupsBindProjectGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
