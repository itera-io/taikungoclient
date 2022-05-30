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

// TicketArticleListReader is a Reader for the TicketArticleList structure.
type TicketArticleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketArticleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketArticleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTicketArticleListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTicketArticleListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTicketArticleListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTicketArticleListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTicketArticleListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTicketArticleListOK creates a TicketArticleListOK with default headers values
func NewTicketArticleListOK() *TicketArticleListOK {
	return &TicketArticleListOK{}
}

/* TicketArticleListOK describes a response with status code 200, with default header values.

Success
*/
type TicketArticleListOK struct {
	Payload *models.ArticleList
}

func (o *TicketArticleListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListOK  %+v", 200, o.Payload)
}
func (o *TicketArticleListOK) GetPayload() *models.ArticleList {
	return o.Payload
}

func (o *TicketArticleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArticleList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListBadRequest creates a TicketArticleListBadRequest with default headers values
func NewTicketArticleListBadRequest() *TicketArticleListBadRequest {
	return &TicketArticleListBadRequest{}
}

/* TicketArticleListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TicketArticleListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *TicketArticleListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListBadRequest  %+v", 400, o.Payload)
}
func (o *TicketArticleListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *TicketArticleListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListUnauthorized creates a TicketArticleListUnauthorized with default headers values
func NewTicketArticleListUnauthorized() *TicketArticleListUnauthorized {
	return &TicketArticleListUnauthorized{}
}

/* TicketArticleListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type TicketArticleListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *TicketArticleListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListUnauthorized  %+v", 401, o.Payload)
}
func (o *TicketArticleListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArticleListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListForbidden creates a TicketArticleListForbidden with default headers values
func NewTicketArticleListForbidden() *TicketArticleListForbidden {
	return &TicketArticleListForbidden{}
}

/* TicketArticleListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TicketArticleListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *TicketArticleListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListForbidden  %+v", 403, o.Payload)
}
func (o *TicketArticleListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArticleListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListNotFound creates a TicketArticleListNotFound with default headers values
func NewTicketArticleListNotFound() *TicketArticleListNotFound {
	return &TicketArticleListNotFound{}
}

/* TicketArticleListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type TicketArticleListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *TicketArticleListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListNotFound  %+v", 404, o.Payload)
}
func (o *TicketArticleListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *TicketArticleListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTicketArticleListInternalServerError creates a TicketArticleListInternalServerError with default headers values
func NewTicketArticleListInternalServerError() *TicketArticleListInternalServerError {
	return &TicketArticleListInternalServerError{}
}

/* TicketArticleListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type TicketArticleListInternalServerError struct {
}

func (o *TicketArticleListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Ticket/{ticketId}/messages][%d] ticketArticleListInternalServerError ", 500)
}

func (o *TicketArticleListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}