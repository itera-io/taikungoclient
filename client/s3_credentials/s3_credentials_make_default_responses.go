// Code generated by go-swagger; DO NOT EDIT.

package s3_credentials

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

// S3CredentialsMakeDefaultReader is a Reader for the S3CredentialsMakeDefault structure.
type S3CredentialsMakeDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3CredentialsMakeDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3CredentialsMakeDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewS3CredentialsMakeDefaultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewS3CredentialsMakeDefaultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewS3CredentialsMakeDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewS3CredentialsMakeDefaultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewS3CredentialsMakeDefaultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewS3CredentialsMakeDefaultOK creates a S3CredentialsMakeDefaultOK with default headers values
func NewS3CredentialsMakeDefaultOK() *S3CredentialsMakeDefaultOK {
	return &S3CredentialsMakeDefaultOK{}
}

/*
S3CredentialsMakeDefaultOK describes a response with status code 200, with default header values.

Success
*/
type S3CredentialsMakeDefaultOK struct {
	Payload interface{}
}

// IsSuccess returns true when this s3 credentials make default o k response has a 2xx status code
func (o *S3CredentialsMakeDefaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 credentials make default o k response has a 3xx status code
func (o *S3CredentialsMakeDefaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default o k response has a 4xx status code
func (o *S3CredentialsMakeDefaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials make default o k response has a 5xx status code
func (o *S3CredentialsMakeDefaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials make default o k response a status code equal to that given
func (o *S3CredentialsMakeDefaultOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3CredentialsMakeDefaultOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsMakeDefaultOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultOK  %+v", 200, o.Payload)
}

func (o *S3CredentialsMakeDefaultOK) GetPayload() interface{} {
	return o.Payload
}

func (o *S3CredentialsMakeDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsMakeDefaultBadRequest creates a S3CredentialsMakeDefaultBadRequest with default headers values
func NewS3CredentialsMakeDefaultBadRequest() *S3CredentialsMakeDefaultBadRequest {
	return &S3CredentialsMakeDefaultBadRequest{}
}

/*
S3CredentialsMakeDefaultBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type S3CredentialsMakeDefaultBadRequest struct {
	Payload []*S3CredentialsMakeDefaultBadRequestBodyItems0
}

// IsSuccess returns true when this s3 credentials make default bad request response has a 2xx status code
func (o *S3CredentialsMakeDefaultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials make default bad request response has a 3xx status code
func (o *S3CredentialsMakeDefaultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default bad request response has a 4xx status code
func (o *S3CredentialsMakeDefaultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials make default bad request response has a 5xx status code
func (o *S3CredentialsMakeDefaultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials make default bad request response a status code equal to that given
func (o *S3CredentialsMakeDefaultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *S3CredentialsMakeDefaultBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsMakeDefaultBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultBadRequest  %+v", 400, o.Payload)
}

func (o *S3CredentialsMakeDefaultBadRequest) GetPayload() []*S3CredentialsMakeDefaultBadRequestBodyItems0 {
	return o.Payload
}

func (o *S3CredentialsMakeDefaultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsMakeDefaultUnauthorized creates a S3CredentialsMakeDefaultUnauthorized with default headers values
func NewS3CredentialsMakeDefaultUnauthorized() *S3CredentialsMakeDefaultUnauthorized {
	return &S3CredentialsMakeDefaultUnauthorized{}
}

/*
S3CredentialsMakeDefaultUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type S3CredentialsMakeDefaultUnauthorized struct {
	Payload *S3CredentialsMakeDefaultUnauthorizedBody
}

// IsSuccess returns true when this s3 credentials make default unauthorized response has a 2xx status code
func (o *S3CredentialsMakeDefaultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials make default unauthorized response has a 3xx status code
func (o *S3CredentialsMakeDefaultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default unauthorized response has a 4xx status code
func (o *S3CredentialsMakeDefaultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials make default unauthorized response has a 5xx status code
func (o *S3CredentialsMakeDefaultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials make default unauthorized response a status code equal to that given
func (o *S3CredentialsMakeDefaultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *S3CredentialsMakeDefaultUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsMakeDefaultUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultUnauthorized  %+v", 401, o.Payload)
}

func (o *S3CredentialsMakeDefaultUnauthorized) GetPayload() *S3CredentialsMakeDefaultUnauthorizedBody {
	return o.Payload
}

func (o *S3CredentialsMakeDefaultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsMakeDefaultUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsMakeDefaultForbidden creates a S3CredentialsMakeDefaultForbidden with default headers values
func NewS3CredentialsMakeDefaultForbidden() *S3CredentialsMakeDefaultForbidden {
	return &S3CredentialsMakeDefaultForbidden{}
}

/*
S3CredentialsMakeDefaultForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type S3CredentialsMakeDefaultForbidden struct {
	Payload *S3CredentialsMakeDefaultForbiddenBody
}

// IsSuccess returns true when this s3 credentials make default forbidden response has a 2xx status code
func (o *S3CredentialsMakeDefaultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials make default forbidden response has a 3xx status code
func (o *S3CredentialsMakeDefaultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default forbidden response has a 4xx status code
func (o *S3CredentialsMakeDefaultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials make default forbidden response has a 5xx status code
func (o *S3CredentialsMakeDefaultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials make default forbidden response a status code equal to that given
func (o *S3CredentialsMakeDefaultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *S3CredentialsMakeDefaultForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsMakeDefaultForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultForbidden  %+v", 403, o.Payload)
}

func (o *S3CredentialsMakeDefaultForbidden) GetPayload() *S3CredentialsMakeDefaultForbiddenBody {
	return o.Payload
}

func (o *S3CredentialsMakeDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsMakeDefaultForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsMakeDefaultNotFound creates a S3CredentialsMakeDefaultNotFound with default headers values
func NewS3CredentialsMakeDefaultNotFound() *S3CredentialsMakeDefaultNotFound {
	return &S3CredentialsMakeDefaultNotFound{}
}

/*
S3CredentialsMakeDefaultNotFound describes a response with status code 404, with default header values.

Not Found
*/
type S3CredentialsMakeDefaultNotFound struct {
	Payload *S3CredentialsMakeDefaultNotFoundBody
}

// IsSuccess returns true when this s3 credentials make default not found response has a 2xx status code
func (o *S3CredentialsMakeDefaultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials make default not found response has a 3xx status code
func (o *S3CredentialsMakeDefaultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default not found response has a 4xx status code
func (o *S3CredentialsMakeDefaultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this s3 credentials make default not found response has a 5xx status code
func (o *S3CredentialsMakeDefaultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 credentials make default not found response a status code equal to that given
func (o *S3CredentialsMakeDefaultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *S3CredentialsMakeDefaultNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsMakeDefaultNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultNotFound  %+v", 404, o.Payload)
}

func (o *S3CredentialsMakeDefaultNotFound) GetPayload() *S3CredentialsMakeDefaultNotFoundBody {
	return o.Payload
}

func (o *S3CredentialsMakeDefaultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(S3CredentialsMakeDefaultNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3CredentialsMakeDefaultInternalServerError creates a S3CredentialsMakeDefaultInternalServerError with default headers values
func NewS3CredentialsMakeDefaultInternalServerError() *S3CredentialsMakeDefaultInternalServerError {
	return &S3CredentialsMakeDefaultInternalServerError{}
}

/*
S3CredentialsMakeDefaultInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type S3CredentialsMakeDefaultInternalServerError struct {
}

// IsSuccess returns true when this s3 credentials make default internal server error response has a 2xx status code
func (o *S3CredentialsMakeDefaultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this s3 credentials make default internal server error response has a 3xx status code
func (o *S3CredentialsMakeDefaultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 credentials make default internal server error response has a 4xx status code
func (o *S3CredentialsMakeDefaultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 credentials make default internal server error response has a 5xx status code
func (o *S3CredentialsMakeDefaultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this s3 credentials make default internal server error response a status code equal to that given
func (o *S3CredentialsMakeDefaultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *S3CredentialsMakeDefaultInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultInternalServerError ", 500)
}

func (o *S3CredentialsMakeDefaultInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/S3Credentials/makedefault][%d] s3CredentialsMakeDefaultInternalServerError ", 500)
}

func (o *S3CredentialsMakeDefaultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
S3CredentialsMakeDefaultBadRequestBodyItems0 s3 credentials make default bad request body items0
swagger:model S3CredentialsMakeDefaultBadRequestBodyItems0
*/
type S3CredentialsMakeDefaultBadRequestBodyItems0 struct {

	// code
	Code string `json:"code,omitempty"`

	// description
	Description string `json:"description,omitempty"`
}

// Validate validates this s3 credentials make default bad request body items0
func (o *S3CredentialsMakeDefaultBadRequestBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials make default bad request body items0 based on context it is used
func (o *S3CredentialsMakeDefaultBadRequestBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultBadRequestBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultBadRequestBodyItems0) UnmarshalBinary(b []byte) error {
	var res S3CredentialsMakeDefaultBadRequestBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsMakeDefaultBody s3 credentials make default body
swagger:model S3CredentialsMakeDefaultBody
*/
type S3CredentialsMakeDefaultBody struct {

	// id
	ID int32 `json:"id,omitempty"`
}

// Validate validates this s3 credentials make default body
func (o *S3CredentialsMakeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials make default body based on context it is used
func (o *S3CredentialsMakeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsMakeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsMakeDefaultForbiddenBody s3 credentials make default forbidden body
swagger:model S3CredentialsMakeDefaultForbiddenBody
*/
type S3CredentialsMakeDefaultForbiddenBody struct {

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

// Validate validates this s3 credentials make default forbidden body
func (o *S3CredentialsMakeDefaultForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials make default forbidden body based on context it is used
func (o *S3CredentialsMakeDefaultForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultForbiddenBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsMakeDefaultForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsMakeDefaultNotFoundBody s3 credentials make default not found body
swagger:model S3CredentialsMakeDefaultNotFoundBody
*/
type S3CredentialsMakeDefaultNotFoundBody struct {

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

// Validate validates this s3 credentials make default not found body
func (o *S3CredentialsMakeDefaultNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials make default not found body based on context it is used
func (o *S3CredentialsMakeDefaultNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultNotFoundBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsMakeDefaultNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
S3CredentialsMakeDefaultUnauthorizedBody s3 credentials make default unauthorized body
swagger:model S3CredentialsMakeDefaultUnauthorizedBody
*/
type S3CredentialsMakeDefaultUnauthorizedBody struct {

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

// Validate validates this s3 credentials make default unauthorized body
func (o *S3CredentialsMakeDefaultUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this s3 credentials make default unauthorized body based on context it is used
func (o *S3CredentialsMakeDefaultUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *S3CredentialsMakeDefaultUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res S3CredentialsMakeDefaultUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
