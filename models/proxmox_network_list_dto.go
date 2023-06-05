// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProxmoxNetworkListDto proxmox network list dto
//
// swagger:model ProxmoxNetworkListDto
type ProxmoxNetworkListDto struct {

	// begin allocation range
	BeginAllocationRange string `json:"beginAllocationRange,omitempty"`

	// bridge
	Bridge string `json:"bridge,omitempty"`

	// end allocation range
	EndAllocationRange string `json:"endAllocationRange,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// is private
	IsPrivate bool `json:"isPrivate"`

	// net mask
	NetMask int32 `json:"netMask,omitempty"`
}

// Validate validates this proxmox network list dto
func (m *ProxmoxNetworkListDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proxmox network list dto based on context it is used
func (m *ProxmoxNetworkListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProxmoxNetworkListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxmoxNetworkListDto) UnmarshalBinary(b []byte) error {
	var res ProxmoxNetworkListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
