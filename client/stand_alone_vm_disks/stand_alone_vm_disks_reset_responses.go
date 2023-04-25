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

// StandAloneVMDisksResetReader is a Reader for the StandAloneVMDisksReset structure.
type StandAloneVMDisksResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksResetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksResetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksResetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksResetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksResetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksResetOK creates a StandAloneVMDisksResetOK with default headers values
func NewStandAloneVMDisksResetOK() *StandAloneVMDisksResetOK {
	return &StandAloneVMDisksResetOK{}
}

/*
StandAloneVMDisksResetOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksResetOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone Vm disks reset o k response has a 2xx status code
func (o *StandAloneVMDisksResetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks reset o k response has a 3xx status code
func (o *StandAloneVMDisksResetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset o k response has a 4xx status code
func (o *StandAloneVMDisksResetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks reset o k response has a 5xx status code
func (o *StandAloneVMDisksResetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks reset o k response a status code equal to that given
func (o *StandAloneVMDisksResetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone Vm disks reset o k response
func (o *StandAloneVMDisksResetOK) Code() int {
	return 200
}

func (o *StandAloneVMDisksResetOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksResetOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksResetOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneVMDisksResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksResetBadRequest creates a StandAloneVMDisksResetBadRequest with default headers values
func NewStandAloneVMDisksResetBadRequest() *StandAloneVMDisksResetBadRequest {
	return &StandAloneVMDisksResetBadRequest{}
}

/*
StandAloneVMDisksResetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksResetBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks reset bad request response has a 2xx status code
func (o *StandAloneVMDisksResetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks reset bad request response has a 3xx status code
func (o *StandAloneVMDisksResetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset bad request response has a 4xx status code
func (o *StandAloneVMDisksResetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks reset bad request response has a 5xx status code
func (o *StandAloneVMDisksResetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks reset bad request response a status code equal to that given
func (o *StandAloneVMDisksResetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone Vm disks reset bad request response
func (o *StandAloneVMDisksResetBadRequest) Code() int {
	return 400
}

func (o *StandAloneVMDisksResetBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksResetBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksResetBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksResetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksResetUnauthorized creates a StandAloneVMDisksResetUnauthorized with default headers values
func NewStandAloneVMDisksResetUnauthorized() *StandAloneVMDisksResetUnauthorized {
	return &StandAloneVMDisksResetUnauthorized{}
}

/*
StandAloneVMDisksResetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksResetUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks reset unauthorized response has a 2xx status code
func (o *StandAloneVMDisksResetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks reset unauthorized response has a 3xx status code
func (o *StandAloneVMDisksResetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset unauthorized response has a 4xx status code
func (o *StandAloneVMDisksResetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks reset unauthorized response has a 5xx status code
func (o *StandAloneVMDisksResetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks reset unauthorized response a status code equal to that given
func (o *StandAloneVMDisksResetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone Vm disks reset unauthorized response
func (o *StandAloneVMDisksResetUnauthorized) Code() int {
	return 401
}

func (o *StandAloneVMDisksResetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksResetUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksResetUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksResetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksResetForbidden creates a StandAloneVMDisksResetForbidden with default headers values
func NewStandAloneVMDisksResetForbidden() *StandAloneVMDisksResetForbidden {
	return &StandAloneVMDisksResetForbidden{}
}

/*
StandAloneVMDisksResetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksResetForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks reset forbidden response has a 2xx status code
func (o *StandAloneVMDisksResetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks reset forbidden response has a 3xx status code
func (o *StandAloneVMDisksResetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset forbidden response has a 4xx status code
func (o *StandAloneVMDisksResetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks reset forbidden response has a 5xx status code
func (o *StandAloneVMDisksResetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks reset forbidden response a status code equal to that given
func (o *StandAloneVMDisksResetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone Vm disks reset forbidden response
func (o *StandAloneVMDisksResetForbidden) Code() int {
	return 403
}

func (o *StandAloneVMDisksResetForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksResetForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksResetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksResetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksResetNotFound creates a StandAloneVMDisksResetNotFound with default headers values
func NewStandAloneVMDisksResetNotFound() *StandAloneVMDisksResetNotFound {
	return &StandAloneVMDisksResetNotFound{}
}

/*
StandAloneVMDisksResetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksResetNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone Vm disks reset not found response has a 2xx status code
func (o *StandAloneVMDisksResetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks reset not found response has a 3xx status code
func (o *StandAloneVMDisksResetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset not found response has a 4xx status code
func (o *StandAloneVMDisksResetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks reset not found response has a 5xx status code
func (o *StandAloneVMDisksResetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks reset not found response a status code equal to that given
func (o *StandAloneVMDisksResetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone Vm disks reset not found response
func (o *StandAloneVMDisksResetNotFound) Code() int {
	return 404
}

func (o *StandAloneVMDisksResetNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksResetNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksResetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneVMDisksResetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksResetInternalServerError creates a StandAloneVMDisksResetInternalServerError with default headers values
func NewStandAloneVMDisksResetInternalServerError() *StandAloneVMDisksResetInternalServerError {
	return &StandAloneVMDisksResetInternalServerError{}
}

/*
StandAloneVMDisksResetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksResetInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks reset internal server error response has a 2xx status code
func (o *StandAloneVMDisksResetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks reset internal server error response has a 3xx status code
func (o *StandAloneVMDisksResetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks reset internal server error response has a 4xx status code
func (o *StandAloneVMDisksResetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks reset internal server error response has a 5xx status code
func (o *StandAloneVMDisksResetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks reset internal server error response a status code equal to that given
func (o *StandAloneVMDisksResetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone Vm disks reset internal server error response
func (o *StandAloneVMDisksResetInternalServerError) Code() int {
	return 500
}

func (o *StandAloneVMDisksResetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetInternalServerError ", 500)
}

func (o *StandAloneVMDisksResetInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/reset][%d] standAloneVmDisksResetInternalServerError ", 500)
}

func (o *StandAloneVMDisksResetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
