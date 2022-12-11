// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// BillingListReader is a Reader for the BillingList structure.
type BillingListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBillingListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBillingListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBillingListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBillingListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBillingListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBillingListOK creates a BillingListOK with default headers values
func NewBillingListOK() *BillingListOK {
	return &BillingListOK{}
}

/*
BillingListOK describes a response with status code 200, with default header values.

Success
*/
type BillingListOK struct {
	Payload *models.BillingInfo
}

// IsSuccess returns true when this billing list o k response has a 2xx status code
func (o *BillingListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing list o k response has a 3xx status code
func (o *BillingListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list o k response has a 4xx status code
func (o *BillingListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing list o k response has a 5xx status code
func (o *BillingListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing list o k response a status code equal to that given
func (o *BillingListOK) IsCode(code int) bool {
	return code == 200
}

func (o *BillingListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListOK  %+v", 200, o.Payload)
}

func (o *BillingListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListOK  %+v", 200, o.Payload)
}

func (o *BillingListOK) GetPayload() *models.BillingInfo {
	return o.Payload
}

func (o *BillingListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BillingInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingListBadRequest creates a BillingListBadRequest with default headers values
func NewBillingListBadRequest() *BillingListBadRequest {
	return &BillingListBadRequest{}
}

/*
BillingListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BillingListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this billing list bad request response has a 2xx status code
func (o *BillingListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing list bad request response has a 3xx status code
func (o *BillingListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list bad request response has a 4xx status code
func (o *BillingListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing list bad request response has a 5xx status code
func (o *BillingListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this billing list bad request response a status code equal to that given
func (o *BillingListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BillingListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListBadRequest  %+v", 400, o.Payload)
}

func (o *BillingListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListBadRequest  %+v", 400, o.Payload)
}

func (o *BillingListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *BillingListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingListUnauthorized creates a BillingListUnauthorized with default headers values
func NewBillingListUnauthorized() *BillingListUnauthorized {
	return &BillingListUnauthorized{}
}

/*
BillingListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BillingListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing list unauthorized response has a 2xx status code
func (o *BillingListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing list unauthorized response has a 3xx status code
func (o *BillingListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list unauthorized response has a 4xx status code
func (o *BillingListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing list unauthorized response has a 5xx status code
func (o *BillingListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this billing list unauthorized response a status code equal to that given
func (o *BillingListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BillingListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListUnauthorized  %+v", 401, o.Payload)
}

func (o *BillingListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListUnauthorized  %+v", 401, o.Payload)
}

func (o *BillingListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingListForbidden creates a BillingListForbidden with default headers values
func NewBillingListForbidden() *BillingListForbidden {
	return &BillingListForbidden{}
}

/*
BillingListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BillingListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing list forbidden response has a 2xx status code
func (o *BillingListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing list forbidden response has a 3xx status code
func (o *BillingListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list forbidden response has a 4xx status code
func (o *BillingListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing list forbidden response has a 5xx status code
func (o *BillingListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this billing list forbidden response a status code equal to that given
func (o *BillingListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BillingListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListForbidden  %+v", 403, o.Payload)
}

func (o *BillingListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListForbidden  %+v", 403, o.Payload)
}

func (o *BillingListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingListNotFound creates a BillingListNotFound with default headers values
func NewBillingListNotFound() *BillingListNotFound {
	return &BillingListNotFound{}
}

/*
BillingListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BillingListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing list not found response has a 2xx status code
func (o *BillingListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing list not found response has a 3xx status code
func (o *BillingListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list not found response has a 4xx status code
func (o *BillingListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing list not found response has a 5xx status code
func (o *BillingListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this billing list not found response a status code equal to that given
func (o *BillingListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BillingListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListNotFound  %+v", 404, o.Payload)
}

func (o *BillingListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListNotFound  %+v", 404, o.Payload)
}

func (o *BillingListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingListInternalServerError creates a BillingListInternalServerError with default headers values
func NewBillingListInternalServerError() *BillingListInternalServerError {
	return &BillingListInternalServerError{}
}

/*
BillingListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BillingListInternalServerError struct {
}

// IsSuccess returns true when this billing list internal server error response has a 2xx status code
func (o *BillingListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing list internal server error response has a 3xx status code
func (o *BillingListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing list internal server error response has a 4xx status code
func (o *BillingListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing list internal server error response has a 5xx status code
func (o *BillingListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this billing list internal server error response a status code equal to that given
func (o *BillingListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BillingListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListInternalServerError ", 500)
}

func (o *BillingListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing][%d] billingListInternalServerError ", 500)
}

func (o *BillingListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
