// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StandAloneProfileUpdateCommand stand alone profile update command
//
// swagger:model StandAloneProfileUpdateCommand
type StandAloneProfileUpdateCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this stand alone profile update command
func (m *StandAloneProfileUpdateCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this stand alone profile update command based on context it is used
func (m *StandAloneProfileUpdateCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StandAloneProfileUpdateCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandAloneProfileUpdateCommand) UnmarshalBinary(b []byte) error {
	var res StandAloneProfileUpdateCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
