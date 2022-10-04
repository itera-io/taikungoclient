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

// AlertingProfilesLockManagerReader is a Reader for the AlertingProfilesLockManager structure.
type AlertingProfilesLockManagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesLockManagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesLockManagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesLockManagerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesLockManagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesLockManagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesLockManagerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesLockManagerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesLockManagerOK creates a AlertingProfilesLockManagerOK with default headers values
func NewAlertingProfilesLockManagerOK() *AlertingProfilesLockManagerOK {
	return &AlertingProfilesLockManagerOK{}
}

/*
AlertingProfilesLockManagerOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesLockManagerOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this alerting profiles lock manager o k response has a 2xx status code
func (o *AlertingProfilesLockManagerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles lock manager o k response has a 3xx status code
func (o *AlertingProfilesLockManagerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager o k response has a 4xx status code
func (o *AlertingProfilesLockManagerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles lock manager o k response has a 5xx status code
func (o *AlertingProfilesLockManagerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles lock manager o k response a status code equal to that given
func (o *AlertingProfilesLockManagerOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesLockManagerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesLockManagerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesLockManagerOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AlertingProfilesLockManagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesLockManagerBadRequest creates a AlertingProfilesLockManagerBadRequest with default headers values
func NewAlertingProfilesLockManagerBadRequest() *AlertingProfilesLockManagerBadRequest {
	return &AlertingProfilesLockManagerBadRequest{}
}

/*
AlertingProfilesLockManagerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesLockManagerBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this alerting profiles lock manager bad request response has a 2xx status code
func (o *AlertingProfilesLockManagerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles lock manager bad request response has a 3xx status code
func (o *AlertingProfilesLockManagerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager bad request response has a 4xx status code
func (o *AlertingProfilesLockManagerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles lock manager bad request response has a 5xx status code
func (o *AlertingProfilesLockManagerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles lock manager bad request response a status code equal to that given
func (o *AlertingProfilesLockManagerBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesLockManagerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesLockManagerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesLockManagerBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesLockManagerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesLockManagerUnauthorized creates a AlertingProfilesLockManagerUnauthorized with default headers values
func NewAlertingProfilesLockManagerUnauthorized() *AlertingProfilesLockManagerUnauthorized {
	return &AlertingProfilesLockManagerUnauthorized{}
}

/*
AlertingProfilesLockManagerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesLockManagerUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting profiles lock manager unauthorized response has a 2xx status code
func (o *AlertingProfilesLockManagerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles lock manager unauthorized response has a 3xx status code
func (o *AlertingProfilesLockManagerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager unauthorized response has a 4xx status code
func (o *AlertingProfilesLockManagerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles lock manager unauthorized response has a 5xx status code
func (o *AlertingProfilesLockManagerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles lock manager unauthorized response a status code equal to that given
func (o *AlertingProfilesLockManagerUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesLockManagerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesLockManagerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesLockManagerUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingProfilesLockManagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesLockManagerForbidden creates a AlertingProfilesLockManagerForbidden with default headers values
func NewAlertingProfilesLockManagerForbidden() *AlertingProfilesLockManagerForbidden {
	return &AlertingProfilesLockManagerForbidden{}
}

/*
AlertingProfilesLockManagerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesLockManagerForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting profiles lock manager forbidden response has a 2xx status code
func (o *AlertingProfilesLockManagerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles lock manager forbidden response has a 3xx status code
func (o *AlertingProfilesLockManagerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager forbidden response has a 4xx status code
func (o *AlertingProfilesLockManagerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles lock manager forbidden response has a 5xx status code
func (o *AlertingProfilesLockManagerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles lock manager forbidden response a status code equal to that given
func (o *AlertingProfilesLockManagerForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesLockManagerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesLockManagerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesLockManagerForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingProfilesLockManagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesLockManagerNotFound creates a AlertingProfilesLockManagerNotFound with default headers values
func NewAlertingProfilesLockManagerNotFound() *AlertingProfilesLockManagerNotFound {
	return &AlertingProfilesLockManagerNotFound{}
}

/*
AlertingProfilesLockManagerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesLockManagerNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this alerting profiles lock manager not found response has a 2xx status code
func (o *AlertingProfilesLockManagerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles lock manager not found response has a 3xx status code
func (o *AlertingProfilesLockManagerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager not found response has a 4xx status code
func (o *AlertingProfilesLockManagerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles lock manager not found response has a 5xx status code
func (o *AlertingProfilesLockManagerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles lock manager not found response a status code equal to that given
func (o *AlertingProfilesLockManagerNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesLockManagerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesLockManagerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesLockManagerNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *AlertingProfilesLockManagerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesLockManagerInternalServerError creates a AlertingProfilesLockManagerInternalServerError with default headers values
func NewAlertingProfilesLockManagerInternalServerError() *AlertingProfilesLockManagerInternalServerError {
	return &AlertingProfilesLockManagerInternalServerError{}
}

/*
AlertingProfilesLockManagerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesLockManagerInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles lock manager internal server error response has a 2xx status code
func (o *AlertingProfilesLockManagerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles lock manager internal server error response has a 3xx status code
func (o *AlertingProfilesLockManagerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles lock manager internal server error response has a 4xx status code
func (o *AlertingProfilesLockManagerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles lock manager internal server error response has a 5xx status code
func (o *AlertingProfilesLockManagerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles lock manager internal server error response a status code equal to that given
func (o *AlertingProfilesLockManagerInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesLockManagerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerInternalServerError ", 500)
}

func (o *AlertingProfilesLockManagerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/lockmanager][%d] alertingProfilesLockManagerInternalServerError ", 500)
}

func (o *AlertingProfilesLockManagerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
