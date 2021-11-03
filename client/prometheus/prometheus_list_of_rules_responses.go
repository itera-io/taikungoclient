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

// PrometheusListOfRulesReader is a Reader for the PrometheusListOfRules structure.
type PrometheusListOfRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrometheusListOfRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrometheusListOfRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPrometheusListOfRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPrometheusListOfRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrometheusListOfRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrometheusListOfRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPrometheusListOfRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPrometheusListOfRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrometheusListOfRulesOK creates a PrometheusListOfRulesOK with default headers values
func NewPrometheusListOfRulesOK() *PrometheusListOfRulesOK {
	return &PrometheusListOfRulesOK{}
}

/* PrometheusListOfRulesOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusListOfRulesOK struct {
	Payload *models.PrometheusRulesList
}

func (o *PrometheusListOfRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesOK  %+v", 200, o.Payload)
}
func (o *PrometheusListOfRulesOK) GetPayload() *models.PrometheusRulesList {
	return o.Payload
}

func (o *PrometheusListOfRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrometheusRulesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesBadRequest creates a PrometheusListOfRulesBadRequest with default headers values
func NewPrometheusListOfRulesBadRequest() *PrometheusListOfRulesBadRequest {
	return &PrometheusListOfRulesBadRequest{}
}

/* PrometheusListOfRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusListOfRulesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *PrometheusListOfRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesBadRequest  %+v", 400, o.Payload)
}
func (o *PrometheusListOfRulesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *PrometheusListOfRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesUnauthorized creates a PrometheusListOfRulesUnauthorized with default headers values
func NewPrometheusListOfRulesUnauthorized() *PrometheusListOfRulesUnauthorized {
	return &PrometheusListOfRulesUnauthorized{}
}

/* PrometheusListOfRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusListOfRulesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusListOfRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesUnauthorized  %+v", 401, o.Payload)
}
func (o *PrometheusListOfRulesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusListOfRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesForbidden creates a PrometheusListOfRulesForbidden with default headers values
func NewPrometheusListOfRulesForbidden() *PrometheusListOfRulesForbidden {
	return &PrometheusListOfRulesForbidden{}
}

/* PrometheusListOfRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusListOfRulesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusListOfRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesForbidden  %+v", 403, o.Payload)
}
func (o *PrometheusListOfRulesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusListOfRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesNotFound creates a PrometheusListOfRulesNotFound with default headers values
func NewPrometheusListOfRulesNotFound() *PrometheusListOfRulesNotFound {
	return &PrometheusListOfRulesNotFound{}
}

/* PrometheusListOfRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusListOfRulesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusListOfRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesNotFound  %+v", 404, o.Payload)
}
func (o *PrometheusListOfRulesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusListOfRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesTooManyRequests creates a PrometheusListOfRulesTooManyRequests with default headers values
func NewPrometheusListOfRulesTooManyRequests() *PrometheusListOfRulesTooManyRequests {
	return &PrometheusListOfRulesTooManyRequests{}
}

/* PrometheusListOfRulesTooManyRequests describes a response with status code 429, with default header values.

Client Error
*/
type PrometheusListOfRulesTooManyRequests struct {
	Payload *models.ProblemDetails
}

func (o *PrometheusListOfRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesTooManyRequests  %+v", 429, o.Payload)
}
func (o *PrometheusListOfRulesTooManyRequests) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *PrometheusListOfRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesInternalServerError creates a PrometheusListOfRulesInternalServerError with default headers values
func NewPrometheusListOfRulesInternalServerError() *PrometheusListOfRulesInternalServerError {
	return &PrometheusListOfRulesInternalServerError{}
}

/* PrometheusListOfRulesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusListOfRulesInternalServerError struct {
}

func (o *PrometheusListOfRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesInternalServerError ", 500)
}

func (o *PrometheusListOfRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
