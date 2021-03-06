// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SlackConfigurationDto slack configuration dto
//
// swagger:model SlackConfigurationDto
type SlackConfigurationDto struct {

	// channel
	Channel string `json:"channel,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// slack type
	SlackType string `json:"slackType,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this slack configuration dto
func (m *SlackConfigurationDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this slack configuration dto based on context it is used
func (m *SlackConfigurationDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SlackConfigurationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SlackConfigurationDto) UnmarshalBinary(b []byte) error {
	var res SlackConfigurationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
