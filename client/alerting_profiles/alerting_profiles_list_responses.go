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

// AlertingProfilesListReader is a Reader for the AlertingProfilesList structure.
type AlertingProfilesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesListOK creates a AlertingProfilesListOK with default headers values
func NewAlertingProfilesListOK() *AlertingProfilesListOK {
	return &AlertingProfilesListOK{}
}

/*
AlertingProfilesListOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesListOK struct {
	Payload *models.AlertingProfilesList
}

// IsSuccess returns true when this alerting profiles list o k response has a 2xx status code
func (o *AlertingProfilesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles list o k response has a 3xx status code
func (o *AlertingProfilesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list o k response has a 4xx status code
func (o *AlertingProfilesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles list o k response has a 5xx status code
func (o *AlertingProfilesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles list o k response a status code equal to that given
func (o *AlertingProfilesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesListOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesListOK) GetPayload() *models.AlertingProfilesList {
	return o.Payload
}

func (o *AlertingProfilesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingProfilesList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesListBadRequest creates a AlertingProfilesListBadRequest with default headers values
func NewAlertingProfilesListBadRequest() *AlertingProfilesListBadRequest {
	return &AlertingProfilesListBadRequest{}
}

/*
AlertingProfilesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesListBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles list bad request response has a 2xx status code
func (o *AlertingProfilesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles list bad request response has a 3xx status code
func (o *AlertingProfilesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list bad request response has a 4xx status code
func (o *AlertingProfilesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles list bad request response has a 5xx status code
func (o *AlertingProfilesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles list bad request response a status code equal to that given
func (o *AlertingProfilesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesListBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesListBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesListUnauthorized creates a AlertingProfilesListUnauthorized with default headers values
func NewAlertingProfilesListUnauthorized() *AlertingProfilesListUnauthorized {
	return &AlertingProfilesListUnauthorized{}
}

/*
AlertingProfilesListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles list unauthorized response has a 2xx status code
func (o *AlertingProfilesListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles list unauthorized response has a 3xx status code
func (o *AlertingProfilesListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list unauthorized response has a 4xx status code
func (o *AlertingProfilesListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles list unauthorized response has a 5xx status code
func (o *AlertingProfilesListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles list unauthorized response a status code equal to that given
func (o *AlertingProfilesListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesListUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesListForbidden creates a AlertingProfilesListForbidden with default headers values
func NewAlertingProfilesListForbidden() *AlertingProfilesListForbidden {
	return &AlertingProfilesListForbidden{}
}

/*
AlertingProfilesListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles list forbidden response has a 2xx status code
func (o *AlertingProfilesListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles list forbidden response has a 3xx status code
func (o *AlertingProfilesListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list forbidden response has a 4xx status code
func (o *AlertingProfilesListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles list forbidden response has a 5xx status code
func (o *AlertingProfilesListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles list forbidden response a status code equal to that given
func (o *AlertingProfilesListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesListForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesListNotFound creates a AlertingProfilesListNotFound with default headers values
func NewAlertingProfilesListNotFound() *AlertingProfilesListNotFound {
	return &AlertingProfilesListNotFound{}
}

/*
AlertingProfilesListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this alerting profiles list not found response has a 2xx status code
func (o *AlertingProfilesListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles list not found response has a 3xx status code
func (o *AlertingProfilesListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list not found response has a 4xx status code
func (o *AlertingProfilesListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles list not found response has a 5xx status code
func (o *AlertingProfilesListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles list not found response a status code equal to that given
func (o *AlertingProfilesListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesListNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesListInternalServerError creates a AlertingProfilesListInternalServerError with default headers values
func NewAlertingProfilesListInternalServerError() *AlertingProfilesListInternalServerError {
	return &AlertingProfilesListInternalServerError{}
}

/*
AlertingProfilesListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesListInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles list internal server error response has a 2xx status code
func (o *AlertingProfilesListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles list internal server error response has a 3xx status code
func (o *AlertingProfilesListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles list internal server error response has a 4xx status code
func (o *AlertingProfilesListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles list internal server error response has a 5xx status code
func (o *AlertingProfilesListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles list internal server error response a status code equal to that given
func (o *AlertingProfilesListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListInternalServerError ", 500)
}

func (o *AlertingProfilesListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/AlertingProfiles][%d] alertingProfilesListInternalServerError ", 500)
}

func (o *AlertingProfilesListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
