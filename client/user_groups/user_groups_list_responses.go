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

/* UserGroupsListOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsListOK struct {
	Payload *models.UserGroupList
}

func (o *UserGroupsListOK) Error() string {
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

/* UserGroupsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *UserGroupsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListBadRequest  %+v", 400, o.Payload)
}
func (o *UserGroupsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UserGroupsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsListUnauthorized creates a UserGroupsListUnauthorized with default headers values
func NewUserGroupsListUnauthorized() *UserGroupsListUnauthorized {
	return &UserGroupsListUnauthorized{}
}

/* UserGroupsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsListUnauthorized) Error() string {
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

/* UserGroupsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsListForbidden) Error() string {
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

/* UserGroupsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsListNotFound) Error() string {
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

/* UserGroupsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsListInternalServerError struct {
}

func (o *UserGroupsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserGroups/list][%d] userGroupsListInternalServerError ", 500)
}

func (o *UserGroupsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}