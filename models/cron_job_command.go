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

// CronJobCommand cron job command
//
// swagger:model CronJobCommand
type CronJobCommand struct {

	// cron period
	// Required: true
	// Min Length: 1
	CronPeriod *string `json:"cronPeriod"`
}

// Validate validates this cron job command
func (m *CronJobCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCronPeriod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CronJobCommand) validateCronPeriod(formats strfmt.Registry) error {

	if err := validate.Required("cronPeriod", "body", m.CronPeriod); err != nil {
		return err
	}

	if err := validate.MinLength("cronPeriod", "body", *m.CronPeriod, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cron job command based on context it is used
func (m *CronJobCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CronJobCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CronJobCommand) UnmarshalBinary(b []byte) error {
	var res CronJobCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
