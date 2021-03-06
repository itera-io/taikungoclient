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
	"github.com/go-openapi/swag"

	"github.com/itera-io/taikungoclient/models"
)

// NewUserGroupsUpdateParams creates a new UserGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupsUpdateParams() *UserGroupsUpdateParams {
	return &UserGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupsUpdateParamsWithTimeout creates a new UserGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewUserGroupsUpdateParamsWithTimeout(timeout time.Duration) *UserGroupsUpdateParams {
	return &UserGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewUserGroupsUpdateParamsWithContext creates a new UserGroupsUpdateParams object
// with the ability to set a context for a request.
func NewUserGroupsUpdateParamsWithContext(ctx context.Context) *UserGroupsUpdateParams {
	return &UserGroupsUpdateParams{
		Context: ctx,
	}
}

// NewUserGroupsUpdateParamsWithHTTPClient creates a new UserGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupsUpdateParamsWithHTTPClient(client *http.Client) *UserGroupsUpdateParams {
	return &UserGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* UserGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the user groups update operation.

   Typically these are written to a http.Request.
*/
type UserGroupsUpdateParams struct {

	// Body.
	Body *models.UpdateUserGroupDto

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

// WithDefaults hydrates default values in the user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsUpdateParams) WithDefaults() *UserGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user groups update params
func (o *UserGroupsUpdateParams) WithTimeout(timeout time.Duration) *UserGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user groups update params
func (o *UserGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user groups update params
func (o *UserGroupsUpdateParams) WithContext(ctx context.Context) *UserGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user groups update params
func (o *UserGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user groups update params
func (o *UserGroupsUpdateParams) WithHTTPClient(client *http.Client) *UserGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user groups update params
func (o *UserGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user groups update params
func (o *UserGroupsUpdateParams) WithBody(body *models.UpdateUserGroupDto) *UserGroupsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user groups update params
func (o *UserGroupsUpdateParams) SetBody(body *models.UpdateUserGroupDto) {
	o.Body = body
}

// WithUserGroupID adds the userGroupID to the user groups update params
func (o *UserGroupsUpdateParams) WithUserGroupID(userGroupID *int32) *UserGroupsUpdateParams {
	o.SetUserGroupID(userGroupID)
	return o
}

// SetUserGroupID adds the userGroupId to the user groups update params
func (o *UserGroupsUpdateParams) SetUserGroupID(userGroupID *int32) {
	o.UserGroupID = userGroupID
}

// WithV adds the v to the user groups update params
func (o *UserGroupsUpdateParams) WithV(v string) *UserGroupsUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the user groups update params
func (o *UserGroupsUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
