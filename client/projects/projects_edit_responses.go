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

// ProjectsEditReader is a Reader for the ProjectsEdit structure.
type ProjectsEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsEditOK creates a ProjectsEditOK with default headers values
func NewProjectsEditOK() *ProjectsEditOK {
	return &ProjectsEditOK{}
}

/*
ProjectsEditOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsEditOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects edit o k response has a 2xx status code
func (o *ProjectsEditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects edit o k response has a 3xx status code
func (o *ProjectsEditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit o k response has a 4xx status code
func (o *ProjectsEditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects edit o k response has a 5xx status code
func (o *ProjectsEditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit o k response a status code equal to that given
func (o *ProjectsEditOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectsEditOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditOK  %+v", 200, o.Payload)
}

func (o *ProjectsEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditBadRequest creates a ProjectsEditBadRequest with default headers values
func NewProjectsEditBadRequest() *ProjectsEditBadRequest {
	return &ProjectsEditBadRequest{}
}

/*
ProjectsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsEditBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this projects edit bad request response has a 2xx status code
func (o *ProjectsEditBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit bad request response has a 3xx status code
func (o *ProjectsEditBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit bad request response has a 4xx status code
func (o *ProjectsEditBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit bad request response has a 5xx status code
func (o *ProjectsEditBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit bad request response a status code equal to that given
func (o *ProjectsEditBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsEditBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsEditBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *ProjectsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditUnauthorized creates a ProjectsEditUnauthorized with default headers values
func NewProjectsEditUnauthorized() *ProjectsEditUnauthorized {
	return &ProjectsEditUnauthorized{}
}

/*
ProjectsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects edit unauthorized response has a 2xx status code
func (o *ProjectsEditUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit unauthorized response has a 3xx status code
func (o *ProjectsEditUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit unauthorized response has a 4xx status code
func (o *ProjectsEditUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit unauthorized response has a 5xx status code
func (o *ProjectsEditUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit unauthorized response a status code equal to that given
func (o *ProjectsEditUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsEditUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditForbidden creates a ProjectsEditForbidden with default headers values
func NewProjectsEditForbidden() *ProjectsEditForbidden {
	return &ProjectsEditForbidden{}
}

/*
ProjectsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsEditForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects edit forbidden response has a 2xx status code
func (o *ProjectsEditForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit forbidden response has a 3xx status code
func (o *ProjectsEditForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit forbidden response has a 4xx status code
func (o *ProjectsEditForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit forbidden response has a 5xx status code
func (o *ProjectsEditForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit forbidden response a status code equal to that given
func (o *ProjectsEditForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsEditForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditNotFound creates a ProjectsEditNotFound with default headers values
func NewProjectsEditNotFound() *ProjectsEditNotFound {
	return &ProjectsEditNotFound{}
}

/*
ProjectsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsEditNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects edit not found response has a 2xx status code
func (o *ProjectsEditNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit not found response has a 3xx status code
func (o *ProjectsEditNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit not found response has a 4xx status code
func (o *ProjectsEditNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects edit not found response has a 5xx status code
func (o *ProjectsEditNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects edit not found response a status code equal to that given
func (o *ProjectsEditNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsEditNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsEditInternalServerError creates a ProjectsEditInternalServerError with default headers values
func NewProjectsEditInternalServerError() *ProjectsEditInternalServerError {
	return &ProjectsEditInternalServerError{}
}

/*
ProjectsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsEditInternalServerError struct {
}

// IsSuccess returns true when this projects edit internal server error response has a 2xx status code
func (o *ProjectsEditInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects edit internal server error response has a 3xx status code
func (o *ProjectsEditInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects edit internal server error response has a 4xx status code
func (o *ProjectsEditInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects edit internal server error response has a 5xx status code
func (o *ProjectsEditInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects edit internal server error response a status code equal to that given
func (o *ProjectsEditInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditInternalServerError ", 500)
}

func (o *ProjectsEditInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/update/{projectId}][%d] projectsEditInternalServerError ", 500)
}

func (o *ProjectsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
