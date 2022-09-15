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

// TicketTransferReader is a Reader for the TicketTransfer structure.
type TicketTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketTransferBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketTransferUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketTransferForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketTransferNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketTransferInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketTransferOK creates a TicketTransferOK with default headers values
func NewTicketTransferOK() *TicketTransferOK {
	return &TicketTransferOK{}
}

/*
TicketTransferOK describes a response with status code 200, with default header values.

Success
*/
type TicketTransferOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ticket transfer o k response has a 2xx status code
func (o *TicketTransferOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket transfer o k response has a 3xx status code
func (o *TicketTransferOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer o k response has a 4xx status code
func (o *TicketTransferOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket transfer o k response has a 5xx status code
func (o *TicketTransferOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer o k response a status code equal to that given
func (o *TicketTransferOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketTransferOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferOK  %+v", 200, o.Payload)
}

func (o *TicketTransferOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferOK  %+v", 200, o.Payload)
}

func (o *TicketTransferOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferBadRequest creates a TicketTransferBadRequest with default headers values
func NewTicketTransferBadRequest() *TicketTransferBadRequest {
	return &TicketTransferBadRequest{}
}

/*
TicketTransferBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketTransferBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this ticket transfer bad request response has a 2xx status code
func (o *TicketTransferBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer bad request response has a 3xx status code
func (o *TicketTransferBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer bad request response has a 4xx status code
func (o *TicketTransferBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer bad request response has a 5xx status code
func (o *TicketTransferBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer bad request response a status code equal to that given
func (o *TicketTransferBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketTransferBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTransferBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTransferBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *TicketTransferBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferUnauthorized creates a TicketTransferUnauthorized with default headers values
func NewTicketTransferUnauthorized() *TicketTransferUnauthorized {
	return &TicketTransferUnauthorized{}
}

/*
TicketTransferUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketTransferUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket transfer unauthorized response has a 2xx status code
func (o *TicketTransferUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer unauthorized response has a 3xx status code
func (o *TicketTransferUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer unauthorized response has a 4xx status code
func (o *TicketTransferUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer unauthorized response has a 5xx status code
func (o *TicketTransferUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer unauthorized response a status code equal to that given
func (o *TicketTransferUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketTransferUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTransferUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTransferUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketTransferUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferForbidden creates a TicketTransferForbidden with default headers values
func NewTicketTransferForbidden() *TicketTransferForbidden {
	return &TicketTransferForbidden{}
}

/*
TicketTransferForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketTransferForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket transfer forbidden response has a 2xx status code
func (o *TicketTransferForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer forbidden response has a 3xx status code
func (o *TicketTransferForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer forbidden response has a 4xx status code
func (o *TicketTransferForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer forbidden response has a 5xx status code
func (o *TicketTransferForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer forbidden response a status code equal to that given
func (o *TicketTransferForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketTransferForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferForbidden  %+v", 403, o.Payload)
}

func (o *TicketTransferForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferForbidden  %+v", 403, o.Payload)
}

func (o *TicketTransferForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketTransferForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferNotFound creates a TicketTransferNotFound with default headers values
func NewTicketTransferNotFound() *TicketTransferNotFound {
	return &TicketTransferNotFound{}
}

/*
TicketTransferNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketTransferNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket transfer not found response has a 2xx status code
func (o *TicketTransferNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer not found response has a 3xx status code
func (o *TicketTransferNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer not found response has a 4xx status code
func (o *TicketTransferNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer not found response has a 5xx status code
func (o *TicketTransferNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer not found response a status code equal to that given
func (o *TicketTransferNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketTransferNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferNotFound  %+v", 404, o.Payload)
}

func (o *TicketTransferNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferNotFound  %+v", 404, o.Payload)
}

func (o *TicketTransferNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketTransferNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferInternalServerError creates a TicketTransferInternalServerError with default headers values
func NewTicketTransferInternalServerError() *TicketTransferInternalServerError {
	return &TicketTransferInternalServerError{}
}

/*
TicketTransferInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketTransferInternalServerError struct {
}

// IsSuccess returns true when this ticket transfer internal server error response has a 2xx status code
func (o *TicketTransferInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer internal server error response has a 3xx status code
func (o *TicketTransferInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer internal server error response has a 4xx status code
func (o *TicketTransferInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket transfer internal server error response has a 5xx status code
func (o *TicketTransferInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket transfer internal server error response a status code equal to that given
func (o *TicketTransferInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketTransferInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferInternalServerError ", 500)
}

func (o *TicketTransferInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/transfer][%d] ticketTransferInternalServerError ", 500)
}

func (o *TicketTransferInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
