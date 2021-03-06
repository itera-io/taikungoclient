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
	"github.com/go-openapi/swag"
)

// NewBackupDescribeBackupParams creates a new BackupDescribeBackupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupDescribeBackupParams() *BackupDescribeBackupParams {
	return &BackupDescribeBackupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupDescribeBackupParamsWithTimeout creates a new BackupDescribeBackupParams object
// with the ability to set a timeout on a request.
func NewBackupDescribeBackupParamsWithTimeout(timeout time.Duration) *BackupDescribeBackupParams {
	return &BackupDescribeBackupParams{
		timeout: timeout,
	}
}

// NewBackupDescribeBackupParamsWithContext creates a new BackupDescribeBackupParams object
// with the ability to set a context for a request.
func NewBackupDescribeBackupParamsWithContext(ctx context.Context) *BackupDescribeBackupParams {
	return &BackupDescribeBackupParams{
		Context: ctx,
	}
}

// NewBackupDescribeBackupParamsWithHTTPClient creates a new BackupDescribeBackupParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupDescribeBackupParamsWithHTTPClient(client *http.Client) *BackupDescribeBackupParams {
	return &BackupDescribeBackupParams{
		HTTPClient: client,
	}
}

/* BackupDescribeBackupParams contains all the parameters to send to the API endpoint
   for the backup describe backup operation.

   Typically these are written to a http.Request.
*/
type BackupDescribeBackupParams struct {

	// Name.
	Name string

	// ProjectID.
	//
	// Format: int32
	ProjectID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup describe backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDescribeBackupParams) WithDefaults() *BackupDescribeBackupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup describe backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDescribeBackupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup describe backup params
func (o *BackupDescribeBackupParams) WithTimeout(timeout time.Duration) *BackupDescribeBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup describe backup params
func (o *BackupDescribeBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup describe backup params
func (o *BackupDescribeBackupParams) WithContext(ctx context.Context) *BackupDescribeBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup describe backup params
func (o *BackupDescribeBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup describe backup params
func (o *BackupDescribeBackupParams) WithHTTPClient(client *http.Client) *BackupDescribeBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup describe backup params
func (o *BackupDescribeBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the backup describe backup params
func (o *BackupDescribeBackupParams) WithName(name string) *BackupDescribeBackupParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the backup describe backup params
func (o *BackupDescribeBackupParams) SetName(name string) {
	o.Name = name
}

// WithProjectID adds the projectID to the backup describe backup params
func (o *BackupDescribeBackupParams) WithProjectID(projectID int32) *BackupDescribeBackupParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the backup describe backup params
func (o *BackupDescribeBackupParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the backup describe backup params
func (o *BackupDescribeBackupParams) WithV(v string) *BackupDescribeBackupParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the backup describe backup params
func (o *BackupDescribeBackupParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *BackupDescribeBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", swag.FormatInt32(o.ProjectID)); err != nil {
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
