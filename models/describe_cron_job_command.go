// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeCronJobCommand describe cron job command
//
// swagger:model DescribeCronJobCommand
type DescribeCronJobCommand struct {

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this describe cron job command
func (m *DescribeCronJobCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe cron job command based on context it is used
func (m *DescribeCronJobCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeCronJobCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeCronJobCommand) UnmarshalBinary(b []byte) error {
	var res DescribeCronJobCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
