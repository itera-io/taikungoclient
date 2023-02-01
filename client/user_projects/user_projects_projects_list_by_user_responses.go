// Code generated by go-swagger; DO NOT EDIT.

package user_projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UserProjectsProjectsListByUserReader is a Reader for the UserProjectsProjectsListByUser structure.
type UserProjectsProjectsListByUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProjectsProjectsListByUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProjectsProjectsListByUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserProjectsProjectsListByUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserProjectsProjectsListByUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserProjectsProjectsListByUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserProjectsProjectsListByUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserProjectsProjectsListByUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserProjectsProjectsListByUserOK creates a UserProjectsProjectsListByUserOK with default headers values
func NewUserProjectsProjectsListByUserOK() *UserProjectsProjectsListByUserOK {
	return &UserProjectsProjectsListByUserOK{}
}

/*
UserProjectsProjectsListByUserOK describes a response with status code 200, with default header values.

Success
*/
type UserProjectsProjectsListByUserOK struct {
	Payload []*models.CommonDropdownDto
}

// IsSuccess returns true when this user projects projects list by user o k response has a 2xx status code
func (o *UserProjectsProjectsListByUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user projects projects list by user o k response has a 3xx status code
func (o *UserProjectsProjectsListByUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user o k response has a 4xx status code
func (o *UserProjectsProjectsListByUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects projects list by user o k response has a 5xx status code
func (o *UserProjectsProjectsListByUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects projects list by user o k response a status code equal to that given
func (o *UserProjectsProjectsListByUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user projects projects list by user o k response
func (o *UserProjectsProjectsListByUserOK) Code() int {
	return 200
}

func (o *UserProjectsProjectsListByUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserOK  %+v", 200, o.Payload)
}

func (o *UserProjectsProjectsListByUserOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserOK  %+v", 200, o.Payload)
}

func (o *UserProjectsProjectsListByUserOK) GetPayload() []*models.CommonDropdownDto {
	return o.Payload
}

func (o *UserProjectsProjectsListByUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsProjectsListByUserBadRequest creates a UserProjectsProjectsListByUserBadRequest with default headers values
func NewUserProjectsProjectsListByUserBadRequest() *UserProjectsProjectsListByUserBadRequest {
	return &UserProjectsProjectsListByUserBadRequest{}
}

/*
UserProjectsProjectsListByUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserProjectsProjectsListByUserBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects projects list by user bad request response has a 2xx status code
func (o *UserProjectsProjectsListByUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects projects list by user bad request response has a 3xx status code
func (o *UserProjectsProjectsListByUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user bad request response has a 4xx status code
func (o *UserProjectsProjectsListByUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects projects list by user bad request response has a 5xx status code
func (o *UserProjectsProjectsListByUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects projects list by user bad request response a status code equal to that given
func (o *UserProjectsProjectsListByUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user projects projects list by user bad request response
func (o *UserProjectsProjectsListByUserBadRequest) Code() int {
	return 400
}

func (o *UserProjectsProjectsListByUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsProjectsListByUserBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsProjectsListByUserBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsProjectsListByUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsProjectsListByUserUnauthorized creates a UserProjectsProjectsListByUserUnauthorized with default headers values
func NewUserProjectsProjectsListByUserUnauthorized() *UserProjectsProjectsListByUserUnauthorized {
	return &UserProjectsProjectsListByUserUnauthorized{}
}

/*
UserProjectsProjectsListByUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserProjectsProjectsListByUserUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects projects list by user unauthorized response has a 2xx status code
func (o *UserProjectsProjectsListByUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects projects list by user unauthorized response has a 3xx status code
func (o *UserProjectsProjectsListByUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user unauthorized response has a 4xx status code
func (o *UserProjectsProjectsListByUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects projects list by user unauthorized response has a 5xx status code
func (o *UserProjectsProjectsListByUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects projects list by user unauthorized response a status code equal to that given
func (o *UserProjectsProjectsListByUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user projects projects list by user unauthorized response
func (o *UserProjectsProjectsListByUserUnauthorized) Code() int {
	return 401
}

func (o *UserProjectsProjectsListByUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsProjectsListByUserUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsProjectsListByUserUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsProjectsListByUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsProjectsListByUserForbidden creates a UserProjectsProjectsListByUserForbidden with default headers values
func NewUserProjectsProjectsListByUserForbidden() *UserProjectsProjectsListByUserForbidden {
	return &UserProjectsProjectsListByUserForbidden{}
}

/*
UserProjectsProjectsListByUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserProjectsProjectsListByUserForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects projects list by user forbidden response has a 2xx status code
func (o *UserProjectsProjectsListByUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects projects list by user forbidden response has a 3xx status code
func (o *UserProjectsProjectsListByUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user forbidden response has a 4xx status code
func (o *UserProjectsProjectsListByUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects projects list by user forbidden response has a 5xx status code
func (o *UserProjectsProjectsListByUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects projects list by user forbidden response a status code equal to that given
func (o *UserProjectsProjectsListByUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user projects projects list by user forbidden response
func (o *UserProjectsProjectsListByUserForbidden) Code() int {
	return 403
}

func (o *UserProjectsProjectsListByUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsProjectsListByUserForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsProjectsListByUserForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsProjectsListByUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsProjectsListByUserNotFound creates a UserProjectsProjectsListByUserNotFound with default headers values
func NewUserProjectsProjectsListByUserNotFound() *UserProjectsProjectsListByUserNotFound {
	return &UserProjectsProjectsListByUserNotFound{}
}

/*
UserProjectsProjectsListByUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserProjectsProjectsListByUserNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects projects list by user not found response has a 2xx status code
func (o *UserProjectsProjectsListByUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects projects list by user not found response has a 3xx status code
func (o *UserProjectsProjectsListByUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user not found response has a 4xx status code
func (o *UserProjectsProjectsListByUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects projects list by user not found response has a 5xx status code
func (o *UserProjectsProjectsListByUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects projects list by user not found response a status code equal to that given
func (o *UserProjectsProjectsListByUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user projects projects list by user not found response
func (o *UserProjectsProjectsListByUserNotFound) Code() int {
	return 404
}

func (o *UserProjectsProjectsListByUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsProjectsListByUserNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsProjectsListByUserNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsProjectsListByUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsProjectsListByUserInternalServerError creates a UserProjectsProjectsListByUserInternalServerError with default headers values
func NewUserProjectsProjectsListByUserInternalServerError() *UserProjectsProjectsListByUserInternalServerError {
	return &UserProjectsProjectsListByUserInternalServerError{}
}

/*
UserProjectsProjectsListByUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserProjectsProjectsListByUserInternalServerError struct {
}

// IsSuccess returns true when this user projects projects list by user internal server error response has a 2xx status code
func (o *UserProjectsProjectsListByUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects projects list by user internal server error response has a 3xx status code
func (o *UserProjectsProjectsListByUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects projects list by user internal server error response has a 4xx status code
func (o *UserProjectsProjectsListByUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects projects list by user internal server error response has a 5xx status code
func (o *UserProjectsProjectsListByUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user projects projects list by user internal server error response a status code equal to that given
func (o *UserProjectsProjectsListByUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user projects projects list by user internal server error response
func (o *UserProjectsProjectsListByUserInternalServerError) Code() int {
	return 500
}

func (o *UserProjectsProjectsListByUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserInternalServerError ", 500)
}

func (o *UserProjectsProjectsListByUserInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/UserProjects/projects/list][%d] userProjectsProjectsListByUserInternalServerError ", 500)
}

func (o *UserProjectsProjectsListByUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
