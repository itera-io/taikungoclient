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

// ProjectsLokiLogsReader is a Reader for the ProjectsLokiLogs structure.
type ProjectsLokiLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsLokiLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsLokiLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsLokiLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsLokiLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsLokiLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsLokiLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsLokiLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsLokiLogsOK creates a ProjectsLokiLogsOK with default headers values
func NewProjectsLokiLogsOK() *ProjectsLokiLogsOK {
	return &ProjectsLokiLogsOK{}
}

/*
ProjectsLokiLogsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsLokiLogsOK struct {
}

// IsSuccess returns true when this projects loki logs o k response has a 2xx status code
func (o *ProjectsLokiLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects loki logs o k response has a 3xx status code
func (o *ProjectsLokiLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs o k response has a 4xx status code
func (o *ProjectsLokiLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects loki logs o k response has a 5xx status code
func (o *ProjectsLokiLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects loki logs o k response a status code equal to that given
func (o *ProjectsLokiLogsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsLokiLogsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsOK ", 200)
}

func (o *ProjectsLokiLogsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsOK ", 200)
}

func (o *ProjectsLokiLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProjectsLokiLogsBadRequest creates a ProjectsLokiLogsBadRequest with default headers values
func NewProjectsLokiLogsBadRequest() *ProjectsLokiLogsBadRequest {
	return &ProjectsLokiLogsBadRequest{}
}

/*
ProjectsLokiLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsLokiLogsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects loki logs bad request response has a 2xx status code
func (o *ProjectsLokiLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects loki logs bad request response has a 3xx status code
func (o *ProjectsLokiLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs bad request response has a 4xx status code
func (o *ProjectsLokiLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects loki logs bad request response has a 5xx status code
func (o *ProjectsLokiLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects loki logs bad request response a status code equal to that given
func (o *ProjectsLokiLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsLokiLogsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsLokiLogsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsLokiLogsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsLokiLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsLokiLogsUnauthorized creates a ProjectsLokiLogsUnauthorized with default headers values
func NewProjectsLokiLogsUnauthorized() *ProjectsLokiLogsUnauthorized {
	return &ProjectsLokiLogsUnauthorized{}
}

/*
ProjectsLokiLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsLokiLogsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects loki logs unauthorized response has a 2xx status code
func (o *ProjectsLokiLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects loki logs unauthorized response has a 3xx status code
func (o *ProjectsLokiLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs unauthorized response has a 4xx status code
func (o *ProjectsLokiLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects loki logs unauthorized response has a 5xx status code
func (o *ProjectsLokiLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects loki logs unauthorized response a status code equal to that given
func (o *ProjectsLokiLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsLokiLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsLokiLogsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsLokiLogsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsLokiLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsLokiLogsForbidden creates a ProjectsLokiLogsForbidden with default headers values
func NewProjectsLokiLogsForbidden() *ProjectsLokiLogsForbidden {
	return &ProjectsLokiLogsForbidden{}
}

/*
ProjectsLokiLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsLokiLogsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects loki logs forbidden response has a 2xx status code
func (o *ProjectsLokiLogsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects loki logs forbidden response has a 3xx status code
func (o *ProjectsLokiLogsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs forbidden response has a 4xx status code
func (o *ProjectsLokiLogsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects loki logs forbidden response has a 5xx status code
func (o *ProjectsLokiLogsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects loki logs forbidden response a status code equal to that given
func (o *ProjectsLokiLogsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsLokiLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsLokiLogsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsLokiLogsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsLokiLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsLokiLogsNotFound creates a ProjectsLokiLogsNotFound with default headers values
func NewProjectsLokiLogsNotFound() *ProjectsLokiLogsNotFound {
	return &ProjectsLokiLogsNotFound{}
}

/*
ProjectsLokiLogsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsLokiLogsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this projects loki logs not found response has a 2xx status code
func (o *ProjectsLokiLogsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects loki logs not found response has a 3xx status code
func (o *ProjectsLokiLogsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs not found response has a 4xx status code
func (o *ProjectsLokiLogsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects loki logs not found response has a 5xx status code
func (o *ProjectsLokiLogsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects loki logs not found response a status code equal to that given
func (o *ProjectsLokiLogsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsLokiLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsLokiLogsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsLokiLogsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *ProjectsLokiLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsLokiLogsInternalServerError creates a ProjectsLokiLogsInternalServerError with default headers values
func NewProjectsLokiLogsInternalServerError() *ProjectsLokiLogsInternalServerError {
	return &ProjectsLokiLogsInternalServerError{}
}

/*
ProjectsLokiLogsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsLokiLogsInternalServerError struct {
}

// IsSuccess returns true when this projects loki logs internal server error response has a 2xx status code
func (o *ProjectsLokiLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects loki logs internal server error response has a 3xx status code
func (o *ProjectsLokiLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects loki logs internal server error response has a 4xx status code
func (o *ProjectsLokiLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects loki logs internal server error response has a 5xx status code
func (o *ProjectsLokiLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects loki logs internal server error response a status code equal to that given
func (o *ProjectsLokiLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsLokiLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsInternalServerError ", 500)
}

func (o *ProjectsLokiLogsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/lokilogs][%d] projectsLokiLogsInternalServerError ", 500)
}

func (o *ProjectsLokiLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
