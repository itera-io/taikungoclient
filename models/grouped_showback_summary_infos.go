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

// GroupedShowbackSummaryInfos grouped showback summary infos
//
// swagger:model GroupedShowbackSummaryInfos
type GroupedShowbackSummaryInfos struct {

	// data
	Data []*GroupedShowbackSummaryInfosDataItems0 `json:"data"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// total price
	TotalPrice float64 `json:"totalPrice,omitempty"`
}

// Validate validates this grouped showback summary infos
func (m *GroupedShowbackSummaryInfos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupedShowbackSummaryInfos) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this grouped showback summary infos based on the context it is used
func (m *GroupedShowbackSummaryInfos) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupedShowbackSummaryInfos) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GroupedShowbackSummaryInfos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupedShowbackSummaryInfos) UnmarshalBinary(b []byte) error {
	var res GroupedShowbackSummaryInfos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GroupedShowbackSummaryInfosDataItems0 grouped showback summary infos data items0
//
// swagger:model GroupedShowbackSummaryInfosDataItems0
type GroupedShowbackSummaryInfosDataItems0 struct {

	// price
	Price float64 `json:"price,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`
}

// Validate validates this grouped showback summary infos data items0
func (m *GroupedShowbackSummaryInfosDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this grouped showback summary infos data items0 based on context it is used
func (m *GroupedShowbackSummaryInfosDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GroupedShowbackSummaryInfosDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupedShowbackSummaryInfosDataItems0) UnmarshalBinary(b []byte) error {
	var res GroupedShowbackSummaryInfosDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
