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

// PaymentGetCardInfoReader is a Reader for the PaymentGetCardInfo structure.
type PaymentGetCardInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetCardInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetCardInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentGetCardInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentGetCardInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentGetCardInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentGetCardInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentGetCardInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentGetCardInfoOK creates a PaymentGetCardInfoOK with default headers values
func NewPaymentGetCardInfoOK() *PaymentGetCardInfoOK {
	return &PaymentGetCardInfoOK{}
}

/*
PaymentGetCardInfoOK describes a response with status code 200, with default header values.

Success
*/
type PaymentGetCardInfoOK struct {
	Payload *models.CardInformationDto
}

// IsSuccess returns true when this payment get card info o k response has a 2xx status code
func (o *PaymentGetCardInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get card info o k response has a 3xx status code
func (o *PaymentGetCardInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info o k response has a 4xx status code
func (o *PaymentGetCardInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get card info o k response has a 5xx status code
func (o *PaymentGetCardInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get card info o k response a status code equal to that given
func (o *PaymentGetCardInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *PaymentGetCardInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoOK  %+v", 200, o.Payload)
}

func (o *PaymentGetCardInfoOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoOK  %+v", 200, o.Payload)
}

func (o *PaymentGetCardInfoOK) GetPayload() *models.CardInformationDto {
	return o.Payload
}

func (o *PaymentGetCardInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CardInformationDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetCardInfoBadRequest creates a PaymentGetCardInfoBadRequest with default headers values
func NewPaymentGetCardInfoBadRequest() *PaymentGetCardInfoBadRequest {
	return &PaymentGetCardInfoBadRequest{}
}

/*
PaymentGetCardInfoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentGetCardInfoBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment get card info bad request response has a 2xx status code
func (o *PaymentGetCardInfoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get card info bad request response has a 3xx status code
func (o *PaymentGetCardInfoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info bad request response has a 4xx status code
func (o *PaymentGetCardInfoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get card info bad request response has a 5xx status code
func (o *PaymentGetCardInfoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get card info bad request response a status code equal to that given
func (o *PaymentGetCardInfoBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PaymentGetCardInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetCardInfoBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetCardInfoBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentGetCardInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetCardInfoUnauthorized creates a PaymentGetCardInfoUnauthorized with default headers values
func NewPaymentGetCardInfoUnauthorized() *PaymentGetCardInfoUnauthorized {
	return &PaymentGetCardInfoUnauthorized{}
}

/*
PaymentGetCardInfoUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentGetCardInfoUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment get card info unauthorized response has a 2xx status code
func (o *PaymentGetCardInfoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get card info unauthorized response has a 3xx status code
func (o *PaymentGetCardInfoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info unauthorized response has a 4xx status code
func (o *PaymentGetCardInfoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get card info unauthorized response has a 5xx status code
func (o *PaymentGetCardInfoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get card info unauthorized response a status code equal to that given
func (o *PaymentGetCardInfoUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PaymentGetCardInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetCardInfoUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetCardInfoUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentGetCardInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetCardInfoForbidden creates a PaymentGetCardInfoForbidden with default headers values
func NewPaymentGetCardInfoForbidden() *PaymentGetCardInfoForbidden {
	return &PaymentGetCardInfoForbidden{}
}

/*
PaymentGetCardInfoForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentGetCardInfoForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment get card info forbidden response has a 2xx status code
func (o *PaymentGetCardInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get card info forbidden response has a 3xx status code
func (o *PaymentGetCardInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info forbidden response has a 4xx status code
func (o *PaymentGetCardInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get card info forbidden response has a 5xx status code
func (o *PaymentGetCardInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get card info forbidden response a status code equal to that given
func (o *PaymentGetCardInfoForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PaymentGetCardInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetCardInfoForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetCardInfoForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentGetCardInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetCardInfoNotFound creates a PaymentGetCardInfoNotFound with default headers values
func NewPaymentGetCardInfoNotFound() *PaymentGetCardInfoNotFound {
	return &PaymentGetCardInfoNotFound{}
}

/*
PaymentGetCardInfoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentGetCardInfoNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this payment get card info not found response has a 2xx status code
func (o *PaymentGetCardInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get card info not found response has a 3xx status code
func (o *PaymentGetCardInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info not found response has a 4xx status code
func (o *PaymentGetCardInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get card info not found response has a 5xx status code
func (o *PaymentGetCardInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get card info not found response a status code equal to that given
func (o *PaymentGetCardInfoNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PaymentGetCardInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetCardInfoNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetCardInfoNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PaymentGetCardInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetCardInfoInternalServerError creates a PaymentGetCardInfoInternalServerError with default headers values
func NewPaymentGetCardInfoInternalServerError() *PaymentGetCardInfoInternalServerError {
	return &PaymentGetCardInfoInternalServerError{}
}

/*
PaymentGetCardInfoInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentGetCardInfoInternalServerError struct {
}

// IsSuccess returns true when this payment get card info internal server error response has a 2xx status code
func (o *PaymentGetCardInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get card info internal server error response has a 3xx status code
func (o *PaymentGetCardInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get card info internal server error response has a 4xx status code
func (o *PaymentGetCardInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get card info internal server error response has a 5xx status code
func (o *PaymentGetCardInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment get card info internal server error response a status code equal to that given
func (o *PaymentGetCardInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PaymentGetCardInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoInternalServerError ", 500)
}

func (o *PaymentGetCardInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/cardinfo][%d] paymentGetCardInfoInternalServerError ", 500)
}

func (o *PaymentGetCardInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
