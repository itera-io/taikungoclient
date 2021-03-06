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

// BillingExportCsvReader is a Reader for the BillingExportCsv structure.
type BillingExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBillingExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBillingExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBillingExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBillingExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBillingExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBillingExportCsvOK creates a BillingExportCsvOK with default headers values
func NewBillingExportCsvOK() *BillingExportCsvOK {
	return &BillingExportCsvOK{}
}

/* BillingExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type BillingExportCsvOK struct {
}

func (o *BillingExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvOK ", 200)
}

func (o *BillingExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBillingExportCsvBadRequest creates a BillingExportCsvBadRequest with default headers values
func NewBillingExportCsvBadRequest() *BillingExportCsvBadRequest {
	return &BillingExportCsvBadRequest{}
}

/* BillingExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BillingExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *BillingExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvBadRequest  %+v", 400, o.Payload)
}
func (o *BillingExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BillingExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingExportCsvUnauthorized creates a BillingExportCsvUnauthorized with default headers values
func NewBillingExportCsvUnauthorized() *BillingExportCsvUnauthorized {
	return &BillingExportCsvUnauthorized{}
}

/* BillingExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BillingExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *BillingExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvUnauthorized  %+v", 401, o.Payload)
}
func (o *BillingExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingExportCsvForbidden creates a BillingExportCsvForbidden with default headers values
func NewBillingExportCsvForbidden() *BillingExportCsvForbidden {
	return &BillingExportCsvForbidden{}
}

/* BillingExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BillingExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

func (o *BillingExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvForbidden  %+v", 403, o.Payload)
}
func (o *BillingExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingExportCsvNotFound creates a BillingExportCsvNotFound with default headers values
func NewBillingExportCsvNotFound() *BillingExportCsvNotFound {
	return &BillingExportCsvNotFound{}
}

/* BillingExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BillingExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

func (o *BillingExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvNotFound  %+v", 404, o.Payload)
}
func (o *BillingExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BillingExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingExportCsvInternalServerError creates a BillingExportCsvInternalServerError with default headers values
func NewBillingExportCsvInternalServerError() *BillingExportCsvInternalServerError {
	return &BillingExportCsvInternalServerError{}
}

/* BillingExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BillingExportCsvInternalServerError struct {
}

func (o *BillingExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Billing/export][%d] billingExportCsvInternalServerError ", 500)
}

func (o *BillingExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
