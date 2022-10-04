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

// PrometheusMetricNameReader is a Reader for the PrometheusMetricName structure.
type PrometheusMetricNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusMetricNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusMetricNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusMetricNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusMetricNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusMetricNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusMetricNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusMetricNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusMetricNameOK creates a PrometheusMetricNameOK with default headers values
func NewPrometheusMetricNameOK() *PrometheusMetricNameOK {
	return &PrometheusMetricNameOK{}
}

/*
PrometheusMetricNameOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusMetricNameOK struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus metric name o k response has a 2xx status code
func (o *PrometheusMetricNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus metric name o k response has a 3xx status code
func (o *PrometheusMetricNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name o k response has a 4xx status code
func (o *PrometheusMetricNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus metric name o k response has a 5xx status code
func (o *PrometheusMetricNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus metric name o k response a status code equal to that given
func (o *PrometheusMetricNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusMetricNameOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameOK  %+v", 200, o.Payload)
}

func (o *PrometheusMetricNameOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameOK  %+v", 200, o.Payload)
}

func (o *PrometheusMetricNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusMetricNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameBadRequest creates a PrometheusMetricNameBadRequest with default headers values
func NewPrometheusMetricNameBadRequest() *PrometheusMetricNameBadRequest {
	return &PrometheusMetricNameBadRequest{}
}

/*
PrometheusMetricNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusMetricNameBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this prometheus metric name bad request response has a 2xx status code
func (o *PrometheusMetricNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus metric name bad request response has a 3xx status code
func (o *PrometheusMetricNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name bad request response has a 4xx status code
func (o *PrometheusMetricNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus metric name bad request response has a 5xx status code
func (o *PrometheusMetricNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus metric name bad request response a status code equal to that given
func (o *PrometheusMetricNameBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusMetricNameBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusMetricNameBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusMetricNameBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusMetricNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameUnauthorized creates a PrometheusMetricNameUnauthorized with default headers values
func NewPrometheusMetricNameUnauthorized() *PrometheusMetricNameUnauthorized {
	return &PrometheusMetricNameUnauthorized{}
}

/*
PrometheusMetricNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusMetricNameUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus metric name unauthorized response has a 2xx status code
func (o *PrometheusMetricNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus metric name unauthorized response has a 3xx status code
func (o *PrometheusMetricNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name unauthorized response has a 4xx status code
func (o *PrometheusMetricNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus metric name unauthorized response has a 5xx status code
func (o *PrometheusMetricNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus metric name unauthorized response a status code equal to that given
func (o *PrometheusMetricNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusMetricNameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusMetricNameUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusMetricNameUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusMetricNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameForbidden creates a PrometheusMetricNameForbidden with default headers values
func NewPrometheusMetricNameForbidden() *PrometheusMetricNameForbidden {
	return &PrometheusMetricNameForbidden{}
}

/*
PrometheusMetricNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusMetricNameForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus metric name forbidden response has a 2xx status code
func (o *PrometheusMetricNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus metric name forbidden response has a 3xx status code
func (o *PrometheusMetricNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name forbidden response has a 4xx status code
func (o *PrometheusMetricNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus metric name forbidden response has a 5xx status code
func (o *PrometheusMetricNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus metric name forbidden response a status code equal to that given
func (o *PrometheusMetricNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusMetricNameForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusMetricNameForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusMetricNameForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusMetricNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameNotFound creates a PrometheusMetricNameNotFound with default headers values
func NewPrometheusMetricNameNotFound() *PrometheusMetricNameNotFound {
	return &PrometheusMetricNameNotFound{}
}

/*
PrometheusMetricNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusMetricNameNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus metric name not found response has a 2xx status code
func (o *PrometheusMetricNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus metric name not found response has a 3xx status code
func (o *PrometheusMetricNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name not found response has a 4xx status code
func (o *PrometheusMetricNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus metric name not found response has a 5xx status code
func (o *PrometheusMetricNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus metric name not found response a status code equal to that given
func (o *PrometheusMetricNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusMetricNameNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusMetricNameNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusMetricNameNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusMetricNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameInternalServerError creates a PrometheusMetricNameInternalServerError with default headers values
func NewPrometheusMetricNameInternalServerError() *PrometheusMetricNameInternalServerError {
	return &PrometheusMetricNameInternalServerError{}
}

/*
PrometheusMetricNameInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusMetricNameInternalServerError struct {
}

// IsSuccess returns true when this prometheus metric name internal server error response has a 2xx status code
func (o *PrometheusMetricNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus metric name internal server error response has a 3xx status code
func (o *PrometheusMetricNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus metric name internal server error response has a 4xx status code
func (o *PrometheusMetricNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus metric name internal server error response has a 5xx status code
func (o *PrometheusMetricNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus metric name internal server error response a status code equal to that given
func (o *PrometheusMetricNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusMetricNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameInternalServerError ", 500)
}

func (o *PrometheusMetricNameInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameInternalServerError ", 500)
}

func (o *PrometheusMetricNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
