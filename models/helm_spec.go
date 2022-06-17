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

// HelmSpec helm spec
//
// swagger:model HelmSpec
type HelmSpec struct {

	// chart
	Chart *Chart `json:"chart,omitempty"`

	// interval
	Interval string `json:"interval,omitempty"`

	// target namespace
	TargetNamespace string `json:"targetNamespace,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// values
	Values *JSONNode `json:"values,omitempty"`
}

// Validate validates this helm spec
func (m *HelmSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmSpec) validateChart(formats strfmt.Registry) error {
	if swag.IsZero(m.Chart) { // not required
		return nil
	}

	if m.Chart != nil {
		if err := m.Chart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *HelmSpec) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if m.Values != nil {
		if err := m.Values.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("values")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this helm spec based on the context it is used
func (m *HelmSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmSpec) contextValidateChart(ctx context.Context, formats strfmt.Registry) error {

	if m.Chart != nil {
		if err := m.Chart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chart")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chart")
			}
			return err
		}
	}

	return nil
}

func (m *HelmSpec) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	if m.Values != nil {
		if err := m.Values.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("values")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HelmSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelmSpec) UnmarshalBinary(b []byte) error {
	var res HelmSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
