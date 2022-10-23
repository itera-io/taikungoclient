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

// AvailablePackagesList available packages list
//
// swagger:model AvailablePackagesList
type AvailablePackagesList struct {

	// data
	Data []*AvailablePackagesListDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this available packages list
func (m *AvailablePackagesList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePackagesList) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this available packages list based on the context it is used
func (m *AvailablePackagesList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailablePackagesList) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AvailablePackagesList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesList) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailablePackagesListDataItems0 available packages list data items0
//
// swagger:model AvailablePackagesListDataItems0
type AvailablePackagesListDataItems0 struct {

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
	Repository *AvailablePackagesListDataItems0Repository `json:"repository,omitempty"`

	// security report summary
	SecurityReportSummary *AvailablePackagesListDataItems0SecurityReportSummary `json:"securityReportSummary,omitempty"`

	// signed
	Signed bool `json:"signed"`

	// stars
	Stars int64 `json:"stars,omitempty"`

	// ts
	Ts string `json:"ts,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this available packages list data items0
func (m *AvailablePackagesListDataItems0) Validate(formats strfmt.Registry) error {
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

func (m *AvailablePackagesListDataItems0) validateRepository(formats strfmt.Registry) error {
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

func (m *AvailablePackagesListDataItems0) validateSecurityReportSummary(formats strfmt.Registry) error {
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

// ContextValidate validate this available packages list data items0 based on the context it is used
func (m *AvailablePackagesListDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *AvailablePackagesListDataItems0) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AvailablePackagesListDataItems0) contextValidateSecurityReportSummary(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AvailablePackagesListDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesListDataItems0) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesListDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailablePackagesListDataItems0Repository available packages list data items0 repository
//
// swagger:model AvailablePackagesListDataItems0Repository
type AvailablePackagesListDataItems0Repository struct {

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

// Validate validates this available packages list data items0 repository
func (m *AvailablePackagesListDataItems0Repository) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available packages list data items0 repository based on context it is used
func (m *AvailablePackagesListDataItems0Repository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePackagesListDataItems0Repository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesListDataItems0Repository) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesListDataItems0Repository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AvailablePackagesListDataItems0SecurityReportSummary available packages list data items0 security report summary
//
// swagger:model AvailablePackagesListDataItems0SecurityReportSummary
type AvailablePackagesListDataItems0SecurityReportSummary struct {

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

// Validate validates this available packages list data items0 security report summary
func (m *AvailablePackagesListDataItems0SecurityReportSummary) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this available packages list data items0 security report summary based on context it is used
func (m *AvailablePackagesListDataItems0SecurityReportSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailablePackagesListDataItems0SecurityReportSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailablePackagesListDataItems0SecurityReportSummary) UnmarshalBinary(b []byte) error {
	var res AvailablePackagesListDataItems0SecurityReportSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
