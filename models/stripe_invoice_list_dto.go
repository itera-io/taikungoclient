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

// StripeInvoiceListDto stripe invoice list dto
//
// swagger:model StripeInvoiceListDto
type StripeInvoiceListDto struct {

	// charge reason
	ChargeReason string `json:"chargeReason,omitempty"`

	// charge status
	ChargeStatus string `json:"chargeStatus,omitempty"`

	// end date
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// invoice status
	InvoiceStatus string `json:"invoiceStatus,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// start date
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`
}

// Validate validates this stripe invoice list dto
func (m *StripeInvoiceListDto) Validate(formats strfmt.Registry) error {
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

func (m *StripeInvoiceListDto) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StripeInvoiceListDto) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stripe invoice list dto based on context it is used
func (m *StripeInvoiceListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StripeInvoiceListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StripeInvoiceListDto) UnmarshalBinary(b []byte) error {
	var res StripeInvoiceListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
