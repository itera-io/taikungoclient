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

// ProjectsListForBillingReader is a Reader for the ProjectsListForBilling structure.
type ProjectsListForBillingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListForBillingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListForBillingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListForBillingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListForBillingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListForBillingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListForBillingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListForBillingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListForBillingOK creates a ProjectsListForBillingOK with default headers values
func NewProjectsListForBillingOK() *ProjectsListForBillingOK {
	return &ProjectsListForBillingOK{}
}

/*
ProjectsListForBillingOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListForBillingOK struct {
	Payload []*models.ProjectsForBillingDto
}

// IsSuccess returns true when this projects list for billing o k response has a 2xx status code
func (o *ProjectsListForBillingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects list for billing o k response has a 3xx status code
func (o *ProjectsListForBillingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing o k response has a 4xx status code
func (o *ProjectsListForBillingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list for billing o k response has a 5xx status code
func (o *ProjectsListForBillingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list for billing o k response a status code equal to that given
func (o *ProjectsListForBillingOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsListForBillingOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingOK  %+v", 200, o.Payload)
}

func (o *ProjectsListForBillingOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingOK  %+v", 200, o.Payload)
}

func (o *ProjectsListForBillingOK) GetPayload() []*models.ProjectsForBillingDto {
	return o.Payload
}

func (o *ProjectsListForBillingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForBillingBadRequest creates a ProjectsListForBillingBadRequest with default headers values
func NewProjectsListForBillingBadRequest() *ProjectsListForBillingBadRequest {
	return &ProjectsListForBillingBadRequest{}
}

/*
ProjectsListForBillingBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListForBillingBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects list for billing bad request response has a 2xx status code
func (o *ProjectsListForBillingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list for billing bad request response has a 3xx status code
func (o *ProjectsListForBillingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing bad request response has a 4xx status code
func (o *ProjectsListForBillingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list for billing bad request response has a 5xx status code
func (o *ProjectsListForBillingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list for billing bad request response a status code equal to that given
func (o *ProjectsListForBillingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsListForBillingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListForBillingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListForBillingBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsListForBillingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForBillingUnauthorized creates a ProjectsListForBillingUnauthorized with default headers values
func NewProjectsListForBillingUnauthorized() *ProjectsListForBillingUnauthorized {
	return &ProjectsListForBillingUnauthorized{}
}

/*
ProjectsListForBillingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListForBillingUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects list for billing unauthorized response has a 2xx status code
func (o *ProjectsListForBillingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list for billing unauthorized response has a 3xx status code
func (o *ProjectsListForBillingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing unauthorized response has a 4xx status code
func (o *ProjectsListForBillingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list for billing unauthorized response has a 5xx status code
func (o *ProjectsListForBillingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list for billing unauthorized response a status code equal to that given
func (o *ProjectsListForBillingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsListForBillingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListForBillingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListForBillingUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsListForBillingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForBillingForbidden creates a ProjectsListForBillingForbidden with default headers values
func NewProjectsListForBillingForbidden() *ProjectsListForBillingForbidden {
	return &ProjectsListForBillingForbidden{}
}

/*
ProjectsListForBillingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListForBillingForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects list for billing forbidden response has a 2xx status code
func (o *ProjectsListForBillingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list for billing forbidden response has a 3xx status code
func (o *ProjectsListForBillingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing forbidden response has a 4xx status code
func (o *ProjectsListForBillingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list for billing forbidden response has a 5xx status code
func (o *ProjectsListForBillingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list for billing forbidden response a status code equal to that given
func (o *ProjectsListForBillingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsListForBillingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListForBillingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListForBillingForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsListForBillingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForBillingNotFound creates a ProjectsListForBillingNotFound with default headers values
func NewProjectsListForBillingNotFound() *ProjectsListForBillingNotFound {
	return &ProjectsListForBillingNotFound{}
}

/*
ProjectsListForBillingNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListForBillingNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects list for billing not found response has a 2xx status code
func (o *ProjectsListForBillingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list for billing not found response has a 3xx status code
func (o *ProjectsListForBillingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing not found response has a 4xx status code
func (o *ProjectsListForBillingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list for billing not found response has a 5xx status code
func (o *ProjectsListForBillingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list for billing not found response a status code equal to that given
func (o *ProjectsListForBillingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsListForBillingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListForBillingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListForBillingNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsListForBillingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListForBillingInternalServerError creates a ProjectsListForBillingInternalServerError with default headers values
func NewProjectsListForBillingInternalServerError() *ProjectsListForBillingInternalServerError {
	return &ProjectsListForBillingInternalServerError{}
}

/*
ProjectsListForBillingInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListForBillingInternalServerError struct {
}

// IsSuccess returns true when this projects list for billing internal server error response has a 2xx status code
func (o *ProjectsListForBillingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list for billing internal server error response has a 3xx status code
func (o *ProjectsListForBillingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list for billing internal server error response has a 4xx status code
func (o *ProjectsListForBillingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list for billing internal server error response has a 5xx status code
func (o *ProjectsListForBillingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects list for billing internal server error response a status code equal to that given
func (o *ProjectsListForBillingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsListForBillingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingInternalServerError ", 500)
}

func (o *ProjectsListForBillingInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/forbilling][%d] projectsListForBillingInternalServerError ", 500)
}

func (o *ProjectsListForBillingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
