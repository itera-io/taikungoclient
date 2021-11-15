// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingSummaryDto billing summary dto
//
// swagger:model BillingSummaryDto
type BillingSummaryDto struct {

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate,omitempty"`

	// is deleted
	IsDeleted bool `json:"isDeleted"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate,omitempty"`

	// tcu
	Tcu float64 `json:"tcu,omitempty"`
}

// Validate validates this billing summary dto
func (m *BillingSummaryDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BillingSummaryDto) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingSummaryDto) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this billing summary dto based on context it is used
func (m *BillingSummaryDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingSummaryDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingSummaryDto) UnmarshalBinary(b []byte) error {
	var res BillingSummaryDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
