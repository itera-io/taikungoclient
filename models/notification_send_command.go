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

// NotificationSendCommand notification send command
//
// swagger:model NotificationSendCommand
type NotificationSendCommand struct {

	// action status
	ActionStatus ActionStatus `json:"actionStatus,omitempty"`

	// action type
	ActionType ActionType `json:"actionType,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project type
	ProjectType ProjectType `json:"projectType,omitempty"`
}

// Validate validates this notification send command
func (m *NotificationSendCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationSendCommand) validateActionStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionStatus) { // not required
		return nil
	}

	if err := m.ActionStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("actionStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("actionStatus")
		}
		return err
	}

	return nil
}

func (m *NotificationSendCommand) validateActionType(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionType) { // not required
		return nil
	}

	if err := m.ActionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("actionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("actionType")
		}
		return err
	}

	return nil
}

func (m *NotificationSendCommand) validateProjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectType) { // not required
		return nil
	}

	if err := m.ProjectType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("projectType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("projectType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this notification send command based on the context it is used
func (m *NotificationSendCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateActionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationSendCommand) contextValidateActionStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActionStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("actionStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("actionStatus")
		}
		return err
	}

	return nil
}

func (m *NotificationSendCommand) contextValidateActionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ActionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("actionType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("actionType")
		}
		return err
	}

	return nil
}

func (m *NotificationSendCommand) contextValidateProjectType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ProjectType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("projectType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("projectType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationSendCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationSendCommand) UnmarshalBinary(b []byte) error {
	var res NotificationSendCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
