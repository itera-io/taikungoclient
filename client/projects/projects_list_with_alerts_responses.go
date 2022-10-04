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

// ProjectsListWithAlertsReader is a Reader for the ProjectsListWithAlerts structure.
type ProjectsListWithAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsListWithAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsListWithAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsListWithAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsListWithAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsListWithAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsListWithAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsListWithAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsListWithAlertsOK creates a ProjectsListWithAlertsOK with default headers values
func NewProjectsListWithAlertsOK() *ProjectsListWithAlertsOK {
	return &ProjectsListWithAlertsOK{}
}

/*
ProjectsListWithAlertsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsListWithAlertsOK struct {
	Payload []*models.ProjectListForAlert
}

// IsSuccess returns true when this projects list with alerts o k response has a 2xx status code
func (o *ProjectsListWithAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects list with alerts o k response has a 3xx status code
func (o *ProjectsListWithAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts o k response has a 4xx status code
func (o *ProjectsListWithAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list with alerts o k response has a 5xx status code
func (o *ProjectsListWithAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts o k response a status code equal to that given
func (o *ProjectsListWithAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsListWithAlertsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsOK  %+v", 200, o.Payload)
}

func (o *ProjectsListWithAlertsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsOK  %+v", 200, o.Payload)
}

func (o *ProjectsListWithAlertsOK) GetPayload() []*models.ProjectListForAlert {
	return o.Payload
}

func (o *ProjectsListWithAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsBadRequest creates a ProjectsListWithAlertsBadRequest with default headers values
func NewProjectsListWithAlertsBadRequest() *ProjectsListWithAlertsBadRequest {
	return &ProjectsListWithAlertsBadRequest{}
}

/*
ProjectsListWithAlertsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsListWithAlertsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this projects list with alerts bad request response has a 2xx status code
func (o *ProjectsListWithAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts bad request response has a 3xx status code
func (o *ProjectsListWithAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts bad request response has a 4xx status code
func (o *ProjectsListWithAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts bad request response has a 5xx status code
func (o *ProjectsListWithAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts bad request response a status code equal to that given
func (o *ProjectsListWithAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsListWithAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListWithAlertsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsListWithAlertsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectsListWithAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsUnauthorized creates a ProjectsListWithAlertsUnauthorized with default headers values
func NewProjectsListWithAlertsUnauthorized() *ProjectsListWithAlertsUnauthorized {
	return &ProjectsListWithAlertsUnauthorized{}
}

/*
ProjectsListWithAlertsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsListWithAlertsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this projects list with alerts unauthorized response has a 2xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts unauthorized response has a 3xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts unauthorized response has a 4xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts unauthorized response has a 5xx status code
func (o *ProjectsListWithAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts unauthorized response a status code equal to that given
func (o *ProjectsListWithAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsListWithAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListWithAlertsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsListWithAlertsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsListWithAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsForbidden creates a ProjectsListWithAlertsForbidden with default headers values
func NewProjectsListWithAlertsForbidden() *ProjectsListWithAlertsForbidden {
	return &ProjectsListWithAlertsForbidden{}
}

/*
ProjectsListWithAlertsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsListWithAlertsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this projects list with alerts forbidden response has a 2xx status code
func (o *ProjectsListWithAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts forbidden response has a 3xx status code
func (o *ProjectsListWithAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts forbidden response has a 4xx status code
func (o *ProjectsListWithAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts forbidden response has a 5xx status code
func (o *ProjectsListWithAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts forbidden response a status code equal to that given
func (o *ProjectsListWithAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsListWithAlertsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListWithAlertsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsListWithAlertsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsListWithAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsNotFound creates a ProjectsListWithAlertsNotFound with default headers values
func NewProjectsListWithAlertsNotFound() *ProjectsListWithAlertsNotFound {
	return &ProjectsListWithAlertsNotFound{}
}

/*
ProjectsListWithAlertsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsListWithAlertsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this projects list with alerts not found response has a 2xx status code
func (o *ProjectsListWithAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts not found response has a 3xx status code
func (o *ProjectsListWithAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts not found response has a 4xx status code
func (o *ProjectsListWithAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects list with alerts not found response has a 5xx status code
func (o *ProjectsListWithAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects list with alerts not found response a status code equal to that given
func (o *ProjectsListWithAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsListWithAlertsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListWithAlertsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsListWithAlertsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsListWithAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsListWithAlertsInternalServerError creates a ProjectsListWithAlertsInternalServerError with default headers values
func NewProjectsListWithAlertsInternalServerError() *ProjectsListWithAlertsInternalServerError {
	return &ProjectsListWithAlertsInternalServerError{}
}

/*
ProjectsListWithAlertsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsListWithAlertsInternalServerError struct {
}

// IsSuccess returns true when this projects list with alerts internal server error response has a 2xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects list with alerts internal server error response has a 3xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects list with alerts internal server error response has a 4xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects list with alerts internal server error response has a 5xx status code
func (o *ProjectsListWithAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects list with alerts internal server error response a status code equal to that given
func (o *ProjectsListWithAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsListWithAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsInternalServerError ", 500)
}

func (o *ProjectsListWithAlertsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Projects/foralerting][%d] projectsListWithAlertsInternalServerError ", 500)
}

func (o *ProjectsListWithAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
