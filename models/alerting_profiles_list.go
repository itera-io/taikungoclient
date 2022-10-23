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

// AlertingProfilesList alerting profiles list
//
// swagger:model AlertingProfilesList
type AlertingProfilesList struct {

	// data
	Data []*AlertingProfilesListDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this alerting profiles list
func (m *AlertingProfilesList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesList) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this alerting profiles list based on the context it is used
func (m *AlertingProfilesList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesList) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AlertingProfilesList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesList) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertingProfilesListDataItems0 alerting profiles list data items0
//
// swagger:model AlertingProfilesListDataItems0
type AlertingProfilesListDataItems0 struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// emails
	Emails []*AlertingProfilesListDataItems0EmailsItems0 `json:"emails"`

	// id
	ID int32 `json:"id,omitempty"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// projects
	Projects []*AlertingProfilesListDataItems0ProjectsItems0 `json:"projects"`

	// reminder
	Reminder string `json:"reminder,omitempty"`

	// slack configuration Id
	SlackConfigurationID int32 `json:"slackConfigurationId,omitempty"`

	// slack configuration name
	SlackConfigurationName string `json:"slackConfigurationName,omitempty"`

	// webhooks
	Webhooks []*AlertingProfilesListDataItems0WebhooksItems0 `json:"webhooks"`
}

// Validate validates this alerting profiles list data items0
func (m *AlertingProfilesListDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesListDataItems0) validateEmails(formats strfmt.Registry) error {
	if swag.IsZero(m.Emails) { // not required
		return nil
	}

	for i := 0; i < len(m.Emails); i++ {
		if swag.IsZero(m.Emails[i]) { // not required
			continue
		}

		if m.Emails[i] != nil {
			if err := m.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingProfilesListDataItems0) validateProjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Projects) { // not required
		return nil
	}

	for i := 0; i < len(m.Projects); i++ {
		if swag.IsZero(m.Projects[i]) { // not required
			continue
		}

		if m.Projects[i] != nil {
			if err := m.Projects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingProfilesListDataItems0) validateWebhooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Webhooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Webhooks); i++ {
		if swag.IsZero(m.Webhooks[i]) { // not required
			continue
		}

		if m.Webhooks[i] != nil {
			if err := m.Webhooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this alerting profiles list data items0 based on the context it is used
func (m *AlertingProfilesListDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesListDataItems0) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Emails); i++ {

		if m.Emails[i] != nil {
			if err := m.Emails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingProfilesListDataItems0) contextValidateProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Projects); i++ {

		if m.Projects[i] != nil {
			if err := m.Projects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("projects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("projects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AlertingProfilesListDataItems0) contextValidateWebhooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Webhooks); i++ {

		if m.Webhooks[i] != nil {
			if err := m.Webhooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesListDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertingProfilesListDataItems0EmailsItems0 alerting profiles list data items0 emails items0
//
// swagger:model AlertingProfilesListDataItems0EmailsItems0
type AlertingProfilesListDataItems0EmailsItems0 struct {

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this alerting profiles list data items0 emails items0
func (m *AlertingProfilesListDataItems0EmailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles list data items0 emails items0 based on context it is used
func (m *AlertingProfilesListDataItems0EmailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0EmailsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0EmailsItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesListDataItems0EmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertingProfilesListDataItems0ProjectsItems0 alerting profiles list data items0 projects items0
//
// swagger:model AlertingProfilesListDataItems0ProjectsItems0
type AlertingProfilesListDataItems0ProjectsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this alerting profiles list data items0 projects items0
func (m *AlertingProfilesListDataItems0ProjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles list data items0 projects items0 based on context it is used
func (m *AlertingProfilesListDataItems0ProjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0ProjectsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0ProjectsItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesListDataItems0ProjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertingProfilesListDataItems0WebhooksItems0 alerting profiles list data items0 webhooks items0
//
// swagger:model AlertingProfilesListDataItems0WebhooksItems0
type AlertingProfilesListDataItems0WebhooksItems0 struct {

	// headers
	Headers []*AlertingProfilesListDataItems0WebhooksItems0HeadersItems0 `json:"headers"`

	// id
	ID int32 `json:"id,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this alerting profiles list data items0 webhooks items0
func (m *AlertingProfilesListDataItems0WebhooksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesListDataItems0WebhooksItems0) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	for i := 0; i < len(m.Headers); i++ {
		if swag.IsZero(m.Headers[i]) { // not required
			continue
		}

		if m.Headers[i] != nil {
			if err := m.Headers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this alerting profiles list data items0 webhooks items0 based on the context it is used
func (m *AlertingProfilesListDataItems0WebhooksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertingProfilesListDataItems0WebhooksItems0) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Headers); i++ {

		if m.Headers[i] != nil {
			if err := m.Headers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("headers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("headers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0WebhooksItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0WebhooksItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesListDataItems0WebhooksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AlertingProfilesListDataItems0WebhooksItems0HeadersItems0 alerting profiles list data items0 webhooks items0 headers items0
//
// swagger:model AlertingProfilesListDataItems0WebhooksItems0HeadersItems0
type AlertingProfilesListDataItems0WebhooksItems0HeadersItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this alerting profiles list data items0 webhooks items0 headers items0
func (m *AlertingProfilesListDataItems0WebhooksItems0HeadersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles list data items0 webhooks items0 headers items0 based on context it is used
func (m *AlertingProfilesListDataItems0WebhooksItems0HeadersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0WebhooksItems0HeadersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertingProfilesListDataItems0WebhooksItems0HeadersItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesListDataItems0WebhooksItems0HeadersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
