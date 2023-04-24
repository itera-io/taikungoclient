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

// PaymentGetBillingInfoReader is a Reader for the PaymentGetBillingInfo structure.
type PaymentGetBillingInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGetBillingInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGetBillingInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPaymentGetBillingInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPaymentGetBillingInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPaymentGetBillingInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPaymentGetBillingInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPaymentGetBillingInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPaymentGetBillingInfoOK creates a PaymentGetBillingInfoOK with default headers values
func NewPaymentGetBillingInfoOK() *PaymentGetBillingInfoOK {
	return &PaymentGetBillingInfoOK{}
}

/*
PaymentGetBillingInfoOK describes a response with status code 200, with default header values.

Success
*/
type PaymentGetBillingInfoOK struct {
	Payload *models.BillingInfoDto
}

// IsSuccess returns true when this payment get billing info o k response has a 2xx status code
func (o *PaymentGetBillingInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment get billing info o k response has a 3xx status code
func (o *PaymentGetBillingInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info o k response has a 4xx status code
func (o *PaymentGetBillingInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get billing info o k response has a 5xx status code
func (o *PaymentGetBillingInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get billing info o k response a status code equal to that given
func (o *PaymentGetBillingInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment get billing info o k response
func (o *PaymentGetBillingInfoOK) Code() int {
	return 200
}

func (o *PaymentGetBillingInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoOK  %+v", 200, o.Payload)
}

func (o *PaymentGetBillingInfoOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoOK  %+v", 200, o.Payload)
}

func (o *PaymentGetBillingInfoOK) GetPayload() *models.BillingInfoDto {
	return o.Payload
}

func (o *PaymentGetBillingInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingInfoDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetBillingInfoBadRequest creates a PaymentGetBillingInfoBadRequest with default headers values
func NewPaymentGetBillingInfoBadRequest() *PaymentGetBillingInfoBadRequest {
	return &PaymentGetBillingInfoBadRequest{}
}

/*
PaymentGetBillingInfoBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PaymentGetBillingInfoBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get billing info bad request response has a 2xx status code
func (o *PaymentGetBillingInfoBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get billing info bad request response has a 3xx status code
func (o *PaymentGetBillingInfoBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info bad request response has a 4xx status code
func (o *PaymentGetBillingInfoBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get billing info bad request response has a 5xx status code
func (o *PaymentGetBillingInfoBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get billing info bad request response a status code equal to that given
func (o *PaymentGetBillingInfoBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the payment get billing info bad request response
func (o *PaymentGetBillingInfoBadRequest) Code() int {
	return 400
}

func (o *PaymentGetBillingInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetBillingInfoBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoBadRequest  %+v", 400, o.Payload)
}

func (o *PaymentGetBillingInfoBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetBillingInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetBillingInfoUnauthorized creates a PaymentGetBillingInfoUnauthorized with default headers values
func NewPaymentGetBillingInfoUnauthorized() *PaymentGetBillingInfoUnauthorized {
	return &PaymentGetBillingInfoUnauthorized{}
}

/*
PaymentGetBillingInfoUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PaymentGetBillingInfoUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get billing info unauthorized response has a 2xx status code
func (o *PaymentGetBillingInfoUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get billing info unauthorized response has a 3xx status code
func (o *PaymentGetBillingInfoUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info unauthorized response has a 4xx status code
func (o *PaymentGetBillingInfoUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get billing info unauthorized response has a 5xx status code
func (o *PaymentGetBillingInfoUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get billing info unauthorized response a status code equal to that given
func (o *PaymentGetBillingInfoUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the payment get billing info unauthorized response
func (o *PaymentGetBillingInfoUnauthorized) Code() int {
	return 401
}

func (o *PaymentGetBillingInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetBillingInfoUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PaymentGetBillingInfoUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetBillingInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetBillingInfoForbidden creates a PaymentGetBillingInfoForbidden with default headers values
func NewPaymentGetBillingInfoForbidden() *PaymentGetBillingInfoForbidden {
	return &PaymentGetBillingInfoForbidden{}
}

/*
PaymentGetBillingInfoForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PaymentGetBillingInfoForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get billing info forbidden response has a 2xx status code
func (o *PaymentGetBillingInfoForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get billing info forbidden response has a 3xx status code
func (o *PaymentGetBillingInfoForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info forbidden response has a 4xx status code
func (o *PaymentGetBillingInfoForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get billing info forbidden response has a 5xx status code
func (o *PaymentGetBillingInfoForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get billing info forbidden response a status code equal to that given
func (o *PaymentGetBillingInfoForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the payment get billing info forbidden response
func (o *PaymentGetBillingInfoForbidden) Code() int {
	return 403
}

func (o *PaymentGetBillingInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetBillingInfoForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoForbidden  %+v", 403, o.Payload)
}

func (o *PaymentGetBillingInfoForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetBillingInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetBillingInfoNotFound creates a PaymentGetBillingInfoNotFound with default headers values
func NewPaymentGetBillingInfoNotFound() *PaymentGetBillingInfoNotFound {
	return &PaymentGetBillingInfoNotFound{}
}

/*
PaymentGetBillingInfoNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PaymentGetBillingInfoNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this payment get billing info not found response has a 2xx status code
func (o *PaymentGetBillingInfoNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get billing info not found response has a 3xx status code
func (o *PaymentGetBillingInfoNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info not found response has a 4xx status code
func (o *PaymentGetBillingInfoNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this payment get billing info not found response has a 5xx status code
func (o *PaymentGetBillingInfoNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this payment get billing info not found response a status code equal to that given
func (o *PaymentGetBillingInfoNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the payment get billing info not found response
func (o *PaymentGetBillingInfoNotFound) Code() int {
	return 404
}

func (o *PaymentGetBillingInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetBillingInfoNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoNotFound  %+v", 404, o.Payload)
}

func (o *PaymentGetBillingInfoNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PaymentGetBillingInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGetBillingInfoInternalServerError creates a PaymentGetBillingInfoInternalServerError with default headers values
func NewPaymentGetBillingInfoInternalServerError() *PaymentGetBillingInfoInternalServerError {
	return &PaymentGetBillingInfoInternalServerError{}
}

/*
PaymentGetBillingInfoInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PaymentGetBillingInfoInternalServerError struct {
}

// IsSuccess returns true when this payment get billing info internal server error response has a 2xx status code
func (o *PaymentGetBillingInfoInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this payment get billing info internal server error response has a 3xx status code
func (o *PaymentGetBillingInfoInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment get billing info internal server error response has a 4xx status code
func (o *PaymentGetBillingInfoInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment get billing info internal server error response has a 5xx status code
func (o *PaymentGetBillingInfoInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this payment get billing info internal server error response a status code equal to that given
func (o *PaymentGetBillingInfoInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the payment get billing info internal server error response
func (o *PaymentGetBillingInfoInternalServerError) Code() int {
	return 500
}

func (o *PaymentGetBillingInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoInternalServerError ", 500)
}

func (o *PaymentGetBillingInfoInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Payment/billing-info][%d] paymentGetBillingInfoInternalServerError ", 500)
}

func (o *PaymentGetBillingInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
