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

// CreateProxmoxCommand create proxmox command
//
// swagger:model CreateProxmoxCommand
type CreateProxmoxCommand struct {

	// continent
	Continent string `json:"continent,omitempty"`

	// hypervisors
	Hypervisors []string `json:"hypervisors"`

	// name
	// Required: true
	// Max Length: 30
	// Min Length: 3
	Name *string `json:"name"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// private network
	PrivateNetwork *CreateProxmoxNetworkDto `json:"privateNetwork,omitempty"`

	// public network
	PublicNetwork *CreateProxmoxNetworkDto `json:"publicNetwork,omitempty"`

	// storage
	// Required: true
	// Min Length: 1
	Storage *string `json:"storage"`

	// token Id
	// Required: true
	// Min Length: 1
	TokenID *string `json:"tokenId"`

	// token secret
	// Required: true
	// Min Length: 1
	TokenSecret *string `json:"tokenSecret"`

	// url
	// Required: true
	// Min Length: 1
	URL *string `json:"url"`

	// vm template name
	// Required: true
	// Min Length: 1
	VMTemplateName *string `json:"vmTemplateName"`
}

// Validate validates this create proxmox command
func (m *CreateProxmoxCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProxmoxCommand) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 30); err != nil {
		return err
	}

	return nil
}

func (m *CreateProxmoxCommand) validatePrivateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivateNetwork) { // not required
		return nil
	}

	if m.PrivateNetwork != nil {
		if err := m.PrivateNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privateNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privateNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *CreateProxmoxCommand) validatePublicNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicNetwork) { // not required
		return nil
	}

	if m.PublicNetwork != nil {
		if err := m.PublicNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *CreateProxmoxCommand) validateStorage(formats strfmt.Registry) error {

	if err := validate.Required("storage", "body", m.Storage); err != nil {
		return err
	}

	if err := validate.MinLength("storage", "body", *m.Storage, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateProxmoxCommand) validateTokenID(formats strfmt.Registry) error {

	if err := validate.Required("tokenId", "body", m.TokenID); err != nil {
		return err
	}

	if err := validate.MinLength("tokenId", "body", *m.TokenID, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateProxmoxCommand) validateTokenSecret(formats strfmt.Registry) error {

	if err := validate.Required("tokenSecret", "body", m.TokenSecret); err != nil {
		return err
	}

	if err := validate.MinLength("tokenSecret", "body", *m.TokenSecret, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateProxmoxCommand) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", *m.URL, 1); err != nil {
		return err
	}

	return nil
}

func (m *CreateProxmoxCommand) validateVMTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("vmTemplateName", "body", m.VMTemplateName); err != nil {
		return err
	}

	if err := validate.MinLength("vmTemplateName", "body", *m.VMTemplateName, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create proxmox command based on the context it is used
func (m *CreateProxmoxCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrivateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateProxmoxCommand) contextValidatePrivateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.PrivateNetwork != nil {
		if err := m.PrivateNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privateNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privateNetwork")
			}
			return err
		}
	}

	return nil
}

func (m *CreateProxmoxCommand) contextValidatePublicNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicNetwork != nil {
		if err := m.PublicNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publicNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publicNetwork")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateProxmoxCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateProxmoxCommand) UnmarshalBinary(b []byte) error {
	var res CreateProxmoxCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
