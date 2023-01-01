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

// CreateProjectCommand create project command
//
// swagger:model CreateProjectCommand
type CreateProjectCommand struct {

	// access profile Id
	AccessProfileID int32 `json:"accessProfileId,omitempty"`

	// alerting profile Id
	AlertingProfileID int32 `json:"alertingProfileId,omitempty"`

	// allow full spot kubernetes
	AllowFullSpotKubernetes bool `json:"allowFullSpotKubernetes"`

	// allow spot v ms
	AllowSpotVMs bool `json:"allowSpotVMs"`

	// allow spot workers
	AllowSpotWorkers bool `json:"allowSpotWorkers"`

	// autoscaling enabled
	AutoscalingEnabled bool `json:"autoscalingEnabled"`

	// autoscaling flavor
	AutoscalingFlavor string `json:"autoscalingFlavor,omitempty"`

	// autoscaling group name
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty"`

	// autoscaling spot enabled
	AutoscalingSpotEnabled bool `json:"autoscalingSpotEnabled"`

	// cidr
	Cidr string `json:"cidr,omitempty"`

	// cloud credential Id
	CloudCredentialID int32 `json:"cloudCredentialId,omitempty"`

	// delete on expiration
	DeleteOnExpiration bool `json:"deleteOnExpiration"`

	// disk size
	DiskSize float64 `json:"diskSize,omitempty"`

	// expired at
	// Format: date-time
	ExpiredAt *strfmt.DateTime `json:"expiredAt,omitempty"`

	// flavors
	Flavors []string `json:"flavors"`

	// is auto upgrade
	IsAutoUpgrade bool `json:"isAutoUpgrade"`

	// is backup enabled
	IsBackupEnabled bool `json:"isBackupEnabled"`

	// is kubernetes
	IsKubernetes bool `json:"isKubernetes"`

	// is monitoring enabled
	IsMonitoringEnabled bool `json:"isMonitoringEnabled"`

	// kubernetes profile Id
	KubernetesProfileID int32 `json:"kubernetesProfileId,omitempty"`

	// kubernetes version
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`

	// max size
	MaxSize int32 `json:"maxSize,omitempty"`

	// max spot price
	MaxSpotPrice float64 `json:"maxSpotPrice,omitempty"`

	// min size
	MinSize int32 `json:"minSize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// net mask
	NetMask int32 `json:"netMask,omitempty"`

	// opa profile Id
	OpaProfileID int32 `json:"opaProfileId,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// router Id end range
	RouterIDEndRange int32 `json:"routerIdEndRange,omitempty"`

	// router Id start range
	RouterIDStartRange int32 `json:"routerIdStartRange,omitempty"`

	// s3 credential Id
	S3CredentialID int32 `json:"s3CredentialId,omitempty"`

	// save as template
	SaveAsTemplate bool `json:"saveAsTemplate"`

	// server templates
	ServerTemplates []*ServerTemplateDto `json:"serverTemplates"`

	// taikun l b flavor
	TaikunLBFlavor string `json:"taikunLBFlavor,omitempty"`

	// template name
	TemplateName string `json:"templateName,omitempty"`

	// users
	Users []string `json:"users"`
}

// Validate validates this create project command
func (m *CreateProjectCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiredAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProjectCommand) validateExpiredAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiredAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expiredAt", "body", "date-time", m.ExpiredAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateProjectCommand) validateServerTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerTemplates); i++ {
		if swag.IsZero(m.ServerTemplates[i]) { // not required
			continue
		}

		if m.ServerTemplates[i] != nil {
			if err := m.ServerTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create project command based on the context it is used
func (m *CreateProjectCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServerTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProjectCommand) contextValidateServerTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerTemplates); i++ {

		if m.ServerTemplates[i] != nil {
			if err := m.ServerTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateProjectCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProjectCommand) UnmarshalBinary(b []byte) error {
	var res CreateProjectCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
