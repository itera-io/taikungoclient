// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

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

	// cloud credential Id
	CloudCredentialID int32 `json:"cloudCredentialId,omitempty"`

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

	// name
	Name string `json:"name,omitempty"`

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

	// taikun l b flavor
	TaikunLBFlavor string `json:"taikunLBFlavor,omitempty"`
}

// Validate validates this create project command
func (m *CreateProjectCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiredAt(formats); err != nil {
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

// ContextValidate validates this create project command based on context it is used
func (m *CreateProjectCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
