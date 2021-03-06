// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureSkusList azure skus list
//
// swagger:model AzureSkusList
type AzureSkusList struct {

	// data
	Data []string `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this azure skus list
func (m *AzureSkusList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure skus list based on context it is used
func (m *AzureSkusList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureSkusList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSkusList) UnmarshalBinary(b []byte) error {
	var res AzureSkusList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
