// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAlonePurgeReader is a Reader for the StandAlonePurge structure.
type StandAlonePurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAlonePurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAlonePurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAlonePurgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAlonePurgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAlonePurgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAlonePurgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAlonePurgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAlonePurgeOK creates a StandAlonePurgeOK with default headers values
func NewStandAlonePurgeOK() *StandAlonePurgeOK {
	return &StandAlonePurgeOK{}
}

/*
StandAlonePurgeOK describes a response with status code 200, with default header values.

Success
*/
type StandAlonePurgeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone purge o k response has a 2xx status code
func (o *StandAlonePurgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone purge o k response has a 3xx status code
func (o *StandAlonePurgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge o k response has a 4xx status code
func (o *StandAlonePurgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone purge o k response has a 5xx status code
func (o *StandAlonePurgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone purge o k response a status code equal to that given
func (o *StandAlonePurgeOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAlonePurgeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeOK  %+v", 200, o.Payload)
}

func (o *StandAlonePurgeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeOK  %+v", 200, o.Payload)
}

func (o *StandAlonePurgeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAlonePurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAlonePurgeBadRequest creates a StandAlonePurgeBadRequest with default headers values
func NewStandAlonePurgeBadRequest() *StandAlonePurgeBadRequest {
	return &StandAlonePurgeBadRequest{}
}

/*
StandAlonePurgeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAlonePurgeBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone purge bad request response has a 2xx status code
func (o *StandAlonePurgeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone purge bad request response has a 3xx status code
func (o *StandAlonePurgeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge bad request response has a 4xx status code
func (o *StandAlonePurgeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone purge bad request response has a 5xx status code
func (o *StandAlonePurgeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone purge bad request response a status code equal to that given
func (o *StandAlonePurgeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAlonePurgeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAlonePurgeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAlonePurgeBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAlonePurgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAlonePurgeUnauthorized creates a StandAlonePurgeUnauthorized with default headers values
func NewStandAlonePurgeUnauthorized() *StandAlonePurgeUnauthorized {
	return &StandAlonePurgeUnauthorized{}
}

/*
StandAlonePurgeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAlonePurgeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone purge unauthorized response has a 2xx status code
func (o *StandAlonePurgeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone purge unauthorized response has a 3xx status code
func (o *StandAlonePurgeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge unauthorized response has a 4xx status code
func (o *StandAlonePurgeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone purge unauthorized response has a 5xx status code
func (o *StandAlonePurgeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone purge unauthorized response a status code equal to that given
func (o *StandAlonePurgeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAlonePurgeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAlonePurgeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAlonePurgeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAlonePurgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAlonePurgeForbidden creates a StandAlonePurgeForbidden with default headers values
func NewStandAlonePurgeForbidden() *StandAlonePurgeForbidden {
	return &StandAlonePurgeForbidden{}
}

/*
StandAlonePurgeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAlonePurgeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone purge forbidden response has a 2xx status code
func (o *StandAlonePurgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone purge forbidden response has a 3xx status code
func (o *StandAlonePurgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge forbidden response has a 4xx status code
func (o *StandAlonePurgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone purge forbidden response has a 5xx status code
func (o *StandAlonePurgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone purge forbidden response a status code equal to that given
func (o *StandAlonePurgeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAlonePurgeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeForbidden  %+v", 403, o.Payload)
}

func (o *StandAlonePurgeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeForbidden  %+v", 403, o.Payload)
}

func (o *StandAlonePurgeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAlonePurgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAlonePurgeNotFound creates a StandAlonePurgeNotFound with default headers values
func NewStandAlonePurgeNotFound() *StandAlonePurgeNotFound {
	return &StandAlonePurgeNotFound{}
}

/*
StandAlonePurgeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAlonePurgeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone purge not found response has a 2xx status code
func (o *StandAlonePurgeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone purge not found response has a 3xx status code
func (o *StandAlonePurgeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge not found response has a 4xx status code
func (o *StandAlonePurgeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone purge not found response has a 5xx status code
func (o *StandAlonePurgeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone purge not found response a status code equal to that given
func (o *StandAlonePurgeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAlonePurgeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeNotFound  %+v", 404, o.Payload)
}

func (o *StandAlonePurgeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeNotFound  %+v", 404, o.Payload)
}

func (o *StandAlonePurgeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAlonePurgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAlonePurgeInternalServerError creates a StandAlonePurgeInternalServerError with default headers values
func NewStandAlonePurgeInternalServerError() *StandAlonePurgeInternalServerError {
	return &StandAlonePurgeInternalServerError{}
}

/*
StandAlonePurgeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAlonePurgeInternalServerError struct {
}

// IsSuccess returns true when this stand alone purge internal server error response has a 2xx status code
func (o *StandAlonePurgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone purge internal server error response has a 3xx status code
func (o *StandAlonePurgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone purge internal server error response has a 4xx status code
func (o *StandAlonePurgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone purge internal server error response has a 5xx status code
func (o *StandAlonePurgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone purge internal server error response a status code equal to that given
func (o *StandAlonePurgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAlonePurgeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeInternalServerError ", 500)
}

func (o *StandAlonePurgeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/purge][%d] standAlonePurgeInternalServerError ", 500)
}

func (o *StandAlonePurgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
