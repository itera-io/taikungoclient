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

// PrometheusCreateReader is a Reader for the PrometheusCreate structure.
type PrometheusCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPrometheusCreateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusCreateOK creates a PrometheusCreateOK with default headers values
func NewPrometheusCreateOK() *PrometheusCreateOK {
	return &PrometheusCreateOK{}
}

/* PrometheusCreateOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusCreateOK struct {
	Payload *models.APIResponse
}

func (o *PrometheusCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateOK  %+v", 200, o.Payload)
}
func (o *PrometheusCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PrometheusCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateBadRequest creates a PrometheusCreateBadRequest with default headers values
func NewPrometheusCreateBadRequest() *PrometheusCreateBadRequest {
	return &PrometheusCreateBadRequest{}
}

/* PrometheusCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusCreateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PrometheusCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateBadRequest  %+v", 400, o.Payload)
}
func (o *PrometheusCreateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateUnauthorized creates a PrometheusCreateUnauthorized with default headers values
func NewPrometheusCreateUnauthorized() *PrometheusCreateUnauthorized {
	return &PrometheusCreateUnauthorized{}
}

/* PrometheusCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateUnauthorized  %+v", 401, o.Payload)
}
func (o *PrometheusCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateForbidden creates a PrometheusCreateForbidden with default headers values
func NewPrometheusCreateForbidden() *PrometheusCreateForbidden {
	return &PrometheusCreateForbidden{}
}

/* PrometheusCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusCreateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateForbidden  %+v", 403, o.Payload)
}
func (o *PrometheusCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateNotFound creates a PrometheusCreateNotFound with default headers values
func NewPrometheusCreateNotFound() *PrometheusCreateNotFound {
	return &PrometheusCreateNotFound{}
}

/* PrometheusCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusCreateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateNotFound  %+v", 404, o.Payload)
}
func (o *PrometheusCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateTooManyRequests creates a PrometheusCreateTooManyRequests with default headers values
func NewPrometheusCreateTooManyRequests() *PrometheusCreateTooManyRequests {
	return &PrometheusCreateTooManyRequests{}
}

/* PrometheusCreateTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type PrometheusCreateTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusCreateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateTooManyRequests  %+v", 429, o.Payload)
}
func (o *PrometheusCreateTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusCreateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusCreateInternalServerError creates a PrometheusCreateInternalServerError with default headers values
func NewPrometheusCreateInternalServerError() *PrometheusCreateInternalServerError {
	return &PrometheusCreateInternalServerError{}
}

/* PrometheusCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusCreateInternalServerError struct {
}

func (o *PrometheusCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus][%d] prometheusCreateInternalServerError ", 500)
}

func (o *PrometheusCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
