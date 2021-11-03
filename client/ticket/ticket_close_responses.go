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
	case 429:
		result := NewTicketCloseTooManyRequests()
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

/* TicketCloseOK describes a response with status code 200, with default header values.

Success
*/
type TicketCloseOK struct {
	Payload models.Unit
}

func (o *TicketCloseOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseOK  %+v", 200, o.Payload)
}
func (o *TicketCloseOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketCloseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseBadRequest creates a TicketCloseBadRequest with default headers values
func NewTicketCloseBadRequest() *TicketCloseBadRequest {
	return &TicketCloseBadRequest{}
}

/* TicketCloseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketCloseBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *TicketCloseBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseBadRequest  %+v", 400, o.Payload)
}
func (o *TicketCloseBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *TicketCloseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseUnauthorized creates a TicketCloseUnauthorized with default headers values
func NewTicketCloseUnauthorized() *TicketCloseUnauthorized {
	return &TicketCloseUnauthorized{}
}

/* TicketCloseUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketCloseUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *TicketCloseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseUnauthorized  %+v", 401, o.Payload)
}
func (o *TicketCloseUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCloseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseForbidden creates a TicketCloseForbidden with default headers values
func NewTicketCloseForbidden() *TicketCloseForbidden {
	return &TicketCloseForbidden{}
}

/* TicketCloseForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketCloseForbidden struct {
	Payload *models.ProblemDetails
}

func (o *TicketCloseForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseForbidden  %+v", 403, o.Payload)
}
func (o *TicketCloseForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCloseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseNotFound creates a TicketCloseNotFound with default headers values
func NewTicketCloseNotFound() *TicketCloseNotFound {
	return &TicketCloseNotFound{}
}

/* TicketCloseNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketCloseNotFound struct {
	Payload *models.ProblemDetails
}

func (o *TicketCloseNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseNotFound  %+v", 404, o.Payload)
}
func (o *TicketCloseNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCloseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseTooManyRequests creates a TicketCloseTooManyRequests with default headers values
func NewTicketCloseTooManyRequests() *TicketCloseTooManyRequests {
	return &TicketCloseTooManyRequests{}
}

/* TicketCloseTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type TicketCloseTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *TicketCloseTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseTooManyRequests  %+v", 429, o.Payload)
}
func (o *TicketCloseTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketCloseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketCloseInternalServerError creates a TicketCloseInternalServerError with default headers values
func NewTicketCloseInternalServerError() *TicketCloseInternalServerError {
	return &TicketCloseInternalServerError{}
}

/* TicketCloseInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketCloseInternalServerError struct {
}

func (o *TicketCloseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/close][%d] ticketCloseInternalServerError ", 500)
}

func (o *TicketCloseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
