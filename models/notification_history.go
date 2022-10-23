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

// NotificationHistory notification history
//
// swagger:model NotificationHistory
type NotificationHistory struct {

	// data
	Data []*NotificationHistoryDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this notification history
func (m *NotificationHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationHistory) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this notification history based on the context it is used
func (m *NotificationHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationHistory) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NotificationHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationHistory) UnmarshalBinary(b []byte) error {
	var res NotificationHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NotificationHistoryDataItems0 notification history data items0
//
// swagger:model NotificationHistoryDataItems0
type NotificationHistoryDataItems0 struct {

	// action message
	ActionMessage string `json:"actionMessage,omitempty"`

	// action status
	ActionStatus string `json:"actionStatus,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// is deleted
	IsDeleted bool `json:"isDeleted"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this notification history data items0
func (m *NotificationHistoryDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notification history data items0 based on context it is used
func (m *NotificationHistoryDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationHistoryDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationHistoryDataItems0) UnmarshalBinary(b []byte) error {
	var res NotificationHistoryDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
