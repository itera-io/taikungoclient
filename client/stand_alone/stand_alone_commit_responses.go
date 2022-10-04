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

// StandAloneCommitReader is a Reader for the StandAloneCommit structure.
type StandAloneCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneCommitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneCommitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneCommitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneCommitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneCommitOK creates a StandAloneCommitOK with default headers values
func NewStandAloneCommitOK() *StandAloneCommitOK {
	return &StandAloneCommitOK{}
}

/*
StandAloneCommitOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneCommitOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone commit o k response has a 2xx status code
func (o *StandAloneCommitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone commit o k response has a 3xx status code
func (o *StandAloneCommitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit o k response has a 4xx status code
func (o *StandAloneCommitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone commit o k response has a 5xx status code
func (o *StandAloneCommitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone commit o k response a status code equal to that given
func (o *StandAloneCommitOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneCommitOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitOK  %+v", 200, o.Payload)
}

func (o *StandAloneCommitOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitOK  %+v", 200, o.Payload)
}

func (o *StandAloneCommitOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCommitBadRequest creates a StandAloneCommitBadRequest with default headers values
func NewStandAloneCommitBadRequest() *StandAloneCommitBadRequest {
	return &StandAloneCommitBadRequest{}
}

/*
StandAloneCommitBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneCommitBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this stand alone commit bad request response has a 2xx status code
func (o *StandAloneCommitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone commit bad request response has a 3xx status code
func (o *StandAloneCommitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit bad request response has a 4xx status code
func (o *StandAloneCommitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone commit bad request response has a 5xx status code
func (o *StandAloneCommitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone commit bad request response a status code equal to that given
func (o *StandAloneCommitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneCommitBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneCommitBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneCommitBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneCommitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCommitUnauthorized creates a StandAloneCommitUnauthorized with default headers values
func NewStandAloneCommitUnauthorized() *StandAloneCommitUnauthorized {
	return &StandAloneCommitUnauthorized{}
}

/*
StandAloneCommitUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneCommitUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone commit unauthorized response has a 2xx status code
func (o *StandAloneCommitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone commit unauthorized response has a 3xx status code
func (o *StandAloneCommitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit unauthorized response has a 4xx status code
func (o *StandAloneCommitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone commit unauthorized response has a 5xx status code
func (o *StandAloneCommitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone commit unauthorized response a status code equal to that given
func (o *StandAloneCommitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneCommitUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneCommitUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneCommitUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneCommitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCommitForbidden creates a StandAloneCommitForbidden with default headers values
func NewStandAloneCommitForbidden() *StandAloneCommitForbidden {
	return &StandAloneCommitForbidden{}
}

/*
StandAloneCommitForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneCommitForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone commit forbidden response has a 2xx status code
func (o *StandAloneCommitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone commit forbidden response has a 3xx status code
func (o *StandAloneCommitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit forbidden response has a 4xx status code
func (o *StandAloneCommitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone commit forbidden response has a 5xx status code
func (o *StandAloneCommitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone commit forbidden response a status code equal to that given
func (o *StandAloneCommitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneCommitForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneCommitForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneCommitForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneCommitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCommitNotFound creates a StandAloneCommitNotFound with default headers values
func NewStandAloneCommitNotFound() *StandAloneCommitNotFound {
	return &StandAloneCommitNotFound{}
}

/*
StandAloneCommitNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneCommitNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone commit not found response has a 2xx status code
func (o *StandAloneCommitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone commit not found response has a 3xx status code
func (o *StandAloneCommitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit not found response has a 4xx status code
func (o *StandAloneCommitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone commit not found response has a 5xx status code
func (o *StandAloneCommitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone commit not found response a status code equal to that given
func (o *StandAloneCommitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneCommitNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneCommitNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneCommitNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneCommitInternalServerError creates a StandAloneCommitInternalServerError with default headers values
func NewStandAloneCommitInternalServerError() *StandAloneCommitInternalServerError {
	return &StandAloneCommitInternalServerError{}
}

/*
StandAloneCommitInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneCommitInternalServerError struct {
}

// IsSuccess returns true when this stand alone commit internal server error response has a 2xx status code
func (o *StandAloneCommitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone commit internal server error response has a 3xx status code
func (o *StandAloneCommitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone commit internal server error response has a 4xx status code
func (o *StandAloneCommitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone commit internal server error response has a 5xx status code
func (o *StandAloneCommitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone commit internal server error response a status code equal to that given
func (o *StandAloneCommitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneCommitInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitInternalServerError ", 500)
}

func (o *StandAloneCommitInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/commit][%d] standAloneCommitInternalServerError ", 500)
}

func (o *StandAloneCommitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
