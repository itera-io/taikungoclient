// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodListDto pod list dto
//
// swagger:model PodListDto
type PodListDto struct {

	// age
	Age string `json:"age,omitempty"`

	// container
	Container []string `json:"container"`

	// metadata name
	MetadataName string `json:"metadataName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// node
	Node string `json:"node,omitempty"`

	// restart count
	RestartCount int32 `json:"restartCount,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this pod list dto
func (m *PodListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pod list dto based on context it is used
func (m *PodListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodListDto) UnmarshalBinary(b []byte) error {
	var res PodListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
