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

// PaymentGetFinalPriceReader is a Reader for the PaymentGetFinalPrice structure.
type PaymentGetFinalPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetFinalPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetFinalPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentGetFinalPriceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentGetFinalPriceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentGetFinalPriceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentGetFinalPriceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentGetFinalPriceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentGetFinalPriceOK creates a PaymentGetFinalPriceOK with default headers values
func NewPaymentGetFinalPriceOK() *PaymentGetFinalPriceOK {
	return &PaymentGetFinalPriceOK{}
}

/*
PaymentGetFinalPriceOK describes a response with status code 200, with default header values.

Success
*/
type PaymentGetFinalPriceOK struct {
	Payload *models.FinalPriceDto
}

// IsSuccess returns true when this payment get final price o k response has a 2xx status code
func (o *PaymentGetFinalPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get final price o k response has a 3xx status code
func (o *PaymentGetFinalPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price o k response has a 4xx status code
func (o *PaymentGetFinalPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get final price o k response has a 5xx status code
func (o *PaymentGetFinalPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price o k response a status code equal to that given
func (o *PaymentGetFinalPriceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentGetFinalPriceOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceOK  %+v", 200, o.Payload)
}

func (o *PaymentGetFinalPriceOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceOK  %+v", 200, o.Payload)
}

func (o *PaymentGetFinalPriceOK) GetPayload() *models.FinalPriceDto {
	return o.Payload
}

func (o *PaymentGetFinalPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FinalPriceDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceBadRequest creates a PaymentGetFinalPriceBadRequest with default headers values
func NewPaymentGetFinalPriceBadRequest() *PaymentGetFinalPriceBadRequest {
	return &PaymentGetFinalPriceBadRequest{}
}

/*
PaymentGetFinalPriceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentGetFinalPriceBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this payment get final price bad request response has a 2xx status code
func (o *PaymentGetFinalPriceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price bad request response has a 3xx status code
func (o *PaymentGetFinalPriceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price bad request response has a 4xx status code
func (o *PaymentGetFinalPriceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price bad request response has a 5xx status code
func (o *PaymentGetFinalPriceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price bad request response a status code equal to that given
func (o *PaymentGetFinalPriceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentGetFinalPriceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetFinalPriceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetFinalPriceBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *PaymentGetFinalPriceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceUnauthorized creates a PaymentGetFinalPriceUnauthorized with default headers values
func NewPaymentGetFinalPriceUnauthorized() *PaymentGetFinalPriceUnauthorized {
	return &PaymentGetFinalPriceUnauthorized{}
}

/*
PaymentGetFinalPriceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentGetFinalPriceUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get final price unauthorized response has a 2xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price unauthorized response has a 3xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price unauthorized response has a 4xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price unauthorized response has a 5xx status code
func (o *PaymentGetFinalPriceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price unauthorized response a status code equal to that given
func (o *PaymentGetFinalPriceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentGetFinalPriceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetFinalPriceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetFinalPriceUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetFinalPriceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceForbidden creates a PaymentGetFinalPriceForbidden with default headers values
func NewPaymentGetFinalPriceForbidden() *PaymentGetFinalPriceForbidden {
	return &PaymentGetFinalPriceForbidden{}
}

/*
PaymentGetFinalPriceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentGetFinalPriceForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get final price forbidden response has a 2xx status code
func (o *PaymentGetFinalPriceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price forbidden response has a 3xx status code
func (o *PaymentGetFinalPriceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price forbidden response has a 4xx status code
func (o *PaymentGetFinalPriceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price forbidden response has a 5xx status code
func (o *PaymentGetFinalPriceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price forbidden response a status code equal to that given
func (o *PaymentGetFinalPriceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentGetFinalPriceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetFinalPriceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetFinalPriceForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetFinalPriceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceNotFound creates a PaymentGetFinalPriceNotFound with default headers values
func NewPaymentGetFinalPriceNotFound() *PaymentGetFinalPriceNotFound {
	return &PaymentGetFinalPriceNotFound{}
}

/*
PaymentGetFinalPriceNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentGetFinalPriceNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get final price not found response has a 2xx status code
func (o *PaymentGetFinalPriceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price not found response has a 3xx status code
func (o *PaymentGetFinalPriceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price not found response has a 4xx status code
func (o *PaymentGetFinalPriceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get final price not found response has a 5xx status code
func (o *PaymentGetFinalPriceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get final price not found response a status code equal to that given
func (o *PaymentGetFinalPriceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentGetFinalPriceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetFinalPriceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetFinalPriceNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetFinalPriceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetFinalPriceInternalServerError creates a PaymentGetFinalPriceInternalServerError with default headers values
func NewPaymentGetFinalPriceInternalServerError() *PaymentGetFinalPriceInternalServerError {
	return &PaymentGetFinalPriceInternalServerError{}
}

/*
PaymentGetFinalPriceInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentGetFinalPriceInternalServerError struct {
}

// IsSuccess returns true when this payment get final price internal server error response has a 2xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get final price internal server error response has a 3xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get final price internal server error response has a 4xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get final price internal server error response has a 5xx status code
func (o *PaymentGetFinalPriceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment get final price internal server error response a status code equal to that given
func (o *PaymentGetFinalPriceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentGetFinalPriceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceInternalServerError ", 500)
}

func (o *PaymentGetFinalPriceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Payment/finalprice][%d] paymentGetFinalPriceInternalServerError ", 500)
}

func (o *PaymentGetFinalPriceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
