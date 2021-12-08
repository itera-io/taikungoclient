// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new projects API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for projects API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectsCommit(params *ProjectsCommitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsCommitOK, error)

	ProjectsCreate(params *ProjectsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsCreateOK, error)

	ProjectsDelete(params *ProjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsDeleteOK, *ProjectsDeleteNoContent, error)

	ProjectsDeleteWholeProject(params *ProjectsDeleteWholeProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsDeleteWholeProjectOK, *ProjectsDeleteWholeProjectNoContent, error)

	ProjectsExtendLifeTime(params *ProjectsExtendLifeTimeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsExtendLifeTimeOK, error)

	ProjectsList(params *ProjectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsListOK, error)

	ProjectsLockManager(params *ProjectsLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsLockManagerOK, error)

	ProjectsLokiLogs(params *ProjectsLokiLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsLokiLogsOK, error)

	ProjectsPrometheusMetrics(params *ProjectsPrometheusMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsPrometheusMetricsOK, error)

	ProjectsRepair(params *ProjectsRepairParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsRepairOK, error)

	ProjectsUpgrade(params *ProjectsUpgradeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsUpgradeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProjectsCommit commits changes for the given project the changes will then be applied and the project will be updated the project must be in the r e a d y state
*/
func (a *Client) ProjectsCommit(params *ProjectsCommitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsCommitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsCommitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_Commit",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/commit/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsCommitReader{formats: a.formats},
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
	success, ok := result.(*ProjectsCommitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_Commit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsCreate onlies available for moderators and admin create a new project in the user s organization admin have the possibility to choose the organization for the newly created project
*/
func (a *Client) ProjectsCreate(params *ProjectsCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsCreateReader{formats: a.formats},
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
	success, ok := result.(*ProjectsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsDelete deletes the project the project must be empty no server and in r e a d y state
*/
func (a *Client) ProjectsDelete(params *ProjectsDeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsDeleteOK, *ProjectsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_Delete",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/delete",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectsDeleteOK:
		return value, nil, nil
	case *ProjectsDeleteNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for projects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsDeleteWholeProject deletes whole project by project Id
*/
func (a *Client) ProjectsDeleteWholeProject(params *ProjectsDeleteWholeProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsDeleteWholeProjectOK, *ProjectsDeleteWholeProjectNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsDeleteWholeProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_DeleteWholeProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/deletewholeproject",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsDeleteWholeProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectsDeleteWholeProjectOK:
		return value, nil, nil
	case *ProjectsDeleteWholeProjectNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for projects: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsExtendLifeTime extends life time of project
*/
func (a *Client) ProjectsExtendLifeTime(params *ProjectsExtendLifeTimeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsExtendLifeTimeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsExtendLifeTimeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_ExtendLifeTime",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/extend/lifetime",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsExtendLifeTimeReader{formats: a.formats},
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
	success, ok := result.(*ProjectsExtendLifeTimeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_ExtendLifeTime: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsList retrieves a list of projects
*/
func (a *Client) ProjectsList(params *ProjectsListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_List",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Projects",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsListReader{formats: a.formats},
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
	success, ok := result.(*ProjectsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_List: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsLockManager locks unlock project
*/
func (a *Client) ProjectsLockManager(params *ProjectsLockManagerParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsLockManagerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsLockManagerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_LockManager",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/lockmanager",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsLockManagerReader{formats: a.formats},
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
	success, ok := result.(*ProjectsLockManagerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_LockManager: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsLokiLogs retrieves loki logs
*/
func (a *Client) ProjectsLokiLogs(params *ProjectsLokiLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsLokiLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsLokiLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_LokiLogs",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/lokilogs",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsLokiLogsReader{formats: a.formats},
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
	success, ok := result.(*ProjectsLokiLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_LokiLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsPrometheusMetrics prometheus metrics data project
*/
func (a *Client) ProjectsPrometheusMetrics(params *ProjectsPrometheusMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsPrometheusMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsPrometheusMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_PrometheusMetrics",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/prometheusmetrics",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsPrometheusMetricsReader{formats: a.formats},
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
	success, ok := result.(*ProjectsPrometheusMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_PrometheusMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsRepair repairs project by Id
*/
func (a *Client) ProjectsRepair(params *ProjectsRepairParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsRepairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsRepairParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_Repair",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/repair/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsRepairReader{formats: a.formats},
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
	success, ok := result.(*ProjectsRepairOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_Repair: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ProjectsUpgrade upgrades the project s kubernetes to the next available version project must be r e a d y
*/
func (a *Client) ProjectsUpgrade(params *ProjectsUpgradeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ProjectsUpgradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectsUpgradeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Projects_Upgrade",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Projects/upgrade/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ProjectsUpgradeReader{formats: a.formats},
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
	success, ok := result.(*ProjectsUpgradeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Projects_Upgrade: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
