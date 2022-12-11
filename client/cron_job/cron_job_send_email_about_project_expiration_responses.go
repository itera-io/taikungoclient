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

// CronJobSendEmailAboutProjectExpirationReader is a Reader for the CronJobSendEmailAboutProjectExpiration structure.
type CronJobSendEmailAboutProjectExpirationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobSendEmailAboutProjectExpirationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobSendEmailAboutProjectExpirationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobSendEmailAboutProjectExpirationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobSendEmailAboutProjectExpirationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobSendEmailAboutProjectExpirationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobSendEmailAboutProjectExpirationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobSendEmailAboutProjectExpirationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobSendEmailAboutProjectExpirationOK creates a CronJobSendEmailAboutProjectExpirationOK with default headers values
func NewCronJobSendEmailAboutProjectExpirationOK() *CronJobSendEmailAboutProjectExpirationOK {
	return &CronJobSendEmailAboutProjectExpirationOK{}
}

/*
CronJobSendEmailAboutProjectExpirationOK describes a response with status code 200, with default header values.

Success
*/
type CronJobSendEmailAboutProjectExpirationOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job send email about project expiration o k response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job send email about project expiration o k response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration o k response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job send email about project expiration o k response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job send email about project expiration o k response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobSendEmailAboutProjectExpirationOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationOK  %+v", 200, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationOK  %+v", 200, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobSendEmailAboutProjectExpirationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSendEmailAboutProjectExpirationBadRequest creates a CronJobSendEmailAboutProjectExpirationBadRequest with default headers values
func NewCronJobSendEmailAboutProjectExpirationBadRequest() *CronJobSendEmailAboutProjectExpirationBadRequest {
	return &CronJobSendEmailAboutProjectExpirationBadRequest{}
}

/*
CronJobSendEmailAboutProjectExpirationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobSendEmailAboutProjectExpirationBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job send email about project expiration bad request response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job send email about project expiration bad request response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration bad request response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job send email about project expiration bad request response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job send email about project expiration bad request response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobSendEmailAboutProjectExpirationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobSendEmailAboutProjectExpirationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSendEmailAboutProjectExpirationUnauthorized creates a CronJobSendEmailAboutProjectExpirationUnauthorized with default headers values
func NewCronJobSendEmailAboutProjectExpirationUnauthorized() *CronJobSendEmailAboutProjectExpirationUnauthorized {
	return &CronJobSendEmailAboutProjectExpirationUnauthorized{}
}

/*
CronJobSendEmailAboutProjectExpirationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobSendEmailAboutProjectExpirationUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job send email about project expiration unauthorized response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job send email about project expiration unauthorized response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration unauthorized response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job send email about project expiration unauthorized response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job send email about project expiration unauthorized response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSendEmailAboutProjectExpirationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSendEmailAboutProjectExpirationForbidden creates a CronJobSendEmailAboutProjectExpirationForbidden with default headers values
func NewCronJobSendEmailAboutProjectExpirationForbidden() *CronJobSendEmailAboutProjectExpirationForbidden {
	return &CronJobSendEmailAboutProjectExpirationForbidden{}
}

/*
CronJobSendEmailAboutProjectExpirationForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobSendEmailAboutProjectExpirationForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job send email about project expiration forbidden response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job send email about project expiration forbidden response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration forbidden response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job send email about project expiration forbidden response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job send email about project expiration forbidden response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobSendEmailAboutProjectExpirationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSendEmailAboutProjectExpirationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSendEmailAboutProjectExpirationNotFound creates a CronJobSendEmailAboutProjectExpirationNotFound with default headers values
func NewCronJobSendEmailAboutProjectExpirationNotFound() *CronJobSendEmailAboutProjectExpirationNotFound {
	return &CronJobSendEmailAboutProjectExpirationNotFound{}
}

/*
CronJobSendEmailAboutProjectExpirationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobSendEmailAboutProjectExpirationNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job send email about project expiration not found response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job send email about project expiration not found response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration not found response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job send email about project expiration not found response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job send email about project expiration not found response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobSendEmailAboutProjectExpirationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSendEmailAboutProjectExpirationNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobSendEmailAboutProjectExpirationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSendEmailAboutProjectExpirationInternalServerError creates a CronJobSendEmailAboutProjectExpirationInternalServerError with default headers values
func NewCronJobSendEmailAboutProjectExpirationInternalServerError() *CronJobSendEmailAboutProjectExpirationInternalServerError {
	return &CronJobSendEmailAboutProjectExpirationInternalServerError{}
}

/*
CronJobSendEmailAboutProjectExpirationInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobSendEmailAboutProjectExpirationInternalServerError struct {
}

// IsSuccess returns true when this cron job send email about project expiration internal server error response has a 2xx status code
func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job send email about project expiration internal server error response has a 3xx status code
func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job send email about project expiration internal server error response has a 4xx status code
func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job send email about project expiration internal server error response has a 5xx status code
func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job send email about project expiration internal server error response a status code equal to that given
func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationInternalServerError ", 500)
}

func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-expiration][%d] cronJobSendEmailAboutProjectExpirationInternalServerError ", 500)
}

func (o *CronJobSendEmailAboutProjectExpirationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
