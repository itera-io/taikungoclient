/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@taikun.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package taikuncore

import (
	"encoding/json"
)

// checks if the ProjectActionVisibilityDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectActionVisibilityDto{}

// ProjectActionVisibilityDto struct for ProjectActionVisibilityDto
type ProjectActionVisibilityDto struct {
	Commit *ButtonStatusDto `json:"commit,omitempty"`
	Repair *ButtonStatusDto `json:"repair,omitempty"`
	Upgrade *ButtonStatusDto `json:"upgrade,omitempty"`
	Monitoring *ButtonStatusDto `json:"monitoring,omitempty"`
	EnableBackup *ButtonStatusDto `json:"enableBackup,omitempty"`
	DisableBackup *ButtonStatusDto `json:"disableBackup,omitempty"`
	EnableOpa *ButtonStatusDto `json:"enableOpa,omitempty"`
	DisableOpa *ButtonStatusDto `json:"disableOpa,omitempty"`
	EnableAutoscaler *ButtonStatusDto `json:"enableAutoscaler,omitempty"`
	DisableAutoscaler *ButtonStatusDto `json:"disableAutoscaler,omitempty"`
	VmRepair *ButtonStatusDto `json:"vmRepair,omitempty"`
	VmCommit *ButtonStatusDto `json:"vmCommit,omitempty"`
	Lock *ButtonStatusDto `json:"lock,omitempty"`
	Unlock *ButtonStatusDto `json:"unlock,omitempty"`
	EnableSpotWorker *ButtonStatusDto `json:"enableSpotWorker,omitempty"`
	DisableSpotWorker *ButtonStatusDto `json:"disableSpotWorker,omitempty"`
	EnableFullSpot *ButtonStatusDto `json:"enableFullSpot,omitempty"`
	DisableFullSpot *ButtonStatusDto `json:"disableFullSpot,omitempty"`
	EnableSpotVm *ButtonStatusDto `json:"enableSpotVm,omitempty"`
	DisableSpotVm *ButtonStatusDto `json:"disableSpotVm,omitempty"`
	AttachAlertingProfile *ButtonStatusDto `json:"attachAlertingProfile,omitempty"`
	DetachAlertingProfile *ButtonStatusDto `json:"detachAlertingProfile,omitempty"`
	EnableAi *ButtonStatusDto `json:"enableAi,omitempty"`
	DisableAi *ButtonStatusDto `json:"disableAi,omitempty"`
	AiAssistant *ButtonStatusDto `json:"aiAssistant,omitempty"`
	ProjectMaintenanceMode *ButtonStatusDto `json:"projectMaintenanceMode,omitempty"`
	AddServer *ButtonStatusDto `json:"addServer,omitempty"`
	AddVm *ButtonStatusDto `json:"addVm,omitempty"`
}

// NewProjectActionVisibilityDto instantiates a new ProjectActionVisibilityDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectActionVisibilityDto() *ProjectActionVisibilityDto {
	this := ProjectActionVisibilityDto{}
	return &this
}

// NewProjectActionVisibilityDtoWithDefaults instantiates a new ProjectActionVisibilityDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectActionVisibilityDtoWithDefaults() *ProjectActionVisibilityDto {
	this := ProjectActionVisibilityDto{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetCommit() ButtonStatusDto {
	if o == nil || IsNil(o.Commit) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetCommitOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given ButtonStatusDto and assigns it to the Commit field.
func (o *ProjectActionVisibilityDto) SetCommit(v ButtonStatusDto) {
	o.Commit = &v
}

// GetRepair returns the Repair field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetRepair() ButtonStatusDto {
	if o == nil || IsNil(o.Repair) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Repair
}

// GetRepairOk returns a tuple with the Repair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetRepairOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Repair) {
		return nil, false
	}
	return o.Repair, true
}

// HasRepair returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasRepair() bool {
	if o != nil && !IsNil(o.Repair) {
		return true
	}

	return false
}

