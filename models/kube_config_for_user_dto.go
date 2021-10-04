// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubeConfigForUserDto kube config for user dto
//
// swagger:model KubeConfigForUserDto
type KubeConfigForUserDto struct {

	// id
	ID int32 `json:"id,omitempty"`

	// is accessible for all
	IsAccessibleForAll bool `json:"isAccessibleForAll,omitempty"`

	// is accessible for manager
	IsAccessibleForManager bool `json:"isAccessibleForManager,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// service account name
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`

	// user role
	UserRole string `json:"userRole,omitempty"`
}

// Validate validates this kube config for user dto
func (m *KubeConfigForUserDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kube config for user dto based on context it is used
func (m *KubeConfigForUserDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeConfigForUserDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeConfigForUserDto) UnmarshalBinary(b []byte) error {
	var res KubeConfigForUserDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
