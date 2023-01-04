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

// SSHUsersListReader is a Reader for the SSHUsersList structure.
type SSHUsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSHUsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSHUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSSHUsersListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSSHUsersListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSSHUsersListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSSHUsersListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSSHUsersListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSSHUsersListOK creates a SSHUsersListOK with default headers values
func NewSSHUsersListOK() *SSHUsersListOK {
	return &SSHUsersListOK{}
}

/*
SSHUsersListOK describes a response with status code 200, with default header values.

Success
*/
type SSHUsersListOK struct {
	Payload []*models.SSHUsersListDto
}

// IsSuccess returns true when this ssh users list o k response has a 2xx status code
func (o *SSHUsersListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ssh users list o k response has a 3xx status code
func (o *SSHUsersListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list o k response has a 4xx status code
func (o *SSHUsersListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users list o k response has a 5xx status code
func (o *SSHUsersListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users list o k response a status code equal to that given
func (o *SSHUsersListOK) IsCode(code int) bool {
	return code == 200
}

func (o *SSHUsersListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListOK  %+v", 200, o.Payload)
}

func (o *SSHUsersListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListOK  %+v", 200, o.Payload)
}

func (o *SSHUsersListOK) GetPayload() []*models.SSHUsersListDto {
	return o.Payload
}

func (o *SSHUsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersListBadRequest creates a SSHUsersListBadRequest with default headers values
func NewSSHUsersListBadRequest() *SSHUsersListBadRequest {
	return &SSHUsersListBadRequest{}
}

/*
SSHUsersListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SSHUsersListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ssh users list bad request response has a 2xx status code
func (o *SSHUsersListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users list bad request response has a 3xx status code
func (o *SSHUsersListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list bad request response has a 4xx status code
func (o *SSHUsersListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users list bad request response has a 5xx status code
func (o *SSHUsersListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users list bad request response a status code equal to that given
func (o *SSHUsersListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SSHUsersListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListBadRequest  %+v", 400, o.Payload)
}

func (o *SSHUsersListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *SSHUsersListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersListUnauthorized creates a SSHUsersListUnauthorized with default headers values
func NewSSHUsersListUnauthorized() *SSHUsersListUnauthorized {
	return &SSHUsersListUnauthorized{}
}

/*
SSHUsersListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SSHUsersListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ssh users list unauthorized response has a 2xx status code
func (o *SSHUsersListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users list unauthorized response has a 3xx status code
func (o *SSHUsersListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list unauthorized response has a 4xx status code
func (o *SSHUsersListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users list unauthorized response has a 5xx status code
func (o *SSHUsersListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users list unauthorized response a status code equal to that given
func (o *SSHUsersListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SSHUsersListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListUnauthorized  %+v", 401, o.Payload)
}

func (o *SSHUsersListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SSHUsersListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersListForbidden creates a SSHUsersListForbidden with default headers values
func NewSSHUsersListForbidden() *SSHUsersListForbidden {
	return &SSHUsersListForbidden{}
}

/*
SSHUsersListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SSHUsersListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ssh users list forbidden response has a 2xx status code
func (o *SSHUsersListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users list forbidden response has a 3xx status code
func (o *SSHUsersListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list forbidden response has a 4xx status code
func (o *SSHUsersListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users list forbidden response has a 5xx status code
func (o *SSHUsersListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users list forbidden response a status code equal to that given
func (o *SSHUsersListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SSHUsersListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListForbidden  %+v", 403, o.Payload)
}

func (o *SSHUsersListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SSHUsersListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersListNotFound creates a SSHUsersListNotFound with default headers values
func NewSSHUsersListNotFound() *SSHUsersListNotFound {
	return &SSHUsersListNotFound{}
}

/*
SSHUsersListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SSHUsersListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this ssh users list not found response has a 2xx status code
func (o *SSHUsersListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users list not found response has a 3xx status code
func (o *SSHUsersListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list not found response has a 4xx status code
func (o *SSHUsersListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ssh users list not found response has a 5xx status code
func (o *SSHUsersListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ssh users list not found response a status code equal to that given
func (o *SSHUsersListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *SSHUsersListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListNotFound  %+v", 404, o.Payload)
}

func (o *SSHUsersListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SSHUsersListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSHUsersListInternalServerError creates a SSHUsersListInternalServerError with default headers values
func NewSSHUsersListInternalServerError() *SSHUsersListInternalServerError {
	return &SSHUsersListInternalServerError{}
}

/*
SSHUsersListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SSHUsersListInternalServerError struct {
}

// IsSuccess returns true when this ssh users list internal server error response has a 2xx status code
func (o *SSHUsersListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ssh users list internal server error response has a 3xx status code
func (o *SSHUsersListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ssh users list internal server error response has a 4xx status code
func (o *SSHUsersListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ssh users list internal server error response has a 5xx status code
func (o *SSHUsersListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ssh users list internal server error response a status code equal to that given
func (o *SSHUsersListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SSHUsersListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListInternalServerError ", 500)
}

func (o *SSHUsersListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/SshUsers/list/{accessProfileId}][%d] sshUsersListInternalServerError ", 500)
}

func (o *SSHUsersListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
