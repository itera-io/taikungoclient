// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoogleCommonImages google common images
//
// swagger:model GoogleCommonImages
type GoogleCommonImages struct {

	// display name
	DisplayName string `json:"displayName"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this google common images
func (m *GoogleCommonImages) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google common images based on context it is used
func (m *GoogleCommonImages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GoogleCommonImages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleCommonImages) UnmarshalBinary(b []byte) error {
	var res GoogleCommonImages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}