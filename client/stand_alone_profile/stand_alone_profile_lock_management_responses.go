// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneProfileLockManagementReader is a Reader for the StandAloneProfileLockManagement structure.
type StandAloneProfileLockManagementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProfileLockManagementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProfileLockManagementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProfileLockManagementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProfileLockManagementUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProfileLockManagementForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProfileLockManagementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProfileLockManagementInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProfileLockManagementOK creates a StandAloneProfileLockManagementOK with default headers values
func NewStandAloneProfileLockManagementOK() *StandAloneProfileLockManagementOK {
	return &StandAloneProfileLockManagementOK{}
}

/* StandAloneProfileLockManagementOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProfileLockManagementOK struct {
	Payload models.Unit
}

func (o *StandAloneProfileLockManagementOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementOK  %+v", 200, o.Payload)
}
func (o *StandAloneProfileLockManagementOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneProfileLockManagementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileLockManagementBadRequest creates a StandAloneProfileLockManagementBadRequest with default headers values
func NewStandAloneProfileLockManagementBadRequest() *StandAloneProfileLockManagementBadRequest {
	return &StandAloneProfileLockManagementBadRequest{}
}

/* StandAloneProfileLockManagementBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProfileLockManagementBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneProfileLockManagementBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneProfileLockManagementBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileLockManagementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileLockManagementUnauthorized creates a StandAloneProfileLockManagementUnauthorized with default headers values
func NewStandAloneProfileLockManagementUnauthorized() *StandAloneProfileLockManagementUnauthorized {
	return &StandAloneProfileLockManagementUnauthorized{}
}

/* StandAloneProfileLockManagementUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProfileLockManagementUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileLockManagementUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneProfileLockManagementUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileLockManagementUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileLockManagementForbidden creates a StandAloneProfileLockManagementForbidden with default headers values
func NewStandAloneProfileLockManagementForbidden() *StandAloneProfileLockManagementForbidden {
	return &StandAloneProfileLockManagementForbidden{}
}

/* StandAloneProfileLockManagementForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProfileLockManagementForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileLockManagementForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneProfileLockManagementForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileLockManagementForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileLockManagementNotFound creates a StandAloneProfileLockManagementNotFound with default headers values
func NewStandAloneProfileLockManagementNotFound() *StandAloneProfileLockManagementNotFound {
	return &StandAloneProfileLockManagementNotFound{}
}

/* StandAloneProfileLockManagementNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProfileLockManagementNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneProfileLockManagementNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneProfileLockManagementNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileLockManagementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileLockManagementInternalServerError creates a StandAloneProfileLockManagementInternalServerError with default headers values
func NewStandAloneProfileLockManagementInternalServerError() *StandAloneProfileLockManagementInternalServerError {
	return &StandAloneProfileLockManagementInternalServerError{}
}

/* StandAloneProfileLockManagementInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProfileLockManagementInternalServerError struct {
}

func (o *StandAloneProfileLockManagementInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/lockmanager][%d] standAloneProfileLockManagementInternalServerError ", 500)
}

func (o *StandAloneProfileLockManagementInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
