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

// NewBackupDeleteRestoreParams creates a new BackupDeleteRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupDeleteRestoreParams() *BackupDeleteRestoreParams {
	return &BackupDeleteRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupDeleteRestoreParamsWithTimeout creates a new BackupDeleteRestoreParams object
// with the ability to set a timeout on a request.
func NewBackupDeleteRestoreParamsWithTimeout(timeout time.Duration) *BackupDeleteRestoreParams {
	return &BackupDeleteRestoreParams{
		timeout: timeout,
	}
}

// NewBackupDeleteRestoreParamsWithContext creates a new BackupDeleteRestoreParams object
// with the ability to set a context for a request.
func NewBackupDeleteRestoreParamsWithContext(ctx context.Context) *BackupDeleteRestoreParams {
	return &BackupDeleteRestoreParams{
		Context: ctx,
	}
}

// NewBackupDeleteRestoreParamsWithHTTPClient creates a new BackupDeleteRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupDeleteRestoreParamsWithHTTPClient(client *http.Client) *BackupDeleteRestoreParams {
	return &BackupDeleteRestoreParams{
		HTTPClient: client,
	}
}

/* BackupDeleteRestoreParams contains all the parameters to send to the API endpoint
   for the backup delete restore operation.

   Typically these are written to a http.Request.
*/
type BackupDeleteRestoreParams struct {

	// Body.
	Body *models.DeleteRestoreCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup delete restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDeleteRestoreParams) WithDefaults() *BackupDeleteRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup delete restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDeleteRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup delete restore params
func (o *BackupDeleteRestoreParams) WithTimeout(timeout time.Duration) *BackupDeleteRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup delete restore params
func (o *BackupDeleteRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup delete restore params
func (o *BackupDeleteRestoreParams) WithContext(ctx context.Context) *BackupDeleteRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup delete restore params
func (o *BackupDeleteRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup delete restore params
func (o *BackupDeleteRestoreParams) WithHTTPClient(client *http.Client) *BackupDeleteRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup delete restore params
func (o *BackupDeleteRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the backup delete restore params
func (o *BackupDeleteRestoreParams) WithBody(body *models.DeleteRestoreCommand) *BackupDeleteRestoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup delete restore params
func (o *BackupDeleteRestoreParams) SetBody(body *models.DeleteRestoreCommand) {
	o.Body = body
}

// WithV adds the v to the backup delete restore params
func (o *BackupDeleteRestoreParams) WithV(v string) *BackupDeleteRestoreParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the backup delete restore params
func (o *BackupDeleteRestoreParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BackupDeleteRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
