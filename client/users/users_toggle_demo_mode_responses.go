// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UsersToggleDemoModeReader is a Reader for the UsersToggleDemoMode structure.
type UsersToggleDemoModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersToggleDemoModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersToggleDemoModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersToggleDemoModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersToggleDemoModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersToggleDemoModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersToggleDemoModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersToggleDemoModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersToggleDemoModeOK creates a UsersToggleDemoModeOK with default headers values
func NewUsersToggleDemoModeOK() *UsersToggleDemoModeOK {
	return &UsersToggleDemoModeOK{}
}

/*
UsersToggleDemoModeOK describes a response with status code 200, with default header values.

Success
*/
type UsersToggleDemoModeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users toggle demo mode o k response has a 2xx status code
func (o *UsersToggleDemoModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users toggle demo mode o k response has a 3xx status code
func (o *UsersToggleDemoModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode o k response has a 4xx status code
func (o *UsersToggleDemoModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users toggle demo mode o k response has a 5xx status code
func (o *UsersToggleDemoModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle demo mode o k response a status code equal to that given
func (o *UsersToggleDemoModeOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersToggleDemoModeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeOK  %+v", 200, o.Payload)
}

func (o *UsersToggleDemoModeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeOK  %+v", 200, o.Payload)
}

func (o *UsersToggleDemoModeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersToggleDemoModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleDemoModeBadRequest creates a UsersToggleDemoModeBadRequest with default headers values
func NewUsersToggleDemoModeBadRequest() *UsersToggleDemoModeBadRequest {
	return &UsersToggleDemoModeBadRequest{}
}

/*
UsersToggleDemoModeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersToggleDemoModeBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users toggle demo mode bad request response has a 2xx status code
func (o *UsersToggleDemoModeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle demo mode bad request response has a 3xx status code
func (o *UsersToggleDemoModeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode bad request response has a 4xx status code
func (o *UsersToggleDemoModeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle demo mode bad request response has a 5xx status code
func (o *UsersToggleDemoModeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle demo mode bad request response a status code equal to that given
func (o *UsersToggleDemoModeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersToggleDemoModeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeBadRequest  %+v", 400, o.Payload)
}

func (o *UsersToggleDemoModeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeBadRequest  %+v", 400, o.Payload)
}

func (o *UsersToggleDemoModeBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersToggleDemoModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleDemoModeUnauthorized creates a UsersToggleDemoModeUnauthorized with default headers values
func NewUsersToggleDemoModeUnauthorized() *UsersToggleDemoModeUnauthorized {
	return &UsersToggleDemoModeUnauthorized{}
}

/*
UsersToggleDemoModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersToggleDemoModeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users toggle demo mode unauthorized response has a 2xx status code
func (o *UsersToggleDemoModeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle demo mode unauthorized response has a 3xx status code
func (o *UsersToggleDemoModeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode unauthorized response has a 4xx status code
func (o *UsersToggleDemoModeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle demo mode unauthorized response has a 5xx status code
func (o *UsersToggleDemoModeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle demo mode unauthorized response a status code equal to that given
func (o *UsersToggleDemoModeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersToggleDemoModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersToggleDemoModeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersToggleDemoModeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleDemoModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleDemoModeForbidden creates a UsersToggleDemoModeForbidden with default headers values
func NewUsersToggleDemoModeForbidden() *UsersToggleDemoModeForbidden {
	return &UsersToggleDemoModeForbidden{}
}

/*
UsersToggleDemoModeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersToggleDemoModeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users toggle demo mode forbidden response has a 2xx status code
func (o *UsersToggleDemoModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle demo mode forbidden response has a 3xx status code
func (o *UsersToggleDemoModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode forbidden response has a 4xx status code
func (o *UsersToggleDemoModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle demo mode forbidden response has a 5xx status code
func (o *UsersToggleDemoModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle demo mode forbidden response a status code equal to that given
func (o *UsersToggleDemoModeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersToggleDemoModeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeForbidden  %+v", 403, o.Payload)
}

func (o *UsersToggleDemoModeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeForbidden  %+v", 403, o.Payload)
}

func (o *UsersToggleDemoModeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleDemoModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleDemoModeNotFound creates a UsersToggleDemoModeNotFound with default headers values
func NewUsersToggleDemoModeNotFound() *UsersToggleDemoModeNotFound {
	return &UsersToggleDemoModeNotFound{}
}

/*
UsersToggleDemoModeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersToggleDemoModeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users toggle demo mode not found response has a 2xx status code
func (o *UsersToggleDemoModeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle demo mode not found response has a 3xx status code
func (o *UsersToggleDemoModeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode not found response has a 4xx status code
func (o *UsersToggleDemoModeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle demo mode not found response has a 5xx status code
func (o *UsersToggleDemoModeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle demo mode not found response a status code equal to that given
func (o *UsersToggleDemoModeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersToggleDemoModeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeNotFound  %+v", 404, o.Payload)
}

func (o *UsersToggleDemoModeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeNotFound  %+v", 404, o.Payload)
}

func (o *UsersToggleDemoModeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleDemoModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleDemoModeInternalServerError creates a UsersToggleDemoModeInternalServerError with default headers values
func NewUsersToggleDemoModeInternalServerError() *UsersToggleDemoModeInternalServerError {
	return &UsersToggleDemoModeInternalServerError{}
}

/*
UsersToggleDemoModeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersToggleDemoModeInternalServerError struct {
}

// IsSuccess returns true when this users toggle demo mode internal server error response has a 2xx status code
func (o *UsersToggleDemoModeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle demo mode internal server error response has a 3xx status code
func (o *UsersToggleDemoModeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle demo mode internal server error response has a 4xx status code
func (o *UsersToggleDemoModeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users toggle demo mode internal server error response has a 5xx status code
func (o *UsersToggleDemoModeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users toggle demo mode internal server error response a status code equal to that given
func (o *UsersToggleDemoModeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersToggleDemoModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeInternalServerError ", 500)
}

func (o *UsersToggleDemoModeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/toggle-demo-mode][%d] usersToggleDemoModeInternalServerError ", 500)
}

func (o *UsersToggleDemoModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