// SetRepair gets a reference to the given ButtonStatusDto and assigns it to the Repair field.
func (o *ProjectActionVisibilityDto) SetRepair(v ButtonStatusDto) {
	o.Repair = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetUpgrade() ButtonStatusDto {
	if o == nil || IsNil(o.Upgrade) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetUpgradeOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Upgrade) {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasUpgrade() bool {
	if o != nil && !IsNil(o.Upgrade) {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given ButtonStatusDto and assigns it to the Upgrade field.
func (o *ProjectActionVisibilityDto) SetUpgrade(v ButtonStatusDto) {
	o.Upgrade = &v
}

// GetMonitoring returns the Monitoring field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetMonitoring() ButtonStatusDto {
	if o == nil || IsNil(o.Monitoring) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Monitoring
}

// GetMonitoringOk returns a tuple with the Monitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetMonitoringOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Monitoring) {
		return nil, false
	}
	return o.Monitoring, true
}

// HasMonitoring returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasMonitoring() bool {
	if o != nil && !IsNil(o.Monitoring) {
		return true
	}

	return false
}

// SetMonitoring gets a reference to the given ButtonStatusDto and assigns it to the Monitoring field.
func (o *ProjectActionVisibilityDto) SetMonitoring(v ButtonStatusDto) {
	o.Monitoring = &v
}

// GetEnableBackup returns the EnableBackup field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableBackup() ButtonStatusDto {
	if o == nil || IsNil(o.EnableBackup) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableBackup
}

// GetEnableBackupOk returns a tuple with the EnableBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableBackupOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableBackup) {
		return nil, false
	}
	return o.EnableBackup, true
}

// HasEnableBackup returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableBackup() bool {
	if o != nil && !IsNil(o.EnableBackup) {
		return true
	}

	return false
}

// SetEnableBackup gets a reference to the given ButtonStatusDto and assigns it to the EnableBackup field.
func (o *ProjectActionVisibilityDto) SetEnableBackup(v ButtonStatusDto) {
	o.EnableBackup = &v
}

// GetDisableBackup returns the DisableBackup field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableBackup() ButtonStatusDto {
	if o == nil || IsNil(o.DisableBackup) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableBackup
}

// GetDisableBackupOk returns a tuple with the DisableBackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableBackupOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableBackup) {
		return nil, false
	}
	return o.DisableBackup, true
}

// HasDisableBackup returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableBackup() bool {
	if o != nil && !IsNil(o.DisableBackup) {
		return true
	}

	return false
}

// SetDisableBackup gets a reference to the given ButtonStatusDto and assigns it to the DisableBackup field.
func (o *ProjectActionVisibilityDto) SetDisableBackup(v ButtonStatusDto) {
	o.DisableBackup = &v
}

// GetEnableOpa returns the EnableOpa field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableOpa() ButtonStatusDto {
	if o == nil || IsNil(o.EnableOpa) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableOpa
}

// GetEnableOpaOk returns a tuple with the EnableOpa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableOpaOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableOpa) {
		return nil, false
	}
	return o.EnableOpa, true
}

// HasEnableOpa returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableOpa() bool {
	if o != nil && !IsNil(o.EnableOpa) {
		return true
	}

	return false
}

// SetEnableOpa gets a reference to the given ButtonStatusDto and assigns it to the EnableOpa field.
func (o *ProjectActionVisibilityDto) SetEnableOpa(v ButtonStatusDto) {
	o.EnableOpa = &v
}

// GetDisableOpa returns the DisableOpa field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableOpa() ButtonStatusDto {
	if o == nil || IsNil(o.DisableOpa) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableOpa
}

// GetDisableOpaOk returns a tuple with the DisableOpa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableOpaOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableOpa) {
		return nil, false
	}
	return o.DisableOpa, true
}

// HasDisableOpa returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableOpa() bool {
	if o != nil && !IsNil(o.DisableOpa) {
		return true
	}

	return false
}

