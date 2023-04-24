// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersForceToResetPasswordReader is a Reader for the UsersForceToResetPassword structure.
type UsersForceToResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersForceToResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersForceToResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersForceToResetPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersForceToResetPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersForceToResetPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersForceToResetPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersForceToResetPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersForceToResetPasswordOK creates a UsersForceToResetPasswordOK with default headers values
func NewUsersForceToResetPasswordOK() *UsersForceToResetPasswordOK {
	return &UsersForceToResetPasswordOK{}
}

/*
UsersForceToResetPasswordOK describes a response with status code 200, with default header values.

Success
*/
type UsersForceToResetPasswordOK struct {
}

// IsSuccess returns true when this users force to reset password o k response has a 2xx status code
func (o *UsersForceToResetPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users force to reset password o k response has a 3xx status code
func (o *UsersForceToResetPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password o k response has a 4xx status code
func (o *UsersForceToResetPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users force to reset password o k response has a 5xx status code
func (o *UsersForceToResetPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users force to reset password o k response a status code equal to that given
func (o *UsersForceToResetPasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users force to reset password o k response
func (o *UsersForceToResetPasswordOK) Code() int {
	return 200
}

func (o *UsersForceToResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordOK ", 200)
}

func (o *UsersForceToResetPasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordOK ", 200)
}

func (o *UsersForceToResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersForceToResetPasswordBadRequest creates a UsersForceToResetPasswordBadRequest with default headers values
func NewUsersForceToResetPasswordBadRequest() *UsersForceToResetPasswordBadRequest {
	return &UsersForceToResetPasswordBadRequest{}
}

/*
UsersForceToResetPasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersForceToResetPasswordBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users force to reset password bad request response has a 2xx status code
func (o *UsersForceToResetPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users force to reset password bad request response has a 3xx status code
func (o *UsersForceToResetPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password bad request response has a 4xx status code
func (o *UsersForceToResetPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users force to reset password bad request response has a 5xx status code
func (o *UsersForceToResetPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users force to reset password bad request response a status code equal to that given
func (o *UsersForceToResetPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the users force to reset password bad request response
func (o *UsersForceToResetPasswordBadRequest) Code() int {
	return 400
}

func (o *UsersForceToResetPasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *UsersForceToResetPasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *UsersForceToResetPasswordBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersForceToResetPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersForceToResetPasswordUnauthorized creates a UsersForceToResetPasswordUnauthorized with default headers values
func NewUsersForceToResetPasswordUnauthorized() *UsersForceToResetPasswordUnauthorized {
	return &UsersForceToResetPasswordUnauthorized{}
}

/*
UsersForceToResetPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersForceToResetPasswordUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this users force to reset password unauthorized response has a 2xx status code
func (o *UsersForceToResetPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users force to reset password unauthorized response has a 3xx status code
func (o *UsersForceToResetPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password unauthorized response has a 4xx status code
func (o *UsersForceToResetPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users force to reset password unauthorized response has a 5xx status code
func (o *UsersForceToResetPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users force to reset password unauthorized response a status code equal to that given
func (o *UsersForceToResetPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the users force to reset password unauthorized response
func (o *UsersForceToResetPasswordUnauthorized) Code() int {
	return 401
}

func (o *UsersForceToResetPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersForceToResetPasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersForceToResetPasswordUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersForceToResetPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersForceToResetPasswordForbidden creates a UsersForceToResetPasswordForbidden with default headers values
func NewUsersForceToResetPasswordForbidden() *UsersForceToResetPasswordForbidden {
	return &UsersForceToResetPasswordForbidden{}
}

/*
UsersForceToResetPasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersForceToResetPasswordForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this users force to reset password forbidden response has a 2xx status code
func (o *UsersForceToResetPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users force to reset password forbidden response has a 3xx status code
func (o *UsersForceToResetPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password forbidden response has a 4xx status code
func (o *UsersForceToResetPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users force to reset password forbidden response has a 5xx status code
func (o *UsersForceToResetPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users force to reset password forbidden response a status code equal to that given
func (o *UsersForceToResetPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the users force to reset password forbidden response
func (o *UsersForceToResetPasswordForbidden) Code() int {
	return 403
}

func (o *UsersForceToResetPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *UsersForceToResetPasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *UsersForceToResetPasswordForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersForceToResetPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersForceToResetPasswordNotFound creates a UsersForceToResetPasswordNotFound with default headers values
func NewUsersForceToResetPasswordNotFound() *UsersForceToResetPasswordNotFound {
	return &UsersForceToResetPasswordNotFound{}
}

/*
UsersForceToResetPasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersForceToResetPasswordNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this users force to reset password not found response has a 2xx status code
func (o *UsersForceToResetPasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users force to reset password not found response has a 3xx status code
func (o *UsersForceToResetPasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password not found response has a 4xx status code
func (o *UsersForceToResetPasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users force to reset password not found response has a 5xx status code
func (o *UsersForceToResetPasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users force to reset password not found response a status code equal to that given
func (o *UsersForceToResetPasswordNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the users force to reset password not found response
func (o *UsersForceToResetPasswordNotFound) Code() int {
	return 404
}

func (o *UsersForceToResetPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *UsersForceToResetPasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *UsersForceToResetPasswordNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersForceToResetPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersForceToResetPasswordInternalServerError creates a UsersForceToResetPasswordInternalServerError with default headers values
func NewUsersForceToResetPasswordInternalServerError() *UsersForceToResetPasswordInternalServerError {
	return &UsersForceToResetPasswordInternalServerError{}
}

/*
UsersForceToResetPasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersForceToResetPasswordInternalServerError struct {
}

// IsSuccess returns true when this users force to reset password internal server error response has a 2xx status code
func (o *UsersForceToResetPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users force to reset password internal server error response has a 3xx status code
func (o *UsersForceToResetPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users force to reset password internal server error response has a 4xx status code
func (o *UsersForceToResetPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users force to reset password internal server error response has a 5xx status code
func (o *UsersForceToResetPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users force to reset password internal server error response a status code equal to that given
func (o *UsersForceToResetPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the users force to reset password internal server error response
func (o *UsersForceToResetPasswordInternalServerError) Code() int {
	return 500
}

func (o *UsersForceToResetPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordInternalServerError ", 500)
}

func (o *UsersForceToResetPasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/force-to-reset][%d] usersForceToResetPasswordInternalServerError ", 500)
}

func (o *UsersForceToResetPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
