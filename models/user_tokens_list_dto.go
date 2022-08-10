// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserTokensListDto user tokens list dto
//
// swagger:model UserTokensListDto
type UserTokensListDto struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// expire date
	ExpireDate string `json:"expireDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// is readonly
	IsReadonly bool `json:"isReadonly"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this user tokens list dto
func (m *UserTokensListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user tokens list dto based on context it is used
func (m *UserTokensListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserTokensListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTokensListDto) UnmarshalBinary(b []byte) error {
	var res UserTokensListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}