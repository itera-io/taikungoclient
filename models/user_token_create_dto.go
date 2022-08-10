// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserTokenCreateDto user token create dto
//
// swagger:model UserTokenCreateDto
type UserTokenCreateDto struct {

	// access key
	AccessKey string `json:"accessKey,omitempty"`

	// secret key
	SecretKey string `json:"secretKey,omitempty"`
}

// Validate validates this user token create dto
func (m *UserTokenCreateDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user token create dto based on context it is used
func (m *UserTokenCreateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserTokenCreateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTokenCreateDto) UnmarshalBinary(b []byte) error {
	var res UserTokenCreateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}