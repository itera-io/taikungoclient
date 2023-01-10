// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectForListDto project for list dto
//
// swagger:model ProjectForListDto
type ProjectForListDto struct {

	// access Ip
	AccessIP string `json:"accessIp,omitempty"`

	// access profile
	AccessProfile *AccessProfilesForProjectListDto `json:"accessProfile,omitempty"`

	// access profile revision
	AccessProfileRevision int32 `json:"accessProfileRevision,omitempty"`

	// availability zones
	AvailabilityZones []string `json:"availabilityZones"`

	// aws project a z subnets
	AwsProjectAZSubnets []*AwsProjectAZSubnetDto `json:"awsProjectAZSubnets"`

	// bastion
	Bastion int32 `json:"bastion,omitempty"`

	// bound users
	BoundUsers []*UserDto `json:"boundUsers"`

	// cidr
	Cidr string `json:"cidr,omitempty"`

	// cloud credential Id
	CloudCredentialID int32 `json:"cloudCredentialId,omitempty"`

	// cloud credential name
	CloudCredentialName string `json:"cloudCredentialName,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// flavors
	Flavors []string `json:"flavors"`

	// google project Id
	GoogleProjectID string `json:"googleProjectId,omitempty"`

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile"`

	// has selected flavors
	HasSelectedFlavors bool `json:"hasSelectedFlavors"`

	// health
	Health string `json:"health,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// is auto upgrade
	IsAutoUpgrade bool `json:"isAutoUpgrade"`

	// is autoscaling enabled
	IsAutoscalingEnabled bool `json:"isAutoscalingEnabled"`

	// is backup enabled
	IsBackupEnabled bool `json:"isBackupEnabled"`

	// is delete cluster
	IsDeleteCluster bool `json:"isDeleteCluster"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes"`

	// is kubevap enabled
	IsKubevapEnabled bool `json:"isKubevapEnabled"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// is monitoring enabled
	IsMonitoringEnabled bool `json:"isMonitoringEnabled"`

	// is opa enabled
	IsOpaEnabled bool `json:"isOpaEnabled"`

	// job Url
	JobURL string `json:"jobUrl,omitempty"`

	// kubernetes alerts
	KubernetesAlerts []*KubernetesAlertDto `json:"kubernetesAlerts"`

	// kubernetes current version
	KubernetesCurrentVersion string `json:"kubernetesCurrentVersion,omitempty"`

	// kubernetes profiles
	KubernetesProfiles *KubernetesProfilesLisForPollerDto `json:"kubernetesProfiles,omitempty"`

	// kubernetes target version
	KubernetesTargetVersion string `json:"kubernetesTargetVersion,omitempty"`

	// kubespray current version
	KubesprayCurrentVersion string `json:"kubesprayCurrentVersion,omitempty"`

	// kubespray target version
	KubesprayTargetVersion string `json:"kubesprayTargetVersion,omitempty"`

	// kubevap enabeled kubernetes versions
	KubevapEnabeledKubernetesVersions []string `json:"kubevapEnabeledKubernetesVersions"`

	// master
	Master int32 `json:"master,omitempty"`

	// master ready
	MasterReady int32 `json:"masterReady,omitempty"`

	// monitoring credential
	MonitoringCredential *MonitoringCredentialsListDto `json:"monitoringCredential,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// net mask
	NetMask int32 `json:"netMask,omitempty"`

	// opa profile
	OpaProfile *OpaProfileListDto `json:"opaProfile,omitempty"`

	// opa profile revision
	OpaProfileRevision int32 `json:"opaProfileRevision,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// partner Id
	PartnerID int32 `json:"partnerId,omitempty"`

	// private Ip
	PrivateIP string `json:"privateIp,omitempty"`

	// public Ip
	PublicIP string `json:"publicIp,omitempty"`

	// quota Id
	QuotaID int32 `json:"quotaId,omitempty"`

	// router Id end range
	RouterIDEndRange int32 `json:"routerIdEndRange,omitempty"`

	// router Id start range
	RouterIDStartRange int32 `json:"routerIdStartRange,omitempty"`

	// s3 access key Id
	S3AccessKeyID string `json:"s3AccessKeyId,omitempty"`

	// s3 bucket name
	S3BucketName string `json:"s3BucketName,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 region
	S3Region string `json:"s3Region,omitempty"`

	// s3 secret key
	S3SecretKey string `json:"s3SecretKey,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// taikun l b flavor
	TaikunLBFlavor string `json:"taikunLBFlavor,omitempty"`

	// taikun l b private key
	TaikunLBPrivateKey string `json:"taikunLBPrivateKey,omitempty"`

	// taikun l b public key
	TaikunLBPublicKey string `json:"taikunLBPublicKey,omitempty"`

	// taikun private SSH key
	TaikunPrivateSSHKey string `json:"taikunPrivateSSHKey,omitempty"`

	// taikun public SSH key
	TaikunPublicSSHKey string `json:"taikunPublicSSHKey,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// topic name
	TopicName string `json:"topicName,omitempty"`

	// total servers count
	TotalServersCount int32 `json:"totalServersCount,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this project for list dto
func (m *ProjectForListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsProjectAZSubnets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBoundUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoringCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpaProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectForListDto) validateAccessProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessProfile) { // not required
		return nil
	}

	if m.AccessProfile != nil {
		if err := m.AccessProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessProfile")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) validateAwsProjectAZSubnets(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsProjectAZSubnets) { // not required
		return nil
	}

	for i := 0; i < len(m.AwsProjectAZSubnets); i++ {
		if swag.IsZero(m.AwsProjectAZSubnets[i]) { // not required
			continue
		}

		if m.AwsProjectAZSubnets[i] != nil {
			if err := m.AwsProjectAZSubnets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("awsProjectAZSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("awsProjectAZSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) validateBoundUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.BoundUsers); i++ {
		if swag.IsZero(m.BoundUsers[i]) { // not required
			continue
		}

		if m.BoundUsers[i] != nil {
			if err := m.BoundUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) validateKubernetesAlerts(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesAlerts) { // not required
		return nil
	}

	for i := 0; i < len(m.KubernetesAlerts); i++ {
		if swag.IsZero(m.KubernetesAlerts[i]) { // not required
			continue
		}

		if m.KubernetesAlerts[i] != nil {
			if err := m.KubernetesAlerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) validateKubernetesProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesProfiles) { // not required
		return nil
	}

	if m.KubernetesProfiles != nil {
		if err := m.KubernetesProfiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesProfiles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesProfiles")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) validateMonitoringCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitoringCredential) { // not required
		return nil
	}

	if m.MonitoringCredential != nil {
		if err := m.MonitoringCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoringCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitoringCredential")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) validateOpaProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.OpaProfile) { // not required
		return nil
	}

	if m.OpaProfile != nil {
		if err := m.OpaProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaProfile")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this project for list dto based on the context it is used
func (m *ProjectForListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsProjectAZSubnets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBoundUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetesAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetesProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitoringCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpaProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectForListDto) contextValidateAccessProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessProfile != nil {
		if err := m.AccessProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessProfile")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) contextValidateAwsProjectAZSubnets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AwsProjectAZSubnets); i++ {

		if m.AwsProjectAZSubnets[i] != nil {
			if err := m.AwsProjectAZSubnets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("awsProjectAZSubnets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("awsProjectAZSubnets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) contextValidateBoundUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BoundUsers); i++ {

		if m.BoundUsers[i] != nil {
			if err := m.BoundUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) contextValidateKubernetesAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.KubernetesAlerts); i++ {

		if m.KubernetesAlerts[i] != nil {
			if err := m.KubernetesAlerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kubernetesAlerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) contextValidateKubernetesProfiles(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesProfiles != nil {
		if err := m.KubernetesProfiles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetesProfiles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetesProfiles")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) contextValidateMonitoringCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.MonitoringCredential != nil {
		if err := m.MonitoringCredential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoringCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("monitoringCredential")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectForListDto) contextValidateOpaProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.OpaProfile != nil {
		if err := m.OpaProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("opaProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("opaProfile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectForListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectForListDto) UnmarshalBinary(b []byte) error {
	var res ProjectForListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
