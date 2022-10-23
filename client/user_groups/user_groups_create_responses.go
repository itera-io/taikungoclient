// Code generated by go-swagger; DO NOT EDIT.

package user_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UserGroupsCreateReader is a Reader for the UserGroupsCreate structure.
type UserGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsCreateOK creates a UserGroupsCreateOK with default headers values
func NewUserGroupsCreateOK() *UserGroupsCreateOK {
	return &UserGroupsCreateOK{}
}

/*
UserGroupsCreateOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this user groups create o k response has a 2xx status code
func (o *UserGroupsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups create o k response has a 3xx status code
func (o *UserGroupsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create o k response has a 4xx status code
func (o *UserGroupsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups create o k response has a 5xx status code
func (o *UserGroupsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups create o k response a status code equal to that given
func (o *UserGroupsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGroupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *UserGroupsCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *UserGroupsCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *UserGroupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsCreateBadRequest creates a UserGroupsCreateBadRequest with default headers values
func NewUserGroupsCreateBadRequest() *UserGroupsCreateBadRequest {
	return &UserGroupsCreateBadRequest{}
}

/*
UserGroupsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsCreateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this user groups create bad request response has a 2xx status code
func (o *UserGroupsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups create bad request response has a 3xx status code
func (o *UserGroupsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create bad request response has a 4xx status code
func (o *UserGroupsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups create bad request response has a 5xx status code
func (o *UserGroupsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups create bad request response a status code equal to that given
func (o *UserGroupsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGroupsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsCreateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *UserGroupsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsCreateUnauthorized creates a UserGroupsCreateUnauthorized with default headers values
func NewUserGroupsCreateUnauthorized() *UserGroupsCreateUnauthorized {
	return &UserGroupsCreateUnauthorized{}
}

/*
UserGroupsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups create unauthorized response has a 2xx status code
func (o *UserGroupsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups create unauthorized response has a 3xx status code
func (o *UserGroupsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create unauthorized response has a 4xx status code
func (o *UserGroupsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups create unauthorized response has a 5xx status code
func (o *UserGroupsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups create unauthorized response a status code equal to that given
func (o *UserGroupsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserGroupsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsCreateForbidden creates a UserGroupsCreateForbidden with default headers values
func NewUserGroupsCreateForbidden() *UserGroupsCreateForbidden {
	return &UserGroupsCreateForbidden{}
}

/*
UserGroupsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups create forbidden response has a 2xx status code
func (o *UserGroupsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups create forbidden response has a 3xx status code
func (o *UserGroupsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create forbidden response has a 4xx status code
func (o *UserGroupsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups create forbidden response has a 5xx status code
func (o *UserGroupsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups create forbidden response a status code equal to that given
func (o *UserGroupsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGroupsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsCreateNotFound creates a UserGroupsCreateNotFound with default headers values
func NewUserGroupsCreateNotFound() *UserGroupsCreateNotFound {
	return &UserGroupsCreateNotFound{}
}

/*
UserGroupsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups create not found response has a 2xx status code
func (o *UserGroupsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups create not found response has a 3xx status code
func (o *UserGroupsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create not found response has a 4xx status code
func (o *UserGroupsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups create not found response has a 5xx status code
func (o *UserGroupsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups create not found response a status code equal to that given
func (o *UserGroupsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGroupsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsCreateInternalServerError creates a UserGroupsCreateInternalServerError with default headers values
func NewUserGroupsCreateInternalServerError() *UserGroupsCreateInternalServerError {
	return &UserGroupsCreateInternalServerError{}
}

/*
UserGroupsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsCreateInternalServerError struct {
}

// IsSuccess returns true when this user groups create internal server error response has a 2xx status code
func (o *UserGroupsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups create internal server error response has a 3xx status code
func (o *UserGroupsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups create internal server error response has a 4xx status code
func (o *UserGroupsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups create internal server error response has a 5xx status code
func (o *UserGroupsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user groups create internal server error response a status code equal to that given
func (o *UserGroupsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGroupsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateInternalServerError ", 500)
}

func (o *UserGroupsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/create][%d] userGroupsCreateInternalServerError ", 500)
}

func (o *UserGroupsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