// SetDisableOpa gets a reference to the given ButtonStatusDto and assigns it to the DisableOpa field.
func (o *ProjectActionVisibilityDto) SetDisableOpa(v ButtonStatusDto) {
	o.DisableOpa = &v
}

// GetEnableAutoscaler returns the EnableAutoscaler field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableAutoscaler() ButtonStatusDto {
	if o == nil || IsNil(o.EnableAutoscaler) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableAutoscaler
}

// GetEnableAutoscalerOk returns a tuple with the EnableAutoscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableAutoscalerOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableAutoscaler) {
		return nil, false
	}
	return o.EnableAutoscaler, true
}

// HasEnableAutoscaler returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableAutoscaler() bool {
	if o != nil && !IsNil(o.EnableAutoscaler) {
		return true
	}

	return false
}

// SetEnableAutoscaler gets a reference to the given ButtonStatusDto and assigns it to the EnableAutoscaler field.
func (o *ProjectActionVisibilityDto) SetEnableAutoscaler(v ButtonStatusDto) {
	o.EnableAutoscaler = &v
}

// GetDisableAutoscaler returns the DisableAutoscaler field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableAutoscaler() ButtonStatusDto {
	if o == nil || IsNil(o.DisableAutoscaler) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableAutoscaler
}

// GetDisableAutoscalerOk returns a tuple with the DisableAutoscaler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableAutoscalerOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableAutoscaler) {
		return nil, false
	}
	return o.DisableAutoscaler, true
}

// HasDisableAutoscaler returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableAutoscaler() bool {
	if o != nil && !IsNil(o.DisableAutoscaler) {
		return true
	}

	return false
}

// SetDisableAutoscaler gets a reference to the given ButtonStatusDto and assigns it to the DisableAutoscaler field.
func (o *ProjectActionVisibilityDto) SetDisableAutoscaler(v ButtonStatusDto) {
	o.DisableAutoscaler = &v
}

// GetVmRepair returns the VmRepair field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetVmRepair() ButtonStatusDto {
	if o == nil || IsNil(o.VmRepair) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.VmRepair
}

// GetVmRepairOk returns a tuple with the VmRepair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetVmRepairOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.VmRepair) {
		return nil, false
	}
	return o.VmRepair, true
}

// HasVmRepair returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasVmRepair() bool {
	if o != nil && !IsNil(o.VmRepair) {
		return true
	}

	return false
}

// SetVmRepair gets a reference to the given ButtonStatusDto and assigns it to the VmRepair field.
func (o *ProjectActionVisibilityDto) SetVmRepair(v ButtonStatusDto) {
	o.VmRepair = &v
}

// GetVmCommit returns the VmCommit field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetVmCommit() ButtonStatusDto {
	if o == nil || IsNil(o.VmCommit) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.VmCommit
}

// GetVmCommitOk returns a tuple with the VmCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetVmCommitOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.VmCommit) {
		return nil, false
	}
	return o.VmCommit, true
}

// HasVmCommit returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasVmCommit() bool {
	if o != nil && !IsNil(o.VmCommit) {
		return true
	}

	return false
}

// SetVmCommit gets a reference to the given ButtonStatusDto and assigns it to the VmCommit field.
func (o *ProjectActionVisibilityDto) SetVmCommit(v ButtonStatusDto) {
	o.VmCommit = &v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetLock() ButtonStatusDto {
	if o == nil || IsNil(o.Lock) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetLockOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasLock() bool {
	if o != nil && !IsNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given ButtonStatusDto and assigns it to the Lock field.
func (o *ProjectActionVisibilityDto) SetLock(v ButtonStatusDto) {
	o.Lock = &v
}

// GetUnlock returns the Unlock field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetUnlock() ButtonStatusDto {
	if o == nil || IsNil(o.Unlock) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.Unlock
}

// GetUnlockOk returns a tuple with the Unlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetUnlockOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.Unlock) {
		return nil, false
	}
	return o.Unlock, true
}

