// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectQuotaUpdateDto project quota update dto
//
// swagger:model ProjectQuotaUpdateDto
type ProjectQuotaUpdateDto struct {

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// disk size
	DiskSize int64 `json:"diskSize,omitempty"`

	// is Cpu unlimited
	IsCPUUnlimited bool `json:"isCpuUnlimited"`

	// is disk size unlimited
	IsDiskSizeUnlimited bool `json:"isDiskSizeUnlimited"`

	// is Ram unlimited
	IsRAMUnlimited bool `json:"isRamUnlimited"`

	// ram
	RAM int64 `json:"ram,omitempty"`
}

// Validate validates this project quota update dto
func (m *ProjectQuotaUpdateDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project quota update dto based on context it is used
func (m *ProjectQuotaUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectQuotaUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectQuotaUpdateDto) UnmarshalBinary(b []byte) error {
	var res ProjectQuotaUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
