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

// CronJobDeleteExpiredAlertsReader is a Reader for the CronJobDeleteExpiredAlerts structure.
type CronJobDeleteExpiredAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteExpiredAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteExpiredAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteExpiredAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteExpiredAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteExpiredAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteExpiredAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteExpiredAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteExpiredAlertsOK creates a CronJobDeleteExpiredAlertsOK with default headers values
func NewCronJobDeleteExpiredAlertsOK() *CronJobDeleteExpiredAlertsOK {
	return &CronJobDeleteExpiredAlertsOK{}
}

/*
CronJobDeleteExpiredAlertsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredAlertsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job delete expired alerts o k response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job delete expired alerts o k response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts o k response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired alerts o k response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired alerts o k response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobDeleteExpiredAlertsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteExpiredAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredAlertsBadRequest creates a CronJobDeleteExpiredAlertsBadRequest with default headers values
func NewCronJobDeleteExpiredAlertsBadRequest() *CronJobDeleteExpiredAlertsBadRequest {
	return &CronJobDeleteExpiredAlertsBadRequest{}
}

/*
CronJobDeleteExpiredAlertsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredAlertsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cron job delete expired alerts bad request response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired alerts bad request response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts bad request response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired alerts bad request response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired alerts bad request response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobDeleteExpiredAlertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CronJobDeleteExpiredAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredAlertsUnauthorized creates a CronJobDeleteExpiredAlertsUnauthorized with default headers values
func NewCronJobDeleteExpiredAlertsUnauthorized() *CronJobDeleteExpiredAlertsUnauthorized {
	return &CronJobDeleteExpiredAlertsUnauthorized{}
}

/*
CronJobDeleteExpiredAlertsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredAlertsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job delete expired alerts unauthorized response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired alerts unauthorized response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts unauthorized response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired alerts unauthorized response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired alerts unauthorized response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobDeleteExpiredAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobDeleteExpiredAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredAlertsForbidden creates a CronJobDeleteExpiredAlertsForbidden with default headers values
func NewCronJobDeleteExpiredAlertsForbidden() *CronJobDeleteExpiredAlertsForbidden {
	return &CronJobDeleteExpiredAlertsForbidden{}
}

/*
CronJobDeleteExpiredAlertsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredAlertsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job delete expired alerts forbidden response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired alerts forbidden response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts forbidden response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired alerts forbidden response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired alerts forbidden response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobDeleteExpiredAlertsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobDeleteExpiredAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredAlertsNotFound creates a CronJobDeleteExpiredAlertsNotFound with default headers values
func NewCronJobDeleteExpiredAlertsNotFound() *CronJobDeleteExpiredAlertsNotFound {
	return &CronJobDeleteExpiredAlertsNotFound{}
}

/*
CronJobDeleteExpiredAlertsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredAlertsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job delete expired alerts not found response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired alerts not found response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts not found response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired alerts not found response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired alerts not found response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobDeleteExpiredAlertsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredAlertsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobDeleteExpiredAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredAlertsInternalServerError creates a CronJobDeleteExpiredAlertsInternalServerError with default headers values
func NewCronJobDeleteExpiredAlertsInternalServerError() *CronJobDeleteExpiredAlertsInternalServerError {
	return &CronJobDeleteExpiredAlertsInternalServerError{}
}

/*
CronJobDeleteExpiredAlertsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredAlertsInternalServerError struct {
}

// IsSuccess returns true when this cron job delete expired alerts internal server error response has a 2xx status code
func (o *CronJobDeleteExpiredAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired alerts internal server error response has a 3xx status code
func (o *CronJobDeleteExpiredAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired alerts internal server error response has a 4xx status code
func (o *CronJobDeleteExpiredAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired alerts internal server error response has a 5xx status code
func (o *CronJobDeleteExpiredAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job delete expired alerts internal server error response a status code equal to that given
func (o *CronJobDeleteExpiredAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobDeleteExpiredAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredAlertsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/alerts][%d] cronJobDeleteExpiredAlertsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
