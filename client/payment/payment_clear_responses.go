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

// PaymentClearReader is a Reader for the PaymentClear structure.
type PaymentClearReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentClearReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentClearOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentClearBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentClearUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentClearForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentClearNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentClearInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentClearOK creates a PaymentClearOK with default headers values
func NewPaymentClearOK() *PaymentClearOK {
	return &PaymentClearOK{}
}

/*
PaymentClearOK describes a response with status code 200, with default header values.

Success
*/
type PaymentClearOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this payment clear o k response has a 2xx status code
func (o *PaymentClearOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment clear o k response has a 3xx status code
func (o *PaymentClearOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear o k response has a 4xx status code
func (o *PaymentClearOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment clear o k response has a 5xx status code
func (o *PaymentClearOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment clear o k response a status code equal to that given
func (o *PaymentClearOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentClearOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearOK  %+v", 200, o.Payload)
}

func (o *PaymentClearOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearOK  %+v", 200, o.Payload)
}

func (o *PaymentClearOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PaymentClearOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentClearBadRequest creates a PaymentClearBadRequest with default headers values
func NewPaymentClearBadRequest() *PaymentClearBadRequest {
	return &PaymentClearBadRequest{}
}

/*
PaymentClearBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentClearBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this payment clear bad request response has a 2xx status code
func (o *PaymentClearBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment clear bad request response has a 3xx status code
func (o *PaymentClearBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear bad request response has a 4xx status code
func (o *PaymentClearBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment clear bad request response has a 5xx status code
func (o *PaymentClearBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment clear bad request response a status code equal to that given
func (o *PaymentClearBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentClearBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentClearBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentClearBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PaymentClearBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentClearUnauthorized creates a PaymentClearUnauthorized with default headers values
func NewPaymentClearUnauthorized() *PaymentClearUnauthorized {
	return &PaymentClearUnauthorized{}
}

/*
PaymentClearUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentClearUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment clear unauthorized response has a 2xx status code
func (o *PaymentClearUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment clear unauthorized response has a 3xx status code
func (o *PaymentClearUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear unauthorized response has a 4xx status code
func (o *PaymentClearUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment clear unauthorized response has a 5xx status code
func (o *PaymentClearUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment clear unauthorized response a status code equal to that given
func (o *PaymentClearUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentClearUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentClearUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentClearUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentClearUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentClearForbidden creates a PaymentClearForbidden with default headers values
func NewPaymentClearForbidden() *PaymentClearForbidden {
	return &PaymentClearForbidden{}
}

/*
PaymentClearForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentClearForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment clear forbidden response has a 2xx status code
func (o *PaymentClearForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment clear forbidden response has a 3xx status code
func (o *PaymentClearForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear forbidden response has a 4xx status code
func (o *PaymentClearForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment clear forbidden response has a 5xx status code
func (o *PaymentClearForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment clear forbidden response a status code equal to that given
func (o *PaymentClearForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentClearForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearForbidden  %+v", 403, o.Payload)
}

func (o *PaymentClearForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearForbidden  %+v", 403, o.Payload)
}

func (o *PaymentClearForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentClearForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentClearNotFound creates a PaymentClearNotFound with default headers values
func NewPaymentClearNotFound() *PaymentClearNotFound {
	return &PaymentClearNotFound{}
}

/*
PaymentClearNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentClearNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment clear not found response has a 2xx status code
func (o *PaymentClearNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment clear not found response has a 3xx status code
func (o *PaymentClearNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear not found response has a 4xx status code
func (o *PaymentClearNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment clear not found response has a 5xx status code
func (o *PaymentClearNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment clear not found response a status code equal to that given
func (o *PaymentClearNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentClearNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearNotFound  %+v", 404, o.Payload)
}

func (o *PaymentClearNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearNotFound  %+v", 404, o.Payload)
}

func (o *PaymentClearNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentClearNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentClearInternalServerError creates a PaymentClearInternalServerError with default headers values
func NewPaymentClearInternalServerError() *PaymentClearInternalServerError {
	return &PaymentClearInternalServerError{}
}

/*
PaymentClearInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentClearInternalServerError struct {
}

// IsSuccess returns true when this payment clear internal server error response has a 2xx status code
func (o *PaymentClearInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment clear internal server error response has a 3xx status code
func (o *PaymentClearInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment clear internal server error response has a 4xx status code
func (o *PaymentClearInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment clear internal server error response has a 5xx status code
func (o *PaymentClearInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment clear internal server error response a status code equal to that given
func (o *PaymentClearInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentClearInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearInternalServerError ", 500)
}

func (o *PaymentClearInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/clear][%d] paymentClearInternalServerError ", 500)
}

func (o *PaymentClearInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
