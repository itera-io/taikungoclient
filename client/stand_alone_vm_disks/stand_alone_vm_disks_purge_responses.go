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

// StandAloneVMDisksPurgeReader is a Reader for the StandAloneVMDisksPurge structure.
type StandAloneVMDisksPurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksPurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksPurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksPurgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksPurgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksPurgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksPurgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksPurgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksPurgeOK creates a StandAloneVMDisksPurgeOK with default headers values
func NewStandAloneVMDisksPurgeOK() *StandAloneVMDisksPurgeOK {
	return &StandAloneVMDisksPurgeOK{}
}

/*
StandAloneVMDisksPurgeOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksPurgeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone Vm disks purge o k response has a 2xx status code
func (o *StandAloneVMDisksPurgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks purge o k response has a 3xx status code
func (o *StandAloneVMDisksPurgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge o k response has a 4xx status code
func (o *StandAloneVMDisksPurgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks purge o k response has a 5xx status code
func (o *StandAloneVMDisksPurgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks purge o k response a status code equal to that given
func (o *StandAloneVMDisksPurgeOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneVMDisksPurgeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksPurgeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksPurgeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneVMDisksPurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksPurgeBadRequest creates a StandAloneVMDisksPurgeBadRequest with default headers values
func NewStandAloneVMDisksPurgeBadRequest() *StandAloneVMDisksPurgeBadRequest {
	return &StandAloneVMDisksPurgeBadRequest{}
}

/*
StandAloneVMDisksPurgeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksPurgeBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this stand alone Vm disks purge bad request response has a 2xx status code
func (o *StandAloneVMDisksPurgeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks purge bad request response has a 3xx status code
func (o *StandAloneVMDisksPurgeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge bad request response has a 4xx status code
func (o *StandAloneVMDisksPurgeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks purge bad request response has a 5xx status code
func (o *StandAloneVMDisksPurgeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks purge bad request response a status code equal to that given
func (o *StandAloneVMDisksPurgeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneVMDisksPurgeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksPurgeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksPurgeBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksPurgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksPurgeUnauthorized creates a StandAloneVMDisksPurgeUnauthorized with default headers values
func NewStandAloneVMDisksPurgeUnauthorized() *StandAloneVMDisksPurgeUnauthorized {
	return &StandAloneVMDisksPurgeUnauthorized{}
}

/*
StandAloneVMDisksPurgeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksPurgeUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks purge unauthorized response has a 2xx status code
func (o *StandAloneVMDisksPurgeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks purge unauthorized response has a 3xx status code
func (o *StandAloneVMDisksPurgeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge unauthorized response has a 4xx status code
func (o *StandAloneVMDisksPurgeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks purge unauthorized response has a 5xx status code
func (o *StandAloneVMDisksPurgeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks purge unauthorized response a status code equal to that given
func (o *StandAloneVMDisksPurgeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneVMDisksPurgeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksPurgeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksPurgeUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksPurgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksPurgeForbidden creates a StandAloneVMDisksPurgeForbidden with default headers values
func NewStandAloneVMDisksPurgeForbidden() *StandAloneVMDisksPurgeForbidden {
	return &StandAloneVMDisksPurgeForbidden{}
}

/*
StandAloneVMDisksPurgeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksPurgeForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks purge forbidden response has a 2xx status code
func (o *StandAloneVMDisksPurgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks purge forbidden response has a 3xx status code
func (o *StandAloneVMDisksPurgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge forbidden response has a 4xx status code
func (o *StandAloneVMDisksPurgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks purge forbidden response has a 5xx status code
func (o *StandAloneVMDisksPurgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks purge forbidden response a status code equal to that given
func (o *StandAloneVMDisksPurgeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneVMDisksPurgeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksPurgeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksPurgeForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksPurgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksPurgeNotFound creates a StandAloneVMDisksPurgeNotFound with default headers values
func NewStandAloneVMDisksPurgeNotFound() *StandAloneVMDisksPurgeNotFound {
	return &StandAloneVMDisksPurgeNotFound{}
}

/*
StandAloneVMDisksPurgeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksPurgeNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks purge not found response has a 2xx status code
func (o *StandAloneVMDisksPurgeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks purge not found response has a 3xx status code
func (o *StandAloneVMDisksPurgeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge not found response has a 4xx status code
func (o *StandAloneVMDisksPurgeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks purge not found response has a 5xx status code
func (o *StandAloneVMDisksPurgeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks purge not found response a status code equal to that given
func (o *StandAloneVMDisksPurgeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneVMDisksPurgeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksPurgeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksPurgeNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksPurgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksPurgeInternalServerError creates a StandAloneVMDisksPurgeInternalServerError with default headers values
func NewStandAloneVMDisksPurgeInternalServerError() *StandAloneVMDisksPurgeInternalServerError {
	return &StandAloneVMDisksPurgeInternalServerError{}
}

/*
StandAloneVMDisksPurgeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksPurgeInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks purge internal server error response has a 2xx status code
func (o *StandAloneVMDisksPurgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks purge internal server error response has a 3xx status code
func (o *StandAloneVMDisksPurgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks purge internal server error response has a 4xx status code
func (o *StandAloneVMDisksPurgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks purge internal server error response has a 5xx status code
func (o *StandAloneVMDisksPurgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks purge internal server error response a status code equal to that given
func (o *StandAloneVMDisksPurgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneVMDisksPurgeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeInternalServerError ", 500)
}

func (o *StandAloneVMDisksPurgeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/purge][%d] standAloneVmDisksPurgeInternalServerError ", 500)
}

func (o *StandAloneVMDisksPurgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
