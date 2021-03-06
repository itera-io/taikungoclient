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

// StandAloneActionsStopReader is a Reader for the StandAloneActionsStop structure.
type StandAloneActionsStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsStopUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsStopForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsStopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsStopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsStopOK creates a StandAloneActionsStopOK with default headers values
func NewStandAloneActionsStopOK() *StandAloneActionsStopOK {
	return &StandAloneActionsStopOK{}
}

/* StandAloneActionsStopOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsStopOK struct {
	Payload models.Unit
}

func (o *StandAloneActionsStopOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopOK  %+v", 200, o.Payload)
}
func (o *StandAloneActionsStopOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneActionsStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStopBadRequest creates a StandAloneActionsStopBadRequest with default headers values
func NewStandAloneActionsStopBadRequest() *StandAloneActionsStopBadRequest {
	return &StandAloneActionsStopBadRequest{}
}

/* StandAloneActionsStopBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsStopBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneActionsStopBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneActionsStopBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStopUnauthorized creates a StandAloneActionsStopUnauthorized with default headers values
func NewStandAloneActionsStopUnauthorized() *StandAloneActionsStopUnauthorized {
	return &StandAloneActionsStopUnauthorized{}
}

/* StandAloneActionsStopUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsStopUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsStopUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneActionsStopUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsStopUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStopForbidden creates a StandAloneActionsStopForbidden with default headers values
func NewStandAloneActionsStopForbidden() *StandAloneActionsStopForbidden {
	return &StandAloneActionsStopForbidden{}
}

/* StandAloneActionsStopForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsStopForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsStopForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneActionsStopForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsStopForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStopNotFound creates a StandAloneActionsStopNotFound with default headers values
func NewStandAloneActionsStopNotFound() *StandAloneActionsStopNotFound {
	return &StandAloneActionsStopNotFound{}
}

/* StandAloneActionsStopNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsStopNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsStopNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneActionsStopNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsStopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStopInternalServerError creates a StandAloneActionsStopInternalServerError with default headers values
func NewStandAloneActionsStopInternalServerError() *StandAloneActionsStopInternalServerError {
	return &StandAloneActionsStopInternalServerError{}
}

/* StandAloneActionsStopInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsStopInternalServerError struct {
}

func (o *StandAloneActionsStopInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/stop][%d] standAloneActionsStopInternalServerError ", 500)
}

func (o *StandAloneActionsStopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
