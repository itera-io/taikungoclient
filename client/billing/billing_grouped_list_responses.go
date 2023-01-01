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

// BillingGroupedListReader is a Reader for the BillingGroupedList structure.
type BillingGroupedListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingGroupedListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingGroupedListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBillingGroupedListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBillingGroupedListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBillingGroupedListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBillingGroupedListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBillingGroupedListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBillingGroupedListOK creates a BillingGroupedListOK with default headers values
func NewBillingGroupedListOK() *BillingGroupedListOK {
	return &BillingGroupedListOK{}
}

/*
BillingGroupedListOK describes a response with status code 200, with default header values.

Success
*/
type BillingGroupedListOK struct {
	Payload []*models.GroupedBillingInfo
}

// IsSuccess returns true when this billing grouped list o k response has a 2xx status code
func (o *BillingGroupedListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing grouped list o k response has a 3xx status code
func (o *BillingGroupedListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list o k response has a 4xx status code
func (o *BillingGroupedListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing grouped list o k response has a 5xx status code
func (o *BillingGroupedListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing grouped list o k response a status code equal to that given
func (o *BillingGroupedListOK) IsCode(code int) bool {
	return code == 200
}

func (o *BillingGroupedListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListOK  %+v", 200, o.Payload)
}

func (o *BillingGroupedListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListOK  %+v", 200, o.Payload)
}

func (o *BillingGroupedListOK) GetPayload() []*models.GroupedBillingInfo {
	return o.Payload
}

func (o *BillingGroupedListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingGroupedListBadRequest creates a BillingGroupedListBadRequest with default headers values
func NewBillingGroupedListBadRequest() *BillingGroupedListBadRequest {
	return &BillingGroupedListBadRequest{}
}

/*
BillingGroupedListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BillingGroupedListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing grouped list bad request response has a 2xx status code
func (o *BillingGroupedListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing grouped list bad request response has a 3xx status code
func (o *BillingGroupedListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list bad request response has a 4xx status code
func (o *BillingGroupedListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing grouped list bad request response has a 5xx status code
func (o *BillingGroupedListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this billing grouped list bad request response a status code equal to that given
func (o *BillingGroupedListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BillingGroupedListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListBadRequest  %+v", 400, o.Payload)
}

func (o *BillingGroupedListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListBadRequest  %+v", 400, o.Payload)
}

func (o *BillingGroupedListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingGroupedListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingGroupedListUnauthorized creates a BillingGroupedListUnauthorized with default headers values
func NewBillingGroupedListUnauthorized() *BillingGroupedListUnauthorized {
	return &BillingGroupedListUnauthorized{}
}

/*
BillingGroupedListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BillingGroupedListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing grouped list unauthorized response has a 2xx status code
func (o *BillingGroupedListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing grouped list unauthorized response has a 3xx status code
func (o *BillingGroupedListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list unauthorized response has a 4xx status code
func (o *BillingGroupedListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing grouped list unauthorized response has a 5xx status code
func (o *BillingGroupedListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this billing grouped list unauthorized response a status code equal to that given
func (o *BillingGroupedListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BillingGroupedListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListUnauthorized  %+v", 401, o.Payload)
}

func (o *BillingGroupedListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListUnauthorized  %+v", 401, o.Payload)
}

func (o *BillingGroupedListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingGroupedListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingGroupedListForbidden creates a BillingGroupedListForbidden with default headers values
func NewBillingGroupedListForbidden() *BillingGroupedListForbidden {
	return &BillingGroupedListForbidden{}
}

/*
BillingGroupedListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BillingGroupedListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing grouped list forbidden response has a 2xx status code
func (o *BillingGroupedListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing grouped list forbidden response has a 3xx status code
func (o *BillingGroupedListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list forbidden response has a 4xx status code
func (o *BillingGroupedListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing grouped list forbidden response has a 5xx status code
func (o *BillingGroupedListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this billing grouped list forbidden response a status code equal to that given
func (o *BillingGroupedListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BillingGroupedListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListForbidden  %+v", 403, o.Payload)
}

func (o *BillingGroupedListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListForbidden  %+v", 403, o.Payload)
}

func (o *BillingGroupedListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingGroupedListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingGroupedListNotFound creates a BillingGroupedListNotFound with default headers values
func NewBillingGroupedListNotFound() *BillingGroupedListNotFound {
	return &BillingGroupedListNotFound{}
}

/*
BillingGroupedListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BillingGroupedListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this billing grouped list not found response has a 2xx status code
func (o *BillingGroupedListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing grouped list not found response has a 3xx status code
func (o *BillingGroupedListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list not found response has a 4xx status code
func (o *BillingGroupedListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this billing grouped list not found response has a 5xx status code
func (o *BillingGroupedListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this billing grouped list not found response a status code equal to that given
func (o *BillingGroupedListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BillingGroupedListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListNotFound  %+v", 404, o.Payload)
}

func (o *BillingGroupedListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListNotFound  %+v", 404, o.Payload)
}

func (o *BillingGroupedListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingGroupedListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingGroupedListInternalServerError creates a BillingGroupedListInternalServerError with default headers values
func NewBillingGroupedListInternalServerError() *BillingGroupedListInternalServerError {
	return &BillingGroupedListInternalServerError{}
}

/*
BillingGroupedListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BillingGroupedListInternalServerError struct {
}

// IsSuccess returns true when this billing grouped list internal server error response has a 2xx status code
func (o *BillingGroupedListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this billing grouped list internal server error response has a 3xx status code
func (o *BillingGroupedListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing grouped list internal server error response has a 4xx status code
func (o *BillingGroupedListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing grouped list internal server error response has a 5xx status code
func (o *BillingGroupedListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this billing grouped list internal server error response a status code equal to that given
func (o *BillingGroupedListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BillingGroupedListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListInternalServerError ", 500)
}

func (o *BillingGroupedListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/grouped][%d] billingGroupedListInternalServerError ", 500)
}

func (o *BillingGroupedListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
