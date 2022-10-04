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

/*
PrometheusDeleteOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this prometheus delete o k response has a 2xx status code
func (o *PrometheusDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus delete o k response has a 3xx status code
func (o *PrometheusDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete o k response has a 4xx status code
func (o *PrometheusDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus delete o k response has a 5xx status code
func (o *PrometheusDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus delete o k response a status code equal to that given
func (o *PrometheusDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteOK  %+v", 200, o.Payload)
}

func (o *PrometheusDeleteOK) String() string {
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

/*
PrometheusDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusDeleteBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus delete bad request response has a 2xx status code
func (o *PrometheusDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus delete bad request response has a 3xx status code
func (o *PrometheusDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete bad request response has a 4xx status code
func (o *PrometheusDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus delete bad request response has a 5xx status code
func (o *PrometheusDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus delete bad request response a status code equal to that given
func (o *PrometheusDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusDeleteBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteUnauthorized creates a PrometheusDeleteUnauthorized with default headers values
func NewPrometheusDeleteUnauthorized() *PrometheusDeleteUnauthorized {
	return &PrometheusDeleteUnauthorized{}
}

/*
PrometheusDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusDeleteUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus delete unauthorized response has a 2xx status code
func (o *PrometheusDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus delete unauthorized response has a 3xx status code
func (o *PrometheusDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete unauthorized response has a 4xx status code
func (o *PrometheusDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus delete unauthorized response has a 5xx status code
func (o *PrometheusDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus delete unauthorized response a status code equal to that given
func (o *PrometheusDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusDeleteUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteForbidden creates a PrometheusDeleteForbidden with default headers values
func NewPrometheusDeleteForbidden() *PrometheusDeleteForbidden {
	return &PrometheusDeleteForbidden{}
}

/*
PrometheusDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusDeleteForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus delete forbidden response has a 2xx status code
func (o *PrometheusDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus delete forbidden response has a 3xx status code
func (o *PrometheusDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete forbidden response has a 4xx status code
func (o *PrometheusDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus delete forbidden response has a 5xx status code
func (o *PrometheusDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus delete forbidden response a status code equal to that given
func (o *PrometheusDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusDeleteForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteNotFound creates a PrometheusDeleteNotFound with default headers values
func NewPrometheusDeleteNotFound() *PrometheusDeleteNotFound {
	return &PrometheusDeleteNotFound{}
}

/*
PrometheusDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusDeleteNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus delete not found response has a 2xx status code
func (o *PrometheusDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus delete not found response has a 3xx status code
func (o *PrometheusDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete not found response has a 4xx status code
func (o *PrometheusDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus delete not found response has a 5xx status code
func (o *PrometheusDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus delete not found response a status code equal to that given
func (o *PrometheusDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusDeleteNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDeleteInternalServerError creates a PrometheusDeleteInternalServerError with default headers values
func NewPrometheusDeleteInternalServerError() *PrometheusDeleteInternalServerError {
	return &PrometheusDeleteInternalServerError{}
}

/*
PrometheusDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusDeleteInternalServerError struct {
}

// IsSuccess returns true when this prometheus delete internal server error response has a 2xx status code
func (o *PrometheusDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus delete internal server error response has a 3xx status code
func (o *PrometheusDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus delete internal server error response has a 4xx status code
func (o *PrometheusDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus delete internal server error response has a 5xx status code
func (o *PrometheusDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus delete internal server error response a status code equal to that given
func (o *PrometheusDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteInternalServerError ", 500)
}

func (o *PrometheusDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v{v}/Prometheus/{id}][%d] prometheusDeleteInternalServerError ", 500)
}

func (o *PrometheusDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
