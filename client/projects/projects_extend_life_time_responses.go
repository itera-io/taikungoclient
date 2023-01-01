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

// ProjectsExtendLifeTimeReader is a Reader for the ProjectsExtendLifeTime structure.
type ProjectsExtendLifeTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsExtendLifeTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsExtendLifeTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsExtendLifeTimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsExtendLifeTimeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsExtendLifeTimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsExtendLifeTimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsExtendLifeTimeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsExtendLifeTimeOK creates a ProjectsExtendLifeTimeOK with default headers values
func NewProjectsExtendLifeTimeOK() *ProjectsExtendLifeTimeOK {
	return &ProjectsExtendLifeTimeOK{}
}

/*
ProjectsExtendLifeTimeOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsExtendLifeTimeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this projects extend life time o k response has a 2xx status code
func (o *ProjectsExtendLifeTimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects extend life time o k response has a 3xx status code
func (o *ProjectsExtendLifeTimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time o k response has a 4xx status code
func (o *ProjectsExtendLifeTimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects extend life time o k response has a 5xx status code
func (o *ProjectsExtendLifeTimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects extend life time o k response a status code equal to that given
func (o *ProjectsExtendLifeTimeOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsExtendLifeTimeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeOK  %+v", 200, o.Payload)
}

func (o *ProjectsExtendLifeTimeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeOK  %+v", 200, o.Payload)
}

func (o *ProjectsExtendLifeTimeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectsExtendLifeTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsExtendLifeTimeBadRequest creates a ProjectsExtendLifeTimeBadRequest with default headers values
func NewProjectsExtendLifeTimeBadRequest() *ProjectsExtendLifeTimeBadRequest {
	return &ProjectsExtendLifeTimeBadRequest{}
}

/*
ProjectsExtendLifeTimeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsExtendLifeTimeBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects extend life time bad request response has a 2xx status code
func (o *ProjectsExtendLifeTimeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects extend life time bad request response has a 3xx status code
func (o *ProjectsExtendLifeTimeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time bad request response has a 4xx status code
func (o *ProjectsExtendLifeTimeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects extend life time bad request response has a 5xx status code
func (o *ProjectsExtendLifeTimeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects extend life time bad request response a status code equal to that given
func (o *ProjectsExtendLifeTimeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsExtendLifeTimeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsExtendLifeTimeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsExtendLifeTimeBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsExtendLifeTimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsExtendLifeTimeUnauthorized creates a ProjectsExtendLifeTimeUnauthorized with default headers values
func NewProjectsExtendLifeTimeUnauthorized() *ProjectsExtendLifeTimeUnauthorized {
	return &ProjectsExtendLifeTimeUnauthorized{}
}

/*
ProjectsExtendLifeTimeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsExtendLifeTimeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects extend life time unauthorized response has a 2xx status code
func (o *ProjectsExtendLifeTimeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects extend life time unauthorized response has a 3xx status code
func (o *ProjectsExtendLifeTimeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time unauthorized response has a 4xx status code
func (o *ProjectsExtendLifeTimeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects extend life time unauthorized response has a 5xx status code
func (o *ProjectsExtendLifeTimeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects extend life time unauthorized response a status code equal to that given
func (o *ProjectsExtendLifeTimeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsExtendLifeTimeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsExtendLifeTimeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsExtendLifeTimeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsExtendLifeTimeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsExtendLifeTimeForbidden creates a ProjectsExtendLifeTimeForbidden with default headers values
func NewProjectsExtendLifeTimeForbidden() *ProjectsExtendLifeTimeForbidden {
	return &ProjectsExtendLifeTimeForbidden{}
}

/*
ProjectsExtendLifeTimeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsExtendLifeTimeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects extend life time forbidden response has a 2xx status code
func (o *ProjectsExtendLifeTimeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects extend life time forbidden response has a 3xx status code
func (o *ProjectsExtendLifeTimeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time forbidden response has a 4xx status code
func (o *ProjectsExtendLifeTimeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects extend life time forbidden response has a 5xx status code
func (o *ProjectsExtendLifeTimeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects extend life time forbidden response a status code equal to that given
func (o *ProjectsExtendLifeTimeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsExtendLifeTimeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsExtendLifeTimeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsExtendLifeTimeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsExtendLifeTimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsExtendLifeTimeNotFound creates a ProjectsExtendLifeTimeNotFound with default headers values
func NewProjectsExtendLifeTimeNotFound() *ProjectsExtendLifeTimeNotFound {
	return &ProjectsExtendLifeTimeNotFound{}
}

/*
ProjectsExtendLifeTimeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsExtendLifeTimeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects extend life time not found response has a 2xx status code
func (o *ProjectsExtendLifeTimeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects extend life time not found response has a 3xx status code
func (o *ProjectsExtendLifeTimeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time not found response has a 4xx status code
func (o *ProjectsExtendLifeTimeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects extend life time not found response has a 5xx status code
func (o *ProjectsExtendLifeTimeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects extend life time not found response a status code equal to that given
func (o *ProjectsExtendLifeTimeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsExtendLifeTimeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsExtendLifeTimeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsExtendLifeTimeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsExtendLifeTimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsExtendLifeTimeInternalServerError creates a ProjectsExtendLifeTimeInternalServerError with default headers values
func NewProjectsExtendLifeTimeInternalServerError() *ProjectsExtendLifeTimeInternalServerError {
	return &ProjectsExtendLifeTimeInternalServerError{}
}

/*
ProjectsExtendLifeTimeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsExtendLifeTimeInternalServerError struct {
}

// IsSuccess returns true when this projects extend life time internal server error response has a 2xx status code
func (o *ProjectsExtendLifeTimeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects extend life time internal server error response has a 3xx status code
func (o *ProjectsExtendLifeTimeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects extend life time internal server error response has a 4xx status code
func (o *ProjectsExtendLifeTimeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects extend life time internal server error response has a 5xx status code
func (o *ProjectsExtendLifeTimeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects extend life time internal server error response a status code equal to that given
func (o *ProjectsExtendLifeTimeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsExtendLifeTimeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeInternalServerError ", 500)
}

func (o *ProjectsExtendLifeTimeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/extend/lifetime][%d] projectsExtendLifeTimeInternalServerError ", 500)
}

func (o *ProjectsExtendLifeTimeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
