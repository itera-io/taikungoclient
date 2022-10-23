// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminUpdateUserEmailReader is a Reader for the AdminUpdateUserEmail structure.
type AdminUpdateUserEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUpdateUserEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUpdateUserEmailOK creates a AdminUpdateUserEmailOK with default headers values
func NewAdminUpdateUserEmailOK() *AdminUpdateUserEmailOK {
	return &AdminUpdateUserEmailOK{}
}

/*
AdminUpdateUserEmailOK describes a response with status code 200, with default header values.

Success
*/
type AdminUpdateUserEmailOK struct {
	Payload interface{}
}

// IsSuccess returns true when this admin update user email o k response has a 2xx status code
func (o *AdminUpdateUserEmailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin update user email o k response has a 3xx status code
func (o *AdminUpdateUserEmailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email o k response has a 4xx status code
func (o *AdminUpdateUserEmailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user email o k response has a 5xx status code
func (o *AdminUpdateUserEmailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user email o k response a status code equal to that given
func (o *AdminUpdateUserEmailOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminUpdateUserEmailOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateUserEmailOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailOK  %+v", 200, o.Payload)
}

func (o *AdminUpdateUserEmailOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AdminUpdateUserEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailBadRequest creates a AdminUpdateUserEmailBadRequest with default headers values
func NewAdminUpdateUserEmailBadRequest() *AdminUpdateUserEmailBadRequest {
	return &AdminUpdateUserEmailBadRequest{}
}

/*
AdminUpdateUserEmailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUpdateUserEmailBadRequest struct {
	Payload []*AdminUpdateUserEmailBadRequestBodyItems0
}

// IsSuccess returns true when this admin update user email bad request response has a 2xx status code
func (o *AdminUpdateUserEmailBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user email bad request response has a 3xx status code
func (o *AdminUpdateUserEmailBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email bad request response has a 4xx status code
func (o *AdminUpdateUserEmailBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user email bad request response has a 5xx status code
func (o *AdminUpdateUserEmailBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user email bad request response a status code equal to that given
func (o *AdminUpdateUserEmailBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminUpdateUserEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateUserEmailBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUpdateUserEmailBadRequest) GetPayload() []*AdminUpdateUserEmailBadRequestBodyItems0 {
	return o.Payload
}

func (o *AdminUpdateUserEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailUnauthorized creates a AdminUpdateUserEmailUnauthorized with default headers values
func NewAdminUpdateUserEmailUnauthorized() *AdminUpdateUserEmailUnauthorized {
	return &AdminUpdateUserEmailUnauthorized{}
}

/*
AdminUpdateUserEmailUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUpdateUserEmailUnauthorized struct {
	Payload *AdminUpdateUserEmailUnauthorizedBody
}

// IsSuccess returns true when this admin update user email unauthorized response has a 2xx status code
func (o *AdminUpdateUserEmailUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user email unauthorized response has a 3xx status code
func (o *AdminUpdateUserEmailUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email unauthorized response has a 4xx status code
func (o *AdminUpdateUserEmailUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user email unauthorized response has a 5xx status code
func (o *AdminUpdateUserEmailUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user email unauthorized response a status code equal to that given
func (o *AdminUpdateUserEmailUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminUpdateUserEmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateUserEmailUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUpdateUserEmailUnauthorized) GetPayload() *AdminUpdateUserEmailUnauthorizedBody {
	return o.Payload
}

func (o *AdminUpdateUserEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AdminUpdateUserEmailUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailForbidden creates a AdminUpdateUserEmailForbidden with default headers values
func NewAdminUpdateUserEmailForbidden() *AdminUpdateUserEmailForbidden {
	return &AdminUpdateUserEmailForbidden{}
}

/*
AdminUpdateUserEmailForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUpdateUserEmailForbidden struct {
	Payload *AdminUpdateUserEmailForbiddenBody
}

// IsSuccess returns true when this admin update user email forbidden response has a 2xx status code
func (o *AdminUpdateUserEmailForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user email forbidden response has a 3xx status code
func (o *AdminUpdateUserEmailForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email forbidden response has a 4xx status code
func (o *AdminUpdateUserEmailForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user email forbidden response has a 5xx status code
func (o *AdminUpdateUserEmailForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user email forbidden response a status code equal to that given
func (o *AdminUpdateUserEmailForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminUpdateUserEmailForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateUserEmailForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailForbidden  %+v", 403, o.Payload)
}

func (o *AdminUpdateUserEmailForbidden) GetPayload() *AdminUpdateUserEmailForbiddenBody {
	return o.Payload
}

func (o *AdminUpdateUserEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AdminUpdateUserEmailForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailNotFound creates a AdminUpdateUserEmailNotFound with default headers values
func NewAdminUpdateUserEmailNotFound() *AdminUpdateUserEmailNotFound {
	return &AdminUpdateUserEmailNotFound{}
}

/*
AdminUpdateUserEmailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUpdateUserEmailNotFound struct {
	Payload *AdminUpdateUserEmailNotFoundBody
}

// IsSuccess returns true when this admin update user email not found response has a 2xx status code
func (o *AdminUpdateUserEmailNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user email not found response has a 3xx status code
func (o *AdminUpdateUserEmailNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email not found response has a 4xx status code
func (o *AdminUpdateUserEmailNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user email not found response has a 5xx status code
func (o *AdminUpdateUserEmailNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user email not found response a status code equal to that given
func (o *AdminUpdateUserEmailNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminUpdateUserEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateUserEmailNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailNotFound  %+v", 404, o.Payload)
}

func (o *AdminUpdateUserEmailNotFound) GetPayload() *AdminUpdateUserEmailNotFoundBody {
	return o.Payload
}

func (o *AdminUpdateUserEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AdminUpdateUserEmailNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserEmailInternalServerError creates a AdminUpdateUserEmailInternalServerError with default headers values
func NewAdminUpdateUserEmailInternalServerError() *AdminUpdateUserEmailInternalServerError {
	return &AdminUpdateUserEmailInternalServerError{}
}

/*
AdminUpdateUserEmailInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUpdateUserEmailInternalServerError struct {
}

// IsSuccess returns true when this admin update user email internal server error response has a 2xx status code
func (o *AdminUpdateUserEmailInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user email internal server error response has a 3xx status code
func (o *AdminUpdateUserEmailInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user email internal server error response has a 4xx status code
func (o *AdminUpdateUserEmailInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user email internal server error response has a 5xx status code
func (o *AdminUpdateUserEmailInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin update user email internal server error response a status code equal to that given
func (o *AdminUpdateUserEmailInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminUpdateUserEmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailInternalServerError ", 500)
}

func (o *AdminUpdateUserEmailInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/update/email][%d] adminUpdateUserEmailInternalServerError ", 500)
}

func (o *AdminUpdateUserEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
AdminUpdateUserEmailBadRequestBodyItems0 admin update user email bad request body items0
swagger:model AdminUpdateUserEmailBadRequestBodyItems0
*/
type AdminUpdateUserEmailBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this admin update user email bad request body items0
func (o *AdminUpdateUserEmailBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user email bad request body items0 based on context it is used
func (o *AdminUpdateUserEmailBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminUpdateUserEmailBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminUpdateUserEmailBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserEmailBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AdminUpdateUserEmailBody admin update user email body
swagger:model AdminUpdateUserEmailBody
*/
type AdminUpdateUserEmailBody struct {

	// email
	Email string `json:"email,omitempty"`

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this admin update user email body
func (o *AdminUpdateUserEmailBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user email body based on context it is used
func (o *AdminUpdateUserEmailBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminUpdateUserEmailBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminUpdateUserEmailBody) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserEmailBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AdminUpdateUserEmailForbiddenBody admin update user email forbidden body
swagger:model AdminUpdateUserEmailForbiddenBody
*/
type AdminUpdateUserEmailForbiddenBody struct {

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

// Validate validates this admin update user email forbidden body
func (o *AdminUpdateUserEmailForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user email forbidden body based on context it is used
func (o *AdminUpdateUserEmailForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminUpdateUserEmailForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminUpdateUserEmailForbiddenBody) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserEmailForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AdminUpdateUserEmailNotFoundBody admin update user email not found body
swagger:model AdminUpdateUserEmailNotFoundBody
*/
type AdminUpdateUserEmailNotFoundBody struct {

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

// Validate validates this admin update user email not found body
func (o *AdminUpdateUserEmailNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user email not found body based on context it is used
func (o *AdminUpdateUserEmailNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminUpdateUserEmailNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminUpdateUserEmailNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserEmailNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
AdminUpdateUserEmailUnauthorizedBody admin update user email unauthorized body
swagger:model AdminUpdateUserEmailUnauthorizedBody
*/
type AdminUpdateUserEmailUnauthorizedBody struct {

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

// Validate validates this admin update user email unauthorized body
func (o *AdminUpdateUserEmailUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin update user email unauthorized body based on context it is used
func (o *AdminUpdateUserEmailUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AdminUpdateUserEmailUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AdminUpdateUserEmailUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res AdminUpdateUserEmailUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
