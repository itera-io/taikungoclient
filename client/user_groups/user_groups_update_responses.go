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

// UserGroupsUpdateReader is a Reader for the UserGroupsUpdate structure.
type UserGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsUpdateOK creates a UserGroupsUpdateOK with default headers values
func NewUserGroupsUpdateOK() *UserGroupsUpdateOK {
	return &UserGroupsUpdateOK{}
}

/*
UserGroupsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user groups update o k response has a 2xx status code
func (o *UserGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups update o k response has a 3xx status code
func (o *UserGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update o k response has a 4xx status code
func (o *UserGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups update o k response has a 5xx status code
func (o *UserGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups update o k response a status code equal to that given
func (o *UserGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UserGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UserGroupsUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsUpdateBadRequest creates a UserGroupsUpdateBadRequest with default headers values
func NewUserGroupsUpdateBadRequest() *UserGroupsUpdateBadRequest {
	return &UserGroupsUpdateBadRequest{}
}

/*
UserGroupsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsUpdateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups update bad request response has a 2xx status code
func (o *UserGroupsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups update bad request response has a 3xx status code
func (o *UserGroupsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update bad request response has a 4xx status code
func (o *UserGroupsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups update bad request response has a 5xx status code
func (o *UserGroupsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups update bad request response a status code equal to that given
func (o *UserGroupsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserGroupsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsUpdateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsUpdateUnauthorized creates a UserGroupsUpdateUnauthorized with default headers values
func NewUserGroupsUpdateUnauthorized() *UserGroupsUpdateUnauthorized {
	return &UserGroupsUpdateUnauthorized{}
}

/*
UserGroupsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups update unauthorized response has a 2xx status code
func (o *UserGroupsUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups update unauthorized response has a 3xx status code
func (o *UserGroupsUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update unauthorized response has a 4xx status code
func (o *UserGroupsUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups update unauthorized response has a 5xx status code
func (o *UserGroupsUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups update unauthorized response a status code equal to that given
func (o *UserGroupsUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserGroupsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsUpdateUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsUpdateForbidden creates a UserGroupsUpdateForbidden with default headers values
func NewUserGroupsUpdateForbidden() *UserGroupsUpdateForbidden {
	return &UserGroupsUpdateForbidden{}
}

/*
UserGroupsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsUpdateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups update forbidden response has a 2xx status code
func (o *UserGroupsUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups update forbidden response has a 3xx status code
func (o *UserGroupsUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update forbidden response has a 4xx status code
func (o *UserGroupsUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups update forbidden response has a 5xx status code
func (o *UserGroupsUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups update forbidden response a status code equal to that given
func (o *UserGroupsUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserGroupsUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsUpdateForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsUpdateNotFound creates a UserGroupsUpdateNotFound with default headers values
func NewUserGroupsUpdateNotFound() *UserGroupsUpdateNotFound {
	return &UserGroupsUpdateNotFound{}
}

/*
UserGroupsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsUpdateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user groups update not found response has a 2xx status code
func (o *UserGroupsUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups update not found response has a 3xx status code
func (o *UserGroupsUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update not found response has a 4xx status code
func (o *UserGroupsUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups update not found response has a 5xx status code
func (o *UserGroupsUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups update not found response a status code equal to that given
func (o *UserGroupsUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserGroupsUpdateNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsUpdateNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserGroupsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsUpdateInternalServerError creates a UserGroupsUpdateInternalServerError with default headers values
func NewUserGroupsUpdateInternalServerError() *UserGroupsUpdateInternalServerError {
	return &UserGroupsUpdateInternalServerError{}
}

/*
UserGroupsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsUpdateInternalServerError struct {
}

// IsSuccess returns true when this user groups update internal server error response has a 2xx status code
func (o *UserGroupsUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups update internal server error response has a 3xx status code
func (o *UserGroupsUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups update internal server error response has a 4xx status code
func (o *UserGroupsUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups update internal server error response has a 5xx status code
func (o *UserGroupsUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user groups update internal server error response a status code equal to that given
func (o *UserGroupsUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserGroupsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateInternalServerError ", 500)
}

func (o *UserGroupsUpdateInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/UserGroups/update][%d] userGroupsUpdateInternalServerError ", 500)
}

func (o *UserGroupsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
