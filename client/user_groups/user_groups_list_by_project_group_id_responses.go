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

// UserGroupsListByProjectGroupIDReader is a Reader for the UserGroupsListByProjectGroupID structure.
type UserGroupsListByProjectGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsListByProjectGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsListByProjectGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsListByProjectGroupIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsListByProjectGroupIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsListByProjectGroupIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsListByProjectGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsListByProjectGroupIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsListByProjectGroupIDOK creates a UserGroupsListByProjectGroupIDOK with default headers values
func NewUserGroupsListByProjectGroupIDOK() *UserGroupsListByProjectGroupIDOK {
	return &UserGroupsListByProjectGroupIDOK{}
}

/*
UserGroupsListByProjectGroupIDOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsListByProjectGroupIDOK struct {
	Payload []*models.CommonDropdownDto
}

// IsSuccess returns true when this user groups list by project group Id o k response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups list by project group Id o k response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id o k response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups list by project group Id o k response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list by project group Id o k response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGroupsListByProjectGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdOK  %+v", 200, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdOK  %+v", 200, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *UserGroupsListByProjectGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListByProjectGroupIDBadRequest creates a UserGroupsListByProjectGroupIDBadRequest with default headers values
func NewUserGroupsListByProjectGroupIDBadRequest() *UserGroupsListByProjectGroupIDBadRequest {
	return &UserGroupsListByProjectGroupIDBadRequest{}
}

/*
UserGroupsListByProjectGroupIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsListByProjectGroupIDBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups list by project group Id bad request response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list by project group Id bad request response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id bad request response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list by project group Id bad request response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list by project group Id bad request response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGroupsListByProjectGroupIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsListByProjectGroupIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListByProjectGroupIDUnauthorized creates a UserGroupsListByProjectGroupIDUnauthorized with default headers values
func NewUserGroupsListByProjectGroupIDUnauthorized() *UserGroupsListByProjectGroupIDUnauthorized {
	return &UserGroupsListByProjectGroupIDUnauthorized{}
}

/*
UserGroupsListByProjectGroupIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsListByProjectGroupIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list by project group Id unauthorized response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list by project group Id unauthorized response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id unauthorized response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list by project group Id unauthorized response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list by project group Id unauthorized response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserGroupsListByProjectGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListByProjectGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListByProjectGroupIDForbidden creates a UserGroupsListByProjectGroupIDForbidden with default headers values
func NewUserGroupsListByProjectGroupIDForbidden() *UserGroupsListByProjectGroupIDForbidden {
	return &UserGroupsListByProjectGroupIDForbidden{}
}

/*
UserGroupsListByProjectGroupIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsListByProjectGroupIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list by project group Id forbidden response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list by project group Id forbidden response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id forbidden response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list by project group Id forbidden response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list by project group Id forbidden response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGroupsListByProjectGroupIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListByProjectGroupIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListByProjectGroupIDNotFound creates a UserGroupsListByProjectGroupIDNotFound with default headers values
func NewUserGroupsListByProjectGroupIDNotFound() *UserGroupsListByProjectGroupIDNotFound {
	return &UserGroupsListByProjectGroupIDNotFound{}
}

/*
UserGroupsListByProjectGroupIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsListByProjectGroupIDNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups list by project group Id not found response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list by project group Id not found response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id not found response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups list by project group Id not found response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups list by project group Id not found response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGroupsListByProjectGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsListByProjectGroupIDNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsListByProjectGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListByProjectGroupIDInternalServerError creates a UserGroupsListByProjectGroupIDInternalServerError with default headers values
func NewUserGroupsListByProjectGroupIDInternalServerError() *UserGroupsListByProjectGroupIDInternalServerError {
	return &UserGroupsListByProjectGroupIDInternalServerError{}
}

/*
UserGroupsListByProjectGroupIDInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsListByProjectGroupIDInternalServerError struct {
}

// IsSuccess returns true when this user groups list by project group Id internal server error response has a 2xx status code
func (o *UserGroupsListByProjectGroupIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups list by project group Id internal server error response has a 3xx status code
func (o *UserGroupsListByProjectGroupIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups list by project group Id internal server error response has a 4xx status code
func (o *UserGroupsListByProjectGroupIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups list by project group Id internal server error response has a 5xx status code
func (o *UserGroupsListByProjectGroupIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user groups list by project group Id internal server error response a status code equal to that given
func (o *UserGroupsListByProjectGroupIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGroupsListByProjectGroupIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdInternalServerError ", 500)
}

func (o *UserGroupsListByProjectGroupIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list-by-project-group-id][%d] userGroupsListByProjectGroupIdInternalServerError ", 500)
}

func (o *UserGroupsListByProjectGroupIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
