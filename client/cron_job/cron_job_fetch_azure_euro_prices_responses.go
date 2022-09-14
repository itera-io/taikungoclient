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

// CronJobFetchAzureEuroPricesReader is a Reader for the CronJobFetchAzureEuroPrices structure.
type CronJobFetchAzureEuroPricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CronJobFetchAzureEuroPricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCronJobFetchAzureEuroPricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCronJobFetchAzureEuroPricesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCronJobFetchAzureEuroPricesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCronJobFetchAzureEuroPricesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCronJobFetchAzureEuroPricesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCronJobFetchAzureEuroPricesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCronJobFetchAzureEuroPricesOK creates a CronJobFetchAzureEuroPricesOK with default headers values
func NewCronJobFetchAzureEuroPricesOK() *CronJobFetchAzureEuroPricesOK {
	return &CronJobFetchAzureEuroPricesOK{}
}

/*
CronJobFetchAzureEuroPricesOK describes a response with status code 200, with default header values.

Success
*/
type CronJobFetchAzureEuroPricesOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this cron job fetch azure euro prices o k response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cron job fetch azure euro prices o k response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices o k response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch azure euro prices o k response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch azure euro prices o k response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesOK) IsCode(code int) bool {
	return code == 200
}

func (o *CronJobFetchAzureEuroPricesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesOK  %+v", 200, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *CronJobFetchAzureEuroPricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchAzureEuroPricesBadRequest creates a CronJobFetchAzureEuroPricesBadRequest with default headers values
func NewCronJobFetchAzureEuroPricesBadRequest() *CronJobFetchAzureEuroPricesBadRequest {
	return &CronJobFetchAzureEuroPricesBadRequest{}
}

/*
CronJobFetchAzureEuroPricesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CronJobFetchAzureEuroPricesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this cron job fetch azure euro prices bad request response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch azure euro prices bad request response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices bad request response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch azure euro prices bad request response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch azure euro prices bad request response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CronJobFetchAzureEuroPricesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesBadRequest  %+v", 400, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *CronJobFetchAzureEuroPricesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchAzureEuroPricesUnauthorized creates a CronJobFetchAzureEuroPricesUnauthorized with default headers values
func NewCronJobFetchAzureEuroPricesUnauthorized() *CronJobFetchAzureEuroPricesUnauthorized {
	return &CronJobFetchAzureEuroPricesUnauthorized{}
}

/*
CronJobFetchAzureEuroPricesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CronJobFetchAzureEuroPricesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch azure euro prices unauthorized response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch azure euro prices unauthorized response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices unauthorized response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch azure euro prices unauthorized response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch azure euro prices unauthorized response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CronJobFetchAzureEuroPricesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesUnauthorized  %+v", 401, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchAzureEuroPricesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchAzureEuroPricesForbidden creates a CronJobFetchAzureEuroPricesForbidden with default headers values
func NewCronJobFetchAzureEuroPricesForbidden() *CronJobFetchAzureEuroPricesForbidden {
	return &CronJobFetchAzureEuroPricesForbidden{}
}

/*
CronJobFetchAzureEuroPricesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CronJobFetchAzureEuroPricesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch azure euro prices forbidden response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch azure euro prices forbidden response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices forbidden response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch azure euro prices forbidden response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch azure euro prices forbidden response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CronJobFetchAzureEuroPricesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesForbidden  %+v", 403, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchAzureEuroPricesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchAzureEuroPricesNotFound creates a CronJobFetchAzureEuroPricesNotFound with default headers values
func NewCronJobFetchAzureEuroPricesNotFound() *CronJobFetchAzureEuroPricesNotFound {
	return &CronJobFetchAzureEuroPricesNotFound{}
}

/*
CronJobFetchAzureEuroPricesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CronJobFetchAzureEuroPricesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this cron job fetch azure euro prices not found response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch azure euro prices not found response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices not found response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this cron job fetch azure euro prices not found response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this cron job fetch azure euro prices not found response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CronJobFetchAzureEuroPricesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesNotFound  %+v", 404, o.Payload)
}

func (o *CronJobFetchAzureEuroPricesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *CronJobFetchAzureEuroPricesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCronJobFetchAzureEuroPricesInternalServerError creates a CronJobFetchAzureEuroPricesInternalServerError with default headers values
func NewCronJobFetchAzureEuroPricesInternalServerError() *CronJobFetchAzureEuroPricesInternalServerError {
	return &CronJobFetchAzureEuroPricesInternalServerError{}
}

/*
CronJobFetchAzureEuroPricesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CronJobFetchAzureEuroPricesInternalServerError struct {
}

// IsSuccess returns true when this cron job fetch azure euro prices internal server error response has a 2xx status code
func (o *CronJobFetchAzureEuroPricesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cron job fetch azure euro prices internal server error response has a 3xx status code
func (o *CronJobFetchAzureEuroPricesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cron job fetch azure euro prices internal server error response has a 4xx status code
func (o *CronJobFetchAzureEuroPricesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this cron job fetch azure euro prices internal server error response has a 5xx status code
func (o *CronJobFetchAzureEuroPricesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this cron job fetch azure euro prices internal server error response a status code equal to that given
func (o *CronJobFetchAzureEuroPricesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CronJobFetchAzureEuroPricesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesInternalServerError ", 500)
}

func (o *CronJobFetchAzureEuroPricesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/CronJob/fetch-azure-euro-prices][%d] cronJobFetchAzureEuroPricesInternalServerError ", 500)
}

func (o *CronJobFetchAzureEuroPricesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
