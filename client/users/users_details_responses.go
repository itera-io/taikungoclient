// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsersDetailsReader is a Reader for the UsersDetails structure.
type UsersDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersDetailsOK creates a UsersDetailsOK with default headers values
func NewUsersDetailsOK() *UsersDetailsOK {
	return &UsersDetailsOK{}
}

/*
UsersDetailsOK describes a response with status code 200, with default header values.

Success
*/
type UsersDetailsOK struct {
	Payload *UsersDetailsOKBody
}

// IsSuccess returns true when this users details o k response has a 2xx status code
func (o *UsersDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users details o k response has a 3xx status code
func (o *UsersDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details o k response has a 4xx status code
func (o *UsersDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users details o k response has a 5xx status code
func (o *UsersDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users details o k response a status code equal to that given
func (o *UsersDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsOK  %+v", 200, o.Payload)
}

func (o *UsersDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsOK  %+v", 200, o.Payload)
}

func (o *UsersDetailsOK) GetPayload() *UsersDetailsOKBody {
	return o.Payload
}

func (o *UsersDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersDetailsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsBadRequest creates a UsersDetailsBadRequest with default headers values
func NewUsersDetailsBadRequest() *UsersDetailsBadRequest {
	return &UsersDetailsBadRequest{}
}

/*
UsersDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersDetailsBadRequest struct {
	Payload []*UsersDetailsBadRequestBodyItems0
}

// IsSuccess returns true when this users details bad request response has a 2xx status code
func (o *UsersDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details bad request response has a 3xx status code
func (o *UsersDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details bad request response has a 4xx status code
func (o *UsersDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details bad request response has a 5xx status code
func (o *UsersDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users details bad request response a status code equal to that given
func (o *UsersDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDetailsBadRequest) GetPayload() []*UsersDetailsBadRequestBodyItems0 {
	return o.Payload
}

func (o *UsersDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsUnauthorized creates a UsersDetailsUnauthorized with default headers values
func NewUsersDetailsUnauthorized() *UsersDetailsUnauthorized {
	return &UsersDetailsUnauthorized{}
}

/*
UsersDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersDetailsUnauthorized struct {
	Payload *UsersDetailsUnauthorizedBody
}

// IsSuccess returns true when this users details unauthorized response has a 2xx status code
func (o *UsersDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details unauthorized response has a 3xx status code
func (o *UsersDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details unauthorized response has a 4xx status code
func (o *UsersDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details unauthorized response has a 5xx status code
func (o *UsersDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users details unauthorized response a status code equal to that given
func (o *UsersDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDetailsUnauthorized) GetPayload() *UsersDetailsUnauthorizedBody {
	return o.Payload
}

func (o *UsersDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersDetailsUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsForbidden creates a UsersDetailsForbidden with default headers values
func NewUsersDetailsForbidden() *UsersDetailsForbidden {
	return &UsersDetailsForbidden{}
}

/*
UsersDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersDetailsForbidden struct {
	Payload *UsersDetailsForbiddenBody
}

// IsSuccess returns true when this users details forbidden response has a 2xx status code
func (o *UsersDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details forbidden response has a 3xx status code
func (o *UsersDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details forbidden response has a 4xx status code
func (o *UsersDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details forbidden response has a 5xx status code
func (o *UsersDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users details forbidden response a status code equal to that given
func (o *UsersDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UsersDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UsersDetailsForbidden) GetPayload() *UsersDetailsForbiddenBody {
	return o.Payload
}

func (o *UsersDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersDetailsForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsNotFound creates a UsersDetailsNotFound with default headers values
func NewUsersDetailsNotFound() *UsersDetailsNotFound {
	return &UsersDetailsNotFound{}
}

/*
UsersDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersDetailsNotFound struct {
	Payload *UsersDetailsNotFoundBody
}

// IsSuccess returns true when this users details not found response has a 2xx status code
func (o *UsersDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details not found response has a 3xx status code
func (o *UsersDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details not found response has a 4xx status code
func (o *UsersDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details not found response has a 5xx status code
func (o *UsersDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users details not found response a status code equal to that given
func (o *UsersDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UsersDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UsersDetailsNotFound) GetPayload() *UsersDetailsNotFoundBody {
	return o.Payload
}

func (o *UsersDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UsersDetailsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsInternalServerError creates a UsersDetailsInternalServerError with default headers values
func NewUsersDetailsInternalServerError() *UsersDetailsInternalServerError {
	return &UsersDetailsInternalServerError{}
}

/*
UsersDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersDetailsInternalServerError struct {
}

// IsSuccess returns true when this users details internal server error response has a 2xx status code
func (o *UsersDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details internal server error response has a 3xx status code
func (o *UsersDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details internal server error response has a 4xx status code
func (o *UsersDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users details internal server error response has a 5xx status code
func (o *UsersDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users details internal server error response a status code equal to that given
func (o *UsersDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsInternalServerError ", 500)
}

func (o *UsersDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsInternalServerError ", 500)
}

func (o *UsersDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
UsersDetailsBadRequestBodyItems0 users details bad request body items0
swagger:model UsersDetailsBadRequestBodyItems0
*/
type UsersDetailsBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this users details bad request body items0
func (o *UsersDetailsBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details bad request body items0 based on context it is used
func (o *UsersDetailsBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res UsersDetailsBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsForbiddenBody users details forbidden body
swagger:model UsersDetailsForbiddenBody
*/
type UsersDetailsForbiddenBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this users details forbidden body
func (o *UsersDetailsForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details forbidden body based on context it is used
func (o *UsersDetailsForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsForbiddenBody) UnmarshalBinary(b []byte) error {
	var res UsersDetailsForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsNotFoundBody users details not found body
swagger:model UsersDetailsNotFoundBody
*/
type UsersDetailsNotFoundBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this users details not found body
func (o *UsersDetailsNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details not found body based on context it is used
func (o *UsersDetailsNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UsersDetailsNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsOKBody users details o k body
swagger:model UsersDetailsOKBody
*/
type UsersDetailsOKBody struct {

	// data
	Data *UsersDetailsOKBodyData `json:"data,omitempty"`

	// demo organization
	DemoOrganization int32 `json:"demoOrganization,omitempty"`

	// is maintenance mode enabled
	IsMaintenanceModeEnabled bool `json:"isMaintenanceModeEnabled"`
}

// Validate validates this users details o k body
func (o *UsersDetailsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersDetailsOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDetailsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDetailsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this users details o k body based on the context it is used
func (o *UsersDetailsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersDetailsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDetailsOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDetailsOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsOKBody) UnmarshalBinary(b []byte) error {
	var res UsersDetailsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsOKBodyData users details o k body data
swagger:model UsersDetailsOKBodyData
*/
type UsersDetailsOKBodyData struct {

	// bound projects
	BoundProjects []*UsersDetailsOKBodyDataBoundProjectsItems0 `json:"boundProjects"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// customer Id
	CustomerID string `json:"customerId,omitempty"`

	// demo mode enabled
	DemoModeEnabled bool `json:"demoModeEnabled"`

	// display name
	DisplayName string `json:"displayName"`

	// email
	Email string `json:"email,omitempty"`

	// has repo
	HasRepo bool `json:"hasRepo"`

	// id
	ID string `json:"id,omitempty"`

	// is approved by partner
	IsApprovedByPartner bool `json:"isApprovedByPartner"`

	// is csm
	IsCsm bool `json:"isCsm"`

	// is eligible update subscription
	IsEligibleUpdateSubscription bool `json:"isEligibleUpdateSubscription"`

	// is email confirmed
	IsEmailConfirmed bool `json:"isEmailConfirmed"`

	// is email notification enabled
	IsEmailNotificationEnabled bool `json:"isEmailNotificationEnabled"`

	// is forced to reset password
	IsForcedToResetPassword bool `json:"isForcedToResetPassword"`

	// is locked
	IsLocked bool `json:"isLocked"`

	// is read only
	IsReadOnly bool `json:"isReadOnly"`

	// last login at
	LastLoginAt string `json:"lastLoginAt,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// organization name
	OrganizationName string `json:"organizationName,omitempty"`

	// owner
	Owner bool `json:"owner"`

	// partner
	Partner *UsersDetailsOKBodyDataPartner `json:"partner,omitempty"`

	// role
	Role string `json:"role,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this users details o k body data
func (o *UsersDetailsOKBodyData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBoundProjects(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePartner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersDetailsOKBodyData) validateBoundProjects(formats strfmt.Registry) error {
	if swag.IsZero(o.BoundProjects) { // not required
		return nil
	}

	for i := 0; i < len(o.BoundProjects); i++ {
		if swag.IsZero(o.BoundProjects[i]) { // not required
			continue
		}

		if o.BoundProjects[i] != nil {
			if err := o.BoundProjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersDetailsOK" + "." + "data" + "." + "boundProjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersDetailsOK" + "." + "data" + "." + "boundProjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UsersDetailsOKBodyData) validatePartner(formats strfmt.Registry) error {
	if swag.IsZero(o.Partner) { // not required
		return nil
	}

	if o.Partner != nil {
		if err := o.Partner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDetailsOK" + "." + "data" + "." + "partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDetailsOK" + "." + "data" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this users details o k body data based on the context it is used
func (o *UsersDetailsOKBodyData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBoundProjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePartner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersDetailsOKBodyData) contextValidateBoundProjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.BoundProjects); i++ {

		if o.BoundProjects[i] != nil {
			if err := o.BoundProjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("usersDetailsOK" + "." + "data" + "." + "boundProjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("usersDetailsOK" + "." + "data" + "." + "boundProjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UsersDetailsOKBodyData) contextValidatePartner(ctx context.Context, formats strfmt.Registry) error {

	if o.Partner != nil {
		if err := o.Partner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usersDetailsOK" + "." + "data" + "." + "partner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usersDetailsOK" + "." + "data" + "." + "partner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsOKBodyData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsOKBodyData) UnmarshalBinary(b []byte) error {
	var res UsersDetailsOKBodyData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsOKBodyDataBoundProjectsItems0 users details o k body data bound projects items0
swagger:model UsersDetailsOKBodyDataBoundProjectsItems0
*/
type UsersDetailsOKBodyDataBoundProjectsItems0 struct {

	// project Id
	ProjectID int32 `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}

// Validate validates this users details o k body data bound projects items0
func (o *UsersDetailsOKBodyDataBoundProjectsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details o k body data bound projects items0 based on context it is used
func (o *UsersDetailsOKBodyDataBoundProjectsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsOKBodyDataBoundProjectsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsOKBodyDataBoundProjectsItems0) UnmarshalBinary(b []byte) error {
	var res UsersDetailsOKBodyDataBoundProjectsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsOKBodyDataPartner users details o k body data partner
swagger:model UsersDetailsOKBodyDataPartner
*/
type UsersDetailsOKBodyDataPartner struct {

	// id
	ID int32 `json:"id,omitempty"`

	// link
	Link string `json:"link,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this users details o k body data partner
func (o *UsersDetailsOKBodyDataPartner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details o k body data partner based on context it is used
func (o *UsersDetailsOKBodyDataPartner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsOKBodyDataPartner) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsOKBodyDataPartner) UnmarshalBinary(b []byte) error {
	var res UsersDetailsOKBodyDataPartner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UsersDetailsUnauthorizedBody users details unauthorized body
swagger:model UsersDetailsUnauthorizedBody
*/
type UsersDetailsUnauthorizedBody struct {

	// detail
	Detail string `json:"detail,omitempty"`

	// instance
	Instance string `json:"instance,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this users details unauthorized body
func (o *UsersDetailsUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this users details unauthorized body based on context it is used
func (o *UsersDetailsUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersDetailsUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersDetailsUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res UsersDetailsUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
