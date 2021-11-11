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

// PrometheusDeleteReader is a Reader for the PrometheusDelete structure.
type PrometheusDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusDeleteOK creates a PrometheusDeleteOK with default headers values
func NewPrometheusDeleteOK() *PrometheusDeleteOK {
	return &PrometheusDeleteOK{}
}

/* PrometheusDeleteOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusDeleteOK struct {
	Payload models.Unit
}

func (o *PrometheusDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteOK  %+v", 200, o.Payload)
}
func (o *PrometheusDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PrometheusDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteBadRequest creates a PrometheusDeleteBadRequest with default headers values
func NewPrometheusDeleteBadRequest() *PrometheusDeleteBadRequest {
	return &PrometheusDeleteBadRequest{}
}

/* PrometheusDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PrometheusDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PrometheusDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteUnauthorized creates a PrometheusDeleteUnauthorized with default headers values
func NewPrometheusDeleteUnauthorized() *PrometheusDeleteUnauthorized {
	return &PrometheusDeleteUnauthorized{}
}

/* PrometheusDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PrometheusDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteForbidden creates a PrometheusDeleteForbidden with default headers values
func NewPrometheusDeleteForbidden() *PrometheusDeleteForbidden {
	return &PrometheusDeleteForbidden{}
}

/* PrometheusDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteForbidden  %+v", 403, o.Payload)
}
func (o *PrometheusDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteNotFound creates a PrometheusDeleteNotFound with default headers values
func NewPrometheusDeleteNotFound() *PrometheusDeleteNotFound {
	return &PrometheusDeleteNotFound{}
}

/* PrometheusDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteNotFound  %+v", 404, o.Payload)
}
func (o *PrometheusDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteInternalServerError creates a PrometheusDeleteInternalServerError with default headers values
func NewPrometheusDeleteInternalServerError() *PrometheusDeleteInternalServerError {
	return &PrometheusDeleteInternalServerError{}
}

/* PrometheusDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusDeleteInternalServerError struct {
}

func (o *PrometheusDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteInternalServerError ", 500)
}

func (o *PrometheusDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
