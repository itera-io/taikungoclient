// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_vm_disks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneVMDisksCreateReader is a Reader for the StandAloneVMDisksCreate structure.
type StandAloneVMDisksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksCreateOK creates a StandAloneVMDisksCreateOK with default headers values
func NewStandAloneVMDisksCreateOK() *StandAloneVMDisksCreateOK {
	return &StandAloneVMDisksCreateOK{}
}

/* StandAloneVMDisksCreateOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksCreateOK struct {
	Payload *models.APIResponse
}

func (o *StandAloneVMDisksCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateOK  %+v", 200, o.Payload)
}
func (o *StandAloneVMDisksCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *StandAloneVMDisksCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateBadRequest creates a StandAloneVMDisksCreateBadRequest with default headers values
func NewStandAloneVMDisksCreateBadRequest() *StandAloneVMDisksCreateBadRequest {
	return &StandAloneVMDisksCreateBadRequest{}
}

/* StandAloneVMDisksCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *StandAloneVMDisksCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateBadRequest  %+v", 400, o.Payload)
}
func (o *StandAloneVMDisksCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateUnauthorized creates a StandAloneVMDisksCreateUnauthorized with default headers values
func NewStandAloneVMDisksCreateUnauthorized() *StandAloneVMDisksCreateUnauthorized {
	return &StandAloneVMDisksCreateUnauthorized{}
}

/* StandAloneVMDisksCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneVMDisksCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *StandAloneVMDisksCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateForbidden creates a StandAloneVMDisksCreateForbidden with default headers values
func NewStandAloneVMDisksCreateForbidden() *StandAloneVMDisksCreateForbidden {
	return &StandAloneVMDisksCreateForbidden{}
}

/* StandAloneVMDisksCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneVMDisksCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateForbidden  %+v", 403, o.Payload)
}
func (o *StandAloneVMDisksCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateNotFound creates a StandAloneVMDisksCreateNotFound with default headers values
func NewStandAloneVMDisksCreateNotFound() *StandAloneVMDisksCreateNotFound {
	return &StandAloneVMDisksCreateNotFound{}
}

/* StandAloneVMDisksCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *StandAloneVMDisksCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateNotFound  %+v", 404, o.Payload)
}
func (o *StandAloneVMDisksCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateInternalServerError creates a StandAloneVMDisksCreateInternalServerError with default headers values
func NewStandAloneVMDisksCreateInternalServerError() *StandAloneVMDisksCreateInternalServerError {
	return &StandAloneVMDisksCreateInternalServerError{}
}

/* StandAloneVMDisksCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksCreateInternalServerError struct {
}

func (o *StandAloneVMDisksCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateInternalServerError ", 500)
}

func (o *StandAloneVMDisksCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
