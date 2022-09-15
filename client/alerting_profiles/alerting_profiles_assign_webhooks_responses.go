// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// AlertingProfilesAssignWebhooksReader is a Reader for the AlertingProfilesAssignWebhooks structure.
type AlertingProfilesAssignWebhooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesAssignWebhooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesAssignWebhooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesAssignWebhooksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesAssignWebhooksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesAssignWebhooksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesAssignWebhooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesAssignWebhooksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesAssignWebhooksOK creates a AlertingProfilesAssignWebhooksOK with default headers values
func NewAlertingProfilesAssignWebhooksOK() *AlertingProfilesAssignWebhooksOK {
	return &AlertingProfilesAssignWebhooksOK{}
}

/*
AlertingProfilesAssignWebhooksOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesAssignWebhooksOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this alerting profiles assign webhooks o k response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles assign webhooks o k response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks o k response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles assign webhooks o k response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles assign webhooks o k response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesAssignWebhooksOK) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksOK) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AlertingProfilesAssignWebhooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAssignWebhooksBadRequest creates a AlertingProfilesAssignWebhooksBadRequest with default headers values
func NewAlertingProfilesAssignWebhooksBadRequest() *AlertingProfilesAssignWebhooksBadRequest {
	return &AlertingProfilesAssignWebhooksBadRequest{}
}

/*
AlertingProfilesAssignWebhooksBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesAssignWebhooksBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this alerting profiles assign webhooks bad request response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles assign webhooks bad request response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks bad request response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles assign webhooks bad request response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles assign webhooks bad request response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesAssignWebhooksBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAssignWebhooksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAssignWebhooksUnauthorized creates a AlertingProfilesAssignWebhooksUnauthorized with default headers values
func NewAlertingProfilesAssignWebhooksUnauthorized() *AlertingProfilesAssignWebhooksUnauthorized {
	return &AlertingProfilesAssignWebhooksUnauthorized{}
}

/*
AlertingProfilesAssignWebhooksUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesAssignWebhooksUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles assign webhooks unauthorized response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles assign webhooks unauthorized response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks unauthorized response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles assign webhooks unauthorized response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles assign webhooks unauthorized response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesAssignWebhooksUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAssignWebhooksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAssignWebhooksForbidden creates a AlertingProfilesAssignWebhooksForbidden with default headers values
func NewAlertingProfilesAssignWebhooksForbidden() *AlertingProfilesAssignWebhooksForbidden {
	return &AlertingProfilesAssignWebhooksForbidden{}
}

/*
AlertingProfilesAssignWebhooksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesAssignWebhooksForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles assign webhooks forbidden response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles assign webhooks forbidden response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks forbidden response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles assign webhooks forbidden response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles assign webhooks forbidden response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesAssignWebhooksForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAssignWebhooksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAssignWebhooksNotFound creates a AlertingProfilesAssignWebhooksNotFound with default headers values
func NewAlertingProfilesAssignWebhooksNotFound() *AlertingProfilesAssignWebhooksNotFound {
	return &AlertingProfilesAssignWebhooksNotFound{}
}

/*
AlertingProfilesAssignWebhooksNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesAssignWebhooksNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles assign webhooks not found response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles assign webhooks not found response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks not found response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles assign webhooks not found response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles assign webhooks not found response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesAssignWebhooksNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesAssignWebhooksNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesAssignWebhooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesAssignWebhooksInternalServerError creates a AlertingProfilesAssignWebhooksInternalServerError with default headers values
func NewAlertingProfilesAssignWebhooksInternalServerError() *AlertingProfilesAssignWebhooksInternalServerError {
	return &AlertingProfilesAssignWebhooksInternalServerError{}
}

/*
AlertingProfilesAssignWebhooksInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesAssignWebhooksInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles assign webhooks internal server error response has a 2xx status code
func (o *AlertingProfilesAssignWebhooksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles assign webhooks internal server error response has a 3xx status code
func (o *AlertingProfilesAssignWebhooksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles assign webhooks internal server error response has a 4xx status code
func (o *AlertingProfilesAssignWebhooksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles assign webhooks internal server error response has a 5xx status code
func (o *AlertingProfilesAssignWebhooksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles assign webhooks internal server error response a status code equal to that given
func (o *AlertingProfilesAssignWebhooksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesAssignWebhooksInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksInternalServerError ", 500)
}

func (o *AlertingProfilesAssignWebhooksInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v{v}/AlertingProfiles/assignwebhooks/{id}][%d] alertingProfilesAssignWebhooksInternalServerError ", 500)
}

func (o *AlertingProfilesAssignWebhooksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
