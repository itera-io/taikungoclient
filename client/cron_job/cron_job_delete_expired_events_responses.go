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

// CronJobDeleteExpiredEventsReader is a Reader for the CronJobDeleteExpiredEvents structure.
type CronJobDeleteExpiredEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteExpiredEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteExpiredEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteExpiredEventsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteExpiredEventsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteExpiredEventsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteExpiredEventsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteExpiredEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteExpiredEventsOK creates a CronJobDeleteExpiredEventsOK with default headers values
func NewCronJobDeleteExpiredEventsOK() *CronJobDeleteExpiredEventsOK {
	return &CronJobDeleteExpiredEventsOK{}
}

/*
CronJobDeleteExpiredEventsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredEventsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job delete expired events o k response has a 2xx status code
func (o *CronJobDeleteExpiredEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job delete expired events o k response has a 3xx status code
func (o *CronJobDeleteExpiredEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events o k response has a 4xx status code
func (o *CronJobDeleteExpiredEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired events o k response has a 5xx status code
func (o *CronJobDeleteExpiredEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired events o k response a status code equal to that given
func (o *CronJobDeleteExpiredEventsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobDeleteExpiredEventsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredEventsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredEventsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteExpiredEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredEventsBadRequest creates a CronJobDeleteExpiredEventsBadRequest with default headers values
func NewCronJobDeleteExpiredEventsBadRequest() *CronJobDeleteExpiredEventsBadRequest {
	return &CronJobDeleteExpiredEventsBadRequest{}
}

/*
CronJobDeleteExpiredEventsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredEventsBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired events bad request response has a 2xx status code
func (o *CronJobDeleteExpiredEventsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired events bad request response has a 3xx status code
func (o *CronJobDeleteExpiredEventsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events bad request response has a 4xx status code
func (o *CronJobDeleteExpiredEventsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired events bad request response has a 5xx status code
func (o *CronJobDeleteExpiredEventsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired events bad request response a status code equal to that given
func (o *CronJobDeleteExpiredEventsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobDeleteExpiredEventsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredEventsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredEventsBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredEventsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredEventsUnauthorized creates a CronJobDeleteExpiredEventsUnauthorized with default headers values
func NewCronJobDeleteExpiredEventsUnauthorized() *CronJobDeleteExpiredEventsUnauthorized {
	return &CronJobDeleteExpiredEventsUnauthorized{}
}

/*
CronJobDeleteExpiredEventsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredEventsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired events unauthorized response has a 2xx status code
func (o *CronJobDeleteExpiredEventsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired events unauthorized response has a 3xx status code
func (o *CronJobDeleteExpiredEventsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events unauthorized response has a 4xx status code
func (o *CronJobDeleteExpiredEventsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired events unauthorized response has a 5xx status code
func (o *CronJobDeleteExpiredEventsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired events unauthorized response a status code equal to that given
func (o *CronJobDeleteExpiredEventsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobDeleteExpiredEventsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredEventsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredEventsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredEventsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredEventsForbidden creates a CronJobDeleteExpiredEventsForbidden with default headers values
func NewCronJobDeleteExpiredEventsForbidden() *CronJobDeleteExpiredEventsForbidden {
	return &CronJobDeleteExpiredEventsForbidden{}
}

/*
CronJobDeleteExpiredEventsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredEventsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired events forbidden response has a 2xx status code
func (o *CronJobDeleteExpiredEventsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired events forbidden response has a 3xx status code
func (o *CronJobDeleteExpiredEventsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events forbidden response has a 4xx status code
func (o *CronJobDeleteExpiredEventsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired events forbidden response has a 5xx status code
func (o *CronJobDeleteExpiredEventsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired events forbidden response a status code equal to that given
func (o *CronJobDeleteExpiredEventsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobDeleteExpiredEventsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredEventsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredEventsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredEventsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredEventsNotFound creates a CronJobDeleteExpiredEventsNotFound with default headers values
func NewCronJobDeleteExpiredEventsNotFound() *CronJobDeleteExpiredEventsNotFound {
	return &CronJobDeleteExpiredEventsNotFound{}
}

/*
CronJobDeleteExpiredEventsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredEventsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete expired events not found response has a 2xx status code
func (o *CronJobDeleteExpiredEventsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired events not found response has a 3xx status code
func (o *CronJobDeleteExpiredEventsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events not found response has a 4xx status code
func (o *CronJobDeleteExpiredEventsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired events not found response has a 5xx status code
func (o *CronJobDeleteExpiredEventsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired events not found response a status code equal to that given
func (o *CronJobDeleteExpiredEventsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobDeleteExpiredEventsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredEventsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredEventsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteExpiredEventsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredEventsInternalServerError creates a CronJobDeleteExpiredEventsInternalServerError with default headers values
func NewCronJobDeleteExpiredEventsInternalServerError() *CronJobDeleteExpiredEventsInternalServerError {
	return &CronJobDeleteExpiredEventsInternalServerError{}
}

/*
CronJobDeleteExpiredEventsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredEventsInternalServerError struct {
}

// IsSuccess returns true when this cron job delete expired events internal server error response has a 2xx status code
func (o *CronJobDeleteExpiredEventsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired events internal server error response has a 3xx status code
func (o *CronJobDeleteExpiredEventsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired events internal server error response has a 4xx status code
func (o *CronJobDeleteExpiredEventsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired events internal server error response has a 5xx status code
func (o *CronJobDeleteExpiredEventsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job delete expired events internal server error response a status code equal to that given
func (o *CronJobDeleteExpiredEventsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobDeleteExpiredEventsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredEventsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/events][%d] cronJobDeleteExpiredEventsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
