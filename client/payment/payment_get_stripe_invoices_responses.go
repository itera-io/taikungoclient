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

// PaymentGetStripeInvoicesReader is a Reader for the PaymentGetStripeInvoices structure.
type PaymentGetStripeInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetStripeInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetStripeInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentGetStripeInvoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentGetStripeInvoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentGetStripeInvoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentGetStripeInvoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentGetStripeInvoicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentGetStripeInvoicesOK creates a PaymentGetStripeInvoicesOK with default headers values
func NewPaymentGetStripeInvoicesOK() *PaymentGetStripeInvoicesOK {
	return &PaymentGetStripeInvoicesOK{}
}

/*
PaymentGetStripeInvoicesOK describes a response with status code 200, with default header values.

Success
*/
type PaymentGetStripeInvoicesOK struct {
	Payload *models.StripeInvoices
}

// IsSuccess returns true when this payment get stripe invoices o k response has a 2xx status code
func (o *PaymentGetStripeInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get stripe invoices o k response has a 3xx status code
func (o *PaymentGetStripeInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices o k response has a 4xx status code
func (o *PaymentGetStripeInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get stripe invoices o k response has a 5xx status code
func (o *PaymentGetStripeInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get stripe invoices o k response a status code equal to that given
func (o *PaymentGetStripeInvoicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment get stripe invoices o k response
func (o *PaymentGetStripeInvoicesOK) Code() int {
	return 200
}

func (o *PaymentGetStripeInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesOK  %+v", 200, o.Payload)
}

func (o *PaymentGetStripeInvoicesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesOK  %+v", 200, o.Payload)
}

func (o *PaymentGetStripeInvoicesOK) GetPayload() *models.StripeInvoices {
	return o.Payload
}

func (o *PaymentGetStripeInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StripeInvoices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetStripeInvoicesBadRequest creates a PaymentGetStripeInvoicesBadRequest with default headers values
func NewPaymentGetStripeInvoicesBadRequest() *PaymentGetStripeInvoicesBadRequest {
	return &PaymentGetStripeInvoicesBadRequest{}
}

/*
PaymentGetStripeInvoicesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentGetStripeInvoicesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get stripe invoices bad request response has a 2xx status code
func (o *PaymentGetStripeInvoicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get stripe invoices bad request response has a 3xx status code
func (o *PaymentGetStripeInvoicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices bad request response has a 4xx status code
func (o *PaymentGetStripeInvoicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get stripe invoices bad request response has a 5xx status code
func (o *PaymentGetStripeInvoicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get stripe invoices bad request response a status code equal to that given
func (o *PaymentGetStripeInvoicesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the payment get stripe invoices bad request response
func (o *PaymentGetStripeInvoicesBadRequest) Code() int {
	return 400
}

func (o *PaymentGetStripeInvoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetStripeInvoicesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetStripeInvoicesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetStripeInvoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetStripeInvoicesUnauthorized creates a PaymentGetStripeInvoicesUnauthorized with default headers values
func NewPaymentGetStripeInvoicesUnauthorized() *PaymentGetStripeInvoicesUnauthorized {
	return &PaymentGetStripeInvoicesUnauthorized{}
}

/*
PaymentGetStripeInvoicesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentGetStripeInvoicesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get stripe invoices unauthorized response has a 2xx status code
func (o *PaymentGetStripeInvoicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get stripe invoices unauthorized response has a 3xx status code
func (o *PaymentGetStripeInvoicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices unauthorized response has a 4xx status code
func (o *PaymentGetStripeInvoicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get stripe invoices unauthorized response has a 5xx status code
func (o *PaymentGetStripeInvoicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get stripe invoices unauthorized response a status code equal to that given
func (o *PaymentGetStripeInvoicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the payment get stripe invoices unauthorized response
func (o *PaymentGetStripeInvoicesUnauthorized) Code() int {
	return 401
}

func (o *PaymentGetStripeInvoicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetStripeInvoicesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetStripeInvoicesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetStripeInvoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetStripeInvoicesForbidden creates a PaymentGetStripeInvoicesForbidden with default headers values
func NewPaymentGetStripeInvoicesForbidden() *PaymentGetStripeInvoicesForbidden {
	return &PaymentGetStripeInvoicesForbidden{}
}

/*
PaymentGetStripeInvoicesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentGetStripeInvoicesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get stripe invoices forbidden response has a 2xx status code
func (o *PaymentGetStripeInvoicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get stripe invoices forbidden response has a 3xx status code
func (o *PaymentGetStripeInvoicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices forbidden response has a 4xx status code
func (o *PaymentGetStripeInvoicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get stripe invoices forbidden response has a 5xx status code
func (o *PaymentGetStripeInvoicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get stripe invoices forbidden response a status code equal to that given
func (o *PaymentGetStripeInvoicesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the payment get stripe invoices forbidden response
func (o *PaymentGetStripeInvoicesForbidden) Code() int {
	return 403
}

func (o *PaymentGetStripeInvoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetStripeInvoicesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetStripeInvoicesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetStripeInvoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetStripeInvoicesNotFound creates a PaymentGetStripeInvoicesNotFound with default headers values
func NewPaymentGetStripeInvoicesNotFound() *PaymentGetStripeInvoicesNotFound {
	return &PaymentGetStripeInvoicesNotFound{}
}

/*
PaymentGetStripeInvoicesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentGetStripeInvoicesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get stripe invoices not found response has a 2xx status code
func (o *PaymentGetStripeInvoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get stripe invoices not found response has a 3xx status code
func (o *PaymentGetStripeInvoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices not found response has a 4xx status code
func (o *PaymentGetStripeInvoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get stripe invoices not found response has a 5xx status code
func (o *PaymentGetStripeInvoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get stripe invoices not found response a status code equal to that given
func (o *PaymentGetStripeInvoicesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the payment get stripe invoices not found response
func (o *PaymentGetStripeInvoicesNotFound) Code() int {
	return 404
}

func (o *PaymentGetStripeInvoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetStripeInvoicesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetStripeInvoicesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetStripeInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetStripeInvoicesInternalServerError creates a PaymentGetStripeInvoicesInternalServerError with default headers values
func NewPaymentGetStripeInvoicesInternalServerError() *PaymentGetStripeInvoicesInternalServerError {
	return &PaymentGetStripeInvoicesInternalServerError{}
}

/*
PaymentGetStripeInvoicesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentGetStripeInvoicesInternalServerError struct {
}

// IsSuccess returns true when this payment get stripe invoices internal server error response has a 2xx status code
func (o *PaymentGetStripeInvoicesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get stripe invoices internal server error response has a 3xx status code
func (o *PaymentGetStripeInvoicesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get stripe invoices internal server error response has a 4xx status code
func (o *PaymentGetStripeInvoicesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get stripe invoices internal server error response has a 5xx status code
func (o *PaymentGetStripeInvoicesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment get stripe invoices internal server error response a status code equal to that given
func (o *PaymentGetStripeInvoicesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the payment get stripe invoices internal server error response
func (o *PaymentGetStripeInvoicesInternalServerError) Code() int {
	return 500
}

func (o *PaymentGetStripeInvoicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesInternalServerError ", 500)
}

func (o *PaymentGetStripeInvoicesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/stripeinvoices/{subscriptionId}][%d] paymentGetStripeInvoicesInternalServerError ", 500)
}

func (o *PaymentGetStripeInvoicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
