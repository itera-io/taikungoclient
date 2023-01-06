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

// CronJobFetchArtifactHubOrganizationsReader is a Reader for the CronJobFetchArtifactHubOrganizations structure.
type CronJobFetchArtifactHubOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobFetchArtifactHubOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobFetchArtifactHubOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobFetchArtifactHubOrganizationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobFetchArtifactHubOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobFetchArtifactHubOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobFetchArtifactHubOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobFetchArtifactHubOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobFetchArtifactHubOrganizationsOK creates a CronJobFetchArtifactHubOrganizationsOK with default headers values
func NewCronJobFetchArtifactHubOrganizationsOK() *CronJobFetchArtifactHubOrganizationsOK {
	return &CronJobFetchArtifactHubOrganizationsOK{}
}

/*
CronJobFetchArtifactHubOrganizationsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobFetchArtifactHubOrganizationsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job fetch artifact hub organizations o k response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job fetch artifact hub organizations o k response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations o k response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch artifact hub organizations o k response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch artifact hub organizations o k response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobFetchArtifactHubOrganizationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobFetchArtifactHubOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchArtifactHubOrganizationsBadRequest creates a CronJobFetchArtifactHubOrganizationsBadRequest with default headers values
func NewCronJobFetchArtifactHubOrganizationsBadRequest() *CronJobFetchArtifactHubOrganizationsBadRequest {
	return &CronJobFetchArtifactHubOrganizationsBadRequest{}
}

/*
CronJobFetchArtifactHubOrganizationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobFetchArtifactHubOrganizationsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch artifact hub organizations bad request response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch artifact hub organizations bad request response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations bad request response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch artifact hub organizations bad request response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch artifact hub organizations bad request response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobFetchArtifactHubOrganizationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchArtifactHubOrganizationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchArtifactHubOrganizationsUnauthorized creates a CronJobFetchArtifactHubOrganizationsUnauthorized with default headers values
func NewCronJobFetchArtifactHubOrganizationsUnauthorized() *CronJobFetchArtifactHubOrganizationsUnauthorized {
	return &CronJobFetchArtifactHubOrganizationsUnauthorized{}
}

/*
CronJobFetchArtifactHubOrganizationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobFetchArtifactHubOrganizationsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch artifact hub organizations unauthorized response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch artifact hub organizations unauthorized response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations unauthorized response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch artifact hub organizations unauthorized response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch artifact hub organizations unauthorized response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchArtifactHubOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchArtifactHubOrganizationsForbidden creates a CronJobFetchArtifactHubOrganizationsForbidden with default headers values
func NewCronJobFetchArtifactHubOrganizationsForbidden() *CronJobFetchArtifactHubOrganizationsForbidden {
	return &CronJobFetchArtifactHubOrganizationsForbidden{}
}

/*
CronJobFetchArtifactHubOrganizationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobFetchArtifactHubOrganizationsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch artifact hub organizations forbidden response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch artifact hub organizations forbidden response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations forbidden response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch artifact hub organizations forbidden response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch artifact hub organizations forbidden response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobFetchArtifactHubOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchArtifactHubOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchArtifactHubOrganizationsNotFound creates a CronJobFetchArtifactHubOrganizationsNotFound with default headers values
func NewCronJobFetchArtifactHubOrganizationsNotFound() *CronJobFetchArtifactHubOrganizationsNotFound {
	return &CronJobFetchArtifactHubOrganizationsNotFound{}
}

/*
CronJobFetchArtifactHubOrganizationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobFetchArtifactHubOrganizationsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch artifact hub organizations not found response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch artifact hub organizations not found response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations not found response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch artifact hub organizations not found response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch artifact hub organizations not found response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobFetchArtifactHubOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchArtifactHubOrganizationsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchArtifactHubOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchArtifactHubOrganizationsInternalServerError creates a CronJobFetchArtifactHubOrganizationsInternalServerError with default headers values
func NewCronJobFetchArtifactHubOrganizationsInternalServerError() *CronJobFetchArtifactHubOrganizationsInternalServerError {
	return &CronJobFetchArtifactHubOrganizationsInternalServerError{}
}

/*
CronJobFetchArtifactHubOrganizationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobFetchArtifactHubOrganizationsInternalServerError struct {
}

// IsSuccess returns true when this cron job fetch artifact hub organizations internal server error response has a 2xx status code
func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch artifact hub organizations internal server error response has a 3xx status code
func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch artifact hub organizations internal server error response has a 4xx status code
func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch artifact hub organizations internal server error response has a 5xx status code
func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job fetch artifact hub organizations internal server error response a status code equal to that given
func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsInternalServerError ", 500)
}

func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-artifact-organizations][%d] cronJobFetchArtifactHubOrganizationsInternalServerError ", 500)
}

func (o *CronJobFetchArtifactHubOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}