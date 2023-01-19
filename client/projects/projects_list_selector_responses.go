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

// ProjectsListSelectorReader is a Reader for the ProjectsListSelector structure.
type ProjectsListSelectorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListSelectorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListSelectorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListSelectorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListSelectorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListSelectorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListSelectorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListSelectorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListSelectorOK creates a ProjectsListSelectorOK with default headers values
func NewProjectsListSelectorOK() *ProjectsListSelectorOK {
	return &ProjectsListSelectorOK{}
}

/*
ProjectsListSelectorOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListSelectorOK struct {
	Payload []*models.ProjectEntity
}

// IsSuccess returns true when this projects list selector o k response has a 2xx status code
func (o *ProjectsListSelectorOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects list selector o k response has a 3xx status code
func (o *ProjectsListSelectorOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector o k response has a 4xx status code
func (o *ProjectsListSelectorOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list selector o k response has a 5xx status code
func (o *ProjectsListSelectorOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list selector o k response a status code equal to that given
func (o *ProjectsListSelectorOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the projects list selector o k response
func (o *ProjectsListSelectorOK) Code() int {
	return 200
}

func (o *ProjectsListSelectorOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorOK  %+v", 200, o.Payload)
}

func (o *ProjectsListSelectorOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorOK  %+v", 200, o.Payload)
}

func (o *ProjectsListSelectorOK) GetPayload() []*models.ProjectEntity {
	return o.Payload
}

func (o *ProjectsListSelectorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListSelectorBadRequest creates a ProjectsListSelectorBadRequest with default headers values
func NewProjectsListSelectorBadRequest() *ProjectsListSelectorBadRequest {
	return &ProjectsListSelectorBadRequest{}
}

/*
ProjectsListSelectorBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListSelectorBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list selector bad request response has a 2xx status code
func (o *ProjectsListSelectorBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list selector bad request response has a 3xx status code
func (o *ProjectsListSelectorBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector bad request response has a 4xx status code
func (o *ProjectsListSelectorBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list selector bad request response has a 5xx status code
func (o *ProjectsListSelectorBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list selector bad request response a status code equal to that given
func (o *ProjectsListSelectorBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the projects list selector bad request response
func (o *ProjectsListSelectorBadRequest) Code() int {
	return 400
}

func (o *ProjectsListSelectorBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListSelectorBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListSelectorBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListSelectorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListSelectorUnauthorized creates a ProjectsListSelectorUnauthorized with default headers values
func NewProjectsListSelectorUnauthorized() *ProjectsListSelectorUnauthorized {
	return &ProjectsListSelectorUnauthorized{}
}

/*
ProjectsListSelectorUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListSelectorUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list selector unauthorized response has a 2xx status code
func (o *ProjectsListSelectorUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list selector unauthorized response has a 3xx status code
func (o *ProjectsListSelectorUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector unauthorized response has a 4xx status code
func (o *ProjectsListSelectorUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list selector unauthorized response has a 5xx status code
func (o *ProjectsListSelectorUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list selector unauthorized response a status code equal to that given
func (o *ProjectsListSelectorUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the projects list selector unauthorized response
func (o *ProjectsListSelectorUnauthorized) Code() int {
	return 401
}

func (o *ProjectsListSelectorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListSelectorUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListSelectorUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListSelectorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListSelectorForbidden creates a ProjectsListSelectorForbidden with default headers values
func NewProjectsListSelectorForbidden() *ProjectsListSelectorForbidden {
	return &ProjectsListSelectorForbidden{}
}

/*
ProjectsListSelectorForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListSelectorForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list selector forbidden response has a 2xx status code
func (o *ProjectsListSelectorForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list selector forbidden response has a 3xx status code
func (o *ProjectsListSelectorForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector forbidden response has a 4xx status code
func (o *ProjectsListSelectorForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list selector forbidden response has a 5xx status code
func (o *ProjectsListSelectorForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list selector forbidden response a status code equal to that given
func (o *ProjectsListSelectorForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the projects list selector forbidden response
func (o *ProjectsListSelectorForbidden) Code() int {
	return 403
}

func (o *ProjectsListSelectorForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListSelectorForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListSelectorForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListSelectorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListSelectorNotFound creates a ProjectsListSelectorNotFound with default headers values
func NewProjectsListSelectorNotFound() *ProjectsListSelectorNotFound {
	return &ProjectsListSelectorNotFound{}
}

/*
ProjectsListSelectorNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListSelectorNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects list selector not found response has a 2xx status code
func (o *ProjectsListSelectorNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list selector not found response has a 3xx status code
func (o *ProjectsListSelectorNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector not found response has a 4xx status code
func (o *ProjectsListSelectorNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list selector not found response has a 5xx status code
func (o *ProjectsListSelectorNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list selector not found response a status code equal to that given
func (o *ProjectsListSelectorNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the projects list selector not found response
func (o *ProjectsListSelectorNotFound) Code() int {
	return 404
}

func (o *ProjectsListSelectorNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListSelectorNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListSelectorNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsListSelectorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListSelectorInternalServerError creates a ProjectsListSelectorInternalServerError with default headers values
func NewProjectsListSelectorInternalServerError() *ProjectsListSelectorInternalServerError {
	return &ProjectsListSelectorInternalServerError{}
}

/*
ProjectsListSelectorInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListSelectorInternalServerError struct {
}

// IsSuccess returns true when this projects list selector internal server error response has a 2xx status code
func (o *ProjectsListSelectorInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list selector internal server error response has a 3xx status code
func (o *ProjectsListSelectorInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list selector internal server error response has a 4xx status code
func (o *ProjectsListSelectorInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list selector internal server error response has a 5xx status code
func (o *ProjectsListSelectorInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects list selector internal server error response a status code equal to that given
func (o *ProjectsListSelectorInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the projects list selector internal server error response
func (o *ProjectsListSelectorInternalServerError) Code() int {
	return 500
}

func (o *ProjectsListSelectorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorInternalServerError ", 500)
}

func (o *ProjectsListSelectorInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/list][%d] projectsListSelectorInternalServerError ", 500)
}

func (o *ProjectsListSelectorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
