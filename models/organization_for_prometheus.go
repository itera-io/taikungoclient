// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OrganizationForPrometheus organization for prometheus
//
// swagger:model OrganizationForPrometheus
type OrganizationForPrometheus struct {

	// global discount rate
	GlobalDiscountRate float64 `json:"globalDiscountRate"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// rule discount rate
	RuleDiscountRate float64 `json:"ruleDiscountRate"`
}

// Validate validates this organization for prometheus
func (m *OrganizationForPrometheus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization for prometheus based on context it is used
func (m *OrganizationForPrometheus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationForPrometheus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationForPrometheus) UnmarshalBinary(b []byte) error {
	var res OrganizationForPrometheus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
