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

// ShowbackRuleList showback rule list
//
// swagger:model ShowbackRuleList
type ShowbackRuleList struct {

	// data
	Data []*ShowbackRuleListDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this showback rule list
func (m *ShowbackRuleList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShowbackRuleList) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this showback rule list based on the context it is used
func (m *ShowbackRuleList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShowbackRuleList) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {
			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShowbackRuleList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowbackRuleList) UnmarshalBinary(b []byte) error {
	var res ShowbackRuleList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ShowbackRuleListDataItems0 showback rule list data items0
//
// swagger:model ShowbackRuleListDataItems0
type ShowbackRuleListDataItems0 struct {

	// billing start date
	BillingStartDate string `json:"billingStartDate,omitempty"`

	// by label
	ByLabel string `json:"byLabel,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// global alert limit
	GlobalAlertLimit int32 `json:"globalAlertLimit,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// labels
	Labels []*ShowbackRuleListDataItems0LabelsItems0 `json:"labels"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// metric name
	MetricName string `json:"metricName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// project alert limit
	ProjectAlertLimit int32 `json:"projectAlertLimit,omitempty"`

	// showback credential Id
	ShowbackCredentialID *int32 `json:"showbackCredentialId,omitempty"`

	// showback credential name
	ShowbackCredentialName string `json:"showbackCredentialName,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this showback rule list data items0
func (m *ShowbackRuleListDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShowbackRuleListDataItems0) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this showback rule list data items0 based on the context it is used
func (m *ShowbackRuleListDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShowbackRuleListDataItems0) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShowbackRuleListDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowbackRuleListDataItems0) UnmarshalBinary(b []byte) error {
	var res ShowbackRuleListDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ShowbackRuleListDataItems0LabelsItems0 showback rule list data items0 labels items0
//
// swagger:model ShowbackRuleListDataItems0LabelsItems0
type ShowbackRuleListDataItems0LabelsItems0 struct {

	// label
	Label string `json:"label,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this showback rule list data items0 labels items0
func (m *ShowbackRuleListDataItems0LabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this showback rule list data items0 labels items0 based on context it is used
func (m *ShowbackRuleListDataItems0LabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ShowbackRuleListDataItems0LabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShowbackRuleListDataItems0LabelsItems0) UnmarshalBinary(b []byte) error {
	var res ShowbackRuleListDataItems0LabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
