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

// CronJobFetchOrganizationDetailsReader is a Reader for the CronJobFetchOrganizationDetails structure.
type CronJobFetchOrganizationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobFetchOrganizationDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobFetchOrganizationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobFetchOrganizationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobFetchOrganizationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobFetchOrganizationDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobFetchOrganizationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobFetchOrganizationDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobFetchOrganizationDetailsOK creates a CronJobFetchOrganizationDetailsOK with default headers values
func NewCronJobFetchOrganizationDetailsOK() *CronJobFetchOrganizationDetailsOK {
	return &CronJobFetchOrganizationDetailsOK{}
}

/*
CronJobFetchOrganizationDetailsOK describes a response with status code 200, with default header values.

Success
*/
type CronJobFetchOrganizationDetailsOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job fetch organization details o k response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job fetch organization details o k response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details o k response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch organization details o k response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch organization details o k response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobFetchOrganizationDetailsOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobFetchOrganizationDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchOrganizationDetailsBadRequest creates a CronJobFetchOrganizationDetailsBadRequest with default headers values
func NewCronJobFetchOrganizationDetailsBadRequest() *CronJobFetchOrganizationDetailsBadRequest {
	return &CronJobFetchOrganizationDetailsBadRequest{}
}

/*
CronJobFetchOrganizationDetailsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobFetchOrganizationDetailsBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this cron job fetch organization details bad request response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch organization details bad request response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details bad request response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch organization details bad request response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch organization details bad request response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobFetchOrganizationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *CronJobFetchOrganizationDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchOrganizationDetailsUnauthorized creates a CronJobFetchOrganizationDetailsUnauthorized with default headers values
func NewCronJobFetchOrganizationDetailsUnauthorized() *CronJobFetchOrganizationDetailsUnauthorized {
	return &CronJobFetchOrganizationDetailsUnauthorized{}
}

/*
CronJobFetchOrganizationDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobFetchOrganizationDetailsUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch organization details unauthorized response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch organization details unauthorized response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details unauthorized response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch organization details unauthorized response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch organization details unauthorized response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobFetchOrganizationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchOrganizationDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchOrganizationDetailsForbidden creates a CronJobFetchOrganizationDetailsForbidden with default headers values
func NewCronJobFetchOrganizationDetailsForbidden() *CronJobFetchOrganizationDetailsForbidden {
	return &CronJobFetchOrganizationDetailsForbidden{}
}

/*
CronJobFetchOrganizationDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobFetchOrganizationDetailsForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch organization details forbidden response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch organization details forbidden response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details forbidden response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch organization details forbidden response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch organization details forbidden response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobFetchOrganizationDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchOrganizationDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchOrganizationDetailsNotFound creates a CronJobFetchOrganizationDetailsNotFound with default headers values
func NewCronJobFetchOrganizationDetailsNotFound() *CronJobFetchOrganizationDetailsNotFound {
	return &CronJobFetchOrganizationDetailsNotFound{}
}

/*
CronJobFetchOrganizationDetailsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobFetchOrganizationDetailsNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this cron job fetch organization details not found response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch organization details not found response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details not found response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch organization details not found response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch organization details not found response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobFetchOrganizationDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchOrganizationDetailsNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *CronJobFetchOrganizationDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchOrganizationDetailsInternalServerError creates a CronJobFetchOrganizationDetailsInternalServerError with default headers values
func NewCronJobFetchOrganizationDetailsInternalServerError() *CronJobFetchOrganizationDetailsInternalServerError {
	return &CronJobFetchOrganizationDetailsInternalServerError{}
}

/*
CronJobFetchOrganizationDetailsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobFetchOrganizationDetailsInternalServerError struct {
}

// IsSuccess returns true when this cron job fetch organization details internal server error response has a 2xx status code
func (o *CronJobFetchOrganizationDetailsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch organization details internal server error response has a 3xx status code
func (o *CronJobFetchOrganizationDetailsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch organization details internal server error response has a 4xx status code
func (o *CronJobFetchOrganizationDetailsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch organization details internal server error response has a 5xx status code
func (o *CronJobFetchOrganizationDetailsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job fetch organization details internal server error response a status code equal to that given
func (o *CronJobFetchOrganizationDetailsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobFetchOrganizationDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsInternalServerError ", 500)
}

func (o *CronJobFetchOrganizationDetailsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-organization-details][%d] cronJobFetchOrganizationDetailsInternalServerError ", 500)
}

func (o *CronJobFetchOrganizationDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
