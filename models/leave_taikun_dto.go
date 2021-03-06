// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LeaveTaikunDto leave taikun dto
//
// swagger:model LeaveTaikunDto
type LeaveTaikunDto struct {

	// payment client secret
	PaymentClientSecret string `json:"paymentClientSecret,omitempty"`

	// payment intent Id
	PaymentIntentID string `json:"paymentIntentId,omitempty"`
}

// Validate validates this leave taikun dto
func (m *LeaveTaikunDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this leave taikun dto based on context it is used
func (m *LeaveTaikunDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeaveTaikunDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaveTaikunDto) UnmarshalBinary(b []byte) error {
	var res LeaveTaikunDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
