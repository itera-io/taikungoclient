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

// CredentialsChart credentials chart
//
// swagger:model CredentialsChart
type CredentialsChart struct {

	// amazon
	Amazon []*AmazonCredentialsListDto `json:"amazon"`

	// azure
	Azure []*AzureCredentialsListDto `json:"azure"`

	// google
	Google []*GoogleCredentialsListDto `json:"google"`

	// openstack
	Openstack []*OpenstackCredentialsListDto `json:"openstack"`

	// tanzu
	Tanzu []*TanzuCredentialsListDto `json:"tanzu"`

	// total count aws
	TotalCountAws int32 `json:"totalCountAws,omitempty"`

	// total count azure
	TotalCountAzure int32 `json:"totalCountAzure,omitempty"`

	// total count google
	TotalCountGoogle int32 `json:"totalCountGoogle,omitempty"`

	// total count openstack
	TotalCountOpenstack int32 `json:"totalCountOpenstack,omitempty"`

	// total count tanzu
	TotalCountTanzu int32 `json:"totalCountTanzu,omitempty"`
}

// Validate validates this credentials chart
func (m *CredentialsChart) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmazon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTanzu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsChart) validateAmazon(formats strfmt.Registry) error {
	if swag.IsZero(m.Amazon) { // not required
		return nil
	}

	for i := 0; i < len(m.Amazon); i++ {
		if swag.IsZero(m.Amazon[i]) { // not required
			continue
		}

		if m.Amazon[i] != nil {
			if err := m.Amazon[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amazon" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amazon" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	for i := 0; i < len(m.Azure); i++ {
		if swag.IsZero(m.Azure[i]) { // not required
			continue
		}

		if m.Azure[i] != nil {
			if err := m.Azure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) validateGoogle(formats strfmt.Registry) error {
	if swag.IsZero(m.Google) { // not required
		return nil
	}

	for i := 0; i < len(m.Google); i++ {
		if swag.IsZero(m.Google[i]) { // not required
			continue
		}

		if m.Google[i] != nil {
			if err := m.Google[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	for i := 0; i < len(m.Openstack); i++ {
		if swag.IsZero(m.Openstack[i]) { // not required
			continue
		}

		if m.Openstack[i] != nil {
			if err := m.Openstack[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openstack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openstack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) validateTanzu(formats strfmt.Registry) error {
	if swag.IsZero(m.Tanzu) { // not required
		return nil
	}

	for i := 0; i < len(m.Tanzu); i++ {
		if swag.IsZero(m.Tanzu[i]) { // not required
			continue
		}

		if m.Tanzu[i] != nil {
			if err := m.Tanzu[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tanzu" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tanzu" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this credentials chart based on the context it is used
func (m *CredentialsChart) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmazon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTanzu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CredentialsChart) contextValidateAmazon(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Amazon); i++ {

		if m.Amazon[i] != nil {
			if err := m.Amazon[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("amazon" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("amazon" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Azure); i++ {

		if m.Azure[i] != nil {
			if err := m.Azure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) contextValidateGoogle(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Google); i++ {

		if m.Google[i] != nil {
			if err := m.Google[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Openstack); i++ {

		if m.Openstack[i] != nil {
			if err := m.Openstack[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openstack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openstack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CredentialsChart) contextValidateTanzu(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tanzu); i++ {

		if m.Tanzu[i] != nil {
			if err := m.Tanzu[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tanzu" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tanzu" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsChart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsChart) UnmarshalBinary(b []byte) error {
	var res CredentialsChart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
