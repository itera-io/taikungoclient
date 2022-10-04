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

/*
KeycloakLoginOK describes a response with status code 200, with default header values.

Success
*/
type KeycloakLoginOK struct {
	Payload *models.GetToken
}

// IsSuccess returns true when this keycloak login o k response has a 2xx status code
func (o *KeycloakLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keycloak login o k response has a 3xx status code
func (o *KeycloakLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login o k response has a 4xx status code
func (o *KeycloakLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak login o k response has a 5xx status code
func (o *KeycloakLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak login o k response a status code equal to that given
func (o *KeycloakLoginOK) IsCode(code int) bool {
	return code == 200
}

func (o *KeycloakLoginOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginOK  %+v", 200, o.Payload)
}

func (o *KeycloakLoginOK) String() string {
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

/*
KeycloakLoginBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type KeycloakLoginBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak login bad request response has a 2xx status code
func (o *KeycloakLoginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak login bad request response has a 3xx status code
func (o *KeycloakLoginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login bad request response has a 4xx status code
func (o *KeycloakLoginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak login bad request response has a 5xx status code
func (o *KeycloakLoginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak login bad request response a status code equal to that given
func (o *KeycloakLoginBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *KeycloakLoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakLoginBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginBadRequest  %+v", 400, o.Payload)
}

func (o *KeycloakLoginBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakLoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginUnauthorized creates a KeycloakLoginUnauthorized with default headers values
func NewKeycloakLoginUnauthorized() *KeycloakLoginUnauthorized {
	return &KeycloakLoginUnauthorized{}
}

/*
KeycloakLoginUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type KeycloakLoginUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak login unauthorized response has a 2xx status code
func (o *KeycloakLoginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak login unauthorized response has a 3xx status code
func (o *KeycloakLoginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login unauthorized response has a 4xx status code
func (o *KeycloakLoginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak login unauthorized response has a 5xx status code
func (o *KeycloakLoginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak login unauthorized response a status code equal to that given
func (o *KeycloakLoginUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *KeycloakLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakLoginUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginUnauthorized  %+v", 401, o.Payload)
}

func (o *KeycloakLoginUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginForbidden creates a KeycloakLoginForbidden with default headers values
func NewKeycloakLoginForbidden() *KeycloakLoginForbidden {
	return &KeycloakLoginForbidden{}
}

/*
KeycloakLoginForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type KeycloakLoginForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak login forbidden response has a 2xx status code
func (o *KeycloakLoginForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak login forbidden response has a 3xx status code
func (o *KeycloakLoginForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login forbidden response has a 4xx status code
func (o *KeycloakLoginForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak login forbidden response has a 5xx status code
func (o *KeycloakLoginForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak login forbidden response a status code equal to that given
func (o *KeycloakLoginForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *KeycloakLoginForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakLoginForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginForbidden  %+v", 403, o.Payload)
}

func (o *KeycloakLoginForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakLoginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginNotFound creates a KeycloakLoginNotFound with default headers values
func NewKeycloakLoginNotFound() *KeycloakLoginNotFound {
	return &KeycloakLoginNotFound{}
}

/*
KeycloakLoginNotFound describes a response with status code 404, with default header values.

Not Found
*/
type KeycloakLoginNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this keycloak login not found response has a 2xx status code
func (o *KeycloakLoginNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak login not found response has a 3xx status code
func (o *KeycloakLoginNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login not found response has a 4xx status code
func (o *KeycloakLoginNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this keycloak login not found response has a 5xx status code
func (o *KeycloakLoginNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this keycloak login not found response a status code equal to that given
func (o *KeycloakLoginNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *KeycloakLoginNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakLoginNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginNotFound  %+v", 404, o.Payload)
}

func (o *KeycloakLoginNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *KeycloakLoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKeycloakLoginInternalServerError creates a KeycloakLoginInternalServerError with default headers values
func NewKeycloakLoginInternalServerError() *KeycloakLoginInternalServerError {
	return &KeycloakLoginInternalServerError{}
}

/*
KeycloakLoginInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type KeycloakLoginInternalServerError struct {
}

// IsSuccess returns true when this keycloak login internal server error response has a 2xx status code
func (o *KeycloakLoginInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this keycloak login internal server error response has a 3xx status code
func (o *KeycloakLoginInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keycloak login internal server error response has a 4xx status code
func (o *KeycloakLoginInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this keycloak login internal server error response has a 5xx status code
func (o *KeycloakLoginInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this keycloak login internal server error response a status code equal to that given
func (o *KeycloakLoginInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *KeycloakLoginInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginInternalServerError ", 500)
}

func (o *KeycloakLoginInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Keycloak/login][%d] keycloakLoginInternalServerError ", 500)
}

func (o *KeycloakLoginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
