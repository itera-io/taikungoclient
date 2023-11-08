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

// checks if the StandaloneVmsListForDetailsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandaloneVmsListForDetailsDto{}

// StandaloneVmsListForDetailsDto struct for StandaloneVmsListForDetailsDto
type StandaloneVmsListForDetailsDto struct {
	Id                  *int32                          `json:"id,omitempty"`
	Name                NullableString                  `json:"name,omitempty"`
	ImageName           NullableString                  `json:"imageName,omitempty"`
	ImageId             NullableString                  `json:"imageId,omitempty"`
	Status              NullableString                  `json:"status,omitempty"`
	CloudInit           NullableString                  `json:"cloudInit,omitempty"`
	VolumeType          NullableString                  `json:"volumeType,omitempty"`
	VolumeSize          *int64                          `json:"volumeSize,omitempty"`
	CreatedAt           NullableString                  `json:"createdAt,omitempty"`
	CreatedBy           NullableString                  `json:"createdBy,omitempty"`
	LastModified        NullableString                  `json:"lastModified,omitempty"`
	LastModifiedBy      NullableString                  `json:"lastModifiedBy,omitempty"`
	SshPublicKey        NullableString                  `json:"sshPublicKey,omitempty"`
	CurrentFlavor       NullableString                  `json:"currentFlavor,omitempty"`
	TargetFlavor        NullableString                  `json:"targetFlavor,omitempty"`
	PublicIpEnabled     *bool                           `json:"publicIpEnabled,omitempty"`
	PublicIp            NullableString                  `json:"publicIp,omitempty"`
	IpAddress           NullableString                  `json:"ipAddress,omitempty"`
	SpotPrice           NullableString                  `json:"spotPrice,omitempty"`
	SpotInstance        *bool                           `json:"spotInstance,omitempty"`
	ActionButtons       *StandaloneVisibilityDto        `json:"actionButtons,omitempty"`
	IsWindows           *bool                           `json:"isWindows,omitempty"`
	Disks               []StandAloneVmDiskForDetailsDto `json:"disks,omitempty"`
	StandAloneMetaDatas []StandAloneMetaDataDtoForVm    `json:"standAloneMetaDatas,omitempty"`
	Profile             *StandAloneProfileForDetailsDto `json:"profile,omitempty"`
}

// NewStandaloneVmsListForDetailsDto instantiates a new StandaloneVmsListForDetailsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandaloneVmsListForDetailsDto() *StandaloneVmsListForDetailsDto {
	this := StandaloneVmsListForDetailsDto{}
	return &this
}

// NewStandaloneVmsListForDetailsDtoWithDefaults instantiates a new StandaloneVmsListForDetailsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandaloneVmsListForDetailsDtoWithDefaults() *StandaloneVmsListForDetailsDto {
	this := StandaloneVmsListForDetailsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StandaloneVmsListForDetailsDto) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *StandaloneVmsListForDetailsDto) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetName() {
	o.Name.Unset()
}

// GetImageName returns the ImageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetImageName() string {
	if o == nil || IsNil(o.ImageName.Get()) {
		var ret string
		return ret
	}
	return *o.ImageName.Get()
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageName.Get(), o.ImageName.IsSet()
}

// HasImageName returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasImageName() bool {
	if o != nil && o.ImageName.IsSet() {
		return true
	}

	return false
}

// SetImageName gets a reference to the given NullableString and assigns it to the ImageName field.
func (o *StandaloneVmsListForDetailsDto) SetImageName(v string) {
	o.ImageName.Set(&v)
}

// SetImageNameNil sets the value for ImageName to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetImageNameNil() {
	o.ImageName.Set(nil)
}

// UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetImageName() {
	o.ImageName.Unset()
}

// GetImageId returns the ImageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetImageId() string {
	if o == nil || IsNil(o.ImageId.Get()) {
		var ret string
		return ret
	}
	return *o.ImageId.Get()
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetImageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageId.Get(), o.ImageId.IsSet()
}

// HasImageId returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasImageId() bool {
	if o != nil && o.ImageId.IsSet() {
		return true
	}

	return false
}

// SetImageId gets a reference to the given NullableString and assigns it to the ImageId field.
func (o *StandaloneVmsListForDetailsDto) SetImageId(v string) {
	o.ImageId.Set(&v)
}

