// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewProjectsLockManagerParams creates a new ProjectsLockManagerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProjectsLockManagerParams() *ProjectsLockManagerParams {
	return &ProjectsLockManagerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsLockManagerParamsWithTimeout creates a new ProjectsLockManagerParams object
// with the ability to set a timeout on a request.
func NewProjectsLockManagerParamsWithTimeout(timeout time.Duration) *ProjectsLockManagerParams {
	return &ProjectsLockManagerParams{
		timeout: timeout,
	}
}

// NewProjectsLockManagerParamsWithContext creates a new ProjectsLockManagerParams object
// with the ability to set a context for a request.
func NewProjectsLockManagerParamsWithContext(ctx context.Context) *ProjectsLockManagerParams {
	return &ProjectsLockManagerParams{
		Context: ctx,
	}
}

// NewProjectsLockManagerParamsWithHTTPClient creates a new ProjectsLockManagerParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsLockManagerParamsWithHTTPClient(client *http.Client) *ProjectsLockManagerParams {
	return &ProjectsLockManagerParams{
		HTTPClient: client,
	}
}

/* ProjectsLockManagerParams contains all the parameters to send to the API endpoint
   for the projects lock manager operation.

   Typically these are written to a http.Request.
*/
type ProjectsLockManagerParams struct {

	// ID.
	//
	// Format: int32
	ID *int32

	// Mode.
	Mode *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsLockManagerParams) WithDefaults() *ProjectsLockManagerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects lock manager params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsLockManagerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects lock manager params
func (o *ProjectsLockManagerParams) WithTimeout(timeout time.Duration) *ProjectsLockManagerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects lock manager params
func (o *ProjectsLockManagerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects lock manager params
func (o *ProjectsLockManagerParams) WithContext(ctx context.Context) *ProjectsLockManagerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects lock manager params
func (o *ProjectsLockManagerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects lock manager params
func (o *ProjectsLockManagerParams) WithHTTPClient(client *http.Client) *ProjectsLockManagerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects lock manager params
func (o *ProjectsLockManagerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the projects lock manager params
func (o *ProjectsLockManagerParams) WithID(id *int32) *ProjectsLockManagerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects lock manager params
func (o *ProjectsLockManagerParams) SetID(id *int32) {
	o.ID = id
}

// WithMode adds the mode to the projects lock manager params
func (o *ProjectsLockManagerParams) WithMode(mode *string) *ProjectsLockManagerParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the projects lock manager params
func (o *ProjectsLockManagerParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithV adds the v to the projects lock manager params
func (o *ProjectsLockManagerParams) WithV(v string) *ProjectsLockManagerParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the projects lock manager params
func (o *ProjectsLockManagerParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsLockManagerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID int32

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt32(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
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
