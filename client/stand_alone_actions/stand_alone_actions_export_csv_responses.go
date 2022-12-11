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

/*
StandAloneActionsExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsExportCsvOK struct {
}

// IsSuccess returns true when this stand alone actions export csv o k response has a 2xx status code
func (o *StandAloneActionsExportCsvOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions export csv o k response has a 3xx status code
func (o *StandAloneActionsExportCsvOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv o k response has a 4xx status code
func (o *StandAloneActionsExportCsvOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions export csv o k response has a 5xx status code
func (o *StandAloneActionsExportCsvOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions export csv o k response a status code equal to that given
func (o *StandAloneActionsExportCsvOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneActionsExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvOK ", 200)
}

func (o *StandAloneActionsExportCsvOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvOK ", 200)
}

func (o *StandAloneActionsExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStandAloneActionsExportCsvBadRequest creates a StandAloneActionsExportCsvBadRequest with default headers values
func NewStandAloneActionsExportCsvBadRequest() *StandAloneActionsExportCsvBadRequest {
	return &StandAloneActionsExportCsvBadRequest{}
}

/*
StandAloneActionsExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsExportCsvBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions export csv bad request response has a 2xx status code
func (o *StandAloneActionsExportCsvBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions export csv bad request response has a 3xx status code
func (o *StandAloneActionsExportCsvBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv bad request response has a 4xx status code
func (o *StandAloneActionsExportCsvBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions export csv bad request response has a 5xx status code
func (o *StandAloneActionsExportCsvBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions export csv bad request response a status code equal to that given
func (o *StandAloneActionsExportCsvBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneActionsExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsExportCsvBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsExportCsvBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsExportCsvUnauthorized creates a StandAloneActionsExportCsvUnauthorized with default headers values
func NewStandAloneActionsExportCsvUnauthorized() *StandAloneActionsExportCsvUnauthorized {
	return &StandAloneActionsExportCsvUnauthorized{}
}

/*
StandAloneActionsExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions export csv unauthorized response has a 2xx status code
func (o *StandAloneActionsExportCsvUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions export csv unauthorized response has a 3xx status code
func (o *StandAloneActionsExportCsvUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv unauthorized response has a 4xx status code
func (o *StandAloneActionsExportCsvUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions export csv unauthorized response has a 5xx status code
func (o *StandAloneActionsExportCsvUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions export csv unauthorized response a status code equal to that given
func (o *StandAloneActionsExportCsvUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneActionsExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsExportCsvUnauthorized) String() string {
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

/*
StandAloneActionsExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions export csv forbidden response has a 2xx status code
func (o *StandAloneActionsExportCsvForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions export csv forbidden response has a 3xx status code
func (o *StandAloneActionsExportCsvForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv forbidden response has a 4xx status code
func (o *StandAloneActionsExportCsvForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions export csv forbidden response has a 5xx status code
func (o *StandAloneActionsExportCsvForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions export csv forbidden response a status code equal to that given
func (o *StandAloneActionsExportCsvForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneActionsExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsExportCsvForbidden) String() string {
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

/*
StandAloneActionsExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions export csv not found response has a 2xx status code
func (o *StandAloneActionsExportCsvNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions export csv not found response has a 3xx status code
func (o *StandAloneActionsExportCsvNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv not found response has a 4xx status code
func (o *StandAloneActionsExportCsvNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions export csv not found response has a 5xx status code
func (o *StandAloneActionsExportCsvNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions export csv not found response a status code equal to that given
func (o *StandAloneActionsExportCsvNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneActionsExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsExportCsvNotFound) String() string {
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

/*
StandAloneActionsExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsExportCsvInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions export csv internal server error response has a 2xx status code
func (o *StandAloneActionsExportCsvInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions export csv internal server error response has a 3xx status code
func (o *StandAloneActionsExportCsvInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions export csv internal server error response has a 4xx status code
func (o *StandAloneActionsExportCsvInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions export csv internal server error response has a 5xx status code
func (o *StandAloneActionsExportCsvInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions export csv internal server error response a status code equal to that given
func (o *StandAloneActionsExportCsvInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneActionsExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvInternalServerError ", 500)
}

func (o *StandAloneActionsExportCsvInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/download/rdp/{id}][%d] standAloneActionsExportCsvInternalServerError ", 500)
}

func (o *StandAloneActionsExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
