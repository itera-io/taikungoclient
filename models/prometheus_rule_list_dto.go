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

// PrometheusRuleListDto prometheus rule list dto
//
// swagger:model PrometheusRuleListDto
type PrometheusRuleListDto struct {

	// billing start date
	BillingStartDate string `json:"billingStartDate,omitempty"`

	// bound organizations
	BoundOrganizations []*PrometheusRuleListDtoBoundOrganizationsItems0 `json:"boundOrganizations"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// is all
	IsAll bool `json:"isAll"`

	// labels
	Labels []*PrometheusRuleListDtoLabelsItems0 `json:"labels"`

	// last modified
	LastModified string `json:"lastModified,omitempty"`

	// last modified by
	LastModifiedBy string `json:"lastModifiedBy,omitempty"`

	// metric name
	MetricName string `json:"metricName,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// operation credential
	OperationCredential *PrometheusRuleListDtoOperationCredential `json:"operationCredential,omitempty"`

	// partner
	Partner *PrometheusRuleListDtoPartner `json:"partner,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this prometheus rule list dto
func (m *PrometheusRuleListDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundOrganizations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationCredential(formats); err != nil {
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

func (m *PrometheusRuleListDto) validateBoundOrganizations(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundOrganizations) { // not required
		return nil
	}

	for i := 0; i < len(m.BoundOrganizations); i++ {
		if swag.IsZero(m.BoundOrganizations[i]) { // not required
			continue
		}

		if m.BoundOrganizations[i] != nil {
			if err := m.BoundOrganizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundOrganizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundOrganizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrometheusRuleListDto) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrometheusRuleListDto) validateOperationCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationCredential) { // not required
		return nil
	}

	if m.OperationCredential != nil {
		if err := m.OperationCredential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationCredential")
			}
			return err
		}
	}

	return nil
}

func (m *PrometheusRuleListDto) validatePartner(formats strfmt.Registry) error {
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

// ContextValidate validate this prometheus rule list dto based on the context it is used
func (m *PrometheusRuleListDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundOrganizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationCredential(ctx, formats); err != nil {
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

func (m *PrometheusRuleListDto) contextValidateBoundOrganizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BoundOrganizations); i++ {

		if m.BoundOrganizations[i] != nil {
			if err := m.BoundOrganizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("boundOrganizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("boundOrganizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrometheusRuleListDto) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PrometheusRuleListDto) contextValidateOperationCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.OperationCredential != nil {
		if err := m.OperationCredential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operationCredential")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operationCredential")
			}
			return err
		}
	}

	return nil
}

func (m *PrometheusRuleListDto) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PrometheusRuleListDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDto) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoBoundOrganizationsItems0 prometheus rule list dto bound organizations items0
//
// swagger:model PrometheusRuleListDtoBoundOrganizationsItems0
type PrometheusRuleListDtoBoundOrganizationsItems0 struct {

	// global discount rate
	GlobalDiscountRate float64 `json:"globalDiscountRate"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// rule discount rate
	RuleDiscountRate float64 `json:"ruleDiscountRate"`
}

// Validate validates this prometheus rule list dto bound organizations items0
func (m *PrometheusRuleListDtoBoundOrganizationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule list dto bound organizations items0 based on context it is used
func (m *PrometheusRuleListDtoBoundOrganizationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRuleListDtoBoundOrganizationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoBoundOrganizationsItems0) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoBoundOrganizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoLabelsItems0 prometheus rule list dto labels items0
//
// swagger:model PrometheusRuleListDtoLabelsItems0
type PrometheusRuleListDtoLabelsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this prometheus rule list dto labels items0
func (m *PrometheusRuleListDtoLabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule list dto labels items0 based on context it is used
func (m *PrometheusRuleListDtoLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRuleListDtoLabelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoLabelsItems0) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoOperationCredential prometheus rule list dto operation credential
//
// swagger:model PrometheusRuleListDtoOperationCredential
type PrometheusRuleListDtoOperationCredential struct {

	// is default
	IsDefault bool `json:"isDefault"`

	// name
	Name string `json:"name,omitempty"`

	// operation credential Id
	OperationCredentialID int32 `json:"operationCredentialId,omitempty"`
}

// Validate validates this prometheus rule list dto operation credential
func (m *PrometheusRuleListDtoOperationCredential) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule list dto operation credential based on context it is used
func (m *PrometheusRuleListDtoOperationCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRuleListDtoOperationCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoOperationCredential) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoOperationCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoPartner prometheus rule list dto partner
//
// swagger:model PrometheusRuleListDtoPartner
type PrometheusRuleListDtoPartner struct {

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
	Organizations []*PrometheusRuleListDtoPartnerOrganizationsItems0 `json:"organizations"`

	// payment enabled
	PaymentEnabled bool `json:"paymentEnabled"`

	// phone
	Phone string `json:"phone,omitempty"`

	// required user approval
	RequiredUserApproval bool `json:"requiredUserApproval"`

	// vat number
	VatNumber string `json:"vatNumber,omitempty"`

	// white list domains
	WhiteListDomains []*PrometheusRuleListDtoPartnerWhiteListDomainsItems0 `json:"whiteListDomains"`
}

// Validate validates this prometheus rule list dto partner
func (m *PrometheusRuleListDtoPartner) Validate(formats strfmt.Registry) error {
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

func (m *PrometheusRuleListDtoPartner) validateOrganizations(formats strfmt.Registry) error {
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

func (m *PrometheusRuleListDtoPartner) validateWhiteListDomains(formats strfmt.Registry) error {
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

// ContextValidate validate this prometheus rule list dto partner based on the context it is used
func (m *PrometheusRuleListDtoPartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PrometheusRuleListDtoPartner) contextValidateOrganizations(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PrometheusRuleListDtoPartner) contextValidateWhiteListDomains(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PrometheusRuleListDtoPartner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoPartner) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoPartnerOrganizationsItems0 prometheus rule list dto partner organizations items0
//
// swagger:model PrometheusRuleListDtoPartnerOrganizationsItems0
type PrometheusRuleListDtoPartnerOrganizationsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this prometheus rule list dto partner organizations items0
func (m *PrometheusRuleListDtoPartnerOrganizationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule list dto partner organizations items0 based on context it is used
func (m *PrometheusRuleListDtoPartnerOrganizationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRuleListDtoPartnerOrganizationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoPartnerOrganizationsItems0) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoPartnerOrganizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PrometheusRuleListDtoPartnerWhiteListDomainsItems0 prometheus rule list dto partner white list domains items0
//
// swagger:model PrometheusRuleListDtoPartnerWhiteListDomainsItems0
type PrometheusRuleListDtoPartnerWhiteListDomainsItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this prometheus rule list dto partner white list domains items0
func (m *PrometheusRuleListDtoPartnerWhiteListDomainsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this prometheus rule list dto partner white list domains items0 based on context it is used
func (m *PrometheusRuleListDtoPartnerWhiteListDomainsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusRuleListDtoPartnerWhiteListDomainsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusRuleListDtoPartnerWhiteListDomainsItems0) UnmarshalBinary(b []byte) error {
	var res PrometheusRuleListDtoPartnerWhiteListDomainsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
