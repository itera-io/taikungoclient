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

// StandAloneVMDisksUpdateReader is a Reader for the StandAloneVMDisksUpdate structure.
type StandAloneVMDisksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksUpdateOK creates a StandAloneVMDisksUpdateOK with default headers values
func NewStandAloneVMDisksUpdateOK() *StandAloneVMDisksUpdateOK {
	return &StandAloneVMDisksUpdateOK{}
}

/*
StandAloneVMDisksUpdateOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone Vm disks update o k response has a 2xx status code
func (o *StandAloneVMDisksUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks update o k response has a 3xx status code
func (o *StandAloneVMDisksUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update o k response has a 4xx status code
func (o *StandAloneVMDisksUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update o k response has a 5xx status code
func (o *StandAloneVMDisksUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update o k response a status code equal to that given
func (o *StandAloneVMDisksUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneVMDisksUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateBadRequest creates a StandAloneVMDisksUpdateBadRequest with default headers values
func NewStandAloneVMDisksUpdateBadRequest() *StandAloneVMDisksUpdateBadRequest {
	return &StandAloneVMDisksUpdateBadRequest{}
}

/*
StandAloneVMDisksUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this stand alone Vm disks update bad request response has a 2xx status code
func (o *StandAloneVMDisksUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update bad request response has a 3xx status code
func (o *StandAloneVMDisksUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update bad request response has a 4xx status code
func (o *StandAloneVMDisksUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update bad request response has a 5xx status code
func (o *StandAloneVMDisksUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update bad request response a status code equal to that given
func (o *StandAloneVMDisksUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneVMDisksUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateUnauthorized creates a StandAloneVMDisksUpdateUnauthorized with default headers values
func NewStandAloneVMDisksUpdateUnauthorized() *StandAloneVMDisksUpdateUnauthorized {
	return &StandAloneVMDisksUpdateUnauthorized{}
}

/*
StandAloneVMDisksUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksUpdateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks update unauthorized response has a 2xx status code
func (o *StandAloneVMDisksUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update unauthorized response has a 3xx status code
func (o *StandAloneVMDisksUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update unauthorized response has a 4xx status code
func (o *StandAloneVMDisksUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update unauthorized response has a 5xx status code
func (o *StandAloneVMDisksUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update unauthorized response a status code equal to that given
func (o *StandAloneVMDisksUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneVMDisksUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateForbidden creates a StandAloneVMDisksUpdateForbidden with default headers values
func NewStandAloneVMDisksUpdateForbidden() *StandAloneVMDisksUpdateForbidden {
	return &StandAloneVMDisksUpdateForbidden{}
}

/*
StandAloneVMDisksUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksUpdateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks update forbidden response has a 2xx status code
func (o *StandAloneVMDisksUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update forbidden response has a 3xx status code
func (o *StandAloneVMDisksUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update forbidden response has a 4xx status code
func (o *StandAloneVMDisksUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update forbidden response has a 5xx status code
func (o *StandAloneVMDisksUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update forbidden response a status code equal to that given
func (o *StandAloneVMDisksUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneVMDisksUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateNotFound creates a StandAloneVMDisksUpdateNotFound with default headers values
func NewStandAloneVMDisksUpdateNotFound() *StandAloneVMDisksUpdateNotFound {
	return &StandAloneVMDisksUpdateNotFound{}
}

/*
StandAloneVMDisksUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksUpdateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks update not found response has a 2xx status code
func (o *StandAloneVMDisksUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update not found response has a 3xx status code
func (o *StandAloneVMDisksUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update not found response has a 4xx status code
func (o *StandAloneVMDisksUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update not found response has a 5xx status code
func (o *StandAloneVMDisksUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update not found response a status code equal to that given
func (o *StandAloneVMDisksUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneVMDisksUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateInternalServerError creates a StandAloneVMDisksUpdateInternalServerError with default headers values
func NewStandAloneVMDisksUpdateInternalServerError() *StandAloneVMDisksUpdateInternalServerError {
	return &StandAloneVMDisksUpdateInternalServerError{}
}

/*
StandAloneVMDisksUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksUpdateInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks update internal server error response has a 2xx status code
func (o *StandAloneVMDisksUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update internal server error response has a 3xx status code
func (o *StandAloneVMDisksUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update internal server error response has a 4xx status code
func (o *StandAloneVMDisksUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update internal server error response has a 5xx status code
func (o *StandAloneVMDisksUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks update internal server error response a status code equal to that given
func (o *StandAloneVMDisksUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneVMDisksUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update][%d] standAloneVmDisksUpdateInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
