// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsBlockDeviceMappingsCommand aws block device mappings command
//
// swagger:model AwsBlockDeviceMappingsCommand
type AwsBlockDeviceMappingsCommand struct {

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this aws block device mappings command
func (m *AwsBlockDeviceMappingsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws block device mappings command based on context it is used
func (m *AwsBlockDeviceMappingsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsBlockDeviceMappingsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsBlockDeviceMappingsCommand) UnmarshalBinary(b []byte) error {
	var res AwsBlockDeviceMappingsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
