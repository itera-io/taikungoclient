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

// StandAloneVMDisksUpdateDiskSizeReader is a Reader for the StandAloneVMDisksUpdateDiskSize structure.
type StandAloneVMDisksUpdateDiskSizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneVMDisksUpdateDiskSizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneVMDisksUpdateDiskSizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneVMDisksUpdateDiskSizeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneVMDisksUpdateDiskSizeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneVMDisksUpdateDiskSizeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneVMDisksUpdateDiskSizeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneVMDisksUpdateDiskSizeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneVMDisksUpdateDiskSizeOK creates a StandAloneVMDisksUpdateDiskSizeOK with default headers values
func NewStandAloneVMDisksUpdateDiskSizeOK() *StandAloneVMDisksUpdateDiskSizeOK {
	return &StandAloneVMDisksUpdateDiskSizeOK{}
}

/*
StandAloneVMDisksUpdateDiskSizeOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneVMDisksUpdateDiskSizeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone Vm disks update disk size o k response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone Vm disks update disk size o k response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size o k response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update disk size o k response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size o k response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeOK  %+v", 200, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeBadRequest creates a StandAloneVMDisksUpdateDiskSizeBadRequest with default headers values
func NewStandAloneVMDisksUpdateDiskSizeBadRequest() *StandAloneVMDisksUpdateDiskSizeBadRequest {
	return &StandAloneVMDisksUpdateDiskSizeBadRequest{}
}

/*
StandAloneVMDisksUpdateDiskSizeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneVMDisksUpdateDiskSizeBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone Vm disks update disk size bad request response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size bad request response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size bad request response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size bad request response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size bad request response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeUnauthorized creates a StandAloneVMDisksUpdateDiskSizeUnauthorized with default headers values
func NewStandAloneVMDisksUpdateDiskSizeUnauthorized() *StandAloneVMDisksUpdateDiskSizeUnauthorized {
	return &StandAloneVMDisksUpdateDiskSizeUnauthorized{}
}

/*
StandAloneVMDisksUpdateDiskSizeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneVMDisksUpdateDiskSizeUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Vm disks update disk size unauthorized response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size unauthorized response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size unauthorized response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size unauthorized response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size unauthorized response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeForbidden creates a StandAloneVMDisksUpdateDiskSizeForbidden with default headers values
func NewStandAloneVMDisksUpdateDiskSizeForbidden() *StandAloneVMDisksUpdateDiskSizeForbidden {
	return &StandAloneVMDisksUpdateDiskSizeForbidden{}
}

/*
StandAloneVMDisksUpdateDiskSizeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneVMDisksUpdateDiskSizeForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Vm disks update disk size forbidden response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size forbidden response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size forbidden response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size forbidden response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size forbidden response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeNotFound creates a StandAloneVMDisksUpdateDiskSizeNotFound with default headers values
func NewStandAloneVMDisksUpdateDiskSizeNotFound() *StandAloneVMDisksUpdateDiskSizeNotFound {
	return &StandAloneVMDisksUpdateDiskSizeNotFound{}
}

/*
StandAloneVMDisksUpdateDiskSizeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneVMDisksUpdateDiskSizeNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone Vm disks update disk size not found response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size not found response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size not found response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone Vm disks update disk size not found response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone Vm disks update disk size not found response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneVMDisksUpdateDiskSizeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneVMDisksUpdateDiskSizeInternalServerError creates a StandAloneVMDisksUpdateDiskSizeInternalServerError with default headers values
func NewStandAloneVMDisksUpdateDiskSizeInternalServerError() *StandAloneVMDisksUpdateDiskSizeInternalServerError {
	return &StandAloneVMDisksUpdateDiskSizeInternalServerError{}
}

/*
StandAloneVMDisksUpdateDiskSizeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneVMDisksUpdateDiskSizeInternalServerError struct {
}

// IsSuccess returns true when this stand alone Vm disks update disk size internal server error response has a 2xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone Vm disks update disk size internal server error response has a 3xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone Vm disks update disk size internal server error response has a 4xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone Vm disks update disk size internal server error response has a 5xx status code
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone Vm disks update disk size internal server error response a status code equal to that given
func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneVmDisks/update-size][%d] standAloneVmDisksUpdateDiskSizeInternalServerError ", 500)
}

func (o *StandAloneVMDisksUpdateDiskSizeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
