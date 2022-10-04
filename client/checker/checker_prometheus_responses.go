// Code generated by go-swagger; DO NOT EDIT.

package checker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CheckerPrometheusReader is a Reader for the CheckerPrometheus structure.
type CheckerPrometheusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerPrometheusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerPrometheusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerPrometheusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerPrometheusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerPrometheusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerPrometheusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerPrometheusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerPrometheusOK creates a CheckerPrometheusOK with default headers values
func NewCheckerPrometheusOK() *CheckerPrometheusOK {
	return &CheckerPrometheusOK{}
}

/*
CheckerPrometheusOK describes a response with status code 200, with default header values.

Success
*/
type CheckerPrometheusOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker prometheus o k response has a 2xx status code
func (o *CheckerPrometheusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker prometheus o k response has a 3xx status code
func (o *CheckerPrometheusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus o k response has a 4xx status code
func (o *CheckerPrometheusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker prometheus o k response has a 5xx status code
func (o *CheckerPrometheusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker prometheus o k response a status code equal to that given
func (o *CheckerPrometheusOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerPrometheusOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusOK  %+v", 200, o.Payload)
}

func (o *CheckerPrometheusOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusOK  %+v", 200, o.Payload)
}

func (o *CheckerPrometheusOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerPrometheusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerPrometheusBadRequest creates a CheckerPrometheusBadRequest with default headers values
func NewCheckerPrometheusBadRequest() *CheckerPrometheusBadRequest {
	return &CheckerPrometheusBadRequest{}
}

/*
CheckerPrometheusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerPrometheusBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this checker prometheus bad request response has a 2xx status code
func (o *CheckerPrometheusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker prometheus bad request response has a 3xx status code
func (o *CheckerPrometheusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus bad request response has a 4xx status code
func (o *CheckerPrometheusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker prometheus bad request response has a 5xx status code
func (o *CheckerPrometheusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker prometheus bad request response a status code equal to that given
func (o *CheckerPrometheusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerPrometheusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerPrometheusBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerPrometheusBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerPrometheusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerPrometheusUnauthorized creates a CheckerPrometheusUnauthorized with default headers values
func NewCheckerPrometheusUnauthorized() *CheckerPrometheusUnauthorized {
	return &CheckerPrometheusUnauthorized{}
}

/*
CheckerPrometheusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerPrometheusUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this checker prometheus unauthorized response has a 2xx status code
func (o *CheckerPrometheusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker prometheus unauthorized response has a 3xx status code
func (o *CheckerPrometheusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus unauthorized response has a 4xx status code
func (o *CheckerPrometheusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker prometheus unauthorized response has a 5xx status code
func (o *CheckerPrometheusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker prometheus unauthorized response a status code equal to that given
func (o *CheckerPrometheusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerPrometheusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerPrometheusUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerPrometheusUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerPrometheusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerPrometheusForbidden creates a CheckerPrometheusForbidden with default headers values
func NewCheckerPrometheusForbidden() *CheckerPrometheusForbidden {
	return &CheckerPrometheusForbidden{}
}

/*
CheckerPrometheusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerPrometheusForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this checker prometheus forbidden response has a 2xx status code
func (o *CheckerPrometheusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker prometheus forbidden response has a 3xx status code
func (o *CheckerPrometheusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus forbidden response has a 4xx status code
func (o *CheckerPrometheusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker prometheus forbidden response has a 5xx status code
func (o *CheckerPrometheusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker prometheus forbidden response a status code equal to that given
func (o *CheckerPrometheusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerPrometheusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusForbidden  %+v", 403, o.Payload)
}

func (o *CheckerPrometheusForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusForbidden  %+v", 403, o.Payload)
}

func (o *CheckerPrometheusForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerPrometheusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerPrometheusNotFound creates a CheckerPrometheusNotFound with default headers values
func NewCheckerPrometheusNotFound() *CheckerPrometheusNotFound {
	return &CheckerPrometheusNotFound{}
}

/*
CheckerPrometheusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerPrometheusNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this checker prometheus not found response has a 2xx status code
func (o *CheckerPrometheusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker prometheus not found response has a 3xx status code
func (o *CheckerPrometheusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus not found response has a 4xx status code
func (o *CheckerPrometheusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker prometheus not found response has a 5xx status code
func (o *CheckerPrometheusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker prometheus not found response a status code equal to that given
func (o *CheckerPrometheusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerPrometheusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusNotFound  %+v", 404, o.Payload)
}

func (o *CheckerPrometheusNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusNotFound  %+v", 404, o.Payload)
}

func (o *CheckerPrometheusNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CheckerPrometheusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerPrometheusInternalServerError creates a CheckerPrometheusInternalServerError with default headers values
func NewCheckerPrometheusInternalServerError() *CheckerPrometheusInternalServerError {
	return &CheckerPrometheusInternalServerError{}
}

/*
CheckerPrometheusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerPrometheusInternalServerError struct {
}

// IsSuccess returns true when this checker prometheus internal server error response has a 2xx status code
func (o *CheckerPrometheusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker prometheus internal server error response has a 3xx status code
func (o *CheckerPrometheusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker prometheus internal server error response has a 4xx status code
func (o *CheckerPrometheusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker prometheus internal server error response has a 5xx status code
func (o *CheckerPrometheusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker prometheus internal server error response a status code equal to that given
func (o *CheckerPrometheusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerPrometheusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusInternalServerError ", 500)
}

func (o *CheckerPrometheusInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/prometheus][%d] checkerPrometheusInternalServerError ", 500)
}

func (o *CheckerPrometheusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
