// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PrometheusEntity prometheus entity
//
// swagger:model PrometheusEntity
type PrometheusEntity struct {

	// prometheus rule Id
	PrometheusRuleID int32 `json:"prometheusRuleId,omitempty"`

	// prometheus rule name
	PrometheusRuleName string `json:"prometheusRuleName,omitempty"`

	// rule discount rate
	RuleDiscountRate float64 `json:"ruleDiscountRate"`
}

// Validate validates this prometheus entity
func (m *PrometheusEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus entity based on context it is used
func (m *PrometheusEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusEntity) UnmarshalBinary(b []byte) error {
	var res PrometheusEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
