// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// OpenstackProjectsReader is a Reader for the OpenstackProjects structure.
type OpenstackProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenstackProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenstackProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOpenstackProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOpenstackProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenstackProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenstackProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenstackProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOpenstackProjectsOK creates a OpenstackProjectsOK with default headers values
func NewOpenstackProjectsOK() *OpenstackProjectsOK {
	return &OpenstackProjectsOK{}
}

/*
OpenstackProjectsOK describes a response with status code 200, with default header values.

Success
*/
type OpenstackProjectsOK struct {
	Payload []*models.CommonStringBasedDropdownDto
}

// IsSuccess returns true when this openstack projects o k response has a 2xx status code
func (o *OpenstackProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this openstack projects o k response has a 3xx status code
func (o *OpenstackProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects o k response has a 4xx status code
func (o *OpenstackProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack projects o k response has a 5xx status code
func (o *OpenstackProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack projects o k response a status code equal to that given
func (o *OpenstackProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenstackProjectsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsOK  %+v", 200, o.Payload)
}

func (o *OpenstackProjectsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsOK  %+v", 200, o.Payload)
}

func (o *OpenstackProjectsOK) GetPayload() []*models.CommonStringBasedDropdownDto {
	return o.Payload
}

func (o *OpenstackProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackProjectsBadRequest creates a OpenstackProjectsBadRequest with default headers values
func NewOpenstackProjectsBadRequest() *OpenstackProjectsBadRequest {
	return &OpenstackProjectsBadRequest{}
}

/*
OpenstackProjectsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type OpenstackProjectsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack projects bad request response has a 2xx status code
func (o *OpenstackProjectsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack projects bad request response has a 3xx status code
func (o *OpenstackProjectsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects bad request response has a 4xx status code
func (o *OpenstackProjectsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack projects bad request response has a 5xx status code
func (o *OpenstackProjectsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack projects bad request response a status code equal to that given
func (o *OpenstackProjectsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *OpenstackProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackProjectsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *OpenstackProjectsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackProjectsUnauthorized creates a OpenstackProjectsUnauthorized with default headers values
func NewOpenstackProjectsUnauthorized() *OpenstackProjectsUnauthorized {
	return &OpenstackProjectsUnauthorized{}
}

/*
OpenstackProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type OpenstackProjectsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack projects unauthorized response has a 2xx status code
func (o *OpenstackProjectsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack projects unauthorized response has a 3xx status code
func (o *OpenstackProjectsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects unauthorized response has a 4xx status code
func (o *OpenstackProjectsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack projects unauthorized response has a 5xx status code
func (o *OpenstackProjectsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack projects unauthorized response a status code equal to that given
func (o *OpenstackProjectsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *OpenstackProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackProjectsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *OpenstackProjectsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackProjectsForbidden creates a OpenstackProjectsForbidden with default headers values
func NewOpenstackProjectsForbidden() *OpenstackProjectsForbidden {
	return &OpenstackProjectsForbidden{}
}

/*
OpenstackProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type OpenstackProjectsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack projects forbidden response has a 2xx status code
func (o *OpenstackProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack projects forbidden response has a 3xx status code
func (o *OpenstackProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects forbidden response has a 4xx status code
func (o *OpenstackProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack projects forbidden response has a 5xx status code
func (o *OpenstackProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack projects forbidden response a status code equal to that given
func (o *OpenstackProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *OpenstackProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackProjectsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsForbidden  %+v", 403, o.Payload)
}

func (o *OpenstackProjectsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackProjectsNotFound creates a OpenstackProjectsNotFound with default headers values
func NewOpenstackProjectsNotFound() *OpenstackProjectsNotFound {
	return &OpenstackProjectsNotFound{}
}

/*
OpenstackProjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type OpenstackProjectsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this openstack projects not found response has a 2xx status code
func (o *OpenstackProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack projects not found response has a 3xx status code
func (o *OpenstackProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects not found response has a 4xx status code
func (o *OpenstackProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this openstack projects not found response has a 5xx status code
func (o *OpenstackProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this openstack projects not found response a status code equal to that given
func (o *OpenstackProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *OpenstackProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackProjectsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsNotFound  %+v", 404, o.Payload)
}

func (o *OpenstackProjectsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *OpenstackProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenstackProjectsInternalServerError creates a OpenstackProjectsInternalServerError with default headers values
func NewOpenstackProjectsInternalServerError() *OpenstackProjectsInternalServerError {
	return &OpenstackProjectsInternalServerError{}
}

/*
OpenstackProjectsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type OpenstackProjectsInternalServerError struct {
}

// IsSuccess returns true when this openstack projects internal server error response has a 2xx status code
func (o *OpenstackProjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this openstack projects internal server error response has a 3xx status code
func (o *OpenstackProjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this openstack projects internal server error response has a 4xx status code
func (o *OpenstackProjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this openstack projects internal server error response has a 5xx status code
func (o *OpenstackProjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this openstack projects internal server error response a status code equal to that given
func (o *OpenstackProjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *OpenstackProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsInternalServerError ", 500)
}

func (o *OpenstackProjectsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Openstack/projects][%d] openstackProjectsInternalServerError ", 500)
}

func (o *OpenstackProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
