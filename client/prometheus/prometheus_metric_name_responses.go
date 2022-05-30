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

/* PrometheusMetricNameOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusMetricNameOK struct {
	Payload interface{}
}

func (o *PrometheusMetricNameOK) Error() string {
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

/* PrometheusMetricNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusMetricNameBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PrometheusMetricNameBadRequest) Error() string {
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

/* PrometheusMetricNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusMetricNameUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusMetricNameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameUnauthorized  %+v", 401, o.Payload)
}
func (o *PrometheusMetricNameUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusMetricNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameForbidden creates a PrometheusMetricNameForbidden with default headers values
func NewPrometheusMetricNameForbidden() *PrometheusMetricNameForbidden {
	return &PrometheusMetricNameForbidden{}
}

/* PrometheusMetricNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusMetricNameForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusMetricNameForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameForbidden  %+v", 403, o.Payload)
}
func (o *PrometheusMetricNameForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusMetricNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameNotFound creates a PrometheusMetricNameNotFound with default headers values
func NewPrometheusMetricNameNotFound() *PrometheusMetricNameNotFound {
	return &PrometheusMetricNameNotFound{}
}

/* PrometheusMetricNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusMetricNameNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusMetricNameNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameNotFound  %+v", 404, o.Payload)
}
func (o *PrometheusMetricNameNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusMetricNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusMetricNameInternalServerError creates a PrometheusMetricNameInternalServerError with default headers values
func NewPrometheusMetricNameInternalServerError() *PrometheusMetricNameInternalServerError {
	return &PrometheusMetricNameInternalServerError{}
}

/* PrometheusMetricNameInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusMetricNameInternalServerError struct {
}

func (o *PrometheusMetricNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/metricname][%d] prometheusMetricNameInternalServerError ", 500)
}

func (o *PrometheusMetricNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}