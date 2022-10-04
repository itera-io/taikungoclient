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

// KeycloakCreateReader is a Reader for the KeycloakCreate structure.
type KeycloakCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeycloakCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeycloakCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewKeycloakCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewKeycloakCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewKeycloakCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewKeycloakCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewKeycloakCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewKeycloakCreateOK creates a KeycloakCreateOK with default headers values
func NewKeycloakCreateOK() *KeycloakCreateOK {
	return &KeycloakCreateOK{}
}

/*
KeycloakCreateOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this keycloak create o k response has a 2xx status code
func (o *KeycloakCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keycloak create o k response has a 3xx status code
func (o *KeycloakCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create o k response has a 4xx status code
func (o *KeycloakCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak create o k response has a 5xx status code
func (o *KeycloakCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak create o k response a status code equal to that given
func (o *KeycloakCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *KeycloakCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateOK  %+v", 200, o.Payload)
}

func (o *KeycloakCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateOK  %+v", 200, o.Payload)
}

func (o *KeycloakCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *KeycloakCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakCreateBadRequest creates a KeycloakCreateBadRequest with default headers values
func NewKeycloakCreateBadRequest() *KeycloakCreateBadRequest {
	return &KeycloakCreateBadRequest{}
}

/*
KeycloakCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakCreateBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak create bad request response has a 2xx status code
func (o *KeycloakCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak create bad request response has a 3xx status code
func (o *KeycloakCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create bad request response has a 4xx status code
func (o *KeycloakCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak create bad request response has a 5xx status code
func (o *KeycloakCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak create bad request response a status code equal to that given
func (o *KeycloakCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KeycloakCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakCreateBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakCreateUnauthorized creates a KeycloakCreateUnauthorized with default headers values
func NewKeycloakCreateUnauthorized() *KeycloakCreateUnauthorized {
	return &KeycloakCreateUnauthorized{}
}

/*
KeycloakCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakCreateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak create unauthorized response has a 2xx status code
func (o *KeycloakCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak create unauthorized response has a 3xx status code
func (o *KeycloakCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create unauthorized response has a 4xx status code
func (o *KeycloakCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak create unauthorized response has a 5xx status code
func (o *KeycloakCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak create unauthorized response a status code equal to that given
func (o *KeycloakCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KeycloakCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakCreateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakCreateForbidden creates a KeycloakCreateForbidden with default headers values
func NewKeycloakCreateForbidden() *KeycloakCreateForbidden {
	return &KeycloakCreateForbidden{}
}

/*
KeycloakCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakCreateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak create forbidden response has a 2xx status code
func (o *KeycloakCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak create forbidden response has a 3xx status code
func (o *KeycloakCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create forbidden response has a 4xx status code
func (o *KeycloakCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak create forbidden response has a 5xx status code
func (o *KeycloakCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak create forbidden response a status code equal to that given
func (o *KeycloakCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KeycloakCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakCreateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakCreateNotFound creates a KeycloakCreateNotFound with default headers values
func NewKeycloakCreateNotFound() *KeycloakCreateNotFound {
	return &KeycloakCreateNotFound{}
}

/*
KeycloakCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakCreateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak create not found response has a 2xx status code
func (o *KeycloakCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak create not found response has a 3xx status code
func (o *KeycloakCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create not found response has a 4xx status code
func (o *KeycloakCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak create not found response has a 5xx status code
func (o *KeycloakCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak create not found response a status code equal to that given
func (o *KeycloakCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KeycloakCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakCreateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakCreateInternalServerError creates a KeycloakCreateInternalServerError with default headers values
func NewKeycloakCreateInternalServerError() *KeycloakCreateInternalServerError {
	return &KeycloakCreateInternalServerError{}
}

/*
KeycloakCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakCreateInternalServerError struct {
}

// IsSuccess returns true when this keycloak create internal server error response has a 2xx status code
func (o *KeycloakCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak create internal server error response has a 3xx status code
func (o *KeycloakCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak create internal server error response has a 4xx status code
func (o *KeycloakCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak create internal server error response has a 5xx status code
func (o *KeycloakCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this keycloak create internal server error response a status code equal to that given
func (o *KeycloakCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KeycloakCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateInternalServerError ", 500)
}

func (o *KeycloakCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/create][%d] keycloakCreateInternalServerError ", 500)
}

func (o *KeycloakCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
