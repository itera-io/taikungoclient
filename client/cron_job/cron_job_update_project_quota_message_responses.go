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

// CronJobUpdateProjectQuotaMessageReader is a Reader for the CronJobUpdateProjectQuotaMessage structure.
type CronJobUpdateProjectQuotaMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobUpdateProjectQuotaMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobUpdateProjectQuotaMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobUpdateProjectQuotaMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobUpdateProjectQuotaMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobUpdateProjectQuotaMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobUpdateProjectQuotaMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobUpdateProjectQuotaMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobUpdateProjectQuotaMessageOK creates a CronJobUpdateProjectQuotaMessageOK with default headers values
func NewCronJobUpdateProjectQuotaMessageOK() *CronJobUpdateProjectQuotaMessageOK {
	return &CronJobUpdateProjectQuotaMessageOK{}
}

/*
CronJobUpdateProjectQuotaMessageOK describes a response with status code 200, with default header values.

Success
*/
type CronJobUpdateProjectQuotaMessageOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job update project quota message o k response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job update project quota message o k response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message o k response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job update project quota message o k response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project quota message o k response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job update project quota message o k response
func (o *CronJobUpdateProjectQuotaMessageOK) Code() int {
	return 200
}

func (o *CronJobUpdateProjectQuotaMessageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageOK  %+v", 200, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageOK  %+v", 200, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobUpdateProjectQuotaMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectQuotaMessageBadRequest creates a CronJobUpdateProjectQuotaMessageBadRequest with default headers values
func NewCronJobUpdateProjectQuotaMessageBadRequest() *CronJobUpdateProjectQuotaMessageBadRequest {
	return &CronJobUpdateProjectQuotaMessageBadRequest{}
}

/*
CronJobUpdateProjectQuotaMessageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobUpdateProjectQuotaMessageBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job update project quota message bad request response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project quota message bad request response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message bad request response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project quota message bad request response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project quota message bad request response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job update project quota message bad request response
func (o *CronJobUpdateProjectQuotaMessageBadRequest) Code() int {
	return 400
}

func (o *CronJobUpdateProjectQuotaMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobUpdateProjectQuotaMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectQuotaMessageUnauthorized creates a CronJobUpdateProjectQuotaMessageUnauthorized with default headers values
func NewCronJobUpdateProjectQuotaMessageUnauthorized() *CronJobUpdateProjectQuotaMessageUnauthorized {
	return &CronJobUpdateProjectQuotaMessageUnauthorized{}
}

/*
CronJobUpdateProjectQuotaMessageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobUpdateProjectQuotaMessageUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job update project quota message unauthorized response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project quota message unauthorized response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message unauthorized response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project quota message unauthorized response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project quota message unauthorized response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job update project quota message unauthorized response
func (o *CronJobUpdateProjectQuotaMessageUnauthorized) Code() int {
	return 401
}

func (o *CronJobUpdateProjectQuotaMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobUpdateProjectQuotaMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectQuotaMessageForbidden creates a CronJobUpdateProjectQuotaMessageForbidden with default headers values
func NewCronJobUpdateProjectQuotaMessageForbidden() *CronJobUpdateProjectQuotaMessageForbidden {
	return &CronJobUpdateProjectQuotaMessageForbidden{}
}

/*
CronJobUpdateProjectQuotaMessageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobUpdateProjectQuotaMessageForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job update project quota message forbidden response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project quota message forbidden response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message forbidden response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project quota message forbidden response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project quota message forbidden response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job update project quota message forbidden response
func (o *CronJobUpdateProjectQuotaMessageForbidden) Code() int {
	return 403
}

func (o *CronJobUpdateProjectQuotaMessageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageForbidden  %+v", 403, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageForbidden  %+v", 403, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobUpdateProjectQuotaMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectQuotaMessageNotFound creates a CronJobUpdateProjectQuotaMessageNotFound with default headers values
func NewCronJobUpdateProjectQuotaMessageNotFound() *CronJobUpdateProjectQuotaMessageNotFound {
	return &CronJobUpdateProjectQuotaMessageNotFound{}
}

/*
CronJobUpdateProjectQuotaMessageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobUpdateProjectQuotaMessageNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this cron job update project quota message not found response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project quota message not found response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message not found response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job update project quota message not found response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job update project quota message not found response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job update project quota message not found response
func (o *CronJobUpdateProjectQuotaMessageNotFound) Code() int {
	return 404
}

func (o *CronJobUpdateProjectQuotaMessageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageNotFound  %+v", 404, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageNotFound  %+v", 404, o.Payload)
}

func (o *CronJobUpdateProjectQuotaMessageNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CronJobUpdateProjectQuotaMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobUpdateProjectQuotaMessageInternalServerError creates a CronJobUpdateProjectQuotaMessageInternalServerError with default headers values
func NewCronJobUpdateProjectQuotaMessageInternalServerError() *CronJobUpdateProjectQuotaMessageInternalServerError {
	return &CronJobUpdateProjectQuotaMessageInternalServerError{}
}

/*
CronJobUpdateProjectQuotaMessageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobUpdateProjectQuotaMessageInternalServerError struct {
}

// IsSuccess returns true when this cron job update project quota message internal server error response has a 2xx status code
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job update project quota message internal server error response has a 3xx status code
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job update project quota message internal server error response has a 4xx status code
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job update project quota message internal server error response has a 5xx status code
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job update project quota message internal server error response a status code equal to that given
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job update project quota message internal server error response
func (o *CronJobUpdateProjectQuotaMessageInternalServerError) Code() int {
	return 500
}

func (o *CronJobUpdateProjectQuotaMessageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageInternalServerError ", 500)
}

func (o *CronJobUpdateProjectQuotaMessageInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/project-quota-message][%d] cronJobUpdateProjectQuotaMessageInternalServerError ", 500)
}

func (o *CronJobUpdateProjectQuotaMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
