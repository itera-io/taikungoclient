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

// AlertingIntegrationDto alerting integration dto
//
// swagger:model AlertingIntegrationDto
type AlertingIntegrationDto struct {

	// alerting integration type
	AlertingIntegrationType AlertingIntegrationType `json:"alertingIntegrationType,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this alerting integration dto
func (m *AlertingIntegrationDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertingIntegrationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingIntegrationDto) validateAlertingIntegrationType(formats strfmt.Registry) error {
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

// ContextValidate validate this alerting integration dto based on the context it is used
func (m *AlertingIntegrationDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertingIntegrationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingIntegrationDto) contextValidateAlertingIntegrationType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AlertingIntegrationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingIntegrationDto) UnmarshalBinary(b []byte) error {
	var res AlertingIntegrationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
