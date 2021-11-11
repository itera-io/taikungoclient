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

// SlackType slack type
//
// swagger:model SlackType
type SlackType int32

// for schema
var slackTypeEnum []interface{}

func init() {
	var res []SlackType
	if err := json.Unmarshal([]byte(`[100,200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		slackTypeEnum = append(slackTypeEnum, v)
	}
}

func (m SlackType) validateSlackTypeEnum(path, location string, value SlackType) error {
	if err := validate.EnumCase(path, location, value, slackTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this slack type
func (m SlackType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSlackTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this slack type based on context it is used
func (m SlackType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
