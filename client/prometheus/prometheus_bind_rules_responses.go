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

// PrometheusBindRulesReader is a Reader for the PrometheusBindRules structure.
type PrometheusBindRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusBindRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusBindRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusBindRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusBindRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusBindRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusBindRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusBindRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusBindRulesOK creates a PrometheusBindRulesOK with default headers values
func NewPrometheusBindRulesOK() *PrometheusBindRulesOK {
	return &PrometheusBindRulesOK{}
}

/*
PrometheusBindRulesOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusBindRulesOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this prometheus bind rules o k response has a 2xx status code
func (o *PrometheusBindRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus bind rules o k response has a 3xx status code
func (o *PrometheusBindRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules o k response has a 4xx status code
func (o *PrometheusBindRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus bind rules o k response has a 5xx status code
func (o *PrometheusBindRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus bind rules o k response a status code equal to that given
func (o *PrometheusBindRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusBindRulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesOK  %+v", 200, o.Payload)
}

func (o *PrometheusBindRulesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesOK  %+v", 200, o.Payload)
}

func (o *PrometheusBindRulesOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *PrometheusBindRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusBindRulesBadRequest creates a PrometheusBindRulesBadRequest with default headers values
func NewPrometheusBindRulesBadRequest() *PrometheusBindRulesBadRequest {
	return &PrometheusBindRulesBadRequest{}
}

/*
PrometheusBindRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusBindRulesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this prometheus bind rules bad request response has a 2xx status code
func (o *PrometheusBindRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus bind rules bad request response has a 3xx status code
func (o *PrometheusBindRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules bad request response has a 4xx status code
func (o *PrometheusBindRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus bind rules bad request response has a 5xx status code
func (o *PrometheusBindRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus bind rules bad request response a status code equal to that given
func (o *PrometheusBindRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusBindRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusBindRulesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusBindRulesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusBindRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusBindRulesUnauthorized creates a PrometheusBindRulesUnauthorized with default headers values
func NewPrometheusBindRulesUnauthorized() *PrometheusBindRulesUnauthorized {
	return &PrometheusBindRulesUnauthorized{}
}

/*
PrometheusBindRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusBindRulesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus bind rules unauthorized response has a 2xx status code
func (o *PrometheusBindRulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus bind rules unauthorized response has a 3xx status code
func (o *PrometheusBindRulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules unauthorized response has a 4xx status code
func (o *PrometheusBindRulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus bind rules unauthorized response has a 5xx status code
func (o *PrometheusBindRulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus bind rules unauthorized response a status code equal to that given
func (o *PrometheusBindRulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusBindRulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusBindRulesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusBindRulesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusBindRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusBindRulesForbidden creates a PrometheusBindRulesForbidden with default headers values
func NewPrometheusBindRulesForbidden() *PrometheusBindRulesForbidden {
	return &PrometheusBindRulesForbidden{}
}

/*
PrometheusBindRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusBindRulesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus bind rules forbidden response has a 2xx status code
func (o *PrometheusBindRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus bind rules forbidden response has a 3xx status code
func (o *PrometheusBindRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules forbidden response has a 4xx status code
func (o *PrometheusBindRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus bind rules forbidden response has a 5xx status code
func (o *PrometheusBindRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus bind rules forbidden response a status code equal to that given
func (o *PrometheusBindRulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusBindRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusBindRulesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusBindRulesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusBindRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusBindRulesNotFound creates a PrometheusBindRulesNotFound with default headers values
func NewPrometheusBindRulesNotFound() *PrometheusBindRulesNotFound {
	return &PrometheusBindRulesNotFound{}
}

/*
PrometheusBindRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusBindRulesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this prometheus bind rules not found response has a 2xx status code
func (o *PrometheusBindRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus bind rules not found response has a 3xx status code
func (o *PrometheusBindRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules not found response has a 4xx status code
func (o *PrometheusBindRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus bind rules not found response has a 5xx status code
func (o *PrometheusBindRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus bind rules not found response a status code equal to that given
func (o *PrometheusBindRulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusBindRulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusBindRulesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusBindRulesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusBindRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusBindRulesInternalServerError creates a PrometheusBindRulesInternalServerError with default headers values
func NewPrometheusBindRulesInternalServerError() *PrometheusBindRulesInternalServerError {
	return &PrometheusBindRulesInternalServerError{}
}

/*
PrometheusBindRulesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusBindRulesInternalServerError struct {
}

// IsSuccess returns true when this prometheus bind rules internal server error response has a 2xx status code
func (o *PrometheusBindRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus bind rules internal server error response has a 3xx status code
func (o *PrometheusBindRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus bind rules internal server error response has a 4xx status code
func (o *PrometheusBindRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus bind rules internal server error response has a 5xx status code
func (o *PrometheusBindRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus bind rules internal server error response a status code equal to that given
func (o *PrometheusBindRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusBindRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesInternalServerError ", 500)
}

func (o *PrometheusBindRulesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Prometheus/bindrules][%d] prometheusBindRulesInternalServerError ", 500)
}

func (o *PrometheusBindRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
