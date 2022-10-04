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

// CronJobSyncProjectAppsReader is a Reader for the CronJobSyncProjectApps structure.
type CronJobSyncProjectAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobSyncProjectAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobSyncProjectAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobSyncProjectAppsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobSyncProjectAppsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobSyncProjectAppsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobSyncProjectAppsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobSyncProjectAppsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobSyncProjectAppsOK creates a CronJobSyncProjectAppsOK with default headers values
func NewCronJobSyncProjectAppsOK() *CronJobSyncProjectAppsOK {
	return &CronJobSyncProjectAppsOK{}
}

/*
CronJobSyncProjectAppsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobSyncProjectAppsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job sync project apps o k response has a 2xx status code
func (o *CronJobSyncProjectAppsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job sync project apps o k response has a 3xx status code
func (o *CronJobSyncProjectAppsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps o k response has a 4xx status code
func (o *CronJobSyncProjectAppsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync project apps o k response has a 5xx status code
func (o *CronJobSyncProjectAppsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync project apps o k response a status code equal to that given
func (o *CronJobSyncProjectAppsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobSyncProjectAppsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncProjectAppsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsOK  %+v", 200, o.Payload)
}

func (o *CronJobSyncProjectAppsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobSyncProjectAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectAppsBadRequest creates a CronJobSyncProjectAppsBadRequest with default headers values
func NewCronJobSyncProjectAppsBadRequest() *CronJobSyncProjectAppsBadRequest {
	return &CronJobSyncProjectAppsBadRequest{}
}

/*
CronJobSyncProjectAppsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobSyncProjectAppsBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job sync project apps bad request response has a 2xx status code
func (o *CronJobSyncProjectAppsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync project apps bad request response has a 3xx status code
func (o *CronJobSyncProjectAppsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps bad request response has a 4xx status code
func (o *CronJobSyncProjectAppsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync project apps bad request response has a 5xx status code
func (o *CronJobSyncProjectAppsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync project apps bad request response a status code equal to that given
func (o *CronJobSyncProjectAppsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobSyncProjectAppsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncProjectAppsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobSyncProjectAppsBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobSyncProjectAppsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectAppsUnauthorized creates a CronJobSyncProjectAppsUnauthorized with default headers values
func NewCronJobSyncProjectAppsUnauthorized() *CronJobSyncProjectAppsUnauthorized {
	return &CronJobSyncProjectAppsUnauthorized{}
}

/*
CronJobSyncProjectAppsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobSyncProjectAppsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job sync project apps unauthorized response has a 2xx status code
func (o *CronJobSyncProjectAppsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync project apps unauthorized response has a 3xx status code
func (o *CronJobSyncProjectAppsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps unauthorized response has a 4xx status code
func (o *CronJobSyncProjectAppsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync project apps unauthorized response has a 5xx status code
func (o *CronJobSyncProjectAppsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync project apps unauthorized response a status code equal to that given
func (o *CronJobSyncProjectAppsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobSyncProjectAppsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncProjectAppsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobSyncProjectAppsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobSyncProjectAppsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectAppsForbidden creates a CronJobSyncProjectAppsForbidden with default headers values
func NewCronJobSyncProjectAppsForbidden() *CronJobSyncProjectAppsForbidden {
	return &CronJobSyncProjectAppsForbidden{}
}

/*
CronJobSyncProjectAppsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobSyncProjectAppsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job sync project apps forbidden response has a 2xx status code
func (o *CronJobSyncProjectAppsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync project apps forbidden response has a 3xx status code
func (o *CronJobSyncProjectAppsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps forbidden response has a 4xx status code
func (o *CronJobSyncProjectAppsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync project apps forbidden response has a 5xx status code
func (o *CronJobSyncProjectAppsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync project apps forbidden response a status code equal to that given
func (o *CronJobSyncProjectAppsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobSyncProjectAppsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncProjectAppsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobSyncProjectAppsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobSyncProjectAppsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectAppsNotFound creates a CronJobSyncProjectAppsNotFound with default headers values
func NewCronJobSyncProjectAppsNotFound() *CronJobSyncProjectAppsNotFound {
	return &CronJobSyncProjectAppsNotFound{}
}

/*
CronJobSyncProjectAppsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobSyncProjectAppsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job sync project apps not found response has a 2xx status code
func (o *CronJobSyncProjectAppsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync project apps not found response has a 3xx status code
func (o *CronJobSyncProjectAppsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps not found response has a 4xx status code
func (o *CronJobSyncProjectAppsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job sync project apps not found response has a 5xx status code
func (o *CronJobSyncProjectAppsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job sync project apps not found response a status code equal to that given
func (o *CronJobSyncProjectAppsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobSyncProjectAppsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncProjectAppsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobSyncProjectAppsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobSyncProjectAppsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobSyncProjectAppsInternalServerError creates a CronJobSyncProjectAppsInternalServerError with default headers values
func NewCronJobSyncProjectAppsInternalServerError() *CronJobSyncProjectAppsInternalServerError {
	return &CronJobSyncProjectAppsInternalServerError{}
}

/*
CronJobSyncProjectAppsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobSyncProjectAppsInternalServerError struct {
}

// IsSuccess returns true when this cron job sync project apps internal server error response has a 2xx status code
func (o *CronJobSyncProjectAppsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job sync project apps internal server error response has a 3xx status code
func (o *CronJobSyncProjectAppsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job sync project apps internal server error response has a 4xx status code
func (o *CronJobSyncProjectAppsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job sync project apps internal server error response has a 5xx status code
func (o *CronJobSyncProjectAppsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job sync project apps internal server error response a status code equal to that given
func (o *CronJobSyncProjectAppsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobSyncProjectAppsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsInternalServerError ", 500)
}

func (o *CronJobSyncProjectAppsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/sync-project-apps][%d] cronJobSyncProjectAppsInternalServerError ", 500)
}

func (o *CronJobSyncProjectAppsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
