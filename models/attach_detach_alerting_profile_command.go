// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttachDetachAlertingProfileCommand attach detach alerting profile command
//
// swagger:model AttachDetachAlertingProfileCommand
type AttachDetachAlertingProfileCommand struct {

	// alerting profile Id
	AlertingProfileID int32 `json:"alertingProfileId,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this attach detach alerting profile command
func (m *AttachDetachAlertingProfileCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attach detach alerting profile command based on context it is used
func (m *AttachDetachAlertingProfileCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachDetachAlertingProfileCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachDetachAlertingProfileCommand) UnmarshalBinary(b []byte) error {
	var res AttachDetachAlertingProfileCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
