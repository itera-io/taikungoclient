// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CheckerAwsReader is a Reader for the CheckerAws structure.
type CheckerAwsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerAwsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerAwsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerAwsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerAwsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerAwsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerAwsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerAwsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerAwsOK creates a CheckerAwsOK with default headers values
func NewCheckerAwsOK() *CheckerAwsOK {
	return &CheckerAwsOK{}
}

/*
CheckerAwsOK describes a response with status code 200, with default header values.

Success
*/
type CheckerAwsOK struct {
}

// IsSuccess returns true when this checker aws o k response has a 2xx status code
func (o *CheckerAwsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker aws o k response has a 3xx status code
func (o *CheckerAwsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws o k response has a 4xx status code
func (o *CheckerAwsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker aws o k response has a 5xx status code
func (o *CheckerAwsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker aws o k response a status code equal to that given
func (o *CheckerAwsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checker aws o k response
func (o *CheckerAwsOK) Code() int {
	return 200
}

func (o *CheckerAwsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsOK ", 200)
}

func (o *CheckerAwsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsOK ", 200)
}

func (o *CheckerAwsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckerAwsBadRequest creates a CheckerAwsBadRequest with default headers values
func NewCheckerAwsBadRequest() *CheckerAwsBadRequest {
	return &CheckerAwsBadRequest{}
}

/*
CheckerAwsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerAwsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this checker aws bad request response has a 2xx status code
func (o *CheckerAwsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker aws bad request response has a 3xx status code
func (o *CheckerAwsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws bad request response has a 4xx status code
func (o *CheckerAwsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker aws bad request response has a 5xx status code
func (o *CheckerAwsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker aws bad request response a status code equal to that given
func (o *CheckerAwsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the checker aws bad request response
func (o *CheckerAwsBadRequest) Code() int {
	return 400
}

func (o *CheckerAwsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAwsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAwsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAwsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAwsUnauthorized creates a CheckerAwsUnauthorized with default headers values
func NewCheckerAwsUnauthorized() *CheckerAwsUnauthorized {
	return &CheckerAwsUnauthorized{}
}

/*
CheckerAwsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerAwsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker aws unauthorized response has a 2xx status code
func (o *CheckerAwsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker aws unauthorized response has a 3xx status code
func (o *CheckerAwsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws unauthorized response has a 4xx status code
func (o *CheckerAwsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker aws unauthorized response has a 5xx status code
func (o *CheckerAwsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker aws unauthorized response a status code equal to that given
func (o *CheckerAwsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the checker aws unauthorized response
func (o *CheckerAwsUnauthorized) Code() int {
	return 401
}

func (o *CheckerAwsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAwsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAwsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAwsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAwsForbidden creates a CheckerAwsForbidden with default headers values
func NewCheckerAwsForbidden() *CheckerAwsForbidden {
	return &CheckerAwsForbidden{}
}

/*
CheckerAwsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerAwsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker aws forbidden response has a 2xx status code
func (o *CheckerAwsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker aws forbidden response has a 3xx status code
func (o *CheckerAwsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws forbidden response has a 4xx status code
func (o *CheckerAwsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker aws forbidden response has a 5xx status code
func (o *CheckerAwsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker aws forbidden response a status code equal to that given
func (o *CheckerAwsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the checker aws forbidden response
func (o *CheckerAwsForbidden) Code() int {
	return 403
}

func (o *CheckerAwsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAwsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAwsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAwsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAwsNotFound creates a CheckerAwsNotFound with default headers values
func NewCheckerAwsNotFound() *CheckerAwsNotFound {
	return &CheckerAwsNotFound{}
}

/*
CheckerAwsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerAwsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker aws not found response has a 2xx status code
func (o *CheckerAwsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker aws not found response has a 3xx status code
func (o *CheckerAwsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws not found response has a 4xx status code
func (o *CheckerAwsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker aws not found response has a 5xx status code
func (o *CheckerAwsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker aws not found response a status code equal to that given
func (o *CheckerAwsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the checker aws not found response
func (o *CheckerAwsNotFound) Code() int {
	return 404
}

func (o *CheckerAwsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAwsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAwsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAwsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAwsInternalServerError creates a CheckerAwsInternalServerError with default headers values
func NewCheckerAwsInternalServerError() *CheckerAwsInternalServerError {
	return &CheckerAwsInternalServerError{}
}

/*
CheckerAwsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerAwsInternalServerError struct {
}

// IsSuccess returns true when this checker aws internal server error response has a 2xx status code
func (o *CheckerAwsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker aws internal server error response has a 3xx status code
func (o *CheckerAwsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker aws internal server error response has a 4xx status code
func (o *CheckerAwsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker aws internal server error response has a 5xx status code
func (o *CheckerAwsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker aws internal server error response a status code equal to that given
func (o *CheckerAwsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the checker aws internal server error response
func (o *CheckerAwsInternalServerError) Code() int {
	return 500
}

func (o *CheckerAwsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsInternalServerError ", 500)
}

func (o *CheckerAwsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/aws][%d] checkerAwsInternalServerError ", 500)
}

func (o *CheckerAwsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
