// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FlavorsListDto flavors list dto
//
// swagger:model FlavorsListDto
type FlavorsListDto struct {

	// cpu
	CPU int32 `json:"cpu,omitempty"`

	// description
	Description interface{} `json:"description,omitempty"`

	// linux price
	LinuxPrice string `json:"linuxPrice,omitempty"`

	// linux spot price
	LinuxSpotPrice string `json:"linuxSpotPrice,omitempty"`

	// max data disk count
	MaxDataDiskCount float64 `json:"maxDataDiskCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ram
	RAM int64 `json:"ram,omitempty"`

	// windows price
	WindowsPrice string `json:"windowsPrice,omitempty"`

	// windows spot price
	WindowsSpotPrice string `json:"windowsSpotPrice,omitempty"`
}

// Validate validates this flavors list dto
func (m *FlavorsListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this flavors list dto based on context it is used
func (m *FlavorsListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FlavorsListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlavorsListDto) UnmarshalBinary(b []byte) error {
	var res FlavorsListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
