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

// UserGroupsListReader is a Reader for the UserGroupsList structure.
type UserGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsListOK creates a UserGroupsListOK with default headers values
func NewUserGroupsListOK() *UserGroupsListOK {
	return &UserGroupsListOK{}
}

/*
UserGroupsListOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsListOK struct {
	Payload *models.UserGroupList
}

// IsSuccess returns true when this user groups list o k response has a 2xx status code
func (o *UserGroupsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups list o k response has a 3xx status code
func (o *UserGroupsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list o k response has a 4xx status code
func (o *UserGroupsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups list o k response has a 5xx status code
func (o *UserGroupsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list o k response a status code equal to that given
func (o *UserGroupsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListOK  %+v", 200, o.Payload)
}

func (o *UserGroupsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListOK  %+v", 200, o.Payload)
}

func (o *UserGroupsListOK) GetPayload() *models.UserGroupList {
	return o.Payload
}

func (o *UserGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListBadRequest creates a UserGroupsListBadRequest with default headers values
func NewUserGroupsListBadRequest() *UserGroupsListBadRequest {
	return &UserGroupsListBadRequest{}
}

/*
UserGroupsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups list bad request response has a 2xx status code
func (o *UserGroupsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list bad request response has a 3xx status code
func (o *UserGroupsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list bad request response has a 4xx status code
func (o *UserGroupsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list bad request response has a 5xx status code
func (o *UserGroupsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list bad request response a status code equal to that given
func (o *UserGroupsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGroupsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListUnauthorized creates a UserGroupsListUnauthorized with default headers values
func NewUserGroupsListUnauthorized() *UserGroupsListUnauthorized {
	return &UserGroupsListUnauthorized{}
}

/*
UserGroupsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list unauthorized response has a 2xx status code
func (o *UserGroupsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list unauthorized response has a 3xx status code
func (o *UserGroupsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list unauthorized response has a 4xx status code
func (o *UserGroupsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list unauthorized response has a 5xx status code
func (o *UserGroupsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list unauthorized response a status code equal to that given
func (o *UserGroupsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserGroupsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListForbidden creates a UserGroupsListForbidden with default headers values
func NewUserGroupsListForbidden() *UserGroupsListForbidden {
	return &UserGroupsListForbidden{}
}

/*
UserGroupsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list forbidden response has a 2xx status code
func (o *UserGroupsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list forbidden response has a 3xx status code
func (o *UserGroupsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list forbidden response has a 4xx status code
func (o *UserGroupsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list forbidden response has a 5xx status code
func (o *UserGroupsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list forbidden response a status code equal to that given
func (o *UserGroupsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGroupsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListNotFound creates a UserGroupsListNotFound with default headers values
func NewUserGroupsListNotFound() *UserGroupsListNotFound {
	return &UserGroupsListNotFound{}
}

/*
UserGroupsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list not found response has a 2xx status code
func (o *UserGroupsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list not found response has a 3xx status code
func (o *UserGroupsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list not found response has a 4xx status code
func (o *UserGroupsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list not found response has a 5xx status code
func (o *UserGroupsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list not found response a status code equal to that given
func (o *UserGroupsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGroupsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListInternalServerError creates a UserGroupsListInternalServerError with default headers values
func NewUserGroupsListInternalServerError() *UserGroupsListInternalServerError {
	return &UserGroupsListInternalServerError{}
}

/*
UserGroupsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsListInternalServerError struct {
}

// IsSuccess returns true when this user groups list internal server error response has a 2xx status code
func (o *UserGroupsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list internal server error response has a 3xx status code
func (o *UserGroupsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list internal server error response has a 4xx status code
func (o *UserGroupsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups list internal server error response has a 5xx status code
func (o *UserGroupsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user groups list internal server error response a status code equal to that given
func (o *UserGroupsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGroupsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListInternalServerError ", 500)
}

func (o *UserGroupsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListInternalServerError ", 500)
}

func (o *UserGroupsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
