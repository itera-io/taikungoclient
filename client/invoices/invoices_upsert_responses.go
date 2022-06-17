// Code generated by go-swagger; DO NOT EDIT.

package invoices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// InvoicesUpsertReader is a Reader for the InvoicesUpsert structure.
type InvoicesUpsertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesUpsertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesUpsertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInvoicesUpsertBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInvoicesUpsertUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInvoicesUpsertForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvoicesUpsertNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInvoicesUpsertInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInvoicesUpsertOK creates a InvoicesUpsertOK with default headers values
func NewInvoicesUpsertOK() *InvoicesUpsertOK {
	return &InvoicesUpsertOK{}
}

/* InvoicesUpsertOK describes a response with status code 200, with default header values.

Success
*/
type InvoicesUpsertOK struct {
	Payload int32
}

func (o *InvoicesUpsertOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertOK  %+v", 200, o.Payload)
}
func (o *InvoicesUpsertOK) GetPayload() int32 {
	return o.Payload
}

func (o *InvoicesUpsertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesUpsertBadRequest creates a InvoicesUpsertBadRequest with default headers values
func NewInvoicesUpsertBadRequest() *InvoicesUpsertBadRequest {
	return &InvoicesUpsertBadRequest{}
}

/* InvoicesUpsertBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InvoicesUpsertBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *InvoicesUpsertBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertBadRequest  %+v", 400, o.Payload)
}
func (o *InvoicesUpsertBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *InvoicesUpsertBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesUpsertUnauthorized creates a InvoicesUpsertUnauthorized with default headers values
func NewInvoicesUpsertUnauthorized() *InvoicesUpsertUnauthorized {
	return &InvoicesUpsertUnauthorized{}
}

/* InvoicesUpsertUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InvoicesUpsertUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *InvoicesUpsertUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertUnauthorized  %+v", 401, o.Payload)
}
func (o *InvoicesUpsertUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesUpsertUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesUpsertForbidden creates a InvoicesUpsertForbidden with default headers values
func NewInvoicesUpsertForbidden() *InvoicesUpsertForbidden {
	return &InvoicesUpsertForbidden{}
}

/* InvoicesUpsertForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InvoicesUpsertForbidden struct {
	Payload *models.ProblemDetails
}

func (o *InvoicesUpsertForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertForbidden  %+v", 403, o.Payload)
}
func (o *InvoicesUpsertForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesUpsertForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesUpsertNotFound creates a InvoicesUpsertNotFound with default headers values
func NewInvoicesUpsertNotFound() *InvoicesUpsertNotFound {
	return &InvoicesUpsertNotFound{}
}

/* InvoicesUpsertNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InvoicesUpsertNotFound struct {
	Payload *models.ProblemDetails
}

func (o *InvoicesUpsertNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertNotFound  %+v", 404, o.Payload)
}
func (o *InvoicesUpsertNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesUpsertNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesUpsertInternalServerError creates a InvoicesUpsertInternalServerError with default headers values
func NewInvoicesUpsertInternalServerError() *InvoicesUpsertInternalServerError {
	return &InvoicesUpsertInternalServerError{}
}

/* InvoicesUpsertInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type InvoicesUpsertInternalServerError struct {
}

func (o *InvoicesUpsertInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices][%d] invoicesUpsertInternalServerError ", 500)
}

func (o *InvoicesUpsertInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
