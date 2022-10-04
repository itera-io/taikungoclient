// Code generated by go-swagger; DO NOT EDIT.

package showback_summaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ShowbackSummariesExportCsvReader is a Reader for the ShowbackSummariesExportCsv structure.
type ShowbackSummariesExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowbackSummariesExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowbackSummariesExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewShowbackSummariesExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewShowbackSummariesExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowbackSummariesExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowbackSummariesExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowbackSummariesExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewShowbackSummariesExportCsvOK creates a ShowbackSummariesExportCsvOK with default headers values
func NewShowbackSummariesExportCsvOK() *ShowbackSummariesExportCsvOK {
	return &ShowbackSummariesExportCsvOK{}
}

/*
ShowbackSummariesExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type ShowbackSummariesExportCsvOK struct {
}

// IsSuccess returns true when this showback summaries export csv o k response has a 2xx status code
func (o *ShowbackSummariesExportCsvOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this showback summaries export csv o k response has a 3xx status code
func (o *ShowbackSummariesExportCsvOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv o k response has a 4xx status code
func (o *ShowbackSummariesExportCsvOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries export csv o k response has a 5xx status code
func (o *ShowbackSummariesExportCsvOK) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries export csv o k response a status code equal to that given
func (o *ShowbackSummariesExportCsvOK) IsCode(code int) bool {
	return code == 200
}

func (o *ShowbackSummariesExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvOK ", 200)
}

func (o *ShowbackSummariesExportCsvOK) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvOK ", 200)
}

func (o *ShowbackSummariesExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowbackSummariesExportCsvBadRequest creates a ShowbackSummariesExportCsvBadRequest with default headers values
func NewShowbackSummariesExportCsvBadRequest() *ShowbackSummariesExportCsvBadRequest {
	return &ShowbackSummariesExportCsvBadRequest{}
}

/*
ShowbackSummariesExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ShowbackSummariesExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this showback summaries export csv bad request response has a 2xx status code
func (o *ShowbackSummariesExportCsvBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries export csv bad request response has a 3xx status code
func (o *ShowbackSummariesExportCsvBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv bad request response has a 4xx status code
func (o *ShowbackSummariesExportCsvBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries export csv bad request response has a 5xx status code
func (o *ShowbackSummariesExportCsvBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries export csv bad request response a status code equal to that given
func (o *ShowbackSummariesExportCsvBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ShowbackSummariesExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesExportCsvBadRequest) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *ShowbackSummariesExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ShowbackSummariesExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesExportCsvUnauthorized creates a ShowbackSummariesExportCsvUnauthorized with default headers values
func NewShowbackSummariesExportCsvUnauthorized() *ShowbackSummariesExportCsvUnauthorized {
	return &ShowbackSummariesExportCsvUnauthorized{}
}

/*
ShowbackSummariesExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowbackSummariesExportCsvUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries export csv unauthorized response has a 2xx status code
func (o *ShowbackSummariesExportCsvUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries export csv unauthorized response has a 3xx status code
func (o *ShowbackSummariesExportCsvUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv unauthorized response has a 4xx status code
func (o *ShowbackSummariesExportCsvUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries export csv unauthorized response has a 5xx status code
func (o *ShowbackSummariesExportCsvUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries export csv unauthorized response a status code equal to that given
func (o *ShowbackSummariesExportCsvUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ShowbackSummariesExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesExportCsvUnauthorized) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *ShowbackSummariesExportCsvUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesExportCsvForbidden creates a ShowbackSummariesExportCsvForbidden with default headers values
func NewShowbackSummariesExportCsvForbidden() *ShowbackSummariesExportCsvForbidden {
	return &ShowbackSummariesExportCsvForbidden{}
}

/*
ShowbackSummariesExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowbackSummariesExportCsvForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries export csv forbidden response has a 2xx status code
func (o *ShowbackSummariesExportCsvForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries export csv forbidden response has a 3xx status code
func (o *ShowbackSummariesExportCsvForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv forbidden response has a 4xx status code
func (o *ShowbackSummariesExportCsvForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries export csv forbidden response has a 5xx status code
func (o *ShowbackSummariesExportCsvForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries export csv forbidden response a status code equal to that given
func (o *ShowbackSummariesExportCsvForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ShowbackSummariesExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesExportCsvForbidden) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *ShowbackSummariesExportCsvForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesExportCsvNotFound creates a ShowbackSummariesExportCsvNotFound with default headers values
func NewShowbackSummariesExportCsvNotFound() *ShowbackSummariesExportCsvNotFound {
	return &ShowbackSummariesExportCsvNotFound{}
}

/*
ShowbackSummariesExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ShowbackSummariesExportCsvNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this showback summaries export csv not found response has a 2xx status code
func (o *ShowbackSummariesExportCsvNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries export csv not found response has a 3xx status code
func (o *ShowbackSummariesExportCsvNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv not found response has a 4xx status code
func (o *ShowbackSummariesExportCsvNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this showback summaries export csv not found response has a 5xx status code
func (o *ShowbackSummariesExportCsvNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this showback summaries export csv not found response a status code equal to that given
func (o *ShowbackSummariesExportCsvNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ShowbackSummariesExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesExportCsvNotFound) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *ShowbackSummariesExportCsvNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ShowbackSummariesExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowbackSummariesExportCsvInternalServerError creates a ShowbackSummariesExportCsvInternalServerError with default headers values
func NewShowbackSummariesExportCsvInternalServerError() *ShowbackSummariesExportCsvInternalServerError {
	return &ShowbackSummariesExportCsvInternalServerError{}
}

/*
ShowbackSummariesExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ShowbackSummariesExportCsvInternalServerError struct {
}

// IsSuccess returns true when this showback summaries export csv internal server error response has a 2xx status code
func (o *ShowbackSummariesExportCsvInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this showback summaries export csv internal server error response has a 3xx status code
func (o *ShowbackSummariesExportCsvInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this showback summaries export csv internal server error response has a 4xx status code
func (o *ShowbackSummariesExportCsvInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this showback summaries export csv internal server error response has a 5xx status code
func (o *ShowbackSummariesExportCsvInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this showback summaries export csv internal server error response a status code equal to that given
func (o *ShowbackSummariesExportCsvInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ShowbackSummariesExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvInternalServerError ", 500)
}

func (o *ShowbackSummariesExportCsvInternalServerError) String() string {
	return fmt.Sprintf("[GET /showback/v{v}/ShowbackSummaries/export][%d] showbackSummariesExportCsvInternalServerError ", 500)
}

func (o *ShowbackSummariesExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
