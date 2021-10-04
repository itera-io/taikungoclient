// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EditTicketCommand edit ticket command
//
// swagger:model EditTicketCommand
type EditTicketCommand struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ticket Id
	TicketID string `json:"ticketId,omitempty"`
}

// Validate validates this edit ticket command
func (m *EditTicketCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edit ticket command based on context it is used
func (m *EditTicketCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditTicketCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditTicketCommand) UnmarshalBinary(b []byte) error {
	var res EditTicketCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
