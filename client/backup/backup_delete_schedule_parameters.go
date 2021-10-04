// Code generated by go-swagger; DO NOT EDIT.

package backup

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

// NewBackupDeleteScheduleParams creates a new BackupDeleteScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupDeleteScheduleParams() *BackupDeleteScheduleParams {
	return &BackupDeleteScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupDeleteScheduleParamsWithTimeout creates a new BackupDeleteScheduleParams object
// with the ability to set a timeout on a request.
func NewBackupDeleteScheduleParamsWithTimeout(timeout time.Duration) *BackupDeleteScheduleParams {
	return &BackupDeleteScheduleParams{
		timeout: timeout,
	}
}

// NewBackupDeleteScheduleParamsWithContext creates a new BackupDeleteScheduleParams object
// with the ability to set a context for a request.
func NewBackupDeleteScheduleParamsWithContext(ctx context.Context) *BackupDeleteScheduleParams {
	return &BackupDeleteScheduleParams{
		Context: ctx,
	}
}

// NewBackupDeleteScheduleParamsWithHTTPClient creates a new BackupDeleteScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupDeleteScheduleParamsWithHTTPClient(client *http.Client) *BackupDeleteScheduleParams {
	return &BackupDeleteScheduleParams{
		HTTPClient: client,
	}
}

/* BackupDeleteScheduleParams contains all the parameters to send to the API endpoint
   for the backup delete schedule operation.

   Typically these are written to a http.Request.
*/
type BackupDeleteScheduleParams struct {

	// Body.
	Body *models.DeleteScheduleCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup delete schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDeleteScheduleParams) WithDefaults() *BackupDeleteScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup delete schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDeleteScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup delete schedule params
func (o *BackupDeleteScheduleParams) WithTimeout(timeout time.Duration) *BackupDeleteScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup delete schedule params
func (o *BackupDeleteScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup delete schedule params
func (o *BackupDeleteScheduleParams) WithContext(ctx context.Context) *BackupDeleteScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup delete schedule params
func (o *BackupDeleteScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup delete schedule params
func (o *BackupDeleteScheduleParams) WithHTTPClient(client *http.Client) *BackupDeleteScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup delete schedule params
func (o *BackupDeleteScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the backup delete schedule params
func (o *BackupDeleteScheduleParams) WithBody(body *models.DeleteScheduleCommand) *BackupDeleteScheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup delete schedule params
func (o *BackupDeleteScheduleParams) SetBody(body *models.DeleteScheduleCommand) {
	o.Body = body
}

// WithV adds the v to the backup delete schedule params
func (o *BackupDeleteScheduleParams) WithV(v string) *BackupDeleteScheduleParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the backup delete schedule params
func (o *BackupDeleteScheduleParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BackupDeleteScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
