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

// CronJobSyncAppProxyReader is a Reader for the CronJobSyncAppProxy structure.
type CronJobSyncAppProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobSyncAppProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobSyncAppProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobSyncAppProxyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobSyncAppProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobSyncAppProxyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobSyncAppProxyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobSyncAppProxyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobSyncAppProxyOK creates a CronJobSyncAppProxyOK with default headers values
func NewCronJobSyncAppProxyOK() *CronJobSyncAppProxyOK {
	return &CronJobSyncAppProxyOK{}
}

/*
CronJobSyncAppProxyOK describes a response with status code 200, with default header values.

Success
*/
type CronJobSyncAppProxyOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job sync app proxy o k response has a 2xx status code
func (o *CronJobSyncAppProxyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job sync app proxy o k response has a 3xx status code
func (o *CronJobSyncAppProxyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy o k response has a 4xx status code
func (o *CronJobSyncAppProxyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync app proxy o k response has a 5xx status code
func (o *CronJobSyncAppProxyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync app proxy o k response a status code equal to that given
func (o *CronJobSyncAppProxyOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobSyncAppProxyOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncAppProxyOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncAppProxyOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobSyncAppProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncAppProxyBadRequest creates a CronJobSyncAppProxyBadRequest with default headers values
func NewCronJobSyncAppProxyBadRequest() *CronJobSyncAppProxyBadRequest {
	return &CronJobSyncAppProxyBadRequest{}
}

/*
CronJobSyncAppProxyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobSyncAppProxyBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync app proxy bad request response has a 2xx status code
func (o *CronJobSyncAppProxyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync app proxy bad request response has a 3xx status code
func (o *CronJobSyncAppProxyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy bad request response has a 4xx status code
func (o *CronJobSyncAppProxyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync app proxy bad request response has a 5xx status code
func (o *CronJobSyncAppProxyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync app proxy bad request response a status code equal to that given
func (o *CronJobSyncAppProxyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobSyncAppProxyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncAppProxyBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncAppProxyBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncAppProxyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncAppProxyUnauthorized creates a CronJobSyncAppProxyUnauthorized with default headers values
func NewCronJobSyncAppProxyUnauthorized() *CronJobSyncAppProxyUnauthorized {
	return &CronJobSyncAppProxyUnauthorized{}
}

/*
CronJobSyncAppProxyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobSyncAppProxyUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync app proxy unauthorized response has a 2xx status code
func (o *CronJobSyncAppProxyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync app proxy unauthorized response has a 3xx status code
func (o *CronJobSyncAppProxyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy unauthorized response has a 4xx status code
func (o *CronJobSyncAppProxyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync app proxy unauthorized response has a 5xx status code
func (o *CronJobSyncAppProxyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync app proxy unauthorized response a status code equal to that given
func (o *CronJobSyncAppProxyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobSyncAppProxyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncAppProxyUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncAppProxyUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncAppProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncAppProxyForbidden creates a CronJobSyncAppProxyForbidden with default headers values
func NewCronJobSyncAppProxyForbidden() *CronJobSyncAppProxyForbidden {
	return &CronJobSyncAppProxyForbidden{}
}

/*
CronJobSyncAppProxyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobSyncAppProxyForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync app proxy forbidden response has a 2xx status code
func (o *CronJobSyncAppProxyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync app proxy forbidden response has a 3xx status code
func (o *CronJobSyncAppProxyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy forbidden response has a 4xx status code
func (o *CronJobSyncAppProxyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync app proxy forbidden response has a 5xx status code
func (o *CronJobSyncAppProxyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync app proxy forbidden response a status code equal to that given
func (o *CronJobSyncAppProxyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobSyncAppProxyForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncAppProxyForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncAppProxyForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncAppProxyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncAppProxyNotFound creates a CronJobSyncAppProxyNotFound with default headers values
func NewCronJobSyncAppProxyNotFound() *CronJobSyncAppProxyNotFound {
	return &CronJobSyncAppProxyNotFound{}
}

/*
CronJobSyncAppProxyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobSyncAppProxyNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync app proxy not found response has a 2xx status code
func (o *CronJobSyncAppProxyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync app proxy not found response has a 3xx status code
func (o *CronJobSyncAppProxyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy not found response has a 4xx status code
func (o *CronJobSyncAppProxyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync app proxy not found response has a 5xx status code
func (o *CronJobSyncAppProxyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync app proxy not found response a status code equal to that given
func (o *CronJobSyncAppProxyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobSyncAppProxyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncAppProxyNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncAppProxyNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncAppProxyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncAppProxyInternalServerError creates a CronJobSyncAppProxyInternalServerError with default headers values
func NewCronJobSyncAppProxyInternalServerError() *CronJobSyncAppProxyInternalServerError {
	return &CronJobSyncAppProxyInternalServerError{}
}

/*
CronJobSyncAppProxyInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobSyncAppProxyInternalServerError struct {
}

// IsSuccess returns true when this cron job sync app proxy internal server error response has a 2xx status code
func (o *CronJobSyncAppProxyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync app proxy internal server error response has a 3xx status code
func (o *CronJobSyncAppProxyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync app proxy internal server error response has a 4xx status code
func (o *CronJobSyncAppProxyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync app proxy internal server error response has a 5xx status code
func (o *CronJobSyncAppProxyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job sync app proxy internal server error response a status code equal to that given
func (o *CronJobSyncAppProxyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobSyncAppProxyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyInternalServerError ", 500)
}

func (o *CronJobSyncAppProxyInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-app-proxy][%d] cronJobSyncAppProxyInternalServerError ", 500)
}

func (o *CronJobSyncAppProxyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}