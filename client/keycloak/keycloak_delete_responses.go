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

// KeycloakDeleteReader is a Reader for the KeycloakDelete structure.
type KeycloakDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeycloakDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeycloakDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeycloakDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKeycloakDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKeycloakDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeycloakDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKeycloakDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeycloakDeleteOK creates a KeycloakDeleteOK with default headers values
func NewKeycloakDeleteOK() *KeycloakDeleteOK {
	return &KeycloakDeleteOK{}
}

/*
KeycloakDeleteOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this keycloak delete o k response has a 2xx status code
func (o *KeycloakDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keycloak delete o k response has a 3xx status code
func (o *KeycloakDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete o k response has a 4xx status code
func (o *KeycloakDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak delete o k response has a 5xx status code
func (o *KeycloakDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak delete o k response a status code equal to that given
func (o *KeycloakDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *KeycloakDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteOK  %+v", 200, o.Payload)
}

func (o *KeycloakDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteOK  %+v", 200, o.Payload)
}

func (o *KeycloakDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KeycloakDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakDeleteBadRequest creates a KeycloakDeleteBadRequest with default headers values
func NewKeycloakDeleteBadRequest() *KeycloakDeleteBadRequest {
	return &KeycloakDeleteBadRequest{}
}

/*
KeycloakDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakDeleteBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak delete bad request response has a 2xx status code
func (o *KeycloakDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak delete bad request response has a 3xx status code
func (o *KeycloakDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete bad request response has a 4xx status code
func (o *KeycloakDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak delete bad request response has a 5xx status code
func (o *KeycloakDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak delete bad request response a status code equal to that given
func (o *KeycloakDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KeycloakDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakDeleteBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakDeleteUnauthorized creates a KeycloakDeleteUnauthorized with default headers values
func NewKeycloakDeleteUnauthorized() *KeycloakDeleteUnauthorized {
	return &KeycloakDeleteUnauthorized{}
}

/*
KeycloakDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak delete unauthorized response has a 2xx status code
func (o *KeycloakDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak delete unauthorized response has a 3xx status code
func (o *KeycloakDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete unauthorized response has a 4xx status code
func (o *KeycloakDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak delete unauthorized response has a 5xx status code
func (o *KeycloakDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak delete unauthorized response a status code equal to that given
func (o *KeycloakDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KeycloakDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakDeleteForbidden creates a KeycloakDeleteForbidden with default headers values
func NewKeycloakDeleteForbidden() *KeycloakDeleteForbidden {
	return &KeycloakDeleteForbidden{}
}

/*
KeycloakDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakDeleteForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak delete forbidden response has a 2xx status code
func (o *KeycloakDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak delete forbidden response has a 3xx status code
func (o *KeycloakDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete forbidden response has a 4xx status code
func (o *KeycloakDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak delete forbidden response has a 5xx status code
func (o *KeycloakDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak delete forbidden response a status code equal to that given
func (o *KeycloakDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KeycloakDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakDeleteNotFound creates a KeycloakDeleteNotFound with default headers values
func NewKeycloakDeleteNotFound() *KeycloakDeleteNotFound {
	return &KeycloakDeleteNotFound{}
}

/*
KeycloakDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakDeleteNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this keycloak delete not found response has a 2xx status code
func (o *KeycloakDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak delete not found response has a 3xx status code
func (o *KeycloakDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete not found response has a 4xx status code
func (o *KeycloakDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak delete not found response has a 5xx status code
func (o *KeycloakDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak delete not found response a status code equal to that given
func (o *KeycloakDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KeycloakDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakDeleteInternalServerError creates a KeycloakDeleteInternalServerError with default headers values
func NewKeycloakDeleteInternalServerError() *KeycloakDeleteInternalServerError {
	return &KeycloakDeleteInternalServerError{}
}

/*
KeycloakDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakDeleteInternalServerError struct {
}

// IsSuccess returns true when this keycloak delete internal server error response has a 2xx status code
func (o *KeycloakDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak delete internal server error response has a 3xx status code
func (o *KeycloakDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak delete internal server error response has a 4xx status code
func (o *KeycloakDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak delete internal server error response has a 5xx status code
func (o *KeycloakDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this keycloak delete internal server error response a status code equal to that given
func (o *KeycloakDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KeycloakDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteInternalServerError ", 500)
}

func (o *KeycloakDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/delete][%d] keycloakDeleteInternalServerError ", 500)
}

func (o *KeycloakDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
