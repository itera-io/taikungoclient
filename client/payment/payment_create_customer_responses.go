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

// PaymentCreateCustomerReader is a Reader for the PaymentCreateCustomer structure.
type PaymentCreateCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentCreateCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentCreateCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentCreateCustomerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentCreateCustomerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentCreateCustomerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentCreateCustomerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentCreateCustomerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentCreateCustomerOK creates a PaymentCreateCustomerOK with default headers values
func NewPaymentCreateCustomerOK() *PaymentCreateCustomerOK {
	return &PaymentCreateCustomerOK{}
}

/*
PaymentCreateCustomerOK describes a response with status code 200, with default header values.

Success
*/
type PaymentCreateCustomerOK struct {
	Payload string
}

// IsSuccess returns true when this payment create customer o k response has a 2xx status code
func (o *PaymentCreateCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment create customer o k response has a 3xx status code
func (o *PaymentCreateCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer o k response has a 4xx status code
func (o *PaymentCreateCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment create customer o k response has a 5xx status code
func (o *PaymentCreateCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment create customer o k response a status code equal to that given
func (o *PaymentCreateCustomerOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentCreateCustomerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerOK  %+v", 200, o.Payload)
}

func (o *PaymentCreateCustomerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerOK  %+v", 200, o.Payload)
}

func (o *PaymentCreateCustomerOK) GetPayload() string {
	return o.Payload
}

func (o *PaymentCreateCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentCreateCustomerBadRequest creates a PaymentCreateCustomerBadRequest with default headers values
func NewPaymentCreateCustomerBadRequest() *PaymentCreateCustomerBadRequest {
	return &PaymentCreateCustomerBadRequest{}
}

/*
PaymentCreateCustomerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentCreateCustomerBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this payment create customer bad request response has a 2xx status code
func (o *PaymentCreateCustomerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment create customer bad request response has a 3xx status code
func (o *PaymentCreateCustomerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer bad request response has a 4xx status code
func (o *PaymentCreateCustomerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment create customer bad request response has a 5xx status code
func (o *PaymentCreateCustomerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment create customer bad request response a status code equal to that given
func (o *PaymentCreateCustomerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentCreateCustomerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentCreateCustomerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentCreateCustomerBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *PaymentCreateCustomerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentCreateCustomerUnauthorized creates a PaymentCreateCustomerUnauthorized with default headers values
func NewPaymentCreateCustomerUnauthorized() *PaymentCreateCustomerUnauthorized {
	return &PaymentCreateCustomerUnauthorized{}
}

/*
PaymentCreateCustomerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentCreateCustomerUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment create customer unauthorized response has a 2xx status code
func (o *PaymentCreateCustomerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment create customer unauthorized response has a 3xx status code
func (o *PaymentCreateCustomerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer unauthorized response has a 4xx status code
func (o *PaymentCreateCustomerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment create customer unauthorized response has a 5xx status code
func (o *PaymentCreateCustomerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment create customer unauthorized response a status code equal to that given
func (o *PaymentCreateCustomerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentCreateCustomerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentCreateCustomerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentCreateCustomerUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentCreateCustomerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentCreateCustomerForbidden creates a PaymentCreateCustomerForbidden with default headers values
func NewPaymentCreateCustomerForbidden() *PaymentCreateCustomerForbidden {
	return &PaymentCreateCustomerForbidden{}
}

/*
PaymentCreateCustomerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentCreateCustomerForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment create customer forbidden response has a 2xx status code
func (o *PaymentCreateCustomerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment create customer forbidden response has a 3xx status code
func (o *PaymentCreateCustomerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer forbidden response has a 4xx status code
func (o *PaymentCreateCustomerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment create customer forbidden response has a 5xx status code
func (o *PaymentCreateCustomerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment create customer forbidden response a status code equal to that given
func (o *PaymentCreateCustomerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentCreateCustomerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerForbidden  %+v", 403, o.Payload)
}

func (o *PaymentCreateCustomerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerForbidden  %+v", 403, o.Payload)
}

func (o *PaymentCreateCustomerForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentCreateCustomerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentCreateCustomerNotFound creates a PaymentCreateCustomerNotFound with default headers values
func NewPaymentCreateCustomerNotFound() *PaymentCreateCustomerNotFound {
	return &PaymentCreateCustomerNotFound{}
}

/*
PaymentCreateCustomerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentCreateCustomerNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment create customer not found response has a 2xx status code
func (o *PaymentCreateCustomerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment create customer not found response has a 3xx status code
func (o *PaymentCreateCustomerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer not found response has a 4xx status code
func (o *PaymentCreateCustomerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment create customer not found response has a 5xx status code
func (o *PaymentCreateCustomerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment create customer not found response a status code equal to that given
func (o *PaymentCreateCustomerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentCreateCustomerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerNotFound  %+v", 404, o.Payload)
}

func (o *PaymentCreateCustomerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerNotFound  %+v", 404, o.Payload)
}

func (o *PaymentCreateCustomerNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentCreateCustomerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentCreateCustomerInternalServerError creates a PaymentCreateCustomerInternalServerError with default headers values
func NewPaymentCreateCustomerInternalServerError() *PaymentCreateCustomerInternalServerError {
	return &PaymentCreateCustomerInternalServerError{}
}

/*
PaymentCreateCustomerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentCreateCustomerInternalServerError struct {
}

// IsSuccess returns true when this payment create customer internal server error response has a 2xx status code
func (o *PaymentCreateCustomerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment create customer internal server error response has a 3xx status code
func (o *PaymentCreateCustomerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment create customer internal server error response has a 4xx status code
func (o *PaymentCreateCustomerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment create customer internal server error response has a 5xx status code
func (o *PaymentCreateCustomerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment create customer internal server error response a status code equal to that given
func (o *PaymentCreateCustomerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentCreateCustomerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerInternalServerError ", 500)
}

func (o *PaymentCreateCustomerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/createcustomer][%d] paymentCreateCustomerInternalServerError ", 500)
}

func (o *PaymentCreateCustomerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
