// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResponse Api response
//
// swagger:model ApiResponse
type APIResponse struct {

	// id
	ID string `json:"id,omitempty"`

	// is error
	IsError bool `json:"isError"`

	// message
	Message string `json:"message,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this Api response
func (m *APIResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api response based on context it is used
func (m *APIResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResponse) UnmarshalBinary(b []byte) error {
	var res APIResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
