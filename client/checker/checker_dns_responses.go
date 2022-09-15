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

// CheckerDNSReader is a Reader for the CheckerDNS structure.
type CheckerDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckerDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckerDNSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckerDNSBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCheckerDNSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckerDNSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckerDNSNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCheckerDNSInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckerDNSOK creates a CheckerDNSOK with default headers values
func NewCheckerDNSOK() *CheckerDNSOK {
	return &CheckerDNSOK{}
}

/*
CheckerDNSOK describes a response with status code 200, with default header values.

Success
*/
type CheckerDNSOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this checker Dns o k response has a 2xx status code
func (o *CheckerDNSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checker Dns o k response has a 3xx status code
func (o *CheckerDNSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns o k response has a 4xx status code
func (o *CheckerDNSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker Dns o k response has a 5xx status code
func (o *CheckerDNSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checker Dns o k response a status code equal to that given
func (o *CheckerDNSOK) IsCode(code int) bool {
	return code == 200
}

func (o *CheckerDNSOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsOK  %+v", 200, o.Payload)
}

func (o *CheckerDNSOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsOK  %+v", 200, o.Payload)
}

func (o *CheckerDNSOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CheckerDNSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDNSBadRequest creates a CheckerDNSBadRequest with default headers values
func NewCheckerDNSBadRequest() *CheckerDNSBadRequest {
	return &CheckerDNSBadRequest{}
}

/*
CheckerDNSBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CheckerDNSBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this checker Dns bad request response has a 2xx status code
func (o *CheckerDNSBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker Dns bad request response has a 3xx status code
func (o *CheckerDNSBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns bad request response has a 4xx status code
func (o *CheckerDNSBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker Dns bad request response has a 5xx status code
func (o *CheckerDNSBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this checker Dns bad request response a status code equal to that given
func (o *CheckerDNSBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CheckerDNSBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerDNSBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsBadRequest  %+v", 400, o.Payload)
}

func (o *CheckerDNSBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CheckerDNSBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDNSUnauthorized creates a CheckerDNSUnauthorized with default headers values
func NewCheckerDNSUnauthorized() *CheckerDNSUnauthorized {
	return &CheckerDNSUnauthorized{}
}

/*
CheckerDNSUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckerDNSUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker Dns unauthorized response has a 2xx status code
func (o *CheckerDNSUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker Dns unauthorized response has a 3xx status code
func (o *CheckerDNSUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns unauthorized response has a 4xx status code
func (o *CheckerDNSUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker Dns unauthorized response has a 5xx status code
func (o *CheckerDNSUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this checker Dns unauthorized response a status code equal to that given
func (o *CheckerDNSUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CheckerDNSUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerDNSUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsUnauthorized  %+v", 401, o.Payload)
}

func (o *CheckerDNSUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerDNSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDNSForbidden creates a CheckerDNSForbidden with default headers values
func NewCheckerDNSForbidden() *CheckerDNSForbidden {
	return &CheckerDNSForbidden{}
}

/*
CheckerDNSForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckerDNSForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker Dns forbidden response has a 2xx status code
func (o *CheckerDNSForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker Dns forbidden response has a 3xx status code
func (o *CheckerDNSForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns forbidden response has a 4xx status code
func (o *CheckerDNSForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker Dns forbidden response has a 5xx status code
func (o *CheckerDNSForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this checker Dns forbidden response a status code equal to that given
func (o *CheckerDNSForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CheckerDNSForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsForbidden  %+v", 403, o.Payload)
}

func (o *CheckerDNSForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsForbidden  %+v", 403, o.Payload)
}

func (o *CheckerDNSForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerDNSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDNSNotFound creates a CheckerDNSNotFound with default headers values
func NewCheckerDNSNotFound() *CheckerDNSNotFound {
	return &CheckerDNSNotFound{}
}

/*
CheckerDNSNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckerDNSNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this checker Dns not found response has a 2xx status code
func (o *CheckerDNSNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker Dns not found response has a 3xx status code
func (o *CheckerDNSNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns not found response has a 4xx status code
func (o *CheckerDNSNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this checker Dns not found response has a 5xx status code
func (o *CheckerDNSNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this checker Dns not found response a status code equal to that given
func (o *CheckerDNSNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CheckerDNSNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsNotFound  %+v", 404, o.Payload)
}

func (o *CheckerDNSNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsNotFound  %+v", 404, o.Payload)
}

func (o *CheckerDNSNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CheckerDNSNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckerDNSInternalServerError creates a CheckerDNSInternalServerError with default headers values
func NewCheckerDNSInternalServerError() *CheckerDNSInternalServerError {
	return &CheckerDNSInternalServerError{}
}

/*
CheckerDNSInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CheckerDNSInternalServerError struct {
}

// IsSuccess returns true when this checker Dns internal server error response has a 2xx status code
func (o *CheckerDNSInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this checker Dns internal server error response has a 3xx status code
func (o *CheckerDNSInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checker Dns internal server error response has a 4xx status code
func (o *CheckerDNSInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this checker Dns internal server error response has a 5xx status code
func (o *CheckerDNSInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this checker Dns internal server error response a status code equal to that given
func (o *CheckerDNSInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CheckerDNSInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsInternalServerError ", 500)
}

func (o *CheckerDNSInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Checker/dns][%d] checkerDnsInternalServerError ", 500)
}

func (o *CheckerDNSInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
