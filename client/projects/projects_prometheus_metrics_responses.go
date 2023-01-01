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

// ProjectsPrometheusMetricsReader is a Reader for the ProjectsPrometheusMetrics structure.
type ProjectsPrometheusMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsPrometheusMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsPrometheusMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsPrometheusMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsPrometheusMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsPrometheusMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectsPrometheusMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsPrometheusMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectsPrometheusMetricsOK creates a ProjectsPrometheusMetricsOK with default headers values
func NewProjectsPrometheusMetricsOK() *ProjectsPrometheusMetricsOK {
	return &ProjectsPrometheusMetricsOK{}
}

/*
ProjectsPrometheusMetricsOK describes a response with status code 200, with default header values.

Success
*/
type ProjectsPrometheusMetricsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this projects prometheus metrics o k response has a 2xx status code
func (o *ProjectsPrometheusMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects prometheus metrics o k response has a 3xx status code
func (o *ProjectsPrometheusMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics o k response has a 4xx status code
func (o *ProjectsPrometheusMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects prometheus metrics o k response has a 5xx status code
func (o *ProjectsPrometheusMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects prometheus metrics o k response a status code equal to that given
func (o *ProjectsPrometheusMetricsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsPrometheusMetricsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsOK  %+v", 200, o.Payload)
}

func (o *ProjectsPrometheusMetricsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsOK  %+v", 200, o.Payload)
}

func (o *ProjectsPrometheusMetricsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectsPrometheusMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPrometheusMetricsBadRequest creates a ProjectsPrometheusMetricsBadRequest with default headers values
func NewProjectsPrometheusMetricsBadRequest() *ProjectsPrometheusMetricsBadRequest {
	return &ProjectsPrometheusMetricsBadRequest{}
}

/*
ProjectsPrometheusMetricsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectsPrometheusMetricsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects prometheus metrics bad request response has a 2xx status code
func (o *ProjectsPrometheusMetricsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects prometheus metrics bad request response has a 3xx status code
func (o *ProjectsPrometheusMetricsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics bad request response has a 4xx status code
func (o *ProjectsPrometheusMetricsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects prometheus metrics bad request response has a 5xx status code
func (o *ProjectsPrometheusMetricsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects prometheus metrics bad request response a status code equal to that given
func (o *ProjectsPrometheusMetricsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsPrometheusMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsPrometheusMetricsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *ProjectsPrometheusMetricsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPrometheusMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPrometheusMetricsUnauthorized creates a ProjectsPrometheusMetricsUnauthorized with default headers values
func NewProjectsPrometheusMetricsUnauthorized() *ProjectsPrometheusMetricsUnauthorized {
	return &ProjectsPrometheusMetricsUnauthorized{}
}

/*
ProjectsPrometheusMetricsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectsPrometheusMetricsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects prometheus metrics unauthorized response has a 2xx status code
func (o *ProjectsPrometheusMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects prometheus metrics unauthorized response has a 3xx status code
func (o *ProjectsPrometheusMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics unauthorized response has a 4xx status code
func (o *ProjectsPrometheusMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects prometheus metrics unauthorized response has a 5xx status code
func (o *ProjectsPrometheusMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects prometheus metrics unauthorized response a status code equal to that given
func (o *ProjectsPrometheusMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsPrometheusMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsPrometheusMetricsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsPrometheusMetricsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPrometheusMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPrometheusMetricsForbidden creates a ProjectsPrometheusMetricsForbidden with default headers values
func NewProjectsPrometheusMetricsForbidden() *ProjectsPrometheusMetricsForbidden {
	return &ProjectsPrometheusMetricsForbidden{}
}

/*
ProjectsPrometheusMetricsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectsPrometheusMetricsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects prometheus metrics forbidden response has a 2xx status code
func (o *ProjectsPrometheusMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects prometheus metrics forbidden response has a 3xx status code
func (o *ProjectsPrometheusMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics forbidden response has a 4xx status code
func (o *ProjectsPrometheusMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects prometheus metrics forbidden response has a 5xx status code
func (o *ProjectsPrometheusMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects prometheus metrics forbidden response a status code equal to that given
func (o *ProjectsPrometheusMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsPrometheusMetricsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsPrometheusMetricsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsPrometheusMetricsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPrometheusMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPrometheusMetricsNotFound creates a ProjectsPrometheusMetricsNotFound with default headers values
func NewProjectsPrometheusMetricsNotFound() *ProjectsPrometheusMetricsNotFound {
	return &ProjectsPrometheusMetricsNotFound{}
}

/*
ProjectsPrometheusMetricsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectsPrometheusMetricsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this projects prometheus metrics not found response has a 2xx status code
func (o *ProjectsPrometheusMetricsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects prometheus metrics not found response has a 3xx status code
func (o *ProjectsPrometheusMetricsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics not found response has a 4xx status code
func (o *ProjectsPrometheusMetricsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects prometheus metrics not found response has a 5xx status code
func (o *ProjectsPrometheusMetricsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this projects prometheus metrics not found response a status code equal to that given
func (o *ProjectsPrometheusMetricsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ProjectsPrometheusMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsPrometheusMetricsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsNotFound  %+v", 404, o.Payload)
}

func (o *ProjectsPrometheusMetricsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectsPrometheusMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsPrometheusMetricsInternalServerError creates a ProjectsPrometheusMetricsInternalServerError with default headers values
func NewProjectsPrometheusMetricsInternalServerError() *ProjectsPrometheusMetricsInternalServerError {
	return &ProjectsPrometheusMetricsInternalServerError{}
}

/*
ProjectsPrometheusMetricsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectsPrometheusMetricsInternalServerError struct {
}

// IsSuccess returns true when this projects prometheus metrics internal server error response has a 2xx status code
func (o *ProjectsPrometheusMetricsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects prometheus metrics internal server error response has a 3xx status code
func (o *ProjectsPrometheusMetricsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects prometheus metrics internal server error response has a 4xx status code
func (o *ProjectsPrometheusMetricsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects prometheus metrics internal server error response has a 5xx status code
func (o *ProjectsPrometheusMetricsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects prometheus metrics internal server error response a status code equal to that given
func (o *ProjectsPrometheusMetricsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsPrometheusMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsInternalServerError ", 500)
}

func (o *ProjectsPrometheusMetricsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Projects/prometheusmetrics][%d] projectsPrometheusMetricsInternalServerError ", 500)
}

func (o *ProjectsPrometheusMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
