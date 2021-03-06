// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupCredentialsUpdateCommand backup credentials update command
//
// swagger:model BackupCredentialsUpdateCommand
type BackupCredentialsUpdateCommand struct {

	// id
	ID int32 `json:"id,omitempty"`

	// s3 access key Id
	S3AccessKeyID string `json:"s3AccessKeyId,omitempty"`

	// s3 name
	S3Name string `json:"s3Name,omitempty"`

	// s3 secret key
	S3SecretKey string `json:"s3SecretKey,omitempty"`
}

// Validate validates this backup credentials update command
func (m *BackupCredentialsUpdateCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup credentials update command based on context it is used
func (m *BackupCredentialsUpdateCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupCredentialsUpdateCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupCredentialsUpdateCommand) UnmarshalBinary(b []byte) error {
	var res BackupCredentialsUpdateCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
