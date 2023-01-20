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

// CronJobCreateKeyPoolReader is a Reader for the CronJobCreateKeyPool structure.
type CronJobCreateKeyPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobCreateKeyPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobCreateKeyPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobCreateKeyPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobCreateKeyPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobCreateKeyPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobCreateKeyPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobCreateKeyPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobCreateKeyPoolOK creates a CronJobCreateKeyPoolOK with default headers values
func NewCronJobCreateKeyPoolOK() *CronJobCreateKeyPoolOK {
	return &CronJobCreateKeyPoolOK{}
}

/*
CronJobCreateKeyPoolOK describes a response with status code 200, with default header values.

Success
*/
type CronJobCreateKeyPoolOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job create key pool o k response has a 2xx status code
func (o *CronJobCreateKeyPoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job create key pool o k response has a 3xx status code
func (o *CronJobCreateKeyPoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool o k response has a 4xx status code
func (o *CronJobCreateKeyPoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job create key pool o k response has a 5xx status code
func (o *CronJobCreateKeyPoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job create key pool o k response a status code equal to that given
func (o *CronJobCreateKeyPoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job create key pool o k response
func (o *CronJobCreateKeyPoolOK) Code() int {
	return 200
}

func (o *CronJobCreateKeyPoolOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolOK  %+v", 200, o.Payload)
}

func (o *CronJobCreateKeyPoolOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolOK  %+v", 200, o.Payload)
}

func (o *CronJobCreateKeyPoolOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobCreateKeyPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobCreateKeyPoolBadRequest creates a CronJobCreateKeyPoolBadRequest with default headers values
func NewCronJobCreateKeyPoolBadRequest() *CronJobCreateKeyPoolBadRequest {
	return &CronJobCreateKeyPoolBadRequest{}
}

/*
CronJobCreateKeyPoolBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobCreateKeyPoolBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job create key pool bad request response has a 2xx status code
func (o *CronJobCreateKeyPoolBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job create key pool bad request response has a 3xx status code
func (o *CronJobCreateKeyPoolBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool bad request response has a 4xx status code
func (o *CronJobCreateKeyPoolBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job create key pool bad request response has a 5xx status code
func (o *CronJobCreateKeyPoolBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job create key pool bad request response a status code equal to that given
func (o *CronJobCreateKeyPoolBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job create key pool bad request response
func (o *CronJobCreateKeyPoolBadRequest) Code() int {
	return 400
}

func (o *CronJobCreateKeyPoolBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobCreateKeyPoolBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobCreateKeyPoolBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobCreateKeyPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobCreateKeyPoolUnauthorized creates a CronJobCreateKeyPoolUnauthorized with default headers values
func NewCronJobCreateKeyPoolUnauthorized() *CronJobCreateKeyPoolUnauthorized {
	return &CronJobCreateKeyPoolUnauthorized{}
}

/*
CronJobCreateKeyPoolUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobCreateKeyPoolUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job create key pool unauthorized response has a 2xx status code
func (o *CronJobCreateKeyPoolUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job create key pool unauthorized response has a 3xx status code
func (o *CronJobCreateKeyPoolUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool unauthorized response has a 4xx status code
func (o *CronJobCreateKeyPoolUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job create key pool unauthorized response has a 5xx status code
func (o *CronJobCreateKeyPoolUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job create key pool unauthorized response a status code equal to that given
func (o *CronJobCreateKeyPoolUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job create key pool unauthorized response
func (o *CronJobCreateKeyPoolUnauthorized) Code() int {
	return 401
}

func (o *CronJobCreateKeyPoolUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobCreateKeyPoolUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobCreateKeyPoolUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobCreateKeyPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobCreateKeyPoolForbidden creates a CronJobCreateKeyPoolForbidden with default headers values
func NewCronJobCreateKeyPoolForbidden() *CronJobCreateKeyPoolForbidden {
	return &CronJobCreateKeyPoolForbidden{}
}

/*
CronJobCreateKeyPoolForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobCreateKeyPoolForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job create key pool forbidden response has a 2xx status code
func (o *CronJobCreateKeyPoolForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job create key pool forbidden response has a 3xx status code
func (o *CronJobCreateKeyPoolForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool forbidden response has a 4xx status code
func (o *CronJobCreateKeyPoolForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job create key pool forbidden response has a 5xx status code
func (o *CronJobCreateKeyPoolForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job create key pool forbidden response a status code equal to that given
func (o *CronJobCreateKeyPoolForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job create key pool forbidden response
func (o *CronJobCreateKeyPoolForbidden) Code() int {
	return 403
}

func (o *CronJobCreateKeyPoolForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolForbidden  %+v", 403, o.Payload)
}

func (o *CronJobCreateKeyPoolForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolForbidden  %+v", 403, o.Payload)
}

func (o *CronJobCreateKeyPoolForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobCreateKeyPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobCreateKeyPoolNotFound creates a CronJobCreateKeyPoolNotFound with default headers values
func NewCronJobCreateKeyPoolNotFound() *CronJobCreateKeyPoolNotFound {
	return &CronJobCreateKeyPoolNotFound{}
}

/*
CronJobCreateKeyPoolNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobCreateKeyPoolNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job create key pool not found response has a 2xx status code
func (o *CronJobCreateKeyPoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job create key pool not found response has a 3xx status code
func (o *CronJobCreateKeyPoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool not found response has a 4xx status code
func (o *CronJobCreateKeyPoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job create key pool not found response has a 5xx status code
func (o *CronJobCreateKeyPoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job create key pool not found response a status code equal to that given
func (o *CronJobCreateKeyPoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job create key pool not found response
func (o *CronJobCreateKeyPoolNotFound) Code() int {
	return 404
}

func (o *CronJobCreateKeyPoolNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolNotFound  %+v", 404, o.Payload)
}

func (o *CronJobCreateKeyPoolNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolNotFound  %+v", 404, o.Payload)
}

func (o *CronJobCreateKeyPoolNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobCreateKeyPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobCreateKeyPoolInternalServerError creates a CronJobCreateKeyPoolInternalServerError with default headers values
func NewCronJobCreateKeyPoolInternalServerError() *CronJobCreateKeyPoolInternalServerError {
	return &CronJobCreateKeyPoolInternalServerError{}
}

/*
CronJobCreateKeyPoolInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobCreateKeyPoolInternalServerError struct {
}

// IsSuccess returns true when this cron job create key pool internal server error response has a 2xx status code
func (o *CronJobCreateKeyPoolInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job create key pool internal server error response has a 3xx status code
func (o *CronJobCreateKeyPoolInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job create key pool internal server error response has a 4xx status code
func (o *CronJobCreateKeyPoolInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job create key pool internal server error response has a 5xx status code
func (o *CronJobCreateKeyPoolInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job create key pool internal server error response a status code equal to that given
func (o *CronJobCreateKeyPoolInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job create key pool internal server error response
func (o *CronJobCreateKeyPoolInternalServerError) Code() int {
	return 500
}

func (o *CronJobCreateKeyPoolInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolInternalServerError ", 500)
}

func (o *CronJobCreateKeyPoolInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/create-key-pool][%d] cronJobCreateKeyPoolInternalServerError ", 500)
}

func (o *CronJobCreateKeyPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
