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

// ProjectsListForPollerReader is a Reader for the ProjectsListForPoller structure.
type ProjectsListForPollerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListForPollerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListForPollerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListForPollerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListForPollerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListForPollerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListForPollerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewProjectsListForPollerTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListForPollerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListForPollerOK creates a ProjectsListForPollerOK with default headers values
func NewProjectsListForPollerOK() *ProjectsListForPollerOK {
	return &ProjectsListForPollerOK{}
}

/* ProjectsListForPollerOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListForPollerOK struct {
	Payload *models.ProjectListForPoller
}

func (o *ProjectsListForPollerOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerOK  %+v", 200, o.Payload)
}
func (o *ProjectsListForPollerOK) GetPayload() *models.ProjectListForPoller {
	return o.Payload
}

func (o *ProjectsListForPollerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectListForPoller)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerBadRequest creates a ProjectsListForPollerBadRequest with default headers values
func NewProjectsListForPollerBadRequest() *ProjectsListForPollerBadRequest {
	return &ProjectsListForPollerBadRequest{}
}

/* ProjectsListForPollerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListForPollerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectsListForPollerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectsListForPollerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsListForPollerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerUnauthorized creates a ProjectsListForPollerUnauthorized with default headers values
func NewProjectsListForPollerUnauthorized() *ProjectsListForPollerUnauthorized {
	return &ProjectsListForPollerUnauthorized{}
}

/* ProjectsListForPollerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListForPollerUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsListForPollerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectsListForPollerUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListForPollerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerForbidden creates a ProjectsListForPollerForbidden with default headers values
func NewProjectsListForPollerForbidden() *ProjectsListForPollerForbidden {
	return &ProjectsListForPollerForbidden{}
}

/* ProjectsListForPollerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListForPollerForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsListForPollerForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerForbidden  %+v", 403, o.Payload)
}
func (o *ProjectsListForPollerForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListForPollerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerNotFound creates a ProjectsListForPollerNotFound with default headers values
func NewProjectsListForPollerNotFound() *ProjectsListForPollerNotFound {
	return &ProjectsListForPollerNotFound{}
}

/* ProjectsListForPollerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListForPollerNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsListForPollerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerNotFound  %+v", 404, o.Payload)
}
func (o *ProjectsListForPollerNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListForPollerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerTooManyRequests creates a ProjectsListForPollerTooManyRequests with default headers values
func NewProjectsListForPollerTooManyRequests() *ProjectsListForPollerTooManyRequests {
	return &ProjectsListForPollerTooManyRequests{}
}

/* ProjectsListForPollerTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type ProjectsListForPollerTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *ProjectsListForPollerTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerTooManyRequests  %+v", 429, o.Payload)
}
func (o *ProjectsListForPollerTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListForPollerTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForPollerInternalServerError creates a ProjectsListForPollerInternalServerError with default headers values
func NewProjectsListForPollerInternalServerError() *ProjectsListForPollerInternalServerError {
	return &ProjectsListForPollerInternalServerError{}
}

/* ProjectsListForPollerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListForPollerInternalServerError struct {
}

func (o *ProjectsListForPollerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forpoller][%d] projectsListForPollerInternalServerError ", 500)
}

func (o *ProjectsListForPollerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
