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

// TicketTransferListReader is a Reader for the TicketTransferList structure.
type TicketTransferListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketTransferListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketTransferListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketTransferListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketTransferListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketTransferListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketTransferListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketTransferListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketTransferListOK creates a TicketTransferListOK with default headers values
func NewTicketTransferListOK() *TicketTransferListOK {
	return &TicketTransferListOK{}
}

/*
TicketTransferListOK describes a response with status code 200, with default header values.

Success
*/
type TicketTransferListOK struct {
	Payload []*models.TransferList
}

// IsSuccess returns true when this ticket transfer list o k response has a 2xx status code
func (o *TicketTransferListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ticket transfer list o k response has a 3xx status code
func (o *TicketTransferListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list o k response has a 4xx status code
func (o *TicketTransferListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket transfer list o k response has a 5xx status code
func (o *TicketTransferListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer list o k response a status code equal to that given
func (o *TicketTransferListOK) IsCode(code int) bool {
	return code == 200
}

func (o *TicketTransferListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListOK  %+v", 200, o.Payload)
}

func (o *TicketTransferListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListOK  %+v", 200, o.Payload)
}

func (o *TicketTransferListOK) GetPayload() []*models.TransferList {
	return o.Payload
}

func (o *TicketTransferListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferListBadRequest creates a TicketTransferListBadRequest with default headers values
func NewTicketTransferListBadRequest() *TicketTransferListBadRequest {
	return &TicketTransferListBadRequest{}
}

/*
TicketTransferListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketTransferListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ticket transfer list bad request response has a 2xx status code
func (o *TicketTransferListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer list bad request response has a 3xx status code
func (o *TicketTransferListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list bad request response has a 4xx status code
func (o *TicketTransferListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer list bad request response has a 5xx status code
func (o *TicketTransferListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer list bad request response a status code equal to that given
func (o *TicketTransferListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TicketTransferListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTransferListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListBadRequest  %+v", 400, o.Payload)
}

func (o *TicketTransferListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *TicketTransferListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferListUnauthorized creates a TicketTransferListUnauthorized with default headers values
func NewTicketTransferListUnauthorized() *TicketTransferListUnauthorized {
	return &TicketTransferListUnauthorized{}
}

/*
TicketTransferListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketTransferListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ticket transfer list unauthorized response has a 2xx status code
func (o *TicketTransferListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer list unauthorized response has a 3xx status code
func (o *TicketTransferListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list unauthorized response has a 4xx status code
func (o *TicketTransferListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer list unauthorized response has a 5xx status code
func (o *TicketTransferListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer list unauthorized response a status code equal to that given
func (o *TicketTransferListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *TicketTransferListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTransferListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListUnauthorized  %+v", 401, o.Payload)
}

func (o *TicketTransferListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *TicketTransferListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferListForbidden creates a TicketTransferListForbidden with default headers values
func NewTicketTransferListForbidden() *TicketTransferListForbidden {
	return &TicketTransferListForbidden{}
}

/*
TicketTransferListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketTransferListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ticket transfer list forbidden response has a 2xx status code
func (o *TicketTransferListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer list forbidden response has a 3xx status code
func (o *TicketTransferListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list forbidden response has a 4xx status code
func (o *TicketTransferListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer list forbidden response has a 5xx status code
func (o *TicketTransferListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer list forbidden response a status code equal to that given
func (o *TicketTransferListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *TicketTransferListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListForbidden  %+v", 403, o.Payload)
}

func (o *TicketTransferListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListForbidden  %+v", 403, o.Payload)
}

func (o *TicketTransferListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *TicketTransferListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferListNotFound creates a TicketTransferListNotFound with default headers values
func NewTicketTransferListNotFound() *TicketTransferListNotFound {
	return &TicketTransferListNotFound{}
}

/*
TicketTransferListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketTransferListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this ticket transfer list not found response has a 2xx status code
func (o *TicketTransferListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer list not found response has a 3xx status code
func (o *TicketTransferListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list not found response has a 4xx status code
func (o *TicketTransferListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ticket transfer list not found response has a 5xx status code
func (o *TicketTransferListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ticket transfer list not found response a status code equal to that given
func (o *TicketTransferListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *TicketTransferListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListNotFound  %+v", 404, o.Payload)
}

func (o *TicketTransferListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListNotFound  %+v", 404, o.Payload)
}

func (o *TicketTransferListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *TicketTransferListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketTransferListInternalServerError creates a TicketTransferListInternalServerError with default headers values
func NewTicketTransferListInternalServerError() *TicketTransferListInternalServerError {
	return &TicketTransferListInternalServerError{}
}

/*
TicketTransferListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketTransferListInternalServerError struct {
}

// IsSuccess returns true when this ticket transfer list internal server error response has a 2xx status code
func (o *TicketTransferListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ticket transfer list internal server error response has a 3xx status code
func (o *TicketTransferListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ticket transfer list internal server error response has a 4xx status code
func (o *TicketTransferListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ticket transfer list internal server error response has a 5xx status code
func (o *TicketTransferListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ticket transfer list internal server error response a status code equal to that given
func (o *TicketTransferListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *TicketTransferListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListInternalServerError ", 500)
}

func (o *TicketTransferListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/transfer/list][%d] ticketTransferListInternalServerError ", 500)
}

func (o *TicketTransferListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
