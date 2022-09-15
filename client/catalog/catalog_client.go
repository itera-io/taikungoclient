// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogAppLockManager(params *CatalogAppLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAppLockManagerOK, error)

	CatalogAvailablePackageDetails(params *CatalogAvailablePackageDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAvailablePackageDetailsOK, error)

	CatalogAvailableVersions(params *CatalogAvailableVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAvailableVersionsOK, error)

	CatalogBindProjects(params *CatalogBindProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogBindProjectsOK, error)

	CatalogCatalogAppDetails(params *CatalogCatalogAppDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppDetailsOK, error)

	CatalogCatalogAppParamsDetails(params *CatalogCatalogAppParamsDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppParamsDetailsOK, error)

	CatalogCatalogAppValue(params *CatalogCatalogAppValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppValueOK, error)

	CatalogCatalogAppValueAutocomplete(params *CatalogCatalogAppValueAutocompleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppValueAutocompleteOK, error)

	CatalogCatalogDropdown(params *CatalogCatalogDropdownParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogDropdownOK, error)

	CatalogCatalogList(params *CatalogCatalogListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogListOK, error)

	CatalogCreateCatalog(params *CatalogCreateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCreateCatalogOK, error)

	CatalogCreateCatalogApp(params *CatalogCreateCatalogAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCreateCatalogAppOK, error)

	CatalogDeleteCatalog(params *CatalogDeleteCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogDeleteCatalogOK, error)

	CatalogDeleteCatalogApp(params *CatalogDeleteCatalogAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogDeleteCatalogAppOK, error)

	CatalogEditCatalog(params *CatalogEditCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogOK, error)

	CatalogEditCatalogAppParams(params *CatalogEditCatalogAppParamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogAppParamsOK, error)

	CatalogEditCatalogAppVersion(params *CatalogEditCatalogAppVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogAppVersionOK, error)

	CatalogList(params *CatalogListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogListOK, error)

	CatalogLockManager(params *CatalogLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogLockManagerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CatalogAppLockManager locks unlock catalog app
*/
func (a *Client) CatalogAppLockManager(params *CatalogAppLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAppLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAppLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_AppLockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/app-lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogAppLockManagerReader{formats: a.formats},
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
	success, ok := result.(*CatalogAppLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_AppLockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogAvailablePackageDetails availables helm package details
*/
func (a *Client) CatalogAvailablePackageDetails(params *CatalogAvailablePackageDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAvailablePackageDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAvailablePackageDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_AvailablePackageDetails",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/available/{repoName}/{packageName}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogAvailablePackageDetailsReader{formats: a.formats},
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
	success, ok := result.(*CatalogAvailablePackageDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_AvailablePackageDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogAvailableVersions gets available versions for catalog app
*/
func (a *Client) CatalogAvailableVersions(params *CatalogAvailableVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogAvailableVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogAvailableVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_AvailableVersions",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/available/versions",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogAvailableVersionsReader{formats: a.formats},
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
	success, ok := result.(*CatalogAvailableVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_AvailableVersions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogBindProjects binds unbind projects to catalog
*/
func (a *Client) CatalogBindProjects(params *CatalogBindProjectsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogBindProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogBindProjectsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_BindProjects",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/bind-project",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogBindProjectsReader{formats: a.formats},
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
	success, ok := result.(*CatalogBindProjectsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_BindProjects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogAppDetails catalogs app details
*/
func (a *Client) CatalogCatalogAppDetails(params *CatalogCatalogAppDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogAppDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogAppDetails",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/catalog-app/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogAppDetailsReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogAppDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogAppDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogAppParamsDetails catalogs app params details
*/
func (a *Client) CatalogCatalogAppParamsDetails(params *CatalogCatalogAppParamsDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppParamsDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogAppParamsDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogAppParamsDetails",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/catalog-app-params/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogAppParamsDetailsReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogAppParamsDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogAppParamsDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogAppValue gets catalog app yaml value
*/
func (a *Client) CatalogCatalogAppValue(params *CatalogCatalogAppValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppValueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogAppValueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogAppValue",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/package-value",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogAppValueReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogAppValueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogAppValue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogAppValueAutocomplete gets autocomplete dictionary for catalog app value
*/
func (a *Client) CatalogCatalogAppValueAutocomplete(params *CatalogCatalogAppValueAutocompleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogAppValueAutocompleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogAppValueAutocompleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogAppValueAutocomplete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/package-value-autocomplete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogAppValueAutocompleteReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogAppValueAutocompleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogAppValueAutocomplete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogDropdown catalogs dropdown list for organization
*/
func (a *Client) CatalogCatalogDropdown(params *CatalogCatalogDropdownParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogDropdownOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogDropdownParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogDropdown",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/dropdown-list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogDropdownReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogDropdownOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogDropdown: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCatalogList catalogs list for organization
*/
func (a *Client) CatalogCatalogList(params *CatalogCatalogListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCatalogListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCatalogListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CatalogList",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/list",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCatalogListReader{formats: a.formats},
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
	success, ok := result.(*CatalogCatalogListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CatalogList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCreateCatalog creates catalog for organization
*/
func (a *Client) CatalogCreateCatalog(params *CatalogCreateCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCreateCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCreateCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CreateCatalog",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/create-catalog",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCreateCatalogReader{formats: a.formats},
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
	success, ok := result.(*CatalogCreateCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CreateCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogCreateCatalogApp adds application to catalog
*/
func (a *Client) CatalogCreateCatalogApp(params *CatalogCreateCatalogAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogCreateCatalogAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogCreateCatalogAppParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_CreateCatalogApp",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/app-to-catalog",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogCreateCatalogAppReader{formats: a.formats},
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
	success, ok := result.(*CatalogCreateCatalogAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_CreateCatalogApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogDeleteCatalog deletes catalog
*/
func (a *Client) CatalogDeleteCatalog(params *CatalogDeleteCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogDeleteCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogDeleteCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_DeleteCatalog",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/Catalog/delete/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogDeleteCatalogReader{formats: a.formats},
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
	success, ok := result.(*CatalogDeleteCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_DeleteCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogDeleteCatalogApp deletes catalog app by catalog app Id
*/
func (a *Client) CatalogDeleteCatalogApp(params *CatalogDeleteCatalogAppParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogDeleteCatalogAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogDeleteCatalogAppParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_DeleteCatalogApp",
		Method:             "DELETE",
		PathPattern:        "/api/v{v}/Catalog/delete-app/{id}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogDeleteCatalogAppReader{formats: a.formats},
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
	success, ok := result.(*CatalogDeleteCatalogAppOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_DeleteCatalogApp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogEditCatalog edits catalog for organization
*/
func (a *Client) CatalogEditCatalog(params *CatalogEditCatalogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogEditCatalogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_EditCatalog",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Catalog/edit-catalog",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogEditCatalogReader{formats: a.formats},
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
	success, ok := result.(*CatalogEditCatalogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_EditCatalog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogEditCatalogAppParams edits catalog app params
*/
func (a *Client) CatalogEditCatalogAppParams(params *CatalogEditCatalogAppParamsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogAppParamsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogEditCatalogAppParamsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_EditCatalogAppParams",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Catalog/edit-catalogapp-params",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogEditCatalogAppParamsReader{formats: a.formats},
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
	success, ok := result.(*CatalogEditCatalogAppParamsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_EditCatalogAppParams: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogEditCatalogAppVersion edits catalog app version
*/
func (a *Client) CatalogEditCatalogAppVersion(params *CatalogEditCatalogAppVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogEditCatalogAppVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogEditCatalogAppVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_EditCatalogAppVersion",
		Method:             "PUT",
		PathPattern:        "/api/v{v}/Catalog/edit-catalogapp-version",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogEditCatalogAppVersionReader{formats: a.formats},
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
	success, ok := result.(*CatalogEditCatalogAppVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_EditCatalogAppVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogList retrieves all available packages
*/
func (a *Client) CatalogList(params *CatalogListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Catalog/available/packages",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogListReader{formats: a.formats},
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
	success, ok := result.(*CatalogListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CatalogLockManager locks unlock catalog
*/
func (a *Client) CatalogLockManager(params *CatalogLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CatalogLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Catalog_LockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Catalog/lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogLockManagerReader{formats: a.formats},
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
	success, ok := result.(*CatalogLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Catalog_LockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
