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

	// access profile revision
	AccessProfileRevision int32 `json:"accessProfileRevision,omitempty"`

	// access profiles
	AccessProfiles *AccessProfilesForProjectListDto `json:"accessProfiles,omitempty"`

	// bastion
	Bastion int32 `json:"bastion,omitempty"`

	// bound users
	BoundUsers []*UserDto `json:"boundUsers"`

	// cloud credential name
	CloudCredentialName string `json:"cloudCredentialName,omitempty"`

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// flavors
	Flavors []string `json:"flavors"`

	// has kube config file
	HasKubeConfigFile bool `json:"hasKubeConfigFile,omitempty"`

	// has selected flavors
	HasSelectedFlavors bool `json:"hasSelectedFlavors,omitempty"`

	// health
	Health string `json:"health,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// image name
	ImageName string `json:"imageName,omitempty"`

	// is auto upgrade
	IsAutoUpgrade bool `json:"isAutoUpgrade,omitempty"`

	// is backup enabled
	IsBackupEnabled bool `json:"isBackupEnabled,omitempty"`

	// is delete cluster
	IsDeleteCluster bool `json:"isDeleteCluster,omitempty"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked,omitempty"`

	// is monitoring enabled
	IsMonitoringEnabled bool `json:"isMonitoringEnabled,omitempty"`

	// job Url
	JobURL string `json:"jobUrl,omitempty"`

	// kubernetes alerts
	KubernetesAlerts []*KubernetesAlertDto `json:"kubernetesAlerts"`

	// kubernetes current version
	KubernetesCurrentVersion string `json:"kubernetesCurrentVersion,omitempty"`

	// kubernetes profiles
	KubernetesProfiles *KubernetesProfilesListDto `json:"kubernetesProfiles,omitempty"`

	// kubernetes target version
	KubernetesTargetVersion string `json:"kubernetesTargetVersion,omitempty"`

	// kubespray current version
	KubesprayCurrentVersion string `json:"kubesprayCurrentVersion,omitempty"`

	// kubespray target version
	KubesprayTargetVersion string `json:"kubesprayTargetVersion,omitempty"`

	// master
	Master int32 `json:"master,omitempty"`

	// master ready
	MasterReady int32 `json:"masterReady,omitempty"`

	// monitoring credentials
	MonitoringCredentials []*MonitoringCredentialsListDto `json:"monitoringCredentials"`

	// name
	Name string `json:"name,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

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

	// standalone vms
	StandaloneVms []*StandAloneVMFullDto `json:"standaloneVms"`

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

	// total servers count
	TotalServersCount int32 `json:"totalServersCount,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this project for list dto
func (m *ProjectForListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessProfiles(formats); err != nil {
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

	if err := m.validateMonitoringCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandaloneVms(formats); err != nil {
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

func (m *ProjectForListDto) validateAccessProfiles(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessProfiles) { // not required
		return nil
	}

	if m.AccessProfiles != nil {
		if err := m.AccessProfiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessProfiles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessProfiles")
			}
			return err
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

func (m *ProjectForListDto) validateMonitoringCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitoringCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.MonitoringCredentials); i++ {
		if swag.IsZero(m.MonitoringCredentials[i]) { // not required
			continue
		}

		if m.MonitoringCredentials[i] != nil {
			if err := m.MonitoringCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) validateStandaloneVms(formats strfmt.Registry) error {
	if swag.IsZero(m.StandaloneVms) { // not required
		return nil
	}

	for i := 0; i < len(m.StandaloneVms); i++ {
		if swag.IsZero(m.StandaloneVms[i]) { // not required
			continue
		}

		if m.StandaloneVms[i] != nil {
			if err := m.StandaloneVms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				}
				return err
			}
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

	if err := m.contextValidateAccessProfiles(ctx, formats); err != nil {
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

	if err := m.contextValidateMonitoringCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandaloneVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectForListDto) contextValidateAccessProfiles(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessProfiles != nil {
		if err := m.AccessProfiles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessProfiles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessProfiles")
			}
			return err
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

func (m *ProjectForListDto) contextValidateMonitoringCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MonitoringCredentials); i++ {

		if m.MonitoringCredentials[i] != nil {
			if err := m.MonitoringCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("monitoringCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProjectForListDto) contextValidateStandaloneVms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StandaloneVms); i++ {

		if m.StandaloneVms[i] != nil {
			if err := m.StandaloneVms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVms" + "." + strconv.Itoa(i))
				}
				return err
			}
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
