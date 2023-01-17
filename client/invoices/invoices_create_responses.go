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

// InvoicesCreateReader is a Reader for the InvoicesCreate structure.
type InvoicesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInvoicesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInvoicesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInvoicesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvoicesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInvoicesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInvoicesCreateOK creates a InvoicesCreateOK with default headers values
func NewInvoicesCreateOK() *InvoicesCreateOK {
	return &InvoicesCreateOK{}
}

/*
InvoicesCreateOK describes a response with status code 200, with default header values.

Success
*/
type InvoicesCreateOK struct {
	Payload int32
}

// IsSuccess returns true when this invoices create o k response has a 2xx status code
func (o *InvoicesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invoices create o k response has a 3xx status code
func (o *InvoicesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create o k response has a 4xx status code
func (o *InvoicesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices create o k response has a 5xx status code
func (o *InvoicesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices create o k response a status code equal to that given
func (o *InvoicesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the invoices create o k response
func (o *InvoicesCreateOK) Code() int {
	return 200
}

func (o *InvoicesCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateOK  %+v", 200, o.Payload)
}

func (o *InvoicesCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateOK  %+v", 200, o.Payload)
}

func (o *InvoicesCreateOK) GetPayload() int32 {
	return o.Payload
}

func (o *InvoicesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesCreateBadRequest creates a InvoicesCreateBadRequest with default headers values
func NewInvoicesCreateBadRequest() *InvoicesCreateBadRequest {
	return &InvoicesCreateBadRequest{}
}

/*
InvoicesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InvoicesCreateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices create bad request response has a 2xx status code
func (o *InvoicesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices create bad request response has a 3xx status code
func (o *InvoicesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create bad request response has a 4xx status code
func (o *InvoicesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices create bad request response has a 5xx status code
func (o *InvoicesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices create bad request response a status code equal to that given
func (o *InvoicesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the invoices create bad request response
func (o *InvoicesCreateBadRequest) Code() int {
	return 400
}

func (o *InvoicesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesCreateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesCreateUnauthorized creates a InvoicesCreateUnauthorized with default headers values
func NewInvoicesCreateUnauthorized() *InvoicesCreateUnauthorized {
	return &InvoicesCreateUnauthorized{}
}

/*
InvoicesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InvoicesCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices create unauthorized response has a 2xx status code
func (o *InvoicesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices create unauthorized response has a 3xx status code
func (o *InvoicesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create unauthorized response has a 4xx status code
func (o *InvoicesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices create unauthorized response has a 5xx status code
func (o *InvoicesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices create unauthorized response a status code equal to that given
func (o *InvoicesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the invoices create unauthorized response
func (o *InvoicesCreateUnauthorized) Code() int {
	return 401
}

func (o *InvoicesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesCreateForbidden creates a InvoicesCreateForbidden with default headers values
func NewInvoicesCreateForbidden() *InvoicesCreateForbidden {
	return &InvoicesCreateForbidden{}
}

/*
InvoicesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InvoicesCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices create forbidden response has a 2xx status code
func (o *InvoicesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices create forbidden response has a 3xx status code
func (o *InvoicesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create forbidden response has a 4xx status code
func (o *InvoicesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices create forbidden response has a 5xx status code
func (o *InvoicesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices create forbidden response a status code equal to that given
func (o *InvoicesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the invoices create forbidden response
func (o *InvoicesCreateForbidden) Code() int {
	return 403
}

func (o *InvoicesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesCreateNotFound creates a InvoicesCreateNotFound with default headers values
func NewInvoicesCreateNotFound() *InvoicesCreateNotFound {
	return &InvoicesCreateNotFound{}
}

/*
InvoicesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InvoicesCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this invoices create not found response has a 2xx status code
func (o *InvoicesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices create not found response has a 3xx status code
func (o *InvoicesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create not found response has a 4xx status code
func (o *InvoicesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices create not found response has a 5xx status code
func (o *InvoicesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices create not found response a status code equal to that given
func (o *InvoicesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the invoices create not found response
func (o *InvoicesCreateNotFound) Code() int {
	return 404
}

func (o *InvoicesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *InvoicesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesCreateInternalServerError creates a InvoicesCreateInternalServerError with default headers values
func NewInvoicesCreateInternalServerError() *InvoicesCreateInternalServerError {
	return &InvoicesCreateInternalServerError{}
}

/*
InvoicesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type InvoicesCreateInternalServerError struct {
}

// IsSuccess returns true when this invoices create internal server error response has a 2xx status code
func (o *InvoicesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices create internal server error response has a 3xx status code
func (o *InvoicesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices create internal server error response has a 4xx status code
func (o *InvoicesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices create internal server error response has a 5xx status code
func (o *InvoicesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this invoices create internal server error response a status code equal to that given
func (o *InvoicesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the invoices create internal server error response
func (o *InvoicesCreateInternalServerError) Code() int {
	return 500
}

func (o *InvoicesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateInternalServerError ", 500)
}

func (o *InvoicesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Invoices/create][%d] invoicesCreateInternalServerError ", 500)
}

func (o *InvoicesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
