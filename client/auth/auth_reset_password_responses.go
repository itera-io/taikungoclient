// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AuthResetPasswordReader is a Reader for the AuthResetPassword structure.
type AuthResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthResetPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthResetPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthResetPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthResetPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthResetPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthResetPasswordOK creates a AuthResetPasswordOK with default headers values
func NewAuthResetPasswordOK() *AuthResetPasswordOK {
	return &AuthResetPasswordOK{}
}

/*
AuthResetPasswordOK describes a response with status code 200, with default header values.

Success
*/
type AuthResetPasswordOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this auth reset password o k response has a 2xx status code
func (o *AuthResetPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth reset password o k response has a 3xx status code
func (o *AuthResetPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password o k response has a 4xx status code
func (o *AuthResetPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth reset password o k response has a 5xx status code
func (o *AuthResetPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth reset password o k response a status code equal to that given
func (o *AuthResetPasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordOK  %+v", 200, o.Payload)
}

func (o *AuthResetPasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordOK  %+v", 200, o.Payload)
}

func (o *AuthResetPasswordOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AuthResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthResetPasswordBadRequest creates a AuthResetPasswordBadRequest with default headers values
func NewAuthResetPasswordBadRequest() *AuthResetPasswordBadRequest {
	return &AuthResetPasswordBadRequest{}
}

/*
AuthResetPasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthResetPasswordBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this auth reset password bad request response has a 2xx status code
func (o *AuthResetPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth reset password bad request response has a 3xx status code
func (o *AuthResetPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password bad request response has a 4xx status code
func (o *AuthResetPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth reset password bad request response has a 5xx status code
func (o *AuthResetPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this auth reset password bad request response a status code equal to that given
func (o *AuthResetPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuthResetPasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AuthResetPasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AuthResetPasswordBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *AuthResetPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthResetPasswordUnauthorized creates a AuthResetPasswordUnauthorized with default headers values
func NewAuthResetPasswordUnauthorized() *AuthResetPasswordUnauthorized {
	return &AuthResetPasswordUnauthorized{}
}

/*
AuthResetPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthResetPasswordUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth reset password unauthorized response has a 2xx status code
func (o *AuthResetPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth reset password unauthorized response has a 3xx status code
func (o *AuthResetPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password unauthorized response has a 4xx status code
func (o *AuthResetPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth reset password unauthorized response has a 5xx status code
func (o *AuthResetPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this auth reset password unauthorized response a status code equal to that given
func (o *AuthResetPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AuthResetPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthResetPasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthResetPasswordUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthResetPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthResetPasswordForbidden creates a AuthResetPasswordForbidden with default headers values
func NewAuthResetPasswordForbidden() *AuthResetPasswordForbidden {
	return &AuthResetPasswordForbidden{}
}

/*
AuthResetPasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthResetPasswordForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth reset password forbidden response has a 2xx status code
func (o *AuthResetPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth reset password forbidden response has a 3xx status code
func (o *AuthResetPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password forbidden response has a 4xx status code
func (o *AuthResetPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth reset password forbidden response has a 5xx status code
func (o *AuthResetPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this auth reset password forbidden response a status code equal to that given
func (o *AuthResetPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AuthResetPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AuthResetPasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AuthResetPasswordForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthResetPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthResetPasswordNotFound creates a AuthResetPasswordNotFound with default headers values
func NewAuthResetPasswordNotFound() *AuthResetPasswordNotFound {
	return &AuthResetPasswordNotFound{}
}

/*
AuthResetPasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthResetPasswordNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this auth reset password not found response has a 2xx status code
func (o *AuthResetPasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth reset password not found response has a 3xx status code
func (o *AuthResetPasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password not found response has a 4xx status code
func (o *AuthResetPasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth reset password not found response has a 5xx status code
func (o *AuthResetPasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this auth reset password not found response a status code equal to that given
func (o *AuthResetPasswordNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AuthResetPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AuthResetPasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AuthResetPasswordNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthResetPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthResetPasswordInternalServerError creates a AuthResetPasswordInternalServerError with default headers values
func NewAuthResetPasswordInternalServerError() *AuthResetPasswordInternalServerError {
	return &AuthResetPasswordInternalServerError{}
}

/*
AuthResetPasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthResetPasswordInternalServerError struct {
}

// IsSuccess returns true when this auth reset password internal server error response has a 2xx status code
func (o *AuthResetPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth reset password internal server error response has a 3xx status code
func (o *AuthResetPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth reset password internal server error response has a 4xx status code
func (o *AuthResetPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth reset password internal server error response has a 5xx status code
func (o *AuthResetPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this auth reset password internal server error response a status code equal to that given
func (o *AuthResetPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuthResetPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordInternalServerError ", 500)
}

func (o *AuthResetPasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/resetpassword][%d] authResetPasswordInternalServerError ", 500)
}

func (o *AuthResetPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
