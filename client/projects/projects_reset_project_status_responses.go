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

// ProjectsResetProjectStatusReader is a Reader for the ProjectsResetProjectStatus structure.
type ProjectsResetProjectStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsResetProjectStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsResetProjectStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsResetProjectStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsResetProjectStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsResetProjectStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsResetProjectStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsResetProjectStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsResetProjectStatusOK creates a ProjectsResetProjectStatusOK with default headers values
func NewProjectsResetProjectStatusOK() *ProjectsResetProjectStatusOK {
	return &ProjectsResetProjectStatusOK{}
}

/*
ProjectsResetProjectStatusOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsResetProjectStatusOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects reset project status o k response has a 2xx status code
func (o *ProjectsResetProjectStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects reset project status o k response has a 3xx status code
func (o *ProjectsResetProjectStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status o k response has a 4xx status code
func (o *ProjectsResetProjectStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects reset project status o k response has a 5xx status code
func (o *ProjectsResetProjectStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status o k response a status code equal to that given
func (o *ProjectsResetProjectStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsResetProjectStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsResetProjectStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusOK  %+v", 200, o.Payload)
}

func (o *ProjectsResetProjectStatusOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsResetProjectStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusBadRequest creates a ProjectsResetProjectStatusBadRequest with default headers values
func NewProjectsResetProjectStatusBadRequest() *ProjectsResetProjectStatusBadRequest {
	return &ProjectsResetProjectStatusBadRequest{}
}

/*
ProjectsResetProjectStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsResetProjectStatusBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects reset project status bad request response has a 2xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status bad request response has a 3xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status bad request response has a 4xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status bad request response has a 5xx status code
func (o *ProjectsResetProjectStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status bad request response a status code equal to that given
func (o *ProjectsResetProjectStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsResetProjectStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsResetProjectStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsResetProjectStatusBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsResetProjectStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusUnauthorized creates a ProjectsResetProjectStatusUnauthorized with default headers values
func NewProjectsResetProjectStatusUnauthorized() *ProjectsResetProjectStatusUnauthorized {
	return &ProjectsResetProjectStatusUnauthorized{}
}

/*
ProjectsResetProjectStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsResetProjectStatusUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects reset project status unauthorized response has a 2xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status unauthorized response has a 3xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status unauthorized response has a 4xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status unauthorized response has a 5xx status code
func (o *ProjectsResetProjectStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status unauthorized response a status code equal to that given
func (o *ProjectsResetProjectStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsResetProjectStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsResetProjectStatusUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsResetProjectStatusUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsResetProjectStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusForbidden creates a ProjectsResetProjectStatusForbidden with default headers values
func NewProjectsResetProjectStatusForbidden() *ProjectsResetProjectStatusForbidden {
	return &ProjectsResetProjectStatusForbidden{}
}

/*
ProjectsResetProjectStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsResetProjectStatusForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects reset project status forbidden response has a 2xx status code
func (o *ProjectsResetProjectStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status forbidden response has a 3xx status code
func (o *ProjectsResetProjectStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status forbidden response has a 4xx status code
func (o *ProjectsResetProjectStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status forbidden response has a 5xx status code
func (o *ProjectsResetProjectStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status forbidden response a status code equal to that given
func (o *ProjectsResetProjectStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsResetProjectStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsResetProjectStatusForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsResetProjectStatusForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsResetProjectStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusNotFound creates a ProjectsResetProjectStatusNotFound with default headers values
func NewProjectsResetProjectStatusNotFound() *ProjectsResetProjectStatusNotFound {
	return &ProjectsResetProjectStatusNotFound{}
}

/*
ProjectsResetProjectStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsResetProjectStatusNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects reset project status not found response has a 2xx status code
func (o *ProjectsResetProjectStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status not found response has a 3xx status code
func (o *ProjectsResetProjectStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status not found response has a 4xx status code
func (o *ProjectsResetProjectStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects reset project status not found response has a 5xx status code
func (o *ProjectsResetProjectStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects reset project status not found response a status code equal to that given
func (o *ProjectsResetProjectStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsResetProjectStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsResetProjectStatusNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsResetProjectStatusNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsResetProjectStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusInternalServerError creates a ProjectsResetProjectStatusInternalServerError with default headers values
func NewProjectsResetProjectStatusInternalServerError() *ProjectsResetProjectStatusInternalServerError {
	return &ProjectsResetProjectStatusInternalServerError{}
}

/*
ProjectsResetProjectStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsResetProjectStatusInternalServerError struct {
}

// IsSuccess returns true when this projects reset project status internal server error response has a 2xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects reset project status internal server error response has a 3xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects reset project status internal server error response has a 4xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects reset project status internal server error response has a 5xx status code
func (o *ProjectsResetProjectStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects reset project status internal server error response a status code equal to that given
func (o *ProjectsResetProjectStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsResetProjectStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusInternalServerError ", 500)
}

func (o *ProjectsResetProjectStatusInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusInternalServerError ", 500)
}

func (o *ProjectsResetProjectStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
