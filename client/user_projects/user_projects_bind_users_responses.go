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

// UserProjectsBindUsersReader is a Reader for the UserProjectsBindUsers structure.
type UserProjectsBindUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProjectsBindUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProjectsBindUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserProjectsBindUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserProjectsBindUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserProjectsBindUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserProjectsBindUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserProjectsBindUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserProjectsBindUsersOK creates a UserProjectsBindUsersOK with default headers values
func NewUserProjectsBindUsersOK() *UserProjectsBindUsersOK {
	return &UserProjectsBindUsersOK{}
}

/*
UserProjectsBindUsersOK describes a response with status code 200, with default header values.

Success
*/
type UserProjectsBindUsersOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user projects bind users o k response has a 2xx status code
func (o *UserProjectsBindUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user projects bind users o k response has a 3xx status code
func (o *UserProjectsBindUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users o k response has a 4xx status code
func (o *UserProjectsBindUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects bind users o k response has a 5xx status code
func (o *UserProjectsBindUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind users o k response a status code equal to that given
func (o *UserProjectsBindUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user projects bind users o k response
func (o *UserProjectsBindUsersOK) Code() int {
	return 200
}

func (o *UserProjectsBindUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersOK  %+v", 200, o.Payload)
}

func (o *UserProjectsBindUsersOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersOK  %+v", 200, o.Payload)
}

func (o *UserProjectsBindUsersOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserProjectsBindUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindUsersBadRequest creates a UserProjectsBindUsersBadRequest with default headers values
func NewUserProjectsBindUsersBadRequest() *UserProjectsBindUsersBadRequest {
	return &UserProjectsBindUsersBadRequest{}
}

/*
UserProjectsBindUsersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserProjectsBindUsersBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects bind users bad request response has a 2xx status code
func (o *UserProjectsBindUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind users bad request response has a 3xx status code
func (o *UserProjectsBindUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users bad request response has a 4xx status code
func (o *UserProjectsBindUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind users bad request response has a 5xx status code
func (o *UserProjectsBindUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind users bad request response a status code equal to that given
func (o *UserProjectsBindUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user projects bind users bad request response
func (o *UserProjectsBindUsersBadRequest) Code() int {
	return 400
}

func (o *UserProjectsBindUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsBindUsersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsBindUsersBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsBindUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindUsersUnauthorized creates a UserProjectsBindUsersUnauthorized with default headers values
func NewUserProjectsBindUsersUnauthorized() *UserProjectsBindUsersUnauthorized {
	return &UserProjectsBindUsersUnauthorized{}
}

/*
UserProjectsBindUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserProjectsBindUsersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects bind users unauthorized response has a 2xx status code
func (o *UserProjectsBindUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind users unauthorized response has a 3xx status code
func (o *UserProjectsBindUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users unauthorized response has a 4xx status code
func (o *UserProjectsBindUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind users unauthorized response has a 5xx status code
func (o *UserProjectsBindUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind users unauthorized response a status code equal to that given
func (o *UserProjectsBindUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user projects bind users unauthorized response
func (o *UserProjectsBindUsersUnauthorized) Code() int {
	return 401
}

func (o *UserProjectsBindUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsBindUsersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsBindUsersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsBindUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindUsersForbidden creates a UserProjectsBindUsersForbidden with default headers values
func NewUserProjectsBindUsersForbidden() *UserProjectsBindUsersForbidden {
	return &UserProjectsBindUsersForbidden{}
}

/*
UserProjectsBindUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserProjectsBindUsersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects bind users forbidden response has a 2xx status code
func (o *UserProjectsBindUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind users forbidden response has a 3xx status code
func (o *UserProjectsBindUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users forbidden response has a 4xx status code
func (o *UserProjectsBindUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind users forbidden response has a 5xx status code
func (o *UserProjectsBindUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind users forbidden response a status code equal to that given
func (o *UserProjectsBindUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user projects bind users forbidden response
func (o *UserProjectsBindUsersForbidden) Code() int {
	return 403
}

func (o *UserProjectsBindUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsBindUsersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsBindUsersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsBindUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindUsersNotFound creates a UserProjectsBindUsersNotFound with default headers values
func NewUserProjectsBindUsersNotFound() *UserProjectsBindUsersNotFound {
	return &UserProjectsBindUsersNotFound{}
}

/*
UserProjectsBindUsersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserProjectsBindUsersNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user projects bind users not found response has a 2xx status code
func (o *UserProjectsBindUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind users not found response has a 3xx status code
func (o *UserProjectsBindUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users not found response has a 4xx status code
func (o *UserProjectsBindUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind users not found response has a 5xx status code
func (o *UserProjectsBindUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind users not found response a status code equal to that given
func (o *UserProjectsBindUsersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user projects bind users not found response
func (o *UserProjectsBindUsersNotFound) Code() int {
	return 404
}

func (o *UserProjectsBindUsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsBindUsersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsBindUsersNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserProjectsBindUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindUsersInternalServerError creates a UserProjectsBindUsersInternalServerError with default headers values
func NewUserProjectsBindUsersInternalServerError() *UserProjectsBindUsersInternalServerError {
	return &UserProjectsBindUsersInternalServerError{}
}

/*
UserProjectsBindUsersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserProjectsBindUsersInternalServerError struct {
}

// IsSuccess returns true when this user projects bind users internal server error response has a 2xx status code
func (o *UserProjectsBindUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind users internal server error response has a 3xx status code
func (o *UserProjectsBindUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind users internal server error response has a 4xx status code
func (o *UserProjectsBindUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects bind users internal server error response has a 5xx status code
func (o *UserProjectsBindUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user projects bind users internal server error response a status code equal to that given
func (o *UserProjectsBindUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user projects bind users internal server error response
func (o *UserProjectsBindUsersInternalServerError) Code() int {
	return 500
}

func (o *UserProjectsBindUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersInternalServerError ", 500)
}

func (o *UserProjectsBindUsersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindusers][%d] userProjectsBindUsersInternalServerError ", 500)
}

func (o *UserProjectsBindUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
