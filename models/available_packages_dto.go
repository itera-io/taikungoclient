// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AvailablePackagesDto available packages dto
//
// swagger:model AvailablePackagesDto
type AvailablePackagesDto struct {

	// app version
	AppVersion string `json:"appVersion,omitempty"`

	// catalog app Id
	CatalogAppID int32 `json:"catalogAppId,omitempty"`

	// deprecated
	Deprecated bool `json:"deprecated"`

	// description
	Description string `json:"description,omitempty"`

	// installed instance count
	InstalledInstanceCount int32 `json:"installedInstanceCount,omitempty"`

	// logo image Id
	LogoImageID string `json:"logoImageId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// normalized name
	NormalizedName string `json:"normalizedName,omitempty"`

	// package Id
	PackageID string `json:"packageId,omitempty"`

	// repository
	Repository *AvailablePackagesDtoRepository `json:"repository,omitempty"`

	// security report summary
	SecurityReportSummary *AvailablePackagesDtoSecurityReportSummary `json:"securityReportSummary,omitempty"`

	// signed
	Signed bool `json:"signed"`

	// stars
	Stars int64 `json:"stars,omitempty"`

	// ts
	Ts string `json:"ts,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this available packages dto
func (m *AvailablePackagesDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityReportSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePackagesDto) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *AvailablePackagesDto) validateSecurityReportSummary(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityReportSummary) { // not required
		return nil
	}

	if m.SecurityReportSummary != nil {
		if err := m.SecurityReportSummary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityReportSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityReportSummary")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this available packages dto based on the context it is used
func (m *AvailablePackagesDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityReportSummary(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePackagesDto) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {
		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *AvailablePackagesDto) contextValidateSecurityReportSummary(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityReportSummary != nil {
		if err := m.SecurityReportSummary.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityReportSummary")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityReportSummary")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePackagesDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesDto) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailablePackagesDtoRepository available packages dto repository
//
// swagger:model AvailablePackagesDtoRepository
type AvailablePackagesDtoRepository struct {

	// kind
	Kind int64 `json:"kind,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// official
	Official bool `json:"official"`

	// organization display name
	OrganizationDisplayName string `json:"organizationDisplayName,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// repository Id
	RepositoryID string `json:"repositoryId,omitempty"`

	// scanner disabled
	ScannerDisabled bool `json:"scannerDisabled"`

	// url
	URL string `json:"url,omitempty"`

	// verified publisher
	VerifiedPublisher bool `json:"verifiedPublisher"`
}

// Validate validates this available packages dto repository
func (m *AvailablePackagesDtoRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available packages dto repository based on context it is used
func (m *AvailablePackagesDtoRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePackagesDtoRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesDtoRepository) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesDtoRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailablePackagesDtoSecurityReportSummary available packages dto security report summary
//
// swagger:model AvailablePackagesDtoSecurityReportSummary
type AvailablePackagesDtoSecurityReportSummary struct {

	// critical
	Critical int64 `json:"critical,omitempty"`

	// high
	High int64 `json:"high,omitempty"`

	// low
	Low int64 `json:"low,omitempty"`

	// medium
	Medium int64 `json:"medium,omitempty"`

	// unknown
	Unknown int64 `json:"unknown,omitempty"`
}

// Validate validates this available packages dto security report summary
func (m *AvailablePackagesDtoSecurityReportSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available packages dto security report summary based on context it is used
func (m *AvailablePackagesDtoSecurityReportSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePackagesDtoSecurityReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesDtoSecurityReportSummary) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesDtoSecurityReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
