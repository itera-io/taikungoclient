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

// TicketArchiveReader is a Reader for the TicketArchive structure.
type TicketArchiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketArchiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketArchiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketArchiveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketArchiveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketArchiveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketArchiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketArchiveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketArchiveOK creates a TicketArchiveOK with default headers values
func NewTicketArchiveOK() *TicketArchiveOK {
	return &TicketArchiveOK{}
}

/*
TicketArchiveOK describes a response with status code 200, with default header values.

Success
*/
type TicketArchiveOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ticket archive o k response has a 2xx status code
func (o *TicketArchiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket archive o k response has a 3xx status code
func (o *TicketArchiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive o k response has a 4xx status code
func (o *TicketArchiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket archive o k response has a 5xx status code
func (o *TicketArchiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket archive o k response a status code equal to that given
func (o *TicketArchiveOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketArchiveOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveOK  %+v", 200, o.Payload)
}

func (o *TicketArchiveOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveOK  %+v", 200, o.Payload)
}

func (o *TicketArchiveOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketArchiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveBadRequest creates a TicketArchiveBadRequest with default headers values
func NewTicketArchiveBadRequest() *TicketArchiveBadRequest {
	return &TicketArchiveBadRequest{}
}

/*
TicketArchiveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketArchiveBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this ticket archive bad request response has a 2xx status code
func (o *TicketArchiveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket archive bad request response has a 3xx status code
func (o *TicketArchiveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive bad request response has a 4xx status code
func (o *TicketArchiveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket archive bad request response has a 5xx status code
func (o *TicketArchiveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket archive bad request response a status code equal to that given
func (o *TicketArchiveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketArchiveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *TicketArchiveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveBadRequest  %+v", 400, o.Payload)
}

func (o *TicketArchiveBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *TicketArchiveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveUnauthorized creates a TicketArchiveUnauthorized with default headers values
func NewTicketArchiveUnauthorized() *TicketArchiveUnauthorized {
	return &TicketArchiveUnauthorized{}
}

/*
TicketArchiveUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketArchiveUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket archive unauthorized response has a 2xx status code
func (o *TicketArchiveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket archive unauthorized response has a 3xx status code
func (o *TicketArchiveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive unauthorized response has a 4xx status code
func (o *TicketArchiveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket archive unauthorized response has a 5xx status code
func (o *TicketArchiveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket archive unauthorized response a status code equal to that given
func (o *TicketArchiveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketArchiveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketArchiveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketArchiveUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketArchiveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveForbidden creates a TicketArchiveForbidden with default headers values
func NewTicketArchiveForbidden() *TicketArchiveForbidden {
	return &TicketArchiveForbidden{}
}

/*
TicketArchiveForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketArchiveForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket archive forbidden response has a 2xx status code
func (o *TicketArchiveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket archive forbidden response has a 3xx status code
func (o *TicketArchiveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive forbidden response has a 4xx status code
func (o *TicketArchiveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket archive forbidden response has a 5xx status code
func (o *TicketArchiveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket archive forbidden response a status code equal to that given
func (o *TicketArchiveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketArchiveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveForbidden  %+v", 403, o.Payload)
}

func (o *TicketArchiveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveForbidden  %+v", 403, o.Payload)
}

func (o *TicketArchiveForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketArchiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveNotFound creates a TicketArchiveNotFound with default headers values
func NewTicketArchiveNotFound() *TicketArchiveNotFound {
	return &TicketArchiveNotFound{}
}

/*
TicketArchiveNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketArchiveNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket archive not found response has a 2xx status code
func (o *TicketArchiveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket archive not found response has a 3xx status code
func (o *TicketArchiveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive not found response has a 4xx status code
func (o *TicketArchiveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket archive not found response has a 5xx status code
func (o *TicketArchiveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket archive not found response a status code equal to that given
func (o *TicketArchiveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketArchiveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveNotFound  %+v", 404, o.Payload)
}

func (o *TicketArchiveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveNotFound  %+v", 404, o.Payload)
}

func (o *TicketArchiveNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveInternalServerError creates a TicketArchiveInternalServerError with default headers values
func NewTicketArchiveInternalServerError() *TicketArchiveInternalServerError {
	return &TicketArchiveInternalServerError{}
}

/*
TicketArchiveInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketArchiveInternalServerError struct {
}

// IsSuccess returns true when this ticket archive internal server error response has a 2xx status code
func (o *TicketArchiveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket archive internal server error response has a 3xx status code
func (o *TicketArchiveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket archive internal server error response has a 4xx status code
func (o *TicketArchiveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket archive internal server error response has a 5xx status code
func (o *TicketArchiveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket archive internal server error response a status code equal to that given
func (o *TicketArchiveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveInternalServerError ", 500)
}

func (o *TicketArchiveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveInternalServerError ", 500)
}

func (o *TicketArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
