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

// CronJobDeleteAwsSpotInstancesReader is a Reader for the CronJobDeleteAwsSpotInstances structure.
type CronJobDeleteAwsSpotInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobDeleteAwsSpotInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobDeleteAwsSpotInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobDeleteAwsSpotInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobDeleteAwsSpotInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobDeleteAwsSpotInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobDeleteAwsSpotInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobDeleteAwsSpotInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobDeleteAwsSpotInstancesOK creates a CronJobDeleteAwsSpotInstancesOK with default headers values
func NewCronJobDeleteAwsSpotInstancesOK() *CronJobDeleteAwsSpotInstancesOK {
	return &CronJobDeleteAwsSpotInstancesOK{}
}

/*
CronJobDeleteAwsSpotInstancesOK describes a response with status code 200, with default header values.

Success
*/
type CronJobDeleteAwsSpotInstancesOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job delete aws spot instances o k response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job delete aws spot instances o k response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances o k response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete aws spot instances o k response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete aws spot instances o k response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cron job delete aws spot instances o k response
func (o *CronJobDeleteAwsSpotInstancesOK) Code() int {
	return 200
}

func (o *CronJobDeleteAwsSpotInstancesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesOK  %+v", 200, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobDeleteAwsSpotInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteAwsSpotInstancesBadRequest creates a CronJobDeleteAwsSpotInstancesBadRequest with default headers values
func NewCronJobDeleteAwsSpotInstancesBadRequest() *CronJobDeleteAwsSpotInstancesBadRequest {
	return &CronJobDeleteAwsSpotInstancesBadRequest{}
}

/*
CronJobDeleteAwsSpotInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobDeleteAwsSpotInstancesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete aws spot instances bad request response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete aws spot instances bad request response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances bad request response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete aws spot instances bad request response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete aws spot instances bad request response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cron job delete aws spot instances bad request response
func (o *CronJobDeleteAwsSpotInstancesBadRequest) Code() int {
	return 400
}

func (o *CronJobDeleteAwsSpotInstancesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteAwsSpotInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteAwsSpotInstancesUnauthorized creates a CronJobDeleteAwsSpotInstancesUnauthorized with default headers values
func NewCronJobDeleteAwsSpotInstancesUnauthorized() *CronJobDeleteAwsSpotInstancesUnauthorized {
	return &CronJobDeleteAwsSpotInstancesUnauthorized{}
}

/*
CronJobDeleteAwsSpotInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobDeleteAwsSpotInstancesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete aws spot instances unauthorized response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete aws spot instances unauthorized response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances unauthorized response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete aws spot instances unauthorized response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete aws spot instances unauthorized response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cron job delete aws spot instances unauthorized response
func (o *CronJobDeleteAwsSpotInstancesUnauthorized) Code() int {
	return 401
}

func (o *CronJobDeleteAwsSpotInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteAwsSpotInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteAwsSpotInstancesForbidden creates a CronJobDeleteAwsSpotInstancesForbidden with default headers values
func NewCronJobDeleteAwsSpotInstancesForbidden() *CronJobDeleteAwsSpotInstancesForbidden {
	return &CronJobDeleteAwsSpotInstancesForbidden{}
}

/*
CronJobDeleteAwsSpotInstancesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobDeleteAwsSpotInstancesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete aws spot instances forbidden response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete aws spot instances forbidden response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances forbidden response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete aws spot instances forbidden response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete aws spot instances forbidden response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cron job delete aws spot instances forbidden response
func (o *CronJobDeleteAwsSpotInstancesForbidden) Code() int {
	return 403
}

func (o *CronJobDeleteAwsSpotInstancesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteAwsSpotInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteAwsSpotInstancesNotFound creates a CronJobDeleteAwsSpotInstancesNotFound with default headers values
func NewCronJobDeleteAwsSpotInstancesNotFound() *CronJobDeleteAwsSpotInstancesNotFound {
	return &CronJobDeleteAwsSpotInstancesNotFound{}
}

/*
CronJobDeleteAwsSpotInstancesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobDeleteAwsSpotInstancesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job delete aws spot instances not found response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete aws spot instances not found response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances not found response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job delete aws spot instances not found response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job delete aws spot instances not found response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the cron job delete aws spot instances not found response
func (o *CronJobDeleteAwsSpotInstancesNotFound) Code() int {
	return 404
}

func (o *CronJobDeleteAwsSpotInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobDeleteAwsSpotInstancesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobDeleteAwsSpotInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobDeleteAwsSpotInstancesInternalServerError creates a CronJobDeleteAwsSpotInstancesInternalServerError with default headers values
func NewCronJobDeleteAwsSpotInstancesInternalServerError() *CronJobDeleteAwsSpotInstancesInternalServerError {
	return &CronJobDeleteAwsSpotInstancesInternalServerError{}
}

/*
CronJobDeleteAwsSpotInstancesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobDeleteAwsSpotInstancesInternalServerError struct {
}

// IsSuccess returns true when this cron job delete aws spot instances internal server error response has a 2xx status code
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job delete aws spot instances internal server error response has a 3xx status code
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job delete aws spot instances internal server error response has a 4xx status code
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job delete aws spot instances internal server error response has a 5xx status code
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job delete aws spot instances internal server error response a status code equal to that given
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the cron job delete aws spot instances internal server error response
func (o *CronJobDeleteAwsSpotInstancesInternalServerError) Code() int {
	return 500
}

func (o *CronJobDeleteAwsSpotInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesInternalServerError ", 500)
}

func (o *CronJobDeleteAwsSpotInstancesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/delete-aws-spot-instances][%d] cronJobDeleteAwsSpotInstancesInternalServerError ", 500)
}

func (o *CronJobDeleteAwsSpotInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
