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

// CronJobUpdateProjectAppStatusReader is a Reader for the CronJobUpdateProjectAppStatus structure.
type CronJobUpdateProjectAppStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobUpdateProjectAppStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobUpdateProjectAppStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobUpdateProjectAppStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobUpdateProjectAppStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobUpdateProjectAppStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobUpdateProjectAppStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobUpdateProjectAppStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobUpdateProjectAppStatusOK creates a CronJobUpdateProjectAppStatusOK with default headers values
func NewCronJobUpdateProjectAppStatusOK() *CronJobUpdateProjectAppStatusOK {
	return &CronJobUpdateProjectAppStatusOK{}
}

/*
CronJobUpdateProjectAppStatusOK describes a response with status code 200, with default header values.

Success
*/
type CronJobUpdateProjectAppStatusOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job update project app status o k response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job update project app status o k response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status o k response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job update project app status o k response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project app status o k response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobUpdateProjectAppStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusOK  %+v", 200, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusOK  %+v", 200, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobUpdateProjectAppStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectAppStatusBadRequest creates a CronJobUpdateProjectAppStatusBadRequest with default headers values
func NewCronJobUpdateProjectAppStatusBadRequest() *CronJobUpdateProjectAppStatusBadRequest {
	return &CronJobUpdateProjectAppStatusBadRequest{}
}

/*
CronJobUpdateProjectAppStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobUpdateProjectAppStatusBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job update project app status bad request response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project app status bad request response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status bad request response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project app status bad request response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project app status bad request response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobUpdateProjectAppStatusBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobUpdateProjectAppStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectAppStatusUnauthorized creates a CronJobUpdateProjectAppStatusUnauthorized with default headers values
func NewCronJobUpdateProjectAppStatusUnauthorized() *CronJobUpdateProjectAppStatusUnauthorized {
	return &CronJobUpdateProjectAppStatusUnauthorized{}
}

/*
CronJobUpdateProjectAppStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobUpdateProjectAppStatusUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job update project app status unauthorized response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project app status unauthorized response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status unauthorized response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project app status unauthorized response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project app status unauthorized response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobUpdateProjectAppStatusUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobUpdateProjectAppStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectAppStatusForbidden creates a CronJobUpdateProjectAppStatusForbidden with default headers values
func NewCronJobUpdateProjectAppStatusForbidden() *CronJobUpdateProjectAppStatusForbidden {
	return &CronJobUpdateProjectAppStatusForbidden{}
}

/*
CronJobUpdateProjectAppStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobUpdateProjectAppStatusForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job update project app status forbidden response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project app status forbidden response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status forbidden response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project app status forbidden response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project app status forbidden response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobUpdateProjectAppStatusForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusForbidden  %+v", 403, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusForbidden  %+v", 403, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobUpdateProjectAppStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectAppStatusNotFound creates a CronJobUpdateProjectAppStatusNotFound with default headers values
func NewCronJobUpdateProjectAppStatusNotFound() *CronJobUpdateProjectAppStatusNotFound {
	return &CronJobUpdateProjectAppStatusNotFound{}
}

/*
CronJobUpdateProjectAppStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobUpdateProjectAppStatusNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job update project app status not found response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project app status not found response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status not found response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project app status not found response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project app status not found response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobUpdateProjectAppStatusNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusNotFound  %+v", 404, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusNotFound  %+v", 404, o.Payload)
}

func (o *CronJobUpdateProjectAppStatusNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobUpdateProjectAppStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectAppStatusInternalServerError creates a CronJobUpdateProjectAppStatusInternalServerError with default headers values
func NewCronJobUpdateProjectAppStatusInternalServerError() *CronJobUpdateProjectAppStatusInternalServerError {
	return &CronJobUpdateProjectAppStatusInternalServerError{}
}

/*
CronJobUpdateProjectAppStatusInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobUpdateProjectAppStatusInternalServerError struct {
}

// IsSuccess returns true when this cron job update project app status internal server error response has a 2xx status code
func (o *CronJobUpdateProjectAppStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project app status internal server error response has a 3xx status code
func (o *CronJobUpdateProjectAppStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project app status internal server error response has a 4xx status code
func (o *CronJobUpdateProjectAppStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job update project app status internal server error response has a 5xx status code
func (o *CronJobUpdateProjectAppStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job update project app status internal server error response a status code equal to that given
func (o *CronJobUpdateProjectAppStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobUpdateProjectAppStatusInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusInternalServerError ", 500)
}

func (o *CronJobUpdateProjectAppStatusInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-app-status][%d] cronJobUpdateProjectAppStatusInternalServerError ", 500)
}

func (o *CronJobUpdateProjectAppStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
