// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectsSpotVmsOperationsReader is a Reader for the ProjectsSpotVmsOperations structure.
type ProjectsSpotVmsOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsSpotVmsOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsSpotVmsOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsSpotVmsOperationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsSpotVmsOperationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsSpotVmsOperationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsSpotVmsOperationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsSpotVmsOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsSpotVmsOperationsOK creates a ProjectsSpotVmsOperationsOK with default headers values
func NewProjectsSpotVmsOperationsOK() *ProjectsSpotVmsOperationsOK {
	return &ProjectsSpotVmsOperationsOK{}
}

/*
ProjectsSpotVmsOperationsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsSpotVmsOperationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects spot vms operations o k response has a 2xx status code
func (o *ProjectsSpotVmsOperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects spot vms operations o k response has a 3xx status code
func (o *ProjectsSpotVmsOperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations o k response has a 4xx status code
func (o *ProjectsSpotVmsOperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects spot vms operations o k response has a 5xx status code
func (o *ProjectsSpotVmsOperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot vms operations o k response a status code equal to that given
func (o *ProjectsSpotVmsOperationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsSpotVmsOperationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsSpotVmsOperationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsSpotVmsOperationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsSpotVmsOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotVmsOperationsBadRequest creates a ProjectsSpotVmsOperationsBadRequest with default headers values
func NewProjectsSpotVmsOperationsBadRequest() *ProjectsSpotVmsOperationsBadRequest {
	return &ProjectsSpotVmsOperationsBadRequest{}
}

/*
ProjectsSpotVmsOperationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsSpotVmsOperationsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects spot vms operations bad request response has a 2xx status code
func (o *ProjectsSpotVmsOperationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot vms operations bad request response has a 3xx status code
func (o *ProjectsSpotVmsOperationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations bad request response has a 4xx status code
func (o *ProjectsSpotVmsOperationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot vms operations bad request response has a 5xx status code
func (o *ProjectsSpotVmsOperationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot vms operations bad request response a status code equal to that given
func (o *ProjectsSpotVmsOperationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsSpotVmsOperationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsSpotVmsOperationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsSpotVmsOperationsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsSpotVmsOperationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotVmsOperationsUnauthorized creates a ProjectsSpotVmsOperationsUnauthorized with default headers values
func NewProjectsSpotVmsOperationsUnauthorized() *ProjectsSpotVmsOperationsUnauthorized {
	return &ProjectsSpotVmsOperationsUnauthorized{}
}

/*
ProjectsSpotVmsOperationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsSpotVmsOperationsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this projects spot vms operations unauthorized response has a 2xx status code
func (o *ProjectsSpotVmsOperationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot vms operations unauthorized response has a 3xx status code
func (o *ProjectsSpotVmsOperationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations unauthorized response has a 4xx status code
func (o *ProjectsSpotVmsOperationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot vms operations unauthorized response has a 5xx status code
func (o *ProjectsSpotVmsOperationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot vms operations unauthorized response a status code equal to that given
func (o *ProjectsSpotVmsOperationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsSpotVmsOperationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsSpotVmsOperationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsSpotVmsOperationsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsSpotVmsOperationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotVmsOperationsForbidden creates a ProjectsSpotVmsOperationsForbidden with default headers values
func NewProjectsSpotVmsOperationsForbidden() *ProjectsSpotVmsOperationsForbidden {
	return &ProjectsSpotVmsOperationsForbidden{}
}

/*
ProjectsSpotVmsOperationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsSpotVmsOperationsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this projects spot vms operations forbidden response has a 2xx status code
func (o *ProjectsSpotVmsOperationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot vms operations forbidden response has a 3xx status code
func (o *ProjectsSpotVmsOperationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations forbidden response has a 4xx status code
func (o *ProjectsSpotVmsOperationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot vms operations forbidden response has a 5xx status code
func (o *ProjectsSpotVmsOperationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot vms operations forbidden response a status code equal to that given
func (o *ProjectsSpotVmsOperationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsSpotVmsOperationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsSpotVmsOperationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsSpotVmsOperationsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsSpotVmsOperationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotVmsOperationsNotFound creates a ProjectsSpotVmsOperationsNotFound with default headers values
func NewProjectsSpotVmsOperationsNotFound() *ProjectsSpotVmsOperationsNotFound {
	return &ProjectsSpotVmsOperationsNotFound{}
}

/*
ProjectsSpotVmsOperationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsSpotVmsOperationsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this projects spot vms operations not found response has a 2xx status code
func (o *ProjectsSpotVmsOperationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot vms operations not found response has a 3xx status code
func (o *ProjectsSpotVmsOperationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations not found response has a 4xx status code
func (o *ProjectsSpotVmsOperationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot vms operations not found response has a 5xx status code
func (o *ProjectsSpotVmsOperationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot vms operations not found response a status code equal to that given
func (o *ProjectsSpotVmsOperationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsSpotVmsOperationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsSpotVmsOperationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsSpotVmsOperationsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsSpotVmsOperationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotVmsOperationsInternalServerError creates a ProjectsSpotVmsOperationsInternalServerError with default headers values
func NewProjectsSpotVmsOperationsInternalServerError() *ProjectsSpotVmsOperationsInternalServerError {
	return &ProjectsSpotVmsOperationsInternalServerError{}
}

/*
ProjectsSpotVmsOperationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsSpotVmsOperationsInternalServerError struct {
}

// IsSuccess returns true when this projects spot vms operations internal server error response has a 2xx status code
func (o *ProjectsSpotVmsOperationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot vms operations internal server error response has a 3xx status code
func (o *ProjectsSpotVmsOperationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot vms operations internal server error response has a 4xx status code
func (o *ProjectsSpotVmsOperationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects spot vms operations internal server error response has a 5xx status code
func (o *ProjectsSpotVmsOperationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects spot vms operations internal server error response a status code equal to that given
func (o *ProjectsSpotVmsOperationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsSpotVmsOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsInternalServerError ", 500)
}

func (o *ProjectsSpotVmsOperationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-vms][%d] projectsSpotVmsOperationsInternalServerError ", 500)
}

func (o *ProjectsSpotVmsOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
