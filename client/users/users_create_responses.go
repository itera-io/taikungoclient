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

// UsersCreateReader is a Reader for the UsersCreate structure.
type UsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersCreateOK creates a UsersCreateOK with default headers values
func NewUsersCreateOK() *UsersCreateOK {
	return &UsersCreateOK{}
}

/*
UsersCreateOK describes a response with status code 200, with default header values.

Success
*/
type UsersCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this users create o k response has a 2xx status code
func (o *UsersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users create o k response has a 3xx status code
func (o *UsersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create o k response has a 4xx status code
func (o *UsersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users create o k response has a 5xx status code
func (o *UsersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users create o k response a status code equal to that given
func (o *UsersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateOK  %+v", 200, o.Payload)
}

func (o *UsersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateOK  %+v", 200, o.Payload)
}

func (o *UsersCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *UsersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateBadRequest creates a UsersCreateBadRequest with default headers values
func NewUsersCreateBadRequest() *UsersCreateBadRequest {
	return &UsersCreateBadRequest{}
}

/*
UsersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users create bad request response has a 2xx status code
func (o *UsersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users create bad request response has a 3xx status code
func (o *UsersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create bad request response has a 4xx status code
func (o *UsersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users create bad request response has a 5xx status code
func (o *UsersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users create bad request response a status code equal to that given
func (o *UsersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateUnauthorized creates a UsersCreateUnauthorized with default headers values
func NewUsersCreateUnauthorized() *UsersCreateUnauthorized {
	return &UsersCreateUnauthorized{}
}

/*
UsersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users create unauthorized response has a 2xx status code
func (o *UsersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users create unauthorized response has a 3xx status code
func (o *UsersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create unauthorized response has a 4xx status code
func (o *UsersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users create unauthorized response has a 5xx status code
func (o *UsersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users create unauthorized response a status code equal to that given
func (o *UsersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateForbidden creates a UsersCreateForbidden with default headers values
func NewUsersCreateForbidden() *UsersCreateForbidden {
	return &UsersCreateForbidden{}
}

/*
UsersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users create forbidden response has a 2xx status code
func (o *UsersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users create forbidden response has a 3xx status code
func (o *UsersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create forbidden response has a 4xx status code
func (o *UsersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users create forbidden response has a 5xx status code
func (o *UsersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users create forbidden response a status code equal to that given
func (o *UsersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateForbidden  %+v", 403, o.Payload)
}

func (o *UsersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateForbidden  %+v", 403, o.Payload)
}

func (o *UsersCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateNotFound creates a UsersCreateNotFound with default headers values
func NewUsersCreateNotFound() *UsersCreateNotFound {
	return &UsersCreateNotFound{}
}

/*
UsersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users create not found response has a 2xx status code
func (o *UsersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users create not found response has a 3xx status code
func (o *UsersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create not found response has a 4xx status code
func (o *UsersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users create not found response has a 5xx status code
func (o *UsersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users create not found response a status code equal to that given
func (o *UsersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateNotFound  %+v", 404, o.Payload)
}

func (o *UsersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateNotFound  %+v", 404, o.Payload)
}

func (o *UsersCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersCreateInternalServerError creates a UsersCreateInternalServerError with default headers values
func NewUsersCreateInternalServerError() *UsersCreateInternalServerError {
	return &UsersCreateInternalServerError{}
}

/*
UsersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersCreateInternalServerError struct {
}

// IsSuccess returns true when this users create internal server error response has a 2xx status code
func (o *UsersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users create internal server error response has a 3xx status code
func (o *UsersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users create internal server error response has a 4xx status code
func (o *UsersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users create internal server error response has a 5xx status code
func (o *UsersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users create internal server error response a status code equal to that given
func (o *UsersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateInternalServerError ", 500)
}

func (o *UsersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users][%d] usersCreateInternalServerError ", 500)
}

func (o *UsersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
