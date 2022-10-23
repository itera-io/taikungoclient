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

// StandAloneActionsUnshelveReader is a Reader for the StandAloneActionsUnshelve structure.
type StandAloneActionsUnshelveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsUnshelveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsUnshelveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsUnshelveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsUnshelveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsUnshelveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsUnshelveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsUnshelveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsUnshelveOK creates a StandAloneActionsUnshelveOK with default headers values
func NewStandAloneActionsUnshelveOK() *StandAloneActionsUnshelveOK {
	return &StandAloneActionsUnshelveOK{}
}

/*
StandAloneActionsUnshelveOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsUnshelveOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone actions unshelve o k response has a 2xx status code
func (o *StandAloneActionsUnshelveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions unshelve o k response has a 3xx status code
func (o *StandAloneActionsUnshelveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve o k response has a 4xx status code
func (o *StandAloneActionsUnshelveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions unshelve o k response has a 5xx status code
func (o *StandAloneActionsUnshelveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions unshelve o k response a status code equal to that given
func (o *StandAloneActionsUnshelveOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneActionsUnshelveOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsUnshelveOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsUnshelveOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneActionsUnshelveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsUnshelveBadRequest creates a StandAloneActionsUnshelveBadRequest with default headers values
func NewStandAloneActionsUnshelveBadRequest() *StandAloneActionsUnshelveBadRequest {
	return &StandAloneActionsUnshelveBadRequest{}
}

/*
StandAloneActionsUnshelveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsUnshelveBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone actions unshelve bad request response has a 2xx status code
func (o *StandAloneActionsUnshelveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions unshelve bad request response has a 3xx status code
func (o *StandAloneActionsUnshelveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve bad request response has a 4xx status code
func (o *StandAloneActionsUnshelveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions unshelve bad request response has a 5xx status code
func (o *StandAloneActionsUnshelveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions unshelve bad request response a status code equal to that given
func (o *StandAloneActionsUnshelveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneActionsUnshelveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsUnshelveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsUnshelveBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneActionsUnshelveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsUnshelveUnauthorized creates a StandAloneActionsUnshelveUnauthorized with default headers values
func NewStandAloneActionsUnshelveUnauthorized() *StandAloneActionsUnshelveUnauthorized {
	return &StandAloneActionsUnshelveUnauthorized{}
}

/*
StandAloneActionsUnshelveUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsUnshelveUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions unshelve unauthorized response has a 2xx status code
func (o *StandAloneActionsUnshelveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions unshelve unauthorized response has a 3xx status code
func (o *StandAloneActionsUnshelveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve unauthorized response has a 4xx status code
func (o *StandAloneActionsUnshelveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions unshelve unauthorized response has a 5xx status code
func (o *StandAloneActionsUnshelveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions unshelve unauthorized response a status code equal to that given
func (o *StandAloneActionsUnshelveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneActionsUnshelveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsUnshelveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsUnshelveUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsUnshelveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsUnshelveForbidden creates a StandAloneActionsUnshelveForbidden with default headers values
func NewStandAloneActionsUnshelveForbidden() *StandAloneActionsUnshelveForbidden {
	return &StandAloneActionsUnshelveForbidden{}
}

/*
StandAloneActionsUnshelveForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsUnshelveForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions unshelve forbidden response has a 2xx status code
func (o *StandAloneActionsUnshelveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions unshelve forbidden response has a 3xx status code
func (o *StandAloneActionsUnshelveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve forbidden response has a 4xx status code
func (o *StandAloneActionsUnshelveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions unshelve forbidden response has a 5xx status code
func (o *StandAloneActionsUnshelveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions unshelve forbidden response a status code equal to that given
func (o *StandAloneActionsUnshelveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneActionsUnshelveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsUnshelveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsUnshelveForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsUnshelveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsUnshelveNotFound creates a StandAloneActionsUnshelveNotFound with default headers values
func NewStandAloneActionsUnshelveNotFound() *StandAloneActionsUnshelveNotFound {
	return &StandAloneActionsUnshelveNotFound{}
}

/*
StandAloneActionsUnshelveNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsUnshelveNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this stand alone actions unshelve not found response has a 2xx status code
func (o *StandAloneActionsUnshelveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions unshelve not found response has a 3xx status code
func (o *StandAloneActionsUnshelveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve not found response has a 4xx status code
func (o *StandAloneActionsUnshelveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions unshelve not found response has a 5xx status code
func (o *StandAloneActionsUnshelveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions unshelve not found response a status code equal to that given
func (o *StandAloneActionsUnshelveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneActionsUnshelveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsUnshelveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsUnshelveNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *StandAloneActionsUnshelveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsUnshelveInternalServerError creates a StandAloneActionsUnshelveInternalServerError with default headers values
func NewStandAloneActionsUnshelveInternalServerError() *StandAloneActionsUnshelveInternalServerError {
	return &StandAloneActionsUnshelveInternalServerError{}
}

/*
StandAloneActionsUnshelveInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsUnshelveInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions unshelve internal server error response has a 2xx status code
func (o *StandAloneActionsUnshelveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions unshelve internal server error response has a 3xx status code
func (o *StandAloneActionsUnshelveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions unshelve internal server error response has a 4xx status code
func (o *StandAloneActionsUnshelveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions unshelve internal server error response has a 5xx status code
func (o *StandAloneActionsUnshelveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions unshelve internal server error response a status code equal to that given
func (o *StandAloneActionsUnshelveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneActionsUnshelveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveInternalServerError ", 500)
}

func (o *StandAloneActionsUnshelveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/unshelve][%d] standAloneActionsUnshelveInternalServerError ", 500)
}

func (o *StandAloneActionsUnshelveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
