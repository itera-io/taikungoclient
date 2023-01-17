// Code generated by go-swagger; DO NOT EDIT.

package cron_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// CronJobDeleteExpiredServersReader is a Reader for the CronJobDeleteExpiredServers structure.
type CronJobDeleteExpiredServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteExpiredServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteExpiredServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteExpiredServersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteExpiredServersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteExpiredServersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteExpiredServersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteExpiredServersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteExpiredServersOK creates a CronJobDeleteExpiredServersOK with default headers values
func NewCronJobDeleteExpiredServersOK() *CronJobDeleteExpiredServersOK {
	return &CronJobDeleteExpiredServersOK{}
}

/*
CronJobDeleteExpiredServersOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredServersOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job delete expired servers o k response has a 2xx status code
func (o *CronJobDeleteExpiredServersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job delete expired servers o k response has a 3xx status code
func (o *CronJobDeleteExpiredServersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers o k response has a 4xx status code
func (o *CronJobDeleteExpiredServersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired servers o k response has a 5xx status code
func (o *CronJobDeleteExpiredServersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired servers o k response a status code equal to that given
func (o *CronJobDeleteExpiredServersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job delete expired servers o k response
func (o *CronJobDeleteExpiredServersOK) Code() int {
	return 200
}

func (o *CronJobDeleteExpiredServersOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredServersOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredServersOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteExpiredServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredServersBadRequest creates a CronJobDeleteExpiredServersBadRequest with default headers values
func NewCronJobDeleteExpiredServersBadRequest() *CronJobDeleteExpiredServersBadRequest {
	return &CronJobDeleteExpiredServersBadRequest{}
}

/*
CronJobDeleteExpiredServersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredServersBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired servers bad request response has a 2xx status code
func (o *CronJobDeleteExpiredServersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired servers bad request response has a 3xx status code
func (o *CronJobDeleteExpiredServersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers bad request response has a 4xx status code
func (o *CronJobDeleteExpiredServersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired servers bad request response has a 5xx status code
func (o *CronJobDeleteExpiredServersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired servers bad request response a status code equal to that given
func (o *CronJobDeleteExpiredServersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job delete expired servers bad request response
func (o *CronJobDeleteExpiredServersBadRequest) Code() int {
	return 400
}

func (o *CronJobDeleteExpiredServersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredServersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredServersBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredServersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredServersUnauthorized creates a CronJobDeleteExpiredServersUnauthorized with default headers values
func NewCronJobDeleteExpiredServersUnauthorized() *CronJobDeleteExpiredServersUnauthorized {
	return &CronJobDeleteExpiredServersUnauthorized{}
}

/*
CronJobDeleteExpiredServersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredServersUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired servers unauthorized response has a 2xx status code
func (o *CronJobDeleteExpiredServersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired servers unauthorized response has a 3xx status code
func (o *CronJobDeleteExpiredServersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers unauthorized response has a 4xx status code
func (o *CronJobDeleteExpiredServersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired servers unauthorized response has a 5xx status code
func (o *CronJobDeleteExpiredServersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired servers unauthorized response a status code equal to that given
func (o *CronJobDeleteExpiredServersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job delete expired servers unauthorized response
func (o *CronJobDeleteExpiredServersUnauthorized) Code() int {
	return 401
}

func (o *CronJobDeleteExpiredServersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredServersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredServersUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredServersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredServersForbidden creates a CronJobDeleteExpiredServersForbidden with default headers values
func NewCronJobDeleteExpiredServersForbidden() *CronJobDeleteExpiredServersForbidden {
	return &CronJobDeleteExpiredServersForbidden{}
}

/*
CronJobDeleteExpiredServersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredServersForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired servers forbidden response has a 2xx status code
func (o *CronJobDeleteExpiredServersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired servers forbidden response has a 3xx status code
func (o *CronJobDeleteExpiredServersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers forbidden response has a 4xx status code
func (o *CronJobDeleteExpiredServersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired servers forbidden response has a 5xx status code
func (o *CronJobDeleteExpiredServersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired servers forbidden response a status code equal to that given
func (o *CronJobDeleteExpiredServersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job delete expired servers forbidden response
func (o *CronJobDeleteExpiredServersForbidden) Code() int {
	return 403
}

func (o *CronJobDeleteExpiredServersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredServersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredServersForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredServersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredServersNotFound creates a CronJobDeleteExpiredServersNotFound with default headers values
func NewCronJobDeleteExpiredServersNotFound() *CronJobDeleteExpiredServersNotFound {
	return &CronJobDeleteExpiredServersNotFound{}
}

/*
CronJobDeleteExpiredServersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredServersNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired servers not found response has a 2xx status code
func (o *CronJobDeleteExpiredServersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired servers not found response has a 3xx status code
func (o *CronJobDeleteExpiredServersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers not found response has a 4xx status code
func (o *CronJobDeleteExpiredServersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired servers not found response has a 5xx status code
func (o *CronJobDeleteExpiredServersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired servers not found response a status code equal to that given
func (o *CronJobDeleteExpiredServersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job delete expired servers not found response
func (o *CronJobDeleteExpiredServersNotFound) Code() int {
	return 404
}

func (o *CronJobDeleteExpiredServersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredServersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredServersNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredServersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredServersInternalServerError creates a CronJobDeleteExpiredServersInternalServerError with default headers values
func NewCronJobDeleteExpiredServersInternalServerError() *CronJobDeleteExpiredServersInternalServerError {
	return &CronJobDeleteExpiredServersInternalServerError{}
}

/*
CronJobDeleteExpiredServersInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredServersInternalServerError struct {
}

// IsSuccess returns true when this cron job delete expired servers internal server error response has a 2xx status code
func (o *CronJobDeleteExpiredServersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired servers internal server error response has a 3xx status code
func (o *CronJobDeleteExpiredServersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired servers internal server error response has a 4xx status code
func (o *CronJobDeleteExpiredServersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired servers internal server error response has a 5xx status code
func (o *CronJobDeleteExpiredServersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job delete expired servers internal server error response a status code equal to that given
func (o *CronJobDeleteExpiredServersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job delete expired servers internal server error response
func (o *CronJobDeleteExpiredServersInternalServerError) Code() int {
	return 500
}

func (o *CronJobDeleteExpiredServersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredServersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/servers][%d] cronJobDeleteExpiredServersInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredServersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
