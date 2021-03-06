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

// KeycloakLoginReader is a Reader for the KeycloakLogin structure.
type KeycloakLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeycloakLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeycloakLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeycloakLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKeycloakLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKeycloakLoginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeycloakLoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKeycloakLoginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeycloakLoginOK creates a KeycloakLoginOK with default headers values
func NewKeycloakLoginOK() *KeycloakLoginOK {
	return &KeycloakLoginOK{}
}

/* KeycloakLoginOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakLoginOK struct {
	Payload *models.GetToken
}

func (o *KeycloakLoginOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginOK  %+v", 200, o.Payload)
}
func (o *KeycloakLoginOK) GetPayload() *models.GetToken {
	return o.Payload
}

func (o *KeycloakLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginBadRequest creates a KeycloakLoginBadRequest with default headers values
func NewKeycloakLoginBadRequest() *KeycloakLoginBadRequest {
	return &KeycloakLoginBadRequest{}
}

/* KeycloakLoginBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakLoginBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KeycloakLoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginBadRequest  %+v", 400, o.Payload)
}
func (o *KeycloakLoginBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *KeycloakLoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginUnauthorized creates a KeycloakLoginUnauthorized with default headers values
func NewKeycloakLoginUnauthorized() *KeycloakLoginUnauthorized {
	return &KeycloakLoginUnauthorized{}
}

/* KeycloakLoginUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakLoginUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginUnauthorized  %+v", 401, o.Payload)
}
func (o *KeycloakLoginUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginForbidden creates a KeycloakLoginForbidden with default headers values
func NewKeycloakLoginForbidden() *KeycloakLoginForbidden {
	return &KeycloakLoginForbidden{}
}

/* KeycloakLoginForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakLoginForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakLoginForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginForbidden  %+v", 403, o.Payload)
}
func (o *KeycloakLoginForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakLoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginNotFound creates a KeycloakLoginNotFound with default headers values
func NewKeycloakLoginNotFound() *KeycloakLoginNotFound {
	return &KeycloakLoginNotFound{}
}

/* KeycloakLoginNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakLoginNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakLoginNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginNotFound  %+v", 404, o.Payload)
}
func (o *KeycloakLoginNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *KeycloakLoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginInternalServerError creates a KeycloakLoginInternalServerError with default headers values
func NewKeycloakLoginInternalServerError() *KeycloakLoginInternalServerError {
	return &KeycloakLoginInternalServerError{}
}

/* KeycloakLoginInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakLoginInternalServerError struct {
}

func (o *KeycloakLoginInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginInternalServerError ", 500)
}

func (o *KeycloakLoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
