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

// CheckerS3Reader is a Reader for the CheckerS3 structure.
type CheckerS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerS3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerS3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerS3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerS3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerS3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerS3OK creates a CheckerS3OK with default headers values
func NewCheckerS3OK() *CheckerS3OK {
	return &CheckerS3OK{}
}

/*
CheckerS3OK describes a response with status code 200, with default header values.

Success
*/
type CheckerS3OK struct {
	Payload interface{}
}

// IsSuccess returns true when this checker s3 o k response has a 2xx status code
func (o *CheckerS3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker s3 o k response has a 3xx status code
func (o *CheckerS3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 o k response has a 4xx status code
func (o *CheckerS3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker s3 o k response has a 5xx status code
func (o *CheckerS3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker s3 o k response a status code equal to that given
func (o *CheckerS3OK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerS3OK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3OK  %+v", 200, o.Payload)
}

func (o *CheckerS3OK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3OK  %+v", 200, o.Payload)
}

func (o *CheckerS3OK) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerS3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerS3BadRequest creates a CheckerS3BadRequest with default headers values
func NewCheckerS3BadRequest() *CheckerS3BadRequest {
	return &CheckerS3BadRequest{}
}

/*
CheckerS3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerS3BadRequest struct {
	Payload []*CheckerS3BadRequestBodyItems0
}

// IsSuccess returns true when this checker s3 bad request response has a 2xx status code
func (o *CheckerS3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker s3 bad request response has a 3xx status code
func (o *CheckerS3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 bad request response has a 4xx status code
func (o *CheckerS3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker s3 bad request response has a 5xx status code
func (o *CheckerS3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker s3 bad request response a status code equal to that given
func (o *CheckerS3BadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerS3BadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3BadRequest  %+v", 400, o.Payload)
}

func (o *CheckerS3BadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3BadRequest  %+v", 400, o.Payload)
}

func (o *CheckerS3BadRequest) GetPayload() []*CheckerS3BadRequestBodyItems0 {
	return o.Payload
}

func (o *CheckerS3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerS3Unauthorized creates a CheckerS3Unauthorized with default headers values
func NewCheckerS3Unauthorized() *CheckerS3Unauthorized {
	return &CheckerS3Unauthorized{}
}

/*
CheckerS3Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerS3Unauthorized struct {
	Payload *CheckerS3UnauthorizedBody
}

// IsSuccess returns true when this checker s3 unauthorized response has a 2xx status code
func (o *CheckerS3Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker s3 unauthorized response has a 3xx status code
func (o *CheckerS3Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 unauthorized response has a 4xx status code
func (o *CheckerS3Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker s3 unauthorized response has a 5xx status code
func (o *CheckerS3Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker s3 unauthorized response a status code equal to that given
func (o *CheckerS3Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerS3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3Unauthorized  %+v", 401, o.Payload)
}

func (o *CheckerS3Unauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3Unauthorized  %+v", 401, o.Payload)
}

func (o *CheckerS3Unauthorized) GetPayload() *CheckerS3UnauthorizedBody {
	return o.Payload
}

func (o *CheckerS3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerS3UnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerS3Forbidden creates a CheckerS3Forbidden with default headers values
func NewCheckerS3Forbidden() *CheckerS3Forbidden {
	return &CheckerS3Forbidden{}
}

/*
CheckerS3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerS3Forbidden struct {
	Payload *CheckerS3ForbiddenBody
}

// IsSuccess returns true when this checker s3 forbidden response has a 2xx status code
func (o *CheckerS3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker s3 forbidden response has a 3xx status code
func (o *CheckerS3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 forbidden response has a 4xx status code
func (o *CheckerS3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker s3 forbidden response has a 5xx status code
func (o *CheckerS3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker s3 forbidden response a status code equal to that given
func (o *CheckerS3Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerS3Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3Forbidden  %+v", 403, o.Payload)
}

func (o *CheckerS3Forbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3Forbidden  %+v", 403, o.Payload)
}

func (o *CheckerS3Forbidden) GetPayload() *CheckerS3ForbiddenBody {
	return o.Payload
}

func (o *CheckerS3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerS3ForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerS3NotFound creates a CheckerS3NotFound with default headers values
func NewCheckerS3NotFound() *CheckerS3NotFound {
	return &CheckerS3NotFound{}
}

/*
CheckerS3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerS3NotFound struct {
	Payload *CheckerS3NotFoundBody
}

// IsSuccess returns true when this checker s3 not found response has a 2xx status code
func (o *CheckerS3NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker s3 not found response has a 3xx status code
func (o *CheckerS3NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 not found response has a 4xx status code
func (o *CheckerS3NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker s3 not found response has a 5xx status code
func (o *CheckerS3NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker s3 not found response a status code equal to that given
func (o *CheckerS3NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerS3NotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3NotFound  %+v", 404, o.Payload)
}

func (o *CheckerS3NotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3NotFound  %+v", 404, o.Payload)
}

func (o *CheckerS3NotFound) GetPayload() *CheckerS3NotFoundBody {
	return o.Payload
}

func (o *CheckerS3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CheckerS3NotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerS3InternalServerError creates a CheckerS3InternalServerError with default headers values
func NewCheckerS3InternalServerError() *CheckerS3InternalServerError {
	return &CheckerS3InternalServerError{}
}

/*
CheckerS3InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerS3InternalServerError struct {
}

// IsSuccess returns true when this checker s3 internal server error response has a 2xx status code
func (o *CheckerS3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker s3 internal server error response has a 3xx status code
func (o *CheckerS3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker s3 internal server error response has a 4xx status code
func (o *CheckerS3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker s3 internal server error response has a 5xx status code
func (o *CheckerS3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker s3 internal server error response a status code equal to that given
func (o *CheckerS3InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerS3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3InternalServerError ", 500)
}

func (o *CheckerS3InternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/s3][%d] checkerS3InternalServerError ", 500)
}

func (o *CheckerS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
CheckerS3BadRequestBodyItems0 checker s3 bad request body items0
swagger:model CheckerS3BadRequestBodyItems0
*/
type CheckerS3BadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this checker s3 bad request body items0
func (o *CheckerS3BadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker s3 bad request body items0 based on context it is used
func (o *CheckerS3BadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerS3BadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerS3BadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res CheckerS3BadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerS3Body checker s3 body
swagger:model CheckerS3Body
*/
type CheckerS3Body struct {

	// s3 access key Id
	S3AccessKeyID string `json:"s3AccessKeyId,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 region
	S3Region string `json:"s3Region,omitempty"`

	// s3 secret key
	S3SecretKey string `json:"s3SecretKey,omitempty"`
}

// Validate validates this checker s3 body
func (o *CheckerS3Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker s3 body based on context it is used
func (o *CheckerS3Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerS3Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerS3Body) UnmarshalBinary(b []byte) error {
	var res CheckerS3Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerS3ForbiddenBody checker s3 forbidden body
swagger:model CheckerS3ForbiddenBody
*/
type CheckerS3ForbiddenBody struct {

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

// Validate validates this checker s3 forbidden body
func (o *CheckerS3ForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker s3 forbidden body based on context it is used
func (o *CheckerS3ForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerS3ForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerS3ForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CheckerS3ForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerS3NotFoundBody checker s3 not found body
swagger:model CheckerS3NotFoundBody
*/
type CheckerS3NotFoundBody struct {

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

// Validate validates this checker s3 not found body
func (o *CheckerS3NotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker s3 not found body based on context it is used
func (o *CheckerS3NotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerS3NotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerS3NotFoundBody) UnmarshalBinary(b []byte) error {
	var res CheckerS3NotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CheckerS3UnauthorizedBody checker s3 unauthorized body
swagger:model CheckerS3UnauthorizedBody
*/
type CheckerS3UnauthorizedBody struct {

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

// Validate validates this checker s3 unauthorized body
func (o *CheckerS3UnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this checker s3 unauthorized body based on context it is used
func (o *CheckerS3UnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CheckerS3UnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CheckerS3UnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CheckerS3UnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
