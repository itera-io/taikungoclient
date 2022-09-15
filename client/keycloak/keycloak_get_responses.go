// Code generated by go-swagger; DO NOT EDIT.

package keycloak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// KeycloakGetReader is a Reader for the KeycloakGet structure.
type KeycloakGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeycloakGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeycloakGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeycloakGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKeycloakGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKeycloakGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeycloakGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKeycloakGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeycloakGetOK creates a KeycloakGetOK with default headers values
func NewKeycloakGetOK() *KeycloakGetOK {
	return &KeycloakGetOK{}
}

/*
KeycloakGetOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakGetOK struct {
	Payload *models.KeycloakListDto
}

// IsSuccess returns true when this keycloak get o k response has a 2xx status code
func (o *KeycloakGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keycloak get o k response has a 3xx status code
func (o *KeycloakGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get o k response has a 4xx status code
func (o *KeycloakGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak get o k response has a 5xx status code
func (o *KeycloakGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak get o k response a status code equal to that given
func (o *KeycloakGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *KeycloakGetOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetOK  %+v", 200, o.Payload)
}

func (o *KeycloakGetOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetOK  %+v", 200, o.Payload)
}

func (o *KeycloakGetOK) GetPayload() *models.KeycloakListDto {
	return o.Payload
}

func (o *KeycloakGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeycloakListDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakGetBadRequest creates a KeycloakGetBadRequest with default headers values
func NewKeycloakGetBadRequest() *KeycloakGetBadRequest {
	return &KeycloakGetBadRequest{}
}

/*
KeycloakGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakGetBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this keycloak get bad request response has a 2xx status code
func (o *KeycloakGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak get bad request response has a 3xx status code
func (o *KeycloakGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get bad request response has a 4xx status code
func (o *KeycloakGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak get bad request response has a 5xx status code
func (o *KeycloakGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak get bad request response a status code equal to that given
func (o *KeycloakGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KeycloakGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakGetBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakGetBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KeycloakGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakGetUnauthorized creates a KeycloakGetUnauthorized with default headers values
func NewKeycloakGetUnauthorized() *KeycloakGetUnauthorized {
	return &KeycloakGetUnauthorized{}
}

/*
KeycloakGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakGetUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak get unauthorized response has a 2xx status code
func (o *KeycloakGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak get unauthorized response has a 3xx status code
func (o *KeycloakGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get unauthorized response has a 4xx status code
func (o *KeycloakGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak get unauthorized response has a 5xx status code
func (o *KeycloakGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak get unauthorized response a status code equal to that given
func (o *KeycloakGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KeycloakGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakGetUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakGetForbidden creates a KeycloakGetForbidden with default headers values
func NewKeycloakGetForbidden() *KeycloakGetForbidden {
	return &KeycloakGetForbidden{}
}

/*
KeycloakGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakGetForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak get forbidden response has a 2xx status code
func (o *KeycloakGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak get forbidden response has a 3xx status code
func (o *KeycloakGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get forbidden response has a 4xx status code
func (o *KeycloakGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak get forbidden response has a 5xx status code
func (o *KeycloakGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak get forbidden response a status code equal to that given
func (o *KeycloakGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KeycloakGetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakGetForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakGetForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakGetNotFound creates a KeycloakGetNotFound with default headers values
func NewKeycloakGetNotFound() *KeycloakGetNotFound {
	return &KeycloakGetNotFound{}
}

/*
KeycloakGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakGetNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak get not found response has a 2xx status code
func (o *KeycloakGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak get not found response has a 3xx status code
func (o *KeycloakGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get not found response has a 4xx status code
func (o *KeycloakGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak get not found response has a 5xx status code
func (o *KeycloakGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak get not found response a status code equal to that given
func (o *KeycloakGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KeycloakGetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakGetNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakGetNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakGetInternalServerError creates a KeycloakGetInternalServerError with default headers values
func NewKeycloakGetInternalServerError() *KeycloakGetInternalServerError {
	return &KeycloakGetInternalServerError{}
}

/*
KeycloakGetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakGetInternalServerError struct {
}

// IsSuccess returns true when this keycloak get internal server error response has a 2xx status code
func (o *KeycloakGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak get internal server error response has a 3xx status code
func (o *KeycloakGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak get internal server error response has a 4xx status code
func (o *KeycloakGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak get internal server error response has a 5xx status code
func (o *KeycloakGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this keycloak get internal server error response a status code equal to that given
func (o *KeycloakGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KeycloakGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetInternalServerError ", 500)
}

func (o *KeycloakGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetInternalServerError ", 500)
}

func (o *KeycloakGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
