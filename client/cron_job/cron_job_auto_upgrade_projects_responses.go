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

// CronJobAutoUpgradeProjectsReader is a Reader for the CronJobAutoUpgradeProjects structure.
type CronJobAutoUpgradeProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobAutoUpgradeProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobAutoUpgradeProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobAutoUpgradeProjectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobAutoUpgradeProjectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobAutoUpgradeProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobAutoUpgradeProjectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobAutoUpgradeProjectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobAutoUpgradeProjectsOK creates a CronJobAutoUpgradeProjectsOK with default headers values
func NewCronJobAutoUpgradeProjectsOK() *CronJobAutoUpgradeProjectsOK {
	return &CronJobAutoUpgradeProjectsOK{}
}

/*
CronJobAutoUpgradeProjectsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobAutoUpgradeProjectsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job auto upgrade projects o k response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job auto upgrade projects o k response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects o k response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job auto upgrade projects o k response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job auto upgrade projects o k response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobAutoUpgradeProjectsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsOK  %+v", 200, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsOK  %+v", 200, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobAutoUpgradeProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobAutoUpgradeProjectsBadRequest creates a CronJobAutoUpgradeProjectsBadRequest with default headers values
func NewCronJobAutoUpgradeProjectsBadRequest() *CronJobAutoUpgradeProjectsBadRequest {
	return &CronJobAutoUpgradeProjectsBadRequest{}
}

/*
CronJobAutoUpgradeProjectsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobAutoUpgradeProjectsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this cron job auto upgrade projects bad request response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job auto upgrade projects bad request response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects bad request response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job auto upgrade projects bad request response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job auto upgrade projects bad request response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobAutoUpgradeProjectsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CronJobAutoUpgradeProjectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobAutoUpgradeProjectsUnauthorized creates a CronJobAutoUpgradeProjectsUnauthorized with default headers values
func NewCronJobAutoUpgradeProjectsUnauthorized() *CronJobAutoUpgradeProjectsUnauthorized {
	return &CronJobAutoUpgradeProjectsUnauthorized{}
}

/*
CronJobAutoUpgradeProjectsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobAutoUpgradeProjectsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job auto upgrade projects unauthorized response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job auto upgrade projects unauthorized response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects unauthorized response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job auto upgrade projects unauthorized response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job auto upgrade projects unauthorized response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobAutoUpgradeProjectsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobAutoUpgradeProjectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobAutoUpgradeProjectsForbidden creates a CronJobAutoUpgradeProjectsForbidden with default headers values
func NewCronJobAutoUpgradeProjectsForbidden() *CronJobAutoUpgradeProjectsForbidden {
	return &CronJobAutoUpgradeProjectsForbidden{}
}

/*
CronJobAutoUpgradeProjectsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobAutoUpgradeProjectsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job auto upgrade projects forbidden response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job auto upgrade projects forbidden response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects forbidden response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job auto upgrade projects forbidden response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job auto upgrade projects forbidden response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobAutoUpgradeProjectsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobAutoUpgradeProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobAutoUpgradeProjectsNotFound creates a CronJobAutoUpgradeProjectsNotFound with default headers values
func NewCronJobAutoUpgradeProjectsNotFound() *CronJobAutoUpgradeProjectsNotFound {
	return &CronJobAutoUpgradeProjectsNotFound{}
}

/*
CronJobAutoUpgradeProjectsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobAutoUpgradeProjectsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job auto upgrade projects not found response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job auto upgrade projects not found response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects not found response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job auto upgrade projects not found response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job auto upgrade projects not found response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobAutoUpgradeProjectsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobAutoUpgradeProjectsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobAutoUpgradeProjectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobAutoUpgradeProjectsInternalServerError creates a CronJobAutoUpgradeProjectsInternalServerError with default headers values
func NewCronJobAutoUpgradeProjectsInternalServerError() *CronJobAutoUpgradeProjectsInternalServerError {
	return &CronJobAutoUpgradeProjectsInternalServerError{}
}

/*
CronJobAutoUpgradeProjectsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobAutoUpgradeProjectsInternalServerError struct {
}

// IsSuccess returns true when this cron job auto upgrade projects internal server error response has a 2xx status code
func (o *CronJobAutoUpgradeProjectsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job auto upgrade projects internal server error response has a 3xx status code
func (o *CronJobAutoUpgradeProjectsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job auto upgrade projects internal server error response has a 4xx status code
func (o *CronJobAutoUpgradeProjectsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job auto upgrade projects internal server error response has a 5xx status code
func (o *CronJobAutoUpgradeProjectsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job auto upgrade projects internal server error response a status code equal to that given
func (o *CronJobAutoUpgradeProjectsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobAutoUpgradeProjectsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsInternalServerError ", 500)
}

func (o *CronJobAutoUpgradeProjectsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/auto-upgrade-projects][%d] cronJobAutoUpgradeProjectsInternalServerError ", 500)
}

func (o *CronJobAutoUpgradeProjectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
