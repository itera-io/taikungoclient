// Code generated by go-swagger; DO NOT EDIT.

package checker

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

// CheckerCronReader is a Reader for the CheckerCron structure.
type CheckerCronReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerCronReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerCronOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerCronBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerCronUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerCronForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerCronNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerCronInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerCronOK creates a CheckerCronOK with default headers values
func NewCheckerCronOK() *CheckerCronOK {
	return &CheckerCronOK{}
}

/*
CheckerCronOK describes a response with status code 200, with default header values.

Success
*/
type CheckerCronOK struct {
	Payload interface{}
}

// IsSuccess returns true when this checker cron o k response has a 2xx status code
func (o *CheckerCronOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker cron o k response has a 3xx status code
func (o *CheckerCronOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron o k response has a 4xx status code
func (o *CheckerCronOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker cron o k response has a 5xx status code
func (o *CheckerCronOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker cron o k response a status code equal to that given
func (o *CheckerCronOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerCronOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronOK  %+v", 200, o.Payload)
}

func (o *CheckerCronOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronOK  %+v", 200, o.Payload)
}

func (o *CheckerCronOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerCronOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCronBadRequest creates a CheckerCronBadRequest with default headers values
func NewCheckerCronBadRequest() *CheckerCronBadRequest {
	return &CheckerCronBadRequest{}
}

/*
CheckerCronBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerCronBadRequest struct {
	Payload []*CheckerCronBadRequestBodyItems0
}

// IsSuccess returns true when this checker cron bad request response has a 2xx status code
func (o *CheckerCronBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker cron bad request response has a 3xx status code
func (o *CheckerCronBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron bad request response has a 4xx status code
func (o *CheckerCronBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker cron bad request response has a 5xx status code
func (o *CheckerCronBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker cron bad request response a status code equal to that given
func (o *CheckerCronBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerCronBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerCronBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerCronBadRequest) GetPayload() []*CheckerCronBadRequestBodyItems0 {
	return o.Payload
}

func (o *CheckerCronBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCronUnauthorized creates a CheckerCronUnauthorized with default headers values
func NewCheckerCronUnauthorized() *CheckerCronUnauthorized {
	return &CheckerCronUnauthorized{}
}

/*
CheckerCronUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerCronUnauthorized struct {
	Payload *CheckerCronUnauthorizedBody
}

// IsSuccess returns true when this checker cron unauthorized response has a 2xx status code
func (o *CheckerCronUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker cron unauthorized response has a 3xx status code
func (o *CheckerCronUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron unauthorized response has a 4xx status code
func (o *CheckerCronUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker cron unauthorized response has a 5xx status code
func (o *CheckerCronUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker cron unauthorized response a status code equal to that given
func (o *CheckerCronUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerCronUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerCronUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerCronUnauthorized) GetPayload() *CheckerCronUnauthorizedBody {
	return o.Payload
}

func (o *CheckerCronUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerCronUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCronForbidden creates a CheckerCronForbidden with default headers values
func NewCheckerCronForbidden() *CheckerCronForbidden {
	return &CheckerCronForbidden{}
}

/*
CheckerCronForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerCronForbidden struct {
	Payload *CheckerCronForbiddenBody
}

// IsSuccess returns true when this checker cron forbidden response has a 2xx status code
func (o *CheckerCronForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker cron forbidden response has a 3xx status code
func (o *CheckerCronForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron forbidden response has a 4xx status code
func (o *CheckerCronForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker cron forbidden response has a 5xx status code
func (o *CheckerCronForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker cron forbidden response a status code equal to that given
func (o *CheckerCronForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerCronForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronForbidden  %+v", 403, o.Payload)
}

func (o *CheckerCronForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronForbidden  %+v", 403, o.Payload)
}

func (o *CheckerCronForbidden) GetPayload() *CheckerCronForbiddenBody {
	return o.Payload
}

func (o *CheckerCronForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerCronForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCronNotFound creates a CheckerCronNotFound with default headers values
func NewCheckerCronNotFound() *CheckerCronNotFound {
	return &CheckerCronNotFound{}
}

/*
CheckerCronNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerCronNotFound struct {
	Payload *CheckerCronNotFoundBody
}

// IsSuccess returns true when this checker cron not found response has a 2xx status code
func (o *CheckerCronNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker cron not found response has a 3xx status code
func (o *CheckerCronNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron not found response has a 4xx status code
func (o *CheckerCronNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker cron not found response has a 5xx status code
func (o *CheckerCronNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker cron not found response a status code equal to that given
func (o *CheckerCronNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerCronNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronNotFound  %+v", 404, o.Payload)
}

func (o *CheckerCronNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronNotFound  %+v", 404, o.Payload)
}

func (o *CheckerCronNotFound) GetPayload() *CheckerCronNotFoundBody {
	return o.Payload
}

func (o *CheckerCronNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerCronNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerCronInternalServerError creates a CheckerCronInternalServerError with default headers values
func NewCheckerCronInternalServerError() *CheckerCronInternalServerError {
	return &CheckerCronInternalServerError{}
}

/*
CheckerCronInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerCronInternalServerError struct {
}

// IsSuccess returns true when this checker cron internal server error response has a 2xx status code
func (o *CheckerCronInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker cron internal server error response has a 3xx status code
func (o *CheckerCronInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker cron internal server error response has a 4xx status code
func (o *CheckerCronInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker cron internal server error response has a 5xx status code
func (o *CheckerCronInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker cron internal server error response a status code equal to that given
func (o *CheckerCronInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerCronInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronInternalServerError ", 500)
}

func (o *CheckerCronInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/cron][%d] checkerCronInternalServerError ", 500)
}

func (o *CheckerCronInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CheckerCronBadRequestBodyItems0 checker cron bad request body items0
swagger:model CheckerCronBadRequestBodyItems0
*/
type CheckerCronBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this checker cron bad request body items0
func (o *CheckerCronBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker cron bad request body items0 based on context it is used
func (o *CheckerCronBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerCronBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerCronBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CheckerCronBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerCronBody checker cron body
swagger:model CheckerCronBody
*/
type CheckerCronBody struct {

	// cron period
	CronPeriod string `json:"cronPeriod,omitempty"`
}

// Validate validates this checker cron body
func (o *CheckerCronBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker cron body based on context it is used
func (o *CheckerCronBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerCronBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerCronBody) UnmarshalBinary(b []byte) error {
	var res CheckerCronBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerCronForbiddenBody checker cron forbidden body
swagger:model CheckerCronForbiddenBody
*/
type CheckerCronForbiddenBody struct {

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

// Validate validates this checker cron forbidden body
func (o *CheckerCronForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker cron forbidden body based on context it is used
func (o *CheckerCronForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerCronForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerCronForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CheckerCronForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerCronNotFoundBody checker cron not found body
swagger:model CheckerCronNotFoundBody
*/
type CheckerCronNotFoundBody struct {

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

// Validate validates this checker cron not found body
func (o *CheckerCronNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker cron not found body based on context it is used
func (o *CheckerCronNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerCronNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerCronNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CheckerCronNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerCronUnauthorizedBody checker cron unauthorized body
swagger:model CheckerCronUnauthorizedBody
*/
type CheckerCronUnauthorizedBody struct {

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

// Validate validates this checker cron unauthorized body
func (o *CheckerCronUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker cron unauthorized body based on context it is used
func (o *CheckerCronUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerCronUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerCronUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CheckerCronUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
