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

/*
AuthRegisterOK describes a response with status code 200, with default header values.

Success
*/
type AuthRegisterOK struct {
}

// IsSuccess returns true when this auth register o k response has a 2xx status code
func (o *AuthRegisterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this auth register o k response has a 3xx status code
func (o *AuthRegisterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register o k response has a 4xx status code
func (o *AuthRegisterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth register o k response has a 5xx status code
func (o *AuthRegisterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this auth register o k response a status code equal to that given
func (o *AuthRegisterOK) IsCode(code int) bool {
	return code == 200
}

func (o *AuthRegisterOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterOK ", 200)
}

func (o *AuthRegisterOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterOK ", 200)
}

func (o *AuthRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuthRegisterBadRequest creates a AuthRegisterBadRequest with default headers values
func NewAuthRegisterBadRequest() *AuthRegisterBadRequest {
	return &AuthRegisterBadRequest{}
}

/*
AuthRegisterBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AuthRegisterBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this auth register bad request response has a 2xx status code
func (o *AuthRegisterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth register bad request response has a 3xx status code
func (o *AuthRegisterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register bad request response has a 4xx status code
func (o *AuthRegisterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth register bad request response has a 5xx status code
func (o *AuthRegisterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this auth register bad request response a status code equal to that given
func (o *AuthRegisterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AuthRegisterBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterBadRequest  %+v", 400, o.Payload)
}

func (o *AuthRegisterBadRequest) String() string {
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

/*
AuthRegisterUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AuthRegisterUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth register unauthorized response has a 2xx status code
func (o *AuthRegisterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth register unauthorized response has a 3xx status code
func (o *AuthRegisterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register unauthorized response has a 4xx status code
func (o *AuthRegisterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth register unauthorized response has a 5xx status code
func (o *AuthRegisterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this auth register unauthorized response a status code equal to that given
func (o *AuthRegisterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AuthRegisterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthRegisterUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterUnauthorized  %+v", 401, o.Payload)
}

func (o *AuthRegisterUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthRegisterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterForbidden creates a AuthRegisterForbidden with default headers values
func NewAuthRegisterForbidden() *AuthRegisterForbidden {
	return &AuthRegisterForbidden{}
}

/*
AuthRegisterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AuthRegisterForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth register forbidden response has a 2xx status code
func (o *AuthRegisterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth register forbidden response has a 3xx status code
func (o *AuthRegisterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register forbidden response has a 4xx status code
func (o *AuthRegisterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth register forbidden response has a 5xx status code
func (o *AuthRegisterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this auth register forbidden response a status code equal to that given
func (o *AuthRegisterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AuthRegisterForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterForbidden  %+v", 403, o.Payload)
}

func (o *AuthRegisterForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterForbidden  %+v", 403, o.Payload)
}

func (o *AuthRegisterForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthRegisterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterNotFound creates a AuthRegisterNotFound with default headers values
func NewAuthRegisterNotFound() *AuthRegisterNotFound {
	return &AuthRegisterNotFound{}
}

/*
AuthRegisterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AuthRegisterNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this auth register not found response has a 2xx status code
func (o *AuthRegisterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth register not found response has a 3xx status code
func (o *AuthRegisterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register not found response has a 4xx status code
func (o *AuthRegisterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this auth register not found response has a 5xx status code
func (o *AuthRegisterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this auth register not found response a status code equal to that given
func (o *AuthRegisterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AuthRegisterNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterNotFound  %+v", 404, o.Payload)
}

func (o *AuthRegisterNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterNotFound  %+v", 404, o.Payload)
}

func (o *AuthRegisterNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AuthRegisterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthRegisterInternalServerError creates a AuthRegisterInternalServerError with default headers values
func NewAuthRegisterInternalServerError() *AuthRegisterInternalServerError {
	return &AuthRegisterInternalServerError{}
}

/*
AuthRegisterInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AuthRegisterInternalServerError struct {
}

// IsSuccess returns true when this auth register internal server error response has a 2xx status code
func (o *AuthRegisterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this auth register internal server error response has a 3xx status code
func (o *AuthRegisterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this auth register internal server error response has a 4xx status code
func (o *AuthRegisterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this auth register internal server error response has a 5xx status code
func (o *AuthRegisterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this auth register internal server error response a status code equal to that given
func (o *AuthRegisterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AuthRegisterInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterInternalServerError ", 500)
}

func (o *AuthRegisterInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Auth/register][%d] authRegisterInternalServerError ", 500)
}

func (o *AuthRegisterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
