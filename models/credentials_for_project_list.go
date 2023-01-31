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

// CredentialsForProjectList credentials for project list
//
// swagger:model CredentialsForProjectList
type CredentialsForProjectList struct {

	// aws
	Aws *AwsCredentialsForProjectDto `json:"aws,omitempty"`

	// azure
	Azure *AzureCredentialsForProjectDto `json:"azure,omitempty"`

	// billing enabled
	BillingEnabled bool `json:"billingEnabled"`

	// cloud credential revision
	CloudCredentialRevision int32 `json:"cloudCredentialRevision,omitempty"`

	// cloud type
	CloudType CloudType `json:"cloudType,omitempty"`

	// continent name
	ContinentName string `json:"continentName,omitempty"`

	// google
	Google *GoogleCredentialForProjectDto `json:"google,omitempty"`

	// openstack
	Openstack *OpenstackCredentialsForProjectDto `json:"openstack,omitempty"`

	// requires v p n
	RequiresVPN bool `json:"requiresVPN"`

	// tanzu
	Tanzu *TanzuCredentialsForProjectDto `json:"tanzu,omitempty"`
}

// Validate validates this credentials for project list
func (m *CredentialsForProjectList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudType(formats); err != nil {
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

func (m *CredentialsForProjectList) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) validateCloudType(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudType) { // not required
		return nil
	}

	if err := m.CloudType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloudType")
		}
		return err
	}

	return nil
}

func (m *CredentialsForProjectList) validateGoogle(formats strfmt.Registry) error {
	if swag.IsZero(m.Google) { // not required
		return nil
	}

	if m.Google != nil {
		if err := m.Google.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("google")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("google")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) validateTanzu(formats strfmt.Registry) error {
	if swag.IsZero(m.Tanzu) { // not required
		return nil
	}

	if m.Tanzu != nil {
		if err := m.Tanzu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tanzu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tanzu")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this credentials for project list based on the context it is used
func (m *CredentialsForProjectList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudType(ctx, formats); err != nil {
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

func (m *CredentialsForProjectList) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {
		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) contextValidateCloudType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.CloudType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cloudType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("cloudType")
		}
		return err
	}

	return nil
}

func (m *CredentialsForProjectList) contextValidateGoogle(ctx context.Context, formats strfmt.Registry) error {

	if m.Google != nil {
		if err := m.Google.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("google")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("google")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	if m.Openstack != nil {
		if err := m.Openstack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *CredentialsForProjectList) contextValidateTanzu(ctx context.Context, formats strfmt.Registry) error {

	if m.Tanzu != nil {
		if err := m.Tanzu.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tanzu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tanzu")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CredentialsForProjectList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CredentialsForProjectList) UnmarshalBinary(b []byte) error {
	var res CredentialsForProjectList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
