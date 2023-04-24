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

// BackupCredentialsUpdateCommand backup credentials update command
//
// swagger:model BackupCredentialsUpdateCommand
type BackupCredentialsUpdateCommand struct {

	// id
	// Required: true
	// Minimum: > 0
	ID *int32 `json:"id"`

	// s3 access key Id
	// Required: true
	// Min Length: 1
	S3AccessKeyID *string `json:"s3AccessKeyId"`

	// s3 name
	// Required: true
	// Max Length: 30
	// Min Length: 3
	S3Name *string `json:"s3Name"`

	// s3 secret key
	// Required: true
	// Min Length: 1
	S3SecretKey *string `json:"s3SecretKey"`
}

// Validate validates this backup credentials update command
func (m *BackupCredentialsUpdateCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3AccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Name(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3SecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupCredentialsUpdateCommand) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, true); err != nil {
		return err
	}

	return nil
}

func (m *BackupCredentialsUpdateCommand) validateS3AccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("s3AccessKeyId", "body", m.S3AccessKeyID); err != nil {
		return err
	}

	if err := validate.MinLength("s3AccessKeyId", "body", *m.S3AccessKeyID, 1); err != nil {
		return err
	}

	return nil
}

func (m *BackupCredentialsUpdateCommand) validateS3Name(formats strfmt.Registry) error {

	if err := validate.Required("s3Name", "body", m.S3Name); err != nil {
		return err
	}

	if err := validate.MinLength("s3Name", "body", *m.S3Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("s3Name", "body", *m.S3Name, 30); err != nil {
		return err
	}

	return nil
}

func (m *BackupCredentialsUpdateCommand) validateS3SecretKey(formats strfmt.Registry) error {

	if err := validate.Required("s3SecretKey", "body", m.S3SecretKey); err != nil {
		return err
	}

	if err := validate.MinLength("s3SecretKey", "body", *m.S3SecretKey, 1); err != nil {
		return err
	}

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
