// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemindUsersByAlertingProfileCommand remind users by alerting profile command
//
// swagger:model RemindUsersByAlertingProfileCommand
type RemindUsersByAlertingProfileCommand struct {

	// reminder
	Reminder AlertingReminder `json:"reminder,omitempty"`
}

// Validate validates this remind users by alerting profile command
func (m *RemindUsersByAlertingProfileCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReminder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemindUsersByAlertingProfileCommand) validateReminder(formats strfmt.Registry) error {
	if swag.IsZero(m.Reminder) { // not required
		return nil
	}

	if err := m.Reminder.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reminder")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reminder")
		}
		return err
	}

	return nil
}

// ContextValidate validate this remind users by alerting profile command based on the context it is used
func (m *RemindUsersByAlertingProfileCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReminder(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemindUsersByAlertingProfileCommand) contextValidateReminder(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Reminder.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("reminder")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("reminder")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemindUsersByAlertingProfileCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemindUsersByAlertingProfileCommand) UnmarshalBinary(b []byte) error {
	var res RemindUsersByAlertingProfileCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
