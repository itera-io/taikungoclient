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

// PaymentUpdateCardReader is a Reader for the PaymentUpdateCard structure.
type PaymentUpdateCardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentUpdateCardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentUpdateCardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentUpdateCardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentUpdateCardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentUpdateCardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentUpdateCardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentUpdateCardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentUpdateCardOK creates a PaymentUpdateCardOK with default headers values
func NewPaymentUpdateCardOK() *PaymentUpdateCardOK {
	return &PaymentUpdateCardOK{}
}

/*
PaymentUpdateCardOK describes a response with status code 200, with default header values.

Success
*/
type PaymentUpdateCardOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this payment update card o k response has a 2xx status code
func (o *PaymentUpdateCardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment update card o k response has a 3xx status code
func (o *PaymentUpdateCardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card o k response has a 4xx status code
func (o *PaymentUpdateCardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment update card o k response has a 5xx status code
func (o *PaymentUpdateCardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment update card o k response a status code equal to that given
func (o *PaymentUpdateCardOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentUpdateCardOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardOK  %+v", 200, o.Payload)
}

func (o *PaymentUpdateCardOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardOK  %+v", 200, o.Payload)
}

func (o *PaymentUpdateCardOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PaymentUpdateCardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentUpdateCardBadRequest creates a PaymentUpdateCardBadRequest with default headers values
func NewPaymentUpdateCardBadRequest() *PaymentUpdateCardBadRequest {
	return &PaymentUpdateCardBadRequest{}
}

/*
PaymentUpdateCardBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentUpdateCardBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this payment update card bad request response has a 2xx status code
func (o *PaymentUpdateCardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment update card bad request response has a 3xx status code
func (o *PaymentUpdateCardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card bad request response has a 4xx status code
func (o *PaymentUpdateCardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment update card bad request response has a 5xx status code
func (o *PaymentUpdateCardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment update card bad request response a status code equal to that given
func (o *PaymentUpdateCardBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentUpdateCardBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentUpdateCardBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentUpdateCardBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *PaymentUpdateCardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentUpdateCardUnauthorized creates a PaymentUpdateCardUnauthorized with default headers values
func NewPaymentUpdateCardUnauthorized() *PaymentUpdateCardUnauthorized {
	return &PaymentUpdateCardUnauthorized{}
}

/*
PaymentUpdateCardUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentUpdateCardUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment update card unauthorized response has a 2xx status code
func (o *PaymentUpdateCardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment update card unauthorized response has a 3xx status code
func (o *PaymentUpdateCardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card unauthorized response has a 4xx status code
func (o *PaymentUpdateCardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment update card unauthorized response has a 5xx status code
func (o *PaymentUpdateCardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment update card unauthorized response a status code equal to that given
func (o *PaymentUpdateCardUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentUpdateCardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentUpdateCardUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentUpdateCardUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentUpdateCardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentUpdateCardForbidden creates a PaymentUpdateCardForbidden with default headers values
func NewPaymentUpdateCardForbidden() *PaymentUpdateCardForbidden {
	return &PaymentUpdateCardForbidden{}
}

/*
PaymentUpdateCardForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentUpdateCardForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment update card forbidden response has a 2xx status code
func (o *PaymentUpdateCardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment update card forbidden response has a 3xx status code
func (o *PaymentUpdateCardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card forbidden response has a 4xx status code
func (o *PaymentUpdateCardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment update card forbidden response has a 5xx status code
func (o *PaymentUpdateCardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment update card forbidden response a status code equal to that given
func (o *PaymentUpdateCardForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentUpdateCardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardForbidden  %+v", 403, o.Payload)
}

func (o *PaymentUpdateCardForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardForbidden  %+v", 403, o.Payload)
}

func (o *PaymentUpdateCardForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentUpdateCardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentUpdateCardNotFound creates a PaymentUpdateCardNotFound with default headers values
func NewPaymentUpdateCardNotFound() *PaymentUpdateCardNotFound {
	return &PaymentUpdateCardNotFound{}
}

/*
PaymentUpdateCardNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentUpdateCardNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment update card not found response has a 2xx status code
func (o *PaymentUpdateCardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment update card not found response has a 3xx status code
func (o *PaymentUpdateCardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card not found response has a 4xx status code
func (o *PaymentUpdateCardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment update card not found response has a 5xx status code
func (o *PaymentUpdateCardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment update card not found response a status code equal to that given
func (o *PaymentUpdateCardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentUpdateCardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardNotFound  %+v", 404, o.Payload)
}

func (o *PaymentUpdateCardNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardNotFound  %+v", 404, o.Payload)
}

func (o *PaymentUpdateCardNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentUpdateCardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentUpdateCardInternalServerError creates a PaymentUpdateCardInternalServerError with default headers values
func NewPaymentUpdateCardInternalServerError() *PaymentUpdateCardInternalServerError {
	return &PaymentUpdateCardInternalServerError{}
}

/*
PaymentUpdateCardInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentUpdateCardInternalServerError struct {
}

// IsSuccess returns true when this payment update card internal server error response has a 2xx status code
func (o *PaymentUpdateCardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment update card internal server error response has a 3xx status code
func (o *PaymentUpdateCardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment update card internal server error response has a 4xx status code
func (o *PaymentUpdateCardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment update card internal server error response has a 5xx status code
func (o *PaymentUpdateCardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment update card internal server error response a status code equal to that given
func (o *PaymentUpdateCardInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentUpdateCardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardInternalServerError ", 500)
}

func (o *PaymentUpdateCardInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/updatecard][%d] paymentUpdateCardInternalServerError ", 500)
}

func (o *PaymentUpdateCardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
