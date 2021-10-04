// Code generated by go-swagger; DO NOT EDIT.

package kube_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new kube config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kube config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	KubeConfigCreate(params *KubeConfigCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigCreateOK, error)

	KubeConfigDelete(params *KubeConfigDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigDeleteOK, error)

	KubeConfigDownload(params *KubeConfigDownloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigDownloadOK, error)

	KubeConfigList(params *KubeConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  KubeConfigCreate creates kube config
*/
func (a *Client) KubeConfigCreate(params *KubeConfigCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubeConfigCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KubeConfig_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/KubeConfig",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KubeConfigCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubeConfigCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KubeConfig_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KubeConfigDelete deletes kube config
*/
func (a *Client) KubeConfigDelete(params *KubeConfigDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubeConfigDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KubeConfig_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/KubeConfig/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KubeConfigDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubeConfigDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KubeConfig_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KubeConfigDownload downloads kube config file for user by project Id
*/
func (a *Client) KubeConfigDownload(params *KubeConfigDownloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigDownloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubeConfigDownloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KubeConfig_Download",
		Method:             "POST",
		PathPattern:        "/api/v{v}/KubeConfig/download",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KubeConfigDownloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubeConfigDownloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KubeConfig_Download: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  KubeConfigList retrieves a list of kube configs for project it s possible to filter and select kube configs by project
*/
func (a *Client) KubeConfigList(params *KubeConfigListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*KubeConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubeConfigListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "KubeConfig_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/KubeConfig",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &KubeConfigListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubeConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for KubeConfig_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
