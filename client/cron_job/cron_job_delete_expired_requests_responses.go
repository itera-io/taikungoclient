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

// CronJobDeleteExpiredRequestsReader is a Reader for the CronJobDeleteExpiredRequests structure.
type CronJobDeleteExpiredRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteExpiredRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteExpiredRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteExpiredRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteExpiredRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteExpiredRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteExpiredRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteExpiredRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteExpiredRequestsOK creates a CronJobDeleteExpiredRequestsOK with default headers values
func NewCronJobDeleteExpiredRequestsOK() *CronJobDeleteExpiredRequestsOK {
	return &CronJobDeleteExpiredRequestsOK{}
}

/* CronJobDeleteExpiredRequestsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredRequestsOK struct {
	Payload models.Unit
}

func (o *CronJobDeleteExpiredRequestsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsOK  %+v", 200, o.Payload)
}
func (o *CronJobDeleteExpiredRequestsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsBadRequest creates a CronJobDeleteExpiredRequestsBadRequest with default headers values
func NewCronJobDeleteExpiredRequestsBadRequest() *CronJobDeleteExpiredRequestsBadRequest {
	return &CronJobDeleteExpiredRequestsBadRequest{}
}

/* CronJobDeleteExpiredRequestsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredRequestsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CronJobDeleteExpiredRequestsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsBadRequest  %+v", 400, o.Payload)
}
func (o *CronJobDeleteExpiredRequestsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsUnauthorized creates a CronJobDeleteExpiredRequestsUnauthorized with default headers values
func NewCronJobDeleteExpiredRequestsUnauthorized() *CronJobDeleteExpiredRequestsUnauthorized {
	return &CronJobDeleteExpiredRequestsUnauthorized{}
}

/* CronJobDeleteExpiredRequestsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredRequestsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsUnauthorized  %+v", 401, o.Payload)
}
func (o *CronJobDeleteExpiredRequestsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsForbidden creates a CronJobDeleteExpiredRequestsForbidden with default headers values
func NewCronJobDeleteExpiredRequestsForbidden() *CronJobDeleteExpiredRequestsForbidden {
	return &CronJobDeleteExpiredRequestsForbidden{}
}

/* CronJobDeleteExpiredRequestsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredRequestsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredRequestsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsForbidden  %+v", 403, o.Payload)
}
func (o *CronJobDeleteExpiredRequestsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsNotFound creates a CronJobDeleteExpiredRequestsNotFound with default headers values
func NewCronJobDeleteExpiredRequestsNotFound() *CronJobDeleteExpiredRequestsNotFound {
	return &CronJobDeleteExpiredRequestsNotFound{}
}

/* CronJobDeleteExpiredRequestsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredRequestsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredRequestsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsNotFound  %+v", 404, o.Payload)
}
func (o *CronJobDeleteExpiredRequestsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsInternalServerError creates a CronJobDeleteExpiredRequestsInternalServerError with default headers values
func NewCronJobDeleteExpiredRequestsInternalServerError() *CronJobDeleteExpiredRequestsInternalServerError {
	return &CronJobDeleteExpiredRequestsInternalServerError{}
}

/* CronJobDeleteExpiredRequestsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredRequestsInternalServerError struct {
}

func (o *CronJobDeleteExpiredRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}