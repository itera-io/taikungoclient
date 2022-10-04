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

// AlertingProfilesDeleteReader is a Reader for the AlertingProfilesDelete structure.
type AlertingProfilesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertingProfilesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertingProfilesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewAlertingProfilesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertingProfilesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAlertingProfilesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertingProfilesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertingProfilesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAlertingProfilesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertingProfilesDeleteOK creates a AlertingProfilesDeleteOK with default headers values
func NewAlertingProfilesDeleteOK() *AlertingProfilesDeleteOK {
	return &AlertingProfilesDeleteOK{}
}

/*
AlertingProfilesDeleteOK describes a response with status code 200, with default header values.

Success
*/
type AlertingProfilesDeleteOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this alerting profiles delete o k response has a 2xx status code
func (o *AlertingProfilesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles delete o k response has a 3xx status code
func (o *AlertingProfilesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete o k response has a 4xx status code
func (o *AlertingProfilesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles delete o k response has a 5xx status code
func (o *AlertingProfilesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete o k response a status code equal to that given
func (o *AlertingProfilesDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *AlertingProfilesDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesDeleteOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteOK  %+v", 200, o.Payload)
}

func (o *AlertingProfilesDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *AlertingProfilesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesDeleteNoContent creates a AlertingProfilesDeleteNoContent with default headers values
func NewAlertingProfilesDeleteNoContent() *AlertingProfilesDeleteNoContent {
	return &AlertingProfilesDeleteNoContent{}
}

/*
AlertingProfilesDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type AlertingProfilesDeleteNoContent struct {
}

// IsSuccess returns true when this alerting profiles delete no content response has a 2xx status code
func (o *AlertingProfilesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alerting profiles delete no content response has a 3xx status code
func (o *AlertingProfilesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete no content response has a 4xx status code
func (o *AlertingProfilesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles delete no content response has a 5xx status code
func (o *AlertingProfilesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete no content response a status code equal to that given
func (o *AlertingProfilesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *AlertingProfilesDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteNoContent ", 204)
}

func (o *AlertingProfilesDeleteNoContent) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteNoContent ", 204)
}

func (o *AlertingProfilesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAlertingProfilesDeleteBadRequest creates a AlertingProfilesDeleteBadRequest with default headers values
func NewAlertingProfilesDeleteBadRequest() *AlertingProfilesDeleteBadRequest {
	return &AlertingProfilesDeleteBadRequest{}
}

/*
AlertingProfilesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type AlertingProfilesDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this alerting profiles delete bad request response has a 2xx status code
func (o *AlertingProfilesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles delete bad request response has a 3xx status code
func (o *AlertingProfilesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete bad request response has a 4xx status code
func (o *AlertingProfilesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles delete bad request response has a 5xx status code
func (o *AlertingProfilesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete bad request response a status code equal to that given
func (o *AlertingProfilesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AlertingProfilesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesDeleteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *AlertingProfilesDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *AlertingProfilesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesDeleteUnauthorized creates a AlertingProfilesDeleteUnauthorized with default headers values
func NewAlertingProfilesDeleteUnauthorized() *AlertingProfilesDeleteUnauthorized {
	return &AlertingProfilesDeleteUnauthorized{}
}

/*
AlertingProfilesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AlertingProfilesDeleteUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this alerting profiles delete unauthorized response has a 2xx status code
func (o *AlertingProfilesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles delete unauthorized response has a 3xx status code
func (o *AlertingProfilesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete unauthorized response has a 4xx status code
func (o *AlertingProfilesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles delete unauthorized response has a 5xx status code
func (o *AlertingProfilesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete unauthorized response a status code equal to that given
func (o *AlertingProfilesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *AlertingProfilesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *AlertingProfilesDeleteUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *AlertingProfilesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesDeleteForbidden creates a AlertingProfilesDeleteForbidden with default headers values
func NewAlertingProfilesDeleteForbidden() *AlertingProfilesDeleteForbidden {
	return &AlertingProfilesDeleteForbidden{}
}

/*
AlertingProfilesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AlertingProfilesDeleteForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this alerting profiles delete forbidden response has a 2xx status code
func (o *AlertingProfilesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles delete forbidden response has a 3xx status code
func (o *AlertingProfilesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete forbidden response has a 4xx status code
func (o *AlertingProfilesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles delete forbidden response has a 5xx status code
func (o *AlertingProfilesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete forbidden response a status code equal to that given
func (o *AlertingProfilesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *AlertingProfilesDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesDeleteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *AlertingProfilesDeleteForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *AlertingProfilesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesDeleteNotFound creates a AlertingProfilesDeleteNotFound with default headers values
func NewAlertingProfilesDeleteNotFound() *AlertingProfilesDeleteNotFound {
	return &AlertingProfilesDeleteNotFound{}
}

/*
AlertingProfilesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AlertingProfilesDeleteNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this alerting profiles delete not found response has a 2xx status code
func (o *AlertingProfilesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles delete not found response has a 3xx status code
func (o *AlertingProfilesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete not found response has a 4xx status code
func (o *AlertingProfilesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this alerting profiles delete not found response has a 5xx status code
func (o *AlertingProfilesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this alerting profiles delete not found response a status code equal to that given
func (o *AlertingProfilesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AlertingProfilesDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesDeleteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *AlertingProfilesDeleteNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *AlertingProfilesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertingProfilesDeleteInternalServerError creates a AlertingProfilesDeleteInternalServerError with default headers values
func NewAlertingProfilesDeleteInternalServerError() *AlertingProfilesDeleteInternalServerError {
	return &AlertingProfilesDeleteInternalServerError{}
}

/*
AlertingProfilesDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type AlertingProfilesDeleteInternalServerError struct {
}

// IsSuccess returns true when this alerting profiles delete internal server error response has a 2xx status code
func (o *AlertingProfilesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this alerting profiles delete internal server error response has a 3xx status code
func (o *AlertingProfilesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alerting profiles delete internal server error response has a 4xx status code
func (o *AlertingProfilesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this alerting profiles delete internal server error response has a 5xx status code
func (o *AlertingProfilesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this alerting profiles delete internal server error response a status code equal to that given
func (o *AlertingProfilesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *AlertingProfilesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteInternalServerError ", 500)
}

func (o *AlertingProfilesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/AlertingProfiles/delete][%d] alertingProfilesDeleteInternalServerError ", 500)
}

func (o *AlertingProfilesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
