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

// UsersToggleNotificationModeReader is a Reader for the UsersToggleNotificationMode structure.
type UsersToggleNotificationModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersToggleNotificationModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersToggleNotificationModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersToggleNotificationModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersToggleNotificationModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersToggleNotificationModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersToggleNotificationModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersToggleNotificationModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersToggleNotificationModeOK creates a UsersToggleNotificationModeOK with default headers values
func NewUsersToggleNotificationModeOK() *UsersToggleNotificationModeOK {
	return &UsersToggleNotificationModeOK{}
}

/* UsersToggleNotificationModeOK describes a response with status code 200, with default header values.

Success
*/
type UsersToggleNotificationModeOK struct {
	Payload models.Unit
}

func (o *UsersToggleNotificationModeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeOK  %+v", 200, o.Payload)
}
func (o *UsersToggleNotificationModeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersToggleNotificationModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleNotificationModeBadRequest creates a UsersToggleNotificationModeBadRequest with default headers values
func NewUsersToggleNotificationModeBadRequest() *UsersToggleNotificationModeBadRequest {
	return &UsersToggleNotificationModeBadRequest{}
}

/* UsersToggleNotificationModeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersToggleNotificationModeBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *UsersToggleNotificationModeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeBadRequest  %+v", 400, o.Payload)
}
func (o *UsersToggleNotificationModeBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UsersToggleNotificationModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleNotificationModeUnauthorized creates a UsersToggleNotificationModeUnauthorized with default headers values
func NewUsersToggleNotificationModeUnauthorized() *UsersToggleNotificationModeUnauthorized {
	return &UsersToggleNotificationModeUnauthorized{}
}

/* UsersToggleNotificationModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersToggleNotificationModeUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *UsersToggleNotificationModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeUnauthorized  %+v", 401, o.Payload)
}
func (o *UsersToggleNotificationModeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleNotificationModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleNotificationModeForbidden creates a UsersToggleNotificationModeForbidden with default headers values
func NewUsersToggleNotificationModeForbidden() *UsersToggleNotificationModeForbidden {
	return &UsersToggleNotificationModeForbidden{}
}

/* UsersToggleNotificationModeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersToggleNotificationModeForbidden struct {
	Payload *models.ProblemDetails
}

func (o *UsersToggleNotificationModeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeForbidden  %+v", 403, o.Payload)
}
func (o *UsersToggleNotificationModeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleNotificationModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleNotificationModeNotFound creates a UsersToggleNotificationModeNotFound with default headers values
func NewUsersToggleNotificationModeNotFound() *UsersToggleNotificationModeNotFound {
	return &UsersToggleNotificationModeNotFound{}
}

/* UsersToggleNotificationModeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersToggleNotificationModeNotFound struct {
	Payload *models.ProblemDetails
}

func (o *UsersToggleNotificationModeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeNotFound  %+v", 404, o.Payload)
}
func (o *UsersToggleNotificationModeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersToggleNotificationModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleNotificationModeInternalServerError creates a UsersToggleNotificationModeInternalServerError with default headers values
func NewUsersToggleNotificationModeInternalServerError() *UsersToggleNotificationModeInternalServerError {
	return &UsersToggleNotificationModeInternalServerError{}
}

/* UsersToggleNotificationModeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersToggleNotificationModeInternalServerError struct {
}

func (o *UsersToggleNotificationModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglenotificationmode][%d] usersToggleNotificationModeInternalServerError ", 500)
}

func (o *UsersToggleNotificationModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}