// SetImageIdNil sets the value for ImageId to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetImageIdNil() {
	o.ImageId.Set(nil)
}

// UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetImageId() {
	o.ImageId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *StandaloneVmsListForDetailsDto) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetStatus() {
	o.Status.Unset()
}

// GetCloudInit returns the CloudInit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetCloudInit() string {
	if o == nil || IsNil(o.CloudInit.Get()) {
		var ret string
		return ret
	}
	return *o.CloudInit.Get()
}

// GetCloudInitOk returns a tuple with the CloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetCloudInitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudInit.Get(), o.CloudInit.IsSet()
}

// HasCloudInit returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasCloudInit() bool {
	if o != nil && o.CloudInit.IsSet() {
		return true
	}

	return false
}

// SetCloudInit gets a reference to the given NullableString and assigns it to the CloudInit field.
func (o *StandaloneVmsListForDetailsDto) SetCloudInit(v string) {
	o.CloudInit.Set(&v)
}

// SetCloudInitNil sets the value for CloudInit to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetCloudInitNil() {
	o.CloudInit.Set(nil)
}

// UnsetCloudInit ensures that no value is present for CloudInit, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetCloudInit() {
	o.CloudInit.Unset()
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetVolumeType() string {
	if o == nil || IsNil(o.VolumeType.Get()) {
		var ret string
		return ret
	}
	return *o.VolumeType.Get()
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetVolumeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeType.Get(), o.VolumeType.IsSet()
}

// HasVolumeType returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasVolumeType() bool {
	if o != nil && o.VolumeType.IsSet() {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given NullableString and assigns it to the VolumeType field.
func (o *StandaloneVmsListForDetailsDto) SetVolumeType(v string) {
	o.VolumeType.Set(&v)
}

// SetVolumeTypeNil sets the value for VolumeType to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetVolumeTypeNil() {
	o.VolumeType.Set(nil)
}

// UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetVolumeType() {
	o.VolumeType.Unset()
}

// GetVolumeSize returns the VolumeSize field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetVolumeSize() int64 {
	if o == nil || IsNil(o.VolumeSize) {
		var ret int64
		return ret
	}
	return *o.VolumeSize
}

// GetVolumeSizeOk returns a tuple with the VolumeSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetVolumeSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.VolumeSize) {
		return nil, false
	}
	return o.VolumeSize, true
}

// HasVolumeSize returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasVolumeSize() bool {
	if o != nil && !IsNil(o.VolumeSize) {
		return true
	}

	return false
}

// SetVolumeSize gets a reference to the given int64 and assigns it to the VolumeSize field.
func (o *StandaloneVmsListForDetailsDto) SetVolumeSize(v int64) {
	o.VolumeSize = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *StandaloneVmsListForDetailsDto) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *StandaloneVmsListForDetailsDto) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}

// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetLastModified returns the LastModified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetLastModified() string {
	if o == nil || IsNil(o.LastModified.Get()) {
		var ret string
		return ret
	}
	return *o.LastModified.Get()
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetLastModifiedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModified.Get(), o.LastModified.IsSet()
}

// HasLastModified returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasLastModified() bool {
	if o != nil && o.LastModified.IsSet() {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given NullableString and assigns it to the LastModified field.
func (o *StandaloneVmsListForDetailsDto) SetLastModified(v string) {
	o.LastModified.Set(&v)
}

// SetLastModifiedNil sets the value for LastModified to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetLastModifiedNil() {
	o.LastModified.Set(nil)
}

// UnsetLastModified ensures that no value is present for LastModified, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetLastModified() {
	o.LastModified.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetLastModifiedBy() string {
	if o == nil || IsNil(o.LastModifiedBy.Get()) {
		var ret string
		return ret
	}
	return *o.LastModifiedBy.Get()
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModifiedBy.Get(), o.LastModifiedBy.IsSet()
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given NullableString and assigns it to the LastModifiedBy field.
func (o *StandaloneVmsListForDetailsDto) SetLastModifiedBy(v string) {
	o.LastModifiedBy.Set(&v)
}

// SetLastModifiedByNil sets the value for LastModifiedBy to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetLastModifiedByNil() {
	o.LastModifiedBy.Set(nil)
}

// UnsetLastModifiedBy ensures that no value is present for LastModifiedBy, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetLastModifiedBy() {
	o.LastModifiedBy.Unset()
}

