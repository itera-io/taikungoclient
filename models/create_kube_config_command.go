// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateKubeConfigCommand create kube config command
//
// swagger:model CreateKubeConfigCommand
type CreateKubeConfigCommand struct {

	// is accessible for all
	IsAccessibleForAll bool `json:"isAccessibleForAll"`

	// is accessible for manager
	IsAccessibleForManager bool `json:"isAccessibleForManager"`

	// kube config role Id
	KubeConfigRoleID int32 `json:"kubeConfigRoleId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// ttl
	TTL int32 `json:"ttl,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this create kube config command
func (m *CreateKubeConfigCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create kube config command based on context it is used
func (m *CreateKubeConfigCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateKubeConfigCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateKubeConfigCommand) UnmarshalBinary(b []byte) error {
	var res CreateKubeConfigCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
