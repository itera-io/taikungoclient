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

// StandAloneActionsStartReader is a Reader for the StandAloneActionsStart structure.
type StandAloneActionsStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneActionsStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneActionsStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneActionsStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneActionsStartUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneActionsStartForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneActionsStartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneActionsStartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneActionsStartOK creates a StandAloneActionsStartOK with default headers values
func NewStandAloneActionsStartOK() *StandAloneActionsStartOK {
	return &StandAloneActionsStartOK{}
}

/*
StandAloneActionsStartOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneActionsStartOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone actions start o k response has a 2xx status code
func (o *StandAloneActionsStartOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone actions start o k response has a 3xx status code
func (o *StandAloneActionsStartOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start o k response has a 4xx status code
func (o *StandAloneActionsStartOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions start o k response has a 5xx status code
func (o *StandAloneActionsStartOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions start o k response a status code equal to that given
func (o *StandAloneActionsStartOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone actions start o k response
func (o *StandAloneActionsStartOK) Code() int {
	return 200
}

func (o *StandAloneActionsStartOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsStartOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartOK  %+v", 200, o.Payload)
}

func (o *StandAloneActionsStartOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneActionsStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStartBadRequest creates a StandAloneActionsStartBadRequest with default headers values
func NewStandAloneActionsStartBadRequest() *StandAloneActionsStartBadRequest {
	return &StandAloneActionsStartBadRequest{}
}

/*
StandAloneActionsStartBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneActionsStartBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions start bad request response has a 2xx status code
func (o *StandAloneActionsStartBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions start bad request response has a 3xx status code
func (o *StandAloneActionsStartBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start bad request response has a 4xx status code
func (o *StandAloneActionsStartBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions start bad request response has a 5xx status code
func (o *StandAloneActionsStartBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions start bad request response a status code equal to that given
func (o *StandAloneActionsStartBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone actions start bad request response
func (o *StandAloneActionsStartBadRequest) Code() int {
	return 400
}

func (o *StandAloneActionsStartBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsStartBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneActionsStartBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsStartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStartUnauthorized creates a StandAloneActionsStartUnauthorized with default headers values
func NewStandAloneActionsStartUnauthorized() *StandAloneActionsStartUnauthorized {
	return &StandAloneActionsStartUnauthorized{}
}

/*
StandAloneActionsStartUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneActionsStartUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions start unauthorized response has a 2xx status code
func (o *StandAloneActionsStartUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions start unauthorized response has a 3xx status code
func (o *StandAloneActionsStartUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start unauthorized response has a 4xx status code
func (o *StandAloneActionsStartUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions start unauthorized response has a 5xx status code
func (o *StandAloneActionsStartUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions start unauthorized response a status code equal to that given
func (o *StandAloneActionsStartUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone actions start unauthorized response
func (o *StandAloneActionsStartUnauthorized) Code() int {
	return 401
}

func (o *StandAloneActionsStartUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsStartUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneActionsStartUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsStartUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStartForbidden creates a StandAloneActionsStartForbidden with default headers values
func NewStandAloneActionsStartForbidden() *StandAloneActionsStartForbidden {
	return &StandAloneActionsStartForbidden{}
}

/*
StandAloneActionsStartForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneActionsStartForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions start forbidden response has a 2xx status code
func (o *StandAloneActionsStartForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions start forbidden response has a 3xx status code
func (o *StandAloneActionsStartForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start forbidden response has a 4xx status code
func (o *StandAloneActionsStartForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions start forbidden response has a 5xx status code
func (o *StandAloneActionsStartForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions start forbidden response a status code equal to that given
func (o *StandAloneActionsStartForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone actions start forbidden response
func (o *StandAloneActionsStartForbidden) Code() int {
	return 403
}

func (o *StandAloneActionsStartForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsStartForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneActionsStartForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsStartForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStartNotFound creates a StandAloneActionsStartNotFound with default headers values
func NewStandAloneActionsStartNotFound() *StandAloneActionsStartNotFound {
	return &StandAloneActionsStartNotFound{}
}

/*
StandAloneActionsStartNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneActionsStartNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone actions start not found response has a 2xx status code
func (o *StandAloneActionsStartNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions start not found response has a 3xx status code
func (o *StandAloneActionsStartNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start not found response has a 4xx status code
func (o *StandAloneActionsStartNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone actions start not found response has a 5xx status code
func (o *StandAloneActionsStartNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone actions start not found response a status code equal to that given
func (o *StandAloneActionsStartNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone actions start not found response
func (o *StandAloneActionsStartNotFound) Code() int {
	return 404
}

func (o *StandAloneActionsStartNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsStartNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneActionsStartNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneActionsStartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneActionsStartInternalServerError creates a StandAloneActionsStartInternalServerError with default headers values
func NewStandAloneActionsStartInternalServerError() *StandAloneActionsStartInternalServerError {
	return &StandAloneActionsStartInternalServerError{}
}

/*
StandAloneActionsStartInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneActionsStartInternalServerError struct {
}

// IsSuccess returns true when this stand alone actions start internal server error response has a 2xx status code
func (o *StandAloneActionsStartInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone actions start internal server error response has a 3xx status code
func (o *StandAloneActionsStartInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone actions start internal server error response has a 4xx status code
func (o *StandAloneActionsStartInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone actions start internal server error response has a 5xx status code
func (o *StandAloneActionsStartInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone actions start internal server error response a status code equal to that given
func (o *StandAloneActionsStartInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone actions start internal server error response
func (o *StandAloneActionsStartInternalServerError) Code() int {
	return 500
}

func (o *StandAloneActionsStartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartInternalServerError ", 500)
}

func (o *StandAloneActionsStartInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneActions/start][%d] standAloneActionsStartInternalServerError ", 500)
}

func (o *StandAloneActionsStartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