// GetSshPublicKey returns the SshPublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetSshPublicKey() string {
	if o == nil || IsNil(o.SshPublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.SshPublicKey.Get()
}

// GetSshPublicKeyOk returns a tuple with the SshPublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetSshPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SshPublicKey.Get(), o.SshPublicKey.IsSet()
}

// HasSshPublicKey returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasSshPublicKey() bool {
	if o != nil && o.SshPublicKey.IsSet() {
		return true
	}

	return false
}

// SetSshPublicKey gets a reference to the given NullableString and assigns it to the SshPublicKey field.
func (o *StandaloneVmsListForDetailsDto) SetSshPublicKey(v string) {
	o.SshPublicKey.Set(&v)
}

// SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetSshPublicKeyNil() {
	o.SshPublicKey.Set(nil)
}

// UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetSshPublicKey() {
	o.SshPublicKey.Unset()
}

// GetCurrentFlavor returns the CurrentFlavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavor() string {
	if o == nil || IsNil(o.CurrentFlavor.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentFlavor.Get()
}

// GetCurrentFlavorOk returns a tuple with the CurrentFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetCurrentFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentFlavor.Get(), o.CurrentFlavor.IsSet()
}

// HasCurrentFlavor returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasCurrentFlavor() bool {
	if o != nil && o.CurrentFlavor.IsSet() {
		return true
	}

	return false
}

// SetCurrentFlavor gets a reference to the given NullableString and assigns it to the CurrentFlavor field.
func (o *StandaloneVmsListForDetailsDto) SetCurrentFlavor(v string) {
	o.CurrentFlavor.Set(&v)
}

// SetCurrentFlavorNil sets the value for CurrentFlavor to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetCurrentFlavorNil() {
	o.CurrentFlavor.Set(nil)
}

// UnsetCurrentFlavor ensures that no value is present for CurrentFlavor, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetCurrentFlavor() {
	o.CurrentFlavor.Unset()
}

// GetTargetFlavor returns the TargetFlavor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetTargetFlavor() string {
	if o == nil || IsNil(o.TargetFlavor.Get()) {
		var ret string
		return ret
	}
	return *o.TargetFlavor.Get()
}

// GetTargetFlavorOk returns a tuple with the TargetFlavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetTargetFlavorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetFlavor.Get(), o.TargetFlavor.IsSet()
}

// HasTargetFlavor returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasTargetFlavor() bool {
	if o != nil && o.TargetFlavor.IsSet() {
		return true
	}

	return false
}

// SetTargetFlavor gets a reference to the given NullableString and assigns it to the TargetFlavor field.
func (o *StandaloneVmsListForDetailsDto) SetTargetFlavor(v string) {
	o.TargetFlavor.Set(&v)
}

// SetTargetFlavorNil sets the value for TargetFlavor to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetTargetFlavorNil() {
	o.TargetFlavor.Set(nil)
}

// UnsetTargetFlavor ensures that no value is present for TargetFlavor, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetTargetFlavor() {
	o.TargetFlavor.Unset()
}

// GetPublicIpEnabled returns the PublicIpEnabled field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabled() bool {
	if o == nil || IsNil(o.PublicIpEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicIpEnabled
}

// GetPublicIpEnabledOk returns a tuple with the PublicIpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetPublicIpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicIpEnabled) {
		return nil, false
	}
	return o.PublicIpEnabled, true
}

// HasPublicIpEnabled returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasPublicIpEnabled() bool {
	if o != nil && !IsNil(o.PublicIpEnabled) {
		return true
	}

	return false
}

// SetPublicIpEnabled gets a reference to the given bool and assigns it to the PublicIpEnabled field.
func (o *StandaloneVmsListForDetailsDto) SetPublicIpEnabled(v bool) {
	o.PublicIpEnabled = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetPublicIp() string {
	if o == nil || IsNil(o.PublicIp.Get()) {
		var ret string
		return ret
	}
	return *o.PublicIp.Get()
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetPublicIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicIp.Get(), o.PublicIp.IsSet()
}

// HasPublicIp returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasPublicIp() bool {
	if o != nil && o.PublicIp.IsSet() {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given NullableString and assigns it to the PublicIp field.
func (o *StandaloneVmsListForDetailsDto) SetPublicIp(v string) {
	o.PublicIp.Set(&v)
}

// SetPublicIpNil sets the value for PublicIp to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetPublicIpNil() {
	o.PublicIp.Set(nil)
}

// UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetPublicIp() {
	o.PublicIp.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress.Get()) {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *StandaloneVmsListForDetailsDto) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}

// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetSpotPrice returns the SpotPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetSpotPrice() string {
	if o == nil || IsNil(o.SpotPrice.Get()) {
		var ret string
		return ret
	}
	return *o.SpotPrice.Get()
}

// GetSpotPriceOk returns a tuple with the SpotPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetSpotPriceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpotPrice.Get(), o.SpotPrice.IsSet()
}

// HasSpotPrice returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasSpotPrice() bool {
	if o != nil && o.SpotPrice.IsSet() {
		return true
	}

	return false
}

// SetSpotPrice gets a reference to the given NullableString and assigns it to the SpotPrice field.
func (o *StandaloneVmsListForDetailsDto) SetSpotPrice(v string) {
	o.SpotPrice.Set(&v)
}

// SetSpotPriceNil sets the value for SpotPrice to be an explicit nil
func (o *StandaloneVmsListForDetailsDto) SetSpotPriceNil() {
	o.SpotPrice.Set(nil)
}

// UnsetSpotPrice ensures that no value is present for SpotPrice, not even an explicit nil
func (o *StandaloneVmsListForDetailsDto) UnsetSpotPrice() {
	o.SpotPrice.Unset()
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetSpotInstance() bool {
	if o == nil || IsNil(o.SpotInstance) {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.SpotInstance) {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasSpotInstance() bool {
	if o != nil && !IsNil(o.SpotInstance) {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *StandaloneVmsListForDetailsDto) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetActionButtons returns the ActionButtons field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetActionButtons() StandaloneVisibilityDto {
	if o == nil || IsNil(o.ActionButtons) {
		var ret StandaloneVisibilityDto
		return ret
	}
	return *o.ActionButtons
}

// GetActionButtonsOk returns a tuple with the ActionButtons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetActionButtonsOk() (*StandaloneVisibilityDto, bool) {
	if o == nil || IsNil(o.ActionButtons) {
		return nil, false
	}
	return o.ActionButtons, true
}

// HasActionButtons returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasActionButtons() bool {
	if o != nil && !IsNil(o.ActionButtons) {
		return true
	}

	return false
}

// SetActionButtons gets a reference to the given StandaloneVisibilityDto and assigns it to the ActionButtons field.
func (o *StandaloneVmsListForDetailsDto) SetActionButtons(v StandaloneVisibilityDto) {
	o.ActionButtons = &v
}

// GetIsWindows returns the IsWindows field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetIsWindows() bool {
	if o == nil || IsNil(o.IsWindows) {
		var ret bool
		return ret
	}
	return *o.IsWindows
}

// GetIsWindowsOk returns a tuple with the IsWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetIsWindowsOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWindows) {
		return nil, false
	}
	return o.IsWindows, true
}

// HasIsWindows returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasIsWindows() bool {
	if o != nil && !IsNil(o.IsWindows) {
		return true
	}

	return false
}

