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

// UsersUpdateUserReader is a Reader for the UsersUpdateUser structure.
type UsersUpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersUpdateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUsersUpdateUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersUpdateUserOK creates a UsersUpdateUserOK with default headers values
func NewUsersUpdateUserOK() *UsersUpdateUserOK {
	return &UsersUpdateUserOK{}
}

/* UsersUpdateUserOK describes a response with status code 200, with default header values.

Success
*/
type UsersUpdateUserOK struct {
}

func (o *UsersUpdateUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserOK ", 200)
}

func (o *UsersUpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersUpdateUserBadRequest creates a UsersUpdateUserBadRequest with default headers values
func NewUsersUpdateUserBadRequest() *UsersUpdateUserBadRequest {
	return &UsersUpdateUserBadRequest{}
}

/* UsersUpdateUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersUpdateUserBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *UsersUpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserBadRequest  %+v", 400, o.Payload)
}
func (o *UsersUpdateUserBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UsersUpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateUserUnauthorized creates a UsersUpdateUserUnauthorized with default headers values
func NewUsersUpdateUserUnauthorized() *UsersUpdateUserUnauthorized {
	return &UsersUpdateUserUnauthorized{}
}

/* UsersUpdateUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersUpdateUserUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *UsersUpdateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserUnauthorized  %+v", 401, o.Payload)
}
func (o *UsersUpdateUserUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersUpdateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateUserForbidden creates a UsersUpdateUserForbidden with default headers values
func NewUsersUpdateUserForbidden() *UsersUpdateUserForbidden {
	return &UsersUpdateUserForbidden{}
}

/* UsersUpdateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersUpdateUserForbidden struct {
	Payload *models.ProblemDetails
}

func (o *UsersUpdateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserForbidden  %+v", 403, o.Payload)
}
func (o *UsersUpdateUserForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersUpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateUserNotFound creates a UsersUpdateUserNotFound with default headers values
func NewUsersUpdateUserNotFound() *UsersUpdateUserNotFound {
	return &UsersUpdateUserNotFound{}
}

/* UsersUpdateUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersUpdateUserNotFound struct {
	Payload *models.ProblemDetails
}

func (o *UsersUpdateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserNotFound  %+v", 404, o.Payload)
}
func (o *UsersUpdateUserNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersUpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateUserTooManyRequests creates a UsersUpdateUserTooManyRequests with default headers values
func NewUsersUpdateUserTooManyRequests() *UsersUpdateUserTooManyRequests {
	return &UsersUpdateUserTooManyRequests{}
}

/* UsersUpdateUserTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type UsersUpdateUserTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *UsersUpdateUserTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserTooManyRequests  %+v", 429, o.Payload)
}
func (o *UsersUpdateUserTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersUpdateUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUpdateUserInternalServerError creates a UsersUpdateUserInternalServerError with default headers values
func NewUsersUpdateUserInternalServerError() *UsersUpdateUserInternalServerError {
	return &UsersUpdateUserInternalServerError{}
}

/* UsersUpdateUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersUpdateUserInternalServerError struct {
}

func (o *UsersUpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/update][%d] usersUpdateUserInternalServerError ", 500)
}

func (o *UsersUpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
