// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AdminKeycloakListReader is a Reader for the AdminKeycloakList structure.
type AdminKeycloakListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminKeycloakListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminKeycloakListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminKeycloakListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminKeycloakListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminKeycloakListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAdminKeycloakListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminKeycloakListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAdminKeycloakListOK creates a AdminKeycloakListOK with default headers values
func NewAdminKeycloakListOK() *AdminKeycloakListOK {
	return &AdminKeycloakListOK{}
}

/*
AdminKeycloakListOK describes a response with status code 200, with default header values.

Success
*/
type AdminKeycloakListOK struct {
	Payload *models.AdminKeycloakList
}

// IsSuccess returns true when this admin keycloak list o k response has a 2xx status code
func (o *AdminKeycloakListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin keycloak list o k response has a 3xx status code
func (o *AdminKeycloakListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list o k response has a 4xx status code
func (o *AdminKeycloakListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin keycloak list o k response has a 5xx status code
func (o *AdminKeycloakListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin keycloak list o k response a status code equal to that given
func (o *AdminKeycloakListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AdminKeycloakListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListOK  %+v", 200, o.Payload)
}

func (o *AdminKeycloakListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListOK  %+v", 200, o.Payload)
}

func (o *AdminKeycloakListOK) GetPayload() *models.AdminKeycloakList {
	return o.Payload
}

func (o *AdminKeycloakListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminKeycloakList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKeycloakListBadRequest creates a AdminKeycloakListBadRequest with default headers values
func NewAdminKeycloakListBadRequest() *AdminKeycloakListBadRequest {
	return &AdminKeycloakListBadRequest{}
}

/*
AdminKeycloakListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AdminKeycloakListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin keycloak list bad request response has a 2xx status code
func (o *AdminKeycloakListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin keycloak list bad request response has a 3xx status code
func (o *AdminKeycloakListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list bad request response has a 4xx status code
func (o *AdminKeycloakListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin keycloak list bad request response has a 5xx status code
func (o *AdminKeycloakListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin keycloak list bad request response a status code equal to that given
func (o *AdminKeycloakListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AdminKeycloakListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminKeycloakListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListBadRequest  %+v", 400, o.Payload)
}

func (o *AdminKeycloakListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminKeycloakListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKeycloakListUnauthorized creates a AdminKeycloakListUnauthorized with default headers values
func NewAdminKeycloakListUnauthorized() *AdminKeycloakListUnauthorized {
	return &AdminKeycloakListUnauthorized{}
}

/*
AdminKeycloakListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AdminKeycloakListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin keycloak list unauthorized response has a 2xx status code
func (o *AdminKeycloakListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin keycloak list unauthorized response has a 3xx status code
func (o *AdminKeycloakListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list unauthorized response has a 4xx status code
func (o *AdminKeycloakListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin keycloak list unauthorized response has a 5xx status code
func (o *AdminKeycloakListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin keycloak list unauthorized response a status code equal to that given
func (o *AdminKeycloakListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AdminKeycloakListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminKeycloakListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminKeycloakListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminKeycloakListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKeycloakListForbidden creates a AdminKeycloakListForbidden with default headers values
func NewAdminKeycloakListForbidden() *AdminKeycloakListForbidden {
	return &AdminKeycloakListForbidden{}
}

/*
AdminKeycloakListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AdminKeycloakListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin keycloak list forbidden response has a 2xx status code
func (o *AdminKeycloakListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin keycloak list forbidden response has a 3xx status code
func (o *AdminKeycloakListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list forbidden response has a 4xx status code
func (o *AdminKeycloakListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin keycloak list forbidden response has a 5xx status code
func (o *AdminKeycloakListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin keycloak list forbidden response a status code equal to that given
func (o *AdminKeycloakListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AdminKeycloakListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListForbidden  %+v", 403, o.Payload)
}

func (o *AdminKeycloakListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListForbidden  %+v", 403, o.Payload)
}

func (o *AdminKeycloakListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminKeycloakListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKeycloakListNotFound creates a AdminKeycloakListNotFound with default headers values
func NewAdminKeycloakListNotFound() *AdminKeycloakListNotFound {
	return &AdminKeycloakListNotFound{}
}

/*
AdminKeycloakListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AdminKeycloakListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this admin keycloak list not found response has a 2xx status code
func (o *AdminKeycloakListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin keycloak list not found response has a 3xx status code
func (o *AdminKeycloakListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list not found response has a 4xx status code
func (o *AdminKeycloakListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin keycloak list not found response has a 5xx status code
func (o *AdminKeycloakListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this admin keycloak list not found response a status code equal to that given
func (o *AdminKeycloakListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AdminKeycloakListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListNotFound  %+v", 404, o.Payload)
}

func (o *AdminKeycloakListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListNotFound  %+v", 404, o.Payload)
}

func (o *AdminKeycloakListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AdminKeycloakListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminKeycloakListInternalServerError creates a AdminKeycloakListInternalServerError with default headers values
func NewAdminKeycloakListInternalServerError() *AdminKeycloakListInternalServerError {
	return &AdminKeycloakListInternalServerError{}
}

/*
AdminKeycloakListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AdminKeycloakListInternalServerError struct {
}

// IsSuccess returns true when this admin keycloak list internal server error response has a 2xx status code
func (o *AdminKeycloakListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin keycloak list internal server error response has a 3xx status code
func (o *AdminKeycloakListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin keycloak list internal server error response has a 4xx status code
func (o *AdminKeycloakListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin keycloak list internal server error response has a 5xx status code
func (o *AdminKeycloakListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin keycloak list internal server error response a status code equal to that given
func (o *AdminKeycloakListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AdminKeycloakListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListInternalServerError ", 500)
}

func (o *AdminKeycloakListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Admin/keycloak/list][%d] adminKeycloakListInternalServerError ", 500)
}

func (o *AdminKeycloakListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
