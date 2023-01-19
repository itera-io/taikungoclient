// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ImagesAwsImagesAsPost(params *ImagesAwsImagesAsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesAwsImagesAsPostOK, error)

	ImagesAzureImages(params *ImagesAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesAzureImagesOK, error)

	ImagesBindImagesToProject(params *ImagesBindImagesToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesBindImagesToProjectOK, error)

	ImagesCommonAwsImages(params *ImagesCommonAwsImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonAwsImagesOK, error)

	ImagesCommonAzureImages(params *ImagesCommonAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonAzureImagesOK, error)

	ImagesCommonGoogleImages(params *ImagesCommonGoogleImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonGoogleImagesOK, error)

	ImagesGetImageDetailsByID(params *ImagesGetImageDetailsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGetImageDetailsByIDOK, error)

	ImagesGetSelectedImagesForProject(params *ImagesGetSelectedImagesForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGetSelectedImagesForProjectOK, error)

	ImagesGoogleImages(params *ImagesGoogleImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGoogleImagesOK, error)

	ImagesOpenstackImages(params *ImagesOpenstackImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesOpenstackImagesOK, error)

	ImagesPersonalAwsImages(params *ImagesPersonalAwsImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesPersonalAwsImagesOK, error)

	ImagesPersonalAzureImages(params *ImagesPersonalAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesPersonalAzureImagesOK, error)

	ImagesUnbindImagesFromProject(params *ImagesUnbindImagesFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesUnbindImagesFromProjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ImagesAwsImagesAsPost retrieves aws images
*/
func (a *Client) ImagesAwsImagesAsPost(params *ImagesAwsImagesAsPostParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesAwsImagesAsPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesAwsImagesAsPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_AwsImagesAsPost",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Images/aws",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesAwsImagesAsPostReader{formats: a.formats},
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
	success, ok := result.(*ImagesAwsImagesAsPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_AwsImagesAsPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesAzureImages retrieves azure images
*/
func (a *Client) ImagesAzureImages(params *ImagesAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesAzureImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesAzureImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_AzureImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/azure/{cloudId}/{publisherName}/{offer}/{sku}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesAzureImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesAzureImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_AzureImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesBindImagesToProject binds images to project
*/
func (a *Client) ImagesBindImagesToProject(params *ImagesBindImagesToProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesBindImagesToProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesBindImagesToProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_BindImagesToProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Images/bind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesBindImagesToProjectReader{formats: a.formats},
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
	success, ok := result.(*ImagesBindImagesToProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_BindImagesToProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesCommonAwsImages commonlies used aws images
*/
func (a *Client) ImagesCommonAwsImages(params *ImagesCommonAwsImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonAwsImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesCommonAwsImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_CommonAwsImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/aws/common/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesCommonAwsImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesCommonAwsImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_CommonAwsImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesCommonAzureImages commonlies used azure images
*/
func (a *Client) ImagesCommonAzureImages(params *ImagesCommonAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonAzureImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesCommonAzureImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_CommonAzureImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/azure/common/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesCommonAzureImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesCommonAzureImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_CommonAzureImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesCommonGoogleImages commonlies used google images
*/
func (a *Client) ImagesCommonGoogleImages(params *ImagesCommonGoogleImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesCommonGoogleImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesCommonGoogleImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_CommonGoogleImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/google/common/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesCommonGoogleImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesCommonGoogleImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_CommonGoogleImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesGetImageDetailsByID gets image details
*/
func (a *Client) ImagesGetImageDetailsByID(params *ImagesGetImageDetailsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGetImageDetailsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesGetImageDetailsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_GetImageDetailsById",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Images/details",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesGetImageDetailsByIDReader{formats: a.formats},
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
	success, ok := result.(*ImagesGetImageDetailsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_GetImageDetailsById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesGetSelectedImagesForProject retrieves selected images for projects
*/
func (a *Client) ImagesGetSelectedImagesForProject(params *ImagesGetSelectedImagesForProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGetSelectedImagesForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesGetSelectedImagesForProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_GetSelectedImagesForProject",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/projects/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesGetSelectedImagesForProjectReader{formats: a.formats},
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
	success, ok := result.(*ImagesGetSelectedImagesForProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_GetSelectedImagesForProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesGoogleImages retrieves google images
*/
func (a *Client) ImagesGoogleImages(params *ImagesGoogleImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesGoogleImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesGoogleImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_GoogleImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/google/{cloudId}/{type}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesGoogleImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesGoogleImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_GoogleImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesOpenstackImages retrieves openstack images
*/
func (a *Client) ImagesOpenstackImages(params *ImagesOpenstackImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesOpenstackImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesOpenstackImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_OpenstackImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/openstack/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesOpenstackImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesOpenstackImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_OpenstackImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesPersonalAwsImages personals aws images
*/
func (a *Client) ImagesPersonalAwsImages(params *ImagesPersonalAwsImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesPersonalAwsImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesPersonalAwsImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_PersonalAwsImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/aws/personal/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesPersonalAwsImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesPersonalAwsImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_PersonalAwsImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesPersonalAzureImages personals azure images
*/
func (a *Client) ImagesPersonalAzureImages(params *ImagesPersonalAzureImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesPersonalAzureImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesPersonalAzureImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_PersonalAzureImages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Images/azure/personal/{cloudId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesPersonalAzureImagesReader{formats: a.formats},
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
	success, ok := result.(*ImagesPersonalAzureImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_PersonalAzureImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ImagesUnbindImagesFromProject unbinds images from project
*/
func (a *Client) ImagesUnbindImagesFromProject(params *ImagesUnbindImagesFromProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ImagesUnbindImagesFromProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImagesUnbindImagesFromProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Images_UnbindImagesFromProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Images/unbind",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImagesUnbindImagesFromProjectReader{formats: a.formats},
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
	success, ok := result.(*ImagesUnbindImagesFromProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Images_UnbindImagesFromProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
