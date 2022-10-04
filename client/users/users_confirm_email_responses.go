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

// UsersConfirmEmailReader is a Reader for the UsersConfirmEmail structure.
type UsersConfirmEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersConfirmEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersConfirmEmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersConfirmEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersConfirmEmailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersConfirmEmailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersConfirmEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersConfirmEmailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersConfirmEmailOK creates a UsersConfirmEmailOK with default headers values
func NewUsersConfirmEmailOK() *UsersConfirmEmailOK {
	return &UsersConfirmEmailOK{}
}

/*
UsersConfirmEmailOK describes a response with status code 200, with default header values.

Success
*/
type UsersConfirmEmailOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users confirm email o k response has a 2xx status code
func (o *UsersConfirmEmailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users confirm email o k response has a 3xx status code
func (o *UsersConfirmEmailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email o k response has a 4xx status code
func (o *UsersConfirmEmailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users confirm email o k response has a 5xx status code
func (o *UsersConfirmEmailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users confirm email o k response a status code equal to that given
func (o *UsersConfirmEmailOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersConfirmEmailOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailOK  %+v", 200, o.Payload)
}

func (o *UsersConfirmEmailOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailOK  %+v", 200, o.Payload)
}

func (o *UsersConfirmEmailOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersConfirmEmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersConfirmEmailBadRequest creates a UsersConfirmEmailBadRequest with default headers values
func NewUsersConfirmEmailBadRequest() *UsersConfirmEmailBadRequest {
	return &UsersConfirmEmailBadRequest{}
}

/*
UsersConfirmEmailBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersConfirmEmailBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users confirm email bad request response has a 2xx status code
func (o *UsersConfirmEmailBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users confirm email bad request response has a 3xx status code
func (o *UsersConfirmEmailBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email bad request response has a 4xx status code
func (o *UsersConfirmEmailBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users confirm email bad request response has a 5xx status code
func (o *UsersConfirmEmailBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users confirm email bad request response a status code equal to that given
func (o *UsersConfirmEmailBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersConfirmEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailBadRequest  %+v", 400, o.Payload)
}

func (o *UsersConfirmEmailBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailBadRequest  %+v", 400, o.Payload)
}

func (o *UsersConfirmEmailBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersConfirmEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersConfirmEmailUnauthorized creates a UsersConfirmEmailUnauthorized with default headers values
func NewUsersConfirmEmailUnauthorized() *UsersConfirmEmailUnauthorized {
	return &UsersConfirmEmailUnauthorized{}
}

/*
UsersConfirmEmailUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersConfirmEmailUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users confirm email unauthorized response has a 2xx status code
func (o *UsersConfirmEmailUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users confirm email unauthorized response has a 3xx status code
func (o *UsersConfirmEmailUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email unauthorized response has a 4xx status code
func (o *UsersConfirmEmailUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users confirm email unauthorized response has a 5xx status code
func (o *UsersConfirmEmailUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users confirm email unauthorized response a status code equal to that given
func (o *UsersConfirmEmailUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersConfirmEmailUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersConfirmEmailUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersConfirmEmailUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersConfirmEmailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersConfirmEmailForbidden creates a UsersConfirmEmailForbidden with default headers values
func NewUsersConfirmEmailForbidden() *UsersConfirmEmailForbidden {
	return &UsersConfirmEmailForbidden{}
}

/*
UsersConfirmEmailForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersConfirmEmailForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users confirm email forbidden response has a 2xx status code
func (o *UsersConfirmEmailForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users confirm email forbidden response has a 3xx status code
func (o *UsersConfirmEmailForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email forbidden response has a 4xx status code
func (o *UsersConfirmEmailForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users confirm email forbidden response has a 5xx status code
func (o *UsersConfirmEmailForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users confirm email forbidden response a status code equal to that given
func (o *UsersConfirmEmailForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersConfirmEmailForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailForbidden  %+v", 403, o.Payload)
}

func (o *UsersConfirmEmailForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailForbidden  %+v", 403, o.Payload)
}

func (o *UsersConfirmEmailForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersConfirmEmailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersConfirmEmailNotFound creates a UsersConfirmEmailNotFound with default headers values
func NewUsersConfirmEmailNotFound() *UsersConfirmEmailNotFound {
	return &UsersConfirmEmailNotFound{}
}

/*
UsersConfirmEmailNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersConfirmEmailNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users confirm email not found response has a 2xx status code
func (o *UsersConfirmEmailNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users confirm email not found response has a 3xx status code
func (o *UsersConfirmEmailNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email not found response has a 4xx status code
func (o *UsersConfirmEmailNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users confirm email not found response has a 5xx status code
func (o *UsersConfirmEmailNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users confirm email not found response a status code equal to that given
func (o *UsersConfirmEmailNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersConfirmEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailNotFound  %+v", 404, o.Payload)
}

func (o *UsersConfirmEmailNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailNotFound  %+v", 404, o.Payload)
}

func (o *UsersConfirmEmailNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersConfirmEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersConfirmEmailInternalServerError creates a UsersConfirmEmailInternalServerError with default headers values
func NewUsersConfirmEmailInternalServerError() *UsersConfirmEmailInternalServerError {
	return &UsersConfirmEmailInternalServerError{}
}

/*
UsersConfirmEmailInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersConfirmEmailInternalServerError struct {
}

// IsSuccess returns true when this users confirm email internal server error response has a 2xx status code
func (o *UsersConfirmEmailInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users confirm email internal server error response has a 3xx status code
func (o *UsersConfirmEmailInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users confirm email internal server error response has a 4xx status code
func (o *UsersConfirmEmailInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users confirm email internal server error response has a 5xx status code
func (o *UsersConfirmEmailInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users confirm email internal server error response a status code equal to that given
func (o *UsersConfirmEmailInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersConfirmEmailInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailInternalServerError ", 500)
}

func (o *UsersConfirmEmailInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/confirmemail][%d] usersConfirmEmailInternalServerError ", 500)
}

func (o *UsersConfirmEmailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
