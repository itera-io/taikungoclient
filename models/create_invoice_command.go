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

// CreateInvoiceCommand create invoice command
//
// swagger:model CreateInvoiceCommand
type CreateInvoiceCommand struct {

	// due date
	// Format: date-time
	DueDate *strfmt.DateTime `json:"dueDate,omitempty"`

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate,omitempty"`

	// is paid
	IsPaid bool `json:"isPaid"`

	// name
	Name string `json:"name,omitempty"`

	// organization subscription Id
	OrganizationSubscriptionID int32 `json:"organizationSubscriptionId,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// required payment action
	RequiredPaymentAction bool `json:"requiredPaymentAction"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate,omitempty"`

	// stripe invoice Id
	StripeInvoiceID string `json:"stripeInvoiceId,omitempty"`
}

// Validate validates this create invoice command
func (m *CreateInvoiceCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDueDate(formats); err != nil {
		res = append(res, err)
	}

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

func (m *CreateInvoiceCommand) validateDueDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dueDate", "body", "date-time", m.DueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateInvoiceCommand) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateInvoiceCommand) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create invoice command based on context it is used
func (m *CreateInvoiceCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateInvoiceCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateInvoiceCommand) UnmarshalBinary(b []byte) error {
	var res CreateInvoiceCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}