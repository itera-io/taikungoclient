// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BackupClearProject(params *BackupClearProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupClearProjectOK, error)

	BackupCreate(params *BackupCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupCreateOK, error)

	BackupDeleteBackup(params *BackupDeleteBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteBackupOK, error)

	BackupDeleteBackupLocation(params *BackupDeleteBackupLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteBackupLocationOK, error)

	BackupDeleteRestore(params *BackupDeleteRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteRestoreOK, error)

	BackupDeleteSchedule(params *BackupDeleteScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteScheduleOK, error)

	BackupDescribeBackup(params *BackupDescribeBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeBackupOK, error)

	BackupDescribeRestore(params *BackupDescribeRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeRestoreOK, error)

	BackupDescribeSchedule(params *BackupDescribeScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeScheduleOK, error)

	BackupDisableBackup(params *BackupDisableBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDisableBackupOK, error)

	BackupEnableBackup(params *BackupEnableBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupEnableBackupOK, error)

	BackupImportBackupStorage(params *BackupImportBackupStorageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupImportBackupStorageOK, error)

	BackupListAllBackupStorages(params *BackupListAllBackupStoragesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllBackupStoragesOK, error)

	BackupListAllBackups(params *BackupListAllBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllBackupsOK, error)

	BackupListAllDeleteBackupRequests(params *BackupListAllDeleteBackupRequestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllDeleteBackupRequestsOK, error)

	BackupListAllRestores(params *BackupListAllRestoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllRestoresOK, error)

	BackupListAllSchedules(params *BackupListAllSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllSchedulesOK, error)

	BackupRestoreBackup(params *BackupRestoreBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupRestoreBackupOK, error)

	BackupScheduleByName(params *BackupScheduleByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupScheduleByNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BackupClearProject deletes unfinished backup for project
*/
func (a *Client) BackupClearProject(params *BackupClearProjectParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupClearProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupClearProjectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ClearProject",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/clear/project",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupClearProjectReader{formats: a.formats},
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
	success, ok := result.(*BackupClearProjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ClearProject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupCreate adds backup policy
*/
func (a *Client) BackupCreate(params *BackupCreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_Create",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/create",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupCreateReader{formats: a.formats},
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
	success, ok := result.(*BackupCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_Create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDeleteBackup removes policy backup
*/
func (a *Client) BackupDeleteBackup(params *BackupDeleteBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDeleteBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DeleteBackup",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/delete/backup",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDeleteBackupReader{formats: a.formats},
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
	success, ok := result.(*BackupDeleteBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DeleteBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDeleteBackupLocation removes backup location from project
*/
func (a *Client) BackupDeleteBackupLocation(params *BackupDeleteBackupLocationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteBackupLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDeleteBackupLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DeleteBackupLocation",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/delete/location",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDeleteBackupLocationReader{formats: a.formats},
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
	success, ok := result.(*BackupDeleteBackupLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DeleteBackupLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDeleteRestore removes policy restore
*/
func (a *Client) BackupDeleteRestore(params *BackupDeleteRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDeleteRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DeleteRestore",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/delete/restore",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDeleteRestoreReader{formats: a.formats},
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
	success, ok := result.(*BackupDeleteRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DeleteRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDeleteSchedule removes policy schedule
*/
func (a *Client) BackupDeleteSchedule(params *BackupDeleteScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDeleteScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDeleteScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DeleteSchedule",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/delete/schedule",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDeleteScheduleReader{formats: a.formats},
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
	success, ok := result.(*BackupDeleteScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DeleteSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDescribeBackup gets backup info by name
*/
func (a *Client) BackupDescribeBackup(params *BackupDescribeBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDescribeBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DescribeBackup",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/describe/backup/{projectId}/{name}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDescribeBackupReader{formats: a.formats},
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
	success, ok := result.(*BackupDescribeBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DescribeBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDescribeRestore gets restore info by name
*/
func (a *Client) BackupDescribeRestore(params *BackupDescribeRestoreParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDescribeRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DescribeRestore",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/describe/restore/{projectId}/{name}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDescribeRestoreReader{formats: a.formats},
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
	success, ok := result.(*BackupDescribeRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DescribeRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDescribeSchedule gets schedule info by name
*/
func (a *Client) BackupDescribeSchedule(params *BackupDescribeScheduleParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDescribeScheduleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDescribeScheduleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DescribeSchedule",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/describe/schedule/{projectId}/{name}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDescribeScheduleReader{formats: a.formats},
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
	success, ok := result.(*BackupDescribeScheduleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DescribeSchedule: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupDisableBackup disables backup by the project Id
*/
func (a *Client) BackupDisableBackup(params *BackupDisableBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupDisableBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDisableBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_DisableBackup",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/disablebackup",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupDisableBackupReader{formats: a.formats},
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
	success, ok := result.(*BackupDisableBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_DisableBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupEnableBackup enables backup by the project Id
*/
func (a *Client) BackupEnableBackup(params *BackupEnableBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupEnableBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupEnableBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_EnableBackup",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/enablebackup",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupEnableBackupReader{formats: a.formats},
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
	success, ok := result.(*BackupEnableBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_EnableBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupImportBackupStorage imports backup storage from source project to target project
*/
func (a *Client) BackupImportBackupStorage(params *BackupImportBackupStorageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupImportBackupStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupImportBackupStorageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ImportBackupStorage",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/location",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupImportBackupStorageReader{formats: a.formats},
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
	success, ok := result.(*BackupImportBackupStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ImportBackupStorage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupListAllBackupStorages lists all backup locations
*/
func (a *Client) BackupListAllBackupStorages(params *BackupListAllBackupStoragesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllBackupStoragesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListAllBackupStoragesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ListAllBackupStorages",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/location/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListAllBackupStoragesReader{formats: a.formats},
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
	success, ok := result.(*BackupListAllBackupStoragesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ListAllBackupStorages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupListAllBackups lists all backups
*/
func (a *Client) BackupListAllBackups(params *BackupListAllBackupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllBackupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListAllBackupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ListAllBackups",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/backups/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListAllBackupsReader{formats: a.formats},
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
	success, ok := result.(*BackupListAllBackupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ListAllBackups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupListAllDeleteBackupRequests lists all delete backup requests
*/
func (a *Client) BackupListAllDeleteBackupRequests(params *BackupListAllDeleteBackupRequestsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllDeleteBackupRequestsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListAllDeleteBackupRequestsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ListAllDeleteBackupRequests",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/delelete-requests/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListAllDeleteBackupRequestsReader{formats: a.formats},
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
	success, ok := result.(*BackupListAllDeleteBackupRequestsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ListAllDeleteBackupRequests: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupListAllRestores lists all restores
*/
func (a *Client) BackupListAllRestores(params *BackupListAllRestoresParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllRestoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListAllRestoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ListAllRestores",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/restores/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListAllRestoresReader{formats: a.formats},
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
	success, ok := result.(*BackupListAllRestoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ListAllRestores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupListAllSchedules lists all schedules
*/
func (a *Client) BackupListAllSchedules(params *BackupListAllSchedulesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupListAllSchedulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupListAllSchedulesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ListAllSchedules",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/schedules/{projectId}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupListAllSchedulesReader{formats: a.formats},
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
	success, ok := result.(*BackupListAllSchedulesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ListAllSchedules: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupRestoreBackup restores backup
*/
func (a *Client) BackupRestoreBackup(params *BackupRestoreBackupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupRestoreBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupRestoreBackupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_RestoreBackup",
		Method:             "POST",
		PathPattern:        "/api/v{v}/Backup/restore",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupRestoreBackupReader{formats: a.formats},
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
	success, ok := result.(*BackupRestoreBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_RestoreBackup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BackupScheduleByName gets schedule info by name
*/
func (a *Client) BackupScheduleByName(params *BackupScheduleByNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BackupScheduleByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupScheduleByNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Backup_ScheduleByName",
		Method:             "GET",
		PathPattern:        "/api/v{v}/Backup/schedule/{projectId}/{name}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BackupScheduleByNameReader{formats: a.formats},
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
	success, ok := result.(*BackupScheduleByNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Backup_ScheduleByName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