// SetIsWindows gets a reference to the given bool and assigns it to the IsWindows field.
func (o *StandaloneVmsListForDetailsDto) SetIsWindows(v bool) {
	o.IsWindows = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetDisks() []StandAloneVmDiskForDetailsDto {
	if o == nil {
		var ret []StandAloneVmDiskForDetailsDto
		return ret
	}
	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetDisksOk() ([]StandAloneVmDiskForDetailsDto, bool) {
	if o == nil || IsNil(o.Disks) {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasDisks() bool {
	if o != nil && IsNil(o.Disks) {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []StandAloneVmDiskForDetailsDto and assigns it to the Disks field.
func (o *StandaloneVmsListForDetailsDto) SetDisks(v []StandAloneVmDiskForDetailsDto) {
	o.Disks = v
}

// GetStandAloneMetaDatas returns the StandAloneMetaDatas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatas() []StandAloneMetaDataDtoForVm {
	if o == nil {
		var ret []StandAloneMetaDataDtoForVm
		return ret
	}
	return o.StandAloneMetaDatas
}

// GetStandAloneMetaDatasOk returns a tuple with the StandAloneMetaDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StandaloneVmsListForDetailsDto) GetStandAloneMetaDatasOk() ([]StandAloneMetaDataDtoForVm, bool) {
	if o == nil || IsNil(o.StandAloneMetaDatas) {
		return nil, false
	}
	return o.StandAloneMetaDatas, true
}

// HasStandAloneMetaDatas returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasStandAloneMetaDatas() bool {
	if o != nil && IsNil(o.StandAloneMetaDatas) {
		return true
	}

	return false
}

// SetStandAloneMetaDatas gets a reference to the given []StandAloneMetaDataDtoForVm and assigns it to the StandAloneMetaDatas field.
func (o *StandaloneVmsListForDetailsDto) SetStandAloneMetaDatas(v []StandAloneMetaDataDtoForVm) {
	o.StandAloneMetaDatas = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *StandaloneVmsListForDetailsDto) GetProfile() StandAloneProfileForDetailsDto {
	if o == nil || IsNil(o.Profile) {
		var ret StandAloneProfileForDetailsDto
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandaloneVmsListForDetailsDto) GetProfileOk() (*StandAloneProfileForDetailsDto, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *StandaloneVmsListForDetailsDto) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given StandAloneProfileForDetailsDto and assigns it to the Profile field.
func (o *StandaloneVmsListForDetailsDto) SetProfile(v StandAloneProfileForDetailsDto) {
	o.Profile = &v
}

func (o StandaloneVmsListForDetailsDto) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StandaloneVmsListForDetailsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ImageName.IsSet() {
		toSerialize["imageName"] = o.ImageName.Get()
	}
	if o.ImageId.IsSet() {
		toSerialize["imageId"] = o.ImageId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.CloudInit.IsSet() {
		toSerialize["cloudInit"] = o.CloudInit.Get()
	}
	if o.VolumeType.IsSet() {
		toSerialize["volumeType"] = o.VolumeType.Get()
	}
	if !IsNil(o.VolumeSize) {
		toSerialize["volumeSize"] = o.VolumeSize
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.LastModified.IsSet() {
		toSerialize["lastModified"] = o.LastModified.Get()
	}
	if o.LastModifiedBy.IsSet() {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy.Get()
	}
	if o.SshPublicKey.IsSet() {
		toSerialize["sshPublicKey"] = o.SshPublicKey.Get()
	}
	if o.CurrentFlavor.IsSet() {
		toSerialize["currentFlavor"] = o.CurrentFlavor.Get()
	}
	if o.TargetFlavor.IsSet() {
		toSerialize["targetFlavor"] = o.TargetFlavor.Get()
	}
	if !IsNil(o.PublicIpEnabled) {
		toSerialize["publicIpEnabled"] = o.PublicIpEnabled
	}
	if o.PublicIp.IsSet() {
		toSerialize["publicIp"] = o.PublicIp.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ipAddress"] = o.IpAddress.Get()
	}
	if o.SpotPrice.IsSet() {
		toSerialize["spotPrice"] = o.SpotPrice.Get()
	}
	if !IsNil(o.SpotInstance) {
		toSerialize["spotInstance"] = o.SpotInstance
	}
	if !IsNil(o.ActionButtons) {
		toSerialize["actionButtons"] = o.ActionButtons
	}
	if !IsNil(o.IsWindows) {
		toSerialize["isWindows"] = o.IsWindows
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	if o.StandAloneMetaDatas != nil {
		toSerialize["standAloneMetaDatas"] = o.StandAloneMetaDatas
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableStandaloneVmsListForDetailsDto struct {
	value *StandaloneVmsListForDetailsDto
	isSet bool
}

func (v NullableStandaloneVmsListForDetailsDto) Get() *StandaloneVmsListForDetailsDto {
	return v.value
}

func (v *NullableStandaloneVmsListForDetailsDto) Set(val *StandaloneVmsListForDetailsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStandaloneVmsListForDetailsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStandaloneVmsListForDetailsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandaloneVmsListForDetailsDto(val *StandaloneVmsListForDetailsDto) *NullableStandaloneVmsListForDetailsDto {
	return &NullableStandaloneVmsListForDetailsDto{value: val, isSet: true}
}

func (v NullableStandaloneVmsListForDetailsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandaloneVmsListForDetailsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
