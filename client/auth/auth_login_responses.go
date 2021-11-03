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

// AuthLoginReader is a Reader for the AuthLogin structure.
type AuthLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthLoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthLoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAuthLoginTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthLoginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthLoginOK creates a AuthLoginOK with default headers values
func NewAuthLoginOK() *AuthLoginOK {
	return &AuthLoginOK{}
}

/* AuthLoginOK describes a response with status code 200, with default header values.

Success
*/
type AuthLoginOK struct {
	Payload *models.GetToken
}

func (o *AuthLoginOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginOK  %+v", 200, o.Payload)
}
func (o *AuthLoginOK) GetPayload() *models.GetToken {
	return o.Payload
}

func (o *AuthLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginBadRequest creates a AuthLoginBadRequest with default headers values
func NewAuthLoginBadRequest() *AuthLoginBadRequest {
	return &AuthLoginBadRequest{}
}

/* AuthLoginBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthLoginBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AuthLoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginBadRequest  %+v", 400, o.Payload)
}
func (o *AuthLoginBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AuthLoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginUnauthorized creates a AuthLoginUnauthorized with default headers values
func NewAuthLoginUnauthorized() *AuthLoginUnauthorized {
	return &AuthLoginUnauthorized{}
}

/* AuthLoginUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthLoginUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AuthLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginUnauthorized  %+v", 401, o.Payload)
}
func (o *AuthLoginUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginForbidden creates a AuthLoginForbidden with default headers values
func NewAuthLoginForbidden() *AuthLoginForbidden {
	return &AuthLoginForbidden{}
}

/* AuthLoginForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthLoginForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AuthLoginForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginForbidden  %+v", 403, o.Payload)
}
func (o *AuthLoginForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthLoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginNotFound creates a AuthLoginNotFound with default headers values
func NewAuthLoginNotFound() *AuthLoginNotFound {
	return &AuthLoginNotFound{}
}

/* AuthLoginNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthLoginNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AuthLoginNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginNotFound  %+v", 404, o.Payload)
}
func (o *AuthLoginNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthLoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginTooManyRequests creates a AuthLoginTooManyRequests with default headers values
func NewAuthLoginTooManyRequests() *AuthLoginTooManyRequests {
	return &AuthLoginTooManyRequests{}
}

/* AuthLoginTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AuthLoginTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AuthLoginTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginTooManyRequests  %+v", 429, o.Payload)
}
func (o *AuthLoginTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthLoginTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthLoginInternalServerError creates a AuthLoginInternalServerError with default headers values
func NewAuthLoginInternalServerError() *AuthLoginInternalServerError {
	return &AuthLoginInternalServerError{}
}

/* AuthLoginInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthLoginInternalServerError struct {
}

func (o *AuthLoginInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/login][%d] authLoginInternalServerError ", 500)
}

func (o *AuthLoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
