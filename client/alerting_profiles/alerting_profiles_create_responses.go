// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertingProfilesCreateReader is a Reader for the AlertingProfilesCreate structure.
type AlertingProfilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesCreateOK creates a AlertingProfilesCreateOK with default headers values
func NewAlertingProfilesCreateOK() *AlertingProfilesCreateOK {
	return &AlertingProfilesCreateOK{}
}

/*
AlertingProfilesCreateOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesCreateOK struct {
	Payload *AlertingProfilesCreateOKBody
}

// IsSuccess returns true when this alerting profiles create o k response has a 2xx status code
func (o *AlertingProfilesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles create o k response has a 3xx status code
func (o *AlertingProfilesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create o k response has a 4xx status code
func (o *AlertingProfilesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles create o k response has a 5xx status code
func (o *AlertingProfilesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create o k response a status code equal to that given
func (o *AlertingProfilesCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesCreateOK) GetPayload() *AlertingProfilesCreateOKBody {
	return o.Payload
}

func (o *AlertingProfilesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AlertingProfilesCreateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateBadRequest creates a AlertingProfilesCreateBadRequest with default headers values
func NewAlertingProfilesCreateBadRequest() *AlertingProfilesCreateBadRequest {
	return &AlertingProfilesCreateBadRequest{}
}

/*
AlertingProfilesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesCreateBadRequest struct {
	Payload []*AlertingProfilesCreateBadRequestBodyItems0
}

// IsSuccess returns true when this alerting profiles create bad request response has a 2xx status code
func (o *AlertingProfilesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create bad request response has a 3xx status code
func (o *AlertingProfilesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create bad request response has a 4xx status code
func (o *AlertingProfilesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create bad request response has a 5xx status code
func (o *AlertingProfilesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create bad request response a status code equal to that given
func (o *AlertingProfilesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesCreateBadRequest) GetPayload() []*AlertingProfilesCreateBadRequestBodyItems0 {
	return o.Payload
}

func (o *AlertingProfilesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateUnauthorized creates a AlertingProfilesCreateUnauthorized with default headers values
func NewAlertingProfilesCreateUnauthorized() *AlertingProfilesCreateUnauthorized {
	return &AlertingProfilesCreateUnauthorized{}
}

/*
AlertingProfilesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesCreateUnauthorized struct {
	Payload *AlertingProfilesCreateUnauthorizedBody
}

// IsSuccess returns true when this alerting profiles create unauthorized response has a 2xx status code
func (o *AlertingProfilesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create unauthorized response has a 3xx status code
func (o *AlertingProfilesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create unauthorized response has a 4xx status code
func (o *AlertingProfilesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create unauthorized response has a 5xx status code
func (o *AlertingProfilesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create unauthorized response a status code equal to that given
func (o *AlertingProfilesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesCreateUnauthorized) GetPayload() *AlertingProfilesCreateUnauthorizedBody {
	return o.Payload
}

func (o *AlertingProfilesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AlertingProfilesCreateUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateForbidden creates a AlertingProfilesCreateForbidden with default headers values
func NewAlertingProfilesCreateForbidden() *AlertingProfilesCreateForbidden {
	return &AlertingProfilesCreateForbidden{}
}

/*
AlertingProfilesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesCreateForbidden struct {
	Payload *AlertingProfilesCreateForbiddenBody
}

// IsSuccess returns true when this alerting profiles create forbidden response has a 2xx status code
func (o *AlertingProfilesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create forbidden response has a 3xx status code
func (o *AlertingProfilesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create forbidden response has a 4xx status code
func (o *AlertingProfilesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create forbidden response has a 5xx status code
func (o *AlertingProfilesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create forbidden response a status code equal to that given
func (o *AlertingProfilesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesCreateForbidden) GetPayload() *AlertingProfilesCreateForbiddenBody {
	return o.Payload
}

func (o *AlertingProfilesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AlertingProfilesCreateForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateNotFound creates a AlertingProfilesCreateNotFound with default headers values
func NewAlertingProfilesCreateNotFound() *AlertingProfilesCreateNotFound {
	return &AlertingProfilesCreateNotFound{}
}

/*
AlertingProfilesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesCreateNotFound struct {
	Payload *AlertingProfilesCreateNotFoundBody
}

// IsSuccess returns true when this alerting profiles create not found response has a 2xx status code
func (o *AlertingProfilesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create not found response has a 3xx status code
func (o *AlertingProfilesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create not found response has a 4xx status code
func (o *AlertingProfilesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create not found response has a 5xx status code
func (o *AlertingProfilesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create not found response a status code equal to that given
func (o *AlertingProfilesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesCreateNotFound) GetPayload() *AlertingProfilesCreateNotFoundBody {
	return o.Payload
}

func (o *AlertingProfilesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AlertingProfilesCreateNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateInternalServerError creates a AlertingProfilesCreateInternalServerError with default headers values
func NewAlertingProfilesCreateInternalServerError() *AlertingProfilesCreateInternalServerError {
	return &AlertingProfilesCreateInternalServerError{}
}

/*
AlertingProfilesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesCreateInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles create internal server error response has a 2xx status code
func (o *AlertingProfilesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create internal server error response has a 3xx status code
func (o *AlertingProfilesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create internal server error response has a 4xx status code
func (o *AlertingProfilesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles create internal server error response has a 5xx status code
func (o *AlertingProfilesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles create internal server error response a status code equal to that given
func (o *AlertingProfilesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateInternalServerError ", 500)
}

func (o *AlertingProfilesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateInternalServerError ", 500)
}

func (o *AlertingProfilesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
AlertingProfilesCreateBadRequestBodyItems0 alerting profiles create bad request body items0
swagger:model AlertingProfilesCreateBadRequestBodyItems0
*/
type AlertingProfilesCreateBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this alerting profiles create bad request body items0
func (o *AlertingProfilesCreateBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create bad request body items0 based on context it is used
func (o *AlertingProfilesCreateBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateBody alerting profiles create body
swagger:model AlertingProfilesCreateBody
*/
type AlertingProfilesCreateBody struct {

	// alerting integrations
	AlertingIntegrations []*AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0 `json:"alertingIntegrations"`

	// emails
	Emails []*AlertingProfilesCreateParamsBodyEmailsItems0 `json:"emails"`

	// name
	Name string `json:"name,omitempty"`

	// organization Id
	OrganizationID int32 `json:"organizationId,omitempty"`

	// reminder
	// Enum: [100 200 300 -1]
	Reminder int32 `json:"reminder,omitempty"`

	// slack configuration Id
	SlackConfigurationID int32 `json:"slackConfigurationId,omitempty"`

	// webhooks
	Webhooks []*AlertingProfilesCreateParamsBodyWebhooksItems0 `json:"webhooks"`
}

// Validate validates this alerting profiles create body
func (o *AlertingProfilesCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlertingIntegrations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmails(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReminder(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWebhooks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AlertingProfilesCreateBody) validateAlertingIntegrations(formats strfmt.Registry) error {
	if swag.IsZero(o.AlertingIntegrations) { // not required
		return nil
	}

	for i := 0; i < len(o.AlertingIntegrations); i++ {
		if swag.IsZero(o.AlertingIntegrations[i]) { // not required
			continue
		}

		if o.AlertingIntegrations[i] != nil {
			if err := o.AlertingIntegrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "alertingIntegrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "alertingIntegrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AlertingProfilesCreateBody) validateEmails(formats strfmt.Registry) error {
	if swag.IsZero(o.Emails) { // not required
		return nil
	}

	for i := 0; i < len(o.Emails); i++ {
		if swag.IsZero(o.Emails[i]) { // not required
			continue
		}

		if o.Emails[i] != nil {
			if err := o.Emails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var alertingProfilesCreateBodyTypeReminderPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[100,200,300,-1]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertingProfilesCreateBodyTypeReminderPropEnum = append(alertingProfilesCreateBodyTypeReminderPropEnum, v)
	}
}

// prop value enum
func (o *AlertingProfilesCreateBody) validateReminderEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, alertingProfilesCreateBodyTypeReminderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AlertingProfilesCreateBody) validateReminder(formats strfmt.Registry) error {
	if swag.IsZero(o.Reminder) { // not required
		return nil
	}

	// value enum
	if err := o.validateReminderEnum("body"+"."+"reminder", "body", o.Reminder); err != nil {
		return err
	}

	return nil
}

func (o *AlertingProfilesCreateBody) validateWebhooks(formats strfmt.Registry) error {
	if swag.IsZero(o.Webhooks) { // not required
		return nil
	}

	for i := 0; i < len(o.Webhooks); i++ {
		if swag.IsZero(o.Webhooks[i]) { // not required
			continue
		}

		if o.Webhooks[i] != nil {
			if err := o.Webhooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this alerting profiles create body based on the context it is used
func (o *AlertingProfilesCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlertingIntegrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEmails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWebhooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AlertingProfilesCreateBody) contextValidateAlertingIntegrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AlertingIntegrations); i++ {

		if o.AlertingIntegrations[i] != nil {
			if err := o.AlertingIntegrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "alertingIntegrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "alertingIntegrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AlertingProfilesCreateBody) contextValidateEmails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Emails); i++ {

		if o.Emails[i] != nil {
			if err := o.Emails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "emails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "emails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *AlertingProfilesCreateBody) contextValidateWebhooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Webhooks); i++ {

		if o.Webhooks[i] != nil {
			if err := o.Webhooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "webhooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "webhooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateBody) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateForbiddenBody alerting profiles create forbidden body
swagger:model AlertingProfilesCreateForbiddenBody
*/
type AlertingProfilesCreateForbiddenBody struct {

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

// Validate validates this alerting profiles create forbidden body
func (o *AlertingProfilesCreateForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create forbidden body based on context it is used
func (o *AlertingProfilesCreateForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateNotFoundBody alerting profiles create not found body
swagger:model AlertingProfilesCreateNotFoundBody
*/
type AlertingProfilesCreateNotFoundBody struct {

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

// Validate validates this alerting profiles create not found body
func (o *AlertingProfilesCreateNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create not found body based on context it is used
func (o *AlertingProfilesCreateNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateOKBody alerting profiles create o k body
swagger:model AlertingProfilesCreateOKBody
*/
type AlertingProfilesCreateOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// is error
	IsError bool `json:"isError"`

	// message
	Message string `json:"message,omitempty"`

	// result
	Result interface{} `json:"result,omitempty"`

	// status
	Status int32 `json:"status,omitempty"`
}

// Validate validates this alerting profiles create o k body
func (o *AlertingProfilesCreateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create o k body based on context it is used
func (o *AlertingProfilesCreateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateOKBody) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0 alerting profiles create params body alerting integrations items0
swagger:model AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0
*/
type AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0 struct {

	// alerting integration type
	// Enum: [100 200 300 400]
	AlertingIntegrationType int32 `json:"alertingIntegrationType,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this alerting profiles create params body alerting integrations items0
func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlertingIntegrationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var alertingProfilesCreateParamsBodyAlertingIntegrationsItems0TypeAlertingIntegrationTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[100,200,300,400]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertingProfilesCreateParamsBodyAlertingIntegrationsItems0TypeAlertingIntegrationTypePropEnum = append(alertingProfilesCreateParamsBodyAlertingIntegrationsItems0TypeAlertingIntegrationTypePropEnum, v)
	}
}

// prop value enum
func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) validateAlertingIntegrationTypeEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, alertingProfilesCreateParamsBodyAlertingIntegrationsItems0TypeAlertingIntegrationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) validateAlertingIntegrationType(formats strfmt.Registry) error {
	if swag.IsZero(o.AlertingIntegrationType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAlertingIntegrationTypeEnum("alertingIntegrationType", "body", o.AlertingIntegrationType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alerting profiles create params body alerting integrations items0 based on context it is used
func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateParamsBodyAlertingIntegrationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateParamsBodyEmailsItems0 alerting profiles create params body emails items0
swagger:model AlertingProfilesCreateParamsBodyEmailsItems0
*/
type AlertingProfilesCreateParamsBodyEmailsItems0 struct {

	// email
	Email string `json:"email,omitempty"`
}

// Validate validates this alerting profiles create params body emails items0
func (o *AlertingProfilesCreateParamsBodyEmailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create params body emails items0 based on context it is used
func (o *AlertingProfilesCreateParamsBodyEmailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyEmailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyEmailsItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateParamsBodyEmailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateParamsBodyWebhooksItems0 alerting profiles create params body webhooks items0
swagger:model AlertingProfilesCreateParamsBodyWebhooksItems0
*/
type AlertingProfilesCreateParamsBodyWebhooksItems0 struct {

	// headers
	Headers []*AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0 `json:"headers"`

	// id
	ID int32 `json:"id,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this alerting profiles create params body webhooks items0
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(o.Headers) { // not required
		return nil
	}

	for i := 0; i < len(o.Headers); i++ {
		if swag.IsZero(o.Headers[i]) { // not required
			continue
		}

		if o.Headers[i] != nil {
			if err := o.Headers[i].Validate(formats); err != nil {
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

// ContextValidate validate this alerting profiles create params body webhooks items0 based on the context it is used
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Headers); i++ {

		if o.Headers[i] != nil {
			if err := o.Headers[i].ContextValidate(ctx, formats); err != nil {
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
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateParamsBodyWebhooksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0 alerting profiles create params body webhooks items0 headers items0
swagger:model AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0
*/
type AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0 struct {

	// id
	ID int32 `json:"id,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this alerting profiles create params body webhooks items0 headers items0
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create params body webhooks items0 headers items0 based on context it is used
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateParamsBodyWebhooksItems0HeadersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AlertingProfilesCreateUnauthorizedBody alerting profiles create unauthorized body
swagger:model AlertingProfilesCreateUnauthorizedBody
*/
type AlertingProfilesCreateUnauthorizedBody struct {

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

// Validate validates this alerting profiles create unauthorized body
func (o *AlertingProfilesCreateUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alerting profiles create unauthorized body based on context it is used
func (o *AlertingProfilesCreateUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AlertingProfilesCreateUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AlertingProfilesCreateUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res AlertingProfilesCreateUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
