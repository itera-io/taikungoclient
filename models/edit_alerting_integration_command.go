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

// EditAlertingIntegrationCommand edit alerting integration command
//
// swagger:model EditAlertingIntegrationCommand
type EditAlertingIntegrationCommand struct {

	// alerting integration type
	AlertingIntegrationType AlertingIntegrationType `json:"alertingIntegrationType,omitempty"`

	// alerting profile Id
	AlertingProfileID int32 `json:"alertingProfileId,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this edit alerting integration command
func (m *EditAlertingIntegrationCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertingIntegrationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditAlertingIntegrationCommand) validateAlertingIntegrationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertingIntegrationType) { // not required
		return nil
	}

	if err := m.AlertingIntegrationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("alertingIntegrationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("alertingIntegrationType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this edit alerting integration command based on the context it is used
func (m *EditAlertingIntegrationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertingIntegrationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditAlertingIntegrationCommand) contextValidateAlertingIntegrationType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AlertingIntegrationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("alertingIntegrationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("alertingIntegrationType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EditAlertingIntegrationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditAlertingIntegrationCommand) UnmarshalBinary(b []byte) error {
	var res EditAlertingIntegrationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
