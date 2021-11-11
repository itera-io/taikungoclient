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

// PrometheusMetricsCommand prometheus metrics command
//
// swagger:model PrometheusMetricsCommand
type PrometheusMetricsCommand struct {

	// end
	// Format: date-time
	End *strfmt.DateTime `json:"end,omitempty"`

	// is auto complete
	IsAutoComplete bool `json:"isAutoComplete,omitempty"`

	// is graph enabled
	IsGraphEnabled bool `json:"isGraphEnabled,omitempty"`

	// parameters
	Parameters string `json:"parameters,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// start
	// Format: date-time
	Start *strfmt.DateTime `json:"start,omitempty"`

	// step
	Step string `json:"step,omitempty"`

	// time
	// Format: date-time
	Time *strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this prometheus metrics command
func (m *PrometheusMetricsCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusMetricsCommand) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PrometheusMetricsCommand) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PrometheusMetricsCommand) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this prometheus metrics command based on context it is used
func (m *PrometheusMetricsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusMetricsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusMetricsCommand) UnmarshalBinary(b []byte) error {
	var res PrometheusMetricsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
