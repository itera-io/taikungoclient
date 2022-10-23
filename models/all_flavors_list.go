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

// AllFlavorsList all flavors list
//
// swagger:model AllFlavorsList
type AllFlavorsList struct {

	// cloud type
	CloudType string `json:"cloudType,omitempty"`

	// data
	Data []*AllFlavorsListDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this all flavors list
func (m *AllFlavorsList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllFlavorsList) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this all flavors list based on the context it is used
func (m *AllFlavorsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllFlavorsList) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllFlavorsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllFlavorsList) UnmarshalBinary(b []byte) error {
	var res AllFlavorsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AllFlavorsListDataItems0 all flavors list data items0
//
// swagger:model AllFlavorsListDataItems0
type AllFlavorsListDataItems0 struct {

	// cpu
	CPU int32 `json:"cpu,omitempty"`

	// description
	Description interface{} `json:"description,omitempty"`

	// max data disk count
	MaxDataDiskCount float64 `json:"maxDataDiskCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ram
	RAM int64 `json:"ram,omitempty"`
}

// Validate validates this all flavors list data items0
func (m *AllFlavorsListDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this all flavors list data items0 based on context it is used
func (m *AllFlavorsListDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AllFlavorsListDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllFlavorsListDataItems0) UnmarshalBinary(b []byte) error {
	var res AllFlavorsListDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
