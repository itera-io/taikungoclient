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

// PaymentWebhookReader is a Reader for the PaymentWebhook structure.
type PaymentWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentWebhookUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentWebhookForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentWebhookInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentWebhookOK creates a PaymentWebhookOK with default headers values
func NewPaymentWebhookOK() *PaymentWebhookOK {
	return &PaymentWebhookOK{}
}

/*
PaymentWebhookOK describes a response with status code 200, with default header values.

Success
*/
type PaymentWebhookOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this payment webhook o k response has a 2xx status code
func (o *PaymentWebhookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment webhook o k response has a 3xx status code
func (o *PaymentWebhookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook o k response has a 4xx status code
func (o *PaymentWebhookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment webhook o k response has a 5xx status code
func (o *PaymentWebhookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment webhook o k response a status code equal to that given
func (o *PaymentWebhookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment webhook o k response
func (o *PaymentWebhookOK) Code() int {
	return 200
}

func (o *PaymentWebhookOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookOK  %+v", 200, o.Payload)
}

func (o *PaymentWebhookOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookOK  %+v", 200, o.Payload)
}

func (o *PaymentWebhookOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PaymentWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentWebhookBadRequest creates a PaymentWebhookBadRequest with default headers values
func NewPaymentWebhookBadRequest() *PaymentWebhookBadRequest {
	return &PaymentWebhookBadRequest{}
}

/*
PaymentWebhookBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentWebhookBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment webhook bad request response has a 2xx status code
func (o *PaymentWebhookBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment webhook bad request response has a 3xx status code
func (o *PaymentWebhookBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook bad request response has a 4xx status code
func (o *PaymentWebhookBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment webhook bad request response has a 5xx status code
func (o *PaymentWebhookBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment webhook bad request response a status code equal to that given
func (o *PaymentWebhookBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the payment webhook bad request response
func (o *PaymentWebhookBadRequest) Code() int {
	return 400
}

func (o *PaymentWebhookBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentWebhookBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentWebhookBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentWebhookUnauthorized creates a PaymentWebhookUnauthorized with default headers values
func NewPaymentWebhookUnauthorized() *PaymentWebhookUnauthorized {
	return &PaymentWebhookUnauthorized{}
}

/*
PaymentWebhookUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentWebhookUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment webhook unauthorized response has a 2xx status code
func (o *PaymentWebhookUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment webhook unauthorized response has a 3xx status code
func (o *PaymentWebhookUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook unauthorized response has a 4xx status code
func (o *PaymentWebhookUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment webhook unauthorized response has a 5xx status code
func (o *PaymentWebhookUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment webhook unauthorized response a status code equal to that given
func (o *PaymentWebhookUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the payment webhook unauthorized response
func (o *PaymentWebhookUnauthorized) Code() int {
	return 401
}

func (o *PaymentWebhookUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentWebhookUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentWebhookUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentWebhookUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentWebhookForbidden creates a PaymentWebhookForbidden with default headers values
func NewPaymentWebhookForbidden() *PaymentWebhookForbidden {
	return &PaymentWebhookForbidden{}
}

/*
PaymentWebhookForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentWebhookForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment webhook forbidden response has a 2xx status code
func (o *PaymentWebhookForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment webhook forbidden response has a 3xx status code
func (o *PaymentWebhookForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook forbidden response has a 4xx status code
func (o *PaymentWebhookForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment webhook forbidden response has a 5xx status code
func (o *PaymentWebhookForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment webhook forbidden response a status code equal to that given
func (o *PaymentWebhookForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the payment webhook forbidden response
func (o *PaymentWebhookForbidden) Code() int {
	return 403
}

func (o *PaymentWebhookForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookForbidden  %+v", 403, o.Payload)
}

func (o *PaymentWebhookForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookForbidden  %+v", 403, o.Payload)
}

func (o *PaymentWebhookForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentWebhookForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentWebhookNotFound creates a PaymentWebhookNotFound with default headers values
func NewPaymentWebhookNotFound() *PaymentWebhookNotFound {
	return &PaymentWebhookNotFound{}
}

/*
PaymentWebhookNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentWebhookNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment webhook not found response has a 2xx status code
func (o *PaymentWebhookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment webhook not found response has a 3xx status code
func (o *PaymentWebhookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook not found response has a 4xx status code
func (o *PaymentWebhookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment webhook not found response has a 5xx status code
func (o *PaymentWebhookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment webhook not found response a status code equal to that given
func (o *PaymentWebhookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the payment webhook not found response
func (o *PaymentWebhookNotFound) Code() int {
	return 404
}

func (o *PaymentWebhookNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookNotFound  %+v", 404, o.Payload)
}

func (o *PaymentWebhookNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookNotFound  %+v", 404, o.Payload)
}

func (o *PaymentWebhookNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentWebhookInternalServerError creates a PaymentWebhookInternalServerError with default headers values
func NewPaymentWebhookInternalServerError() *PaymentWebhookInternalServerError {
	return &PaymentWebhookInternalServerError{}
}

/*
PaymentWebhookInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentWebhookInternalServerError struct {
}

// IsSuccess returns true when this payment webhook internal server error response has a 2xx status code
func (o *PaymentWebhookInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment webhook internal server error response has a 3xx status code
func (o *PaymentWebhookInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment webhook internal server error response has a 4xx status code
func (o *PaymentWebhookInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment webhook internal server error response has a 5xx status code
func (o *PaymentWebhookInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment webhook internal server error response a status code equal to that given
func (o *PaymentWebhookInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the payment webhook internal server error response
func (o *PaymentWebhookInternalServerError) Code() int {
	return 500
}

func (o *PaymentWebhookInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookInternalServerError ", 500)
}

func (o *PaymentWebhookInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/webhook][%d] paymentWebhookInternalServerError ", 500)
}

func (o *PaymentWebhookInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
