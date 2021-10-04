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
)

// NewKubernetesDeletePodParams creates a new KubernetesDeletePodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKubernetesDeletePodParams() *KubernetesDeletePodParams {
	return &KubernetesDeletePodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKubernetesDeletePodParamsWithTimeout creates a new KubernetesDeletePodParams object
// with the ability to set a timeout on a request.
func NewKubernetesDeletePodParamsWithTimeout(timeout time.Duration) *KubernetesDeletePodParams {
	return &KubernetesDeletePodParams{
		timeout: timeout,
	}
}

// NewKubernetesDeletePodParamsWithContext creates a new KubernetesDeletePodParams object
// with the ability to set a context for a request.
func NewKubernetesDeletePodParamsWithContext(ctx context.Context) *KubernetesDeletePodParams {
	return &KubernetesDeletePodParams{
		Context: ctx,
	}
}

// NewKubernetesDeletePodParamsWithHTTPClient creates a new KubernetesDeletePodParams object
// with the ability to set a custom HTTPClient for a request.
func NewKubernetesDeletePodParamsWithHTTPClient(client *http.Client) *KubernetesDeletePodParams {
	return &KubernetesDeletePodParams{
		HTTPClient: client,
	}
}

/* KubernetesDeletePodParams contains all the parameters to send to the API endpoint
   for the kubernetes delete pod operation.

   Typically these are written to a http.Request.
*/
type KubernetesDeletePodParams struct {

	// MetadataName.
	MetadataName string

	// PodNamespace.
	PodNamespace string

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

// WithDefaults hydrates default values in the kubernetes delete pod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDeletePodParams) WithDefaults() *KubernetesDeletePodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kubernetes delete pod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KubernetesDeletePodParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithTimeout(timeout time.Duration) *KubernetesDeletePodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithContext(ctx context.Context) *KubernetesDeletePodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithHTTPClient(client *http.Client) *KubernetesDeletePodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetadataName adds the metadataName to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithMetadataName(metadataName string) *KubernetesDeletePodParams {
	o.SetMetadataName(metadataName)
	return o
}

// SetMetadataName adds the metadataName to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetMetadataName(metadataName string) {
	o.MetadataName = metadataName
}

// WithPodNamespace adds the podNamespace to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithPodNamespace(podNamespace string) *KubernetesDeletePodParams {
	o.SetPodNamespace(podNamespace)
	return o
}

// SetPodNamespace adds the podNamespace to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetPodNamespace(podNamespace string) {
	o.PodNamespace = podNamespace
}

// WithProjectID adds the projectID to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithProjectID(projectID int32) *KubernetesDeletePodParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetProjectID(projectID int32) {
	o.ProjectID = projectID
}

// WithV adds the v to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) WithV(v string) *KubernetesDeletePodParams {
	o.SetV(v)
	return o
}

// SetV adds the v to the kubernetes delete pod params
func (o *KubernetesDeletePodParams) SetV(v string) {
	o.V = v
}

// WriteToRequest writes these params to a swagger request
func (o *KubernetesDeletePodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metadataName
	if err := r.SetPathParam("metadataName", o.MetadataName); err != nil {
		return err
	}

	// path param podNamespace
	if err := r.SetPathParam("podNamespace", o.PodNamespace); err != nil {
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
