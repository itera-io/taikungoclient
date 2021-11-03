// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// RequestExportCsvReader is a Reader for the RequestExportCsv structure.
type RequestExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRequestExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRequestExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRequestExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRequestExportCsvTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequestExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequestExportCsvOK creates a RequestExportCsvOK with default headers values
func NewRequestExportCsvOK() *RequestExportCsvOK {
	return &RequestExportCsvOK{}
}

/* RequestExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type RequestExportCsvOK struct {
}

func (o *RequestExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvOK ", 200)
}

func (o *RequestExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestExportCsvBadRequest creates a RequestExportCsvBadRequest with default headers values
func NewRequestExportCsvBadRequest() *RequestExportCsvBadRequest {
	return &RequestExportCsvBadRequest{}
}

/* RequestExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RequestExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *RequestExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvBadRequest  %+v", 400, o.Payload)
}
func (o *RequestExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *RequestExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestExportCsvUnauthorized creates a RequestExportCsvUnauthorized with default headers values
func NewRequestExportCsvUnauthorized() *RequestExportCsvUnauthorized {
	return &RequestExportCsvUnauthorized{}
}

/* RequestExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RequestExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *RequestExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvUnauthorized  %+v", 401, o.Payload)
}
func (o *RequestExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RequestExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestExportCsvForbidden creates a RequestExportCsvForbidden with default headers values
func NewRequestExportCsvForbidden() *RequestExportCsvForbidden {
	return &RequestExportCsvForbidden{}
}

/* RequestExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RequestExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

func (o *RequestExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvForbidden  %+v", 403, o.Payload)
}
func (o *RequestExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RequestExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestExportCsvNotFound creates a RequestExportCsvNotFound with default headers values
func NewRequestExportCsvNotFound() *RequestExportCsvNotFound {
	return &RequestExportCsvNotFound{}
}

/* RequestExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RequestExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

func (o *RequestExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvNotFound  %+v", 404, o.Payload)
}
func (o *RequestExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RequestExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestExportCsvTooManyRequests creates a RequestExportCsvTooManyRequests with default headers values
func NewRequestExportCsvTooManyRequests() *RequestExportCsvTooManyRequests {
	return &RequestExportCsvTooManyRequests{}
}

/* RequestExportCsvTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type RequestExportCsvTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *RequestExportCsvTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvTooManyRequests  %+v", 429, o.Payload)
}
func (o *RequestExportCsvTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *RequestExportCsvTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestExportCsvInternalServerError creates a RequestExportCsvInternalServerError with default headers values
func NewRequestExportCsvInternalServerError() *RequestExportCsvInternalServerError {
	return &RequestExportCsvInternalServerError{}
}

/* RequestExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type RequestExportCsvInternalServerError struct {
}

func (o *RequestExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Request/download][%d] requestExportCsvInternalServerError ", 500)
}

func (o *RequestExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
