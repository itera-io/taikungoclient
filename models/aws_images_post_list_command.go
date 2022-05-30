// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsImagesPostListCommand aws images post list command
//
// swagger:model AwsImagesPostListCommand
type AwsImagesPostListCommand struct {

	// cloud Id
	CloudID int32 `json:"cloudId,omitempty"`

	// latest
	Latest bool `json:"latest"`

	// limit
	Limit int32 `json:"limit,omitempty"`

	// offset
	Offset int32 `json:"offset,omitempty"`

	// owners
	Owners []string `json:"owners"`

	// search
	Search string `json:"search,omitempty"`

	// sort by
	SortBy string `json:"sortBy,omitempty"`

	// sort direction
	SortDirection string `json:"sortDirection,omitempty"`
}

// Validate validates this aws images post list command
func (m *AwsImagesPostListCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws images post list command based on context it is used
func (m *AwsImagesPostListCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsImagesPostListCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsImagesPostListCommand) UnmarshalBinary(b []byte) error {
	var res AwsImagesPostListCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}