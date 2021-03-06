// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureQuotaListRecordDto azure quota list record dto
//
// swagger:model AzureQuotaListRecordDto
type AzureQuotaListRecordDto struct {

	// current usage
	CurrentUsage int32 `json:"currentUsage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// total cores
	TotalCores int64 `json:"totalCores,omitempty"`
}

// Validate validates this azure quota list record dto
func (m *AzureQuotaListRecordDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure quota list record dto based on context it is used
func (m *AzureQuotaListRecordDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureQuotaListRecordDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureQuotaListRecordDto) UnmarshalBinary(b []byte) error {
	var res AzureQuotaListRecordDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
