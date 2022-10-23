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

// InvoicesInvoicesReader is a Reader for the InvoicesInvoices structure.
type InvoicesInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInvoicesInvoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInvoicesInvoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInvoicesInvoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvoicesInvoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInvoicesInvoicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInvoicesInvoicesOK creates a InvoicesInvoicesOK with default headers values
func NewInvoicesInvoicesOK() *InvoicesInvoicesOK {
	return &InvoicesInvoicesOK{}
}

/*
InvoicesInvoicesOK describes a response with status code 200, with default header values.

Success
*/
type InvoicesInvoicesOK struct {
	Payload *models.Invoices
}

// IsSuccess returns true when this invoices invoices o k response has a 2xx status code
func (o *InvoicesInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invoices invoices o k response has a 3xx status code
func (o *InvoicesInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices o k response has a 4xx status code
func (o *InvoicesInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices invoices o k response has a 5xx status code
func (o *InvoicesInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices invoices o k response a status code equal to that given
func (o *InvoicesInvoicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *InvoicesInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesOK  %+v", 200, o.Payload)
}

func (o *InvoicesInvoicesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesOK  %+v", 200, o.Payload)
}

func (o *InvoicesInvoicesOK) GetPayload() *models.Invoices {
	return o.Payload
}

func (o *InvoicesInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invoices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesInvoicesBadRequest creates a InvoicesInvoicesBadRequest with default headers values
func NewInvoicesInvoicesBadRequest() *InvoicesInvoicesBadRequest {
	return &InvoicesInvoicesBadRequest{}
}

/*
InvoicesInvoicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InvoicesInvoicesBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this invoices invoices bad request response has a 2xx status code
func (o *InvoicesInvoicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices invoices bad request response has a 3xx status code
func (o *InvoicesInvoicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices bad request response has a 4xx status code
func (o *InvoicesInvoicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices invoices bad request response has a 5xx status code
func (o *InvoicesInvoicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices invoices bad request response a status code equal to that given
func (o *InvoicesInvoicesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InvoicesInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesInvoicesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesInvoicesBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *InvoicesInvoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesInvoicesUnauthorized creates a InvoicesInvoicesUnauthorized with default headers values
func NewInvoicesInvoicesUnauthorized() *InvoicesInvoicesUnauthorized {
	return &InvoicesInvoicesUnauthorized{}
}

/*
InvoicesInvoicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InvoicesInvoicesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices invoices unauthorized response has a 2xx status code
func (o *InvoicesInvoicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices invoices unauthorized response has a 3xx status code
func (o *InvoicesInvoicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices unauthorized response has a 4xx status code
func (o *InvoicesInvoicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices invoices unauthorized response has a 5xx status code
func (o *InvoicesInvoicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices invoices unauthorized response a status code equal to that given
func (o *InvoicesInvoicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *InvoicesInvoicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesInvoicesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesInvoicesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesInvoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesInvoicesForbidden creates a InvoicesInvoicesForbidden with default headers values
func NewInvoicesInvoicesForbidden() *InvoicesInvoicesForbidden {
	return &InvoicesInvoicesForbidden{}
}

/*
InvoicesInvoicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InvoicesInvoicesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices invoices forbidden response has a 2xx status code
func (o *InvoicesInvoicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices invoices forbidden response has a 3xx status code
func (o *InvoicesInvoicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices forbidden response has a 4xx status code
func (o *InvoicesInvoicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices invoices forbidden response has a 5xx status code
func (o *InvoicesInvoicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices invoices forbidden response a status code equal to that given
func (o *InvoicesInvoicesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InvoicesInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesInvoicesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesInvoicesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesInvoicesNotFound creates a InvoicesInvoicesNotFound with default headers values
func NewInvoicesInvoicesNotFound() *InvoicesInvoicesNotFound {
	return &InvoicesInvoicesNotFound{}
}

/*
InvoicesInvoicesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InvoicesInvoicesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices invoices not found response has a 2xx status code
func (o *InvoicesInvoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices invoices not found response has a 3xx status code
func (o *InvoicesInvoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices not found response has a 4xx status code
func (o *InvoicesInvoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices invoices not found response has a 5xx status code
func (o *InvoicesInvoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices invoices not found response a status code equal to that given
func (o *InvoicesInvoicesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *InvoicesInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesInvoicesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesInvoicesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesInvoicesInternalServerError creates a InvoicesInvoicesInternalServerError with default headers values
func NewInvoicesInvoicesInternalServerError() *InvoicesInvoicesInternalServerError {
	return &InvoicesInvoicesInternalServerError{}
}

/*
InvoicesInvoicesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type InvoicesInvoicesInternalServerError struct {
}

// IsSuccess returns true when this invoices invoices internal server error response has a 2xx status code
func (o *InvoicesInvoicesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices invoices internal server error response has a 3xx status code
func (o *InvoicesInvoicesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices invoices internal server error response has a 4xx status code
func (o *InvoicesInvoicesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices invoices internal server error response has a 5xx status code
func (o *InvoicesInvoicesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this invoices invoices internal server error response a status code equal to that given
func (o *InvoicesInvoicesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *InvoicesInvoicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesInternalServerError ", 500)
}

func (o *InvoicesInvoicesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Invoices/list][%d] invoicesInvoicesInternalServerError ", 500)
}

func (o *InvoicesInvoicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
