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

// StandAloneUpdateReader is a Reader for the StandAloneUpdate structure.
type StandAloneUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneUpdateOK creates a StandAloneUpdateOK with default headers values
func NewStandAloneUpdateOK() *StandAloneUpdateOK {
	return &StandAloneUpdateOK{}
}

/*
StandAloneUpdateOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone update o k response has a 2xx status code
func (o *StandAloneUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone update o k response has a 3xx status code
func (o *StandAloneUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update o k response has a 4xx status code
func (o *StandAloneUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone update o k response has a 5xx status code
func (o *StandAloneUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update o k response a status code equal to that given
func (o *StandAloneUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateOK  %+v", 200, o.Payload)
}

func (o *StandAloneUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateOK  %+v", 200, o.Payload)
}

func (o *StandAloneUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateBadRequest creates a StandAloneUpdateBadRequest with default headers values
func NewStandAloneUpdateBadRequest() *StandAloneUpdateBadRequest {
	return &StandAloneUpdateBadRequest{}
}

/*
StandAloneUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone update bad request response has a 2xx status code
func (o *StandAloneUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update bad request response has a 3xx status code
func (o *StandAloneUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update bad request response has a 4xx status code
func (o *StandAloneUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update bad request response has a 5xx status code
func (o *StandAloneUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update bad request response a status code equal to that given
func (o *StandAloneUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateUnauthorized creates a StandAloneUpdateUnauthorized with default headers values
func NewStandAloneUpdateUnauthorized() *StandAloneUpdateUnauthorized {
	return &StandAloneUpdateUnauthorized{}
}

/*
StandAloneUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update unauthorized response has a 2xx status code
func (o *StandAloneUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update unauthorized response has a 3xx status code
func (o *StandAloneUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update unauthorized response has a 4xx status code
func (o *StandAloneUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update unauthorized response has a 5xx status code
func (o *StandAloneUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update unauthorized response a status code equal to that given
func (o *StandAloneUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateForbidden creates a StandAloneUpdateForbidden with default headers values
func NewStandAloneUpdateForbidden() *StandAloneUpdateForbidden {
	return &StandAloneUpdateForbidden{}
}

/*
StandAloneUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update forbidden response has a 2xx status code
func (o *StandAloneUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update forbidden response has a 3xx status code
func (o *StandAloneUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update forbidden response has a 4xx status code
func (o *StandAloneUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update forbidden response has a 5xx status code
func (o *StandAloneUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update forbidden response a status code equal to that given
func (o *StandAloneUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateNotFound creates a StandAloneUpdateNotFound with default headers values
func NewStandAloneUpdateNotFound() *StandAloneUpdateNotFound {
	return &StandAloneUpdateNotFound{}
}

/*
StandAloneUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone update not found response has a 2xx status code
func (o *StandAloneUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update not found response has a 3xx status code
func (o *StandAloneUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update not found response has a 4xx status code
func (o *StandAloneUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone update not found response has a 5xx status code
func (o *StandAloneUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone update not found response a status code equal to that given
func (o *StandAloneUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneUpdateInternalServerError creates a StandAloneUpdateInternalServerError with default headers values
func NewStandAloneUpdateInternalServerError() *StandAloneUpdateInternalServerError {
	return &StandAloneUpdateInternalServerError{}
}

/*
StandAloneUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneUpdateInternalServerError struct {
}

// IsSuccess returns true when this stand alone update internal server error response has a 2xx status code
func (o *StandAloneUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone update internal server error response has a 3xx status code
func (o *StandAloneUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone update internal server error response has a 4xx status code
func (o *StandAloneUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone update internal server error response has a 5xx status code
func (o *StandAloneUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone update internal server error response a status code equal to that given
func (o *StandAloneUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateInternalServerError ", 500)
}

func (o *StandAloneUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/update][%d] standAloneUpdateInternalServerError ", 500)
}

func (o *StandAloneUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
