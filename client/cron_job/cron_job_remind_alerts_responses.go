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

// CronJobRemindAlertsReader is a Reader for the CronJobRemindAlerts structure.
type CronJobRemindAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobRemindAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobRemindAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobRemindAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobRemindAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobRemindAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobRemindAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobRemindAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobRemindAlertsOK creates a CronJobRemindAlertsOK with default headers values
func NewCronJobRemindAlertsOK() *CronJobRemindAlertsOK {
	return &CronJobRemindAlertsOK{}
}

/*
CronJobRemindAlertsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobRemindAlertsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job remind alerts o k response has a 2xx status code
func (o *CronJobRemindAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job remind alerts o k response has a 3xx status code
func (o *CronJobRemindAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts o k response has a 4xx status code
func (o *CronJobRemindAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job remind alerts o k response has a 5xx status code
func (o *CronJobRemindAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job remind alerts o k response a status code equal to that given
func (o *CronJobRemindAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobRemindAlertsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobRemindAlertsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobRemindAlertsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobRemindAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobRemindAlertsBadRequest creates a CronJobRemindAlertsBadRequest with default headers values
func NewCronJobRemindAlertsBadRequest() *CronJobRemindAlertsBadRequest {
	return &CronJobRemindAlertsBadRequest{}
}

/*
CronJobRemindAlertsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobRemindAlertsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cron job remind alerts bad request response has a 2xx status code
func (o *CronJobRemindAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job remind alerts bad request response has a 3xx status code
func (o *CronJobRemindAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts bad request response has a 4xx status code
func (o *CronJobRemindAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job remind alerts bad request response has a 5xx status code
func (o *CronJobRemindAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job remind alerts bad request response a status code equal to that given
func (o *CronJobRemindAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobRemindAlertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobRemindAlertsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobRemindAlertsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CronJobRemindAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobRemindAlertsUnauthorized creates a CronJobRemindAlertsUnauthorized with default headers values
func NewCronJobRemindAlertsUnauthorized() *CronJobRemindAlertsUnauthorized {
	return &CronJobRemindAlertsUnauthorized{}
}

/*
CronJobRemindAlertsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobRemindAlertsUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job remind alerts unauthorized response has a 2xx status code
func (o *CronJobRemindAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job remind alerts unauthorized response has a 3xx status code
func (o *CronJobRemindAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts unauthorized response has a 4xx status code
func (o *CronJobRemindAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job remind alerts unauthorized response has a 5xx status code
func (o *CronJobRemindAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job remind alerts unauthorized response a status code equal to that given
func (o *CronJobRemindAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobRemindAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobRemindAlertsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobRemindAlertsUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobRemindAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobRemindAlertsForbidden creates a CronJobRemindAlertsForbidden with default headers values
func NewCronJobRemindAlertsForbidden() *CronJobRemindAlertsForbidden {
	return &CronJobRemindAlertsForbidden{}
}

/*
CronJobRemindAlertsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobRemindAlertsForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job remind alerts forbidden response has a 2xx status code
func (o *CronJobRemindAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job remind alerts forbidden response has a 3xx status code
func (o *CronJobRemindAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts forbidden response has a 4xx status code
func (o *CronJobRemindAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job remind alerts forbidden response has a 5xx status code
func (o *CronJobRemindAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job remind alerts forbidden response a status code equal to that given
func (o *CronJobRemindAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobRemindAlertsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobRemindAlertsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobRemindAlertsForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobRemindAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobRemindAlertsNotFound creates a CronJobRemindAlertsNotFound with default headers values
func NewCronJobRemindAlertsNotFound() *CronJobRemindAlertsNotFound {
	return &CronJobRemindAlertsNotFound{}
}

/*
CronJobRemindAlertsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobRemindAlertsNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job remind alerts not found response has a 2xx status code
func (o *CronJobRemindAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job remind alerts not found response has a 3xx status code
func (o *CronJobRemindAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts not found response has a 4xx status code
func (o *CronJobRemindAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job remind alerts not found response has a 5xx status code
func (o *CronJobRemindAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job remind alerts not found response a status code equal to that given
func (o *CronJobRemindAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobRemindAlertsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobRemindAlertsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobRemindAlertsNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobRemindAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobRemindAlertsInternalServerError creates a CronJobRemindAlertsInternalServerError with default headers values
func NewCronJobRemindAlertsInternalServerError() *CronJobRemindAlertsInternalServerError {
	return &CronJobRemindAlertsInternalServerError{}
}

/*
CronJobRemindAlertsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobRemindAlertsInternalServerError struct {
}

// IsSuccess returns true when this cron job remind alerts internal server error response has a 2xx status code
func (o *CronJobRemindAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job remind alerts internal server error response has a 3xx status code
func (o *CronJobRemindAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job remind alerts internal server error response has a 4xx status code
func (o *CronJobRemindAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job remind alerts internal server error response has a 5xx status code
func (o *CronJobRemindAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job remind alerts internal server error response a status code equal to that given
func (o *CronJobRemindAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobRemindAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsInternalServerError ", 500)
}

func (o *CronJobRemindAlertsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/remind-alerts][%d] cronJobRemindAlertsInternalServerError ", 500)
}

func (o *CronJobRemindAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
