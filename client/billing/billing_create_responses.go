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

// BillingCreateReader is a Reader for the BillingCreate structure.
type BillingCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBillingCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBillingCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBillingCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBillingCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBillingCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBillingCreateOK creates a BillingCreateOK with default headers values
func NewBillingCreateOK() *BillingCreateOK {
	return &BillingCreateOK{}
}

/* BillingCreateOK describes a response with status code 200, with default header values.

Success
*/
type BillingCreateOK struct {
	Payload models.Unit
}

func (o *BillingCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateOK  %+v", 200, o.Payload)
}
func (o *BillingCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BillingCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCreateBadRequest creates a BillingCreateBadRequest with default headers values
func NewBillingCreateBadRequest() *BillingCreateBadRequest {
	return &BillingCreateBadRequest{}
}

/* BillingCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BillingCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *BillingCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateBadRequest  %+v", 400, o.Payload)
}
func (o *BillingCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BillingCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCreateUnauthorized creates a BillingCreateUnauthorized with default headers values
func NewBillingCreateUnauthorized() *BillingCreateUnauthorized {
	return &BillingCreateUnauthorized{}
}

/* BillingCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BillingCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *BillingCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *BillingCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCreateForbidden creates a BillingCreateForbidden with default headers values
func NewBillingCreateForbidden() *BillingCreateForbidden {
	return &BillingCreateForbidden{}
}

/* BillingCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BillingCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *BillingCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateForbidden  %+v", 403, o.Payload)
}
func (o *BillingCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCreateNotFound creates a BillingCreateNotFound with default headers values
func NewBillingCreateNotFound() *BillingCreateNotFound {
	return &BillingCreateNotFound{}
}

/* BillingCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BillingCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *BillingCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateNotFound  %+v", 404, o.Payload)
}
func (o *BillingCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingCreateInternalServerError creates a BillingCreateInternalServerError with default headers values
func NewBillingCreateInternalServerError() *BillingCreateInternalServerError {
	return &BillingCreateInternalServerError{}
}

/* BillingCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BillingCreateInternalServerError struct {
}

func (o *BillingCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Billing/add][%d] billingCreateInternalServerError ", 500)
}

func (o *BillingCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}