// HasUnlock returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasUnlock() bool {
	if o != nil && !IsNil(o.Unlock) {
		return true
	}

	return false
}

// SetUnlock gets a reference to the given ButtonStatusDto and assigns it to the Unlock field.
func (o *ProjectActionVisibilityDto) SetUnlock(v ButtonStatusDto) {
	o.Unlock = &v
}

// GetEnableSpotWorker returns the EnableSpotWorker field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableSpotWorker() ButtonStatusDto {
	if o == nil || IsNil(o.EnableSpotWorker) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableSpotWorker
}

// GetEnableSpotWorkerOk returns a tuple with the EnableSpotWorker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableSpotWorkerOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableSpotWorker) {
		return nil, false
	}
	return o.EnableSpotWorker, true
}

// HasEnableSpotWorker returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableSpotWorker() bool {
	if o != nil && !IsNil(o.EnableSpotWorker) {
		return true
	}

	return false
}

// SetEnableSpotWorker gets a reference to the given ButtonStatusDto and assigns it to the EnableSpotWorker field.
func (o *ProjectActionVisibilityDto) SetEnableSpotWorker(v ButtonStatusDto) {
	o.EnableSpotWorker = &v
}

// GetDisableSpotWorker returns the DisableSpotWorker field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableSpotWorker() ButtonStatusDto {
	if o == nil || IsNil(o.DisableSpotWorker) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableSpotWorker
}

// GetDisableSpotWorkerOk returns a tuple with the DisableSpotWorker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableSpotWorkerOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableSpotWorker) {
		return nil, false
	}
	return o.DisableSpotWorker, true
}

// HasDisableSpotWorker returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableSpotWorker() bool {
	if o != nil && !IsNil(o.DisableSpotWorker) {
		return true
	}

	return false
}

// SetDisableSpotWorker gets a reference to the given ButtonStatusDto and assigns it to the DisableSpotWorker field.
func (o *ProjectActionVisibilityDto) SetDisableSpotWorker(v ButtonStatusDto) {
	o.DisableSpotWorker = &v
}

// GetEnableFullSpot returns the EnableFullSpot field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableFullSpot() ButtonStatusDto {
	if o == nil || IsNil(o.EnableFullSpot) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableFullSpot
}

// GetEnableFullSpotOk returns a tuple with the EnableFullSpot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableFullSpotOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableFullSpot) {
		return nil, false
	}
	return o.EnableFullSpot, true
}

// HasEnableFullSpot returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableFullSpot() bool {
	if o != nil && !IsNil(o.EnableFullSpot) {
		return true
	}

	return false
}

// SetEnableFullSpot gets a reference to the given ButtonStatusDto and assigns it to the EnableFullSpot field.
func (o *ProjectActionVisibilityDto) SetEnableFullSpot(v ButtonStatusDto) {
	o.EnableFullSpot = &v
}

// GetDisableFullSpot returns the DisableFullSpot field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableFullSpot() ButtonStatusDto {
	if o == nil || IsNil(o.DisableFullSpot) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableFullSpot
}

// GetDisableFullSpotOk returns a tuple with the DisableFullSpot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableFullSpotOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableFullSpot) {
		return nil, false
	}
	return o.DisableFullSpot, true
}

// HasDisableFullSpot returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableFullSpot() bool {
	if o != nil && !IsNil(o.DisableFullSpot) {
		return true
	}

	return false
}

// SetDisableFullSpot gets a reference to the given ButtonStatusDto and assigns it to the DisableFullSpot field.
func (o *ProjectActionVisibilityDto) SetDisableFullSpot(v ButtonStatusDto) {
	o.DisableFullSpot = &v
}

// GetEnableSpotVm returns the EnableSpotVm field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableSpotVm() ButtonStatusDto {
	if o == nil || IsNil(o.EnableSpotVm) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableSpotVm
}

// GetEnableSpotVmOk returns a tuple with the EnableSpotVm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableSpotVmOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableSpotVm) {
		return nil, false
	}
	return o.EnableSpotVm, true
}

