// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccessProfilesSearchCommand access profiles search command
//
// swagger:model AccessProfilesSearchCommand
type AccessProfilesSearchCommand struct {

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// search term
	SearchTerm string `json:"searchTerm,omitempty"`
}

// Validate validates this access profiles search command
func (m *AccessProfilesSearchCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this access profiles search command based on context it is used
func (m *AccessProfilesSearchCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessProfilesSearchCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessProfilesSearchCommand) UnmarshalBinary(b []byte) error {
	var res AccessProfilesSearchCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
