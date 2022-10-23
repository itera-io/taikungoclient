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

// StandAloneActionsRebootReader is a Reader for the StandAloneActionsReboot structure.
type StandAloneActionsRebootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsRebootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsRebootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsRebootBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsRebootUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsRebootForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsRebootNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsRebootInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsRebootOK creates a StandAloneActionsRebootOK with default headers values
func NewStandAloneActionsRebootOK() *StandAloneActionsRebootOK {
	return &StandAloneActionsRebootOK{}
}

/*
StandAloneActionsRebootOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsRebootOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone actions reboot o k response has a 2xx status code
func (o *StandAloneActionsRebootOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions reboot o k response has a 3xx status code
func (o *StandAloneActionsRebootOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot o k response has a 4xx status code
func (o *StandAloneActionsRebootOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions reboot o k response has a 5xx status code
func (o *StandAloneActionsRebootOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions reboot o k response a status code equal to that given
func (o *StandAloneActionsRebootOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneActionsRebootOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsRebootOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsRebootOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneActionsRebootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsRebootBadRequest creates a StandAloneActionsRebootBadRequest with default headers values
func NewStandAloneActionsRebootBadRequest() *StandAloneActionsRebootBadRequest {
	return &StandAloneActionsRebootBadRequest{}
}

/*
StandAloneActionsRebootBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsRebootBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone actions reboot bad request response has a 2xx status code
func (o *StandAloneActionsRebootBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions reboot bad request response has a 3xx status code
func (o *StandAloneActionsRebootBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot bad request response has a 4xx status code
func (o *StandAloneActionsRebootBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions reboot bad request response has a 5xx status code
func (o *StandAloneActionsRebootBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions reboot bad request response a status code equal to that given
func (o *StandAloneActionsRebootBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneActionsRebootBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsRebootBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsRebootBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneActionsRebootBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsRebootUnauthorized creates a StandAloneActionsRebootUnauthorized with default headers values
func NewStandAloneActionsRebootUnauthorized() *StandAloneActionsRebootUnauthorized {
	return &StandAloneActionsRebootUnauthorized{}
}

/*
StandAloneActionsRebootUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsRebootUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions reboot unauthorized response has a 2xx status code
func (o *StandAloneActionsRebootUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions reboot unauthorized response has a 3xx status code
func (o *StandAloneActionsRebootUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot unauthorized response has a 4xx status code
func (o *StandAloneActionsRebootUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions reboot unauthorized response has a 5xx status code
func (o *StandAloneActionsRebootUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions reboot unauthorized response a status code equal to that given
func (o *StandAloneActionsRebootUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneActionsRebootUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsRebootUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsRebootUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsRebootUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsRebootForbidden creates a StandAloneActionsRebootForbidden with default headers values
func NewStandAloneActionsRebootForbidden() *StandAloneActionsRebootForbidden {
	return &StandAloneActionsRebootForbidden{}
}

/*
StandAloneActionsRebootForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsRebootForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions reboot forbidden response has a 2xx status code
func (o *StandAloneActionsRebootForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions reboot forbidden response has a 3xx status code
func (o *StandAloneActionsRebootForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot forbidden response has a 4xx status code
func (o *StandAloneActionsRebootForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions reboot forbidden response has a 5xx status code
func (o *StandAloneActionsRebootForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions reboot forbidden response a status code equal to that given
func (o *StandAloneActionsRebootForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneActionsRebootForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsRebootForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsRebootForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsRebootForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsRebootNotFound creates a StandAloneActionsRebootNotFound with default headers values
func NewStandAloneActionsRebootNotFound() *StandAloneActionsRebootNotFound {
	return &StandAloneActionsRebootNotFound{}
}

/*
StandAloneActionsRebootNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsRebootNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions reboot not found response has a 2xx status code
func (o *StandAloneActionsRebootNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions reboot not found response has a 3xx status code
func (o *StandAloneActionsRebootNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot not found response has a 4xx status code
func (o *StandAloneActionsRebootNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions reboot not found response has a 5xx status code
func (o *StandAloneActionsRebootNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions reboot not found response a status code equal to that given
func (o *StandAloneActionsRebootNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneActionsRebootNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsRebootNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsRebootNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsRebootNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsRebootInternalServerError creates a StandAloneActionsRebootInternalServerError with default headers values
func NewStandAloneActionsRebootInternalServerError() *StandAloneActionsRebootInternalServerError {
	return &StandAloneActionsRebootInternalServerError{}
}

/*
StandAloneActionsRebootInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsRebootInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions reboot internal server error response has a 2xx status code
func (o *StandAloneActionsRebootInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions reboot internal server error response has a 3xx status code
func (o *StandAloneActionsRebootInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions reboot internal server error response has a 4xx status code
func (o *StandAloneActionsRebootInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions reboot internal server error response has a 5xx status code
func (o *StandAloneActionsRebootInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions reboot internal server error response a status code equal to that given
func (o *StandAloneActionsRebootInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneActionsRebootInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootInternalServerError ", 500)
}

func (o *StandAloneActionsRebootInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/reboot][%d] standAloneActionsRebootInternalServerError ", 500)
}

func (o *StandAloneActionsRebootInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
