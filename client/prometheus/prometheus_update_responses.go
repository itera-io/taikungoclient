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

// PrometheusUpdateReader is a Reader for the PrometheusUpdate structure.
type PrometheusUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusUpdateOK creates a PrometheusUpdateOK with default headers values
func NewPrometheusUpdateOK() *PrometheusUpdateOK {
	return &PrometheusUpdateOK{}
}

/*
PrometheusUpdateOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusUpdateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this prometheus update o k response has a 2xx status code
func (o *PrometheusUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus update o k response has a 3xx status code
func (o *PrometheusUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update o k response has a 4xx status code
func (o *PrometheusUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus update o k response has a 5xx status code
func (o *PrometheusUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus update o k response a status code equal to that given
func (o *PrometheusUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateOK  %+v", 200, o.Payload)
}

func (o *PrometheusUpdateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateOK  %+v", 200, o.Payload)
}

func (o *PrometheusUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PrometheusUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusUpdateBadRequest creates a PrometheusUpdateBadRequest with default headers values
func NewPrometheusUpdateBadRequest() *PrometheusUpdateBadRequest {
	return &PrometheusUpdateBadRequest{}
}

/*
PrometheusUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this prometheus update bad request response has a 2xx status code
func (o *PrometheusUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus update bad request response has a 3xx status code
func (o *PrometheusUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update bad request response has a 4xx status code
func (o *PrometheusUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus update bad request response has a 5xx status code
func (o *PrometheusUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus update bad request response a status code equal to that given
func (o *PrometheusUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusUpdateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusUpdateUnauthorized creates a PrometheusUpdateUnauthorized with default headers values
func NewPrometheusUpdateUnauthorized() *PrometheusUpdateUnauthorized {
	return &PrometheusUpdateUnauthorized{}
}

/*
PrometheusUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusUpdateUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus update unauthorized response has a 2xx status code
func (o *PrometheusUpdateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus update unauthorized response has a 3xx status code
func (o *PrometheusUpdateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update unauthorized response has a 4xx status code
func (o *PrometheusUpdateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus update unauthorized response has a 5xx status code
func (o *PrometheusUpdateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus update unauthorized response a status code equal to that given
func (o *PrometheusUpdateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusUpdateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusUpdateUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusUpdateForbidden creates a PrometheusUpdateForbidden with default headers values
func NewPrometheusUpdateForbidden() *PrometheusUpdateForbidden {
	return &PrometheusUpdateForbidden{}
}

/*
PrometheusUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusUpdateForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus update forbidden response has a 2xx status code
func (o *PrometheusUpdateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus update forbidden response has a 3xx status code
func (o *PrometheusUpdateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update forbidden response has a 4xx status code
func (o *PrometheusUpdateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus update forbidden response has a 5xx status code
func (o *PrometheusUpdateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus update forbidden response a status code equal to that given
func (o *PrometheusUpdateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusUpdateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusUpdateForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusUpdateNotFound creates a PrometheusUpdateNotFound with default headers values
func NewPrometheusUpdateNotFound() *PrometheusUpdateNotFound {
	return &PrometheusUpdateNotFound{}
}

/*
PrometheusUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusUpdateNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus update not found response has a 2xx status code
func (o *PrometheusUpdateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus update not found response has a 3xx status code
func (o *PrometheusUpdateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update not found response has a 4xx status code
func (o *PrometheusUpdateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus update not found response has a 5xx status code
func (o *PrometheusUpdateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus update not found response a status code equal to that given
func (o *PrometheusUpdateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusUpdateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusUpdateNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusUpdateInternalServerError creates a PrometheusUpdateInternalServerError with default headers values
func NewPrometheusUpdateInternalServerError() *PrometheusUpdateInternalServerError {
	return &PrometheusUpdateInternalServerError{}
}

/*
PrometheusUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusUpdateInternalServerError struct {
}

// IsSuccess returns true when this prometheus update internal server error response has a 2xx status code
func (o *PrometheusUpdateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus update internal server error response has a 3xx status code
func (o *PrometheusUpdateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus update internal server error response has a 4xx status code
func (o *PrometheusUpdateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus update internal server error response has a 5xx status code
func (o *PrometheusUpdateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus update internal server error response a status code equal to that given
func (o *PrometheusUpdateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateInternalServerError ", 500)
}

func (o *PrometheusUpdateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/update/{id}][%d] prometheusUpdateInternalServerError ", 500)
}

func (o *PrometheusUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
