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

// OpenstackQuotaList openstack quota list
//
// swagger:model OpenstackQuotaList
type OpenstackQuotaList struct {

	// compute
	Compute *OpenstackQuotaListCompute `json:"compute,omitempty"`

	// network
	Network *OpenstackQuotaListNetwork `json:"network,omitempty"`

	// volume
	Volume *OpenstackQuotaListVolume `json:"volume,omitempty"`
}

// Validate validates this openstack quota list
func (m *OpenstackQuotaList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackQuotaList) validateCompute(formats strfmt.Registry) error {
	if swag.IsZero(m.Compute) { // not required
		return nil
	}

	if m.Compute != nil {
		if err := m.Compute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compute")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackQuotaList) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackQuotaList) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this openstack quota list based on the context it is used
func (m *OpenstackQuotaList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompute(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackQuotaList) contextValidateCompute(ctx context.Context, formats strfmt.Registry) error {

	if m.Compute != nil {
		if err := m.Compute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("compute")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackQuotaList) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if m.Network != nil {
		if err := m.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackQuotaList) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackQuotaList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackQuotaList) UnmarshalBinary(b []byte) error {
	var res OpenstackQuotaList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenstackQuotaListCompute openstack quota list compute
//
// swagger:model OpenstackQuotaListCompute
type OpenstackQuotaListCompute struct {

	// max server groups
	MaxServerGroups int64 `json:"maxServerGroups,omitempty"`

	// max total cores
	MaxTotalCores int64 `json:"maxTotalCores,omitempty"`

	// max total instances
	MaxTotalInstances int64 `json:"maxTotalInstances,omitempty"`

	// max total Ram size
	MaxTotalRAMSize int64 `json:"maxTotalRamSize,omitempty"`

	// used Cpu size
	UsedCPUSize int64 `json:"usedCpuSize,omitempty"`

	// used instance size
	UsedInstanceSize int64 `json:"usedInstanceSize,omitempty"`

	// used Ram size
	UsedRAMSize int64 `json:"usedRamSize,omitempty"`

	// used server groups
	UsedServerGroups int64 `json:"usedServerGroups,omitempty"`
}

// Validate validates this openstack quota list compute
func (m *OpenstackQuotaListCompute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack quota list compute based on context it is used
func (m *OpenstackQuotaListCompute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackQuotaListCompute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackQuotaListCompute) UnmarshalBinary(b []byte) error {
	var res OpenstackQuotaListCompute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenstackQuotaListNetwork openstack quota list network
//
// swagger:model OpenstackQuotaListNetwork
type OpenstackQuotaListNetwork struct {

	// floating Ip limit
	FloatingIPLimit int64 `json:"floatingIpLimit,omitempty"`

	// floating Ip used
	FloatingIPUsed int64 `json:"floatingIpUsed,omitempty"`

	// network limit
	NetworkLimit int64 `json:"networkLimit,omitempty"`

	// network used
	NetworkUsed int64 `json:"networkUsed,omitempty"`

	// port limit
	PortLimit int64 `json:"portLimit,omitempty"`

	// port used
	PortUsed int64 `json:"portUsed,omitempty"`

	// router limit
	RouterLimit int64 `json:"routerLimit,omitempty"`

	// router used
	RouterUsed int64 `json:"routerUsed,omitempty"`

	// security group limit
	SecurityGroupLimit int64 `json:"securityGroupLimit,omitempty"`

	// security group rule limit
	SecurityGroupRuleLimit int64 `json:"securityGroupRuleLimit,omitempty"`

	// security group rule used
	SecurityGroupRuleUsed int64 `json:"securityGroupRuleUsed,omitempty"`

	// security group used
	SecurityGroupUsed int64 `json:"securityGroupUsed,omitempty"`

	// subnet limit
	SubnetLimit int64 `json:"subnetLimit,omitempty"`

	// subnet used
	SubnetUsed int64 `json:"subnetUsed,omitempty"`
}

// Validate validates this openstack quota list network
func (m *OpenstackQuotaListNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack quota list network based on context it is used
func (m *OpenstackQuotaListNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackQuotaListNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackQuotaListNetwork) UnmarshalBinary(b []byte) error {
	var res OpenstackQuotaListNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OpenstackQuotaListVolume openstack quota list volume
//
// swagger:model OpenstackQuotaListVolume
type OpenstackQuotaListVolume struct {

	// count volume size
	CountVolumeSize int64 `json:"countVolumeSize,omitempty"`

	// max count volume size
	MaxCountVolumeSize int64 `json:"maxCountVolumeSize,omitempty"`

	// max total volume size
	MaxTotalVolumeSize int64 `json:"maxTotalVolumeSize,omitempty"`

	// used volume size
	UsedVolumeSize int64 `json:"usedVolumeSize,omitempty"`
}

// Validate validates this openstack quota list volume
func (m *OpenstackQuotaListVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this openstack quota list volume based on context it is used
func (m *OpenstackQuotaListVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackQuotaListVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackQuotaListVolume) UnmarshalBinary(b []byte) error {
	var res OpenstackQuotaListVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
