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

// TicketDeleteMessageReader is a Reader for the TicketDeleteMessage structure.
type TicketDeleteMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketDeleteMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketDeleteMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketDeleteMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketDeleteMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketDeleteMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketDeleteMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketDeleteMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketDeleteMessageOK creates a TicketDeleteMessageOK with default headers values
func NewTicketDeleteMessageOK() *TicketDeleteMessageOK {
	return &TicketDeleteMessageOK{}
}

/* TicketDeleteMessageOK describes a response with status code 200, with default header values.

Success
*/
type TicketDeleteMessageOK struct {
	Payload models.Unit
}

func (o *TicketDeleteMessageOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageOK  %+v", 200, o.Payload)
}
func (o *TicketDeleteMessageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketDeleteMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketDeleteMessageBadRequest creates a TicketDeleteMessageBadRequest with default headers values
func NewTicketDeleteMessageBadRequest() *TicketDeleteMessageBadRequest {
	return &TicketDeleteMessageBadRequest{}
}

/* TicketDeleteMessageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketDeleteMessageBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *TicketDeleteMessageBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageBadRequest  %+v", 400, o.Payload)
}
func (o *TicketDeleteMessageBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *TicketDeleteMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketDeleteMessageUnauthorized creates a TicketDeleteMessageUnauthorized with default headers values
func NewTicketDeleteMessageUnauthorized() *TicketDeleteMessageUnauthorized {
	return &TicketDeleteMessageUnauthorized{}
}

/* TicketDeleteMessageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketDeleteMessageUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *TicketDeleteMessageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageUnauthorized  %+v", 401, o.Payload)
}
func (o *TicketDeleteMessageUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketDeleteMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketDeleteMessageForbidden creates a TicketDeleteMessageForbidden with default headers values
func NewTicketDeleteMessageForbidden() *TicketDeleteMessageForbidden {
	return &TicketDeleteMessageForbidden{}
}

/* TicketDeleteMessageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketDeleteMessageForbidden struct {
	Payload *models.ProblemDetails
}

func (o *TicketDeleteMessageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageForbidden  %+v", 403, o.Payload)
}
func (o *TicketDeleteMessageForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketDeleteMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketDeleteMessageNotFound creates a TicketDeleteMessageNotFound with default headers values
func NewTicketDeleteMessageNotFound() *TicketDeleteMessageNotFound {
	return &TicketDeleteMessageNotFound{}
}

/* TicketDeleteMessageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketDeleteMessageNotFound struct {
	Payload *models.ProblemDetails
}

func (o *TicketDeleteMessageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageNotFound  %+v", 404, o.Payload)
}
func (o *TicketDeleteMessageNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketDeleteMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketDeleteMessageInternalServerError creates a TicketDeleteMessageInternalServerError with default headers values
func NewTicketDeleteMessageInternalServerError() *TicketDeleteMessageInternalServerError {
	return &TicketDeleteMessageInternalServerError{}
}

/* TicketDeleteMessageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketDeleteMessageInternalServerError struct {
}

func (o *TicketDeleteMessageInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Ticket/delete/message/{messageId}][%d] ticketDeleteMessageInternalServerError ", 500)
}

func (o *TicketDeleteMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
