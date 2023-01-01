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

// CronJobSyncOrganizationsReader is a Reader for the CronJobSyncOrganizations structure.
type CronJobSyncOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobSyncOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobSyncOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobSyncOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobSyncOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobSyncOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobSyncOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobSyncOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobSyncOrganizationsOK creates a CronJobSyncOrganizationsOK with default headers values
func NewCronJobSyncOrganizationsOK() *CronJobSyncOrganizationsOK {
	return &CronJobSyncOrganizationsOK{}
}

/*
CronJobSyncOrganizationsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobSyncOrganizationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job sync organizations o k response has a 2xx status code
func (o *CronJobSyncOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job sync organizations o k response has a 3xx status code
func (o *CronJobSyncOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations o k response has a 4xx status code
func (o *CronJobSyncOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync organizations o k response has a 5xx status code
func (o *CronJobSyncOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync organizations o k response a status code equal to that given
func (o *CronJobSyncOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobSyncOrganizationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncOrganizationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncOrganizationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobSyncOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncOrganizationsBadRequest creates a CronJobSyncOrganizationsBadRequest with default headers values
func NewCronJobSyncOrganizationsBadRequest() *CronJobSyncOrganizationsBadRequest {
	return &CronJobSyncOrganizationsBadRequest{}
}

/*
CronJobSyncOrganizationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobSyncOrganizationsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync organizations bad request response has a 2xx status code
func (o *CronJobSyncOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync organizations bad request response has a 3xx status code
func (o *CronJobSyncOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations bad request response has a 4xx status code
func (o *CronJobSyncOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync organizations bad request response has a 5xx status code
func (o *CronJobSyncOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync organizations bad request response a status code equal to that given
func (o *CronJobSyncOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobSyncOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncOrganizationsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncOrganizationsUnauthorized creates a CronJobSyncOrganizationsUnauthorized with default headers values
func NewCronJobSyncOrganizationsUnauthorized() *CronJobSyncOrganizationsUnauthorized {
	return &CronJobSyncOrganizationsUnauthorized{}
}

/*
CronJobSyncOrganizationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobSyncOrganizationsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync organizations unauthorized response has a 2xx status code
func (o *CronJobSyncOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync organizations unauthorized response has a 3xx status code
func (o *CronJobSyncOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations unauthorized response has a 4xx status code
func (o *CronJobSyncOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync organizations unauthorized response has a 5xx status code
func (o *CronJobSyncOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync organizations unauthorized response a status code equal to that given
func (o *CronJobSyncOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobSyncOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncOrganizationsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncOrganizationsForbidden creates a CronJobSyncOrganizationsForbidden with default headers values
func NewCronJobSyncOrganizationsForbidden() *CronJobSyncOrganizationsForbidden {
	return &CronJobSyncOrganizationsForbidden{}
}

/*
CronJobSyncOrganizationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobSyncOrganizationsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync organizations forbidden response has a 2xx status code
func (o *CronJobSyncOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync organizations forbidden response has a 3xx status code
func (o *CronJobSyncOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations forbidden response has a 4xx status code
func (o *CronJobSyncOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync organizations forbidden response has a 5xx status code
func (o *CronJobSyncOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync organizations forbidden response a status code equal to that given
func (o *CronJobSyncOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobSyncOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncOrganizationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncOrganizationsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncOrganizationsNotFound creates a CronJobSyncOrganizationsNotFound with default headers values
func NewCronJobSyncOrganizationsNotFound() *CronJobSyncOrganizationsNotFound {
	return &CronJobSyncOrganizationsNotFound{}
}

/*
CronJobSyncOrganizationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobSyncOrganizationsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job sync organizations not found response has a 2xx status code
func (o *CronJobSyncOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync organizations not found response has a 3xx status code
func (o *CronJobSyncOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations not found response has a 4xx status code
func (o *CronJobSyncOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync organizations not found response has a 5xx status code
func (o *CronJobSyncOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync organizations not found response a status code equal to that given
func (o *CronJobSyncOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobSyncOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncOrganizationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncOrganizationsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSyncOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncOrganizationsInternalServerError creates a CronJobSyncOrganizationsInternalServerError with default headers values
func NewCronJobSyncOrganizationsInternalServerError() *CronJobSyncOrganizationsInternalServerError {
	return &CronJobSyncOrganizationsInternalServerError{}
}

/*
CronJobSyncOrganizationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobSyncOrganizationsInternalServerError struct {
}

// IsSuccess returns true when this cron job sync organizations internal server error response has a 2xx status code
func (o *CronJobSyncOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync organizations internal server error response has a 3xx status code
func (o *CronJobSyncOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync organizations internal server error response has a 4xx status code
func (o *CronJobSyncOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync organizations internal server error response has a 5xx status code
func (o *CronJobSyncOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job sync organizations internal server error response a status code equal to that given
func (o *CronJobSyncOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobSyncOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsInternalServerError ", 500)
}

func (o *CronJobSyncOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-organizations][%d] cronJobSyncOrganizationsInternalServerError ", 500)
}

func (o *CronJobSyncOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
