// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PaymentPayInvoiceReader is a Reader for the PaymentPayInvoice structure.
type PaymentPayInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentPayInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentPayInvoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentPayInvoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentPayInvoiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentPayInvoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentPayInvoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentPayInvoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentPayInvoiceOK creates a PaymentPayInvoiceOK with default headers values
func NewPaymentPayInvoiceOK() *PaymentPayInvoiceOK {
	return &PaymentPayInvoiceOK{}
}

/*
PaymentPayInvoiceOK describes a response with status code 200, with default header values.

Success
*/
type PaymentPayInvoiceOK struct {
}

// IsSuccess returns true when this payment pay invoice o k response has a 2xx status code
func (o *PaymentPayInvoiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment pay invoice o k response has a 3xx status code
func (o *PaymentPayInvoiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice o k response has a 4xx status code
func (o *PaymentPayInvoiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment pay invoice o k response has a 5xx status code
func (o *PaymentPayInvoiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment pay invoice o k response a status code equal to that given
func (o *PaymentPayInvoiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentPayInvoiceOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceOK ", 200)
}

func (o *PaymentPayInvoiceOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceOK ", 200)
}

func (o *PaymentPayInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPaymentPayInvoiceBadRequest creates a PaymentPayInvoiceBadRequest with default headers values
func NewPaymentPayInvoiceBadRequest() *PaymentPayInvoiceBadRequest {
	return &PaymentPayInvoiceBadRequest{}
}

/*
PaymentPayInvoiceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentPayInvoiceBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this payment pay invoice bad request response has a 2xx status code
func (o *PaymentPayInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment pay invoice bad request response has a 3xx status code
func (o *PaymentPayInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice bad request response has a 4xx status code
func (o *PaymentPayInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment pay invoice bad request response has a 5xx status code
func (o *PaymentPayInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment pay invoice bad request response a status code equal to that given
func (o *PaymentPayInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentPayInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentPayInvoiceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentPayInvoiceBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PaymentPayInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentPayInvoiceUnauthorized creates a PaymentPayInvoiceUnauthorized with default headers values
func NewPaymentPayInvoiceUnauthorized() *PaymentPayInvoiceUnauthorized {
	return &PaymentPayInvoiceUnauthorized{}
}

/*
PaymentPayInvoiceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentPayInvoiceUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this payment pay invoice unauthorized response has a 2xx status code
func (o *PaymentPayInvoiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment pay invoice unauthorized response has a 3xx status code
func (o *PaymentPayInvoiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice unauthorized response has a 4xx status code
func (o *PaymentPayInvoiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment pay invoice unauthorized response has a 5xx status code
func (o *PaymentPayInvoiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment pay invoice unauthorized response a status code equal to that given
func (o *PaymentPayInvoiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentPayInvoiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentPayInvoiceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentPayInvoiceUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PaymentPayInvoiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentPayInvoiceForbidden creates a PaymentPayInvoiceForbidden with default headers values
func NewPaymentPayInvoiceForbidden() *PaymentPayInvoiceForbidden {
	return &PaymentPayInvoiceForbidden{}
}

/*
PaymentPayInvoiceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentPayInvoiceForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this payment pay invoice forbidden response has a 2xx status code
func (o *PaymentPayInvoiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment pay invoice forbidden response has a 3xx status code
func (o *PaymentPayInvoiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice forbidden response has a 4xx status code
func (o *PaymentPayInvoiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment pay invoice forbidden response has a 5xx status code
func (o *PaymentPayInvoiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment pay invoice forbidden response a status code equal to that given
func (o *PaymentPayInvoiceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentPayInvoiceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentPayInvoiceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentPayInvoiceForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PaymentPayInvoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentPayInvoiceNotFound creates a PaymentPayInvoiceNotFound with default headers values
func NewPaymentPayInvoiceNotFound() *PaymentPayInvoiceNotFound {
	return &PaymentPayInvoiceNotFound{}
}

/*
PaymentPayInvoiceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentPayInvoiceNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this payment pay invoice not found response has a 2xx status code
func (o *PaymentPayInvoiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment pay invoice not found response has a 3xx status code
func (o *PaymentPayInvoiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice not found response has a 4xx status code
func (o *PaymentPayInvoiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment pay invoice not found response has a 5xx status code
func (o *PaymentPayInvoiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment pay invoice not found response a status code equal to that given
func (o *PaymentPayInvoiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentPayInvoiceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentPayInvoiceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentPayInvoiceNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PaymentPayInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentPayInvoiceInternalServerError creates a PaymentPayInvoiceInternalServerError with default headers values
func NewPaymentPayInvoiceInternalServerError() *PaymentPayInvoiceInternalServerError {
	return &PaymentPayInvoiceInternalServerError{}
}

/*
PaymentPayInvoiceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentPayInvoiceInternalServerError struct {
}

// IsSuccess returns true when this payment pay invoice internal server error response has a 2xx status code
func (o *PaymentPayInvoiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment pay invoice internal server error response has a 3xx status code
func (o *PaymentPayInvoiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment pay invoice internal server error response has a 4xx status code
func (o *PaymentPayInvoiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment pay invoice internal server error response has a 5xx status code
func (o *PaymentPayInvoiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment pay invoice internal server error response a status code equal to that given
func (o *PaymentPayInvoiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentPayInvoiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceInternalServerError ", 500)
}

func (o *PaymentPayInvoiceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/pay][%d] paymentPayInvoiceInternalServerError ", 500)
}

func (o *PaymentPayInvoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
