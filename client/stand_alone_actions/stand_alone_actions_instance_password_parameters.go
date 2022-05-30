// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_actions

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

// NewStandAloneActionsInstancePasswordParams creates a new StandAloneActionsInstancePasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStandAloneActionsInstancePasswordParams() *StandAloneActionsInstancePasswordParams {
	return &StandAloneActionsInstancePasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStandAloneActionsInstancePasswordParamsWithTimeout creates a new StandAloneActionsInstancePasswordParams object
// with the ability to set a timeout on a request.
func NewStandAloneActionsInstancePasswordParamsWithTimeout(timeout time.Duration) *StandAloneActionsInstancePasswordParams {
	return &StandAloneActionsInstancePasswordParams{
		timeout: timeout,
	}
}

// NewStandAloneActionsInstancePasswordParamsWithContext creates a new StandAloneActionsInstancePasswordParams object
// with the ability to set a context for a request.
func NewStandAloneActionsInstancePasswordParamsWithContext(ctx context.Context) *StandAloneActionsInstancePasswordParams {
	return &StandAloneActionsInstancePasswordParams{
		Context: ctx,
	}
}

// NewStandAloneActionsInstancePasswordParamsWithHTTPClient creates a new StandAloneActionsInstancePasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewStandAloneActionsInstancePasswordParamsWithHTTPClient(client *http.Client) *StandAloneActionsInstancePasswordParams {
	return &StandAloneActionsInstancePasswordParams{
		HTTPClient: client,
	}
}

/* StandAloneActionsInstancePasswordParams contains all the parameters to send to the API endpoint
   for the stand alone actions instance password operation.

   Typically these are written to a http.Request.
*/
type StandAloneActionsInstancePasswordParams struct {

	// Config.
	Config runtime.NamedReadCloser

	// ID.
	//
	// Format: int32
	ID *int32

	// Key.
	Key *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stand alone actions instance password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsInstancePasswordParams) WithDefaults() *StandAloneActionsInstancePasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stand alone actions instance password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StandAloneActionsInstancePasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithTimeout(timeout time.Duration) *StandAloneActionsInstancePasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithContext(ctx context.Context) *StandAloneActionsInstancePasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithHTTPClient(client *http.Client) *StandAloneActionsInstancePasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithConfig(config runtime.NamedReadCloser) *StandAloneActionsInstancePasswordParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetConfig(config runtime.NamedReadCloser) {
	o.Config = config
}

// WithID adds the id to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithID(id *int32) *StandAloneActionsInstancePasswordParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetID(id *int32) {
	o.ID = id
}

// WithKey adds the key to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithKey(key *string) *StandAloneActionsInstancePasswordParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetKey(key *string) {
	o.Key = key
}

// WithV adds the v to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) WithV(v string) *StandAloneActionsInstancePasswordParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the stand alone actions instance password params
func (o *StandAloneActionsInstancePasswordParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *StandAloneActionsInstancePasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Config != nil {

		if o.Config != nil {
			// form file param Config
			if err := r.SetFileParam("Config", o.Config); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// form param Id
		var frID int32
		if o.ID != nil {
			frID = *o.ID
		}
		fID := swag.FormatInt32(frID)
		if fID != "" {
			if err := r.SetFormParam("Id", fID); err != nil {
				return err
			}
		}
	}

	if o.Key != nil {

		// form param Key
		var frKey string
		if o.Key != nil {
			frKey = *o.Key
		}
		fKey := frKey
		if fKey != "" {
			if err := r.SetFormParam("Key", fKey); err != nil {
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