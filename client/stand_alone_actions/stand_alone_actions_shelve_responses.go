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

// StandAloneActionsShelveReader is a Reader for the StandAloneActionsShelve structure.
type StandAloneActionsShelveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsShelveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsShelveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsShelveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsShelveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsShelveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsShelveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsShelveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsShelveOK creates a StandAloneActionsShelveOK with default headers values
func NewStandAloneActionsShelveOK() *StandAloneActionsShelveOK {
	return &StandAloneActionsShelveOK{}
}

/*
StandAloneActionsShelveOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsShelveOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone actions shelve o k response has a 2xx status code
func (o *StandAloneActionsShelveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions shelve o k response has a 3xx status code
func (o *StandAloneActionsShelveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve o k response has a 4xx status code
func (o *StandAloneActionsShelveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions shelve o k response has a 5xx status code
func (o *StandAloneActionsShelveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions shelve o k response a status code equal to that given
func (o *StandAloneActionsShelveOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneActionsShelveOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsShelveOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsShelveOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneActionsShelveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShelveBadRequest creates a StandAloneActionsShelveBadRequest with default headers values
func NewStandAloneActionsShelveBadRequest() *StandAloneActionsShelveBadRequest {
	return &StandAloneActionsShelveBadRequest{}
}

/*
StandAloneActionsShelveBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsShelveBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this stand alone actions shelve bad request response has a 2xx status code
func (o *StandAloneActionsShelveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions shelve bad request response has a 3xx status code
func (o *StandAloneActionsShelveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve bad request response has a 4xx status code
func (o *StandAloneActionsShelveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions shelve bad request response has a 5xx status code
func (o *StandAloneActionsShelveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions shelve bad request response a status code equal to that given
func (o *StandAloneActionsShelveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneActionsShelveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsShelveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsShelveBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *StandAloneActionsShelveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShelveUnauthorized creates a StandAloneActionsShelveUnauthorized with default headers values
func NewStandAloneActionsShelveUnauthorized() *StandAloneActionsShelveUnauthorized {
	return &StandAloneActionsShelveUnauthorized{}
}

/*
StandAloneActionsShelveUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsShelveUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions shelve unauthorized response has a 2xx status code
func (o *StandAloneActionsShelveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions shelve unauthorized response has a 3xx status code
func (o *StandAloneActionsShelveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve unauthorized response has a 4xx status code
func (o *StandAloneActionsShelveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions shelve unauthorized response has a 5xx status code
func (o *StandAloneActionsShelveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions shelve unauthorized response a status code equal to that given
func (o *StandAloneActionsShelveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneActionsShelveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsShelveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsShelveUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsShelveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShelveForbidden creates a StandAloneActionsShelveForbidden with default headers values
func NewStandAloneActionsShelveForbidden() *StandAloneActionsShelveForbidden {
	return &StandAloneActionsShelveForbidden{}
}

/*
StandAloneActionsShelveForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsShelveForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions shelve forbidden response has a 2xx status code
func (o *StandAloneActionsShelveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions shelve forbidden response has a 3xx status code
func (o *StandAloneActionsShelveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve forbidden response has a 4xx status code
func (o *StandAloneActionsShelveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions shelve forbidden response has a 5xx status code
func (o *StandAloneActionsShelveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions shelve forbidden response a status code equal to that given
func (o *StandAloneActionsShelveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneActionsShelveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsShelveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsShelveForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsShelveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShelveNotFound creates a StandAloneActionsShelveNotFound with default headers values
func NewStandAloneActionsShelveNotFound() *StandAloneActionsShelveNotFound {
	return &StandAloneActionsShelveNotFound{}
}

/*
StandAloneActionsShelveNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsShelveNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this stand alone actions shelve not found response has a 2xx status code
func (o *StandAloneActionsShelveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions shelve not found response has a 3xx status code
func (o *StandAloneActionsShelveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve not found response has a 4xx status code
func (o *StandAloneActionsShelveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions shelve not found response has a 5xx status code
func (o *StandAloneActionsShelveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions shelve not found response a status code equal to that given
func (o *StandAloneActionsShelveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneActionsShelveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsShelveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsShelveNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *StandAloneActionsShelveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsShelveInternalServerError creates a StandAloneActionsShelveInternalServerError with default headers values
func NewStandAloneActionsShelveInternalServerError() *StandAloneActionsShelveInternalServerError {
	return &StandAloneActionsShelveInternalServerError{}
}

/*
StandAloneActionsShelveInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsShelveInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions shelve internal server error response has a 2xx status code
func (o *StandAloneActionsShelveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions shelve internal server error response has a 3xx status code
func (o *StandAloneActionsShelveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions shelve internal server error response has a 4xx status code
func (o *StandAloneActionsShelveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions shelve internal server error response has a 5xx status code
func (o *StandAloneActionsShelveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions shelve internal server error response a status code equal to that given
func (o *StandAloneActionsShelveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneActionsShelveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveInternalServerError ", 500)
}

func (o *StandAloneActionsShelveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/shelve][%d] standAloneActionsShelveInternalServerError ", 500)
}

func (o *StandAloneActionsShelveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
