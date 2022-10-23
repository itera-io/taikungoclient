// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateCatalogAppCommand create catalog app command
//
// swagger:model CreateCatalogAppCommand
type CreateCatalogAppCommand struct {

	// catalog Id
	CatalogID int32 `json:"catalogId,omitempty"`

	// package name
	PackageName string `json:"packageName,omitempty"`

	// parameters
	Parameters []*CreateCatalogAppCommandParametersItems0 `json:"parameters"`

	// repo name
	RepoName string `json:"repoName,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this create catalog app command
func (m *CreateCatalogAppCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCatalogAppCommand) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create catalog app command based on the context it is used
func (m *CreateCatalogAppCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCatalogAppCommand) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCatalogAppCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCatalogAppCommand) UnmarshalBinary(b []byte) error {
	var res CreateCatalogAppCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateCatalogAppCommandParametersItems0 create catalog app command parameters items0
//
// swagger:model CreateCatalogAppCommandParametersItems0
type CreateCatalogAppCommandParametersItems0 struct {

	// is changeable
	IsChangeable bool `json:"isChangeable"`

	// is readonly
	IsReadonly bool `json:"isReadonly"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this create catalog app command parameters items0
func (m *CreateCatalogAppCommandParametersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create catalog app command parameters items0 based on context it is used
func (m *CreateCatalogAppCommandParametersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateCatalogAppCommandParametersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCatalogAppCommandParametersItems0) UnmarshalBinary(b []byte) error {
	var res CreateCatalogAppCommandParametersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
