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

// TicketEditReader is a Reader for the TicketEdit structure.
type TicketEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketEditOK creates a TicketEditOK with default headers values
func NewTicketEditOK() *TicketEditOK {
	return &TicketEditOK{}
}

/*
TicketEditOK describes a response with status code 200, with default header values.

Success
*/
type TicketEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this ticket edit o k response has a 2xx status code
func (o *TicketEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket edit o k response has a 3xx status code
func (o *TicketEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit o k response has a 4xx status code
func (o *TicketEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket edit o k response has a 5xx status code
func (o *TicketEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket edit o k response a status code equal to that given
func (o *TicketEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditOK  %+v", 200, o.Payload)
}

func (o *TicketEditOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditOK  %+v", 200, o.Payload)
}

func (o *TicketEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *TicketEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketEditBadRequest creates a TicketEditBadRequest with default headers values
func NewTicketEditBadRequest() *TicketEditBadRequest {
	return &TicketEditBadRequest{}
}

/*
TicketEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketEditBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ticket edit bad request response has a 2xx status code
func (o *TicketEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket edit bad request response has a 3xx status code
func (o *TicketEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit bad request response has a 4xx status code
func (o *TicketEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket edit bad request response has a 5xx status code
func (o *TicketEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket edit bad request response a status code equal to that given
func (o *TicketEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditBadRequest  %+v", 400, o.Payload)
}

func (o *TicketEditBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditBadRequest  %+v", 400, o.Payload)
}

func (o *TicketEditBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketEditUnauthorized creates a TicketEditUnauthorized with default headers values
func NewTicketEditUnauthorized() *TicketEditUnauthorized {
	return &TicketEditUnauthorized{}
}

/*
TicketEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketEditUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket edit unauthorized response has a 2xx status code
func (o *TicketEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket edit unauthorized response has a 3xx status code
func (o *TicketEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit unauthorized response has a 4xx status code
func (o *TicketEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket edit unauthorized response has a 5xx status code
func (o *TicketEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket edit unauthorized response a status code equal to that given
func (o *TicketEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketEditUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketEditForbidden creates a TicketEditForbidden with default headers values
func NewTicketEditForbidden() *TicketEditForbidden {
	return &TicketEditForbidden{}
}

/*
TicketEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketEditForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket edit forbidden response has a 2xx status code
func (o *TicketEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket edit forbidden response has a 3xx status code
func (o *TicketEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit forbidden response has a 4xx status code
func (o *TicketEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket edit forbidden response has a 5xx status code
func (o *TicketEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket edit forbidden response a status code equal to that given
func (o *TicketEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditForbidden  %+v", 403, o.Payload)
}

func (o *TicketEditForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditForbidden  %+v", 403, o.Payload)
}

func (o *TicketEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketEditNotFound creates a TicketEditNotFound with default headers values
func NewTicketEditNotFound() *TicketEditNotFound {
	return &TicketEditNotFound{}
}

/*
TicketEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketEditNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ticket edit not found response has a 2xx status code
func (o *TicketEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket edit not found response has a 3xx status code
func (o *TicketEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit not found response has a 4xx status code
func (o *TicketEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket edit not found response has a 5xx status code
func (o *TicketEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket edit not found response a status code equal to that given
func (o *TicketEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditNotFound  %+v", 404, o.Payload)
}

func (o *TicketEditNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditNotFound  %+v", 404, o.Payload)
}

func (o *TicketEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketEditInternalServerError creates a TicketEditInternalServerError with default headers values
func NewTicketEditInternalServerError() *TicketEditInternalServerError {
	return &TicketEditInternalServerError{}
}

/*
TicketEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketEditInternalServerError struct {
}

// IsSuccess returns true when this ticket edit internal server error response has a 2xx status code
func (o *TicketEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket edit internal server error response has a 3xx status code
func (o *TicketEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket edit internal server error response has a 4xx status code
func (o *TicketEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket edit internal server error response has a 5xx status code
func (o *TicketEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket edit internal server error response a status code equal to that given
func (o *TicketEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditInternalServerError ", 500)
}

func (o *TicketEditInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Ticket/edit][%d] ticketEditInternalServerError ", 500)
}

func (o *TicketEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
