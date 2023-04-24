// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMTemplateListCommand Vm template list command
//
// swagger:model VmTemplateListCommand
type VMTemplateListCommand struct {

	// password
	Password string `json:"password,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this Vm template list command
func (m *VMTemplateListCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Vm template list command based on context it is used
func (m *VMTemplateListCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplateListCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplateListCommand) UnmarshalBinary(b []byte) error {
	var res VMTemplateListCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
