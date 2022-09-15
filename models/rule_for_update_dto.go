// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RuleForUpdateDto rule for update dto
//
// swagger:model RuleForUpdateDto
type RuleForUpdateDto struct {

	// labels to add
	LabelsToAdd []*PrometheusLabelListDto `json:"labelsToAdd"`

	// labels to delete
	LabelsToDelete []*PrometheusLabelDeleteDto `json:"labelsToDelete"`

	// labels to update
	LabelsToUpdate []*PrometheusLabelUpdateDto `json:"labelsToUpdate"`

	// metric name
	MetricName string `json:"metricName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operation credential Id
	OperationCredentialID int32 `json:"operationCredentialId,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// rule discount rate
	RuleDiscountRate int32 `json:"ruleDiscountRate"`

	// type
	Type PrometheusType `json:"type,omitempty"`
}

// Validate validates this rule for update dto
func (m *RuleForUpdateDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelsToAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsToDelete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsToUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleForUpdateDto) validateLabelsToAdd(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsToAdd) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelsToAdd); i++ {
		if swag.IsZero(m.LabelsToAdd[i]) { // not required
			continue
		}

		if m.LabelsToAdd[i] != nil {
			if err := m.LabelsToAdd[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToAdd" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToAdd" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) validateLabelsToDelete(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsToDelete) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelsToDelete); i++ {
		if swag.IsZero(m.LabelsToDelete[i]) { // not required
			continue
		}

		if m.LabelsToDelete[i] != nil {
			if err := m.LabelsToDelete[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToDelete" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToDelete" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) validateLabelsToUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsToUpdate) { // not required
		return nil
	}

	for i := 0; i < len(m.LabelsToUpdate); i++ {
		if swag.IsZero(m.LabelsToUpdate[i]) { // not required
			continue
		}

		if m.LabelsToUpdate[i] != nil {
			if err := m.LabelsToUpdate[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToUpdate" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToUpdate" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this rule for update dto based on the context it is used
func (m *RuleForUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabelsToAdd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsToDelete(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsToUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleForUpdateDto) contextValidateLabelsToAdd(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelsToAdd); i++ {

		if m.LabelsToAdd[i] != nil {
			if err := m.LabelsToAdd[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToAdd" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToAdd" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) contextValidateLabelsToDelete(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelsToDelete); i++ {

		if m.LabelsToDelete[i] != nil {
			if err := m.LabelsToDelete[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToDelete" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToDelete" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) contextValidateLabelsToUpdate(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LabelsToUpdate); i++ {

		if m.LabelsToUpdate[i] != nil {
			if err := m.LabelsToUpdate[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labelsToUpdate" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labelsToUpdate" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RuleForUpdateDto) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleForUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleForUpdateDto) UnmarshalBinary(b []byte) error {
	var res RuleForUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
