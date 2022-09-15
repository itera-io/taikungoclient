// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OpenstackFlavorListDto openstack flavor list dto
//
// swagger:model OpenstackFlavorListDto
type OpenstackFlavorListDto struct {

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ram
	RAM int64 `json:"ram,omitempty"`
}

// Validate validates this openstack flavor list dto
func (m *OpenstackFlavorListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack flavor list dto based on context it is used
func (m *OpenstackFlavorListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackFlavorListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackFlavorListDto) UnmarshalBinary(b []byte) error {
	var res OpenstackFlavorListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
