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

// CheckerAzureReader is a Reader for the CheckerAzure structure.
type CheckerAzureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerAzureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerAzureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerAzureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerAzureUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerAzureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerAzureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerAzureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerAzureOK creates a CheckerAzureOK with default headers values
func NewCheckerAzureOK() *CheckerAzureOK {
	return &CheckerAzureOK{}
}

/*
CheckerAzureOK describes a response with status code 200, with default header values.

Success
*/
type CheckerAzureOK struct {
}

// IsSuccess returns true when this checker azure o k response has a 2xx status code
func (o *CheckerAzureOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker azure o k response has a 3xx status code
func (o *CheckerAzureOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure o k response has a 4xx status code
func (o *CheckerAzureOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker azure o k response has a 5xx status code
func (o *CheckerAzureOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure o k response a status code equal to that given
func (o *CheckerAzureOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checker azure o k response
func (o *CheckerAzureOK) Code() int {
	return 200
}

func (o *CheckerAzureOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureOK ", 200)
}

func (o *CheckerAzureOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureOK ", 200)
}

func (o *CheckerAzureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckerAzureBadRequest creates a CheckerAzureBadRequest with default headers values
func NewCheckerAzureBadRequest() *CheckerAzureBadRequest {
	return &CheckerAzureBadRequest{}
}

/*
CheckerAzureBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerAzureBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this checker azure bad request response has a 2xx status code
func (o *CheckerAzureBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure bad request response has a 3xx status code
func (o *CheckerAzureBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure bad request response has a 4xx status code
func (o *CheckerAzureBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure bad request response has a 5xx status code
func (o *CheckerAzureBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure bad request response a status code equal to that given
func (o *CheckerAzureBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the checker azure bad request response
func (o *CheckerAzureBadRequest) Code() int {
	return 400
}

func (o *CheckerAzureBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAzureBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerAzureBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAzureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureUnauthorized creates a CheckerAzureUnauthorized with default headers values
func NewCheckerAzureUnauthorized() *CheckerAzureUnauthorized {
	return &CheckerAzureUnauthorized{}
}

/*
CheckerAzureUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerAzureUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker azure unauthorized response has a 2xx status code
func (o *CheckerAzureUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure unauthorized response has a 3xx status code
func (o *CheckerAzureUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure unauthorized response has a 4xx status code
func (o *CheckerAzureUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure unauthorized response has a 5xx status code
func (o *CheckerAzureUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure unauthorized response a status code equal to that given
func (o *CheckerAzureUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the checker azure unauthorized response
func (o *CheckerAzureUnauthorized) Code() int {
	return 401
}

func (o *CheckerAzureUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAzureUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerAzureUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAzureUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureForbidden creates a CheckerAzureForbidden with default headers values
func NewCheckerAzureForbidden() *CheckerAzureForbidden {
	return &CheckerAzureForbidden{}
}

/*
CheckerAzureForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerAzureForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker azure forbidden response has a 2xx status code
func (o *CheckerAzureForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure forbidden response has a 3xx status code
func (o *CheckerAzureForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure forbidden response has a 4xx status code
func (o *CheckerAzureForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure forbidden response has a 5xx status code
func (o *CheckerAzureForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure forbidden response a status code equal to that given
func (o *CheckerAzureForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the checker azure forbidden response
func (o *CheckerAzureForbidden) Code() int {
	return 403
}

func (o *CheckerAzureForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAzureForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureForbidden  %+v", 403, o.Payload)
}

func (o *CheckerAzureForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAzureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureNotFound creates a CheckerAzureNotFound with default headers values
func NewCheckerAzureNotFound() *CheckerAzureNotFound {
	return &CheckerAzureNotFound{}
}

/*
CheckerAzureNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerAzureNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker azure not found response has a 2xx status code
func (o *CheckerAzureNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure not found response has a 3xx status code
func (o *CheckerAzureNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure not found response has a 4xx status code
func (o *CheckerAzureNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker azure not found response has a 5xx status code
func (o *CheckerAzureNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker azure not found response a status code equal to that given
func (o *CheckerAzureNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the checker azure not found response
func (o *CheckerAzureNotFound) Code() int {
	return 404
}

func (o *CheckerAzureNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAzureNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureNotFound  %+v", 404, o.Payload)
}

func (o *CheckerAzureNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerAzureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerAzureInternalServerError creates a CheckerAzureInternalServerError with default headers values
func NewCheckerAzureInternalServerError() *CheckerAzureInternalServerError {
	return &CheckerAzureInternalServerError{}
}

/*
CheckerAzureInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerAzureInternalServerError struct {
}

// IsSuccess returns true when this checker azure internal server error response has a 2xx status code
func (o *CheckerAzureInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker azure internal server error response has a 3xx status code
func (o *CheckerAzureInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker azure internal server error response has a 4xx status code
func (o *CheckerAzureInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker azure internal server error response has a 5xx status code
func (o *CheckerAzureInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker azure internal server error response a status code equal to that given
func (o *CheckerAzureInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the checker azure internal server error response
func (o *CheckerAzureInternalServerError) Code() int {
	return 500
}

func (o *CheckerAzureInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureInternalServerError ", 500)
}

func (o *CheckerAzureInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/azure][%d] checkerAzureInternalServerError ", 500)
}

func (o *CheckerAzureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
