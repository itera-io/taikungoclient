// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GroupedShowbackSummaries grouped showback summaries
//
// swagger:model GroupedShowbackSummaries
type GroupedShowbackSummaries struct {

	// price
	Price float64 `json:"price,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`
}

// Validate validates this grouped showback summaries
func (m *GroupedShowbackSummaries) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this grouped showback summaries based on context it is used
func (m *GroupedShowbackSummaries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupedShowbackSummaries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupedShowbackSummaries) UnmarshalBinary(b []byte) error {
	var res GroupedShowbackSummaries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
