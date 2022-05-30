// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

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

// NewKubernetesUpdateKubernetesAlertParams creates a new KubernetesUpdateKubernetesAlertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesUpdateKubernetesAlertParams() *KubernetesUpdateKubernetesAlertParams {
	return &KubernetesUpdateKubernetesAlertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesUpdateKubernetesAlertParamsWithTimeout creates a new KubernetesUpdateKubernetesAlertParams object
// with the ability to set a timeout on a request.
func NewKubernetesUpdateKubernetesAlertParamsWithTimeout(timeout time.Duration) *KubernetesUpdateKubernetesAlertParams {
	return &KubernetesUpdateKubernetesAlertParams{
		timeout: timeout,
	}
}

// NewKubernetesUpdateKubernetesAlertParamsWithContext creates a new KubernetesUpdateKubernetesAlertParams object
// with the ability to set a context for a request.
func NewKubernetesUpdateKubernetesAlertParamsWithContext(ctx context.Context) *KubernetesUpdateKubernetesAlertParams {
	return &KubernetesUpdateKubernetesAlertParams{
		Context: ctx,
	}
}

// NewKubernetesUpdateKubernetesAlertParamsWithHTTPClient creates a new KubernetesUpdateKubernetesAlertParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesUpdateKubernetesAlertParamsWithHTTPClient(client *http.Client) *KubernetesUpdateKubernetesAlertParams {
	return &KubernetesUpdateKubernetesAlertParams{
		HTTPClient: client,
	}
}

/* KubernetesUpdateKubernetesAlertParams contains all the parameters to send to the API endpoint
   for the kubernetes update kubernetes alert operation.

   Typically these are written to a http.Request.
*/
type KubernetesUpdateKubernetesAlertParams struct {

	// AlertID.
	//
	// Format: int32
	AlertID int32

	// Body.
	Body *models.UpdateKubernetesAlertDto

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the kubernetes update kubernetes alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesUpdateKubernetesAlertParams) WithDefaults() *KubernetesUpdateKubernetesAlertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes update kubernetes alert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesUpdateKubernetesAlertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithTimeout(timeout time.Duration) *KubernetesUpdateKubernetesAlertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithContext(ctx context.Context) *KubernetesUpdateKubernetesAlertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithHTTPClient(client *http.Client) *KubernetesUpdateKubernetesAlertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertID adds the alertID to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithAlertID(alertID int32) *KubernetesUpdateKubernetesAlertParams {
	o.SetAlertID(alertID)
	return o
}

// SetAlertID adds the alertId to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetAlertID(alertID int32) {
	o.AlertID = alertID
}

// WithBody adds the body to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithBody(body *models.UpdateKubernetesAlertDto) *KubernetesUpdateKubernetesAlertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetBody(body *models.UpdateKubernetesAlertDto) {
	o.Body = body
}

// WithV adds the v to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) WithV(v string) *KubernetesUpdateKubernetesAlertParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes update kubernetes alert params
func (o *KubernetesUpdateKubernetesAlertParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesUpdateKubernetesAlertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", swag.FormatInt32(o.AlertID)); err != nil {
		return err
	}
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