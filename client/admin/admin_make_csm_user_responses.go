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

// AdminMakeCsmUserReader is a Reader for the AdminMakeCsmUser structure.
type AdminMakeCsmUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminMakeCsmUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminMakeCsmUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminMakeCsmUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminMakeCsmUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminMakeCsmUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminMakeCsmUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminMakeCsmUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminMakeCsmUserOK creates a AdminMakeCsmUserOK with default headers values
func NewAdminMakeCsmUserOK() *AdminMakeCsmUserOK {
	return &AdminMakeCsmUserOK{}
}

/*
AdminMakeCsmUserOK describes a response with status code 200, with default header values.

Success
*/
type AdminMakeCsmUserOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this admin make csm user o k response has a 2xx status code
func (o *AdminMakeCsmUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin make csm user o k response has a 3xx status code
func (o *AdminMakeCsmUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user o k response has a 4xx status code
func (o *AdminMakeCsmUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin make csm user o k response has a 5xx status code
func (o *AdminMakeCsmUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin make csm user o k response a status code equal to that given
func (o *AdminMakeCsmUserOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminMakeCsmUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserOK  %+v", 200, o.Payload)
}

func (o *AdminMakeCsmUserOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserOK  %+v", 200, o.Payload)
}

func (o *AdminMakeCsmUserOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AdminMakeCsmUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeCsmUserBadRequest creates a AdminMakeCsmUserBadRequest with default headers values
func NewAdminMakeCsmUserBadRequest() *AdminMakeCsmUserBadRequest {
	return &AdminMakeCsmUserBadRequest{}
}

/*
AdminMakeCsmUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminMakeCsmUserBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this admin make csm user bad request response has a 2xx status code
func (o *AdminMakeCsmUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin make csm user bad request response has a 3xx status code
func (o *AdminMakeCsmUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user bad request response has a 4xx status code
func (o *AdminMakeCsmUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin make csm user bad request response has a 5xx status code
func (o *AdminMakeCsmUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin make csm user bad request response a status code equal to that given
func (o *AdminMakeCsmUserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminMakeCsmUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminMakeCsmUserBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserBadRequest  %+v", 400, o.Payload)
}

func (o *AdminMakeCsmUserBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AdminMakeCsmUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeCsmUserUnauthorized creates a AdminMakeCsmUserUnauthorized with default headers values
func NewAdminMakeCsmUserUnauthorized() *AdminMakeCsmUserUnauthorized {
	return &AdminMakeCsmUserUnauthorized{}
}

/*
AdminMakeCsmUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminMakeCsmUserUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this admin make csm user unauthorized response has a 2xx status code
func (o *AdminMakeCsmUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin make csm user unauthorized response has a 3xx status code
func (o *AdminMakeCsmUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user unauthorized response has a 4xx status code
func (o *AdminMakeCsmUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin make csm user unauthorized response has a 5xx status code
func (o *AdminMakeCsmUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin make csm user unauthorized response a status code equal to that given
func (o *AdminMakeCsmUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminMakeCsmUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminMakeCsmUserUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminMakeCsmUserUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AdminMakeCsmUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeCsmUserForbidden creates a AdminMakeCsmUserForbidden with default headers values
func NewAdminMakeCsmUserForbidden() *AdminMakeCsmUserForbidden {
	return &AdminMakeCsmUserForbidden{}
}

/*
AdminMakeCsmUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminMakeCsmUserForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this admin make csm user forbidden response has a 2xx status code
func (o *AdminMakeCsmUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin make csm user forbidden response has a 3xx status code
func (o *AdminMakeCsmUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user forbidden response has a 4xx status code
func (o *AdminMakeCsmUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin make csm user forbidden response has a 5xx status code
func (o *AdminMakeCsmUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin make csm user forbidden response a status code equal to that given
func (o *AdminMakeCsmUserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminMakeCsmUserForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminMakeCsmUserForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserForbidden  %+v", 403, o.Payload)
}

func (o *AdminMakeCsmUserForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AdminMakeCsmUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeCsmUserNotFound creates a AdminMakeCsmUserNotFound with default headers values
func NewAdminMakeCsmUserNotFound() *AdminMakeCsmUserNotFound {
	return &AdminMakeCsmUserNotFound{}
}

/*
AdminMakeCsmUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminMakeCsmUserNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this admin make csm user not found response has a 2xx status code
func (o *AdminMakeCsmUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin make csm user not found response has a 3xx status code
func (o *AdminMakeCsmUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user not found response has a 4xx status code
func (o *AdminMakeCsmUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin make csm user not found response has a 5xx status code
func (o *AdminMakeCsmUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin make csm user not found response a status code equal to that given
func (o *AdminMakeCsmUserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminMakeCsmUserNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminMakeCsmUserNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserNotFound  %+v", 404, o.Payload)
}

func (o *AdminMakeCsmUserNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AdminMakeCsmUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminMakeCsmUserInternalServerError creates a AdminMakeCsmUserInternalServerError with default headers values
func NewAdminMakeCsmUserInternalServerError() *AdminMakeCsmUserInternalServerError {
	return &AdminMakeCsmUserInternalServerError{}
}

/*
AdminMakeCsmUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminMakeCsmUserInternalServerError struct {
}

// IsSuccess returns true when this admin make csm user internal server error response has a 2xx status code
func (o *AdminMakeCsmUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin make csm user internal server error response has a 3xx status code
func (o *AdminMakeCsmUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin make csm user internal server error response has a 4xx status code
func (o *AdminMakeCsmUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin make csm user internal server error response has a 5xx status code
func (o *AdminMakeCsmUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin make csm user internal server error response a status code equal to that given
func (o *AdminMakeCsmUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminMakeCsmUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserInternalServerError ", 500)
}

func (o *AdminMakeCsmUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Admin/users/make/csm][%d] adminMakeCsmUserInternalServerError ", 500)
}

func (o *AdminMakeCsmUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
