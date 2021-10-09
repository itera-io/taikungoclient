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

// FilteringElement filtering element
//
// swagger:model FilteringElement
type FilteringElement int32

// for schema
var filteringElementEnum []interface{}

func init() {
	var res []FilteringElement
	if err := json.Unmarshal([]byte(`[1,2,3,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filteringElementEnum = append(filteringElementEnum, v)
	}
}

func (m FilteringElement) validateFilteringElementEnum(path, location string, value FilteringElement) error {
	if err := validate.EnumCase(path, location, value, filteringElementEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this filtering element
func (m FilteringElement) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFilteringElementEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this filtering element based on context it is used
func (m FilteringElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}