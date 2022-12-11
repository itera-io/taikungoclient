// Code generated by go-swagger; DO NOT EDIT.

package ticket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// TicketCreateReader is a Reader for the TicketCreate structure.
type TicketCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketCreateOK creates a TicketCreateOK with default headers values
func NewTicketCreateOK() *TicketCreateOK {
	return &TicketCreateOK{}
}

/*
TicketCreateOK describes a response with status code 200, with default header values.

Success
*/
type TicketCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ticket create o k response has a 2xx status code
func (o *TicketCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket create o k response has a 3xx status code
func (o *TicketCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create o k response has a 4xx status code
func (o *TicketCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket create o k response has a 5xx status code
func (o *TicketCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket create o k response a status code equal to that given
func (o *TicketCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateOK  %+v", 200, o.Payload)
}

func (o *TicketCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateOK  %+v", 200, o.Payload)
}

func (o *TicketCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCreateBadRequest creates a TicketCreateBadRequest with default headers values
func NewTicketCreateBadRequest() *TicketCreateBadRequest {
	return &TicketCreateBadRequest{}
}

/*
TicketCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket create bad request response has a 2xx status code
func (o *TicketCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket create bad request response has a 3xx status code
func (o *TicketCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create bad request response has a 4xx status code
func (o *TicketCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket create bad request response has a 5xx status code
func (o *TicketCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket create bad request response a status code equal to that given
func (o *TicketCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TicketCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TicketCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCreateUnauthorized creates a TicketCreateUnauthorized with default headers values
func NewTicketCreateUnauthorized() *TicketCreateUnauthorized {
	return &TicketCreateUnauthorized{}
}

/*
TicketCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket create unauthorized response has a 2xx status code
func (o *TicketCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket create unauthorized response has a 3xx status code
func (o *TicketCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create unauthorized response has a 4xx status code
func (o *TicketCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket create unauthorized response has a 5xx status code
func (o *TicketCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket create unauthorized response a status code equal to that given
func (o *TicketCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCreateForbidden creates a TicketCreateForbidden with default headers values
func NewTicketCreateForbidden() *TicketCreateForbidden {
	return &TicketCreateForbidden{}
}

/*
TicketCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket create forbidden response has a 2xx status code
func (o *TicketCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket create forbidden response has a 3xx status code
func (o *TicketCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create forbidden response has a 4xx status code
func (o *TicketCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket create forbidden response has a 5xx status code
func (o *TicketCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket create forbidden response a status code equal to that given
func (o *TicketCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateForbidden  %+v", 403, o.Payload)
}

func (o *TicketCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateForbidden  %+v", 403, o.Payload)
}

func (o *TicketCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCreateNotFound creates a TicketCreateNotFound with default headers values
func NewTicketCreateNotFound() *TicketCreateNotFound {
	return &TicketCreateNotFound{}
}

/*
TicketCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket create not found response has a 2xx status code
func (o *TicketCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket create not found response has a 3xx status code
func (o *TicketCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create not found response has a 4xx status code
func (o *TicketCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket create not found response has a 5xx status code
func (o *TicketCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket create not found response a status code equal to that given
func (o *TicketCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateNotFound  %+v", 404, o.Payload)
}

func (o *TicketCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateNotFound  %+v", 404, o.Payload)
}

func (o *TicketCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCreateInternalServerError creates a TicketCreateInternalServerError with default headers values
func NewTicketCreateInternalServerError() *TicketCreateInternalServerError {
	return &TicketCreateInternalServerError{}
}

/*
TicketCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketCreateInternalServerError struct {
}

// IsSuccess returns true when this ticket create internal server error response has a 2xx status code
func (o *TicketCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket create internal server error response has a 3xx status code
func (o *TicketCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket create internal server error response has a 4xx status code
func (o *TicketCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket create internal server error response has a 5xx status code
func (o *TicketCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket create internal server error response a status code equal to that given
func (o *TicketCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateInternalServerError ", 500)
}

func (o *TicketCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/create][%d] ticketCreateInternalServerError ", 500)
}

func (o *TicketCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
