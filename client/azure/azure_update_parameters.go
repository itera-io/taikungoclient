// Code generated by go-swagger; DO NOT EDIT.

package azure

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

// NewAzureUpdateParams creates a new AzureUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureUpdateParams() *AzureUpdateParams {
	return &AzureUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureUpdateParamsWithTimeout creates a new AzureUpdateParams object
// with the ability to set a timeout on a request.
func NewAzureUpdateParamsWithTimeout(timeout time.Duration) *AzureUpdateParams {
	return &AzureUpdateParams{
		timeout: timeout,
	}
}

// NewAzureUpdateParamsWithContext creates a new AzureUpdateParams object
// with the ability to set a context for a request.
func NewAzureUpdateParamsWithContext(ctx context.Context) *AzureUpdateParams {
	return &AzureUpdateParams{
		Context: ctx,
	}
}

// NewAzureUpdateParamsWithHTTPClient creates a new AzureUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureUpdateParamsWithHTTPClient(client *http.Client) *AzureUpdateParams {
	return &AzureUpdateParams{
		HTTPClient: client,
	}
}

/* AzureUpdateParams contains all the parameters to send to the API endpoint
   for the azure update operation.

   Typically these are written to a http.Request.
*/
type AzureUpdateParams struct {

	// Body.
	Body *models.UpdateAzureCommand

	// V.
	V string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureUpdateParams) WithDefaults() *AzureUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the azure update params
func (o *AzureUpdateParams) WithTimeout(timeout time.Duration) *AzureUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure update params
func (o *AzureUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure update params
func (o *AzureUpdateParams) WithContext(ctx context.Context) *AzureUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure update params
func (o *AzureUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure update params
func (o *AzureUpdateParams) WithHTTPClient(client *http.Client) *AzureUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure update params
func (o *AzureUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the azure update params
func (o *AzureUpdateParams) WithBody(body *models.UpdateAzureCommand) *AzureUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the azure update params
func (o *AzureUpdateParams) SetBody(body *models.UpdateAzureCommand) {
	o.Body = body
}

// WithV adds the v to the azure update params
func (o *AzureUpdateParams) WithV(v string) *AzureUpdateParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the azure update params
func (o *AzureUpdateParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *AzureUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
