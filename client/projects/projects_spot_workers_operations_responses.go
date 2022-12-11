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

// ProjectsSpotWorkersOperationsReader is a Reader for the ProjectsSpotWorkersOperations structure.
type ProjectsSpotWorkersOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsSpotWorkersOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsSpotWorkersOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsSpotWorkersOperationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsSpotWorkersOperationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsSpotWorkersOperationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsSpotWorkersOperationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsSpotWorkersOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsSpotWorkersOperationsOK creates a ProjectsSpotWorkersOperationsOK with default headers values
func NewProjectsSpotWorkersOperationsOK() *ProjectsSpotWorkersOperationsOK {
	return &ProjectsSpotWorkersOperationsOK{}
}

/*
ProjectsSpotWorkersOperationsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsSpotWorkersOperationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects spot workers operations o k response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects spot workers operations o k response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations o k response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects spot workers operations o k response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot workers operations o k response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsSpotWorkersOperationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsSpotWorkersOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotWorkersOperationsBadRequest creates a ProjectsSpotWorkersOperationsBadRequest with default headers values
func NewProjectsSpotWorkersOperationsBadRequest() *ProjectsSpotWorkersOperationsBadRequest {
	return &ProjectsSpotWorkersOperationsBadRequest{}
}

/*
ProjectsSpotWorkersOperationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsSpotWorkersOperationsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this projects spot workers operations bad request response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot workers operations bad request response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations bad request response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot workers operations bad request response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot workers operations bad request response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsSpotWorkersOperationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsSpotWorkersOperationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotWorkersOperationsUnauthorized creates a ProjectsSpotWorkersOperationsUnauthorized with default headers values
func NewProjectsSpotWorkersOperationsUnauthorized() *ProjectsSpotWorkersOperationsUnauthorized {
	return &ProjectsSpotWorkersOperationsUnauthorized{}
}

/*
ProjectsSpotWorkersOperationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsSpotWorkersOperationsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects spot workers operations unauthorized response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot workers operations unauthorized response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations unauthorized response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot workers operations unauthorized response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot workers operations unauthorized response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsSpotWorkersOperationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsSpotWorkersOperationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotWorkersOperationsForbidden creates a ProjectsSpotWorkersOperationsForbidden with default headers values
func NewProjectsSpotWorkersOperationsForbidden() *ProjectsSpotWorkersOperationsForbidden {
	return &ProjectsSpotWorkersOperationsForbidden{}
}

/*
ProjectsSpotWorkersOperationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsSpotWorkersOperationsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects spot workers operations forbidden response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot workers operations forbidden response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations forbidden response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot workers operations forbidden response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot workers operations forbidden response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsSpotWorkersOperationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsSpotWorkersOperationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotWorkersOperationsNotFound creates a ProjectsSpotWorkersOperationsNotFound with default headers values
func NewProjectsSpotWorkersOperationsNotFound() *ProjectsSpotWorkersOperationsNotFound {
	return &ProjectsSpotWorkersOperationsNotFound{}
}

/*
ProjectsSpotWorkersOperationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsSpotWorkersOperationsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects spot workers operations not found response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot workers operations not found response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations not found response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects spot workers operations not found response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects spot workers operations not found response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsSpotWorkersOperationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsSpotWorkersOperationsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsSpotWorkersOperationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsSpotWorkersOperationsInternalServerError creates a ProjectsSpotWorkersOperationsInternalServerError with default headers values
func NewProjectsSpotWorkersOperationsInternalServerError() *ProjectsSpotWorkersOperationsInternalServerError {
	return &ProjectsSpotWorkersOperationsInternalServerError{}
}

/*
ProjectsSpotWorkersOperationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsSpotWorkersOperationsInternalServerError struct {
}

// IsSuccess returns true when this projects spot workers operations internal server error response has a 2xx status code
func (o *ProjectsSpotWorkersOperationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects spot workers operations internal server error response has a 3xx status code
func (o *ProjectsSpotWorkersOperationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects spot workers operations internal server error response has a 4xx status code
func (o *ProjectsSpotWorkersOperationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects spot workers operations internal server error response has a 5xx status code
func (o *ProjectsSpotWorkersOperationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects spot workers operations internal server error response a status code equal to that given
func (o *ProjectsSpotWorkersOperationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsSpotWorkersOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsInternalServerError ", 500)
}

func (o *ProjectsSpotWorkersOperationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/toggle-spot-workers][%d] projectsSpotWorkersOperationsInternalServerError ", 500)
}

func (o *ProjectsSpotWorkersOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
