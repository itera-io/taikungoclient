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

// UsersDeleteMyAccountReader is a Reader for the UsersDeleteMyAccount structure.
type UsersDeleteMyAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDeleteMyAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersDeleteMyAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersDeleteMyAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersDeleteMyAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersDeleteMyAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersDeleteMyAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersDeleteMyAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersDeleteMyAccountOK creates a UsersDeleteMyAccountOK with default headers values
func NewUsersDeleteMyAccountOK() *UsersDeleteMyAccountOK {
	return &UsersDeleteMyAccountOK{}
}

/*
UsersDeleteMyAccountOK describes a response with status code 200, with default header values.

Success
*/
type UsersDeleteMyAccountOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users delete my account o k response has a 2xx status code
func (o *UsersDeleteMyAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users delete my account o k response has a 3xx status code
func (o *UsersDeleteMyAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account o k response has a 4xx status code
func (o *UsersDeleteMyAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users delete my account o k response has a 5xx status code
func (o *UsersDeleteMyAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete my account o k response a status code equal to that given
func (o *UsersDeleteMyAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersDeleteMyAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountOK  %+v", 200, o.Payload)
}

func (o *UsersDeleteMyAccountOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountOK  %+v", 200, o.Payload)
}

func (o *UsersDeleteMyAccountOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersDeleteMyAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDeleteMyAccountBadRequest creates a UsersDeleteMyAccountBadRequest with default headers values
func NewUsersDeleteMyAccountBadRequest() *UsersDeleteMyAccountBadRequest {
	return &UsersDeleteMyAccountBadRequest{}
}

/*
UsersDeleteMyAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersDeleteMyAccountBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this users delete my account bad request response has a 2xx status code
func (o *UsersDeleteMyAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users delete my account bad request response has a 3xx status code
func (o *UsersDeleteMyAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account bad request response has a 4xx status code
func (o *UsersDeleteMyAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users delete my account bad request response has a 5xx status code
func (o *UsersDeleteMyAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete my account bad request response a status code equal to that given
func (o *UsersDeleteMyAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersDeleteMyAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDeleteMyAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDeleteMyAccountBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UsersDeleteMyAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDeleteMyAccountUnauthorized creates a UsersDeleteMyAccountUnauthorized with default headers values
func NewUsersDeleteMyAccountUnauthorized() *UsersDeleteMyAccountUnauthorized {
	return &UsersDeleteMyAccountUnauthorized{}
}

/*
UsersDeleteMyAccountUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersDeleteMyAccountUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this users delete my account unauthorized response has a 2xx status code
func (o *UsersDeleteMyAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users delete my account unauthorized response has a 3xx status code
func (o *UsersDeleteMyAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account unauthorized response has a 4xx status code
func (o *UsersDeleteMyAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users delete my account unauthorized response has a 5xx status code
func (o *UsersDeleteMyAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete my account unauthorized response a status code equal to that given
func (o *UsersDeleteMyAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersDeleteMyAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDeleteMyAccountUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDeleteMyAccountUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersDeleteMyAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDeleteMyAccountForbidden creates a UsersDeleteMyAccountForbidden with default headers values
func NewUsersDeleteMyAccountForbidden() *UsersDeleteMyAccountForbidden {
	return &UsersDeleteMyAccountForbidden{}
}

/*
UsersDeleteMyAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersDeleteMyAccountForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this users delete my account forbidden response has a 2xx status code
func (o *UsersDeleteMyAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users delete my account forbidden response has a 3xx status code
func (o *UsersDeleteMyAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account forbidden response has a 4xx status code
func (o *UsersDeleteMyAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users delete my account forbidden response has a 5xx status code
func (o *UsersDeleteMyAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete my account forbidden response a status code equal to that given
func (o *UsersDeleteMyAccountForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersDeleteMyAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountForbidden  %+v", 403, o.Payload)
}

func (o *UsersDeleteMyAccountForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountForbidden  %+v", 403, o.Payload)
}

func (o *UsersDeleteMyAccountForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersDeleteMyAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDeleteMyAccountNotFound creates a UsersDeleteMyAccountNotFound with default headers values
func NewUsersDeleteMyAccountNotFound() *UsersDeleteMyAccountNotFound {
	return &UsersDeleteMyAccountNotFound{}
}

/*
UsersDeleteMyAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersDeleteMyAccountNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this users delete my account not found response has a 2xx status code
func (o *UsersDeleteMyAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users delete my account not found response has a 3xx status code
func (o *UsersDeleteMyAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account not found response has a 4xx status code
func (o *UsersDeleteMyAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users delete my account not found response has a 5xx status code
func (o *UsersDeleteMyAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users delete my account not found response a status code equal to that given
func (o *UsersDeleteMyAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersDeleteMyAccountNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountNotFound  %+v", 404, o.Payload)
}

func (o *UsersDeleteMyAccountNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountNotFound  %+v", 404, o.Payload)
}

func (o *UsersDeleteMyAccountNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersDeleteMyAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDeleteMyAccountInternalServerError creates a UsersDeleteMyAccountInternalServerError with default headers values
func NewUsersDeleteMyAccountInternalServerError() *UsersDeleteMyAccountInternalServerError {
	return &UsersDeleteMyAccountInternalServerError{}
}

/*
UsersDeleteMyAccountInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersDeleteMyAccountInternalServerError struct {
}

// IsSuccess returns true when this users delete my account internal server error response has a 2xx status code
func (o *UsersDeleteMyAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users delete my account internal server error response has a 3xx status code
func (o *UsersDeleteMyAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users delete my account internal server error response has a 4xx status code
func (o *UsersDeleteMyAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users delete my account internal server error response has a 5xx status code
func (o *UsersDeleteMyAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users delete my account internal server error response a status code equal to that given
func (o *UsersDeleteMyAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersDeleteMyAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountInternalServerError ", 500)
}

func (o *UsersDeleteMyAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/delete][%d] usersDeleteMyAccountInternalServerError ", 500)
}

func (o *UsersDeleteMyAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
