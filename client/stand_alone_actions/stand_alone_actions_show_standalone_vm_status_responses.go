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

// StandAloneActionsShowStandaloneVMStatusReader is a Reader for the StandAloneActionsShowStandaloneVMStatus structure.
type StandAloneActionsShowStandaloneVMStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsShowStandaloneVMStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsShowStandaloneVMStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsShowStandaloneVMStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsShowStandaloneVMStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsShowStandaloneVMStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsShowStandaloneVMStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsShowStandaloneVMStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsShowStandaloneVMStatusOK creates a StandAloneActionsShowStandaloneVMStatusOK with default headers values
func NewStandAloneActionsShowStandaloneVMStatusOK() *StandAloneActionsShowStandaloneVMStatusOK {
	return &StandAloneActionsShowStandaloneVMStatusOK{}
}

/* StandAloneActionsShowStandaloneVMStatusOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsShowStandaloneVMStatusOK struct {
	Payload string
}

func (o *StandAloneActionsShowStandaloneVMStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusOK  %+v", 200, o.Payload)
}
func (o *StandAloneActionsShowStandaloneVMStatusOK) GetPayload() string {
	return o.Payload
}

func (o *StandAloneActionsShowStandaloneVMStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShowStandaloneVMStatusBadRequest creates a StandAloneActionsShowStandaloneVMStatusBadRequest with default headers values
func NewStandAloneActionsShowStandaloneVMStatusBadRequest() *StandAloneActionsShowStandaloneVMStatusBadRequest {
	return &StandAloneActionsShowStandaloneVMStatusBadRequest{}
}

/* StandAloneActionsShowStandaloneVMStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsShowStandaloneVMStatusBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneActionsShowStandaloneVMStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneActionsShowStandaloneVMStatusBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsShowStandaloneVMStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShowStandaloneVMStatusUnauthorized creates a StandAloneActionsShowStandaloneVMStatusUnauthorized with default headers values
func NewStandAloneActionsShowStandaloneVMStatusUnauthorized() *StandAloneActionsShowStandaloneVMStatusUnauthorized {
	return &StandAloneActionsShowStandaloneVMStatusUnauthorized{}
}

/* StandAloneActionsShowStandaloneVMStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsShowStandaloneVMStatusUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsShowStandaloneVMStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneActionsShowStandaloneVMStatusUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsShowStandaloneVMStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShowStandaloneVMStatusForbidden creates a StandAloneActionsShowStandaloneVMStatusForbidden with default headers values
func NewStandAloneActionsShowStandaloneVMStatusForbidden() *StandAloneActionsShowStandaloneVMStatusForbidden {
	return &StandAloneActionsShowStandaloneVMStatusForbidden{}
}

/* StandAloneActionsShowStandaloneVMStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsShowStandaloneVMStatusForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsShowStandaloneVMStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneActionsShowStandaloneVMStatusForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsShowStandaloneVMStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShowStandaloneVMStatusNotFound creates a StandAloneActionsShowStandaloneVMStatusNotFound with default headers values
func NewStandAloneActionsShowStandaloneVMStatusNotFound() *StandAloneActionsShowStandaloneVMStatusNotFound {
	return &StandAloneActionsShowStandaloneVMStatusNotFound{}
}

/* StandAloneActionsShowStandaloneVMStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsShowStandaloneVMStatusNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneActionsShowStandaloneVMStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneActionsShowStandaloneVMStatusNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsShowStandaloneVMStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShowStandaloneVMStatusInternalServerError creates a StandAloneActionsShowStandaloneVMStatusInternalServerError with default headers values
func NewStandAloneActionsShowStandaloneVMStatusInternalServerError() *StandAloneActionsShowStandaloneVMStatusInternalServerError {
	return &StandAloneActionsShowStandaloneVMStatusInternalServerError{}
}

/* StandAloneActionsShowStandaloneVMStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsShowStandaloneVMStatusInternalServerError struct {
}

func (o *StandAloneActionsShowStandaloneVMStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/StandAloneActions/status/{id}][%d] standAloneActionsShowStandaloneVmStatusInternalServerError ", 500)
}

func (o *StandAloneActionsShowStandaloneVMStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
