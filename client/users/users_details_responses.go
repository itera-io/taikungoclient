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

// UsersDetailsReader is a Reader for the UsersDetails structure.
type UsersDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersDetailsOK creates a UsersDetailsOK with default headers values
func NewUsersDetailsOK() *UsersDetailsOK {
	return &UsersDetailsOK{}
}

/*
UsersDetailsOK describes a response with status code 200, with default header values.

Success
*/
type UsersDetailsOK struct {
	Payload *models.UserDetails
}

// IsSuccess returns true when this users details o k response has a 2xx status code
func (o *UsersDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users details o k response has a 3xx status code
func (o *UsersDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details o k response has a 4xx status code
func (o *UsersDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users details o k response has a 5xx status code
func (o *UsersDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users details o k response a status code equal to that given
func (o *UsersDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsOK  %+v", 200, o.Payload)
}

func (o *UsersDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsOK  %+v", 200, o.Payload)
}

func (o *UsersDetailsOK) GetPayload() *models.UserDetails {
	return o.Payload
}

func (o *UsersDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsBadRequest creates a UsersDetailsBadRequest with default headers values
func NewUsersDetailsBadRequest() *UsersDetailsBadRequest {
	return &UsersDetailsBadRequest{}
}

/*
UsersDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersDetailsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this users details bad request response has a 2xx status code
func (o *UsersDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details bad request response has a 3xx status code
func (o *UsersDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details bad request response has a 4xx status code
func (o *UsersDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details bad request response has a 5xx status code
func (o *UsersDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users details bad request response a status code equal to that given
func (o *UsersDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *UsersDetailsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UsersDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsUnauthorized creates a UsersDetailsUnauthorized with default headers values
func NewUsersDetailsUnauthorized() *UsersDetailsUnauthorized {
	return &UsersDetailsUnauthorized{}
}

/*
UsersDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersDetailsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users details unauthorized response has a 2xx status code
func (o *UsersDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details unauthorized response has a 3xx status code
func (o *UsersDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details unauthorized response has a 4xx status code
func (o *UsersDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details unauthorized response has a 5xx status code
func (o *UsersDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users details unauthorized response a status code equal to that given
func (o *UsersDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersDetailsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsForbidden creates a UsersDetailsForbidden with default headers values
func NewUsersDetailsForbidden() *UsersDetailsForbidden {
	return &UsersDetailsForbidden{}
}

/*
UsersDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersDetailsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users details forbidden response has a 2xx status code
func (o *UsersDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details forbidden response has a 3xx status code
func (o *UsersDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details forbidden response has a 4xx status code
func (o *UsersDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details forbidden response has a 5xx status code
func (o *UsersDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users details forbidden response a status code equal to that given
func (o *UsersDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UsersDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsForbidden  %+v", 403, o.Payload)
}

func (o *UsersDetailsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsNotFound creates a UsersDetailsNotFound with default headers values
func NewUsersDetailsNotFound() *UsersDetailsNotFound {
	return &UsersDetailsNotFound{}
}

/*
UsersDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersDetailsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this users details not found response has a 2xx status code
func (o *UsersDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details not found response has a 3xx status code
func (o *UsersDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details not found response has a 4xx status code
func (o *UsersDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users details not found response has a 5xx status code
func (o *UsersDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users details not found response a status code equal to that given
func (o *UsersDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UsersDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsNotFound  %+v", 404, o.Payload)
}

func (o *UsersDetailsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UsersDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersDetailsInternalServerError creates a UsersDetailsInternalServerError with default headers values
func NewUsersDetailsInternalServerError() *UsersDetailsInternalServerError {
	return &UsersDetailsInternalServerError{}
}

/*
UsersDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersDetailsInternalServerError struct {
}

// IsSuccess returns true when this users details internal server error response has a 2xx status code
func (o *UsersDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users details internal server error response has a 3xx status code
func (o *UsersDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users details internal server error response has a 4xx status code
func (o *UsersDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users details internal server error response has a 5xx status code
func (o *UsersDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users details internal server error response a status code equal to that given
func (o *UsersDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsInternalServerError ", 500)
}

func (o *UsersDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Users/userinfo][%d] usersDetailsInternalServerError ", 500)
}

func (o *UsersDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
