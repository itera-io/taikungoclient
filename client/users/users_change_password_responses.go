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

// UsersChangePasswordReader is a Reader for the UsersChangePassword structure.
type UsersChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersChangePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersChangePasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersChangePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersChangePasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersChangePasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersChangePasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersChangePasswordOK creates a UsersChangePasswordOK with default headers values
func NewUsersChangePasswordOK() *UsersChangePasswordOK {
	return &UsersChangePasswordOK{}
}

/*
UsersChangePasswordOK describes a response with status code 200, with default header values.

Success
*/
type UsersChangePasswordOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users change password o k response has a 2xx status code
func (o *UsersChangePasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users change password o k response has a 3xx status code
func (o *UsersChangePasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password o k response has a 4xx status code
func (o *UsersChangePasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users change password o k response has a 5xx status code
func (o *UsersChangePasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users change password o k response a status code equal to that given
func (o *UsersChangePasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersChangePasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordOK  %+v", 200, o.Payload)
}

func (o *UsersChangePasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordOK  %+v", 200, o.Payload)
}

func (o *UsersChangePasswordOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersChangePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersChangePasswordBadRequest creates a UsersChangePasswordBadRequest with default headers values
func NewUsersChangePasswordBadRequest() *UsersChangePasswordBadRequest {
	return &UsersChangePasswordBadRequest{}
}

/*
UsersChangePasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersChangePasswordBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this users change password bad request response has a 2xx status code
func (o *UsersChangePasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users change password bad request response has a 3xx status code
func (o *UsersChangePasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password bad request response has a 4xx status code
func (o *UsersChangePasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users change password bad request response has a 5xx status code
func (o *UsersChangePasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users change password bad request response a status code equal to that given
func (o *UsersChangePasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersChangePasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *UsersChangePasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *UsersChangePasswordBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *UsersChangePasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersChangePasswordUnauthorized creates a UsersChangePasswordUnauthorized with default headers values
func NewUsersChangePasswordUnauthorized() *UsersChangePasswordUnauthorized {
	return &UsersChangePasswordUnauthorized{}
}

/*
UsersChangePasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersChangePasswordUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users change password unauthorized response has a 2xx status code
func (o *UsersChangePasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users change password unauthorized response has a 3xx status code
func (o *UsersChangePasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password unauthorized response has a 4xx status code
func (o *UsersChangePasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users change password unauthorized response has a 5xx status code
func (o *UsersChangePasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users change password unauthorized response a status code equal to that given
func (o *UsersChangePasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersChangePasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersChangePasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersChangePasswordUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersChangePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersChangePasswordForbidden creates a UsersChangePasswordForbidden with default headers values
func NewUsersChangePasswordForbidden() *UsersChangePasswordForbidden {
	return &UsersChangePasswordForbidden{}
}

/*
UsersChangePasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersChangePasswordForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users change password forbidden response has a 2xx status code
func (o *UsersChangePasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users change password forbidden response has a 3xx status code
func (o *UsersChangePasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password forbidden response has a 4xx status code
func (o *UsersChangePasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users change password forbidden response has a 5xx status code
func (o *UsersChangePasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users change password forbidden response a status code equal to that given
func (o *UsersChangePasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersChangePasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordForbidden  %+v", 403, o.Payload)
}

func (o *UsersChangePasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordForbidden  %+v", 403, o.Payload)
}

func (o *UsersChangePasswordForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersChangePasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersChangePasswordNotFound creates a UsersChangePasswordNotFound with default headers values
func NewUsersChangePasswordNotFound() *UsersChangePasswordNotFound {
	return &UsersChangePasswordNotFound{}
}

/*
UsersChangePasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersChangePasswordNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users change password not found response has a 2xx status code
func (o *UsersChangePasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users change password not found response has a 3xx status code
func (o *UsersChangePasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password not found response has a 4xx status code
func (o *UsersChangePasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users change password not found response has a 5xx status code
func (o *UsersChangePasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users change password not found response a status code equal to that given
func (o *UsersChangePasswordNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersChangePasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordNotFound  %+v", 404, o.Payload)
}

func (o *UsersChangePasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordNotFound  %+v", 404, o.Payload)
}

func (o *UsersChangePasswordNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersChangePasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersChangePasswordInternalServerError creates a UsersChangePasswordInternalServerError with default headers values
func NewUsersChangePasswordInternalServerError() *UsersChangePasswordInternalServerError {
	return &UsersChangePasswordInternalServerError{}
}

/*
UsersChangePasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersChangePasswordInternalServerError struct {
}

// IsSuccess returns true when this users change password internal server error response has a 2xx status code
func (o *UsersChangePasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users change password internal server error response has a 3xx status code
func (o *UsersChangePasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users change password internal server error response has a 4xx status code
func (o *UsersChangePasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users change password internal server error response has a 5xx status code
func (o *UsersChangePasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users change password internal server error response a status code equal to that given
func (o *UsersChangePasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersChangePasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordInternalServerError ", 500)
}

func (o *UsersChangePasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/changepassword][%d] usersChangePasswordInternalServerError ", 500)
}

func (o *UsersChangePasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
