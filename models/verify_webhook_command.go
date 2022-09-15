// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VerifyWebhookCommand verify webhook command
//
// swagger:model VerifyWebhookCommand
type VerifyWebhookCommand struct {

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this verify webhook command
func (m *VerifyWebhookCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this verify webhook command based on context it is used
func (m *VerifyWebhookCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VerifyWebhookCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerifyWebhookCommand) UnmarshalBinary(b []byte) error {
	var res VerifyWebhookCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
