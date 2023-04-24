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

// PrometheusDetailsReader is a Reader for the PrometheusDetails structure.
type PrometheusDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusDetailsOK creates a PrometheusDetailsOK with default headers values
func NewPrometheusDetailsOK() *PrometheusDetailsOK {
	return &PrometheusDetailsOK{}
}

/*
PrometheusDetailsOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusDetailsOK struct {
	Payload []*models.CommonDropdownIsBoundDto
}

// IsSuccess returns true when this prometheus details o k response has a 2xx status code
func (o *PrometheusDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus details o k response has a 3xx status code
func (o *PrometheusDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details o k response has a 4xx status code
func (o *PrometheusDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus details o k response has a 5xx status code
func (o *PrometheusDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus details o k response a status code equal to that given
func (o *PrometheusDetailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the prometheus details o k response
func (o *PrometheusDetailsOK) Code() int {
	return 200
}

func (o *PrometheusDetailsOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsOK  %+v", 200, o.Payload)
}

func (o *PrometheusDetailsOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsOK  %+v", 200, o.Payload)
}

func (o *PrometheusDetailsOK) GetPayload() []*models.CommonDropdownIsBoundDto {
	return o.Payload
}

func (o *PrometheusDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDetailsBadRequest creates a PrometheusDetailsBadRequest with default headers values
func NewPrometheusDetailsBadRequest() *PrometheusDetailsBadRequest {
	return &PrometheusDetailsBadRequest{}
}

/*
PrometheusDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusDetailsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus details bad request response has a 2xx status code
func (o *PrometheusDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus details bad request response has a 3xx status code
func (o *PrometheusDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details bad request response has a 4xx status code
func (o *PrometheusDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus details bad request response has a 5xx status code
func (o *PrometheusDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus details bad request response a status code equal to that given
func (o *PrometheusDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the prometheus details bad request response
func (o *PrometheusDetailsBadRequest) Code() int {
	return 400
}

func (o *PrometheusDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusDetailsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusDetailsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDetailsUnauthorized creates a PrometheusDetailsUnauthorized with default headers values
func NewPrometheusDetailsUnauthorized() *PrometheusDetailsUnauthorized {
	return &PrometheusDetailsUnauthorized{}
}

/*
PrometheusDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusDetailsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus details unauthorized response has a 2xx status code
func (o *PrometheusDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus details unauthorized response has a 3xx status code
func (o *PrometheusDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details unauthorized response has a 4xx status code
func (o *PrometheusDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus details unauthorized response has a 5xx status code
func (o *PrometheusDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus details unauthorized response a status code equal to that given
func (o *PrometheusDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the prometheus details unauthorized response
func (o *PrometheusDetailsUnauthorized) Code() int {
	return 401
}

func (o *PrometheusDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusDetailsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusDetailsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDetailsForbidden creates a PrometheusDetailsForbidden with default headers values
func NewPrometheusDetailsForbidden() *PrometheusDetailsForbidden {
	return &PrometheusDetailsForbidden{}
}

/*
PrometheusDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusDetailsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus details forbidden response has a 2xx status code
func (o *PrometheusDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus details forbidden response has a 3xx status code
func (o *PrometheusDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details forbidden response has a 4xx status code
func (o *PrometheusDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus details forbidden response has a 5xx status code
func (o *PrometheusDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus details forbidden response a status code equal to that given
func (o *PrometheusDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the prometheus details forbidden response
func (o *PrometheusDetailsForbidden) Code() int {
	return 403
}

func (o *PrometheusDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusDetailsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusDetailsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDetailsNotFound creates a PrometheusDetailsNotFound with default headers values
func NewPrometheusDetailsNotFound() *PrometheusDetailsNotFound {
	return &PrometheusDetailsNotFound{}
}

/*
PrometheusDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusDetailsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this prometheus details not found response has a 2xx status code
func (o *PrometheusDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus details not found response has a 3xx status code
func (o *PrometheusDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details not found response has a 4xx status code
func (o *PrometheusDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus details not found response has a 5xx status code
func (o *PrometheusDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus details not found response a status code equal to that given
func (o *PrometheusDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the prometheus details not found response
func (o *PrometheusDetailsNotFound) Code() int {
	return 404
}

func (o *PrometheusDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusDetailsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusDetailsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PrometheusDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusDetailsInternalServerError creates a PrometheusDetailsInternalServerError with default headers values
func NewPrometheusDetailsInternalServerError() *PrometheusDetailsInternalServerError {
	return &PrometheusDetailsInternalServerError{}
}

/*
PrometheusDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusDetailsInternalServerError struct {
}

// IsSuccess returns true when this prometheus details internal server error response has a 2xx status code
func (o *PrometheusDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus details internal server error response has a 3xx status code
func (o *PrometheusDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus details internal server error response has a 4xx status code
func (o *PrometheusDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus details internal server error response has a 5xx status code
func (o *PrometheusDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus details internal server error response a status code equal to that given
func (o *PrometheusDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the prometheus details internal server error response
func (o *PrometheusDetailsInternalServerError) Code() int {
	return 500
}

func (o *PrometheusDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsInternalServerError ", 500)
}

func (o *PrometheusDetailsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus/details][%d] prometheusDetailsInternalServerError ", 500)
}

func (o *PrometheusDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
