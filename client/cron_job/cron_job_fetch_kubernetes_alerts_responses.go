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

// CronJobFetchKubernetesAlertsReader is a Reader for the CronJobFetchKubernetesAlerts structure.
type CronJobFetchKubernetesAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobFetchKubernetesAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobFetchKubernetesAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobFetchKubernetesAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobFetchKubernetesAlertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobFetchKubernetesAlertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobFetchKubernetesAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobFetchKubernetesAlertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobFetchKubernetesAlertsOK creates a CronJobFetchKubernetesAlertsOK with default headers values
func NewCronJobFetchKubernetesAlertsOK() *CronJobFetchKubernetesAlertsOK {
	return &CronJobFetchKubernetesAlertsOK{}
}

/*
CronJobFetchKubernetesAlertsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobFetchKubernetesAlertsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job fetch kubernetes alerts o k response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job fetch kubernetes alerts o k response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts o k response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch kubernetes alerts o k response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch kubernetes alerts o k response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobFetchKubernetesAlertsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobFetchKubernetesAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchKubernetesAlertsBadRequest creates a CronJobFetchKubernetesAlertsBadRequest with default headers values
func NewCronJobFetchKubernetesAlertsBadRequest() *CronJobFetchKubernetesAlertsBadRequest {
	return &CronJobFetchKubernetesAlertsBadRequest{}
}

/*
CronJobFetchKubernetesAlertsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobFetchKubernetesAlertsBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this cron job fetch kubernetes alerts bad request response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch kubernetes alerts bad request response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts bad request response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch kubernetes alerts bad request response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch kubernetes alerts bad request response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobFetchKubernetesAlertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CronJobFetchKubernetesAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchKubernetesAlertsUnauthorized creates a CronJobFetchKubernetesAlertsUnauthorized with default headers values
func NewCronJobFetchKubernetesAlertsUnauthorized() *CronJobFetchKubernetesAlertsUnauthorized {
	return &CronJobFetchKubernetesAlertsUnauthorized{}
}

/*
CronJobFetchKubernetesAlertsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobFetchKubernetesAlertsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch kubernetes alerts unauthorized response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch kubernetes alerts unauthorized response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts unauthorized response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch kubernetes alerts unauthorized response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch kubernetes alerts unauthorized response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobFetchKubernetesAlertsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchKubernetesAlertsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchKubernetesAlertsForbidden creates a CronJobFetchKubernetesAlertsForbidden with default headers values
func NewCronJobFetchKubernetesAlertsForbidden() *CronJobFetchKubernetesAlertsForbidden {
	return &CronJobFetchKubernetesAlertsForbidden{}
}

/*
CronJobFetchKubernetesAlertsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobFetchKubernetesAlertsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch kubernetes alerts forbidden response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch kubernetes alerts forbidden response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts forbidden response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch kubernetes alerts forbidden response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch kubernetes alerts forbidden response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobFetchKubernetesAlertsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchKubernetesAlertsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchKubernetesAlertsNotFound creates a CronJobFetchKubernetesAlertsNotFound with default headers values
func NewCronJobFetchKubernetesAlertsNotFound() *CronJobFetchKubernetesAlertsNotFound {
	return &CronJobFetchKubernetesAlertsNotFound{}
}

/*
CronJobFetchKubernetesAlertsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobFetchKubernetesAlertsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch kubernetes alerts not found response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch kubernetes alerts not found response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts not found response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch kubernetes alerts not found response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch kubernetes alerts not found response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobFetchKubernetesAlertsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchKubernetesAlertsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchKubernetesAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchKubernetesAlertsInternalServerError creates a CronJobFetchKubernetesAlertsInternalServerError with default headers values
func NewCronJobFetchKubernetesAlertsInternalServerError() *CronJobFetchKubernetesAlertsInternalServerError {
	return &CronJobFetchKubernetesAlertsInternalServerError{}
}

/*
CronJobFetchKubernetesAlertsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobFetchKubernetesAlertsInternalServerError struct {
}

// IsSuccess returns true when this cron job fetch kubernetes alerts internal server error response has a 2xx status code
func (o *CronJobFetchKubernetesAlertsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch kubernetes alerts internal server error response has a 3xx status code
func (o *CronJobFetchKubernetesAlertsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch kubernetes alerts internal server error response has a 4xx status code
func (o *CronJobFetchKubernetesAlertsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch kubernetes alerts internal server error response has a 5xx status code
func (o *CronJobFetchKubernetesAlertsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job fetch kubernetes alerts internal server error response a status code equal to that given
func (o *CronJobFetchKubernetesAlertsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobFetchKubernetesAlertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsInternalServerError ", 500)
}

func (o *CronJobFetchKubernetesAlertsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-kubernetes-alerts][%d] cronJobFetchKubernetesAlertsInternalServerError ", 500)
}

func (o *CronJobFetchKubernetesAlertsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
