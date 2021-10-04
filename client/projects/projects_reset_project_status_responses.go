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

/* ProjectsResetProjectStatusOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsResetProjectStatusOK struct {
	Payload models.Unit
}

func (o *ProjectsResetProjectStatusOK) Error() string {
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

/* ProjectsResetProjectStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsResetProjectStatusBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectsResetProjectStatusBadRequest) Error() string {
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

/* ProjectsResetProjectStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsResetProjectStatusUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsResetProjectStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectsResetProjectStatusUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsResetProjectStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusForbidden creates a ProjectsResetProjectStatusForbidden with default headers values
func NewProjectsResetProjectStatusForbidden() *ProjectsResetProjectStatusForbidden {
	return &ProjectsResetProjectStatusForbidden{}
}

/* ProjectsResetProjectStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsResetProjectStatusForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsResetProjectStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusForbidden  %+v", 403, o.Payload)
}
func (o *ProjectsResetProjectStatusForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsResetProjectStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusNotFound creates a ProjectsResetProjectStatusNotFound with default headers values
func NewProjectsResetProjectStatusNotFound() *ProjectsResetProjectStatusNotFound {
	return &ProjectsResetProjectStatusNotFound{}
}

/* ProjectsResetProjectStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsResetProjectStatusNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsResetProjectStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusNotFound  %+v", 404, o.Payload)
}
func (o *ProjectsResetProjectStatusNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsResetProjectStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsResetProjectStatusInternalServerError creates a ProjectsResetProjectStatusInternalServerError with default headers values
func NewProjectsResetProjectStatusInternalServerError() *ProjectsResetProjectStatusInternalServerError {
	return &ProjectsResetProjectStatusInternalServerError{}
}

/* ProjectsResetProjectStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsResetProjectStatusInternalServerError struct {
}

func (o *ProjectsResetProjectStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/reset][%d] projectsResetProjectStatusInternalServerError ", 500)
}

func (o *ProjectsResetProjectStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
