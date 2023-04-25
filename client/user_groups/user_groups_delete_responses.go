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

// UserGroupsDeleteReader is a Reader for the UserGroupsDelete structure.
type UserGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUserGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserGroupsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserGroupsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserGroupsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserGroupsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserGroupsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserGroupsDeleteOK creates a UserGroupsDeleteOK with default headers values
func NewUserGroupsDeleteOK() *UserGroupsDeleteOK {
	return &UserGroupsDeleteOK{}
}

/*
UserGroupsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type UserGroupsDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user groups delete o k response has a 2xx status code
func (o *UserGroupsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups delete o k response has a 3xx status code
func (o *UserGroupsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete o k response has a 4xx status code
func (o *UserGroupsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups delete o k response has a 5xx status code
func (o *UserGroupsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete o k response a status code equal to that given
func (o *UserGroupsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user groups delete o k response
func (o *UserGroupsDeleteOK) Code() int {
	return 200
}

func (o *UserGroupsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteOK  %+v", 200, o.Payload)
}

func (o *UserGroupsDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteOK  %+v", 200, o.Payload)
}

func (o *UserGroupsDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserGroupsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsDeleteNoContent creates a UserGroupsDeleteNoContent with default headers values
func NewUserGroupsDeleteNoContent() *UserGroupsDeleteNoContent {
	return &UserGroupsDeleteNoContent{}
}

/*
UserGroupsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type UserGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this user groups delete no content response has a 2xx status code
func (o *UserGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups delete no content response has a 3xx status code
func (o *UserGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete no content response has a 4xx status code
func (o *UserGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups delete no content response has a 5xx status code
func (o *UserGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete no content response a status code equal to that given
func (o *UserGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user groups delete no content response
func (o *UserGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *UserGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteNoContent ", 204)
}

func (o *UserGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteNoContent ", 204)
}

func (o *UserGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserGroupsDeleteBadRequest creates a UserGroupsDeleteBadRequest with default headers values
func NewUserGroupsDeleteBadRequest() *UserGroupsDeleteBadRequest {
	return &UserGroupsDeleteBadRequest{}
}

/*
UserGroupsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserGroupsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups delete bad request response has a 2xx status code
func (o *UserGroupsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups delete bad request response has a 3xx status code
func (o *UserGroupsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete bad request response has a 4xx status code
func (o *UserGroupsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups delete bad request response has a 5xx status code
func (o *UserGroupsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete bad request response a status code equal to that given
func (o *UserGroupsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user groups delete bad request response
func (o *UserGroupsDeleteBadRequest) Code() int {
	return 400
}

func (o *UserGroupsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *UserGroupsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsDeleteUnauthorized creates a UserGroupsDeleteUnauthorized with default headers values
func NewUserGroupsDeleteUnauthorized() *UserGroupsDeleteUnauthorized {
	return &UserGroupsDeleteUnauthorized{}
}

/*
UserGroupsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserGroupsDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups delete unauthorized response has a 2xx status code
func (o *UserGroupsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups delete unauthorized response has a 3xx status code
func (o *UserGroupsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete unauthorized response has a 4xx status code
func (o *UserGroupsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups delete unauthorized response has a 5xx status code
func (o *UserGroupsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete unauthorized response a status code equal to that given
func (o *UserGroupsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user groups delete unauthorized response
func (o *UserGroupsDeleteUnauthorized) Code() int {
	return 401
}

func (o *UserGroupsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *UserGroupsDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsDeleteForbidden creates a UserGroupsDeleteForbidden with default headers values
func NewUserGroupsDeleteForbidden() *UserGroupsDeleteForbidden {
	return &UserGroupsDeleteForbidden{}
}

/*
UserGroupsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserGroupsDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups delete forbidden response has a 2xx status code
func (o *UserGroupsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups delete forbidden response has a 3xx status code
func (o *UserGroupsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete forbidden response has a 4xx status code
func (o *UserGroupsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups delete forbidden response has a 5xx status code
func (o *UserGroupsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete forbidden response a status code equal to that given
func (o *UserGroupsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user groups delete forbidden response
func (o *UserGroupsDeleteForbidden) Code() int {
	return 403
}

func (o *UserGroupsDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteForbidden  %+v", 403, o.Payload)
}

func (o *UserGroupsDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsDeleteNotFound creates a UserGroupsDeleteNotFound with default headers values
func NewUserGroupsDeleteNotFound() *UserGroupsDeleteNotFound {
	return &UserGroupsDeleteNotFound{}
}

/*
UserGroupsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserGroupsDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this user groups delete not found response has a 2xx status code
func (o *UserGroupsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups delete not found response has a 3xx status code
func (o *UserGroupsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete not found response has a 4xx status code
func (o *UserGroupsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user groups delete not found response has a 5xx status code
func (o *UserGroupsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups delete not found response a status code equal to that given
func (o *UserGroupsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user groups delete not found response
func (o *UserGroupsDeleteNotFound) Code() int {
	return 404
}

func (o *UserGroupsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *UserGroupsDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGroupsDeleteInternalServerError creates a UserGroupsDeleteInternalServerError with default headers values
func NewUserGroupsDeleteInternalServerError() *UserGroupsDeleteInternalServerError {
	return &UserGroupsDeleteInternalServerError{}
}

/*
UserGroupsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserGroupsDeleteInternalServerError struct {
}

// IsSuccess returns true when this user groups delete internal server error response has a 2xx status code
func (o *UserGroupsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user groups delete internal server error response has a 3xx status code
func (o *UserGroupsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups delete internal server error response has a 4xx status code
func (o *UserGroupsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups delete internal server error response has a 5xx status code
func (o *UserGroupsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user groups delete internal server error response a status code equal to that given
func (o *UserGroupsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user groups delete internal server error response
func (o *UserGroupsDeleteInternalServerError) Code() int {
	return 500
}

func (o *UserGroupsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteInternalServerError ", 500)
}

func (o *UserGroupsDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/UserGroups][%d] userGroupsDeleteInternalServerError ", 500)
}

func (o *UserGroupsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
