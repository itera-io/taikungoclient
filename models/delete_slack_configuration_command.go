// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteSlackConfigurationCommand delete slack configuration command
//
// swagger:model DeleteSlackConfigurationCommand
type DeleteSlackConfigurationCommand struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this delete slack configuration command
func (m *DeleteSlackConfigurationCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete slack configuration command based on context it is used
func (m *DeleteSlackConfigurationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteSlackConfigurationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteSlackConfigurationCommand) UnmarshalBinary(b []byte) error {
	var res DeleteSlackConfigurationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
