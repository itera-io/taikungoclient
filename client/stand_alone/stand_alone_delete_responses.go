// Code generated by go-swagger; DO NOT EDIT.

package stand_alone

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneDeleteReader is a Reader for the StandAloneDelete structure.
type StandAloneDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneDeleteOK creates a StandAloneDeleteOK with default headers values
func NewStandAloneDeleteOK() *StandAloneDeleteOK {
	return &StandAloneDeleteOK{}
}

/*
StandAloneDeleteOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone delete o k response has a 2xx status code
func (o *StandAloneDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone delete o k response has a 3xx status code
func (o *StandAloneDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete o k response has a 4xx status code
func (o *StandAloneDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone delete o k response has a 5xx status code
func (o *StandAloneDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone delete o k response a status code equal to that given
func (o *StandAloneDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stand alone delete o k response
func (o *StandAloneDeleteOK) Code() int {
	return 200
}

func (o *StandAloneDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteOK  %+v", 200, o.Payload)
}

func (o *StandAloneDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteOK  %+v", 200, o.Payload)
}

func (o *StandAloneDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDeleteBadRequest creates a StandAloneDeleteBadRequest with default headers values
func NewStandAloneDeleteBadRequest() *StandAloneDeleteBadRequest {
	return &StandAloneDeleteBadRequest{}
}

/*
StandAloneDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone delete bad request response has a 2xx status code
func (o *StandAloneDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone delete bad request response has a 3xx status code
func (o *StandAloneDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete bad request response has a 4xx status code
func (o *StandAloneDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone delete bad request response has a 5xx status code
func (o *StandAloneDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone delete bad request response a status code equal to that given
func (o *StandAloneDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stand alone delete bad request response
func (o *StandAloneDeleteBadRequest) Code() int {
	return 400
}

func (o *StandAloneDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDeleteUnauthorized creates a StandAloneDeleteUnauthorized with default headers values
func NewStandAloneDeleteUnauthorized() *StandAloneDeleteUnauthorized {
	return &StandAloneDeleteUnauthorized{}
}

/*
StandAloneDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone delete unauthorized response has a 2xx status code
func (o *StandAloneDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone delete unauthorized response has a 3xx status code
func (o *StandAloneDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete unauthorized response has a 4xx status code
func (o *StandAloneDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone delete unauthorized response has a 5xx status code
func (o *StandAloneDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone delete unauthorized response a status code equal to that given
func (o *StandAloneDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stand alone delete unauthorized response
func (o *StandAloneDeleteUnauthorized) Code() int {
	return 401
}

func (o *StandAloneDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDeleteForbidden creates a StandAloneDeleteForbidden with default headers values
func NewStandAloneDeleteForbidden() *StandAloneDeleteForbidden {
	return &StandAloneDeleteForbidden{}
}

/*
StandAloneDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone delete forbidden response has a 2xx status code
func (o *StandAloneDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone delete forbidden response has a 3xx status code
func (o *StandAloneDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete forbidden response has a 4xx status code
func (o *StandAloneDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone delete forbidden response has a 5xx status code
func (o *StandAloneDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone delete forbidden response a status code equal to that given
func (o *StandAloneDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stand alone delete forbidden response
func (o *StandAloneDeleteForbidden) Code() int {
	return 403
}

func (o *StandAloneDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDeleteNotFound creates a StandAloneDeleteNotFound with default headers values
func NewStandAloneDeleteNotFound() *StandAloneDeleteNotFound {
	return &StandAloneDeleteNotFound{}
}

/*
StandAloneDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone delete not found response has a 2xx status code
func (o *StandAloneDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone delete not found response has a 3xx status code
func (o *StandAloneDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete not found response has a 4xx status code
func (o *StandAloneDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone delete not found response has a 5xx status code
func (o *StandAloneDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone delete not found response a status code equal to that given
func (o *StandAloneDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stand alone delete not found response
func (o *StandAloneDeleteNotFound) Code() int {
	return 404
}

func (o *StandAloneDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneDeleteInternalServerError creates a StandAloneDeleteInternalServerError with default headers values
func NewStandAloneDeleteInternalServerError() *StandAloneDeleteInternalServerError {
	return &StandAloneDeleteInternalServerError{}
}

/*
StandAloneDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneDeleteInternalServerError struct {
}

// IsSuccess returns true when this stand alone delete internal server error response has a 2xx status code
func (o *StandAloneDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone delete internal server error response has a 3xx status code
func (o *StandAloneDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone delete internal server error response has a 4xx status code
func (o *StandAloneDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone delete internal server error response has a 5xx status code
func (o *StandAloneDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone delete internal server error response a status code equal to that given
func (o *StandAloneDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stand alone delete internal server error response
func (o *StandAloneDeleteInternalServerError) Code() int {
	return 500
}

func (o *StandAloneDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteInternalServerError ", 500)
}

func (o *StandAloneDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAlone/delete][%d] standAloneDeleteInternalServerError ", 500)
}

func (o *StandAloneDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
