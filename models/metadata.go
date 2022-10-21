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

// Metadata metadata
//
// swagger:model Metadata
type Metadata struct {

	// infracost command
	InfracostCommand string `json:"infracostCommand,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// vcs branch
	VcsBranch string `json:"vcsBranch,omitempty"`

	// vcs commit author email
	VcsCommitAuthorEmail string `json:"vcsCommitAuthorEmail,omitempty"`

	// vcs commit author name
	VcsCommitAuthorName string `json:"vcsCommitAuthorName,omitempty"`

	// vcs commit message
	VcsCommitMessage string `json:"vcsCommitMessage,omitempty"`

	// vcs commit sha
	VcsCommitSha string `json:"vcsCommitSha,omitempty"`

	// vcs commit timestamp
	// Format: date-time
	VcsCommitTimestamp *strfmt.DateTime `json:"vcsCommitTimestamp,omitempty"`

	// vcs repository Url
	VcsRepositoryURL string `json:"vcsRepositoryUrl,omitempty"`

	// vcs sub path
	VcsSubPath string `json:"vcsSubPath,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcsCommitTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metadata) validateVcsCommitTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.VcsCommitTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("vcsCommitTimestamp", "body", "date-time", m.VcsCommitTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this metadata based on context it is used
func (m *Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
