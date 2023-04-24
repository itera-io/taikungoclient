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

// CreateTicketCommand create ticket command
//
// swagger:model CreateTicketCommand
type CreateTicketCommand struct {

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

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// priority
	Priority TicketPriority `json:"priority,omitempty"`
}

// Validate validates this create ticket command
func (m *CreateTicketCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTicketCommand) validateDescription(formats strfmt.Registry) error {

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

func (m *CreateTicketCommand) validateName(formats strfmt.Registry) error {

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

func (m *CreateTicketCommand) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

// ContextValidate validate this create ticket command based on the context it is used
func (m *CreateTicketCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateTicketCommand) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Priority.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateTicketCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTicketCommand) UnmarshalBinary(b []byte) error {
	var res CreateTicketCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
