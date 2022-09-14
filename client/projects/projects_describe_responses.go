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

// ProjectsDescribeReader is a Reader for the ProjectsDescribe structure.
type ProjectsDescribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDescribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsDescribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsDescribeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsDescribeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsDescribeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsDescribeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsDescribeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsDescribeOK creates a ProjectsDescribeOK with default headers values
func NewProjectsDescribeOK() *ProjectsDescribeOK {
	return &ProjectsDescribeOK{}
}

/*
ProjectsDescribeOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsDescribeOK struct {
	Payload interface{}
}

// IsSuccess returns true when this projects describe o k response has a 2xx status code
func (o *ProjectsDescribeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects describe o k response has a 3xx status code
func (o *ProjectsDescribeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe o k response has a 4xx status code
func (o *ProjectsDescribeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects describe o k response has a 5xx status code
func (o *ProjectsDescribeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects describe o k response a status code equal to that given
func (o *ProjectsDescribeOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsDescribeOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeOK  %+v", 200, o.Payload)
}

func (o *ProjectsDescribeOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeOK  %+v", 200, o.Payload)
}

func (o *ProjectsDescribeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDescribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDescribeBadRequest creates a ProjectsDescribeBadRequest with default headers values
func NewProjectsDescribeBadRequest() *ProjectsDescribeBadRequest {
	return &ProjectsDescribeBadRequest{}
}

/*
ProjectsDescribeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsDescribeBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects describe bad request response has a 2xx status code
func (o *ProjectsDescribeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects describe bad request response has a 3xx status code
func (o *ProjectsDescribeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe bad request response has a 4xx status code
func (o *ProjectsDescribeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects describe bad request response has a 5xx status code
func (o *ProjectsDescribeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects describe bad request response a status code equal to that given
func (o *ProjectsDescribeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsDescribeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDescribeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDescribeBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsDescribeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDescribeUnauthorized creates a ProjectsDescribeUnauthorized with default headers values
func NewProjectsDescribeUnauthorized() *ProjectsDescribeUnauthorized {
	return &ProjectsDescribeUnauthorized{}
}

/*
ProjectsDescribeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsDescribeUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects describe unauthorized response has a 2xx status code
func (o *ProjectsDescribeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects describe unauthorized response has a 3xx status code
func (o *ProjectsDescribeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe unauthorized response has a 4xx status code
func (o *ProjectsDescribeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects describe unauthorized response has a 5xx status code
func (o *ProjectsDescribeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects describe unauthorized response a status code equal to that given
func (o *ProjectsDescribeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsDescribeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDescribeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDescribeUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDescribeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDescribeForbidden creates a ProjectsDescribeForbidden with default headers values
func NewProjectsDescribeForbidden() *ProjectsDescribeForbidden {
	return &ProjectsDescribeForbidden{}
}

/*
ProjectsDescribeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsDescribeForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects describe forbidden response has a 2xx status code
func (o *ProjectsDescribeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects describe forbidden response has a 3xx status code
func (o *ProjectsDescribeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe forbidden response has a 4xx status code
func (o *ProjectsDescribeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects describe forbidden response has a 5xx status code
func (o *ProjectsDescribeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects describe forbidden response a status code equal to that given
func (o *ProjectsDescribeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsDescribeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDescribeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDescribeForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDescribeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDescribeNotFound creates a ProjectsDescribeNotFound with default headers values
func NewProjectsDescribeNotFound() *ProjectsDescribeNotFound {
	return &ProjectsDescribeNotFound{}
}

/*
ProjectsDescribeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsDescribeNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects describe not found response has a 2xx status code
func (o *ProjectsDescribeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects describe not found response has a 3xx status code
func (o *ProjectsDescribeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe not found response has a 4xx status code
func (o *ProjectsDescribeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects describe not found response has a 5xx status code
func (o *ProjectsDescribeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects describe not found response a status code equal to that given
func (o *ProjectsDescribeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsDescribeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDescribeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDescribeNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsDescribeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDescribeInternalServerError creates a ProjectsDescribeInternalServerError with default headers values
func NewProjectsDescribeInternalServerError() *ProjectsDescribeInternalServerError {
	return &ProjectsDescribeInternalServerError{}
}

/*
ProjectsDescribeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsDescribeInternalServerError struct {
}

// IsSuccess returns true when this projects describe internal server error response has a 2xx status code
func (o *ProjectsDescribeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects describe internal server error response has a 3xx status code
func (o *ProjectsDescribeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects describe internal server error response has a 4xx status code
func (o *ProjectsDescribeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects describe internal server error response has a 5xx status code
func (o *ProjectsDescribeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects describe internal server error response a status code equal to that given
func (o *ProjectsDescribeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsDescribeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeInternalServerError ", 500)
}

func (o *ProjectsDescribeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/describe/{projectId}][%d] projectsDescribeInternalServerError ", 500)
}

func (o *ProjectsDescribeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
