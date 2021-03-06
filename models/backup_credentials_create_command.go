// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupCredentialsCreateCommand backup credentials create command
//
// swagger:model BackupCredentialsCreateCommand
type BackupCredentialsCreateCommand struct {

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// s3 access key Id
	S3AccessKeyID string `json:"s3AccessKeyId,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 name
	S3Name string `json:"s3Name,omitempty"`

	// s3 region
	S3Region string `json:"s3Region,omitempty"`

	// s3 secret key
	S3SecretKey string `json:"s3SecretKey,omitempty"`
}

// Validate validates this backup credentials create command
func (m *BackupCredentialsCreateCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this backup credentials create command based on context it is used
func (m *BackupCredentialsCreateCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupCredentialsCreateCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupCredentialsCreateCommand) UnmarshalBinary(b []byte) error {
	var res BackupCredentialsCreateCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