// HasEnableSpotVm returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableSpotVm() bool {
	if o != nil && !IsNil(o.EnableSpotVm) {
		return true
	}

	return false
}

// SetEnableSpotVm gets a reference to the given ButtonStatusDto and assigns it to the EnableSpotVm field.
func (o *ProjectActionVisibilityDto) SetEnableSpotVm(v ButtonStatusDto) {
	o.EnableSpotVm = &v
}

// GetDisableSpotVm returns the DisableSpotVm field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableSpotVm() ButtonStatusDto {
	if o == nil || IsNil(o.DisableSpotVm) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableSpotVm
}

// GetDisableSpotVmOk returns a tuple with the DisableSpotVm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableSpotVmOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableSpotVm) {
		return nil, false
	}
	return o.DisableSpotVm, true
}

// HasDisableSpotVm returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableSpotVm() bool {
	if o != nil && !IsNil(o.DisableSpotVm) {
		return true
	}

	return false
}

// SetDisableSpotVm gets a reference to the given ButtonStatusDto and assigns it to the DisableSpotVm field.
func (o *ProjectActionVisibilityDto) SetDisableSpotVm(v ButtonStatusDto) {
	o.DisableSpotVm = &v
}

// GetAttachAlertingProfile returns the AttachAlertingProfile field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetAttachAlertingProfile() ButtonStatusDto {
	if o == nil || IsNil(o.AttachAlertingProfile) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.AttachAlertingProfile
}

// GetAttachAlertingProfileOk returns a tuple with the AttachAlertingProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetAttachAlertingProfileOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.AttachAlertingProfile) {
		return nil, false
	}
	return o.AttachAlertingProfile, true
}

// HasAttachAlertingProfile returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasAttachAlertingProfile() bool {
	if o != nil && !IsNil(o.AttachAlertingProfile) {
		return true
	}

	return false
}

// SetAttachAlertingProfile gets a reference to the given ButtonStatusDto and assigns it to the AttachAlertingProfile field.
func (o *ProjectActionVisibilityDto) SetAttachAlertingProfile(v ButtonStatusDto) {
	o.AttachAlertingProfile = &v
}

// GetDetachAlertingProfile returns the DetachAlertingProfile field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDetachAlertingProfile() ButtonStatusDto {
	if o == nil || IsNil(o.DetachAlertingProfile) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DetachAlertingProfile
}

// GetDetachAlertingProfileOk returns a tuple with the DetachAlertingProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDetachAlertingProfileOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DetachAlertingProfile) {
		return nil, false
	}
	return o.DetachAlertingProfile, true
}

// HasDetachAlertingProfile returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDetachAlertingProfile() bool {
	if o != nil && !IsNil(o.DetachAlertingProfile) {
		return true
	}

	return false
}

// SetDetachAlertingProfile gets a reference to the given ButtonStatusDto and assigns it to the DetachAlertingProfile field.
func (o *ProjectActionVisibilityDto) SetDetachAlertingProfile(v ButtonStatusDto) {
	o.DetachAlertingProfile = &v
}

// GetEnableAi returns the EnableAi field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetEnableAi() ButtonStatusDto {
	if o == nil || IsNil(o.EnableAi) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.EnableAi
}

// GetEnableAiOk returns a tuple with the EnableAi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetEnableAiOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.EnableAi) {
		return nil, false
	}
	return o.EnableAi, true
}

// HasEnableAi returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasEnableAi() bool {
	if o != nil && !IsNil(o.EnableAi) {
		return true
	}

	return false
}

// SetEnableAi gets a reference to the given ButtonStatusDto and assigns it to the EnableAi field.
func (o *ProjectActionVisibilityDto) SetEnableAi(v ButtonStatusDto) {
	o.EnableAi = &v
}

// GetDisableAi returns the DisableAi field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetDisableAi() ButtonStatusDto {
	if o == nil || IsNil(o.DisableAi) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.DisableAi
}

