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

// ProjectsDeleteReader is a Reader for the ProjectsDelete structure.
type ProjectsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsDeleteOK creates a ProjectsDeleteOK with default headers values
func NewProjectsDeleteOK() *ProjectsDeleteOK {
	return &ProjectsDeleteOK{}
}

/* ProjectsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsDeleteOK struct {
}

func (o *ProjectsDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteOK ", 200)
}

func (o *ProjectsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsDeleteNoContent creates a ProjectsDeleteNoContent with default headers values
func NewProjectsDeleteNoContent() *ProjectsDeleteNoContent {
	return &ProjectsDeleteNoContent{}
}

/* ProjectsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type ProjectsDeleteNoContent struct {
}

func (o *ProjectsDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteNoContent ", 204)
}

func (o *ProjectsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsDeleteBadRequest creates a ProjectsDeleteBadRequest with default headers values
func NewProjectsDeleteBadRequest() *ProjectsDeleteBadRequest {
	return &ProjectsDeleteBadRequest{}
}

/* ProjectsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectsDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteUnauthorized creates a ProjectsDeleteUnauthorized with default headers values
func NewProjectsDeleteUnauthorized() *ProjectsDeleteUnauthorized {
	return &ProjectsDeleteUnauthorized{}
}

/* ProjectsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectsDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteForbidden creates a ProjectsDeleteForbidden with default headers values
func NewProjectsDeleteForbidden() *ProjectsDeleteForbidden {
	return &ProjectsDeleteForbidden{}
}

/* ProjectsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteForbidden  %+v", 403, o.Payload)
}
func (o *ProjectsDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteNotFound creates a ProjectsDeleteNotFound with default headers values
func NewProjectsDeleteNotFound() *ProjectsDeleteNotFound {
	return &ProjectsDeleteNotFound{}
}

/* ProjectsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteNotFound  %+v", 404, o.Payload)
}
func (o *ProjectsDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteInternalServerError creates a ProjectsDeleteInternalServerError with default headers values
func NewProjectsDeleteInternalServerError() *ProjectsDeleteInternalServerError {
	return &ProjectsDeleteInternalServerError{}
}

/* ProjectsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsDeleteInternalServerError struct {
}

func (o *ProjectsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/delete][%d] projectsDeleteInternalServerError ", 500)
}

func (o *ProjectsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
