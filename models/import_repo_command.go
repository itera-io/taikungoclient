// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImportRepoCommand import repo command
//
// swagger:model ImportRepoCommand
type ImportRepoCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this import repo command
func (m *ImportRepoCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this import repo command based on context it is used
func (m *ImportRepoCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportRepoCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportRepoCommand) UnmarshalBinary(b []byte) error {
	var res ImportRepoCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
