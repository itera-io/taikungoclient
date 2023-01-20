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

// ProjectsDetailsReader is a Reader for the ProjectsDetails structure.
type ProjectsDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsDetailsOK creates a ProjectsDetailsOK with default headers values
func NewProjectsDetailsOK() *ProjectsDetailsOK {
	return &ProjectsDetailsOK{}
}

/*
ProjectsDetailsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsDetailsOK struct {
	Payload *models.ProjectForListDto
}

// IsSuccess returns true when this projects details o k response has a 2xx status code
func (o *ProjectsDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects details o k response has a 3xx status code
func (o *ProjectsDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details o k response has a 4xx status code
func (o *ProjectsDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects details o k response has a 5xx status code
func (o *ProjectsDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects details o k response a status code equal to that given
func (o *ProjectsDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the projects details o k response
func (o *ProjectsDetailsOK) Code() int {
	return 200
}

func (o *ProjectsDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsOK  %+v", 200, o.Payload)
}

func (o *ProjectsDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsOK  %+v", 200, o.Payload)
}

func (o *ProjectsDetailsOK) GetPayload() *models.ProjectForListDto {
	return o.Payload
}

func (o *ProjectsDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectForListDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDetailsBadRequest creates a ProjectsDetailsBadRequest with default headers values
func NewProjectsDetailsBadRequest() *ProjectsDetailsBadRequest {
	return &ProjectsDetailsBadRequest{}
}

/*
ProjectsDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsDetailsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this projects details bad request response has a 2xx status code
func (o *ProjectsDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects details bad request response has a 3xx status code
func (o *ProjectsDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details bad request response has a 4xx status code
func (o *ProjectsDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects details bad request response has a 5xx status code
func (o *ProjectsDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects details bad request response a status code equal to that given
func (o *ProjectsDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the projects details bad request response
func (o *ProjectsDetailsBadRequest) Code() int {
	return 400
}

func (o *ProjectsDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsDetailsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDetailsUnauthorized creates a ProjectsDetailsUnauthorized with default headers values
func NewProjectsDetailsUnauthorized() *ProjectsDetailsUnauthorized {
	return &ProjectsDetailsUnauthorized{}
}

/*
ProjectsDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsDetailsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this projects details unauthorized response has a 2xx status code
func (o *ProjectsDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects details unauthorized response has a 3xx status code
func (o *ProjectsDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details unauthorized response has a 4xx status code
func (o *ProjectsDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects details unauthorized response has a 5xx status code
func (o *ProjectsDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects details unauthorized response a status code equal to that given
func (o *ProjectsDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the projects details unauthorized response
func (o *ProjectsDetailsUnauthorized) Code() int {
	return 401
}

func (o *ProjectsDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsDetailsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDetailsForbidden creates a ProjectsDetailsForbidden with default headers values
func NewProjectsDetailsForbidden() *ProjectsDetailsForbidden {
	return &ProjectsDetailsForbidden{}
}

/*
ProjectsDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsDetailsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this projects details forbidden response has a 2xx status code
func (o *ProjectsDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects details forbidden response has a 3xx status code
func (o *ProjectsDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details forbidden response has a 4xx status code
func (o *ProjectsDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects details forbidden response has a 5xx status code
func (o *ProjectsDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects details forbidden response a status code equal to that given
func (o *ProjectsDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the projects details forbidden response
func (o *ProjectsDetailsForbidden) Code() int {
	return 403
}

func (o *ProjectsDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsDetailsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDetailsNotFound creates a ProjectsDetailsNotFound with default headers values
func NewProjectsDetailsNotFound() *ProjectsDetailsNotFound {
	return &ProjectsDetailsNotFound{}
}

/*
ProjectsDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsDetailsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this projects details not found response has a 2xx status code
func (o *ProjectsDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects details not found response has a 3xx status code
func (o *ProjectsDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details not found response has a 4xx status code
func (o *ProjectsDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects details not found response has a 5xx status code
func (o *ProjectsDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects details not found response a status code equal to that given
func (o *ProjectsDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the projects details not found response
func (o *ProjectsDetailsNotFound) Code() int {
	return 404
}

func (o *ProjectsDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsDetailsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDetailsInternalServerError creates a ProjectsDetailsInternalServerError with default headers values
func NewProjectsDetailsInternalServerError() *ProjectsDetailsInternalServerError {
	return &ProjectsDetailsInternalServerError{}
}

/*
ProjectsDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsDetailsInternalServerError struct {
}

// IsSuccess returns true when this projects details internal server error response has a 2xx status code
func (o *ProjectsDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects details internal server error response has a 3xx status code
func (o *ProjectsDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects details internal server error response has a 4xx status code
func (o *ProjectsDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects details internal server error response has a 5xx status code
func (o *ProjectsDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects details internal server error response a status code equal to that given
func (o *ProjectsDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the projects details internal server error response
func (o *ProjectsDetailsInternalServerError) Code() int {
	return 500
}

func (o *ProjectsDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsInternalServerError ", 500)
}

func (o *ProjectsDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/{projectId}][%d] projectsDetailsInternalServerError ", 500)
}

func (o *ProjectsDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
