// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkPolicyDto network policy dto
//
// swagger:model NetworkPolicyDto
type NetworkPolicyDto struct {

	// age
	Age string `json:"age,omitempty"`

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// pod selector
	PodSelector map[string]string `json:"podSelector,omitempty"`
}

// Validate validates this network policy dto
func (m *NetworkPolicyDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network policy dto based on context it is used
func (m *NetworkPolicyDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPolicyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPolicyDto) UnmarshalBinary(b []byte) error {
	var res NetworkPolicyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
