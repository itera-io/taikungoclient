// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PartnerRecordDto partner record dto
//
// swagger:model PartnerRecordDto
type PartnerRecordDto struct {

	// allow registration
	AllowRegistration bool `json:"allowRegistration"`

	// id
	ID int32 `json:"id,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// payment enabled
	PaymentEnabled bool `json:"paymentEnabled"`
}

// Validate validates this partner record dto
func (m *PartnerRecordDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this partner record dto based on context it is used
func (m *PartnerRecordDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PartnerRecordDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PartnerRecordDto) UnmarshalBinary(b []byte) error {
	var res PartnerRecordDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
