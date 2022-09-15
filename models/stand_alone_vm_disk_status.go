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

// StandAloneVMDiskStatus stand alone Vm disk status
//
// swagger:model StandAloneVmDiskStatus
type StandAloneVMDiskStatus int32

// for schema
var standAloneVmDiskStatusEnum []interface{}

func init() {
	var res []StandAloneVMDiskStatus
	if err := json.Unmarshal([]byte(`[100,200,300,400,500,600]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		standAloneVmDiskStatusEnum = append(standAloneVmDiskStatusEnum, v)
	}
}

func (m StandAloneVMDiskStatus) validateStandAloneVMDiskStatusEnum(path, location string, value StandAloneVMDiskStatus) error {
	if err := validate.EnumCase(path, location, value, standAloneVmDiskStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this stand alone Vm disk status
func (m StandAloneVMDiskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStandAloneVMDiskStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this stand alone Vm disk status based on context it is used
func (m StandAloneVMDiskStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
