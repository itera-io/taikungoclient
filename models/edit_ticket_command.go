// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EditTicketCommand edit ticket command
//
// swagger:model EditTicketCommand
type EditTicketCommand struct {

	// description
	// Required: true
	// Max Length: 2000
	// Min Length: 3
	Description *string `json:"description"`

	// name
	// Required: true
	// Max Length: 200
	// Min Length: 3
	Name *string `json:"name"`

	// ticket Id
	// Required: true
	// Min Length: 1
	TicketID *string `json:"ticketId"`
}

// Validate validates this edit ticket command
func (m *EditTicketCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditTicketCommand) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 2000); err != nil {
		return err
	}

	return nil
}

func (m *EditTicketCommand) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 200); err != nil {
		return err
	}

	return nil
}

func (m *EditTicketCommand) validateTicketID(formats strfmt.Registry) error {

	if err := validate.Required("ticketId", "body", m.TicketID); err != nil {
		return err
	}

	if err := validate.MinLength("ticketId", "body", *m.TicketID, 1); err != nil {
		return err
	}

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
