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

// EnumList enum list
//
// swagger:model EnumList
type EnumList struct {

	// alert types
	AlertTypes []*CommonDropdownDto `json:"alertTypes"`

	// alerting inntegration types
	AlertingInntegrationTypes []*CommonDropdownDto `json:"alertingInntegrationTypes"`

	// audit logs
	AuditLogs []*CommonDropdownDto `json:"auditLogs"`

	// availability
	Availability []*CommonAvailabilityDto `json:"availability"`

	// aws platforms
	AwsPlatforms []*CommonStringBasedDropdownDto `json:"awsPlatforms"`

	// azure quotas
	AzureQuotas []*CommonDropdownDto `json:"azureQuotas"`

	// cloud types
	CloudTypes interface{} `json:"cloudTypes,omitempty"`

	// cron periods
	CronPeriods []*CommonStringBasedDropdownDto `json:"cronPeriods"`

	// google image types
	GoogleImageTypes []*CommonDropdownDto `json:"googleImageTypes"`

	// project statuses
	ProjectStatuses []*CommonDropdownDto `json:"projectStatuses"`

	// prometheus types
	PrometheusTypes []*CommonDropdownDto `json:"prometheusTypes"`

	// reboot options
	RebootOptions []*CommonDropdownDto `json:"rebootOptions"`

	// reminder types
	ReminderTypes []*CommonDropdownDto `json:"reminderTypes"`

	// request logs
	RequestLogs []*CommonDropdownDto `json:"requestLogs"`

	// security group rules
	SecurityGroupRules []*CommonDropdownDto `json:"securityGroupRules"`

	// server roles
	ServerRoles []*CommonDropdownDto `json:"serverRoles"`

	// server statuses
	ServerStatuses []*CommonDropdownDto `json:"serverStatuses"`

	// showback kinds
	ShowbackKinds []*CommonDropdownDto `json:"showbackKinds"`

	// slack types
	SlackTypes []*CommonDropdownDto `json:"slackTypes"`

	// standalone Vm statuses
	StandaloneVMStatuses []*CommonDropdownDto `json:"standaloneVmStatuses"`

	// user roles
	UserRoles []*CommonDropdownDto `json:"userRoles"`
}

