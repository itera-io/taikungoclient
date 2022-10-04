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

// UsersListReader is a Reader for the UsersList structure.
type UsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersListOK creates a UsersListOK with default headers values
func NewUsersListOK() *UsersListOK {
	return &UsersListOK{}
}

/*
UsersListOK describes a response with status code 200, with default header values.

Success
*/
type UsersListOK struct {
	Payload *models.UsersList
}

// IsSuccess returns true when this users list o k response has a 2xx status code
func (o *UsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users list o k response has a 3xx status code
func (o *UsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list o k response has a 4xx status code
func (o *UsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users list o k response has a 5xx status code
func (o *UsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users list o k response a status code equal to that given
func (o *UsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListOK  %+v", 200, o.Payload)
}

func (o *UsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListOK  %+v", 200, o.Payload)
}

func (o *UsersListOK) GetPayload() *models.UsersList {
	return o.Payload
}

func (o *UsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListBadRequest creates a UsersListBadRequest with default headers values
func NewUsersListBadRequest() *UsersListBadRequest {
	return &UsersListBadRequest{}
}

/*
UsersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersListBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list bad request response has a 2xx status code
func (o *UsersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list bad request response has a 3xx status code
func (o *UsersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list bad request response has a 4xx status code
func (o *UsersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list bad request response has a 5xx status code
func (o *UsersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users list bad request response a status code equal to that given
func (o *UsersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListBadRequest  %+v", 400, o.Payload)
}

func (o *UsersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListBadRequest  %+v", 400, o.Payload)
}

func (o *UsersListBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListUnauthorized creates a UsersListUnauthorized with default headers values
func NewUsersListUnauthorized() *UsersListUnauthorized {
	return &UsersListUnauthorized{}
}

/*
UsersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersListUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list unauthorized response has a 2xx status code
func (o *UsersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list unauthorized response has a 3xx status code
func (o *UsersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list unauthorized response has a 4xx status code
func (o *UsersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list unauthorized response has a 5xx status code
func (o *UsersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users list unauthorized response a status code equal to that given
func (o *UsersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersListUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListForbidden creates a UsersListForbidden with default headers values
func NewUsersListForbidden() *UsersListForbidden {
	return &UsersListForbidden{}
}

/*
UsersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersListForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list forbidden response has a 2xx status code
func (o *UsersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list forbidden response has a 3xx status code
func (o *UsersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list forbidden response has a 4xx status code
func (o *UsersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list forbidden response has a 5xx status code
func (o *UsersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users list forbidden response a status code equal to that given
func (o *UsersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListForbidden  %+v", 403, o.Payload)
}

func (o *UsersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListForbidden  %+v", 403, o.Payload)
}

func (o *UsersListForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListNotFound creates a UsersListNotFound with default headers values
func NewUsersListNotFound() *UsersListNotFound {
	return &UsersListNotFound{}
}

/*
UsersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersListNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users list not found response has a 2xx status code
func (o *UsersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list not found response has a 3xx status code
func (o *UsersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list not found response has a 4xx status code
func (o *UsersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users list not found response has a 5xx status code
func (o *UsersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users list not found response a status code equal to that given
func (o *UsersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListNotFound  %+v", 404, o.Payload)
}

func (o *UsersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListNotFound  %+v", 404, o.Payload)
}

func (o *UsersListNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersListInternalServerError creates a UsersListInternalServerError with default headers values
func NewUsersListInternalServerError() *UsersListInternalServerError {
	return &UsersListInternalServerError{}
}

/*
UsersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersListInternalServerError struct {
}

// IsSuccess returns true when this users list internal server error response has a 2xx status code
func (o *UsersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users list internal server error response has a 3xx status code
func (o *UsersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users list internal server error response has a 4xx status code
func (o *UsersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users list internal server error response has a 5xx status code
func (o *UsersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users list internal server error response a status code equal to that given
func (o *UsersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListInternalServerError ", 500)
}

func (o *UsersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users][%d] usersListInternalServerError ", 500)
}

func (o *UsersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
