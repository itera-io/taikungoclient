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

// ProjectsDeleteWholeProjectReader is a Reader for the ProjectsDeleteWholeProject structure.
type ProjectsDeleteWholeProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDeleteWholeProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsDeleteWholeProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewProjectsDeleteWholeProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsDeleteWholeProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsDeleteWholeProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsDeleteWholeProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsDeleteWholeProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsDeleteWholeProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsDeleteWholeProjectOK creates a ProjectsDeleteWholeProjectOK with default headers values
func NewProjectsDeleteWholeProjectOK() *ProjectsDeleteWholeProjectOK {
	return &ProjectsDeleteWholeProjectOK{}
}

/*
ProjectsDeleteWholeProjectOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsDeleteWholeProjectOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects delete whole project o k response has a 2xx status code
func (o *ProjectsDeleteWholeProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects delete whole project o k response has a 3xx status code
func (o *ProjectsDeleteWholeProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project o k response has a 4xx status code
func (o *ProjectsDeleteWholeProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete whole project o k response has a 5xx status code
func (o *ProjectsDeleteWholeProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project o k response a status code equal to that given
func (o *ProjectsDeleteWholeProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsDeleteWholeProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectOK  %+v", 200, o.Payload)
}

func (o *ProjectsDeleteWholeProjectOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectOK  %+v", 200, o.Payload)
}

func (o *ProjectsDeleteWholeProjectOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsDeleteWholeProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteWholeProjectNoContent creates a ProjectsDeleteWholeProjectNoContent with default headers values
func NewProjectsDeleteWholeProjectNoContent() *ProjectsDeleteWholeProjectNoContent {
	return &ProjectsDeleteWholeProjectNoContent{}
}

/*
ProjectsDeleteWholeProjectNoContent describes a response with status code 204, with default header values.

Success
*/
type ProjectsDeleteWholeProjectNoContent struct {
}

// IsSuccess returns true when this projects delete whole project no content response has a 2xx status code
func (o *ProjectsDeleteWholeProjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects delete whole project no content response has a 3xx status code
func (o *ProjectsDeleteWholeProjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project no content response has a 4xx status code
func (o *ProjectsDeleteWholeProjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete whole project no content response has a 5xx status code
func (o *ProjectsDeleteWholeProjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project no content response a status code equal to that given
func (o *ProjectsDeleteWholeProjectNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ProjectsDeleteWholeProjectNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectNoContent ", 204)
}

func (o *ProjectsDeleteWholeProjectNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectNoContent ", 204)
}

func (o *ProjectsDeleteWholeProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsDeleteWholeProjectBadRequest creates a ProjectsDeleteWholeProjectBadRequest with default headers values
func NewProjectsDeleteWholeProjectBadRequest() *ProjectsDeleteWholeProjectBadRequest {
	return &ProjectsDeleteWholeProjectBadRequest{}
}

/*
ProjectsDeleteWholeProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsDeleteWholeProjectBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this projects delete whole project bad request response has a 2xx status code
func (o *ProjectsDeleteWholeProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete whole project bad request response has a 3xx status code
func (o *ProjectsDeleteWholeProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project bad request response has a 4xx status code
func (o *ProjectsDeleteWholeProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete whole project bad request response has a 5xx status code
func (o *ProjectsDeleteWholeProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project bad request response a status code equal to that given
func (o *ProjectsDeleteWholeProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsDeleteWholeProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDeleteWholeProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDeleteWholeProjectBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDeleteWholeProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteWholeProjectUnauthorized creates a ProjectsDeleteWholeProjectUnauthorized with default headers values
func NewProjectsDeleteWholeProjectUnauthorized() *ProjectsDeleteWholeProjectUnauthorized {
	return &ProjectsDeleteWholeProjectUnauthorized{}
}

/*
ProjectsDeleteWholeProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsDeleteWholeProjectUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects delete whole project unauthorized response has a 2xx status code
func (o *ProjectsDeleteWholeProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete whole project unauthorized response has a 3xx status code
func (o *ProjectsDeleteWholeProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project unauthorized response has a 4xx status code
func (o *ProjectsDeleteWholeProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete whole project unauthorized response has a 5xx status code
func (o *ProjectsDeleteWholeProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project unauthorized response a status code equal to that given
func (o *ProjectsDeleteWholeProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsDeleteWholeProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDeleteWholeProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDeleteWholeProjectUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteWholeProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteWholeProjectForbidden creates a ProjectsDeleteWholeProjectForbidden with default headers values
func NewProjectsDeleteWholeProjectForbidden() *ProjectsDeleteWholeProjectForbidden {
	return &ProjectsDeleteWholeProjectForbidden{}
}

/*
ProjectsDeleteWholeProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsDeleteWholeProjectForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects delete whole project forbidden response has a 2xx status code
func (o *ProjectsDeleteWholeProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete whole project forbidden response has a 3xx status code
func (o *ProjectsDeleteWholeProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project forbidden response has a 4xx status code
func (o *ProjectsDeleteWholeProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete whole project forbidden response has a 5xx status code
func (o *ProjectsDeleteWholeProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project forbidden response a status code equal to that given
func (o *ProjectsDeleteWholeProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsDeleteWholeProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDeleteWholeProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDeleteWholeProjectForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteWholeProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteWholeProjectNotFound creates a ProjectsDeleteWholeProjectNotFound with default headers values
func NewProjectsDeleteWholeProjectNotFound() *ProjectsDeleteWholeProjectNotFound {
	return &ProjectsDeleteWholeProjectNotFound{}
}

/*
ProjectsDeleteWholeProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsDeleteWholeProjectNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects delete whole project not found response has a 2xx status code
func (o *ProjectsDeleteWholeProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete whole project not found response has a 3xx status code
func (o *ProjectsDeleteWholeProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project not found response has a 4xx status code
func (o *ProjectsDeleteWholeProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects delete whole project not found response has a 5xx status code
func (o *ProjectsDeleteWholeProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects delete whole project not found response a status code equal to that given
func (o *ProjectsDeleteWholeProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsDeleteWholeProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeleteWholeProjectNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDeleteWholeProjectNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDeleteWholeProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDeleteWholeProjectInternalServerError creates a ProjectsDeleteWholeProjectInternalServerError with default headers values
func NewProjectsDeleteWholeProjectInternalServerError() *ProjectsDeleteWholeProjectInternalServerError {
	return &ProjectsDeleteWholeProjectInternalServerError{}
}

/*
ProjectsDeleteWholeProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsDeleteWholeProjectInternalServerError struct {
}

// IsSuccess returns true when this projects delete whole project internal server error response has a 2xx status code
func (o *ProjectsDeleteWholeProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects delete whole project internal server error response has a 3xx status code
func (o *ProjectsDeleteWholeProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects delete whole project internal server error response has a 4xx status code
func (o *ProjectsDeleteWholeProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects delete whole project internal server error response has a 5xx status code
func (o *ProjectsDeleteWholeProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects delete whole project internal server error response a status code equal to that given
func (o *ProjectsDeleteWholeProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsDeleteWholeProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectInternalServerError ", 500)
}

func (o *ProjectsDeleteWholeProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/deletewholeproject][%d] projectsDeleteWholeProjectInternalServerError ", 500)
}

func (o *ProjectsDeleteWholeProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
