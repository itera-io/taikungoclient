// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectDetailsForServersDto project details for servers dto
//
// swagger:model ProjectDetailsForServersDto
type ProjectDetailsForServersDto struct {

	// access Ip
	AccessIP string `json:"accessIp,omitempty"`

	// access profile Id
	AccessProfileID int32 `json:"accessProfileId,omitempty"`

	// access profile name
	AccessProfileName string `json:"accessProfileName,omitempty"`

	// access profile revision
	AccessProfileRevision int32 `json:"accessProfileRevision,omitempty"`

	// alerting profile Id
	AlertingProfileID int32 `json:"alertingProfileId,omitempty"`

	// alerting profile name
	AlertingProfileName string `json:"alertingProfileName,omitempty"`

	// alerts total count
	AlertsTotalCount int32 `json:"alertsTotalCount,omitempty"`

	// allow full spot kubernetes
	AllowFullSpotKubernetes bool `json:"allowFullSpotKubernetes"`

	// allow spot v ms
	AllowSpotVMs bool `json:"allowSpotVMs"`

	// allow spot workers
	AllowSpotWorkers bool `json:"allowSpotWorkers"`

	// autoscaling group name
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty"`

	// availability zones
	AvailabilityZones []string `json:"availabilityZones"`

	// bastion
	Bastion int32 `json:"bastion,omitempty"`

	// certification expired at
	CertificationExpiredAt string `json:"certificationExpiredAt,omitempty"`

	// cloud credential revision
	CloudCredentialRevision int32 `json:"cloudCredentialRevision,omitempty"`

	// cloud Id
	CloudID int32 `json:"cloudId,omitempty"`

	// cloud name
	CloudName string `json:"cloudName,omitempty"`

	// cloud provider message
	CloudProviderMessage string `json:"cloudProviderMessage,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// cpu limit
	CPULimit int64 `json:"cpuLimit,omitempty"`

	// delete on expiration
	DeleteOnExpiration bool `json:"deleteOnExpiration"`

	// disk size
	DiskSize float64 `json:"diskSize,omitempty"`

	// disk size limit
	DiskSizeLimit int64 `json:"diskSizeLimit,omitempty"`

	// expired at
	ExpiredAt string `json:"expiredAt,omitempty"`

	// failure reason
	FailureReason string `json:"failureReason,omitempty"`

	// flavor
	Flavor string `json:"flavor,omitempty"`

	// has alerting profile
	HasAlertingProfile bool `json:"hasAlertingProfile"`

	// has expiration warning
	HasExpirationWarning bool `json:"hasExpirationWarning"`

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`

	// has next version
	HasNextVersion bool `json:"hasNextVersion"`

	// has nfs server
	HasNfsServer bool `json:"hasNfsServer"`

	// has selected flavors
	HasSelectedFlavors bool `json:"hasSelectedFlavors"`

	// hypervisors
	Hypervisors []string `json:"hypervisors"`

	// is all failed upgrade
	IsAllFailedUpgrade bool `json:"isAllFailedUpgrade"`

	// is all ready
	IsAllReady bool `json:"isAllReady"`

	// is auto upgrade
	IsAutoUpgrade bool `json:"isAutoUpgrade"`

	// is autoscaling enabled
	IsAutoscalingEnabled bool `json:"isAutoscalingEnabled"`

	// is backup enabled
	IsBackupEnabled bool `json:"isBackupEnabled"`

	// is deprecated
	IsDeprecated bool `json:"isDeprecated"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes"`

	// is kubernetes current version kubevap enabled
	IsKubernetesCurrentVersionKubevapEnabled bool `json:"isKubernetesCurrentVersionKubevapEnabled"`

	// is kubevap enabled
	IsKubevapEnabled bool `json:"isKubevapEnabled"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// is maintenance mode enabled
	IsMaintenanceModeEnabled bool `json:"isMaintenanceModeEnabled"`

	// is monitoring enabled
	IsMonitoringEnabled bool `json:"isMonitoringEnabled"`

	// is opa enabled
	IsOpaEnabled bool `json:"isOpaEnabled"`

	// kube current version
	KubeCurrentVersion string `json:"kubeCurrentVersion,omitempty"`

	// kubernetes current version
	KubernetesCurrentVersion string `json:"kubernetesCurrentVersion,omitempty"`

	// kubernetes profile Id
	KubernetesProfileID int32 `json:"kubernetesProfileId,omitempty"`

	// kubernetes profile name
	KubernetesProfileName string `json:"kubernetesProfileName,omitempty"`

	// master
	Master int32 `json:"master,omitempty"`

	// master ready
	MasterReady int32 `json:"masterReady,omitempty"`

	// max size
	MaxSize int32 `json:"maxSize,omitempty"`

	// max spot price
	MaxSpotPrice float64 `json:"maxSpotPrice,omitempty"`

	// min size
	MinSize int32 `json:"minSize,omitempty"`

	// opa profile Id
	OpaProfileID int32 `json:"opaProfileId,omitempty"`

	// opa profile name
	OpaProfileName string `json:"opaProfileName,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// project cloud revision
	ProjectCloudRevision int32 `json:"projectCloudRevision,omitempty"`

	// project health
	ProjectHealth string `json:"projectHealth,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// project revision
	ProjectRevision int32 `json:"projectRevision,omitempty"`

	// project status
	ProjectStatus string `json:"projectStatus,omitempty"`

	// quota Id
	QuotaID int32 `json:"quotaId,omitempty"`

	// quota message
	QuotaMessage string `json:"quotaMessage,omitempty"`

	// ram limit
	RAMLimit int64 `json:"ramLimit,omitempty"`

	// s3 credential Id
	S3CredentialID int32 `json:"s3CredentialId,omitempty"`

	// spot enabled
	SpotEnabled bool `json:"spotEnabled"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`

	// total hourly cost
	TotalHourlyCost float64 `json:"totalHourlyCost,omitempty"`

	// used Cpu
	UsedCPU int64 `json:"usedCpu,omitempty"`

	// used disk size
	UsedDiskSize int64 `json:"usedDiskSize,omitempty"`

	// used Ram
	UsedRAM int64 `json:"usedRam,omitempty"`

	// vm Cpu limit
	VMCPULimit int64 `json:"vmCpuLimit,omitempty"`

	// vm Ram limit
	VMRAMLimit int64 `json:"vmRamLimit,omitempty"`

	// vm used Cpu
	VMUsedCPU int64 `json:"vmUsedCpu,omitempty"`

	// vm used Ram
	VMUsedRAM int64 `json:"vmUsedRam,omitempty"`

	// vm used volume size
	VMUsedVolumeSize int64 `json:"vmUsedVolumeSize,omitempty"`

	// vm volume size limit
	VMVolumeSizeLimit int64 `json:"vmVolumeSizeLimit,omitempty"`

	// warning message
	WarningMessage string `json:"warningMessage,omitempty"`

	// worker
	Worker int32 `json:"worker,omitempty"`
}

// Validate validates this project details for servers dto
func (m *ProjectDetailsForServersDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project details for servers dto based on context it is used
func (m *ProjectDetailsForServersDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectDetailsForServersDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectDetailsForServersDto) UnmarshalBinary(b []byte) error {
	var res ProjectDetailsForServersDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
