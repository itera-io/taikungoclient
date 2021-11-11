// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CloudStatus cloud status
//
// swagger:model CloudStatus
type CloudStatus int32

// for schema
var cloudStatusEnum []interface{}

func init() {
	var res []CloudStatus
	if err := json.Unmarshal([]byte(`[100,200,250,300,400,500,600,700,800]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudStatusEnum = append(cloudStatusEnum, v)
	}
}

func (m CloudStatus) validateCloudStatusEnum(path, location string, value CloudStatus) error {
	if err := validate.EnumCase(path, location, value, cloudStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cloud status
func (m CloudStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCloudStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cloud status based on context it is used
func (m CloudStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
