// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesVersionListDto kubernetes version list dto
//
// swagger:model KubernetesVersionListDto
type KubernetesVersionListDto struct {

	// is kubevap enabled
	IsKubevapEnabled bool `json:"isKubevapEnabled"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this kubernetes version list dto
func (m *KubernetesVersionListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes version list dto based on context it is used
func (m *KubernetesVersionListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesVersionListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesVersionListDto) UnmarshalBinary(b []byte) error {
	var res KubernetesVersionListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}