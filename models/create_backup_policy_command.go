// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateBackupPolicyCommand create backup policy command
//
// swagger:model CreateBackupPolicyCommand
type CreateBackupPolicyCommand struct {

	// cron period
	CronPeriod string `json:"cronPeriod,omitempty"`

	// exclude namespaces
	ExcludeNamespaces []string `json:"excludeNamespaces"`

	// include namespaces
	IncludeNamespaces []string `json:"includeNamespaces"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// retention period
	RetentionPeriod string `json:"retentionPeriod,omitempty"`
}

// Validate validates this create backup policy command
func (m *CreateBackupPolicyCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create backup policy command based on context it is used
func (m *CreateBackupPolicyCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBackupPolicyCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBackupPolicyCommand) UnmarshalBinary(b []byte) error {
	var res CreateBackupPolicyCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
