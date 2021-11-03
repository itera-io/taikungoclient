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

/* KeycloakGetOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakGetOK struct {
	Payload *models.KeycloakListDto
}

func (o *KeycloakGetOK) Error() string {
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

/* KeycloakGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakGetBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *KeycloakGetBadRequest) Error() string {
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

/* KeycloakGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakGetUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakGetUnauthorized) Error() string {
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

/* KeycloakGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakGetForbidden struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakGetForbidden) Error() string {
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

/* KeycloakGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakGetNotFound struct {
	Payload *models.ProblemDetails
}

func (o *KeycloakGetNotFound) Error() string {
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

/* KeycloakGetInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakGetInternalServerError struct {
}

func (o *KeycloakGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Keycloak][%d] keycloakGetInternalServerError ", 500)
}

func (o *KeycloakGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
