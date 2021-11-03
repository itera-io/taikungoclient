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

// AuthRegisterReader is a Reader for the AuthRegister structure.
type AuthRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAuthRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAuthRegisterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAuthRegisterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAuthRegisterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAuthRegisterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAuthRegisterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAuthRegisterOK creates a AuthRegisterOK with default headers values
func NewAuthRegisterOK() *AuthRegisterOK {
	return &AuthRegisterOK{}
}

/* AuthRegisterOK describes a response with status code 200, with default header values.

Success
*/
type AuthRegisterOK struct {
}

func (o *AuthRegisterOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterOK ", 200)
}

func (o *AuthRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthRegisterBadRequest creates a AuthRegisterBadRequest with default headers values
func NewAuthRegisterBadRequest() *AuthRegisterBadRequest {
	return &AuthRegisterBadRequest{}
}

/* AuthRegisterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthRegisterBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *AuthRegisterBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterBadRequest  %+v", 400, o.Payload)
}
func (o *AuthRegisterBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AuthRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterUnauthorized creates a AuthRegisterUnauthorized with default headers values
func NewAuthRegisterUnauthorized() *AuthRegisterUnauthorized {
	return &AuthRegisterUnauthorized{}
}

/* AuthRegisterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthRegisterUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *AuthRegisterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterUnauthorized  %+v", 401, o.Payload)
}
func (o *AuthRegisterUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRegisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterForbidden creates a AuthRegisterForbidden with default headers values
func NewAuthRegisterForbidden() *AuthRegisterForbidden {
	return &AuthRegisterForbidden{}
}

/* AuthRegisterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthRegisterForbidden struct {
	Payload *models.ProblemDetails
}

func (o *AuthRegisterForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterForbidden  %+v", 403, o.Payload)
}
func (o *AuthRegisterForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRegisterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterNotFound creates a AuthRegisterNotFound with default headers values
func NewAuthRegisterNotFound() *AuthRegisterNotFound {
	return &AuthRegisterNotFound{}
}

/* AuthRegisterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthRegisterNotFound struct {
	Payload *models.ProblemDetails
}

func (o *AuthRegisterNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterNotFound  %+v", 404, o.Payload)
}
func (o *AuthRegisterNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRegisterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterTooManyRequests creates a AuthRegisterTooManyRequests with default headers values
func NewAuthRegisterTooManyRequests() *AuthRegisterTooManyRequests {
	return &AuthRegisterTooManyRequests{}
}

/* AuthRegisterTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type AuthRegisterTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *AuthRegisterTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterTooManyRequests  %+v", 429, o.Payload)
}
func (o *AuthRegisterTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AuthRegisterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterInternalServerError creates a AuthRegisterInternalServerError with default headers values
func NewAuthRegisterInternalServerError() *AuthRegisterInternalServerError {
	return &AuthRegisterInternalServerError{}
}

/* AuthRegisterInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthRegisterInternalServerError struct {
}

func (o *AuthRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterInternalServerError ", 500)
}

func (o *AuthRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
