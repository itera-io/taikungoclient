// Code generated by go-swagger; DO NOT EDIT.

package showback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackExportCsvReader is a Reader for the ShowbackExportCsv structure.
type ShowbackExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackExportCsvOK creates a ShowbackExportCsvOK with default headers values
func NewShowbackExportCsvOK() *ShowbackExportCsvOK {
	return &ShowbackExportCsvOK{}
}

/* ShowbackExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackExportCsvOK struct {
}

func (o *ShowbackExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvOK ", 200)
}

func (o *ShowbackExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowbackExportCsvBadRequest creates a ShowbackExportCsvBadRequest with default headers values
func NewShowbackExportCsvBadRequest() *ShowbackExportCsvBadRequest {
	return &ShowbackExportCsvBadRequest{}
}

/* ShowbackExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ShowbackExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvBadRequest  %+v", 400, o.Payload)
}
func (o *ShowbackExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackExportCsvUnauthorized creates a ShowbackExportCsvUnauthorized with default headers values
func NewShowbackExportCsvUnauthorized() *ShowbackExportCsvUnauthorized {
	return &ShowbackExportCsvUnauthorized{}
}

/* ShowbackExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvUnauthorized  %+v", 401, o.Payload)
}
func (o *ShowbackExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackExportCsvForbidden creates a ShowbackExportCsvForbidden with default headers values
func NewShowbackExportCsvForbidden() *ShowbackExportCsvForbidden {
	return &ShowbackExportCsvForbidden{}
}

/* ShowbackExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvForbidden  %+v", 403, o.Payload)
}
func (o *ShowbackExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackExportCsvNotFound creates a ShowbackExportCsvNotFound with default headers values
func NewShowbackExportCsvNotFound() *ShowbackExportCsvNotFound {
	return &ShowbackExportCsvNotFound{}
}

/* ShowbackExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ShowbackExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvNotFound  %+v", 404, o.Payload)
}
func (o *ShowbackExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ShowbackExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackExportCsvInternalServerError creates a ShowbackExportCsvInternalServerError with default headers values
func NewShowbackExportCsvInternalServerError() *ShowbackExportCsvInternalServerError {
	return &ShowbackExportCsvInternalServerError{}
}

/* ShowbackExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackExportCsvInternalServerError struct {
}

func (o *ShowbackExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Showback/export][%d] showbackExportCsvInternalServerError ", 500)
}

func (o *ShowbackExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
