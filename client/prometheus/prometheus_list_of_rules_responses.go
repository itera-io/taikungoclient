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

/*
PrometheusListOfRulesOK describes a response with status code 200, with default header values.

Success
*/
type PrometheusListOfRulesOK struct {
	Payload *models.PrometheusRulesList
}

// IsSuccess returns true when this prometheus list of rules o k response has a 2xx status code
func (o *PrometheusListOfRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prometheus list of rules o k response has a 3xx status code
func (o *PrometheusListOfRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules o k response has a 4xx status code
func (o *PrometheusListOfRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus list of rules o k response has a 5xx status code
func (o *PrometheusListOfRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus list of rules o k response a status code equal to that given
func (o *PrometheusListOfRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PrometheusListOfRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesOK  %+v", 200, o.Payload)
}

func (o *PrometheusListOfRulesOK) String() string {
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

/*
PrometheusListOfRulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PrometheusListOfRulesBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus list of rules bad request response has a 2xx status code
func (o *PrometheusListOfRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus list of rules bad request response has a 3xx status code
func (o *PrometheusListOfRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules bad request response has a 4xx status code
func (o *PrometheusListOfRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus list of rules bad request response has a 5xx status code
func (o *PrometheusListOfRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus list of rules bad request response a status code equal to that given
func (o *PrometheusListOfRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PrometheusListOfRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusListOfRulesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PrometheusListOfRulesBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusListOfRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesUnauthorized creates a PrometheusListOfRulesUnauthorized with default headers values
func NewPrometheusListOfRulesUnauthorized() *PrometheusListOfRulesUnauthorized {
	return &PrometheusListOfRulesUnauthorized{}
}

/*
PrometheusListOfRulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PrometheusListOfRulesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus list of rules unauthorized response has a 2xx status code
func (o *PrometheusListOfRulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus list of rules unauthorized response has a 3xx status code
func (o *PrometheusListOfRulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules unauthorized response has a 4xx status code
func (o *PrometheusListOfRulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus list of rules unauthorized response has a 5xx status code
func (o *PrometheusListOfRulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus list of rules unauthorized response a status code equal to that given
func (o *PrometheusListOfRulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PrometheusListOfRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusListOfRulesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PrometheusListOfRulesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusListOfRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesForbidden creates a PrometheusListOfRulesForbidden with default headers values
func NewPrometheusListOfRulesForbidden() *PrometheusListOfRulesForbidden {
	return &PrometheusListOfRulesForbidden{}
}

/*
PrometheusListOfRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PrometheusListOfRulesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus list of rules forbidden response has a 2xx status code
func (o *PrometheusListOfRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus list of rules forbidden response has a 3xx status code
func (o *PrometheusListOfRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules forbidden response has a 4xx status code
func (o *PrometheusListOfRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus list of rules forbidden response has a 5xx status code
func (o *PrometheusListOfRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus list of rules forbidden response a status code equal to that given
func (o *PrometheusListOfRulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PrometheusListOfRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusListOfRulesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesForbidden  %+v", 403, o.Payload)
}

func (o *PrometheusListOfRulesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusListOfRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesNotFound creates a PrometheusListOfRulesNotFound with default headers values
func NewPrometheusListOfRulesNotFound() *PrometheusListOfRulesNotFound {
	return &PrometheusListOfRulesNotFound{}
}

/*
PrometheusListOfRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PrometheusListOfRulesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this prometheus list of rules not found response has a 2xx status code
func (o *PrometheusListOfRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus list of rules not found response has a 3xx status code
func (o *PrometheusListOfRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules not found response has a 4xx status code
func (o *PrometheusListOfRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this prometheus list of rules not found response has a 5xx status code
func (o *PrometheusListOfRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this prometheus list of rules not found response a status code equal to that given
func (o *PrometheusListOfRulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PrometheusListOfRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusListOfRulesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesNotFound  %+v", 404, o.Payload)
}

func (o *PrometheusListOfRulesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *PrometheusListOfRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrometheusListOfRulesInternalServerError creates a PrometheusListOfRulesInternalServerError with default headers values
func NewPrometheusListOfRulesInternalServerError() *PrometheusListOfRulesInternalServerError {
	return &PrometheusListOfRulesInternalServerError{}
}

/*
PrometheusListOfRulesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PrometheusListOfRulesInternalServerError struct {
}

// IsSuccess returns true when this prometheus list of rules internal server error response has a 2xx status code
func (o *PrometheusListOfRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this prometheus list of rules internal server error response has a 3xx status code
func (o *PrometheusListOfRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prometheus list of rules internal server error response has a 4xx status code
func (o *PrometheusListOfRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this prometheus list of rules internal server error response has a 5xx status code
func (o *PrometheusListOfRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this prometheus list of rules internal server error response a status code equal to that given
func (o *PrometheusListOfRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PrometheusListOfRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesInternalServerError ", 500)
}

func (o *PrometheusListOfRulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Prometheus][%d] prometheusListOfRulesInternalServerError ", 500)
}

func (o *PrometheusListOfRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
