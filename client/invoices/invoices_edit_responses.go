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

// InvoicesEditReader is a Reader for the InvoicesEdit structure.
type InvoicesEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoicesEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoicesEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInvoicesEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInvoicesEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInvoicesEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInvoicesEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInvoicesEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInvoicesEditOK creates a InvoicesEditOK with default headers values
func NewInvoicesEditOK() *InvoicesEditOK {
	return &InvoicesEditOK{}
}

/*
InvoicesEditOK describes a response with status code 200, with default header values.

Success
*/
type InvoicesEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this invoices edit o k response has a 2xx status code
func (o *InvoicesEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invoices edit o k response has a 3xx status code
func (o *InvoicesEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit o k response has a 4xx status code
func (o *InvoicesEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices edit o k response has a 5xx status code
func (o *InvoicesEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices edit o k response a status code equal to that given
func (o *InvoicesEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *InvoicesEditOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditOK  %+v", 200, o.Payload)
}

func (o *InvoicesEditOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditOK  %+v", 200, o.Payload)
}

func (o *InvoicesEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *InvoicesEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesEditBadRequest creates a InvoicesEditBadRequest with default headers values
func NewInvoicesEditBadRequest() *InvoicesEditBadRequest {
	return &InvoicesEditBadRequest{}
}

/*
InvoicesEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InvoicesEditBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this invoices edit bad request response has a 2xx status code
func (o *InvoicesEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices edit bad request response has a 3xx status code
func (o *InvoicesEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit bad request response has a 4xx status code
func (o *InvoicesEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices edit bad request response has a 5xx status code
func (o *InvoicesEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices edit bad request response a status code equal to that given
func (o *InvoicesEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *InvoicesEditBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesEditBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditBadRequest  %+v", 400, o.Payload)
}

func (o *InvoicesEditBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *InvoicesEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesEditUnauthorized creates a InvoicesEditUnauthorized with default headers values
func NewInvoicesEditUnauthorized() *InvoicesEditUnauthorized {
	return &InvoicesEditUnauthorized{}
}

/*
InvoicesEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InvoicesEditUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this invoices edit unauthorized response has a 2xx status code
func (o *InvoicesEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices edit unauthorized response has a 3xx status code
func (o *InvoicesEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit unauthorized response has a 4xx status code
func (o *InvoicesEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices edit unauthorized response has a 5xx status code
func (o *InvoicesEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices edit unauthorized response a status code equal to that given
func (o *InvoicesEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *InvoicesEditUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesEditUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditUnauthorized  %+v", 401, o.Payload)
}

func (o *InvoicesEditUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *InvoicesEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesEditForbidden creates a InvoicesEditForbidden with default headers values
func NewInvoicesEditForbidden() *InvoicesEditForbidden {
	return &InvoicesEditForbidden{}
}

/*
InvoicesEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InvoicesEditForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this invoices edit forbidden response has a 2xx status code
func (o *InvoicesEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices edit forbidden response has a 3xx status code
func (o *InvoicesEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit forbidden response has a 4xx status code
func (o *InvoicesEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices edit forbidden response has a 5xx status code
func (o *InvoicesEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices edit forbidden response a status code equal to that given
func (o *InvoicesEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *InvoicesEditForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesEditForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditForbidden  %+v", 403, o.Payload)
}

func (o *InvoicesEditForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *InvoicesEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesEditNotFound creates a InvoicesEditNotFound with default headers values
func NewInvoicesEditNotFound() *InvoicesEditNotFound {
	return &InvoicesEditNotFound{}
}

/*
InvoicesEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InvoicesEditNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this invoices edit not found response has a 2xx status code
func (o *InvoicesEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices edit not found response has a 3xx status code
func (o *InvoicesEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit not found response has a 4xx status code
func (o *InvoicesEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this invoices edit not found response has a 5xx status code
func (o *InvoicesEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this invoices edit not found response a status code equal to that given
func (o *InvoicesEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *InvoicesEditNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesEditNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditNotFound  %+v", 404, o.Payload)
}

func (o *InvoicesEditNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *InvoicesEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoicesEditInternalServerError creates a InvoicesEditInternalServerError with default headers values
func NewInvoicesEditInternalServerError() *InvoicesEditInternalServerError {
	return &InvoicesEditInternalServerError{}
}

/*
InvoicesEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type InvoicesEditInternalServerError struct {
}

// IsSuccess returns true when this invoices edit internal server error response has a 2xx status code
func (o *InvoicesEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this invoices edit internal server error response has a 3xx status code
func (o *InvoicesEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoices edit internal server error response has a 4xx status code
func (o *InvoicesEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoices edit internal server error response has a 5xx status code
func (o *InvoicesEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this invoices edit internal server error response a status code equal to that given
func (o *InvoicesEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *InvoicesEditInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditInternalServerError ", 500)
}

func (o *InvoicesEditInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/Invoices/update/{invoiceId}][%d] invoicesEditInternalServerError ", 500)
}

func (o *InvoicesEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
