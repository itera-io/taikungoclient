// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureZonesCommand azure zones command
//
// swagger:model AzureZonesCommand
type AzureZonesCommand struct {

	// azure client Id
	AzureClientID string `json:"azureClientId,omitempty"`

	// azure client secret
	AzureClientSecret string `json:"azureClientSecret,omitempty"`

	// azure location
	AzureLocation string `json:"azureLocation,omitempty"`

	// azure subscription Id
	AzureSubscriptionID string `json:"azureSubscriptionId,omitempty"`

	// azure tenant Id
	AzureTenantID string `json:"azureTenantId,omitempty"`
}

// Validate validates this azure zones command
func (m *AzureZonesCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure zones command based on context it is used
func (m *AzureZonesCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureZonesCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureZonesCommand) UnmarshalBinary(b []byte) error {
	var res AzureZonesCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