// GetDisableAiOk returns a tuple with the DisableAi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetDisableAiOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.DisableAi) {
		return nil, false
	}
	return o.DisableAi, true
}

// HasDisableAi returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasDisableAi() bool {
	if o != nil && !IsNil(o.DisableAi) {
		return true
	}

	return false
}

// SetDisableAi gets a reference to the given ButtonStatusDto and assigns it to the DisableAi field.
func (o *ProjectActionVisibilityDto) SetDisableAi(v ButtonStatusDto) {
	o.DisableAi = &v
}

// GetAiAssistant returns the AiAssistant field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetAiAssistant() ButtonStatusDto {
	if o == nil || IsNil(o.AiAssistant) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.AiAssistant
}

// GetAiAssistantOk returns a tuple with the AiAssistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetAiAssistantOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.AiAssistant) {
		return nil, false
	}
	return o.AiAssistant, true
}

// HasAiAssistant returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasAiAssistant() bool {
	if o != nil && !IsNil(o.AiAssistant) {
		return true
	}

	return false
}

// SetAiAssistant gets a reference to the given ButtonStatusDto and assigns it to the AiAssistant field.
func (o *ProjectActionVisibilityDto) SetAiAssistant(v ButtonStatusDto) {
	o.AiAssistant = &v
}

// GetProjectMaintenanceMode returns the ProjectMaintenanceMode field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetProjectMaintenanceMode() ButtonStatusDto {
	if o == nil || IsNil(o.ProjectMaintenanceMode) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.ProjectMaintenanceMode
}

// GetProjectMaintenanceModeOk returns a tuple with the ProjectMaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetProjectMaintenanceModeOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.ProjectMaintenanceMode) {
		return nil, false
	}
	return o.ProjectMaintenanceMode, true
}

// HasProjectMaintenanceMode returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasProjectMaintenanceMode() bool {
	if o != nil && !IsNil(o.ProjectMaintenanceMode) {
		return true
	}

	return false
}

// SetProjectMaintenanceMode gets a reference to the given ButtonStatusDto and assigns it to the ProjectMaintenanceMode field.
func (o *ProjectActionVisibilityDto) SetProjectMaintenanceMode(v ButtonStatusDto) {
	o.ProjectMaintenanceMode = &v
}

// GetAddServer returns the AddServer field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetAddServer() ButtonStatusDto {
	if o == nil || IsNil(o.AddServer) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.AddServer
}

// GetAddServerOk returns a tuple with the AddServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetAddServerOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.AddServer) {
		return nil, false
	}
	return o.AddServer, true
}

// HasAddServer returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasAddServer() bool {
	if o != nil && !IsNil(o.AddServer) {
		return true
	}

	return false
}

// SetAddServer gets a reference to the given ButtonStatusDto and assigns it to the AddServer field.
func (o *ProjectActionVisibilityDto) SetAddServer(v ButtonStatusDto) {
	o.AddServer = &v
}

// GetAddVm returns the AddVm field value if set, zero value otherwise.
func (o *ProjectActionVisibilityDto) GetAddVm() ButtonStatusDto {
	if o == nil || IsNil(o.AddVm) {
		var ret ButtonStatusDto
		return ret
	}
	return *o.AddVm
}

// GetAddVmOk returns a tuple with the AddVm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectActionVisibilityDto) GetAddVmOk() (*ButtonStatusDto, bool) {
	if o == nil || IsNil(o.AddVm) {
		return nil, false
	}
	return o.AddVm, true
}

// HasAddVm returns a boolean if a field has been set.
func (o *ProjectActionVisibilityDto) HasAddVm() bool {
	if o != nil && !IsNil(o.AddVm) {
		return true
	}

	return false
}

// SetAddVm gets a reference to the given ButtonStatusDto and assigns it to the AddVm field.
func (o *ProjectActionVisibilityDto) SetAddVm(v ButtonStatusDto) {
	o.AddVm = &v
}

