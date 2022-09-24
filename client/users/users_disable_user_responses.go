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

// UsersDisableUserReader is a Reader for the UsersDisableUser structure.
type UsersDisableUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDisableUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersDisableUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersDisableUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersDisableUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersDisableUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersDisableUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersDisableUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersDisableUserOK creates a UsersDisableUserOK with default headers values
func NewUsersDisableUserOK() *UsersDisableUserOK {
	return &UsersDisableUserOK{}
}

/*
UsersDisableUserOK describes a response with status code 200, with default header values.

Success
*/
type UsersDisableUserOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users disable user o k response has a 2xx status code
func (o *UsersDisableUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users disable user o k response has a 3xx status code
func (o *UsersDisableUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user o k response has a 4xx status code
func (o *UsersDisableUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users disable user o k response has a 5xx status code
func (o *UsersDisableUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users disable user o k response a status code equal to that given
func (o *UsersDisableUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersDisableUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserOK  %+v", 200, o.Payload)
}

func (o *UsersDisableUserOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserOK  %+v", 200, o.Payload)
}

func (o *UsersDisableUserOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersDisableUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDisableUserBadRequest creates a UsersDisableUserBadRequest with default headers values
func NewUsersDisableUserBadRequest() *UsersDisableUserBadRequest {
	return &UsersDisableUserBadRequest{}
}

/*
UsersDisableUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersDisableUserBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this users disable user bad request response has a 2xx status code
func (o *UsersDisableUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users disable user bad request response has a 3xx status code
func (o *UsersDisableUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user bad request response has a 4xx status code
func (o *UsersDisableUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users disable user bad request response has a 5xx status code
func (o *UsersDisableUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users disable user bad request response a status code equal to that given
func (o *UsersDisableUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersDisableUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDisableUserBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDisableUserBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UsersDisableUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDisableUserUnauthorized creates a UsersDisableUserUnauthorized with default headers values
func NewUsersDisableUserUnauthorized() *UsersDisableUserUnauthorized {
	return &UsersDisableUserUnauthorized{}
}

/*
UsersDisableUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersDisableUserUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users disable user unauthorized response has a 2xx status code
func (o *UsersDisableUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users disable user unauthorized response has a 3xx status code
func (o *UsersDisableUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user unauthorized response has a 4xx status code
func (o *UsersDisableUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users disable user unauthorized response has a 5xx status code
func (o *UsersDisableUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users disable user unauthorized response a status code equal to that given
func (o *UsersDisableUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersDisableUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDisableUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDisableUserUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDisableUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDisableUserForbidden creates a UsersDisableUserForbidden with default headers values
func NewUsersDisableUserForbidden() *UsersDisableUserForbidden {
	return &UsersDisableUserForbidden{}
}

/*
UsersDisableUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersDisableUserForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users disable user forbidden response has a 2xx status code
func (o *UsersDisableUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users disable user forbidden response has a 3xx status code
func (o *UsersDisableUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user forbidden response has a 4xx status code
func (o *UsersDisableUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users disable user forbidden response has a 5xx status code
func (o *UsersDisableUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users disable user forbidden response a status code equal to that given
func (o *UsersDisableUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersDisableUserForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserForbidden  %+v", 403, o.Payload)
}

func (o *UsersDisableUserForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserForbidden  %+v", 403, o.Payload)
}

func (o *UsersDisableUserForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDisableUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDisableUserNotFound creates a UsersDisableUserNotFound with default headers values
func NewUsersDisableUserNotFound() *UsersDisableUserNotFound {
	return &UsersDisableUserNotFound{}
}

/*
UsersDisableUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersDisableUserNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users disable user not found response has a 2xx status code
func (o *UsersDisableUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users disable user not found response has a 3xx status code
func (o *UsersDisableUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user not found response has a 4xx status code
func (o *UsersDisableUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users disable user not found response has a 5xx status code
func (o *UsersDisableUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users disable user not found response a status code equal to that given
func (o *UsersDisableUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersDisableUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserNotFound  %+v", 404, o.Payload)
}

func (o *UsersDisableUserNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserNotFound  %+v", 404, o.Payload)
}

func (o *UsersDisableUserNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDisableUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDisableUserInternalServerError creates a UsersDisableUserInternalServerError with default headers values
func NewUsersDisableUserInternalServerError() *UsersDisableUserInternalServerError {
	return &UsersDisableUserInternalServerError{}
}

/*
UsersDisableUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersDisableUserInternalServerError struct {
}

// IsSuccess returns true when this users disable user internal server error response has a 2xx status code
func (o *UsersDisableUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users disable user internal server error response has a 3xx status code
func (o *UsersDisableUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users disable user internal server error response has a 4xx status code
func (o *UsersDisableUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users disable user internal server error response has a 5xx status code
func (o *UsersDisableUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users disable user internal server error response a status code equal to that given
func (o *UsersDisableUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersDisableUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserInternalServerError ", 500)
}

func (o *UsersDisableUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/disable][%d] usersDisableUserInternalServerError ", 500)
}

func (o *UsersDisableUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
