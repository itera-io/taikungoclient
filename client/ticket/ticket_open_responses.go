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

// TicketOpenReader is a Reader for the TicketOpen structure.
type TicketOpenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketOpenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketOpenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketOpenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketOpenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketOpenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketOpenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketOpenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketOpenOK creates a TicketOpenOK with default headers values
func NewTicketOpenOK() *TicketOpenOK {
	return &TicketOpenOK{}
}

/*
TicketOpenOK describes a response with status code 200, with default header values.

Success
*/
type TicketOpenOK struct {
}

// IsSuccess returns true when this ticket open o k response has a 2xx status code
func (o *TicketOpenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket open o k response has a 3xx status code
func (o *TicketOpenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open o k response has a 4xx status code
func (o *TicketOpenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket open o k response has a 5xx status code
func (o *TicketOpenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket open o k response a status code equal to that given
func (o *TicketOpenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ticket open o k response
func (o *TicketOpenOK) Code() int {
	return 200
}

func (o *TicketOpenOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenOK ", 200)
}

func (o *TicketOpenOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenOK ", 200)
}

func (o *TicketOpenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTicketOpenBadRequest creates a TicketOpenBadRequest with default headers values
func NewTicketOpenBadRequest() *TicketOpenBadRequest {
	return &TicketOpenBadRequest{}
}

/*
TicketOpenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketOpenBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket open bad request response has a 2xx status code
func (o *TicketOpenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket open bad request response has a 3xx status code
func (o *TicketOpenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open bad request response has a 4xx status code
func (o *TicketOpenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket open bad request response has a 5xx status code
func (o *TicketOpenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket open bad request response a status code equal to that given
func (o *TicketOpenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ticket open bad request response
func (o *TicketOpenBadRequest) Code() int {
	return 400
}

func (o *TicketOpenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenBadRequest  %+v", 400, o.Payload)
}

func (o *TicketOpenBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenBadRequest  %+v", 400, o.Payload)
}

func (o *TicketOpenBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketOpenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketOpenUnauthorized creates a TicketOpenUnauthorized with default headers values
func NewTicketOpenUnauthorized() *TicketOpenUnauthorized {
	return &TicketOpenUnauthorized{}
}

/*
TicketOpenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketOpenUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket open unauthorized response has a 2xx status code
func (o *TicketOpenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket open unauthorized response has a 3xx status code
func (o *TicketOpenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open unauthorized response has a 4xx status code
func (o *TicketOpenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket open unauthorized response has a 5xx status code
func (o *TicketOpenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket open unauthorized response a status code equal to that given
func (o *TicketOpenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ticket open unauthorized response
func (o *TicketOpenUnauthorized) Code() int {
	return 401
}

func (o *TicketOpenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketOpenUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketOpenUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketOpenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketOpenForbidden creates a TicketOpenForbidden with default headers values
func NewTicketOpenForbidden() *TicketOpenForbidden {
	return &TicketOpenForbidden{}
}

/*
TicketOpenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketOpenForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket open forbidden response has a 2xx status code
func (o *TicketOpenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket open forbidden response has a 3xx status code
func (o *TicketOpenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open forbidden response has a 4xx status code
func (o *TicketOpenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket open forbidden response has a 5xx status code
func (o *TicketOpenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket open forbidden response a status code equal to that given
func (o *TicketOpenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the ticket open forbidden response
func (o *TicketOpenForbidden) Code() int {
	return 403
}

func (o *TicketOpenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenForbidden  %+v", 403, o.Payload)
}

func (o *TicketOpenForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenForbidden  %+v", 403, o.Payload)
}

func (o *TicketOpenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketOpenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketOpenNotFound creates a TicketOpenNotFound with default headers values
func NewTicketOpenNotFound() *TicketOpenNotFound {
	return &TicketOpenNotFound{}
}

/*
TicketOpenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketOpenNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket open not found response has a 2xx status code
func (o *TicketOpenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket open not found response has a 3xx status code
func (o *TicketOpenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open not found response has a 4xx status code
func (o *TicketOpenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket open not found response has a 5xx status code
func (o *TicketOpenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket open not found response a status code equal to that given
func (o *TicketOpenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ticket open not found response
func (o *TicketOpenNotFound) Code() int {
	return 404
}

func (o *TicketOpenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenNotFound  %+v", 404, o.Payload)
}

func (o *TicketOpenNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenNotFound  %+v", 404, o.Payload)
}

func (o *TicketOpenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketOpenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketOpenInternalServerError creates a TicketOpenInternalServerError with default headers values
func NewTicketOpenInternalServerError() *TicketOpenInternalServerError {
	return &TicketOpenInternalServerError{}
}

/*
TicketOpenInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketOpenInternalServerError struct {
}

// IsSuccess returns true when this ticket open internal server error response has a 2xx status code
func (o *TicketOpenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket open internal server error response has a 3xx status code
func (o *TicketOpenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket open internal server error response has a 4xx status code
func (o *TicketOpenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket open internal server error response has a 5xx status code
func (o *TicketOpenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket open internal server error response a status code equal to that given
func (o *TicketOpenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ticket open internal server error response
func (o *TicketOpenInternalServerError) Code() int {
	return 500
}

func (o *TicketOpenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenInternalServerError ", 500)
}

func (o *TicketOpenInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/open][%d] ticketOpenInternalServerError ", 500)
}

func (o *TicketOpenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
