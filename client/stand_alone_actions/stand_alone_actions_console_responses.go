// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneActionsConsoleReader is a Reader for the StandAloneActionsConsole structure.
type StandAloneActionsConsoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsConsoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsConsoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsConsoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsConsoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsConsoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsConsoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsConsoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsConsoleOK creates a StandAloneActionsConsoleOK with default headers values
func NewStandAloneActionsConsoleOK() *StandAloneActionsConsoleOK {
	return &StandAloneActionsConsoleOK{}
}

/*
StandAloneActionsConsoleOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsConsoleOK struct {
	Payload string
}

// IsSuccess returns true when this stand alone actions console o k response has a 2xx status code
func (o *StandAloneActionsConsoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions console o k response has a 3xx status code
func (o *StandAloneActionsConsoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console o k response has a 4xx status code
func (o *StandAloneActionsConsoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions console o k response has a 5xx status code
func (o *StandAloneActionsConsoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions console o k response a status code equal to that given
func (o *StandAloneActionsConsoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneActionsConsoleOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsConsoleOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsConsoleOK) GetPayload() string {
	return o.Payload
}

func (o *StandAloneActionsConsoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsConsoleBadRequest creates a StandAloneActionsConsoleBadRequest with default headers values
func NewStandAloneActionsConsoleBadRequest() *StandAloneActionsConsoleBadRequest {
	return &StandAloneActionsConsoleBadRequest{}
}

/*
StandAloneActionsConsoleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsConsoleBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this stand alone actions console bad request response has a 2xx status code
func (o *StandAloneActionsConsoleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions console bad request response has a 3xx status code
func (o *StandAloneActionsConsoleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console bad request response has a 4xx status code
func (o *StandAloneActionsConsoleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions console bad request response has a 5xx status code
func (o *StandAloneActionsConsoleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions console bad request response a status code equal to that given
func (o *StandAloneActionsConsoleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneActionsConsoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsConsoleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsConsoleBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsConsoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsConsoleUnauthorized creates a StandAloneActionsConsoleUnauthorized with default headers values
func NewStandAloneActionsConsoleUnauthorized() *StandAloneActionsConsoleUnauthorized {
	return &StandAloneActionsConsoleUnauthorized{}
}

/*
StandAloneActionsConsoleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsConsoleUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions console unauthorized response has a 2xx status code
func (o *StandAloneActionsConsoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions console unauthorized response has a 3xx status code
func (o *StandAloneActionsConsoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console unauthorized response has a 4xx status code
func (o *StandAloneActionsConsoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions console unauthorized response has a 5xx status code
func (o *StandAloneActionsConsoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions console unauthorized response a status code equal to that given
func (o *StandAloneActionsConsoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneActionsConsoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsConsoleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsConsoleUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsConsoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsConsoleForbidden creates a StandAloneActionsConsoleForbidden with default headers values
func NewStandAloneActionsConsoleForbidden() *StandAloneActionsConsoleForbidden {
	return &StandAloneActionsConsoleForbidden{}
}

/*
StandAloneActionsConsoleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsConsoleForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions console forbidden response has a 2xx status code
func (o *StandAloneActionsConsoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions console forbidden response has a 3xx status code
func (o *StandAloneActionsConsoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console forbidden response has a 4xx status code
func (o *StandAloneActionsConsoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions console forbidden response has a 5xx status code
func (o *StandAloneActionsConsoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions console forbidden response a status code equal to that given
func (o *StandAloneActionsConsoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneActionsConsoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsConsoleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsConsoleForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsConsoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsConsoleNotFound creates a StandAloneActionsConsoleNotFound with default headers values
func NewStandAloneActionsConsoleNotFound() *StandAloneActionsConsoleNotFound {
	return &StandAloneActionsConsoleNotFound{}
}

/*
StandAloneActionsConsoleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsConsoleNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions console not found response has a 2xx status code
func (o *StandAloneActionsConsoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions console not found response has a 3xx status code
func (o *StandAloneActionsConsoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console not found response has a 4xx status code
func (o *StandAloneActionsConsoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions console not found response has a 5xx status code
func (o *StandAloneActionsConsoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions console not found response a status code equal to that given
func (o *StandAloneActionsConsoleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneActionsConsoleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsConsoleNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsConsoleNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsConsoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsConsoleInternalServerError creates a StandAloneActionsConsoleInternalServerError with default headers values
func NewStandAloneActionsConsoleInternalServerError() *StandAloneActionsConsoleInternalServerError {
	return &StandAloneActionsConsoleInternalServerError{}
}

/*
StandAloneActionsConsoleInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsConsoleInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions console internal server error response has a 2xx status code
func (o *StandAloneActionsConsoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions console internal server error response has a 3xx status code
func (o *StandAloneActionsConsoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions console internal server error response has a 4xx status code
func (o *StandAloneActionsConsoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions console internal server error response has a 5xx status code
func (o *StandAloneActionsConsoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions console internal server error response a status code equal to that given
func (o *StandAloneActionsConsoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneActionsConsoleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleInternalServerError ", 500)
}

func (o *StandAloneActionsConsoleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/console][%d] standAloneActionsConsoleInternalServerError ", 500)
}

func (o *StandAloneActionsConsoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
