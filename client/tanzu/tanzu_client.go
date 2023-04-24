// Code generated by go-swagger; DO NOT EDIT.

package tanzu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tanzu API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tanzu API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	TanzuCreate(params *TanzuCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuCreateOK, error)

	TanzuKubernetesVersions(params *TanzuKubernetesVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuKubernetesVersionsOK, error)

	TanzuList(params *TanzuListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuListOK, error)

	TanzuUpdate(params *TanzuUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
TanzuCreate adds tanzu credentials
*/
func (a *Client) TanzuCreate(params *TanzuCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTanzuCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Tanzu_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Tanzu/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TanzuCreateReader{formats: a.formats},
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
	success, ok := result.(*TanzuCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tanzu_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TanzuKubernetesVersions tanzus available k8s version list
*/
func (a *Client) TanzuKubernetesVersions(params *TanzuKubernetesVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuKubernetesVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTanzuKubernetesVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Tanzu_KubernetesVersions",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Tanzu/kubernetes-versions/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TanzuKubernetesVersionsReader{formats: a.formats},
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
	success, ok := result.(*TanzuKubernetesVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tanzu_KubernetesVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TanzuList retrieves list of tanzu cloud credentials
*/
func (a *Client) TanzuList(params *TanzuListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTanzuListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Tanzu_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Tanzu/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TanzuListReader{formats: a.formats},
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
	success, ok := result.(*TanzuListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tanzu_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
TanzuUpdate updates tanzu credentials
*/
func (a *Client) TanzuUpdate(params *TanzuUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TanzuUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTanzuUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Tanzu_Update",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Tanzu/update",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TanzuUpdateReader{formats: a.formats},
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
	success, ok := result.(*TanzuUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Tanzu_Update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
