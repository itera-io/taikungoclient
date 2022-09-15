// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OrganizationsToggleKeycloakReader is a Reader for the OrganizationsToggleKeycloak structure.
type OrganizationsToggleKeycloakReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationsToggleKeycloakReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationsToggleKeycloakOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrganizationsToggleKeycloakBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrganizationsToggleKeycloakUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOrganizationsToggleKeycloakForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrganizationsToggleKeycloakNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrganizationsToggleKeycloakInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOrganizationsToggleKeycloakOK creates a OrganizationsToggleKeycloakOK with default headers values
func NewOrganizationsToggleKeycloakOK() *OrganizationsToggleKeycloakOK {
	return &OrganizationsToggleKeycloakOK{}
}

/*
OrganizationsToggleKeycloakOK describes a response with status code 200, with default header values.

Success
*/
type OrganizationsToggleKeycloakOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this organizations toggle keycloak o k response has a 2xx status code
func (o *OrganizationsToggleKeycloakOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organizations toggle keycloak o k response has a 3xx status code
func (o *OrganizationsToggleKeycloakOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak o k response has a 4xx status code
func (o *OrganizationsToggleKeycloakOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations toggle keycloak o k response has a 5xx status code
func (o *OrganizationsToggleKeycloakOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations toggle keycloak o k response a status code equal to that given
func (o *OrganizationsToggleKeycloakOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationsToggleKeycloakOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakOK  %+v", 200, o.Payload)
}

func (o *OrganizationsToggleKeycloakOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakOK  %+v", 200, o.Payload)
}

func (o *OrganizationsToggleKeycloakOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *OrganizationsToggleKeycloakOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsToggleKeycloakBadRequest creates a OrganizationsToggleKeycloakBadRequest with default headers values
func NewOrganizationsToggleKeycloakBadRequest() *OrganizationsToggleKeycloakBadRequest {
	return &OrganizationsToggleKeycloakBadRequest{}
}

/*
OrganizationsToggleKeycloakBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OrganizationsToggleKeycloakBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this organizations toggle keycloak bad request response has a 2xx status code
func (o *OrganizationsToggleKeycloakBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations toggle keycloak bad request response has a 3xx status code
func (o *OrganizationsToggleKeycloakBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak bad request response has a 4xx status code
func (o *OrganizationsToggleKeycloakBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations toggle keycloak bad request response has a 5xx status code
func (o *OrganizationsToggleKeycloakBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations toggle keycloak bad request response a status code equal to that given
func (o *OrganizationsToggleKeycloakBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OrganizationsToggleKeycloakBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsToggleKeycloakBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakBadRequest  %+v", 400, o.Payload)
}

func (o *OrganizationsToggleKeycloakBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *OrganizationsToggleKeycloakBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsToggleKeycloakUnauthorized creates a OrganizationsToggleKeycloakUnauthorized with default headers values
func NewOrganizationsToggleKeycloakUnauthorized() *OrganizationsToggleKeycloakUnauthorized {
	return &OrganizationsToggleKeycloakUnauthorized{}
}

/*
OrganizationsToggleKeycloakUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OrganizationsToggleKeycloakUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations toggle keycloak unauthorized response has a 2xx status code
func (o *OrganizationsToggleKeycloakUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations toggle keycloak unauthorized response has a 3xx status code
func (o *OrganizationsToggleKeycloakUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak unauthorized response has a 4xx status code
func (o *OrganizationsToggleKeycloakUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations toggle keycloak unauthorized response has a 5xx status code
func (o *OrganizationsToggleKeycloakUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations toggle keycloak unauthorized response a status code equal to that given
func (o *OrganizationsToggleKeycloakUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OrganizationsToggleKeycloakUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsToggleKeycloakUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakUnauthorized  %+v", 401, o.Payload)
}

func (o *OrganizationsToggleKeycloakUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsToggleKeycloakUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsToggleKeycloakForbidden creates a OrganizationsToggleKeycloakForbidden with default headers values
func NewOrganizationsToggleKeycloakForbidden() *OrganizationsToggleKeycloakForbidden {
	return &OrganizationsToggleKeycloakForbidden{}
}

/*
OrganizationsToggleKeycloakForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OrganizationsToggleKeycloakForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations toggle keycloak forbidden response has a 2xx status code
func (o *OrganizationsToggleKeycloakForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations toggle keycloak forbidden response has a 3xx status code
func (o *OrganizationsToggleKeycloakForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak forbidden response has a 4xx status code
func (o *OrganizationsToggleKeycloakForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations toggle keycloak forbidden response has a 5xx status code
func (o *OrganizationsToggleKeycloakForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations toggle keycloak forbidden response a status code equal to that given
func (o *OrganizationsToggleKeycloakForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OrganizationsToggleKeycloakForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsToggleKeycloakForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakForbidden  %+v", 403, o.Payload)
}

func (o *OrganizationsToggleKeycloakForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsToggleKeycloakForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsToggleKeycloakNotFound creates a OrganizationsToggleKeycloakNotFound with default headers values
func NewOrganizationsToggleKeycloakNotFound() *OrganizationsToggleKeycloakNotFound {
	return &OrganizationsToggleKeycloakNotFound{}
}

/*
OrganizationsToggleKeycloakNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OrganizationsToggleKeycloakNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this organizations toggle keycloak not found response has a 2xx status code
func (o *OrganizationsToggleKeycloakNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations toggle keycloak not found response has a 3xx status code
func (o *OrganizationsToggleKeycloakNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak not found response has a 4xx status code
func (o *OrganizationsToggleKeycloakNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this organizations toggle keycloak not found response has a 5xx status code
func (o *OrganizationsToggleKeycloakNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this organizations toggle keycloak not found response a status code equal to that given
func (o *OrganizationsToggleKeycloakNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OrganizationsToggleKeycloakNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsToggleKeycloakNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakNotFound  %+v", 404, o.Payload)
}

func (o *OrganizationsToggleKeycloakNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OrganizationsToggleKeycloakNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationsToggleKeycloakInternalServerError creates a OrganizationsToggleKeycloakInternalServerError with default headers values
func NewOrganizationsToggleKeycloakInternalServerError() *OrganizationsToggleKeycloakInternalServerError {
	return &OrganizationsToggleKeycloakInternalServerError{}
}

/*
OrganizationsToggleKeycloakInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OrganizationsToggleKeycloakInternalServerError struct {
}

// IsSuccess returns true when this organizations toggle keycloak internal server error response has a 2xx status code
func (o *OrganizationsToggleKeycloakInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this organizations toggle keycloak internal server error response has a 3xx status code
func (o *OrganizationsToggleKeycloakInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organizations toggle keycloak internal server error response has a 4xx status code
func (o *OrganizationsToggleKeycloakInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this organizations toggle keycloak internal server error response has a 5xx status code
func (o *OrganizationsToggleKeycloakInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this organizations toggle keycloak internal server error response a status code equal to that given
func (o *OrganizationsToggleKeycloakInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OrganizationsToggleKeycloakInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakInternalServerError ", 500)
}

func (o *OrganizationsToggleKeycloakInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Organizations/toggle/keycloak][%d] organizationsToggleKeycloakInternalServerError ", 500)
}

func (o *OrganizationsToggleKeycloakInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
