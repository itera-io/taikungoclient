// Code generated by go-swagger; DO NOT EDIT.

package ticket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TicketCloseReader is a Reader for the TicketClose structure.
type TicketCloseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketCloseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketCloseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketCloseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketCloseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketCloseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketCloseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketCloseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketCloseOK creates a TicketCloseOK with default headers values
func NewTicketCloseOK() *TicketCloseOK {
	return &TicketCloseOK{}
}

/*
TicketCloseOK describes a response with status code 200, with default header values.

Success
*/
type TicketCloseOK struct {
}

// IsSuccess returns true when this ticket close o k response has a 2xx status code
func (o *TicketCloseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket close o k response has a 3xx status code
func (o *TicketCloseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close o k response has a 4xx status code
func (o *TicketCloseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket close o k response has a 5xx status code
func (o *TicketCloseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket close o k response a status code equal to that given
func (o *TicketCloseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ticket close o k response
func (o *TicketCloseOK) Code() int {
	return 200
}

func (o *TicketCloseOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseOK ", 200)
}

func (o *TicketCloseOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseOK ", 200)
}

func (o *TicketCloseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTicketCloseBadRequest creates a TicketCloseBadRequest with default headers values
func NewTicketCloseBadRequest() *TicketCloseBadRequest {
	return &TicketCloseBadRequest{}
}

/*
TicketCloseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketCloseBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket close bad request response has a 2xx status code
func (o *TicketCloseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket close bad request response has a 3xx status code
func (o *TicketCloseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close bad request response has a 4xx status code
func (o *TicketCloseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket close bad request response has a 5xx status code
func (o *TicketCloseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket close bad request response a status code equal to that given
func (o *TicketCloseBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ticket close bad request response
func (o *TicketCloseBadRequest) Code() int {
	return 400
}

func (o *TicketCloseBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseBadRequest  %+v", 400, o.Payload)
}

func (o *TicketCloseBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseBadRequest  %+v", 400, o.Payload)
}

func (o *TicketCloseBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCloseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseUnauthorized creates a TicketCloseUnauthorized with default headers values
func NewTicketCloseUnauthorized() *TicketCloseUnauthorized {
	return &TicketCloseUnauthorized{}
}

/*
TicketCloseUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketCloseUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket close unauthorized response has a 2xx status code
func (o *TicketCloseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket close unauthorized response has a 3xx status code
func (o *TicketCloseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close unauthorized response has a 4xx status code
func (o *TicketCloseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket close unauthorized response has a 5xx status code
func (o *TicketCloseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket close unauthorized response a status code equal to that given
func (o *TicketCloseUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ticket close unauthorized response
func (o *TicketCloseUnauthorized) Code() int {
	return 401
}

func (o *TicketCloseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketCloseUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketCloseUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCloseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseForbidden creates a TicketCloseForbidden with default headers values
func NewTicketCloseForbidden() *TicketCloseForbidden {
	return &TicketCloseForbidden{}
}

/*
TicketCloseForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketCloseForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket close forbidden response has a 2xx status code
func (o *TicketCloseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket close forbidden response has a 3xx status code
func (o *TicketCloseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close forbidden response has a 4xx status code
func (o *TicketCloseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket close forbidden response has a 5xx status code
func (o *TicketCloseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket close forbidden response a status code equal to that given
func (o *TicketCloseForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ticket close forbidden response
func (o *TicketCloseForbidden) Code() int {
	return 403
}

func (o *TicketCloseForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseForbidden  %+v", 403, o.Payload)
}

func (o *TicketCloseForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseForbidden  %+v", 403, o.Payload)
}

func (o *TicketCloseForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCloseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseNotFound creates a TicketCloseNotFound with default headers values
func NewTicketCloseNotFound() *TicketCloseNotFound {
	return &TicketCloseNotFound{}
}

/*
TicketCloseNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketCloseNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket close not found response has a 2xx status code
func (o *TicketCloseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket close not found response has a 3xx status code
func (o *TicketCloseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close not found response has a 4xx status code
func (o *TicketCloseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket close not found response has a 5xx status code
func (o *TicketCloseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket close not found response a status code equal to that given
func (o *TicketCloseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ticket close not found response
func (o *TicketCloseNotFound) Code() int {
	return 404
}

func (o *TicketCloseNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseNotFound  %+v", 404, o.Payload)
}

func (o *TicketCloseNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseNotFound  %+v", 404, o.Payload)
}

func (o *TicketCloseNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCloseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseInternalServerError creates a TicketCloseInternalServerError with default headers values
func NewTicketCloseInternalServerError() *TicketCloseInternalServerError {
	return &TicketCloseInternalServerError{}
}

/*
TicketCloseInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketCloseInternalServerError struct {
}

// IsSuccess returns true when this ticket close internal server error response has a 2xx status code
func (o *TicketCloseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket close internal server error response has a 3xx status code
func (o *TicketCloseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket close internal server error response has a 4xx status code
func (o *TicketCloseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket close internal server error response has a 5xx status code
func (o *TicketCloseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket close internal server error response a status code equal to that given
func (o *TicketCloseInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ticket close internal server error response
func (o *TicketCloseInternalServerError) Code() int {
	return 500
}

func (o *TicketCloseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseInternalServerError ", 500)
}

func (o *TicketCloseInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseInternalServerError ", 500)
}

func (o *TicketCloseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