// Validate validates this enum list
func (m *EnumList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlertingInntegrationTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAwsPlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureQuotas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCronPeriods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogleImageTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrometheusTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebootOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReminderTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroupRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShowbackKinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlackTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandaloneVMStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnumList) validateAlertTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.AlertTypes); i++ {
		if swag.IsZero(m.AlertTypes[i]) { // not required
			continue
		}

		if m.AlertTypes[i] != nil {
			if err := m.AlertTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alertTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alertTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateAlertingInntegrationTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.AlertingInntegrationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.AlertingInntegrationTypes); i++ {
		if swag.IsZero(m.AlertingInntegrationTypes[i]) { // not required
			continue
		}

		if m.AlertingInntegrationTypes[i] != nil {
			if err := m.AlertingInntegrationTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alertingInntegrationTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alertingInntegrationTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateAvailability(formats strfmt.Registry) error {
	if swag.IsZero(m.Availability) { // not required
		return nil
	}

	for i := 0; i < len(m.Availability); i++ {
		if swag.IsZero(m.Availability[i]) { // not required
			continue
		}

		if m.Availability[i] != nil {
			if err := m.Availability[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availability" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateAwsPlatforms(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsPlatforms) { // not required
		return nil
	}

	for i := 0; i < len(m.AwsPlatforms); i++ {
		if swag.IsZero(m.AwsPlatforms[i]) { // not required
			continue
		}

		if m.AwsPlatforms[i] != nil {
			if err := m.AwsPlatforms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("awsPlatforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("awsPlatforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateAzureQuotas(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureQuotas) { // not required
		return nil
	}

	for i := 0; i < len(m.AzureQuotas); i++ {
		if swag.IsZero(m.AzureQuotas[i]) { // not required
			continue
		}

		if m.AzureQuotas[i] != nil {
			if err := m.AzureQuotas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureQuotas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureQuotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateCronPeriods(formats strfmt.Registry) error {
	if swag.IsZero(m.CronPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.CronPeriods); i++ {
		if swag.IsZero(m.CronPeriods[i]) { // not required
			continue
		}

		if m.CronPeriods[i] != nil {
			if err := m.CronPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cronPeriods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cronPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateGoogleImageTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.GoogleImageTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.GoogleImageTypes); i++ {
		if swag.IsZero(m.GoogleImageTypes[i]) { // not required
			continue
		}

		if m.GoogleImageTypes[i] != nil {
			if err := m.GoogleImageTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("googleImageTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("googleImageTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateProjectStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.ProjectStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.ProjectStatuses); i++ {
		if swag.IsZero(m.ProjectStatuses[i]) { // not required
			continue
		}

		if m.ProjectStatuses[i] != nil {
			if err := m.ProjectStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validatePrometheusTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.PrometheusTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.PrometheusTypes); i++ {
		if swag.IsZero(m.PrometheusTypes[i]) { // not required
			continue
		}

		if m.PrometheusTypes[i] != nil {
			if err := m.PrometheusTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prometheusTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prometheusTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateRebootOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.RebootOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.RebootOptions); i++ {
		if swag.IsZero(m.RebootOptions[i]) { // not required
			continue
		}

		if m.RebootOptions[i] != nil {
			if err := m.RebootOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rebootOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rebootOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateReminderTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ReminderTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ReminderTypes); i++ {
		if swag.IsZero(m.ReminderTypes[i]) { // not required
			continue
		}

		if m.ReminderTypes[i] != nil {
			if err := m.ReminderTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reminderTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reminderTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateRequestLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestLogs); i++ {
		if swag.IsZero(m.RequestLogs[i]) { // not required
			continue
		}

		if m.RequestLogs[i] != nil {
			if err := m.RequestLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateSecurityGroupRules(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityGroupRules) { // not required
		return nil
	}

	for i := 0; i < len(m.SecurityGroupRules); i++ {
		if swag.IsZero(m.SecurityGroupRules[i]) { // not required
			continue
		}

		if m.SecurityGroupRules[i] != nil {
			if err := m.SecurityGroupRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateServerRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerRoles); i++ {
		if swag.IsZero(m.ServerRoles[i]) { // not required
			continue
		}

		if m.ServerRoles[i] != nil {
			if err := m.ServerRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateServerStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerStatuses); i++ {
		if swag.IsZero(m.ServerStatuses[i]) { // not required
			continue
		}

		if m.ServerStatuses[i] != nil {
			if err := m.ServerStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateShowbackKinds(formats strfmt.Registry) error {
	if swag.IsZero(m.ShowbackKinds) { // not required
		return nil
	}

	for i := 0; i < len(m.ShowbackKinds); i++ {
		if swag.IsZero(m.ShowbackKinds[i]) { // not required
			continue
		}

		if m.ShowbackKinds[i] != nil {
			if err := m.ShowbackKinds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("showbackKinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("showbackKinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateSlackTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.SlackTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.SlackTypes); i++ {
		if swag.IsZero(m.SlackTypes[i]) { // not required
			continue
		}

		if m.SlackTypes[i] != nil {
			if err := m.SlackTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slackTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slackTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateStandaloneVMStatuses(formats strfmt.Registry) error {
	if swag.IsZero(m.StandaloneVMStatuses) { // not required
		return nil
	}

	for i := 0; i < len(m.StandaloneVMStatuses); i++ {
		if swag.IsZero(m.StandaloneVMStatuses[i]) { // not required
			continue
		}

		if m.StandaloneVMStatuses[i] != nil {
			if err := m.StandaloneVMStatuses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVmStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVmStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) validateUserRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.UserRoles) { // not required
		return nil
	}

	for i := 0; i < len(m.UserRoles); i++ {
		if swag.IsZero(m.UserRoles[i]) { // not required
			continue
		}

		if m.UserRoles[i] != nil {
			if err := m.UserRoles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this enum list based on the context it is used
func (m *EnumList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAlertTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAlertingInntegrationTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAvailability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAwsPlatforms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureQuotas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCronPeriods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogleImageTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjectStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrometheusTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRebootOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReminderTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRequestLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityGroupRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShowbackKinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSlackTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandaloneVMStatuses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnumList) contextValidateAlertTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AlertTypes); i++ {

		if m.AlertTypes[i] != nil {
			if err := m.AlertTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alertTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alertTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateAlertingInntegrationTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AlertingInntegrationTypes); i++ {

		if m.AlertingInntegrationTypes[i] != nil {
			if err := m.AlertingInntegrationTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("alertingInntegrationTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("alertingInntegrationTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateAvailability(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Availability); i++ {

		if m.Availability[i] != nil {
			if err := m.Availability[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availability" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateAwsPlatforms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AwsPlatforms); i++ {

		if m.AwsPlatforms[i] != nil {
			if err := m.AwsPlatforms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("awsPlatforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("awsPlatforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateAzureQuotas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AzureQuotas); i++ {

		if m.AzureQuotas[i] != nil {
			if err := m.AzureQuotas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureQuotas" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureQuotas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateCronPeriods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CronPeriods); i++ {

		if m.CronPeriods[i] != nil {
			if err := m.CronPeriods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cronPeriods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cronPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateGoogleImageTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.GoogleImageTypes); i++ {

		if m.GoogleImageTypes[i] != nil {
			if err := m.GoogleImageTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("googleImageTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("googleImageTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateProjectStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProjectStatuses); i++ {

		if m.ProjectStatuses[i] != nil {
			if err := m.ProjectStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projectStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projectStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidatePrometheusTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrometheusTypes); i++ {

		if m.PrometheusTypes[i] != nil {
			if err := m.PrometheusTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prometheusTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prometheusTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateRebootOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RebootOptions); i++ {

		if m.RebootOptions[i] != nil {
			if err := m.RebootOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rebootOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rebootOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateReminderTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ReminderTypes); i++ {

		if m.ReminderTypes[i] != nil {
			if err := m.ReminderTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reminderTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reminderTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateRequestLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RequestLogs); i++ {

		if m.RequestLogs[i] != nil {
			if err := m.RequestLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("requestLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateSecurityGroupRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecurityGroupRules); i++ {

		if m.SecurityGroupRules[i] != nil {
			if err := m.SecurityGroupRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("securityGroupRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateServerRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerRoles); i++ {

		if m.ServerRoles[i] != nil {
			if err := m.ServerRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateServerStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerStatuses); i++ {

		if m.ServerStatuses[i] != nil {
			if err := m.ServerStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateShowbackKinds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ShowbackKinds); i++ {

		if m.ShowbackKinds[i] != nil {
			if err := m.ShowbackKinds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("showbackKinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("showbackKinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateSlackTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SlackTypes); i++ {

		if m.SlackTypes[i] != nil {
			if err := m.SlackTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slackTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("slackTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateStandaloneVMStatuses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StandaloneVMStatuses); i++ {

		if m.StandaloneVMStatuses[i] != nil {
			if err := m.StandaloneVMStatuses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("standaloneVmStatuses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("standaloneVmStatuses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EnumList) contextValidateUserRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserRoles); i++ {

		if m.UserRoles[i] != nil {
			if err := m.UserRoles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userRoles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userRoles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnumList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnumList) UnmarshalBinary(b []byte) error {
	var res EnumList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
