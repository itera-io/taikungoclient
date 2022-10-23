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

// ServerChartDto server chart dto
//
// swagger:model ServerChartDto
type ServerChartDto struct {

	// aws
	Aws []*ServerChartDtoAwsItems0 `json:"aws"`

	// azure
	Azure []*ServerChartDtoAzureItems0 `json:"azure"`

	// failed
	Failed []*ServerChartDtoFailedItems0 `json:"failed"`

	// google
	Google []*ServerChartDtoGoogleItems0 `json:"google"`

	// openstack
	Openstack []*ServerChartDtoOpenstackItems0 `json:"openstack"`

	// succeeded
	Succeeded []*ServerChartDtoSucceededItems0 `json:"succeeded"`

	// total aws count
	TotalAwsCount int32 `json:"totalAwsCount,omitempty"`

	// total azure count
	TotalAzureCount int32 `json:"totalAzureCount,omitempty"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`

	// total Cpu
	TotalCPU int32 `json:"totalCpu,omitempty"`

	// total disk size
	TotalDiskSize int64 `json:"totalDiskSize,omitempty"`

	// total failed count
	TotalFailedCount int32 `json:"totalFailedCount,omitempty"`

	// total google count
	TotalGoogleCount int32 `json:"totalGoogleCount,omitempty"`

	// total openstack count
	TotalOpenstackCount int32 `json:"totalOpenstackCount,omitempty"`

	// total pending count
	TotalPendingCount int32 `json:"totalPendingCount,omitempty"`

	// total Ram
	TotalRAM int64 `json:"totalRam,omitempty"`

	// total succeeded count
	TotalSucceededCount int32 `json:"totalSucceededCount,omitempty"`

	// total updating count
	TotalUpdatingCount int32 `json:"totalUpdatingCount,omitempty"`

	// updating
	Updating []*ServerChartDtoUpdatingItems0 `json:"updating"`

	// used resources
	UsedResources []*ServerChartDtoUsedResourcesItems0 `json:"usedResources"`

	// waiting
	Waiting []*ServerChartDtoWaitingItems0 `json:"waiting"`
}

