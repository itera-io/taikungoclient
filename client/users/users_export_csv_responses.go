// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UsersExportCsvReader is a Reader for the UsersExportCsv structure.
type UsersExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersExportCsvOK creates a UsersExportCsvOK with default headers values
func NewUsersExportCsvOK() *UsersExportCsvOK {
	return &UsersExportCsvOK{}
}

/*
UsersExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type UsersExportCsvOK struct {
}

// IsSuccess returns true when this users export csv o k response has a 2xx status code
func (o *UsersExportCsvOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users export csv o k response has a 3xx status code
func (o *UsersExportCsvOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv o k response has a 4xx status code
func (o *UsersExportCsvOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users export csv o k response has a 5xx status code
func (o *UsersExportCsvOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users export csv o k response a status code equal to that given
func (o *UsersExportCsvOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvOK ", 200)
}

func (o *UsersExportCsvOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvOK ", 200)
}

func (o *UsersExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersExportCsvBadRequest creates a UsersExportCsvBadRequest with default headers values
func NewUsersExportCsvBadRequest() *UsersExportCsvBadRequest {
	return &UsersExportCsvBadRequest{}
}

/*
UsersExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersExportCsvBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users export csv bad request response has a 2xx status code
func (o *UsersExportCsvBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users export csv bad request response has a 3xx status code
func (o *UsersExportCsvBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv bad request response has a 4xx status code
func (o *UsersExportCsvBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users export csv bad request response has a 5xx status code
func (o *UsersExportCsvBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users export csv bad request response a status code equal to that given
func (o *UsersExportCsvBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *UsersExportCsvBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *UsersExportCsvBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersExportCsvUnauthorized creates a UsersExportCsvUnauthorized with default headers values
func NewUsersExportCsvUnauthorized() *UsersExportCsvUnauthorized {
	return &UsersExportCsvUnauthorized{}
}

/*
UsersExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersExportCsvUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users export csv unauthorized response has a 2xx status code
func (o *UsersExportCsvUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users export csv unauthorized response has a 3xx status code
func (o *UsersExportCsvUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv unauthorized response has a 4xx status code
func (o *UsersExportCsvUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users export csv unauthorized response has a 5xx status code
func (o *UsersExportCsvUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users export csv unauthorized response a status code equal to that given
func (o *UsersExportCsvUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersExportCsvUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersExportCsvUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersExportCsvForbidden creates a UsersExportCsvForbidden with default headers values
func NewUsersExportCsvForbidden() *UsersExportCsvForbidden {
	return &UsersExportCsvForbidden{}
}

/*
UsersExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersExportCsvForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users export csv forbidden response has a 2xx status code
func (o *UsersExportCsvForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users export csv forbidden response has a 3xx status code
func (o *UsersExportCsvForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv forbidden response has a 4xx status code
func (o *UsersExportCsvForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users export csv forbidden response has a 5xx status code
func (o *UsersExportCsvForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users export csv forbidden response a status code equal to that given
func (o *UsersExportCsvForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *UsersExportCsvForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *UsersExportCsvForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersExportCsvNotFound creates a UsersExportCsvNotFound with default headers values
func NewUsersExportCsvNotFound() *UsersExportCsvNotFound {
	return &UsersExportCsvNotFound{}
}

/*
UsersExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersExportCsvNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users export csv not found response has a 2xx status code
func (o *UsersExportCsvNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users export csv not found response has a 3xx status code
func (o *UsersExportCsvNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv not found response has a 4xx status code
func (o *UsersExportCsvNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users export csv not found response has a 5xx status code
func (o *UsersExportCsvNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users export csv not found response a status code equal to that given
func (o *UsersExportCsvNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *UsersExportCsvNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *UsersExportCsvNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersExportCsvInternalServerError creates a UsersExportCsvInternalServerError with default headers values
func NewUsersExportCsvInternalServerError() *UsersExportCsvInternalServerError {
	return &UsersExportCsvInternalServerError{}
}

/*
UsersExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersExportCsvInternalServerError struct {
}

// IsSuccess returns true when this users export csv internal server error response has a 2xx status code
func (o *UsersExportCsvInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users export csv internal server error response has a 3xx status code
func (o *UsersExportCsvInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users export csv internal server error response has a 4xx status code
func (o *UsersExportCsvInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users export csv internal server error response has a 5xx status code
func (o *UsersExportCsvInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users export csv internal server error response a status code equal to that given
func (o *UsersExportCsvInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvInternalServerError ", 500)
}

func (o *UsersExportCsvInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/export][%d] usersExportCsvInternalServerError ", 500)
}

func (o *UsersExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
