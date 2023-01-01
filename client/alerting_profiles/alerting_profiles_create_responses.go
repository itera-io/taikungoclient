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

// AlertingProfilesCreateReader is a Reader for the AlertingProfilesCreate structure.
type AlertingProfilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesCreateOK creates a AlertingProfilesCreateOK with default headers values
func NewAlertingProfilesCreateOK() *AlertingProfilesCreateOK {
	return &AlertingProfilesCreateOK{}
}

/*
AlertingProfilesCreateOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesCreateOK struct {
	Payload *models.APIResponse
}

// IsSuccess returns true when this alerting profiles create o k response has a 2xx status code
func (o *AlertingProfilesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles create o k response has a 3xx status code
func (o *AlertingProfilesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create o k response has a 4xx status code
func (o *AlertingProfilesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles create o k response has a 5xx status code
func (o *AlertingProfilesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create o k response a status code equal to that given
func (o *AlertingProfilesCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesCreateOK) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *AlertingProfilesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateBadRequest creates a AlertingProfilesCreateBadRequest with default headers values
func NewAlertingProfilesCreateBadRequest() *AlertingProfilesCreateBadRequest {
	return &AlertingProfilesCreateBadRequest{}
}

/*
AlertingProfilesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesCreateBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles create bad request response has a 2xx status code
func (o *AlertingProfilesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create bad request response has a 3xx status code
func (o *AlertingProfilesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create bad request response has a 4xx status code
func (o *AlertingProfilesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create bad request response has a 5xx status code
func (o *AlertingProfilesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create bad request response a status code equal to that given
func (o *AlertingProfilesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesCreateBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateUnauthorized creates a AlertingProfilesCreateUnauthorized with default headers values
func NewAlertingProfilesCreateUnauthorized() *AlertingProfilesCreateUnauthorized {
	return &AlertingProfilesCreateUnauthorized{}
}

/*
AlertingProfilesCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles create unauthorized response has a 2xx status code
func (o *AlertingProfilesCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create unauthorized response has a 3xx status code
func (o *AlertingProfilesCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create unauthorized response has a 4xx status code
func (o *AlertingProfilesCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create unauthorized response has a 5xx status code
func (o *AlertingProfilesCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create unauthorized response a status code equal to that given
func (o *AlertingProfilesCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateForbidden creates a AlertingProfilesCreateForbidden with default headers values
func NewAlertingProfilesCreateForbidden() *AlertingProfilesCreateForbidden {
	return &AlertingProfilesCreateForbidden{}
}

/*
AlertingProfilesCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles create forbidden response has a 2xx status code
func (o *AlertingProfilesCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create forbidden response has a 3xx status code
func (o *AlertingProfilesCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create forbidden response has a 4xx status code
func (o *AlertingProfilesCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create forbidden response has a 5xx status code
func (o *AlertingProfilesCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create forbidden response a status code equal to that given
func (o *AlertingProfilesCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateNotFound creates a AlertingProfilesCreateNotFound with default headers values
func NewAlertingProfilesCreateNotFound() *AlertingProfilesCreateNotFound {
	return &AlertingProfilesCreateNotFound{}
}

/*
AlertingProfilesCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles create not found response has a 2xx status code
func (o *AlertingProfilesCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create not found response has a 3xx status code
func (o *AlertingProfilesCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create not found response has a 4xx status code
func (o *AlertingProfilesCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles create not found response has a 5xx status code
func (o *AlertingProfilesCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles create not found response a status code equal to that given
func (o *AlertingProfilesCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesCreateInternalServerError creates a AlertingProfilesCreateInternalServerError with default headers values
func NewAlertingProfilesCreateInternalServerError() *AlertingProfilesCreateInternalServerError {
	return &AlertingProfilesCreateInternalServerError{}
}

/*
AlertingProfilesCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesCreateInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles create internal server error response has a 2xx status code
func (o *AlertingProfilesCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles create internal server error response has a 3xx status code
func (o *AlertingProfilesCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles create internal server error response has a 4xx status code
func (o *AlertingProfilesCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles create internal server error response has a 5xx status code
func (o *AlertingProfilesCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles create internal server error response a status code equal to that given
func (o *AlertingProfilesCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateInternalServerError ", 500)
}

func (o *AlertingProfilesCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/create][%d] alertingProfilesCreateInternalServerError ", 500)
}

func (o *AlertingProfilesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
