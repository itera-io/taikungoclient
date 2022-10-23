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

// ProjectsMonitoringOperationsReader is a Reader for the ProjectsMonitoringOperations structure.
type ProjectsMonitoringOperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsMonitoringOperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsMonitoringOperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsMonitoringOperationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsMonitoringOperationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsMonitoringOperationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsMonitoringOperationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsMonitoringOperationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsMonitoringOperationsOK creates a ProjectsMonitoringOperationsOK with default headers values
func NewProjectsMonitoringOperationsOK() *ProjectsMonitoringOperationsOK {
	return &ProjectsMonitoringOperationsOK{}
}

/*
ProjectsMonitoringOperationsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsMonitoringOperationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects monitoring operations o k response has a 2xx status code
func (o *ProjectsMonitoringOperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects monitoring operations o k response has a 3xx status code
func (o *ProjectsMonitoringOperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations o k response has a 4xx status code
func (o *ProjectsMonitoringOperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects monitoring operations o k response has a 5xx status code
func (o *ProjectsMonitoringOperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects monitoring operations o k response a status code equal to that given
func (o *ProjectsMonitoringOperationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsMonitoringOperationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsMonitoringOperationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsOK  %+v", 200, o.Payload)
}

func (o *ProjectsMonitoringOperationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsMonitoringOperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsMonitoringOperationsBadRequest creates a ProjectsMonitoringOperationsBadRequest with default headers values
func NewProjectsMonitoringOperationsBadRequest() *ProjectsMonitoringOperationsBadRequest {
	return &ProjectsMonitoringOperationsBadRequest{}
}

/*
ProjectsMonitoringOperationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsMonitoringOperationsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this projects monitoring operations bad request response has a 2xx status code
func (o *ProjectsMonitoringOperationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects monitoring operations bad request response has a 3xx status code
func (o *ProjectsMonitoringOperationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations bad request response has a 4xx status code
func (o *ProjectsMonitoringOperationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects monitoring operations bad request response has a 5xx status code
func (o *ProjectsMonitoringOperationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects monitoring operations bad request response a status code equal to that given
func (o *ProjectsMonitoringOperationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsMonitoringOperationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsMonitoringOperationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsMonitoringOperationsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectsMonitoringOperationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsMonitoringOperationsUnauthorized creates a ProjectsMonitoringOperationsUnauthorized with default headers values
func NewProjectsMonitoringOperationsUnauthorized() *ProjectsMonitoringOperationsUnauthorized {
	return &ProjectsMonitoringOperationsUnauthorized{}
}

/*
ProjectsMonitoringOperationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsMonitoringOperationsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects monitoring operations unauthorized response has a 2xx status code
func (o *ProjectsMonitoringOperationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects monitoring operations unauthorized response has a 3xx status code
func (o *ProjectsMonitoringOperationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations unauthorized response has a 4xx status code
func (o *ProjectsMonitoringOperationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects monitoring operations unauthorized response has a 5xx status code
func (o *ProjectsMonitoringOperationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects monitoring operations unauthorized response a status code equal to that given
func (o *ProjectsMonitoringOperationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsMonitoringOperationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsMonitoringOperationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsMonitoringOperationsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsMonitoringOperationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsMonitoringOperationsForbidden creates a ProjectsMonitoringOperationsForbidden with default headers values
func NewProjectsMonitoringOperationsForbidden() *ProjectsMonitoringOperationsForbidden {
	return &ProjectsMonitoringOperationsForbidden{}
}

/*
ProjectsMonitoringOperationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsMonitoringOperationsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects monitoring operations forbidden response has a 2xx status code
func (o *ProjectsMonitoringOperationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects monitoring operations forbidden response has a 3xx status code
func (o *ProjectsMonitoringOperationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations forbidden response has a 4xx status code
func (o *ProjectsMonitoringOperationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects monitoring operations forbidden response has a 5xx status code
func (o *ProjectsMonitoringOperationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects monitoring operations forbidden response a status code equal to that given
func (o *ProjectsMonitoringOperationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsMonitoringOperationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsMonitoringOperationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsMonitoringOperationsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsMonitoringOperationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsMonitoringOperationsNotFound creates a ProjectsMonitoringOperationsNotFound with default headers values
func NewProjectsMonitoringOperationsNotFound() *ProjectsMonitoringOperationsNotFound {
	return &ProjectsMonitoringOperationsNotFound{}
}

/*
ProjectsMonitoringOperationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsMonitoringOperationsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects monitoring operations not found response has a 2xx status code
func (o *ProjectsMonitoringOperationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects monitoring operations not found response has a 3xx status code
func (o *ProjectsMonitoringOperationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations not found response has a 4xx status code
func (o *ProjectsMonitoringOperationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects monitoring operations not found response has a 5xx status code
func (o *ProjectsMonitoringOperationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects monitoring operations not found response a status code equal to that given
func (o *ProjectsMonitoringOperationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsMonitoringOperationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsMonitoringOperationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsMonitoringOperationsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsMonitoringOperationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsMonitoringOperationsInternalServerError creates a ProjectsMonitoringOperationsInternalServerError with default headers values
func NewProjectsMonitoringOperationsInternalServerError() *ProjectsMonitoringOperationsInternalServerError {
	return &ProjectsMonitoringOperationsInternalServerError{}
}

/*
ProjectsMonitoringOperationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsMonitoringOperationsInternalServerError struct {
}

// IsSuccess returns true when this projects monitoring operations internal server error response has a 2xx status code
func (o *ProjectsMonitoringOperationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects monitoring operations internal server error response has a 3xx status code
func (o *ProjectsMonitoringOperationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects monitoring operations internal server error response has a 4xx status code
func (o *ProjectsMonitoringOperationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects monitoring operations internal server error response has a 5xx status code
func (o *ProjectsMonitoringOperationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects monitoring operations internal server error response a status code equal to that given
func (o *ProjectsMonitoringOperationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsMonitoringOperationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsInternalServerError ", 500)
}

func (o *ProjectsMonitoringOperationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/monitoring][%d] projectsMonitoringOperationsInternalServerError ", 500)
}

func (o *ProjectsMonitoringOperationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
