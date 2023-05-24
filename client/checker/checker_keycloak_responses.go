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

// CheckerKeycloakReader is a Reader for the CheckerKeycloak structure.
type CheckerKeycloakReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerKeycloakReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerKeycloakOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerKeycloakBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerKeycloakUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerKeycloakForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerKeycloakNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerKeycloakInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerKeycloakOK creates a CheckerKeycloakOK with default headers values
func NewCheckerKeycloakOK() *CheckerKeycloakOK {
	return &CheckerKeycloakOK{}
}

/*
CheckerKeycloakOK describes a response with status code 200, with default header values.

Success
*/
type CheckerKeycloakOK struct {
}

// IsSuccess returns true when this checker keycloak o k response has a 2xx status code
func (o *CheckerKeycloakOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker keycloak o k response has a 3xx status code
func (o *CheckerKeycloakOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak o k response has a 4xx status code
func (o *CheckerKeycloakOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker keycloak o k response has a 5xx status code
func (o *CheckerKeycloakOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker keycloak o k response a status code equal to that given
func (o *CheckerKeycloakOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checker keycloak o k response
func (o *CheckerKeycloakOK) Code() int {
	return 200
}

func (o *CheckerKeycloakOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakOK ", 200)
}

func (o *CheckerKeycloakOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakOK ", 200)
}

func (o *CheckerKeycloakOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckerKeycloakBadRequest creates a CheckerKeycloakBadRequest with default headers values
func NewCheckerKeycloakBadRequest() *CheckerKeycloakBadRequest {
	return &CheckerKeycloakBadRequest{}
}

/*
CheckerKeycloakBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerKeycloakBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this checker keycloak bad request response has a 2xx status code
func (o *CheckerKeycloakBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker keycloak bad request response has a 3xx status code
func (o *CheckerKeycloakBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak bad request response has a 4xx status code
func (o *CheckerKeycloakBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker keycloak bad request response has a 5xx status code
func (o *CheckerKeycloakBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker keycloak bad request response a status code equal to that given
func (o *CheckerKeycloakBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the checker keycloak bad request response
func (o *CheckerKeycloakBadRequest) Code() int {
	return 400
}

func (o *CheckerKeycloakBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerKeycloakBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerKeycloakBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerKeycloakBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerKeycloakUnauthorized creates a CheckerKeycloakUnauthorized with default headers values
func NewCheckerKeycloakUnauthorized() *CheckerKeycloakUnauthorized {
	return &CheckerKeycloakUnauthorized{}
}

/*
CheckerKeycloakUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerKeycloakUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker keycloak unauthorized response has a 2xx status code
func (o *CheckerKeycloakUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker keycloak unauthorized response has a 3xx status code
func (o *CheckerKeycloakUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak unauthorized response has a 4xx status code
func (o *CheckerKeycloakUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker keycloak unauthorized response has a 5xx status code
func (o *CheckerKeycloakUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker keycloak unauthorized response a status code equal to that given
func (o *CheckerKeycloakUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the checker keycloak unauthorized response
func (o *CheckerKeycloakUnauthorized) Code() int {
	return 401
}

func (o *CheckerKeycloakUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerKeycloakUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerKeycloakUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerKeycloakUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerKeycloakForbidden creates a CheckerKeycloakForbidden with default headers values
func NewCheckerKeycloakForbidden() *CheckerKeycloakForbidden {
	return &CheckerKeycloakForbidden{}
}

/*
CheckerKeycloakForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerKeycloakForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker keycloak forbidden response has a 2xx status code
func (o *CheckerKeycloakForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker keycloak forbidden response has a 3xx status code
func (o *CheckerKeycloakForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak forbidden response has a 4xx status code
func (o *CheckerKeycloakForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker keycloak forbidden response has a 5xx status code
func (o *CheckerKeycloakForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker keycloak forbidden response a status code equal to that given
func (o *CheckerKeycloakForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the checker keycloak forbidden response
func (o *CheckerKeycloakForbidden) Code() int {
	return 403
}

func (o *CheckerKeycloakForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakForbidden  %+v", 403, o.Payload)
}

func (o *CheckerKeycloakForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakForbidden  %+v", 403, o.Payload)
}

func (o *CheckerKeycloakForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerKeycloakForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerKeycloakNotFound creates a CheckerKeycloakNotFound with default headers values
func NewCheckerKeycloakNotFound() *CheckerKeycloakNotFound {
	return &CheckerKeycloakNotFound{}
}

/*
CheckerKeycloakNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerKeycloakNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker keycloak not found response has a 2xx status code
func (o *CheckerKeycloakNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker keycloak not found response has a 3xx status code
func (o *CheckerKeycloakNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak not found response has a 4xx status code
func (o *CheckerKeycloakNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker keycloak not found response has a 5xx status code
func (o *CheckerKeycloakNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker keycloak not found response a status code equal to that given
func (o *CheckerKeycloakNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the checker keycloak not found response
func (o *CheckerKeycloakNotFound) Code() int {
	return 404
}

func (o *CheckerKeycloakNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakNotFound  %+v", 404, o.Payload)
}

func (o *CheckerKeycloakNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakNotFound  %+v", 404, o.Payload)
}

func (o *CheckerKeycloakNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerKeycloakNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerKeycloakInternalServerError creates a CheckerKeycloakInternalServerError with default headers values
func NewCheckerKeycloakInternalServerError() *CheckerKeycloakInternalServerError {
	return &CheckerKeycloakInternalServerError{}
}

/*
CheckerKeycloakInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerKeycloakInternalServerError struct {
}

// IsSuccess returns true when this checker keycloak internal server error response has a 2xx status code
func (o *CheckerKeycloakInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker keycloak internal server error response has a 3xx status code
func (o *CheckerKeycloakInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker keycloak internal server error response has a 4xx status code
func (o *CheckerKeycloakInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker keycloak internal server error response has a 5xx status code
func (o *CheckerKeycloakInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker keycloak internal server error response a status code equal to that given
func (o *CheckerKeycloakInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the checker keycloak internal server error response
func (o *CheckerKeycloakInternalServerError) Code() int {
	return 500
}

func (o *CheckerKeycloakInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakInternalServerError ", 500)
}

func (o *CheckerKeycloakInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/keycloak][%d] checkerKeycloakInternalServerError ", 500)
}

func (o *CheckerKeycloakInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}