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

// CheckerDuplicateReader is a Reader for the CheckerDuplicate structure.
type CheckerDuplicateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerDuplicateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerDuplicateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerDuplicateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerDuplicateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerDuplicateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerDuplicateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerDuplicateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerDuplicateOK creates a CheckerDuplicateOK with default headers values
func NewCheckerDuplicateOK() *CheckerDuplicateOK {
	return &CheckerDuplicateOK{}
}

/*
CheckerDuplicateOK describes a response with status code 200, with default header values.

Success
*/
type CheckerDuplicateOK struct {
	Payload bool
}

// IsSuccess returns true when this checker duplicate o k response has a 2xx status code
func (o *CheckerDuplicateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker duplicate o k response has a 3xx status code
func (o *CheckerDuplicateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate o k response has a 4xx status code
func (o *CheckerDuplicateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker duplicate o k response has a 5xx status code
func (o *CheckerDuplicateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker duplicate o k response a status code equal to that given
func (o *CheckerDuplicateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checker duplicate o k response
func (o *CheckerDuplicateOK) Code() int {
	return 200
}

func (o *CheckerDuplicateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateOK  %+v", 200, o.Payload)
}

func (o *CheckerDuplicateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateOK  %+v", 200, o.Payload)
}

func (o *CheckerDuplicateOK) GetPayload() bool {
	return o.Payload
}

func (o *CheckerDuplicateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDuplicateBadRequest creates a CheckerDuplicateBadRequest with default headers values
func NewCheckerDuplicateBadRequest() *CheckerDuplicateBadRequest {
	return &CheckerDuplicateBadRequest{}
}

/*
CheckerDuplicateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerDuplicateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this checker duplicate bad request response has a 2xx status code
func (o *CheckerDuplicateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker duplicate bad request response has a 3xx status code
func (o *CheckerDuplicateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate bad request response has a 4xx status code
func (o *CheckerDuplicateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker duplicate bad request response has a 5xx status code
func (o *CheckerDuplicateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker duplicate bad request response a status code equal to that given
func (o *CheckerDuplicateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the checker duplicate bad request response
func (o *CheckerDuplicateBadRequest) Code() int {
	return 400
}

func (o *CheckerDuplicateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerDuplicateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerDuplicateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerDuplicateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDuplicateUnauthorized creates a CheckerDuplicateUnauthorized with default headers values
func NewCheckerDuplicateUnauthorized() *CheckerDuplicateUnauthorized {
	return &CheckerDuplicateUnauthorized{}
}

/*
CheckerDuplicateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerDuplicateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker duplicate unauthorized response has a 2xx status code
func (o *CheckerDuplicateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker duplicate unauthorized response has a 3xx status code
func (o *CheckerDuplicateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate unauthorized response has a 4xx status code
func (o *CheckerDuplicateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker duplicate unauthorized response has a 5xx status code
func (o *CheckerDuplicateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker duplicate unauthorized response a status code equal to that given
func (o *CheckerDuplicateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the checker duplicate unauthorized response
func (o *CheckerDuplicateUnauthorized) Code() int {
	return 401
}

func (o *CheckerDuplicateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerDuplicateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerDuplicateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerDuplicateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDuplicateForbidden creates a CheckerDuplicateForbidden with default headers values
func NewCheckerDuplicateForbidden() *CheckerDuplicateForbidden {
	return &CheckerDuplicateForbidden{}
}

/*
CheckerDuplicateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerDuplicateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker duplicate forbidden response has a 2xx status code
func (o *CheckerDuplicateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker duplicate forbidden response has a 3xx status code
func (o *CheckerDuplicateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate forbidden response has a 4xx status code
func (o *CheckerDuplicateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker duplicate forbidden response has a 5xx status code
func (o *CheckerDuplicateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker duplicate forbidden response a status code equal to that given
func (o *CheckerDuplicateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the checker duplicate forbidden response
func (o *CheckerDuplicateForbidden) Code() int {
	return 403
}

func (o *CheckerDuplicateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateForbidden  %+v", 403, o.Payload)
}

func (o *CheckerDuplicateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateForbidden  %+v", 403, o.Payload)
}

func (o *CheckerDuplicateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerDuplicateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDuplicateNotFound creates a CheckerDuplicateNotFound with default headers values
func NewCheckerDuplicateNotFound() *CheckerDuplicateNotFound {
	return &CheckerDuplicateNotFound{}
}

/*
CheckerDuplicateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerDuplicateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker duplicate not found response has a 2xx status code
func (o *CheckerDuplicateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker duplicate not found response has a 3xx status code
func (o *CheckerDuplicateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate not found response has a 4xx status code
func (o *CheckerDuplicateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker duplicate not found response has a 5xx status code
func (o *CheckerDuplicateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker duplicate not found response a status code equal to that given
func (o *CheckerDuplicateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the checker duplicate not found response
func (o *CheckerDuplicateNotFound) Code() int {
	return 404
}

func (o *CheckerDuplicateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateNotFound  %+v", 404, o.Payload)
}

func (o *CheckerDuplicateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateNotFound  %+v", 404, o.Payload)
}

func (o *CheckerDuplicateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerDuplicateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDuplicateInternalServerError creates a CheckerDuplicateInternalServerError with default headers values
func NewCheckerDuplicateInternalServerError() *CheckerDuplicateInternalServerError {
	return &CheckerDuplicateInternalServerError{}
}

/*
CheckerDuplicateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerDuplicateInternalServerError struct {
}

// IsSuccess returns true when this checker duplicate internal server error response has a 2xx status code
func (o *CheckerDuplicateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker duplicate internal server error response has a 3xx status code
func (o *CheckerDuplicateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker duplicate internal server error response has a 4xx status code
func (o *CheckerDuplicateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker duplicate internal server error response has a 5xx status code
func (o *CheckerDuplicateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker duplicate internal server error response a status code equal to that given
func (o *CheckerDuplicateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the checker duplicate internal server error response
func (o *CheckerDuplicateInternalServerError) Code() int {
	return 500
}

func (o *CheckerDuplicateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateInternalServerError ", 500)
}

func (o *CheckerDuplicateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/duplicate][%d] checkerDuplicateInternalServerError ", 500)
}

func (o *CheckerDuplicateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
