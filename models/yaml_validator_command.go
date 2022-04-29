// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// YamlValidatorCommand yaml validator command
//
// swagger:model YamlValidatorCommand
type YamlValidatorCommand struct {

	// yaml
	Yaml string `json:"yaml,omitempty"`
}

// Validate validates this yaml validator command
func (m *YamlValidatorCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this yaml validator command based on context it is used
func (m *YamlValidatorCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *YamlValidatorCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *YamlValidatorCommand) UnmarshalBinary(b []byte) error {
	var res YamlValidatorCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
