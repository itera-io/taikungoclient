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

// ProjectsEditHealthStatusReader is a Reader for the ProjectsEditHealthStatus structure.
type ProjectsEditHealthStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsEditHealthStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsEditHealthStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsEditHealthStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsEditHealthStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsEditHealthStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsEditHealthStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsEditHealthStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsEditHealthStatusOK creates a ProjectsEditHealthStatusOK with default headers values
func NewProjectsEditHealthStatusOK() *ProjectsEditHealthStatusOK {
	return &ProjectsEditHealthStatusOK{}
}

/*
ProjectsEditHealthStatusOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsEditHealthStatusOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects edit health status o k response has a 2xx status code
func (o *ProjectsEditHealthStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects edit health status o k response has a 3xx status code
func (o *ProjectsEditHealthStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status o k response has a 4xx status code
func (o *ProjectsEditHealthStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects edit health status o k response has a 5xx status code
func (o *ProjectsEditHealthStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit health status o k response a status code equal to that given
func (o *ProjectsEditHealthStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsEditHealthStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsEditHealthStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsEditHealthStatusOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsEditHealthStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditHealthStatusBadRequest creates a ProjectsEditHealthStatusBadRequest with default headers values
func NewProjectsEditHealthStatusBadRequest() *ProjectsEditHealthStatusBadRequest {
	return &ProjectsEditHealthStatusBadRequest{}
}

/*
ProjectsEditHealthStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsEditHealthStatusBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects edit health status bad request response has a 2xx status code
func (o *ProjectsEditHealthStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit health status bad request response has a 3xx status code
func (o *ProjectsEditHealthStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status bad request response has a 4xx status code
func (o *ProjectsEditHealthStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit health status bad request response has a 5xx status code
func (o *ProjectsEditHealthStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit health status bad request response a status code equal to that given
func (o *ProjectsEditHealthStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsEditHealthStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsEditHealthStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsEditHealthStatusBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsEditHealthStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditHealthStatusUnauthorized creates a ProjectsEditHealthStatusUnauthorized with default headers values
func NewProjectsEditHealthStatusUnauthorized() *ProjectsEditHealthStatusUnauthorized {
	return &ProjectsEditHealthStatusUnauthorized{}
}

/*
ProjectsEditHealthStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsEditHealthStatusUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects edit health status unauthorized response has a 2xx status code
func (o *ProjectsEditHealthStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit health status unauthorized response has a 3xx status code
func (o *ProjectsEditHealthStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status unauthorized response has a 4xx status code
func (o *ProjectsEditHealthStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit health status unauthorized response has a 5xx status code
func (o *ProjectsEditHealthStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit health status unauthorized response a status code equal to that given
func (o *ProjectsEditHealthStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsEditHealthStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsEditHealthStatusUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsEditHealthStatusUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsEditHealthStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditHealthStatusForbidden creates a ProjectsEditHealthStatusForbidden with default headers values
func NewProjectsEditHealthStatusForbidden() *ProjectsEditHealthStatusForbidden {
	return &ProjectsEditHealthStatusForbidden{}
}

/*
ProjectsEditHealthStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsEditHealthStatusForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects edit health status forbidden response has a 2xx status code
func (o *ProjectsEditHealthStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit health status forbidden response has a 3xx status code
func (o *ProjectsEditHealthStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status forbidden response has a 4xx status code
func (o *ProjectsEditHealthStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit health status forbidden response has a 5xx status code
func (o *ProjectsEditHealthStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit health status forbidden response a status code equal to that given
func (o *ProjectsEditHealthStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsEditHealthStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsEditHealthStatusForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsEditHealthStatusForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsEditHealthStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditHealthStatusNotFound creates a ProjectsEditHealthStatusNotFound with default headers values
func NewProjectsEditHealthStatusNotFound() *ProjectsEditHealthStatusNotFound {
	return &ProjectsEditHealthStatusNotFound{}
}

/*
ProjectsEditHealthStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsEditHealthStatusNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects edit health status not found response has a 2xx status code
func (o *ProjectsEditHealthStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit health status not found response has a 3xx status code
func (o *ProjectsEditHealthStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status not found response has a 4xx status code
func (o *ProjectsEditHealthStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit health status not found response has a 5xx status code
func (o *ProjectsEditHealthStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit health status not found response a status code equal to that given
func (o *ProjectsEditHealthStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsEditHealthStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsEditHealthStatusNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsEditHealthStatusNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsEditHealthStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditHealthStatusInternalServerError creates a ProjectsEditHealthStatusInternalServerError with default headers values
func NewProjectsEditHealthStatusInternalServerError() *ProjectsEditHealthStatusInternalServerError {
	return &ProjectsEditHealthStatusInternalServerError{}
}

/*
ProjectsEditHealthStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsEditHealthStatusInternalServerError struct {
}

// IsSuccess returns true when this projects edit health status internal server error response has a 2xx status code
func (o *ProjectsEditHealthStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit health status internal server error response has a 3xx status code
func (o *ProjectsEditHealthStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit health status internal server error response has a 4xx status code
func (o *ProjectsEditHealthStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects edit health status internal server error response has a 5xx status code
func (o *ProjectsEditHealthStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects edit health status internal server error response a status code equal to that given
func (o *ProjectsEditHealthStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsEditHealthStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusInternalServerError ", 500)
}

func (o *ProjectsEditHealthStatusInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/updatehealth/{projectId}][%d] projectsEditHealthStatusInternalServerError ", 500)
}

func (o *ProjectsEditHealthStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
