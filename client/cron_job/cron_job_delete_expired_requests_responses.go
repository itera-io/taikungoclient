// Code generated by go-swagger; DO NOT EDIT.

package cron_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
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

/*
CronJobDeleteExpiredRequestsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteExpiredRequestsOK struct {
}

// IsSuccess returns true when this cron job delete expired requests o k response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job delete expired requests o k response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests o k response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired requests o k response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired requests o k response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job delete expired requests o k response
func (o *CronJobDeleteExpiredRequestsOK) Code() int {
	return 200
}

func (o *CronJobDeleteExpiredRequestsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsOK ", 200)
}

func (o *CronJobDeleteExpiredRequestsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsOK ", 200)
}

func (o *CronJobDeleteExpiredRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCronJobDeleteExpiredRequestsBadRequest creates a CronJobDeleteExpiredRequestsBadRequest with default headers values
func NewCronJobDeleteExpiredRequestsBadRequest() *CronJobDeleteExpiredRequestsBadRequest {
	return &CronJobDeleteExpiredRequestsBadRequest{}
}

/*
CronJobDeleteExpiredRequestsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteExpiredRequestsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job delete expired requests bad request response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired requests bad request response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests bad request response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired requests bad request response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired requests bad request response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job delete expired requests bad request response
func (o *CronJobDeleteExpiredRequestsBadRequest) Code() int {
	return 400
}

func (o *CronJobDeleteExpiredRequestsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsUnauthorized creates a CronJobDeleteExpiredRequestsUnauthorized with default headers values
func NewCronJobDeleteExpiredRequestsUnauthorized() *CronJobDeleteExpiredRequestsUnauthorized {
	return &CronJobDeleteExpiredRequestsUnauthorized{}
}

/*
CronJobDeleteExpiredRequestsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteExpiredRequestsUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job delete expired requests unauthorized response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired requests unauthorized response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests unauthorized response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired requests unauthorized response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired requests unauthorized response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job delete expired requests unauthorized response
func (o *CronJobDeleteExpiredRequestsUnauthorized) Code() int {
	return 401
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsForbidden creates a CronJobDeleteExpiredRequestsForbidden with default headers values
func NewCronJobDeleteExpiredRequestsForbidden() *CronJobDeleteExpiredRequestsForbidden {
	return &CronJobDeleteExpiredRequestsForbidden{}
}

/*
CronJobDeleteExpiredRequestsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteExpiredRequestsForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job delete expired requests forbidden response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired requests forbidden response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests forbidden response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired requests forbidden response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired requests forbidden response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job delete expired requests forbidden response
func (o *CronJobDeleteExpiredRequestsForbidden) Code() int {
	return 403
}

func (o *CronJobDeleteExpiredRequestsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsNotFound creates a CronJobDeleteExpiredRequestsNotFound with default headers values
func NewCronJobDeleteExpiredRequestsNotFound() *CronJobDeleteExpiredRequestsNotFound {
	return &CronJobDeleteExpiredRequestsNotFound{}
}

/*
CronJobDeleteExpiredRequestsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteExpiredRequestsNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job delete expired requests not found response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired requests not found response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests not found response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete expired requests not found response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete expired requests not found response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job delete expired requests not found response
func (o *CronJobDeleteExpiredRequestsNotFound) Code() int {
	return 404
}

func (o *CronJobDeleteExpiredRequestsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteExpiredRequestsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobDeleteExpiredRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteExpiredRequestsInternalServerError creates a CronJobDeleteExpiredRequestsInternalServerError with default headers values
func NewCronJobDeleteExpiredRequestsInternalServerError() *CronJobDeleteExpiredRequestsInternalServerError {
	return &CronJobDeleteExpiredRequestsInternalServerError{}
}

/*
CronJobDeleteExpiredRequestsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteExpiredRequestsInternalServerError struct {
}

// IsSuccess returns true when this cron job delete expired requests internal server error response has a 2xx status code
func (o *CronJobDeleteExpiredRequestsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete expired requests internal server error response has a 3xx status code
func (o *CronJobDeleteExpiredRequestsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete expired requests internal server error response has a 4xx status code
func (o *CronJobDeleteExpiredRequestsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete expired requests internal server error response has a 5xx status code
func (o *CronJobDeleteExpiredRequestsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job delete expired requests internal server error response a status code equal to that given
func (o *CronJobDeleteExpiredRequestsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job delete expired requests internal server error response
func (o *CronJobDeleteExpiredRequestsInternalServerError) Code() int {
	return 500
}

func (o *CronJobDeleteExpiredRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredRequestsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/taikun-requests][%d] cronJobDeleteExpiredRequestsInternalServerError ", 500)
}

func (o *CronJobDeleteExpiredRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
