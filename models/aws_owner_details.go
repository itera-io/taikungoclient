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

// AwsOwnerDetails aws owner details
//
// swagger:model AwsOwnerDetails
type AwsOwnerDetails struct {

	// image
	Image *AwsOwnerDetailsImage `json:"image,omitempty"`

	// owner Id
	OwnerID string `json:"ownerId,omitempty"`

	// owner name
	OwnerName string `json:"ownerName,omitempty"`
}

// Validate validates this aws owner details
func (m *AwsOwnerDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsOwnerDetails) validateImage(formats strfmt.Registry) error {
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

// ContextValidate validate this aws owner details based on the context it is used
func (m *AwsOwnerDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsOwnerDetails) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AwsOwnerDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsOwnerDetails) UnmarshalBinary(b []byte) error {
	var res AwsOwnerDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AwsOwnerDetailsImage aws owner details image
//
// swagger:model AwsOwnerDetailsImage
type AwsOwnerDetailsImage struct {

	// display name
	DisplayName string `json:"displayName"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this aws owner details image
func (m *AwsOwnerDetailsImage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws owner details image based on context it is used
func (m *AwsOwnerDetailsImage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsOwnerDetailsImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsOwnerDetailsImage) UnmarshalBinary(b []byte) error {
	var res AwsOwnerDetailsImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