// Validate validates this server chart dto
func (m *ServerChartDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSucceeded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdating(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaiting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerChartDto) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	for i := 0; i < len(m.Aws); i++ {
		if swag.IsZero(m.Aws[i]) { // not required
			continue
		}

		if m.Aws[i] != nil {
			if err := m.Aws[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aws" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aws" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	for i := 0; i < len(m.Azure); i++ {
		if swag.IsZero(m.Azure[i]) { // not required
			continue
		}

		if m.Azure[i] != nil {
			if err := m.Azure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateFailed(formats strfmt.Registry) error {
	if swag.IsZero(m.Failed) { // not required
		return nil
	}

	for i := 0; i < len(m.Failed); i++ {
		if swag.IsZero(m.Failed[i]) { // not required
			continue
		}

		if m.Failed[i] != nil {
			if err := m.Failed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateGoogle(formats strfmt.Registry) error {
	if swag.IsZero(m.Google) { // not required
		return nil
	}

	for i := 0; i < len(m.Google); i++ {
		if swag.IsZero(m.Google[i]) { // not required
			continue
		}

		if m.Google[i] != nil {
			if err := m.Google[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	for i := 0; i < len(m.Openstack); i++ {
		if swag.IsZero(m.Openstack[i]) { // not required
			continue
		}

		if m.Openstack[i] != nil {
			if err := m.Openstack[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openstack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openstack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateSucceeded(formats strfmt.Registry) error {
	if swag.IsZero(m.Succeeded) { // not required
		return nil
	}

	for i := 0; i < len(m.Succeeded); i++ {
		if swag.IsZero(m.Succeeded[i]) { // not required
			continue
		}

		if m.Succeeded[i] != nil {
			if err := m.Succeeded[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("succeeded" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("succeeded" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateUpdating(formats strfmt.Registry) error {
	if swag.IsZero(m.Updating) { // not required
		return nil
	}

	for i := 0; i < len(m.Updating); i++ {
		if swag.IsZero(m.Updating[i]) { // not required
			continue
		}

		if m.Updating[i] != nil {
			if err := m.Updating[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updating" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updating" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateUsedResources(formats strfmt.Registry) error {
	if swag.IsZero(m.UsedResources) { // not required
		return nil
	}

	for i := 0; i < len(m.UsedResources); i++ {
		if swag.IsZero(m.UsedResources[i]) { // not required
			continue
		}

		if m.UsedResources[i] != nil {
			if err := m.UsedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) validateWaiting(formats strfmt.Registry) error {
	if swag.IsZero(m.Waiting) { // not required
		return nil
	}

	for i := 0; i < len(m.Waiting); i++ {
		if swag.IsZero(m.Waiting[i]) { // not required
			continue
		}

		if m.Waiting[i] != nil {
			if err := m.Waiting[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("waiting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("waiting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this server chart dto based on the context it is used
func (m *ServerChartDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSucceeded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdating(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWaiting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerChartDto) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aws); i++ {

		if m.Aws[i] != nil {
			if err := m.Aws[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aws" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aws" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Azure); i++ {

		if m.Azure[i] != nil {
			if err := m.Azure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateFailed(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Failed); i++ {

		if m.Failed[i] != nil {
			if err := m.Failed[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateGoogle(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Google); i++ {

		if m.Google[i] != nil {
			if err := m.Google[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("google" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("google" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Openstack); i++ {

		if m.Openstack[i] != nil {
			if err := m.Openstack[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openstack" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openstack" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateSucceeded(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Succeeded); i++ {

		if m.Succeeded[i] != nil {
			if err := m.Succeeded[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("succeeded" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("succeeded" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateUpdating(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Updating); i++ {

		if m.Updating[i] != nil {
			if err := m.Updating[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updating" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updating" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateUsedResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UsedResources); i++ {

		if m.UsedResources[i] != nil {
			if err := m.UsedResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usedResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usedResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerChartDto) contextValidateWaiting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Waiting); i++ {

		if m.Waiting[i] != nil {
			if err := m.Waiting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("waiting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("waiting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDto) UnmarshalBinary(b []byte) error {
	var res ServerChartDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoAwsItems0 server chart dto aws items0
//
// swagger:model ServerChartDtoAwsItems0
type ServerChartDtoAwsItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto aws items0
func (m *ServerChartDtoAwsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto aws items0 based on context it is used
func (m *ServerChartDtoAwsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoAwsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoAwsItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoAwsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoAzureItems0 server chart dto azure items0
//
// swagger:model ServerChartDtoAzureItems0
type ServerChartDtoAzureItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto azure items0
func (m *ServerChartDtoAzureItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto azure items0 based on context it is used
func (m *ServerChartDtoAzureItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoAzureItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoAzureItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoAzureItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoFailedItems0 server chart dto failed items0
//
// swagger:model ServerChartDtoFailedItems0
type ServerChartDtoFailedItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto failed items0
func (m *ServerChartDtoFailedItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto failed items0 based on context it is used
func (m *ServerChartDtoFailedItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoFailedItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoFailedItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoFailedItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoGoogleItems0 server chart dto google items0
//
// swagger:model ServerChartDtoGoogleItems0
type ServerChartDtoGoogleItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto google items0
func (m *ServerChartDtoGoogleItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto google items0 based on context it is used
func (m *ServerChartDtoGoogleItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoGoogleItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoGoogleItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoGoogleItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoOpenstackItems0 server chart dto openstack items0
//
// swagger:model ServerChartDtoOpenstackItems0
type ServerChartDtoOpenstackItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto openstack items0
func (m *ServerChartDtoOpenstackItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto openstack items0 based on context it is used
func (m *ServerChartDtoOpenstackItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoOpenstackItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoOpenstackItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoOpenstackItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoSucceededItems0 server chart dto succeeded items0
//
// swagger:model ServerChartDtoSucceededItems0
type ServerChartDtoSucceededItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto succeeded items0
func (m *ServerChartDtoSucceededItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto succeeded items0 based on context it is used
func (m *ServerChartDtoSucceededItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoSucceededItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoSucceededItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoSucceededItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoUpdatingItems0 server chart dto updating items0
//
// swagger:model ServerChartDtoUpdatingItems0
type ServerChartDtoUpdatingItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto updating items0
func (m *ServerChartDtoUpdatingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto updating items0 based on context it is used
func (m *ServerChartDtoUpdatingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoUpdatingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoUpdatingItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoUpdatingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoUsedResourcesItems0 server chart dto used resources items0
//
// swagger:model ServerChartDtoUsedResourcesItems0
type ServerChartDtoUsedResourcesItems0 struct {

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// disk size
	DiskSize int64 `json:"diskSize,omitempty"`

	// max Cpu
	MaxCPU int64 `json:"maxCpu,omitempty"`

	// max disk size
	MaxDiskSize int64 `json:"maxDiskSize,omitempty"`

	// max Ram
	MaxRAM int64 `json:"maxRam,omitempty"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`

	// ram
	RAM int64 `json:"ram,omitempty"`
}

// Validate validates this server chart dto used resources items0
func (m *ServerChartDtoUsedResourcesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto used resources items0 based on context it is used
func (m *ServerChartDtoUsedResourcesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoUsedResourcesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoUsedResourcesItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoUsedResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServerChartDtoWaitingItems0 server chart dto waiting items0
//
// swagger:model ServerChartDtoWaitingItems0
type ServerChartDtoWaitingItems0 struct {

	// names
	Names []string `json:"names"`

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this server chart dto waiting items0
func (m *ServerChartDtoWaitingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this server chart dto waiting items0 based on context it is used
func (m *ServerChartDtoWaitingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServerChartDtoWaitingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerChartDtoWaitingItems0) UnmarshalBinary(b []byte) error {
	var res ServerChartDtoWaitingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
