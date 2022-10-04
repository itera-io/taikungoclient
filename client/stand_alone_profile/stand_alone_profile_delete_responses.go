// Code generated by go-swagger; DO NOT EDIT.

package stand_alone_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// StandAloneProfileDeleteReader is a Reader for the StandAloneProfileDelete structure.
type StandAloneProfileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StandAloneProfileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStandAloneProfileDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStandAloneProfileDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStandAloneProfileDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStandAloneProfileDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStandAloneProfileDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStandAloneProfileDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStandAloneProfileDeleteOK creates a StandAloneProfileDeleteOK with default headers values
func NewStandAloneProfileDeleteOK() *StandAloneProfileDeleteOK {
	return &StandAloneProfileDeleteOK{}
}

/*
StandAloneProfileDeleteOK describes a response with status code 200, with default header values.

Success
*/
type StandAloneProfileDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this stand alone profile delete o k response has a 2xx status code
func (o *StandAloneProfileDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stand alone profile delete o k response has a 3xx status code
func (o *StandAloneProfileDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete o k response has a 4xx status code
func (o *StandAloneProfileDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile delete o k response has a 5xx status code
func (o *StandAloneProfileDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile delete o k response a status code equal to that given
func (o *StandAloneProfileDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *StandAloneProfileDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteOK  %+v", 200, o.Payload)
}

func (o *StandAloneProfileDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *StandAloneProfileDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDeleteBadRequest creates a StandAloneProfileDeleteBadRequest with default headers values
func NewStandAloneProfileDeleteBadRequest() *StandAloneProfileDeleteBadRequest {
	return &StandAloneProfileDeleteBadRequest{}
}

/*
StandAloneProfileDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type StandAloneProfileDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this stand alone profile delete bad request response has a 2xx status code
func (o *StandAloneProfileDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile delete bad request response has a 3xx status code
func (o *StandAloneProfileDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete bad request response has a 4xx status code
func (o *StandAloneProfileDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile delete bad request response has a 5xx status code
func (o *StandAloneProfileDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile delete bad request response a status code equal to that given
func (o *StandAloneProfileDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *StandAloneProfileDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *StandAloneProfileDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *StandAloneProfileDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDeleteUnauthorized creates a StandAloneProfileDeleteUnauthorized with default headers values
func NewStandAloneProfileDeleteUnauthorized() *StandAloneProfileDeleteUnauthorized {
	return &StandAloneProfileDeleteUnauthorized{}
}

/*
StandAloneProfileDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StandAloneProfileDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone profile delete unauthorized response has a 2xx status code
func (o *StandAloneProfileDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile delete unauthorized response has a 3xx status code
func (o *StandAloneProfileDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete unauthorized response has a 4xx status code
func (o *StandAloneProfileDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile delete unauthorized response has a 5xx status code
func (o *StandAloneProfileDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile delete unauthorized response a status code equal to that given
func (o *StandAloneProfileDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StandAloneProfileDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *StandAloneProfileDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneProfileDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDeleteForbidden creates a StandAloneProfileDeleteForbidden with default headers values
func NewStandAloneProfileDeleteForbidden() *StandAloneProfileDeleteForbidden {
	return &StandAloneProfileDeleteForbidden{}
}

/*
StandAloneProfileDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StandAloneProfileDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone profile delete forbidden response has a 2xx status code
func (o *StandAloneProfileDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile delete forbidden response has a 3xx status code
func (o *StandAloneProfileDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete forbidden response has a 4xx status code
func (o *StandAloneProfileDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile delete forbidden response has a 5xx status code
func (o *StandAloneProfileDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile delete forbidden response a status code equal to that given
func (o *StandAloneProfileDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StandAloneProfileDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteForbidden  %+v", 403, o.Payload)
}

func (o *StandAloneProfileDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneProfileDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDeleteNotFound creates a StandAloneProfileDeleteNotFound with default headers values
func NewStandAloneProfileDeleteNotFound() *StandAloneProfileDeleteNotFound {
	return &StandAloneProfileDeleteNotFound{}
}

/*
StandAloneProfileDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type StandAloneProfileDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this stand alone profile delete not found response has a 2xx status code
func (o *StandAloneProfileDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile delete not found response has a 3xx status code
func (o *StandAloneProfileDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete not found response has a 4xx status code
func (o *StandAloneProfileDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stand alone profile delete not found response has a 5xx status code
func (o *StandAloneProfileDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stand alone profile delete not found response a status code equal to that given
func (o *StandAloneProfileDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StandAloneProfileDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteNotFound  %+v", 404, o.Payload)
}

func (o *StandAloneProfileDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StandAloneProfileDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStandAloneProfileDeleteInternalServerError creates a StandAloneProfileDeleteInternalServerError with default headers values
func NewStandAloneProfileDeleteInternalServerError() *StandAloneProfileDeleteInternalServerError {
	return &StandAloneProfileDeleteInternalServerError{}
}

/*
StandAloneProfileDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type StandAloneProfileDeleteInternalServerError struct {
}

// IsSuccess returns true when this stand alone profile delete internal server error response has a 2xx status code
func (o *StandAloneProfileDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stand alone profile delete internal server error response has a 3xx status code
func (o *StandAloneProfileDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stand alone profile delete internal server error response has a 4xx status code
func (o *StandAloneProfileDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stand alone profile delete internal server error response has a 5xx status code
func (o *StandAloneProfileDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stand alone profile delete internal server error response a status code equal to that given
func (o *StandAloneProfileDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StandAloneProfileDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteInternalServerError ", 500)
}

func (o *StandAloneProfileDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/StandAloneProfile/delete][%d] standAloneProfileDeleteInternalServerError ", 500)
}

func (o *StandAloneProfileDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
