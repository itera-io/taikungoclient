// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminCreateUserReader is a Reader for the AdminCreateUser structure.
type AdminCreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminCreateUserOK creates a AdminCreateUserOK with default headers values
func NewAdminCreateUserOK() *AdminCreateUserOK {
	return &AdminCreateUserOK{}
}

/*
AdminCreateUserOK describes a response with status code 200, with default header values.

Success
*/
type AdminCreateUserOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this admin create user o k response has a 2xx status code
func (o *AdminCreateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin create user o k response has a 3xx status code
func (o *AdminCreateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user o k response has a 4xx status code
func (o *AdminCreateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin create user o k response has a 5xx status code
func (o *AdminCreateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user o k response a status code equal to that given
func (o *AdminCreateUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminCreateUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserOK  %+v", 200, o.Payload)
}

func (o *AdminCreateUserOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserOK  %+v", 200, o.Payload)
}

func (o *AdminCreateUserOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminCreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserBadRequest creates a AdminCreateUserBadRequest with default headers values
func NewAdminCreateUserBadRequest() *AdminCreateUserBadRequest {
	return &AdminCreateUserBadRequest{}
}

/*
AdminCreateUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminCreateUserBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin create user bad request response has a 2xx status code
func (o *AdminCreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user bad request response has a 3xx status code
func (o *AdminCreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user bad request response has a 4xx status code
func (o *AdminCreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user bad request response has a 5xx status code
func (o *AdminCreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user bad request response a status code equal to that given
func (o *AdminCreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminCreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminCreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminCreateUserBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminCreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserUnauthorized creates a AdminCreateUserUnauthorized with default headers values
func NewAdminCreateUserUnauthorized() *AdminCreateUserUnauthorized {
	return &AdminCreateUserUnauthorized{}
}

/*
AdminCreateUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminCreateUserUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin create user unauthorized response has a 2xx status code
func (o *AdminCreateUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user unauthorized response has a 3xx status code
func (o *AdminCreateUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user unauthorized response has a 4xx status code
func (o *AdminCreateUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user unauthorized response has a 5xx status code
func (o *AdminCreateUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user unauthorized response a status code equal to that given
func (o *AdminCreateUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminCreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminCreateUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminCreateUserUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminCreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserForbidden creates a AdminCreateUserForbidden with default headers values
func NewAdminCreateUserForbidden() *AdminCreateUserForbidden {
	return &AdminCreateUserForbidden{}
}

/*
AdminCreateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminCreateUserForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin create user forbidden response has a 2xx status code
func (o *AdminCreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user forbidden response has a 3xx status code
func (o *AdminCreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user forbidden response has a 4xx status code
func (o *AdminCreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user forbidden response has a 5xx status code
func (o *AdminCreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user forbidden response a status code equal to that given
func (o *AdminCreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminCreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminCreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminCreateUserForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminCreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserNotFound creates a AdminCreateUserNotFound with default headers values
func NewAdminCreateUserNotFound() *AdminCreateUserNotFound {
	return &AdminCreateUserNotFound{}
}

/*
AdminCreateUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminCreateUserNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin create user not found response has a 2xx status code
func (o *AdminCreateUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user not found response has a 3xx status code
func (o *AdminCreateUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user not found response has a 4xx status code
func (o *AdminCreateUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create user not found response has a 5xx status code
func (o *AdminCreateUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create user not found response a status code equal to that given
func (o *AdminCreateUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminCreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminCreateUserNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminCreateUserNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminCreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateUserInternalServerError creates a AdminCreateUserInternalServerError with default headers values
func NewAdminCreateUserInternalServerError() *AdminCreateUserInternalServerError {
	return &AdminCreateUserInternalServerError{}
}

/*
AdminCreateUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminCreateUserInternalServerError struct {
}

// IsSuccess returns true when this admin create user internal server error response has a 2xx status code
func (o *AdminCreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create user internal server error response has a 3xx status code
func (o *AdminCreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create user internal server error response has a 4xx status code
func (o *AdminCreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin create user internal server error response has a 5xx status code
func (o *AdminCreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin create user internal server error response a status code equal to that given
func (o *AdminCreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminCreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserInternalServerError ", 500)
}

func (o *AdminCreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/create][%d] adminCreateUserInternalServerError ", 500)
}

func (o *AdminCreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
