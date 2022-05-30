// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewProjectGroupsListByUserGroupIDParams creates a new ProjectGroupsListByUserGroupIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectGroupsListByUserGroupIDParams() *ProjectGroupsListByUserGroupIDParams {
	return &ProjectGroupsListByUserGroupIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectGroupsListByUserGroupIDParamsWithTimeout creates a new ProjectGroupsListByUserGroupIDParams object
// with the ability to set a timeout on a request.
func NewProjectGroupsListByUserGroupIDParamsWithTimeout(timeout time.Duration) *ProjectGroupsListByUserGroupIDParams {
	return &ProjectGroupsListByUserGroupIDParams{
		timeout: timeout,
	}
}

// NewProjectGroupsListByUserGroupIDParamsWithContext creates a new ProjectGroupsListByUserGroupIDParams object
// with the ability to set a context for a request.
func NewProjectGroupsListByUserGroupIDParamsWithContext(ctx context.Context) *ProjectGroupsListByUserGroupIDParams {
	return &ProjectGroupsListByUserGroupIDParams{
		Context: ctx,
	}
}

// NewProjectGroupsListByUserGroupIDParamsWithHTTPClient creates a new ProjectGroupsListByUserGroupIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectGroupsListByUserGroupIDParamsWithHTTPClient(client *http.Client) *ProjectGroupsListByUserGroupIDParams {
	return &ProjectGroupsListByUserGroupIDParams{
		HTTPClient: client,
	}
}

/* ProjectGroupsListByUserGroupIDParams contains all the parameters to send to the API endpoint
   for the project groups list by user group Id operation.

   Typically these are written to a http.Request.
*/
type ProjectGroupsListByUserGroupIDParams struct {

	// UserGroupID.
	//
	// Format: int32
	UserGroupID *int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the project groups list by user group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsListByUserGroupIDParams) WithDefaults() *ProjectGroupsListByUserGroupIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the project groups list by user group Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectGroupsListByUserGroupIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) WithTimeout(timeout time.Duration) *ProjectGroupsListByUserGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) WithContext(ctx context.Context) *ProjectGroupsListByUserGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) WithHTTPClient(client *http.Client) *ProjectGroupsListByUserGroupIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserGroupID adds the userGroupID to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) WithUserGroupID(userGroupID *int32) *ProjectGroupsListByUserGroupIDParams {
	o.SetUserGroupID(userGroupID)
	return o
}

// SetUserGroupID adds the userGroupId to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) SetUserGroupID(userGroupID *int32) {
	o.UserGroupID = userGroupID
}

// WithV adds the v to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) WithV(v string) *ProjectGroupsListByUserGroupIDParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the project groups list by user group Id params
func (o *ProjectGroupsListByUserGroupIDParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectGroupsListByUserGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserGroupID != nil {

		// query param userGroupId
		var qrUserGroupID int32

		if o.UserGroupID != nil {
			qrUserGroupID = *o.UserGroupID
		}
		qUserGroupID := swag.FormatInt32(qrUserGroupID)
		if qUserGroupID != "" {

			if err := r.SetQueryParam("userGroupId", qUserGroupID); err != nil {
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