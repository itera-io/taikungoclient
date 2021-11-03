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

// CronJobDeleteExpiredHistoryLogsReader is a Reader for the CronJobDeleteExpiredHistoryLogs structure.
type CronJobDeleteExpiredHistoryLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteExpiredHistoryLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteExpiredHistoryLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteExpiredHistoryLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteExpiredHistoryLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteExpiredHistoryLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteExpiredHistoryLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteExpiredHistoryLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteExpiredHistoryLogsOK creates a CronJobDeleteExpiredHistoryLogsOK with default headers values
func NewCronJobDeleteExpiredHistoryLogsOK() *CronJobDeleteExpiredHistoryLogsOK {
	return &CronJobDeleteExpiredHistoryLogsOK{}
}

/* CronJobDeleteExpiredHistoryLogsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredHistoryLogsOK struct {
	Payload models.Unit
}

func (o *CronJobDeleteExpiredHistoryLogsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsOK  %+v", 200, o.Payload)
}
func (o *CronJobDeleteExpiredHistoryLogsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteExpiredHistoryLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredHistoryLogsBadRequest creates a CronJobDeleteExpiredHistoryLogsBadRequest with default headers values
func NewCronJobDeleteExpiredHistoryLogsBadRequest() *CronJobDeleteExpiredHistoryLogsBadRequest {
	return &CronJobDeleteExpiredHistoryLogsBadRequest{}
}

/* CronJobDeleteExpiredHistoryLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredHistoryLogsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *CronJobDeleteExpiredHistoryLogsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsBadRequest  %+v", 400, o.Payload)
}
func (o *CronJobDeleteExpiredHistoryLogsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredHistoryLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredHistoryLogsUnauthorized creates a CronJobDeleteExpiredHistoryLogsUnauthorized with default headers values
func NewCronJobDeleteExpiredHistoryLogsUnauthorized() *CronJobDeleteExpiredHistoryLogsUnauthorized {
	return &CronJobDeleteExpiredHistoryLogsUnauthorized{}
}

/* CronJobDeleteExpiredHistoryLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredHistoryLogsUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredHistoryLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsUnauthorized  %+v", 401, o.Payload)
}
func (o *CronJobDeleteExpiredHistoryLogsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredHistoryLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredHistoryLogsForbidden creates a CronJobDeleteExpiredHistoryLogsForbidden with default headers values
func NewCronJobDeleteExpiredHistoryLogsForbidden() *CronJobDeleteExpiredHistoryLogsForbidden {
	return &CronJobDeleteExpiredHistoryLogsForbidden{}
}

/* CronJobDeleteExpiredHistoryLogsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredHistoryLogsForbidden struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredHistoryLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsForbidden  %+v", 403, o.Payload)
}
func (o *CronJobDeleteExpiredHistoryLogsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredHistoryLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredHistoryLogsNotFound creates a CronJobDeleteExpiredHistoryLogsNotFound with default headers values
func NewCronJobDeleteExpiredHistoryLogsNotFound() *CronJobDeleteExpiredHistoryLogsNotFound {
	return &CronJobDeleteExpiredHistoryLogsNotFound{}
}

/* CronJobDeleteExpiredHistoryLogsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredHistoryLogsNotFound struct {
	Payload *models.ProblemDetails
}

func (o *CronJobDeleteExpiredHistoryLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsNotFound  %+v", 404, o.Payload)
}
func (o *CronJobDeleteExpiredHistoryLogsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredHistoryLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredHistoryLogsInternalServerError creates a CronJobDeleteExpiredHistoryLogsInternalServerError with default headers values
func NewCronJobDeleteExpiredHistoryLogsInternalServerError() *CronJobDeleteExpiredHistoryLogsInternalServerError {
	return &CronJobDeleteExpiredHistoryLogsInternalServerError{}
}

/* CronJobDeleteExpiredHistoryLogsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredHistoryLogsInternalServerError struct {
}

func (o *CronJobDeleteExpiredHistoryLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/history-logs][%d] cronJobDeleteExpiredHistoryLogsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredHistoryLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
