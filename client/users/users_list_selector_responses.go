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

// UsersListSelectorReader is a Reader for the UsersListSelector structure.
type UsersListSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersListSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersListSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersListSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersListSelectorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersListSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersListSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersListSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersListSelectorOK creates a UsersListSelectorOK with default headers values
func NewUsersListSelectorOK() *UsersListSelectorOK {
	return &UsersListSelectorOK{}
}

/*
UsersListSelectorOK describes a response with status code 200, with default header values.

Success
*/
type UsersListSelectorOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

// IsSuccess returns true when this users list selector o k response has a 2xx status code
func (o *UsersListSelectorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users list selector o k response has a 3xx status code
func (o *UsersListSelectorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector o k response has a 4xx status code
func (o *UsersListSelectorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users list selector o k response has a 5xx status code
func (o *UsersListSelectorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users list selector o k response a status code equal to that given
func (o *UsersListSelectorOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersListSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorOK  %+v", 200, o.Payload)
}

func (o *UsersListSelectorOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorOK  %+v", 200, o.Payload)
}

func (o *UsersListSelectorOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *UsersListSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListSelectorBadRequest creates a UsersListSelectorBadRequest with default headers values
func NewUsersListSelectorBadRequest() *UsersListSelectorBadRequest {
	return &UsersListSelectorBadRequest{}
}

/*
UsersListSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersListSelectorBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this users list selector bad request response has a 2xx status code
func (o *UsersListSelectorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list selector bad request response has a 3xx status code
func (o *UsersListSelectorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector bad request response has a 4xx status code
func (o *UsersListSelectorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list selector bad request response has a 5xx status code
func (o *UsersListSelectorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users list selector bad request response a status code equal to that given
func (o *UsersListSelectorBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersListSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *UsersListSelectorBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *UsersListSelectorBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *UsersListSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListSelectorUnauthorized creates a UsersListSelectorUnauthorized with default headers values
func NewUsersListSelectorUnauthorized() *UsersListSelectorUnauthorized {
	return &UsersListSelectorUnauthorized{}
}

/*
UsersListSelectorUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersListSelectorUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list selector unauthorized response has a 2xx status code
func (o *UsersListSelectorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list selector unauthorized response has a 3xx status code
func (o *UsersListSelectorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector unauthorized response has a 4xx status code
func (o *UsersListSelectorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list selector unauthorized response has a 5xx status code
func (o *UsersListSelectorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users list selector unauthorized response a status code equal to that given
func (o *UsersListSelectorUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersListSelectorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersListSelectorUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersListSelectorUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListSelectorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListSelectorForbidden creates a UsersListSelectorForbidden with default headers values
func NewUsersListSelectorForbidden() *UsersListSelectorForbidden {
	return &UsersListSelectorForbidden{}
}

/*
UsersListSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersListSelectorForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list selector forbidden response has a 2xx status code
func (o *UsersListSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list selector forbidden response has a 3xx status code
func (o *UsersListSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector forbidden response has a 4xx status code
func (o *UsersListSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list selector forbidden response has a 5xx status code
func (o *UsersListSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users list selector forbidden response a status code equal to that given
func (o *UsersListSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersListSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorForbidden  %+v", 403, o.Payload)
}

func (o *UsersListSelectorForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorForbidden  %+v", 403, o.Payload)
}

func (o *UsersListSelectorForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListSelectorNotFound creates a UsersListSelectorNotFound with default headers values
func NewUsersListSelectorNotFound() *UsersListSelectorNotFound {
	return &UsersListSelectorNotFound{}
}

/*
UsersListSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersListSelectorNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list selector not found response has a 2xx status code
func (o *UsersListSelectorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list selector not found response has a 3xx status code
func (o *UsersListSelectorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector not found response has a 4xx status code
func (o *UsersListSelectorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list selector not found response has a 5xx status code
func (o *UsersListSelectorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users list selector not found response a status code equal to that given
func (o *UsersListSelectorNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersListSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorNotFound  %+v", 404, o.Payload)
}

func (o *UsersListSelectorNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorNotFound  %+v", 404, o.Payload)
}

func (o *UsersListSelectorNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListSelectorInternalServerError creates a UsersListSelectorInternalServerError with default headers values
func NewUsersListSelectorInternalServerError() *UsersListSelectorInternalServerError {
	return &UsersListSelectorInternalServerError{}
}

/*
UsersListSelectorInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersListSelectorInternalServerError struct {
}

// IsSuccess returns true when this users list selector internal server error response has a 2xx status code
func (o *UsersListSelectorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list selector internal server error response has a 3xx status code
func (o *UsersListSelectorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list selector internal server error response has a 4xx status code
func (o *UsersListSelectorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users list selector internal server error response has a 5xx status code
func (o *UsersListSelectorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users list selector internal server error response a status code equal to that given
func (o *UsersListSelectorInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersListSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorInternalServerError ", 500)
}

func (o *UsersListSelectorInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/list][%d] usersListSelectorInternalServerError ", 500)
}

func (o *UsersListSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
