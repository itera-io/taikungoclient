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

// CronJobTriggerTemplatesReader is a Reader for the CronJobTriggerTemplates structure.
type CronJobTriggerTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobTriggerTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobTriggerTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobTriggerTemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobTriggerTemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobTriggerTemplatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobTriggerTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobTriggerTemplatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobTriggerTemplatesOK creates a CronJobTriggerTemplatesOK with default headers values
func NewCronJobTriggerTemplatesOK() *CronJobTriggerTemplatesOK {
	return &CronJobTriggerTemplatesOK{}
}

/*
CronJobTriggerTemplatesOK describes a response with status code 200, with default header values.

Success
*/
type CronJobTriggerTemplatesOK struct {
}

// IsSuccess returns true when this cron job trigger templates o k response has a 2xx status code
func (o *CronJobTriggerTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job trigger templates o k response has a 3xx status code
func (o *CronJobTriggerTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates o k response has a 4xx status code
func (o *CronJobTriggerTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job trigger templates o k response has a 5xx status code
func (o *CronJobTriggerTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job trigger templates o k response a status code equal to that given
func (o *CronJobTriggerTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job trigger templates o k response
func (o *CronJobTriggerTemplatesOK) Code() int {
	return 200
}

func (o *CronJobTriggerTemplatesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesOK ", 200)
}

func (o *CronJobTriggerTemplatesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesOK ", 200)
}

func (o *CronJobTriggerTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCronJobTriggerTemplatesBadRequest creates a CronJobTriggerTemplatesBadRequest with default headers values
func NewCronJobTriggerTemplatesBadRequest() *CronJobTriggerTemplatesBadRequest {
	return &CronJobTriggerTemplatesBadRequest{}
}

/*
CronJobTriggerTemplatesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobTriggerTemplatesBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job trigger templates bad request response has a 2xx status code
func (o *CronJobTriggerTemplatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job trigger templates bad request response has a 3xx status code
func (o *CronJobTriggerTemplatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates bad request response has a 4xx status code
func (o *CronJobTriggerTemplatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job trigger templates bad request response has a 5xx status code
func (o *CronJobTriggerTemplatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job trigger templates bad request response a status code equal to that given
func (o *CronJobTriggerTemplatesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job trigger templates bad request response
func (o *CronJobTriggerTemplatesBadRequest) Code() int {
	return 400
}

func (o *CronJobTriggerTemplatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobTriggerTemplatesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobTriggerTemplatesBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobTriggerTemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobTriggerTemplatesUnauthorized creates a CronJobTriggerTemplatesUnauthorized with default headers values
func NewCronJobTriggerTemplatesUnauthorized() *CronJobTriggerTemplatesUnauthorized {
	return &CronJobTriggerTemplatesUnauthorized{}
}

/*
CronJobTriggerTemplatesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobTriggerTemplatesUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job trigger templates unauthorized response has a 2xx status code
func (o *CronJobTriggerTemplatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job trigger templates unauthorized response has a 3xx status code
func (o *CronJobTriggerTemplatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates unauthorized response has a 4xx status code
func (o *CronJobTriggerTemplatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job trigger templates unauthorized response has a 5xx status code
func (o *CronJobTriggerTemplatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job trigger templates unauthorized response a status code equal to that given
func (o *CronJobTriggerTemplatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job trigger templates unauthorized response
func (o *CronJobTriggerTemplatesUnauthorized) Code() int {
	return 401
}

func (o *CronJobTriggerTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobTriggerTemplatesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobTriggerTemplatesUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobTriggerTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobTriggerTemplatesForbidden creates a CronJobTriggerTemplatesForbidden with default headers values
func NewCronJobTriggerTemplatesForbidden() *CronJobTriggerTemplatesForbidden {
	return &CronJobTriggerTemplatesForbidden{}
}

/*
CronJobTriggerTemplatesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobTriggerTemplatesForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job trigger templates forbidden response has a 2xx status code
func (o *CronJobTriggerTemplatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job trigger templates forbidden response has a 3xx status code
func (o *CronJobTriggerTemplatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates forbidden response has a 4xx status code
func (o *CronJobTriggerTemplatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job trigger templates forbidden response has a 5xx status code
func (o *CronJobTriggerTemplatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job trigger templates forbidden response a status code equal to that given
func (o *CronJobTriggerTemplatesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job trigger templates forbidden response
func (o *CronJobTriggerTemplatesForbidden) Code() int {
	return 403
}

func (o *CronJobTriggerTemplatesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobTriggerTemplatesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobTriggerTemplatesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobTriggerTemplatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobTriggerTemplatesNotFound creates a CronJobTriggerTemplatesNotFound with default headers values
func NewCronJobTriggerTemplatesNotFound() *CronJobTriggerTemplatesNotFound {
	return &CronJobTriggerTemplatesNotFound{}
}

/*
CronJobTriggerTemplatesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobTriggerTemplatesNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job trigger templates not found response has a 2xx status code
func (o *CronJobTriggerTemplatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job trigger templates not found response has a 3xx status code
func (o *CronJobTriggerTemplatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates not found response has a 4xx status code
func (o *CronJobTriggerTemplatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job trigger templates not found response has a 5xx status code
func (o *CronJobTriggerTemplatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job trigger templates not found response a status code equal to that given
func (o *CronJobTriggerTemplatesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job trigger templates not found response
func (o *CronJobTriggerTemplatesNotFound) Code() int {
	return 404
}

func (o *CronJobTriggerTemplatesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobTriggerTemplatesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobTriggerTemplatesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobTriggerTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobTriggerTemplatesInternalServerError creates a CronJobTriggerTemplatesInternalServerError with default headers values
func NewCronJobTriggerTemplatesInternalServerError() *CronJobTriggerTemplatesInternalServerError {
	return &CronJobTriggerTemplatesInternalServerError{}
}

/*
CronJobTriggerTemplatesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobTriggerTemplatesInternalServerError struct {
}

// IsSuccess returns true when this cron job trigger templates internal server error response has a 2xx status code
func (o *CronJobTriggerTemplatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job trigger templates internal server error response has a 3xx status code
func (o *CronJobTriggerTemplatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job trigger templates internal server error response has a 4xx status code
func (o *CronJobTriggerTemplatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job trigger templates internal server error response has a 5xx status code
func (o *CronJobTriggerTemplatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job trigger templates internal server error response a status code equal to that given
func (o *CronJobTriggerTemplatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job trigger templates internal server error response
func (o *CronJobTriggerTemplatesInternalServerError) Code() int {
	return 500
}

func (o *CronJobTriggerTemplatesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesInternalServerError ", 500)
}

func (o *CronJobTriggerTemplatesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/trigger-templates][%d] cronJobTriggerTemplatesInternalServerError ", 500)
}

func (o *CronJobTriggerTemplatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}