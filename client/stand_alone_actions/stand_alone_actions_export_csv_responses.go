// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneActionsExportCsvReader is a Reader for the StandAloneActionsExportCsv structure.
type StandAloneActionsExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsExportCsvOK creates a StandAloneActionsExportCsvOK with default headers values
func NewStandAloneActionsExportCsvOK() *StandAloneActionsExportCsvOK {
	return &StandAloneActionsExportCsvOK{}
}

/* StandAloneActionsExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsExportCsvOK struct {
}

func (o *StandAloneActionsExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvOK ", 200)
}

func (o *StandAloneActionsExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStandAloneActionsExportCsvBadRequest creates a StandAloneActionsExportCsvBadRequest with default headers values
func NewStandAloneActionsExportCsvBadRequest() *StandAloneActionsExportCsvBadRequest {
	return &StandAloneActionsExportCsvBadRequest{}
}

/* StandAloneActionsExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneActionsExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneActionsExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsExportCsvUnauthorized creates a StandAloneActionsExportCsvUnauthorized with default headers values
func NewStandAloneActionsExportCsvUnauthorized() *StandAloneActionsExportCsvUnauthorized {
	return &StandAloneActionsExportCsvUnauthorized{}
}

/* StandAloneActionsExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneActionsExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsExportCsvForbidden creates a StandAloneActionsExportCsvForbidden with default headers values
func NewStandAloneActionsExportCsvForbidden() *StandAloneActionsExportCsvForbidden {
	return &StandAloneActionsExportCsvForbidden{}
}

/* StandAloneActionsExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneActionsExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsExportCsvNotFound creates a StandAloneActionsExportCsvNotFound with default headers values
func NewStandAloneActionsExportCsvNotFound() *StandAloneActionsExportCsvNotFound {
	return &StandAloneActionsExportCsvNotFound{}
}

/* StandAloneActionsExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneActionsExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsExportCsvInternalServerError creates a StandAloneActionsExportCsvInternalServerError with default headers values
func NewStandAloneActionsExportCsvInternalServerError() *StandAloneActionsExportCsvInternalServerError {
	return &StandAloneActionsExportCsvInternalServerError{}
}

/* StandAloneActionsExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsExportCsvInternalServerError struct {
}

func (o *StandAloneActionsExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvInternalServerError ", 500)
}

func (o *StandAloneActionsExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
