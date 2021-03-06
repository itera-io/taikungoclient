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

// UserGroupsBindProjectGroupsReader is a Reader for the UserGroupsBindProjectGroups structure.
type UserGroupsBindProjectGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsBindProjectGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsBindProjectGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsBindProjectGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsBindProjectGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsBindProjectGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsBindProjectGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsBindProjectGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsBindProjectGroupsOK creates a UserGroupsBindProjectGroupsOK with default headers values
func NewUserGroupsBindProjectGroupsOK() *UserGroupsBindProjectGroupsOK {
	return &UserGroupsBindProjectGroupsOK{}
}

/* UserGroupsBindProjectGroupsOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsBindProjectGroupsOK struct {
	Payload models.Unit
}

func (o *UserGroupsBindProjectGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsOK  %+v", 200, o.Payload)
}
func (o *UserGroupsBindProjectGroupsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserGroupsBindProjectGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsBindProjectGroupsBadRequest creates a UserGroupsBindProjectGroupsBadRequest with default headers values
func NewUserGroupsBindProjectGroupsBadRequest() *UserGroupsBindProjectGroupsBadRequest {
	return &UserGroupsBindProjectGroupsBadRequest{}
}

/* UserGroupsBindProjectGroupsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsBindProjectGroupsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *UserGroupsBindProjectGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsBadRequest  %+v", 400, o.Payload)
}
func (o *UserGroupsBindProjectGroupsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *UserGroupsBindProjectGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsBindProjectGroupsUnauthorized creates a UserGroupsBindProjectGroupsUnauthorized with default headers values
func NewUserGroupsBindProjectGroupsUnauthorized() *UserGroupsBindProjectGroupsUnauthorized {
	return &UserGroupsBindProjectGroupsUnauthorized{}
}

/* UserGroupsBindProjectGroupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsBindProjectGroupsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsBindProjectGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsUnauthorized  %+v", 401, o.Payload)
}
func (o *UserGroupsBindProjectGroupsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsBindProjectGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsBindProjectGroupsForbidden creates a UserGroupsBindProjectGroupsForbidden with default headers values
func NewUserGroupsBindProjectGroupsForbidden() *UserGroupsBindProjectGroupsForbidden {
	return &UserGroupsBindProjectGroupsForbidden{}
}

/* UserGroupsBindProjectGroupsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsBindProjectGroupsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsBindProjectGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsForbidden  %+v", 403, o.Payload)
}
func (o *UserGroupsBindProjectGroupsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsBindProjectGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsBindProjectGroupsNotFound creates a UserGroupsBindProjectGroupsNotFound with default headers values
func NewUserGroupsBindProjectGroupsNotFound() *UserGroupsBindProjectGroupsNotFound {
	return &UserGroupsBindProjectGroupsNotFound{}
}

/* UserGroupsBindProjectGroupsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsBindProjectGroupsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *UserGroupsBindProjectGroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsNotFound  %+v", 404, o.Payload)
}
func (o *UserGroupsBindProjectGroupsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsBindProjectGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsBindProjectGroupsInternalServerError creates a UserGroupsBindProjectGroupsInternalServerError with default headers values
func NewUserGroupsBindProjectGroupsInternalServerError() *UserGroupsBindProjectGroupsInternalServerError {
	return &UserGroupsBindProjectGroupsInternalServerError{}
}

/* UserGroupsBindProjectGroupsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsBindProjectGroupsInternalServerError struct {
}

func (o *UserGroupsBindProjectGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserGroups/bind-project-groups][%d] userGroupsBindProjectGroupsInternalServerError ", 500)
}

func (o *UserGroupsBindProjectGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
