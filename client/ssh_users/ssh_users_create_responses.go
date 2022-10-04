// Code generated by go-swagger; DO NOT EDIT.

package ssh_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SSHUsersCreateReader is a Reader for the SSHUsersCreate structure.
type SSHUsersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSHUsersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSHUsersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSSHUsersCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSSHUsersCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSSHUsersCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSSHUsersCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSSHUsersCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSSHUsersCreateOK creates a SSHUsersCreateOK with default headers values
func NewSSHUsersCreateOK() *SSHUsersCreateOK {
	return &SSHUsersCreateOK{}
}

/*
SSHUsersCreateOK describes a response with status code 200, with default header values.

Success
*/
type SSHUsersCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this ssh users create o k response has a 2xx status code
func (o *SSHUsersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ssh users create o k response has a 3xx status code
func (o *SSHUsersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create o k response has a 4xx status code
func (o *SSHUsersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users create o k response has a 5xx status code
func (o *SSHUsersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create o k response a status code equal to that given
func (o *SSHUsersCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SSHUsersCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateOK  %+v", 200, o.Payload)
}

func (o *SSHUsersCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateOK  %+v", 200, o.Payload)
}

func (o *SSHUsersCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *SSHUsersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateBadRequest creates a SSHUsersCreateBadRequest with default headers values
func NewSSHUsersCreateBadRequest() *SSHUsersCreateBadRequest {
	return &SSHUsersCreateBadRequest{}
}

/*
SSHUsersCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SSHUsersCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this ssh users create bad request response has a 2xx status code
func (o *SSHUsersCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create bad request response has a 3xx status code
func (o *SSHUsersCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create bad request response has a 4xx status code
func (o *SSHUsersCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create bad request response has a 5xx status code
func (o *SSHUsersCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create bad request response a status code equal to that given
func (o *SSHUsersCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SSHUsersCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SSHUsersCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateUnauthorized creates a SSHUsersCreateUnauthorized with default headers values
func NewSSHUsersCreateUnauthorized() *SSHUsersCreateUnauthorized {
	return &SSHUsersCreateUnauthorized{}
}

/*
SSHUsersCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SSHUsersCreateUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this ssh users create unauthorized response has a 2xx status code
func (o *SSHUsersCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create unauthorized response has a 3xx status code
func (o *SSHUsersCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create unauthorized response has a 4xx status code
func (o *SSHUsersCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create unauthorized response has a 5xx status code
func (o *SSHUsersCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create unauthorized response a status code equal to that given
func (o *SSHUsersCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SSHUsersCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersCreateUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *SSHUsersCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateForbidden creates a SSHUsersCreateForbidden with default headers values
func NewSSHUsersCreateForbidden() *SSHUsersCreateForbidden {
	return &SSHUsersCreateForbidden{}
}

/*
SSHUsersCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SSHUsersCreateForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this ssh users create forbidden response has a 2xx status code
func (o *SSHUsersCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create forbidden response has a 3xx status code
func (o *SSHUsersCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create forbidden response has a 4xx status code
func (o *SSHUsersCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create forbidden response has a 5xx status code
func (o *SSHUsersCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create forbidden response a status code equal to that given
func (o *SSHUsersCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SSHUsersCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersCreateForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *SSHUsersCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateNotFound creates a SSHUsersCreateNotFound with default headers values
func NewSSHUsersCreateNotFound() *SSHUsersCreateNotFound {
	return &SSHUsersCreateNotFound{}
}

/*
SSHUsersCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SSHUsersCreateNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this ssh users create not found response has a 2xx status code
func (o *SSHUsersCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create not found response has a 3xx status code
func (o *SSHUsersCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create not found response has a 4xx status code
func (o *SSHUsersCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users create not found response has a 5xx status code
func (o *SSHUsersCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users create not found response a status code equal to that given
func (o *SSHUsersCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SSHUsersCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersCreateNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *SSHUsersCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersCreateInternalServerError creates a SSHUsersCreateInternalServerError with default headers values
func NewSSHUsersCreateInternalServerError() *SSHUsersCreateInternalServerError {
	return &SSHUsersCreateInternalServerError{}
}

/*
SSHUsersCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SSHUsersCreateInternalServerError struct {
}

// IsSuccess returns true when this ssh users create internal server error response has a 2xx status code
func (o *SSHUsersCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users create internal server error response has a 3xx status code
func (o *SSHUsersCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users create internal server error response has a 4xx status code
func (o *SSHUsersCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users create internal server error response has a 5xx status code
func (o *SSHUsersCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ssh users create internal server error response a status code equal to that given
func (o *SSHUsersCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SSHUsersCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateInternalServerError ", 500)
}

func (o *SSHUsersCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/SshUsers/create][%d] sshUsersCreateInternalServerError ", 500)
}

func (o *SSHUsersCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
