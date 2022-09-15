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

// NewImagesPersonalAwsImagesParams creates a new ImagesPersonalAwsImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImagesPersonalAwsImagesParams() *ImagesPersonalAwsImagesParams {
	return &ImagesPersonalAwsImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImagesPersonalAwsImagesParamsWithTimeout creates a new ImagesPersonalAwsImagesParams object
// with the ability to set a timeout on a request.
func NewImagesPersonalAwsImagesParamsWithTimeout(timeout time.Duration) *ImagesPersonalAwsImagesParams {
	return &ImagesPersonalAwsImagesParams{
		timeout: timeout,
	}
}

// NewImagesPersonalAwsImagesParamsWithContext creates a new ImagesPersonalAwsImagesParams object
// with the ability to set a context for a request.
func NewImagesPersonalAwsImagesParamsWithContext(ctx context.Context) *ImagesPersonalAwsImagesParams {
	return &ImagesPersonalAwsImagesParams{
		Context: ctx,
	}
}

// NewImagesPersonalAwsImagesParamsWithHTTPClient creates a new ImagesPersonalAwsImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewImagesPersonalAwsImagesParamsWithHTTPClient(client *http.Client) *ImagesPersonalAwsImagesParams {
	return &ImagesPersonalAwsImagesParams{
		HTTPClient: client,
	}
}

/*
ImagesPersonalAwsImagesParams contains all the parameters to send to the API endpoint

	for the images personal aws images operation.

	Typically these are written to a http.Request.
*/
type ImagesPersonalAwsImagesParams struct {

	// CloudID.
	//
	// Format: int32
	CloudID int32

	// Search.
	Search *string

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the images personal aws images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesPersonalAwsImagesParams) WithDefaults() *ImagesPersonalAwsImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the images personal aws images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImagesPersonalAwsImagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithTimeout(timeout time.Duration) *ImagesPersonalAwsImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithContext(ctx context.Context) *ImagesPersonalAwsImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithHTTPClient(client *http.Client) *ImagesPersonalAwsImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudID adds the cloudID to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithCloudID(cloudID int32) *ImagesPersonalAwsImagesParams {
	o.SetCloudID(cloudID)
	return o
}

// SetCloudID adds the cloudId to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetCloudID(cloudID int32) {
	o.CloudID = cloudID
}

// WithSearch adds the search to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithSearch(search *string) *ImagesPersonalAwsImagesParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetSearch(search *string) {
	o.Search = search
}

// WithV adds the v to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) WithV(v string) *ImagesPersonalAwsImagesParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the images personal aws images params
func (o *ImagesPersonalAwsImagesParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *ImagesPersonalAwsImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloudId
	if err := r.SetPathParam("cloudId", swag.FormatInt32(o.CloudID)); err != nil {
		return err
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
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
