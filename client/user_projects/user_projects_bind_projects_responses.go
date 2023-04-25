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

// UserProjectsBindProjectsReader is a Reader for the UserProjectsBindProjects structure.
type UserProjectsBindProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserProjectsBindProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserProjectsBindProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserProjectsBindProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserProjectsBindProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserProjectsBindProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserProjectsBindProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserProjectsBindProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserProjectsBindProjectsOK creates a UserProjectsBindProjectsOK with default headers values
func NewUserProjectsBindProjectsOK() *UserProjectsBindProjectsOK {
	return &UserProjectsBindProjectsOK{}
}

/*
UserProjectsBindProjectsOK describes a response with status code 200, with default header values.

Success
*/
type UserProjectsBindProjectsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user projects bind projects o k response has a 2xx status code
func (o *UserProjectsBindProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user projects bind projects o k response has a 3xx status code
func (o *UserProjectsBindProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects o k response has a 4xx status code
func (o *UserProjectsBindProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects bind projects o k response has a 5xx status code
func (o *UserProjectsBindProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind projects o k response a status code equal to that given
func (o *UserProjectsBindProjectsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user projects bind projects o k response
func (o *UserProjectsBindProjectsOK) Code() int {
	return 200
}

func (o *UserProjectsBindProjectsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsOK  %+v", 200, o.Payload)
}

func (o *UserProjectsBindProjectsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsOK  %+v", 200, o.Payload)
}

func (o *UserProjectsBindProjectsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserProjectsBindProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindProjectsBadRequest creates a UserProjectsBindProjectsBadRequest with default headers values
func NewUserProjectsBindProjectsBadRequest() *UserProjectsBindProjectsBadRequest {
	return &UserProjectsBindProjectsBadRequest{}
}

/*
UserProjectsBindProjectsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserProjectsBindProjectsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this user projects bind projects bad request response has a 2xx status code
func (o *UserProjectsBindProjectsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind projects bad request response has a 3xx status code
func (o *UserProjectsBindProjectsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects bad request response has a 4xx status code
func (o *UserProjectsBindProjectsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind projects bad request response has a 5xx status code
func (o *UserProjectsBindProjectsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind projects bad request response a status code equal to that given
func (o *UserProjectsBindProjectsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user projects bind projects bad request response
func (o *UserProjectsBindProjectsBadRequest) Code() int {
	return 400
}

func (o *UserProjectsBindProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsBindProjectsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *UserProjectsBindProjectsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UserProjectsBindProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindProjectsUnauthorized creates a UserProjectsBindProjectsUnauthorized with default headers values
func NewUserProjectsBindProjectsUnauthorized() *UserProjectsBindProjectsUnauthorized {
	return &UserProjectsBindProjectsUnauthorized{}
}

/*
UserProjectsBindProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserProjectsBindProjectsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this user projects bind projects unauthorized response has a 2xx status code
func (o *UserProjectsBindProjectsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind projects unauthorized response has a 3xx status code
func (o *UserProjectsBindProjectsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects unauthorized response has a 4xx status code
func (o *UserProjectsBindProjectsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind projects unauthorized response has a 5xx status code
func (o *UserProjectsBindProjectsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind projects unauthorized response a status code equal to that given
func (o *UserProjectsBindProjectsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user projects bind projects unauthorized response
func (o *UserProjectsBindProjectsUnauthorized) Code() int {
	return 401
}

func (o *UserProjectsBindProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsBindProjectsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *UserProjectsBindProjectsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *UserProjectsBindProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindProjectsForbidden creates a UserProjectsBindProjectsForbidden with default headers values
func NewUserProjectsBindProjectsForbidden() *UserProjectsBindProjectsForbidden {
	return &UserProjectsBindProjectsForbidden{}
}

/*
UserProjectsBindProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserProjectsBindProjectsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this user projects bind projects forbidden response has a 2xx status code
func (o *UserProjectsBindProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind projects forbidden response has a 3xx status code
func (o *UserProjectsBindProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects forbidden response has a 4xx status code
func (o *UserProjectsBindProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind projects forbidden response has a 5xx status code
func (o *UserProjectsBindProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind projects forbidden response a status code equal to that given
func (o *UserProjectsBindProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user projects bind projects forbidden response
func (o *UserProjectsBindProjectsForbidden) Code() int {
	return 403
}

func (o *UserProjectsBindProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsBindProjectsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsForbidden  %+v", 403, o.Payload)
}

func (o *UserProjectsBindProjectsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UserProjectsBindProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindProjectsNotFound creates a UserProjectsBindProjectsNotFound with default headers values
func NewUserProjectsBindProjectsNotFound() *UserProjectsBindProjectsNotFound {
	return &UserProjectsBindProjectsNotFound{}
}

/*
UserProjectsBindProjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserProjectsBindProjectsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this user projects bind projects not found response has a 2xx status code
func (o *UserProjectsBindProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind projects not found response has a 3xx status code
func (o *UserProjectsBindProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects not found response has a 4xx status code
func (o *UserProjectsBindProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user projects bind projects not found response has a 5xx status code
func (o *UserProjectsBindProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user projects bind projects not found response a status code equal to that given
func (o *UserProjectsBindProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user projects bind projects not found response
func (o *UserProjectsBindProjectsNotFound) Code() int {
	return 404
}

func (o *UserProjectsBindProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsBindProjectsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsNotFound  %+v", 404, o.Payload)
}

func (o *UserProjectsBindProjectsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UserProjectsBindProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserProjectsBindProjectsInternalServerError creates a UserProjectsBindProjectsInternalServerError with default headers values
func NewUserProjectsBindProjectsInternalServerError() *UserProjectsBindProjectsInternalServerError {
	return &UserProjectsBindProjectsInternalServerError{}
}

/*
UserProjectsBindProjectsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserProjectsBindProjectsInternalServerError struct {
}

// IsSuccess returns true when this user projects bind projects internal server error response has a 2xx status code
func (o *UserProjectsBindProjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user projects bind projects internal server error response has a 3xx status code
func (o *UserProjectsBindProjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user projects bind projects internal server error response has a 4xx status code
func (o *UserProjectsBindProjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user projects bind projects internal server error response has a 5xx status code
func (o *UserProjectsBindProjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user projects bind projects internal server error response a status code equal to that given
func (o *UserProjectsBindProjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user projects bind projects internal server error response
func (o *UserProjectsBindProjectsInternalServerError) Code() int {
	return 500
}

func (o *UserProjectsBindProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsInternalServerError ", 500)
}

func (o *UserProjectsBindProjectsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserProjects/bindprojects][%d] userProjectsBindProjectsInternalServerError ", 500)
}

func (o *UserProjectsBindProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
