// Code generated by go-swagger; DO NOT EDIT.

package user_token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UserTokenBindUnbindReader is a Reader for the UserTokenBindUnbind structure.
type UserTokenBindUnbindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserTokenBindUnbindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserTokenBindUnbindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserTokenBindUnbindBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUserTokenBindUnbindUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserTokenBindUnbindForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserTokenBindUnbindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserTokenBindUnbindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUserTokenBindUnbindOK creates a UserTokenBindUnbindOK with default headers values
func NewUserTokenBindUnbindOK() *UserTokenBindUnbindOK {
	return &UserTokenBindUnbindOK{}
}

/*
UserTokenBindUnbindOK describes a response with status code 200, with default header values.

Success
*/
type UserTokenBindUnbindOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this user token bind unbind o k response has a 2xx status code
func (o *UserTokenBindUnbindOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user token bind unbind o k response has a 3xx status code
func (o *UserTokenBindUnbindOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind o k response has a 4xx status code
func (o *UserTokenBindUnbindOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token bind unbind o k response has a 5xx status code
func (o *UserTokenBindUnbindOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user token bind unbind o k response a status code equal to that given
func (o *UserTokenBindUnbindOK) IsCode(code int) bool {
	return code == 200
}

func (o *UserTokenBindUnbindOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindOK  %+v", 200, o.Payload)
}

func (o *UserTokenBindUnbindOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindOK  %+v", 200, o.Payload)
}

func (o *UserTokenBindUnbindOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UserTokenBindUnbindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenBindUnbindBadRequest creates a UserTokenBindUnbindBadRequest with default headers values
func NewUserTokenBindUnbindBadRequest() *UserTokenBindUnbindBadRequest {
	return &UserTokenBindUnbindBadRequest{}
}

/*
UserTokenBindUnbindBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UserTokenBindUnbindBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this user token bind unbind bad request response has a 2xx status code
func (o *UserTokenBindUnbindBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token bind unbind bad request response has a 3xx status code
func (o *UserTokenBindUnbindBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind bad request response has a 4xx status code
func (o *UserTokenBindUnbindBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token bind unbind bad request response has a 5xx status code
func (o *UserTokenBindUnbindBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user token bind unbind bad request response a status code equal to that given
func (o *UserTokenBindUnbindBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UserTokenBindUnbindBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenBindUnbindBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindBadRequest  %+v", 400, o.Payload)
}

func (o *UserTokenBindUnbindBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *UserTokenBindUnbindBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenBindUnbindUnauthorized creates a UserTokenBindUnbindUnauthorized with default headers values
func NewUserTokenBindUnbindUnauthorized() *UserTokenBindUnbindUnauthorized {
	return &UserTokenBindUnbindUnauthorized{}
}

/*
UserTokenBindUnbindUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UserTokenBindUnbindUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user token bind unbind unauthorized response has a 2xx status code
func (o *UserTokenBindUnbindUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token bind unbind unauthorized response has a 3xx status code
func (o *UserTokenBindUnbindUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind unauthorized response has a 4xx status code
func (o *UserTokenBindUnbindUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token bind unbind unauthorized response has a 5xx status code
func (o *UserTokenBindUnbindUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user token bind unbind unauthorized response a status code equal to that given
func (o *UserTokenBindUnbindUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UserTokenBindUnbindUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenBindUnbindUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindUnauthorized  %+v", 401, o.Payload)
}

func (o *UserTokenBindUnbindUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenBindUnbindUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenBindUnbindForbidden creates a UserTokenBindUnbindForbidden with default headers values
func NewUserTokenBindUnbindForbidden() *UserTokenBindUnbindForbidden {
	return &UserTokenBindUnbindForbidden{}
}

/*
UserTokenBindUnbindForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UserTokenBindUnbindForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user token bind unbind forbidden response has a 2xx status code
func (o *UserTokenBindUnbindForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token bind unbind forbidden response has a 3xx status code
func (o *UserTokenBindUnbindForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind forbidden response has a 4xx status code
func (o *UserTokenBindUnbindForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token bind unbind forbidden response has a 5xx status code
func (o *UserTokenBindUnbindForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user token bind unbind forbidden response a status code equal to that given
func (o *UserTokenBindUnbindForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UserTokenBindUnbindForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenBindUnbindForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindForbidden  %+v", 403, o.Payload)
}

func (o *UserTokenBindUnbindForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenBindUnbindForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenBindUnbindNotFound creates a UserTokenBindUnbindNotFound with default headers values
func NewUserTokenBindUnbindNotFound() *UserTokenBindUnbindNotFound {
	return &UserTokenBindUnbindNotFound{}
}

/*
UserTokenBindUnbindNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UserTokenBindUnbindNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this user token bind unbind not found response has a 2xx status code
func (o *UserTokenBindUnbindNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token bind unbind not found response has a 3xx status code
func (o *UserTokenBindUnbindNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind not found response has a 4xx status code
func (o *UserTokenBindUnbindNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user token bind unbind not found response has a 5xx status code
func (o *UserTokenBindUnbindNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user token bind unbind not found response a status code equal to that given
func (o *UserTokenBindUnbindNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UserTokenBindUnbindNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenBindUnbindNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindNotFound  %+v", 404, o.Payload)
}

func (o *UserTokenBindUnbindNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *UserTokenBindUnbindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserTokenBindUnbindInternalServerError creates a UserTokenBindUnbindInternalServerError with default headers values
func NewUserTokenBindUnbindInternalServerError() *UserTokenBindUnbindInternalServerError {
	return &UserTokenBindUnbindInternalServerError{}
}

/*
UserTokenBindUnbindInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UserTokenBindUnbindInternalServerError struct {
}

// IsSuccess returns true when this user token bind unbind internal server error response has a 2xx status code
func (o *UserTokenBindUnbindInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user token bind unbind internal server error response has a 3xx status code
func (o *UserTokenBindUnbindInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user token bind unbind internal server error response has a 4xx status code
func (o *UserTokenBindUnbindInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user token bind unbind internal server error response has a 5xx status code
func (o *UserTokenBindUnbindInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user token bind unbind internal server error response a status code equal to that given
func (o *UserTokenBindUnbindInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UserTokenBindUnbindInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindInternalServerError ", 500)
}

func (o *UserTokenBindUnbindInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/UserToken/bind-unbind][%d] userTokenBindUnbindInternalServerError ", 500)
}

func (o *UserTokenBindUnbindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
