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

// CheckS3Command check s3 command
//
// swagger:model CheckS3Command
type CheckS3Command struct {

	// s3 access key Id
	// Required: true
	// Min Length: 1
	S3AccessKeyID *string `json:"s3AccessKeyId"`

	// s3 endpoint
	// Required: true
	// Min Length: 1
	S3Endpoint *string `json:"s3Endpoint"`

	// s3 region
	// Required: true
	// Min Length: 1
	S3Region *string `json:"s3Region"`

	// s3 secret key
	// Required: true
	// Min Length: 1
	S3SecretKey *string `json:"s3SecretKey"`
}

// Validate validates this check s3 command
func (m *CheckS3Command) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateS3AccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Endpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Region(formats); err != nil {
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

func (m *CheckS3Command) validateS3AccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("s3AccessKeyId", "body", m.S3AccessKeyID); err != nil {
		return err
	}

	if err := validate.MinLength("s3AccessKeyId", "body", *m.S3AccessKeyID, 1); err != nil {
		return err
	}

	return nil
}

func (m *CheckS3Command) validateS3Endpoint(formats strfmt.Registry) error {

	if err := validate.Required("s3Endpoint", "body", m.S3Endpoint); err != nil {
		return err
	}

	if err := validate.MinLength("s3Endpoint", "body", *m.S3Endpoint, 1); err != nil {
		return err
	}

	return nil
}

func (m *CheckS3Command) validateS3Region(formats strfmt.Registry) error {

	if err := validate.Required("s3Region", "body", m.S3Region); err != nil {
		return err
	}

	if err := validate.MinLength("s3Region", "body", *m.S3Region, 1); err != nil {
		return err
	}

	return nil
}

func (m *CheckS3Command) validateS3SecretKey(formats strfmt.Registry) error {

	if err := validate.Required("s3SecretKey", "body", m.S3SecretKey); err != nil {
		return err
	}

	if err := validate.MinLength("s3SecretKey", "body", *m.S3SecretKey, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this check s3 command based on context it is used
func (m *CheckS3Command) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CheckS3Command) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckS3Command) UnmarshalBinary(b []byte) error {
	var res CheckS3Command
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
