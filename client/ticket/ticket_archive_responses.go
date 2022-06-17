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

/* TicketArchiveOK describes a response with status code 200, with default header values.

Success
*/
type TicketArchiveOK struct {
	Payload models.Unit
}

func (o *TicketArchiveOK) Error() string {
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

/* TicketArchiveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketArchiveBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *TicketArchiveBadRequest) Error() string {
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

/* TicketArchiveUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketArchiveUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *TicketArchiveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveUnauthorized  %+v", 401, o.Payload)
}
func (o *TicketArchiveUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArchiveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveForbidden creates a TicketArchiveForbidden with default headers values
func NewTicketArchiveForbidden() *TicketArchiveForbidden {
	return &TicketArchiveForbidden{}
}

/* TicketArchiveForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketArchiveForbidden struct {
	Payload *models.ProblemDetails
}

func (o *TicketArchiveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveForbidden  %+v", 403, o.Payload)
}
func (o *TicketArchiveForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArchiveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveNotFound creates a TicketArchiveNotFound with default headers values
func NewTicketArchiveNotFound() *TicketArchiveNotFound {
	return &TicketArchiveNotFound{}
}

/* TicketArchiveNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketArchiveNotFound struct {
	Payload *models.ProblemDetails
}

func (o *TicketArchiveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveNotFound  %+v", 404, o.Payload)
}
func (o *TicketArchiveNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArchiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArchiveInternalServerError creates a TicketArchiveInternalServerError with default headers values
func NewTicketArchiveInternalServerError() *TicketArchiveInternalServerError {
	return &TicketArchiveInternalServerError{}
}

/* TicketArchiveInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketArchiveInternalServerError struct {
}

func (o *TicketArchiveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/archive][%d] ticketArchiveInternalServerError ", 500)
}

func (o *TicketArchiveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
