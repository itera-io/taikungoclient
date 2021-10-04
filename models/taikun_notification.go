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

// TaikunNotification taikun notification
//
// swagger:model TaikunNotification
type TaikunNotification struct {

	// action message
	ActionMessage string `json:"actionMessage,omitempty"`

	// action status
	ActionStatus string `json:"actionStatus,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// estimated time
	EstimatedTime string `json:"estimatedTime,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is deleted
	IsDeleted bool `json:"isDeleted,omitempty"`

	// is seen
	IsSeen bool `json:"isSeen,omitempty"`

	// is unread
	IsUnread bool `json:"isUnread,omitempty"`

	// kubernetes start date
	KubernetesStartDate string `json:"kubernetesStartDate,omitempty"`

	// last modified
	// Format: date-time
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// operation project Id
	OperationProjectID int32 `json:"operationProjectId,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// ticket Id
	TicketID string `json:"ticketId,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this taikun notification
func (m *TaikunNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaikunNotification) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaikunNotification) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this taikun notification based on context it is used
func (m *TaikunNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaikunNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaikunNotification) UnmarshalBinary(b []byte) error {
	var res TaikunNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
