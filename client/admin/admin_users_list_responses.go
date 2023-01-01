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

// AdminUsersListReader is a Reader for the AdminUsersList structure.
type AdminUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUsersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUsersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUsersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminUsersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUsersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminUsersListOK creates a AdminUsersListOK with default headers values
func NewAdminUsersListOK() *AdminUsersListOK {
	return &AdminUsersListOK{}
}

/*
AdminUsersListOK describes a response with status code 200, with default header values.

Success
*/
type AdminUsersListOK struct {
	Payload *models.AdminUsersList
}

// IsSuccess returns true when this admin users list o k response has a 2xx status code
func (o *AdminUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin users list o k response has a 3xx status code
func (o *AdminUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list o k response has a 4xx status code
func (o *AdminUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin users list o k response has a 5xx status code
func (o *AdminUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin users list o k response a status code equal to that given
func (o *AdminUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListOK  %+v", 200, o.Payload)
}

func (o *AdminUsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListOK  %+v", 200, o.Payload)
}

func (o *AdminUsersListOK) GetPayload() *models.AdminUsersList {
	return o.Payload
}

func (o *AdminUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminUsersList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUsersListBadRequest creates a AdminUsersListBadRequest with default headers values
func NewAdminUsersListBadRequest() *AdminUsersListBadRequest {
	return &AdminUsersListBadRequest{}
}

/*
AdminUsersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminUsersListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin users list bad request response has a 2xx status code
func (o *AdminUsersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin users list bad request response has a 3xx status code
func (o *AdminUsersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list bad request response has a 4xx status code
func (o *AdminUsersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin users list bad request response has a 5xx status code
func (o *AdminUsersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin users list bad request response a status code equal to that given
func (o *AdminUsersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminUsersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUsersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminUsersListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUsersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUsersListUnauthorized creates a AdminUsersListUnauthorized with default headers values
func NewAdminUsersListUnauthorized() *AdminUsersListUnauthorized {
	return &AdminUsersListUnauthorized{}
}

/*
AdminUsersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminUsersListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin users list unauthorized response has a 2xx status code
func (o *AdminUsersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin users list unauthorized response has a 3xx status code
func (o *AdminUsersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list unauthorized response has a 4xx status code
func (o *AdminUsersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin users list unauthorized response has a 5xx status code
func (o *AdminUsersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin users list unauthorized response a status code equal to that given
func (o *AdminUsersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminUsersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUsersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminUsersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUsersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUsersListForbidden creates a AdminUsersListForbidden with default headers values
func NewAdminUsersListForbidden() *AdminUsersListForbidden {
	return &AdminUsersListForbidden{}
}

/*
AdminUsersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminUsersListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin users list forbidden response has a 2xx status code
func (o *AdminUsersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin users list forbidden response has a 3xx status code
func (o *AdminUsersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list forbidden response has a 4xx status code
func (o *AdminUsersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin users list forbidden response has a 5xx status code
func (o *AdminUsersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin users list forbidden response a status code equal to that given
func (o *AdminUsersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminUsersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListForbidden  %+v", 403, o.Payload)
}

func (o *AdminUsersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListForbidden  %+v", 403, o.Payload)
}

func (o *AdminUsersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUsersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUsersListNotFound creates a AdminUsersListNotFound with default headers values
func NewAdminUsersListNotFound() *AdminUsersListNotFound {
	return &AdminUsersListNotFound{}
}

/*
AdminUsersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminUsersListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin users list not found response has a 2xx status code
func (o *AdminUsersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin users list not found response has a 3xx status code
func (o *AdminUsersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list not found response has a 4xx status code
func (o *AdminUsersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin users list not found response has a 5xx status code
func (o *AdminUsersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin users list not found response a status code equal to that given
func (o *AdminUsersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminUsersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListNotFound  %+v", 404, o.Payload)
}

func (o *AdminUsersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListNotFound  %+v", 404, o.Payload)
}

func (o *AdminUsersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminUsersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUsersListInternalServerError creates a AdminUsersListInternalServerError with default headers values
func NewAdminUsersListInternalServerError() *AdminUsersListInternalServerError {
	return &AdminUsersListInternalServerError{}
}

/*
AdminUsersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminUsersListInternalServerError struct {
}

// IsSuccess returns true when this admin users list internal server error response has a 2xx status code
func (o *AdminUsersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin users list internal server error response has a 3xx status code
func (o *AdminUsersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin users list internal server error response has a 4xx status code
func (o *AdminUsersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin users list internal server error response has a 5xx status code
func (o *AdminUsersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin users list internal server error response a status code equal to that given
func (o *AdminUsersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminUsersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListInternalServerError ", 500)
}

func (o *AdminUsersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/users/list][%d] adminUsersListInternalServerError ", 500)
}

func (o *AdminUsersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
