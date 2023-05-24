// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StandaloneVisibilityDto standalone visibility dto
//
// swagger:model StandaloneVisibilityDto
type StandaloneVisibilityDto struct {

	// reboot
	Reboot bool `json:"reboot"`

	// shelve
	Shelve bool `json:"shelve"`

	// show console
	ShowConsole bool `json:"showConsole"`

	// show status
	ShowStatus bool `json:"showStatus"`

	// start
	Start bool `json:"start"`

	// stop
	Stop bool `json:"stop"`

	// unshelve
	Unshelve bool `json:"unshelve"`
}

// Validate validates this standalone visibility dto
func (m *StandaloneVisibilityDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this standalone visibility dto based on context it is used
func (m *StandaloneVisibilityDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StandaloneVisibilityDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StandaloneVisibilityDto) UnmarshalBinary(b []byte) error {
	var res StandaloneVisibilityDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}