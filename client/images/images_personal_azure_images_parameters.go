// Code generated by go-swagger; DO NOT EDIT.

package images

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

// NewImagesPersonalAzureImagesParams creates a new ImagesPersonalAzureImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImagesPersonalAzureImagesParams() *ImagesPersonalAzureImagesParams {
	return &ImagesPersonalAzureImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImagesPersonalAzureImagesParamsWithTimeout creates a new ImagesPersonalAzureImagesParams object
// with the ability to set a timeout on a request.
func NewImagesPersonalAzureImagesParamsWithTimeout(timeout time.Duration) *ImagesPersonalAzureImagesParams {
	return &ImagesPersonalAzureImagesParams{
		timeout: timeout,
	}
}

// NewImagesPersonalAzureImagesParamsWithContext creates a new ImagesPersonalAzureImagesParams object
// with the ability to set a context for a request.
func NewImagesPersonalAzureImagesParamsWithContext(ctx context.Context) *ImagesPersonalAzureImagesParams {
	return &ImagesPersonalAzureImagesParams{
		Context: ctx,
	}
}

// NewImagesPersonalAzureImagesParamsWithHTTPClient creates a new ImagesPersonalAzureImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImagesPersonalAzureImagesParamsWithHTTPClient(client *http.Client) *ImagesPersonalAzureImagesParams {
	return &ImagesPersonalAzureImagesParams{
		HTTPClient: client,
	}
}

/* ImagesPersonalAzureImagesParams contains all the parameters to send to the API endpoint
   for the images personal azure images operation.

   Typically these are written to a http.Request.
*/
type ImagesPersonalAzureImagesParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the images personal azure images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesPersonalAzureImagesParams) WithDefaults() *ImagesPersonalAzureImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the images personal azure images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesPersonalAzureImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) WithTimeout(timeout time.Duration) *ImagesPersonalAzureImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) WithContext(ctx context.Context) *ImagesPersonalAzureImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) WithHTTPClient(client *http.Client) *ImagesPersonalAzureImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) WithCloudID(cloudID int32) *ImagesPersonalAzureImagesParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithV adds the v to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) WithV(v string) *ImagesPersonalAzureImagesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the images personal azure images params
func (o *ImagesPersonalAzureImagesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesPersonalAzureImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudId
	if err := r.SetPathParam("cloudId", swag.FormatInt32(o.CloudID)); err != nil {
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
