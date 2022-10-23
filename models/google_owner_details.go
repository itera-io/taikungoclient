// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GoogleOwnerDetails google owner details
//
// swagger:model GoogleOwnerDetails
type GoogleOwnerDetails struct {

	// image
	Image *GoogleOwnerDetailsImage `json:"image,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`
}

// Validate validates this google owner details
func (m *GoogleOwnerDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GoogleOwnerDetails) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this google owner details based on the context it is used
func (m *GoogleOwnerDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GoogleOwnerDetails) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {
		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GoogleOwnerDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleOwnerDetails) UnmarshalBinary(b []byte) error {
	var res GoogleOwnerDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GoogleOwnerDetailsImage google owner details image
//
// swagger:model GoogleOwnerDetailsImage
type GoogleOwnerDetailsImage struct {

	// display name
	DisplayName string `json:"displayName"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this google owner details image
func (m *GoogleOwnerDetailsImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this google owner details image based on context it is used
func (m *GoogleOwnerDetailsImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GoogleOwnerDetailsImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GoogleOwnerDetailsImage) UnmarshalBinary(b []byte) error {
	var res GoogleOwnerDetailsImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
