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

// OrganizationsList organizations list
//
// swagger:model OrganizationsList
type OrganizationsList struct {

	// data
	Data []*OrganizationsListDataItems0 `json:"data"`

	// total count
	TotalCount int32 `json:"totalCount,omitempty"`
}

// Validate validates this organizations list
func (m *OrganizationsList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsList) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this organizations list based on the context it is used
func (m *OrganizationsList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsList) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (m *OrganizationsList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsList) UnmarshalBinary(b []byte) error {
	var res OrganizationsList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrganizationsListDataItems0 organizations list data items0
//
// swagger:model OrganizationsListDataItems0
type OrganizationsListDataItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// billing email
	BillingEmail string `json:"billingEmail,omitempty"`

	// bound rules
	BoundRules []*OrganizationsListDataItems0BoundRulesItems0 `json:"boundRules"`

	// city
	City string `json:"city,omitempty"`

	// cloud credentials
	CloudCredentials int32 `json:"cloudCredentials,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// discount rate
	DiscountRate float64 `json:"discountRate"`

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"fullName,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is eligible update subscription
	IsEligibleUpdateSubscription bool `json:"isEligibleUpdateSubscription"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// is read only
	IsReadOnly bool `json:"isReadOnly"`

	// name
	Name string `json:"name,omitempty"`

	// partner
	Partner *OrganizationsListDataItems0Partner `json:"partner,omitempty"`

	// partner Id
	PartnerID int32 `json:"partnerId,omitempty"`

	// partner name
	PartnerName string `json:"partnerName,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// projects
	Projects int32 `json:"projects,omitempty"`

	// servers
	Servers int32 `json:"servers,omitempty"`

	// users
	Users int32 `json:"users,omitempty"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`
}

// Validate validates this organizations list data items0
func (m *OrganizationsListDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsListDataItems0) validateBoundRules(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundRules) { // not required
		return nil
	}

	for i := 0; i < len(m.BoundRules); i++ {
		if swag.IsZero(m.BoundRules[i]) { // not required
			continue
		}

		if m.BoundRules[i] != nil {
			if err := m.BoundRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationsListDataItems0) validatePartner(formats strfmt.Registry) error {
	if swag.IsZero(m.Partner) { // not required
		return nil
	}

	if m.Partner != nil {
		if err := m.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this organizations list data items0 based on the context it is used
func (m *OrganizationsListDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsListDataItems0) contextValidateBoundRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BoundRules); i++ {

		if m.BoundRules[i] != nil {
			if err := m.BoundRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationsListDataItems0) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

	if m.Partner != nil {
		if err := m.Partner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsListDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsListDataItems0) UnmarshalBinary(b []byte) error {
	var res OrganizationsListDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrganizationsListDataItems0BoundRulesItems0 organizations list data items0 bound rules items0
//
// swagger:model OrganizationsListDataItems0BoundRulesItems0
type OrganizationsListDataItems0BoundRulesItems0 struct {

	// prometheus rule Id
	PrometheusRuleID int32 `json:"prometheusRuleId,omitempty"`

	// prometheus rule name
	PrometheusRuleName string `json:"prometheusRuleName,omitempty"`

	// rule discount rate
	RuleDiscountRate float64 `json:"ruleDiscountRate"`
}

// Validate validates this organizations list data items0 bound rules items0
func (m *OrganizationsListDataItems0BoundRulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organizations list data items0 bound rules items0 based on context it is used
func (m *OrganizationsListDataItems0BoundRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsListDataItems0BoundRulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsListDataItems0BoundRulesItems0) UnmarshalBinary(b []byte) error {
	var res OrganizationsListDataItems0BoundRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrganizationsListDataItems0Partner organizations list data items0 partner
//
// swagger:model OrganizationsListDataItems0Partner
type OrganizationsListDataItems0Partner struct {

	// address
	Address string `json:"address,omitempty"`

	// allow registration
	AllowRegistration bool `json:"allowRegistration"`

	// city
	City string `json:"city,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organizations
	Organizations []*OrganizationsListDataItems0PartnerOrganizationsItems0 `json:"organizations"`

	// payment enabled
	PaymentEnabled bool `json:"paymentEnabled"`

	// phone
	Phone string `json:"phone,omitempty"`

	// required user approval
	RequiredUserApproval bool `json:"requiredUserApproval"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`

	// white list domains
	WhiteListDomains []*OrganizationsListDataItems0PartnerWhiteListDomainsItems0 `json:"whiteListDomains"`
}

// Validate validates this organizations list data items0 partner
func (m *OrganizationsListDataItems0Partner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhiteListDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsListDataItems0Partner) validateOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.Organizations) { // not required
		return nil
	}

	for i := 0; i < len(m.Organizations); i++ {
		if swag.IsZero(m.Organizations[i]) { // not required
			continue
		}

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partner" + "." + "organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partner" + "." + "organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationsListDataItems0Partner) validateWhiteListDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.WhiteListDomains) { // not required
		return nil
	}

	for i := 0; i < len(m.WhiteListDomains); i++ {
		if swag.IsZero(m.WhiteListDomains[i]) { // not required
			continue
		}

		if m.WhiteListDomains[i] != nil {
			if err := m.WhiteListDomains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partner" + "." + "whiteListDomains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partner" + "." + "whiteListDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this organizations list data items0 partner based on the context it is used
func (m *OrganizationsListDataItems0Partner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhiteListDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationsListDataItems0Partner) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Organizations); i++ {

		if m.Organizations[i] != nil {
			if err := m.Organizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partner" + "." + "organizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partner" + "." + "organizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OrganizationsListDataItems0Partner) contextValidateWhiteListDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WhiteListDomains); i++ {

		if m.WhiteListDomains[i] != nil {
			if err := m.WhiteListDomains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("partner" + "." + "whiteListDomains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("partner" + "." + "whiteListDomains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsListDataItems0Partner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsListDataItems0Partner) UnmarshalBinary(b []byte) error {
	var res OrganizationsListDataItems0Partner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrganizationsListDataItems0PartnerOrganizationsItems0 organizations list data items0 partner organizations items0
//
// swagger:model OrganizationsListDataItems0PartnerOrganizationsItems0
type OrganizationsListDataItems0PartnerOrganizationsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this organizations list data items0 partner organizations items0
func (m *OrganizationsListDataItems0PartnerOrganizationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organizations list data items0 partner organizations items0 based on context it is used
func (m *OrganizationsListDataItems0PartnerOrganizationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsListDataItems0PartnerOrganizationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsListDataItems0PartnerOrganizationsItems0) UnmarshalBinary(b []byte) error {
	var res OrganizationsListDataItems0PartnerOrganizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OrganizationsListDataItems0PartnerWhiteListDomainsItems0 organizations list data items0 partner white list domains items0
//
// swagger:model OrganizationsListDataItems0PartnerWhiteListDomainsItems0
type OrganizationsListDataItems0PartnerWhiteListDomainsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this organizations list data items0 partner white list domains items0
func (m *OrganizationsListDataItems0PartnerWhiteListDomainsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organizations list data items0 partner white list domains items0 based on context it is used
func (m *OrganizationsListDataItems0PartnerWhiteListDomainsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationsListDataItems0PartnerWhiteListDomainsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationsListDataItems0PartnerWhiteListDomainsItems0) UnmarshalBinary(b []byte) error {
	var res OrganizationsListDataItems0PartnerWhiteListDomainsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
