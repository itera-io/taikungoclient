// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DownloadKubeConfigCommand download kube config command
//
// swagger:model DownloadKubeConfigCommand
type DownloadKubeConfigCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`
}

// Validate validates this download kube config command
func (m *DownloadKubeConfigCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this download kube config command based on context it is used
func (m *DownloadKubeConfigCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadKubeConfigCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadKubeConfigCommand) UnmarshalBinary(b []byte) error {
	var res DownloadKubeConfigCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
