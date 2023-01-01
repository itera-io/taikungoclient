// Code generated by go-swagger; DO NOT EDIT.

package prometheus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// PrometheusExportCsvReader is a Reader for the PrometheusExportCsv structure.
type PrometheusExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusExportCsvOK creates a PrometheusExportCsvOK with default headers values
func NewPrometheusExportCsvOK() *PrometheusExportCsvOK {
	return &PrometheusExportCsvOK{}
}

/*
PrometheusExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusExportCsvOK struct {
}

// IsSuccess returns true when this prometheus export csv o k response has a 2xx status code
func (o *PrometheusExportCsvOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus export csv o k response has a 3xx status code
func (o *PrometheusExportCsvOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv o k response has a 4xx status code
func (o *PrometheusExportCsvOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus export csv o k response has a 5xx status code
func (o *PrometheusExportCsvOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus export csv o k response a status code equal to that given
func (o *PrometheusExportCsvOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvOK ", 200)
}

func (o *PrometheusExportCsvOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvOK ", 200)
}

func (o *PrometheusExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrometheusExportCsvBadRequest creates a PrometheusExportCsvBadRequest with default headers values
func NewPrometheusExportCsvBadRequest() *PrometheusExportCsvBadRequest {
	return &PrometheusExportCsvBadRequest{}
}

/*
PrometheusExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusExportCsvBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus export csv bad request response has a 2xx status code
func (o *PrometheusExportCsvBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus export csv bad request response has a 3xx status code
func (o *PrometheusExportCsvBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv bad request response has a 4xx status code
func (o *PrometheusExportCsvBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus export csv bad request response has a 5xx status code
func (o *PrometheusExportCsvBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus export csv bad request response a status code equal to that given
func (o *PrometheusExportCsvBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusExportCsvBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusExportCsvBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusExportCsvUnauthorized creates a PrometheusExportCsvUnauthorized with default headers values
func NewPrometheusExportCsvUnauthorized() *PrometheusExportCsvUnauthorized {
	return &PrometheusExportCsvUnauthorized{}
}

/*
PrometheusExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus export csv unauthorized response has a 2xx status code
func (o *PrometheusExportCsvUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus export csv unauthorized response has a 3xx status code
func (o *PrometheusExportCsvUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv unauthorized response has a 4xx status code
func (o *PrometheusExportCsvUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus export csv unauthorized response has a 5xx status code
func (o *PrometheusExportCsvUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus export csv unauthorized response a status code equal to that given
func (o *PrometheusExportCsvUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusExportCsvUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusExportCsvForbidden creates a PrometheusExportCsvForbidden with default headers values
func NewPrometheusExportCsvForbidden() *PrometheusExportCsvForbidden {
	return &PrometheusExportCsvForbidden{}
}

/*
PrometheusExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus export csv forbidden response has a 2xx status code
func (o *PrometheusExportCsvForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus export csv forbidden response has a 3xx status code
func (o *PrometheusExportCsvForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv forbidden response has a 4xx status code
func (o *PrometheusExportCsvForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus export csv forbidden response has a 5xx status code
func (o *PrometheusExportCsvForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus export csv forbidden response a status code equal to that given
func (o *PrometheusExportCsvForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusExportCsvForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusExportCsvNotFound creates a PrometheusExportCsvNotFound with default headers values
func NewPrometheusExportCsvNotFound() *PrometheusExportCsvNotFound {
	return &PrometheusExportCsvNotFound{}
}

/*
PrometheusExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus export csv not found response has a 2xx status code
func (o *PrometheusExportCsvNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus export csv not found response has a 3xx status code
func (o *PrometheusExportCsvNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv not found response has a 4xx status code
func (o *PrometheusExportCsvNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus export csv not found response has a 5xx status code
func (o *PrometheusExportCsvNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus export csv not found response a status code equal to that given
func (o *PrometheusExportCsvNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusExportCsvNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusExportCsvInternalServerError creates a PrometheusExportCsvInternalServerError with default headers values
func NewPrometheusExportCsvInternalServerError() *PrometheusExportCsvInternalServerError {
	return &PrometheusExportCsvInternalServerError{}
}

/*
PrometheusExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusExportCsvInternalServerError struct {
}

// IsSuccess returns true when this prometheus export csv internal server error response has a 2xx status code
func (o *PrometheusExportCsvInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus export csv internal server error response has a 3xx status code
func (o *PrometheusExportCsvInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus export csv internal server error response has a 4xx status code
func (o *PrometheusExportCsvInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus export csv internal server error response has a 5xx status code
func (o *PrometheusExportCsvInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus export csv internal server error response a status code equal to that given
func (o *PrometheusExportCsvInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvInternalServerError ", 500)
}

func (o *PrometheusExportCsvInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/export][%d] prometheusExportCsvInternalServerError ", 500)
}

func (o *PrometheusExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
