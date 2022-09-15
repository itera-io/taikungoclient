// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AuthTryDemoReader is a Reader for the AuthTryDemo structure.
type AuthTryDemoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthTryDemoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthTryDemoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthTryDemoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthTryDemoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthTryDemoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthTryDemoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthTryDemoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthTryDemoOK creates a AuthTryDemoOK with default headers values
func NewAuthTryDemoOK() *AuthTryDemoOK {
	return &AuthTryDemoOK{}
}

/*
AuthTryDemoOK describes a response with status code 200, with default header values.

Success
*/
type AuthTryDemoOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this auth try demo o k response has a 2xx status code
func (o *AuthTryDemoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth try demo o k response has a 3xx status code
func (o *AuthTryDemoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo o k response has a 4xx status code
func (o *AuthTryDemoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth try demo o k response has a 5xx status code
func (o *AuthTryDemoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth try demo o k response a status code equal to that given
func (o *AuthTryDemoOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthTryDemoOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoOK  %+v", 200, o.Payload)
}

func (o *AuthTryDemoOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoOK  %+v", 200, o.Payload)
}

func (o *AuthTryDemoOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AuthTryDemoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthTryDemoBadRequest creates a AuthTryDemoBadRequest with default headers values
func NewAuthTryDemoBadRequest() *AuthTryDemoBadRequest {
	return &AuthTryDemoBadRequest{}
}

/*
AuthTryDemoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthTryDemoBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this auth try demo bad request response has a 2xx status code
func (o *AuthTryDemoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth try demo bad request response has a 3xx status code
func (o *AuthTryDemoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo bad request response has a 4xx status code
func (o *AuthTryDemoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth try demo bad request response has a 5xx status code
func (o *AuthTryDemoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this auth try demo bad request response a status code equal to that given
func (o *AuthTryDemoBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuthTryDemoBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoBadRequest  %+v", 400, o.Payload)
}

func (o *AuthTryDemoBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoBadRequest  %+v", 400, o.Payload)
}

func (o *AuthTryDemoBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AuthTryDemoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthTryDemoUnauthorized creates a AuthTryDemoUnauthorized with default headers values
func NewAuthTryDemoUnauthorized() *AuthTryDemoUnauthorized {
	return &AuthTryDemoUnauthorized{}
}

/*
AuthTryDemoUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthTryDemoUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth try demo unauthorized response has a 2xx status code
func (o *AuthTryDemoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth try demo unauthorized response has a 3xx status code
func (o *AuthTryDemoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo unauthorized response has a 4xx status code
func (o *AuthTryDemoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth try demo unauthorized response has a 5xx status code
func (o *AuthTryDemoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this auth try demo unauthorized response a status code equal to that given
func (o *AuthTryDemoUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AuthTryDemoUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthTryDemoUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthTryDemoUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthTryDemoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthTryDemoForbidden creates a AuthTryDemoForbidden with default headers values
func NewAuthTryDemoForbidden() *AuthTryDemoForbidden {
	return &AuthTryDemoForbidden{}
}

/*
AuthTryDemoForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthTryDemoForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth try demo forbidden response has a 2xx status code
func (o *AuthTryDemoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth try demo forbidden response has a 3xx status code
func (o *AuthTryDemoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo forbidden response has a 4xx status code
func (o *AuthTryDemoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth try demo forbidden response has a 5xx status code
func (o *AuthTryDemoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this auth try demo forbidden response a status code equal to that given
func (o *AuthTryDemoForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AuthTryDemoForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoForbidden  %+v", 403, o.Payload)
}

func (o *AuthTryDemoForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoForbidden  %+v", 403, o.Payload)
}

func (o *AuthTryDemoForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthTryDemoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthTryDemoNotFound creates a AuthTryDemoNotFound with default headers values
func NewAuthTryDemoNotFound() *AuthTryDemoNotFound {
	return &AuthTryDemoNotFound{}
}

/*
AuthTryDemoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthTryDemoNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth try demo not found response has a 2xx status code
func (o *AuthTryDemoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth try demo not found response has a 3xx status code
func (o *AuthTryDemoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo not found response has a 4xx status code
func (o *AuthTryDemoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth try demo not found response has a 5xx status code
func (o *AuthTryDemoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this auth try demo not found response a status code equal to that given
func (o *AuthTryDemoNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AuthTryDemoNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoNotFound  %+v", 404, o.Payload)
}

func (o *AuthTryDemoNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoNotFound  %+v", 404, o.Payload)
}

func (o *AuthTryDemoNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthTryDemoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthTryDemoInternalServerError creates a AuthTryDemoInternalServerError with default headers values
func NewAuthTryDemoInternalServerError() *AuthTryDemoInternalServerError {
	return &AuthTryDemoInternalServerError{}
}

/*
AuthTryDemoInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthTryDemoInternalServerError struct {
}

// IsSuccess returns true when this auth try demo internal server error response has a 2xx status code
func (o *AuthTryDemoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth try demo internal server error response has a 3xx status code
func (o *AuthTryDemoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth try demo internal server error response has a 4xx status code
func (o *AuthTryDemoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth try demo internal server error response has a 5xx status code
func (o *AuthTryDemoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this auth try demo internal server error response a status code equal to that given
func (o *AuthTryDemoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuthTryDemoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoInternalServerError ", 500)
}

func (o *AuthTryDemoInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/try-demo][%d] authTryDemoInternalServerError ", 500)
}

func (o *AuthTryDemoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
