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

// ProjectsListReader is a Reader for the ProjectsList structure.
type ProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListOK creates a ProjectsListOK with default headers values
func NewProjectsListOK() *ProjectsListOK {
	return &ProjectsListOK{}
}

/*
ProjectsListOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListOK struct {
	Payload *models.ProjectsList
}

// IsSuccess returns true when this projects list o k response has a 2xx status code
func (o *ProjectsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects list o k response has a 3xx status code
func (o *ProjectsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list o k response has a 4xx status code
func (o *ProjectsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list o k response has a 5xx status code
func (o *ProjectsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list o k response a status code equal to that given
func (o *ProjectsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListOK  %+v", 200, o.Payload)
}

func (o *ProjectsListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListOK  %+v", 200, o.Payload)
}

func (o *ProjectsListOK) GetPayload() *models.ProjectsList {
	return o.Payload
}

func (o *ProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListBadRequest creates a ProjectsListBadRequest with default headers values
func NewProjectsListBadRequest() *ProjectsListBadRequest {
	return &ProjectsListBadRequest{}
}

/*
ProjectsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this projects list bad request response has a 2xx status code
func (o *ProjectsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list bad request response has a 3xx status code
func (o *ProjectsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list bad request response has a 4xx status code
func (o *ProjectsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list bad request response has a 5xx status code
func (o *ProjectsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list bad request response a status code equal to that given
func (o *ProjectsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListUnauthorized creates a ProjectsListUnauthorized with default headers values
func NewProjectsListUnauthorized() *ProjectsListUnauthorized {
	return &ProjectsListUnauthorized{}
}

/*
ProjectsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list unauthorized response has a 2xx status code
func (o *ProjectsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list unauthorized response has a 3xx status code
func (o *ProjectsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list unauthorized response has a 4xx status code
func (o *ProjectsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list unauthorized response has a 5xx status code
func (o *ProjectsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list unauthorized response a status code equal to that given
func (o *ProjectsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForbidden creates a ProjectsListForbidden with default headers values
func NewProjectsListForbidden() *ProjectsListForbidden {
	return &ProjectsListForbidden{}
}

/*
ProjectsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list forbidden response has a 2xx status code
func (o *ProjectsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list forbidden response has a 3xx status code
func (o *ProjectsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list forbidden response has a 4xx status code
func (o *ProjectsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list forbidden response has a 5xx status code
func (o *ProjectsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list forbidden response a status code equal to that given
func (o *ProjectsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListNotFound creates a ProjectsListNotFound with default headers values
func NewProjectsListNotFound() *ProjectsListNotFound {
	return &ProjectsListNotFound{}
}

/*
ProjectsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list not found response has a 2xx status code
func (o *ProjectsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list not found response has a 3xx status code
func (o *ProjectsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list not found response has a 4xx status code
func (o *ProjectsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list not found response has a 5xx status code
func (o *ProjectsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list not found response a status code equal to that given
func (o *ProjectsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListInternalServerError creates a ProjectsListInternalServerError with default headers values
func NewProjectsListInternalServerError() *ProjectsListInternalServerError {
	return &ProjectsListInternalServerError{}
}

/*
ProjectsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListInternalServerError struct {
}

// IsSuccess returns true when this projects list internal server error response has a 2xx status code
func (o *ProjectsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list internal server error response has a 3xx status code
func (o *ProjectsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list internal server error response has a 4xx status code
func (o *ProjectsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list internal server error response has a 5xx status code
func (o *ProjectsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects list internal server error response a status code equal to that given
func (o *ProjectsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListInternalServerError ", 500)
}

func (o *ProjectsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects][%d] projectsListInternalServerError ", 500)
}

func (o *ProjectsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
