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

// TicketUpdateMessageReader is a Reader for the TicketUpdateMessage structure.
type TicketUpdateMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketUpdateMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketUpdateMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketUpdateMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketUpdateMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketUpdateMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketUpdateMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketUpdateMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketUpdateMessageOK creates a TicketUpdateMessageOK with default headers values
func NewTicketUpdateMessageOK() *TicketUpdateMessageOK {
	return &TicketUpdateMessageOK{}
}

/*
TicketUpdateMessageOK describes a response with status code 200, with default header values.

Success
*/
type TicketUpdateMessageOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ticket update message o k response has a 2xx status code
func (o *TicketUpdateMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket update message o k response has a 3xx status code
func (o *TicketUpdateMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message o k response has a 4xx status code
func (o *TicketUpdateMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket update message o k response has a 5xx status code
func (o *TicketUpdateMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket update message o k response a status code equal to that given
func (o *TicketUpdateMessageOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketUpdateMessageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageOK  %+v", 200, o.Payload)
}

func (o *TicketUpdateMessageOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageOK  %+v", 200, o.Payload)
}

func (o *TicketUpdateMessageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketUpdateMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketUpdateMessageBadRequest creates a TicketUpdateMessageBadRequest with default headers values
func NewTicketUpdateMessageBadRequest() *TicketUpdateMessageBadRequest {
	return &TicketUpdateMessageBadRequest{}
}

/*
TicketUpdateMessageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketUpdateMessageBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket update message bad request response has a 2xx status code
func (o *TicketUpdateMessageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket update message bad request response has a 3xx status code
func (o *TicketUpdateMessageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message bad request response has a 4xx status code
func (o *TicketUpdateMessageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket update message bad request response has a 5xx status code
func (o *TicketUpdateMessageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket update message bad request response a status code equal to that given
func (o *TicketUpdateMessageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketUpdateMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageBadRequest  %+v", 400, o.Payload)
}

func (o *TicketUpdateMessageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageBadRequest  %+v", 400, o.Payload)
}

func (o *TicketUpdateMessageBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketUpdateMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketUpdateMessageUnauthorized creates a TicketUpdateMessageUnauthorized with default headers values
func NewTicketUpdateMessageUnauthorized() *TicketUpdateMessageUnauthorized {
	return &TicketUpdateMessageUnauthorized{}
}

/*
TicketUpdateMessageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketUpdateMessageUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket update message unauthorized response has a 2xx status code
func (o *TicketUpdateMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket update message unauthorized response has a 3xx status code
func (o *TicketUpdateMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message unauthorized response has a 4xx status code
func (o *TicketUpdateMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket update message unauthorized response has a 5xx status code
func (o *TicketUpdateMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket update message unauthorized response a status code equal to that given
func (o *TicketUpdateMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketUpdateMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketUpdateMessageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketUpdateMessageUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketUpdateMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketUpdateMessageForbidden creates a TicketUpdateMessageForbidden with default headers values
func NewTicketUpdateMessageForbidden() *TicketUpdateMessageForbidden {
	return &TicketUpdateMessageForbidden{}
}

/*
TicketUpdateMessageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketUpdateMessageForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket update message forbidden response has a 2xx status code
func (o *TicketUpdateMessageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket update message forbidden response has a 3xx status code
func (o *TicketUpdateMessageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message forbidden response has a 4xx status code
func (o *TicketUpdateMessageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket update message forbidden response has a 5xx status code
func (o *TicketUpdateMessageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket update message forbidden response a status code equal to that given
func (o *TicketUpdateMessageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketUpdateMessageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageForbidden  %+v", 403, o.Payload)
}

func (o *TicketUpdateMessageForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageForbidden  %+v", 403, o.Payload)
}

func (o *TicketUpdateMessageForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketUpdateMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketUpdateMessageNotFound creates a TicketUpdateMessageNotFound with default headers values
func NewTicketUpdateMessageNotFound() *TicketUpdateMessageNotFound {
	return &TicketUpdateMessageNotFound{}
}

/*
TicketUpdateMessageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketUpdateMessageNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket update message not found response has a 2xx status code
func (o *TicketUpdateMessageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket update message not found response has a 3xx status code
func (o *TicketUpdateMessageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message not found response has a 4xx status code
func (o *TicketUpdateMessageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket update message not found response has a 5xx status code
func (o *TicketUpdateMessageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket update message not found response a status code equal to that given
func (o *TicketUpdateMessageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketUpdateMessageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageNotFound  %+v", 404, o.Payload)
}

func (o *TicketUpdateMessageNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageNotFound  %+v", 404, o.Payload)
}

func (o *TicketUpdateMessageNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketUpdateMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketUpdateMessageInternalServerError creates a TicketUpdateMessageInternalServerError with default headers values
func NewTicketUpdateMessageInternalServerError() *TicketUpdateMessageInternalServerError {
	return &TicketUpdateMessageInternalServerError{}
}

/*
TicketUpdateMessageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketUpdateMessageInternalServerError struct {
}

// IsSuccess returns true when this ticket update message internal server error response has a 2xx status code
func (o *TicketUpdateMessageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket update message internal server error response has a 3xx status code
func (o *TicketUpdateMessageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket update message internal server error response has a 4xx status code
func (o *TicketUpdateMessageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket update message internal server error response has a 5xx status code
func (o *TicketUpdateMessageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket update message internal server error response a status code equal to that given
func (o *TicketUpdateMessageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketUpdateMessageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageInternalServerError ", 500)
}

func (o *TicketUpdateMessageInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit/message][%d] ticketUpdateMessageInternalServerError ", 500)
}

func (o *TicketUpdateMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
