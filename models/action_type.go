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

// ActionType action type
//
// swagger:model ActionType
type ActionType int32

// for schema
var actionTypeEnum []interface{}

func init() {
	var res []ActionType
	if err := json.Unmarshal([]byte(`[10,20,30,32,34,36,38,40,42,44,50,60,70,80,90,100,105,106,107,108,109,110,120,130,140,150,155,156,160,165,170,175,176,177,180,190,200,210,220,230,240,250,260,270,280,290,300,310,320,330,340,350,360,370,380,390,400,410]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionTypeEnum = append(actionTypeEnum, v)
	}
}

func (m ActionType) validateActionTypeEnum(path, location string, value ActionType) error {
	if err := validate.EnumCase(path, location, value, actionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this action type
func (m ActionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateActionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this action type based on context it is used
func (m ActionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
