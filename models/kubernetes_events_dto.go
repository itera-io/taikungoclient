// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesEventsDto kubernetes events dto
//
// swagger:model KubernetesEventsDto
type KubernetesEventsDto struct {

	// count
	Count int32 `json:"count,omitempty"`

	// first time stamp
	FirstTimeStamp string `json:"firstTimeStamp,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// involved object
	InvolvedObject interface{} `json:"involvedObject,omitempty"`

	// last time stamp
	LastTimeStamp string `json:"lastTimeStamp,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// source
	Source interface{} `json:"source,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this kubernetes events dto
func (m *KubernetesEventsDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kubernetes events dto based on context it is used
func (m *KubernetesEventsDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesEventsDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesEventsDto) UnmarshalBinary(b []byte) error {
	var res KubernetesEventsDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
