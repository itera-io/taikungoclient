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

// AuthForgotPasswordReader is a Reader for the AuthForgotPassword structure.
type AuthForgotPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthForgotPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthForgotPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthForgotPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthForgotPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthForgotPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthForgotPasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthForgotPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthForgotPasswordOK creates a AuthForgotPasswordOK with default headers values
func NewAuthForgotPasswordOK() *AuthForgotPasswordOK {
	return &AuthForgotPasswordOK{}
}

/*
AuthForgotPasswordOK describes a response with status code 200, with default header values.

Success
*/
type AuthForgotPasswordOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this auth forgot password o k response has a 2xx status code
func (o *AuthForgotPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth forgot password o k response has a 3xx status code
func (o *AuthForgotPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password o k response has a 4xx status code
func (o *AuthForgotPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth forgot password o k response has a 5xx status code
func (o *AuthForgotPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth forgot password o k response a status code equal to that given
func (o *AuthForgotPasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthForgotPasswordOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordOK  %+v", 200, o.Payload)
}

func (o *AuthForgotPasswordOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordOK  %+v", 200, o.Payload)
}

func (o *AuthForgotPasswordOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AuthForgotPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthForgotPasswordBadRequest creates a AuthForgotPasswordBadRequest with default headers values
func NewAuthForgotPasswordBadRequest() *AuthForgotPasswordBadRequest {
	return &AuthForgotPasswordBadRequest{}
}

/*
AuthForgotPasswordBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthForgotPasswordBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this auth forgot password bad request response has a 2xx status code
func (o *AuthForgotPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth forgot password bad request response has a 3xx status code
func (o *AuthForgotPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password bad request response has a 4xx status code
func (o *AuthForgotPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth forgot password bad request response has a 5xx status code
func (o *AuthForgotPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this auth forgot password bad request response a status code equal to that given
func (o *AuthForgotPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuthForgotPasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AuthForgotPasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *AuthForgotPasswordBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AuthForgotPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthForgotPasswordUnauthorized creates a AuthForgotPasswordUnauthorized with default headers values
func NewAuthForgotPasswordUnauthorized() *AuthForgotPasswordUnauthorized {
	return &AuthForgotPasswordUnauthorized{}
}

/*
AuthForgotPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthForgotPasswordUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth forgot password unauthorized response has a 2xx status code
func (o *AuthForgotPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth forgot password unauthorized response has a 3xx status code
func (o *AuthForgotPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password unauthorized response has a 4xx status code
func (o *AuthForgotPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth forgot password unauthorized response has a 5xx status code
func (o *AuthForgotPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this auth forgot password unauthorized response a status code equal to that given
func (o *AuthForgotPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AuthForgotPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthForgotPasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthForgotPasswordUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthForgotPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthForgotPasswordForbidden creates a AuthForgotPasswordForbidden with default headers values
func NewAuthForgotPasswordForbidden() *AuthForgotPasswordForbidden {
	return &AuthForgotPasswordForbidden{}
}

/*
AuthForgotPasswordForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthForgotPasswordForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth forgot password forbidden response has a 2xx status code
func (o *AuthForgotPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth forgot password forbidden response has a 3xx status code
func (o *AuthForgotPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password forbidden response has a 4xx status code
func (o *AuthForgotPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth forgot password forbidden response has a 5xx status code
func (o *AuthForgotPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this auth forgot password forbidden response a status code equal to that given
func (o *AuthForgotPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AuthForgotPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AuthForgotPasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordForbidden  %+v", 403, o.Payload)
}

func (o *AuthForgotPasswordForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthForgotPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthForgotPasswordNotFound creates a AuthForgotPasswordNotFound with default headers values
func NewAuthForgotPasswordNotFound() *AuthForgotPasswordNotFound {
	return &AuthForgotPasswordNotFound{}
}

/*
AuthForgotPasswordNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthForgotPasswordNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth forgot password not found response has a 2xx status code
func (o *AuthForgotPasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth forgot password not found response has a 3xx status code
func (o *AuthForgotPasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password not found response has a 4xx status code
func (o *AuthForgotPasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth forgot password not found response has a 5xx status code
func (o *AuthForgotPasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this auth forgot password not found response a status code equal to that given
func (o *AuthForgotPasswordNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AuthForgotPasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AuthForgotPasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordNotFound  %+v", 404, o.Payload)
}

func (o *AuthForgotPasswordNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthForgotPasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthForgotPasswordInternalServerError creates a AuthForgotPasswordInternalServerError with default headers values
func NewAuthForgotPasswordInternalServerError() *AuthForgotPasswordInternalServerError {
	return &AuthForgotPasswordInternalServerError{}
}

/*
AuthForgotPasswordInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthForgotPasswordInternalServerError struct {
}

// IsSuccess returns true when this auth forgot password internal server error response has a 2xx status code
func (o *AuthForgotPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth forgot password internal server error response has a 3xx status code
func (o *AuthForgotPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth forgot password internal server error response has a 4xx status code
func (o *AuthForgotPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth forgot password internal server error response has a 5xx status code
func (o *AuthForgotPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this auth forgot password internal server error response a status code equal to that given
func (o *AuthForgotPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuthForgotPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordInternalServerError ", 500)
}

func (o *AuthForgotPasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/forgotpassword][%d] authForgotPasswordInternalServerError ", 500)
}

func (o *AuthForgotPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
