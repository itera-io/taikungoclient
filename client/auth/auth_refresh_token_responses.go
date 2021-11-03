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

// AuthRefreshTokenReader is a Reader for the AuthRefreshToken structure.
type AuthRefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthRefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthRefreshTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthRefreshTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthRefreshTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthRefreshTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAuthRefreshTokenTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthRefreshTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthRefreshTokenOK creates a AuthRefreshTokenOK with default headers values
func NewAuthRefreshTokenOK() *AuthRefreshTokenOK {
	return &AuthRefreshTokenOK{}
}

/* AuthRefreshTokenOK describes a response with status code 200, with default header values.

Success
*/
type AuthRefreshTokenOK struct {
	Payload *models.GetToken
}

func (o *AuthRefreshTokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenOK  %+v", 200, o.Payload)
}
func (o *AuthRefreshTokenOK) GetPayload() *models.GetToken {
	return o.Payload
}

func (o *AuthRefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenBadRequest creates a AuthRefreshTokenBadRequest with default headers values
func NewAuthRefreshTokenBadRequest() *AuthRefreshTokenBadRequest {
	return &AuthRefreshTokenBadRequest{}
}

/* AuthRefreshTokenBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthRefreshTokenBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AuthRefreshTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenBadRequest  %+v", 400, o.Payload)
}
func (o *AuthRefreshTokenBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AuthRefreshTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenUnauthorized creates a AuthRefreshTokenUnauthorized with default headers values
func NewAuthRefreshTokenUnauthorized() *AuthRefreshTokenUnauthorized {
	return &AuthRefreshTokenUnauthorized{}
}

/* AuthRefreshTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthRefreshTokenUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AuthRefreshTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenUnauthorized  %+v", 401, o.Payload)
}
func (o *AuthRefreshTokenUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRefreshTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenForbidden creates a AuthRefreshTokenForbidden with default headers values
func NewAuthRefreshTokenForbidden() *AuthRefreshTokenForbidden {
	return &AuthRefreshTokenForbidden{}
}

/* AuthRefreshTokenForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthRefreshTokenForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AuthRefreshTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenForbidden  %+v", 403, o.Payload)
}
func (o *AuthRefreshTokenForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRefreshTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenNotFound creates a AuthRefreshTokenNotFound with default headers values
func NewAuthRefreshTokenNotFound() *AuthRefreshTokenNotFound {
	return &AuthRefreshTokenNotFound{}
}

/* AuthRefreshTokenNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthRefreshTokenNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AuthRefreshTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenNotFound  %+v", 404, o.Payload)
}
func (o *AuthRefreshTokenNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRefreshTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenTooManyRequests creates a AuthRefreshTokenTooManyRequests with default headers values
func NewAuthRefreshTokenTooManyRequests() *AuthRefreshTokenTooManyRequests {
	return &AuthRefreshTokenTooManyRequests{}
}

/* AuthRefreshTokenTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AuthRefreshTokenTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AuthRefreshTokenTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenTooManyRequests  %+v", 429, o.Payload)
}
func (o *AuthRefreshTokenTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRefreshTokenTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRefreshTokenInternalServerError creates a AuthRefreshTokenInternalServerError with default headers values
func NewAuthRefreshTokenInternalServerError() *AuthRefreshTokenInternalServerError {
	return &AuthRefreshTokenInternalServerError{}
}

/* AuthRefreshTokenInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthRefreshTokenInternalServerError struct {
}

func (o *AuthRefreshTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/refresh][%d] authRefreshTokenInternalServerError ", 500)
}

func (o *AuthRefreshTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
