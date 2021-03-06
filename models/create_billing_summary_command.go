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

// CreateBillingSummaryCommand create billing summary command
//
// swagger:model CreateBillingSummaryCommand
type CreateBillingSummaryCommand struct {

	// begin apply
	// Format: date-time
	BeginApply *strfmt.DateTime `json:"beginApply,omitempty"`

	// icu
	Icu int32 `json:"icu,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this create billing summary command
func (m *CreateBillingSummaryCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeginApply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBillingSummaryCommand) validateBeginApply(formats strfmt.Registry) error {
	if swag.IsZero(m.BeginApply) { // not required
		return nil
	}

	if err := validate.FormatOf("beginApply", "body", "date-time", m.BeginApply.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create billing summary command based on context it is used
func (m *CreateBillingSummaryCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBillingSummaryCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBillingSummaryCommand) UnmarshalBinary(b []byte) error {
	var res CreateBillingSummaryCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
