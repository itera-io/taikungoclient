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

// AccessProfilesForProjectListDto access profiles for project list dto
//
// swagger:model AccessProfilesForProjectListDto
type AccessProfilesForProjectListDto struct {

	// dns servers
	DNSServers []*DNSServerListDto `json:"dnsServers"`

	// http proxy
	HTTPProxy string `json:"httpProxy,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ntp servers
	NtpServers []*NtpServerListDto `json:"ntpServers"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// revision
	Revision int32 `json:"revision,omitempty"`

	// ssh users
	SSHUsers []*SSHUserListDto `json:"sshUsers"`
}

// Validate validates this access profiles for project list dto
func (m *AccessProfilesForProjectListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNtpServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSSHUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessProfilesForProjectListDto) validateDNSServers(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSServers) { // not required
		return nil
	}

	for i := 0; i < len(m.DNSServers); i++ {
		if swag.IsZero(m.DNSServers[i]) { // not required
			continue
		}

		if m.DNSServers[i] != nil {
			if err := m.DNSServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccessProfilesForProjectListDto) validateNtpServers(formats strfmt.Registry) error {
	if swag.IsZero(m.NtpServers) { // not required
		return nil
	}

	for i := 0; i < len(m.NtpServers); i++ {
		if swag.IsZero(m.NtpServers[i]) { // not required
			continue
		}

		if m.NtpServers[i] != nil {
			if err := m.NtpServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntpServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntpServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccessProfilesForProjectListDto) validateSSHUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.SSHUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.SSHUsers); i++ {
		if swag.IsZero(m.SSHUsers[i]) { // not required
			continue
		}

		if m.SSHUsers[i] != nil {
			if err := m.SSHUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sshUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this access profiles for project list dto based on the context it is used
func (m *AccessProfilesForProjectListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNSServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNtpServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSSHUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessProfilesForProjectListDto) contextValidateDNSServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNSServers); i++ {

		if m.DNSServers[i] != nil {
			if err := m.DNSServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dnsServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccessProfilesForProjectListDto) contextValidateNtpServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NtpServers); i++ {

		if m.NtpServers[i] != nil {
			if err := m.NtpServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ntpServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ntpServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccessProfilesForProjectListDto) contextValidateSSHUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SSHUsers); i++ {

		if m.SSHUsers[i] != nil {
			if err := m.SSHUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sshUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessProfilesForProjectListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessProfilesForProjectListDto) UnmarshalBinary(b []byte) error {
	var res AccessProfilesForProjectListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
