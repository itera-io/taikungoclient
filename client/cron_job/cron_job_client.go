// Code generated by go-swagger; DO NOT EDIT.

package cron_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cron job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cron job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CronJobCreateKeyPool(params *CronJobCreateKeyPoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobCreateKeyPoolOK, error)

	CronJobDeleteAwsSpotInstances(params *CronJobDeleteAwsSpotInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteAwsSpotInstancesOK, error)

	CronJobDeleteExpiredAlerts(params *CronJobDeleteExpiredAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredAlertsOK, error)

	CronJobDeleteExpiredEvents(params *CronJobDeleteExpiredEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredEventsOK, error)

	CronJobDeleteExpiredHistoryLogs(params *CronJobDeleteExpiredHistoryLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredHistoryLogsOK, error)

	CronJobDeleteExpiredRefreshTokens(params *CronJobDeleteExpiredRefreshTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredRefreshTokensOK, error)

	CronJobDeleteExpiredRequests(params *CronJobDeleteExpiredRequestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredRequestsOK, error)

	CronJobDeleteExpiredServers(params *CronJobDeleteExpiredServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredServersOK, error)

	CronJobDeleteKubeConfigs(params *CronJobDeleteKubeConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteKubeConfigsOK, error)

	CronJobDeletePendingOrganizations(params *CronJobDeletePendingOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeletePendingOrganizationsOK, error)

	CronJobDeleteRedundantProjectActions(params *CronJobDeleteRedundantProjectActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteRedundantProjectActionsOK, error)

	CronJobFetchAzurePrices(params *CronJobFetchAzurePricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchAzurePricesOK, error)

	CronJobFetchKubernetesAlerts(params *CronJobFetchKubernetesAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchKubernetesAlertsOK, error)

	CronJobFetchKubernetesOverview(params *CronJobFetchKubernetesOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchKubernetesOverviewOK, error)

	CronJobFetchOrganizationDetails(params *CronJobFetchOrganizationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchOrganizationDetailsOK, error)

	CronJobRemindAlerts(params *CronJobRemindAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobRemindAlertsOK, error)

	CronJobSyncBackupCredentials(params *CronJobSyncBackupCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobSyncBackupCredentialsOK, error)

	CronJobSyncOpaProfiles(params *CronJobSyncOpaProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobSyncOpaProfilesOK, error)

	CronJobUpdateProjectQuotaMessage(params *CronJobUpdateProjectQuotaMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobUpdateProjectQuotaMessageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CronJobCreateKeyPool creates key pool
*/
func (a *Client) CronJobCreateKeyPool(params *CronJobCreateKeyPoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobCreateKeyPoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobCreateKeyPoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_CreateKeyPool",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/create-key-pool",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobCreateKeyPoolReader{formats: a.formats},
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
	success, ok := result.(*CronJobCreateKeyPoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_CreateKeyPool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteAwsSpotInstances deletes removed aws spot instances
*/
func (a *Client) CronJobDeleteAwsSpotInstances(params *CronJobDeleteAwsSpotInstancesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteAwsSpotInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteAwsSpotInstancesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteAwsSpotInstances",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/delete-aws-spot-instances",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteAwsSpotInstancesReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteAwsSpotInstancesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteAwsSpotInstances: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredAlerts deletes expired alerts
*/
func (a *Client) CronJobDeleteExpiredAlerts(params *CronJobDeleteExpiredAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredAlerts",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/alerts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredAlertsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredEvents deletes expired events
*/
func (a *Client) CronJobDeleteExpiredEvents(params *CronJobDeleteExpiredEventsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredEvents",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/events",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredEventsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredHistoryLogs deletes expired history logs
*/
func (a *Client) CronJobDeleteExpiredHistoryLogs(params *CronJobDeleteExpiredHistoryLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredHistoryLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredHistoryLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredHistoryLogs",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/history-logs",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredHistoryLogsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredHistoryLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredHistoryLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredRefreshTokens deletes expired refresh tokens
*/
func (a *Client) CronJobDeleteExpiredRefreshTokens(params *CronJobDeleteExpiredRefreshTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredRefreshTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredRefreshTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredRefreshTokens",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/refresh-tokens",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredRefreshTokensReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredRefreshTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredRefreshTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredRequests deletes expired requests
*/
func (a *Client) CronJobDeleteExpiredRequests(params *CronJobDeleteExpiredRequestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredRequestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredRequests",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/taikun-requests",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredRequestsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredRequestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredRequests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteExpiredServers deletes expired servers
*/
func (a *Client) CronJobDeleteExpiredServers(params *CronJobDeleteExpiredServersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteExpiredServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteExpiredServersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteExpiredServers",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/servers",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteExpiredServersReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteExpiredServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteExpiredServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteKubeConfigs removes deleted user s kube configs
*/
func (a *Client) CronJobDeleteKubeConfigs(params *CronJobDeleteKubeConfigsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteKubeConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteKubeConfigsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteKubeConfigs",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/delete-kube-configs",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteKubeConfigsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteKubeConfigsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteKubeConfigs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeletePendingOrganizations deletes pending organizations
*/
func (a *Client) CronJobDeletePendingOrganizations(params *CronJobDeletePendingOrganizationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeletePendingOrganizationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeletePendingOrganizationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeletePendingOrganizations",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/organizations",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeletePendingOrganizationsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeletePendingOrganizationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeletePendingOrganizations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobDeleteRedundantProjectActions deletes useless project actions
*/
func (a *Client) CronJobDeleteRedundantProjectActions(params *CronJobDeleteRedundantProjectActionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobDeleteRedundantProjectActionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobDeleteRedundantProjectActionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_DeleteRedundantProjectActions",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/project-actions",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobDeleteRedundantProjectActionsReader{formats: a.formats},
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
	success, ok := result.(*CronJobDeleteRedundantProjectActionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_DeleteRedundantProjectActions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobFetchAzurePrices fetches azure prices
*/
func (a *Client) CronJobFetchAzurePrices(params *CronJobFetchAzurePricesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchAzurePricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobFetchAzurePricesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_FetchAzurePrices",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/fetch-azure-prices",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobFetchAzurePricesReader{formats: a.formats},
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
	success, ok := result.(*CronJobFetchAzurePricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_FetchAzurePrices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobFetchKubernetesAlerts fetches kubernetes alerts for organization
*/
func (a *Client) CronJobFetchKubernetesAlerts(params *CronJobFetchKubernetesAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchKubernetesAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobFetchKubernetesAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_FetchKubernetesAlerts",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/fetch-kubernetes-alerts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobFetchKubernetesAlertsReader{formats: a.formats},
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
	success, ok := result.(*CronJobFetchKubernetesAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_FetchKubernetesAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobFetchKubernetesOverview fetches kubernetes overview for organization
*/
func (a *Client) CronJobFetchKubernetesOverview(params *CronJobFetchKubernetesOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchKubernetesOverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobFetchKubernetesOverviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_FetchKubernetesOverview",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/fetch-kubernetes-overview",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobFetchKubernetesOverviewReader{formats: a.formats},
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
	success, ok := result.(*CronJobFetchKubernetesOverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_FetchKubernetesOverview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobFetchOrganizationDetails fetches details of organizations
*/
func (a *Client) CronJobFetchOrganizationDetails(params *CronJobFetchOrganizationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobFetchOrganizationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobFetchOrganizationDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_FetchOrganizationDetails",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/fetch-organization-details",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobFetchOrganizationDetailsReader{formats: a.formats},
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
	success, ok := result.(*CronJobFetchOrganizationDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_FetchOrganizationDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobRemindAlerts reminds alerts
*/
func (a *Client) CronJobRemindAlerts(params *CronJobRemindAlertsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobRemindAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobRemindAlertsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_RemindAlerts",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/remind-alerts",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobRemindAlertsReader{formats: a.formats},
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
	success, ok := result.(*CronJobRemindAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_RemindAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobSyncBackupCredentials syncs backup credentials
*/
func (a *Client) CronJobSyncBackupCredentials(params *CronJobSyncBackupCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobSyncBackupCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobSyncBackupCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_SyncBackupCredentials",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/sync-backup-credentials",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobSyncBackupCredentialsReader{formats: a.formats},
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
	success, ok := result.(*CronJobSyncBackupCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_SyncBackupCredentials: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobSyncOpaProfiles syncs opa profiles
*/
func (a *Client) CronJobSyncOpaProfiles(params *CronJobSyncOpaProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobSyncOpaProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobSyncOpaProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_SyncOpaProfiles",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/sync-opa-profiles",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobSyncOpaProfilesReader{formats: a.formats},
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
	success, ok := result.(*CronJobSyncOpaProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_SyncOpaProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CronJobUpdateProjectQuotaMessage updates project quota message
*/
func (a *Client) CronJobUpdateProjectQuotaMessage(params *CronJobUpdateProjectQuotaMessageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CronJobUpdateProjectQuotaMessageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCronJobUpdateProjectQuotaMessageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CronJob_UpdateProjectQuotaMessage",
		Method:             "POST",
		PathPattern:        "/api/v{v}/CronJob/project-quota-message",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CronJobUpdateProjectQuotaMessageReader{formats: a.formats},
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
	success, ok := result.(*CronJobUpdateProjectQuotaMessageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CronJob_UpdateProjectQuotaMessage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
