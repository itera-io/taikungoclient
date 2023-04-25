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

/*
StandAloneVMDisksCreateOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this stand alone Vm disks create o k response has a 2xx status code
func (o *StandAloneVMDisksCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks create o k response has a 3xx status code
func (o *StandAloneVMDisksCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create o k response has a 4xx status code
func (o *StandAloneVMDisksCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks create o k response has a 5xx status code
func (o *StandAloneVMDisksCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks create o k response a status code equal to that given
func (o *StandAloneVMDisksCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone Vm disks create o k response
func (o *StandAloneVMDisksCreateOK) Code() int {
	return 200
}

func (o *StandAloneVMDisksCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksCreateOK) String() string {
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

/*
StandAloneVMDisksCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks create bad request response has a 2xx status code
func (o *StandAloneVMDisksCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks create bad request response has a 3xx status code
func (o *StandAloneVMDisksCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create bad request response has a 4xx status code
func (o *StandAloneVMDisksCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks create bad request response has a 5xx status code
func (o *StandAloneVMDisksCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks create bad request response a status code equal to that given
func (o *StandAloneVMDisksCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone Vm disks create bad request response
func (o *StandAloneVMDisksCreateBadRequest) Code() int {
	return 400
}

func (o *StandAloneVMDisksCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateUnauthorized creates a StandAloneVMDisksCreateUnauthorized with default headers values
func NewStandAloneVMDisksCreateUnauthorized() *StandAloneVMDisksCreateUnauthorized {
	return &StandAloneVMDisksCreateUnauthorized{}
}

/*
StandAloneVMDisksCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks create unauthorized response has a 2xx status code
func (o *StandAloneVMDisksCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks create unauthorized response has a 3xx status code
func (o *StandAloneVMDisksCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create unauthorized response has a 4xx status code
func (o *StandAloneVMDisksCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks create unauthorized response has a 5xx status code
func (o *StandAloneVMDisksCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks create unauthorized response a status code equal to that given
func (o *StandAloneVMDisksCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone Vm disks create unauthorized response
func (o *StandAloneVMDisksCreateUnauthorized) Code() int {
	return 401
}

func (o *StandAloneVMDisksCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateForbidden creates a StandAloneVMDisksCreateForbidden with default headers values
func NewStandAloneVMDisksCreateForbidden() *StandAloneVMDisksCreateForbidden {
	return &StandAloneVMDisksCreateForbidden{}
}

/*
StandAloneVMDisksCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks create forbidden response has a 2xx status code
func (o *StandAloneVMDisksCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks create forbidden response has a 3xx status code
func (o *StandAloneVMDisksCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create forbidden response has a 4xx status code
func (o *StandAloneVMDisksCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks create forbidden response has a 5xx status code
func (o *StandAloneVMDisksCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks create forbidden response a status code equal to that given
func (o *StandAloneVMDisksCreateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone Vm disks create forbidden response
func (o *StandAloneVMDisksCreateForbidden) Code() int {
	return 403
}

func (o *StandAloneVMDisksCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateNotFound creates a StandAloneVMDisksCreateNotFound with default headers values
func NewStandAloneVMDisksCreateNotFound() *StandAloneVMDisksCreateNotFound {
	return &StandAloneVMDisksCreateNotFound{}
}

/*
StandAloneVMDisksCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks create not found response has a 2xx status code
func (o *StandAloneVMDisksCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks create not found response has a 3xx status code
func (o *StandAloneVMDisksCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create not found response has a 4xx status code
func (o *StandAloneVMDisksCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks create not found response has a 5xx status code
func (o *StandAloneVMDisksCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks create not found response a status code equal to that given
func (o *StandAloneVMDisksCreateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone Vm disks create not found response
func (o *StandAloneVMDisksCreateNotFound) Code() int {
	return 404
}

func (o *StandAloneVMDisksCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksCreateInternalServerError creates a StandAloneVMDisksCreateInternalServerError with default headers values
func NewStandAloneVMDisksCreateInternalServerError() *StandAloneVMDisksCreateInternalServerError {
	return &StandAloneVMDisksCreateInternalServerError{}
}

/*
StandAloneVMDisksCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksCreateInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks create internal server error response has a 2xx status code
func (o *StandAloneVMDisksCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks create internal server error response has a 3xx status code
func (o *StandAloneVMDisksCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks create internal server error response has a 4xx status code
func (o *StandAloneVMDisksCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks create internal server error response has a 5xx status code
func (o *StandAloneVMDisksCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks create internal server error response a status code equal to that given
func (o *StandAloneVMDisksCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone Vm disks create internal server error response
func (o *StandAloneVMDisksCreateInternalServerError) Code() int {
	return 500
}

func (o *StandAloneVMDisksCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateInternalServerError ", 500)
}

func (o *StandAloneVMDisksCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/create][%d] standAloneVmDisksCreateInternalServerError ", 500)
}

func (o *StandAloneVMDisksCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