func (o ProjectActionVisibilityDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectActionVisibilityDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.Repair) {
		toSerialize["repair"] = o.Repair
	}
	if !IsNil(o.Upgrade) {
		toSerialize["upgrade"] = o.Upgrade
	}
	if !IsNil(o.Monitoring) {
		toSerialize["monitoring"] = o.Monitoring
	}
	if !IsNil(o.EnableBackup) {
		toSerialize["enableBackup"] = o.EnableBackup
	}
	if !IsNil(o.DisableBackup) {
		toSerialize["disableBackup"] = o.DisableBackup
	}
	if !IsNil(o.EnableOpa) {
		toSerialize["enableOpa"] = o.EnableOpa
	}
	if !IsNil(o.DisableOpa) {
		toSerialize["disableOpa"] = o.DisableOpa
	}
	if !IsNil(o.EnableAutoscaler) {
		toSerialize["enableAutoscaler"] = o.EnableAutoscaler
	}
	if !IsNil(o.DisableAutoscaler) {
		toSerialize["disableAutoscaler"] = o.DisableAutoscaler
	}
	if !IsNil(o.VmRepair) {
		toSerialize["vmRepair"] = o.VmRepair
	}
	if !IsNil(o.VmCommit) {
		toSerialize["vmCommit"] = o.VmCommit
	}
	if !IsNil(o.Lock) {
		toSerialize["lock"] = o.Lock
	}
	if !IsNil(o.Unlock) {
		toSerialize["unlock"] = o.Unlock
	}
	if !IsNil(o.EnableSpotWorker) {
		toSerialize["enableSpotWorker"] = o.EnableSpotWorker
	}
	if !IsNil(o.DisableSpotWorker) {
		toSerialize["disableSpotWorker"] = o.DisableSpotWorker
	}
	if !IsNil(o.EnableFullSpot) {
		toSerialize["enableFullSpot"] = o.EnableFullSpot
	}
	if !IsNil(o.DisableFullSpot) {
		toSerialize["disableFullSpot"] = o.DisableFullSpot
	}
	if !IsNil(o.EnableSpotVm) {
		toSerialize["enableSpotVm"] = o.EnableSpotVm
	}
	if !IsNil(o.DisableSpotVm) {
		toSerialize["disableSpotVm"] = o.DisableSpotVm
	}
	if !IsNil(o.AttachAlertingProfile) {
		toSerialize["attachAlertingProfile"] = o.AttachAlertingProfile
	}
	if !IsNil(o.DetachAlertingProfile) {
		toSerialize["detachAlertingProfile"] = o.DetachAlertingProfile
	}
	if !IsNil(o.EnableAi) {
		toSerialize["enableAi"] = o.EnableAi
	}
	if !IsNil(o.DisableAi) {
		toSerialize["disableAi"] = o.DisableAi
	}
	if !IsNil(o.AiAssistant) {
		toSerialize["aiAssistant"] = o.AiAssistant
	}
	if !IsNil(o.ProjectMaintenanceMode) {
		toSerialize["projectMaintenanceMode"] = o.ProjectMaintenanceMode
	}
	if !IsNil(o.AddServer) {
		toSerialize["addServer"] = o.AddServer
	}
	if !IsNil(o.AddVm) {
		toSerialize["addVm"] = o.AddVm
	}
	return toSerialize, nil
}

type NullableProjectActionVisibilityDto struct {
	value *ProjectActionVisibilityDto
	isSet bool
}

func (v NullableProjectActionVisibilityDto) Get() *ProjectActionVisibilityDto {
	return v.value
}

func (v *NullableProjectActionVisibilityDto) Set(val *ProjectActionVisibilityDto) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectActionVisibilityDto) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectActionVisibilityDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectActionVisibilityDto(val *ProjectActionVisibilityDto) *NullableProjectActionVisibilityDto {
	return &NullableProjectActionVisibilityDto{value: val, isSet: true}
}

func (v NullableProjectActionVisibilityDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectActionVisibilityDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


