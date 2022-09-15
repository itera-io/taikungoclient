// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// NotificationsGetProjectOperationMessagesReader is a Reader for the NotificationsGetProjectOperationMessages structure.
type NotificationsGetProjectOperationMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsGetProjectOperationMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsGetProjectOperationMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsGetProjectOperationMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsGetProjectOperationMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsGetProjectOperationMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsGetProjectOperationMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsGetProjectOperationMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsGetProjectOperationMessagesOK creates a NotificationsGetProjectOperationMessagesOK with default headers values
func NewNotificationsGetProjectOperationMessagesOK() *NotificationsGetProjectOperationMessagesOK {
	return &NotificationsGetProjectOperationMessagesOK{}
}

/*
NotificationsGetProjectOperationMessagesOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsGetProjectOperationMessagesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this notifications get project operation messages o k response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications get project operation messages o k response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages o k response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications get project operation messages o k response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications get project operation messages o k response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationsGetProjectOperationMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesOK  %+v", 200, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesOK  %+v", 200, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsGetProjectOperationMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsGetProjectOperationMessagesBadRequest creates a NotificationsGetProjectOperationMessagesBadRequest with default headers values
func NewNotificationsGetProjectOperationMessagesBadRequest() *NotificationsGetProjectOperationMessagesBadRequest {
	return &NotificationsGetProjectOperationMessagesBadRequest{}
}

/*
NotificationsGetProjectOperationMessagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsGetProjectOperationMessagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this notifications get project operation messages bad request response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications get project operation messages bad request response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages bad request response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications get project operation messages bad request response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications get project operation messages bad request response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NotificationsGetProjectOperationMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsGetProjectOperationMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsGetProjectOperationMessagesUnauthorized creates a NotificationsGetProjectOperationMessagesUnauthorized with default headers values
func NewNotificationsGetProjectOperationMessagesUnauthorized() *NotificationsGetProjectOperationMessagesUnauthorized {
	return &NotificationsGetProjectOperationMessagesUnauthorized{}
}

/*
NotificationsGetProjectOperationMessagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsGetProjectOperationMessagesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications get project operation messages unauthorized response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications get project operation messages unauthorized response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages unauthorized response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications get project operation messages unauthorized response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications get project operation messages unauthorized response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NotificationsGetProjectOperationMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsGetProjectOperationMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsGetProjectOperationMessagesForbidden creates a NotificationsGetProjectOperationMessagesForbidden with default headers values
func NewNotificationsGetProjectOperationMessagesForbidden() *NotificationsGetProjectOperationMessagesForbidden {
	return &NotificationsGetProjectOperationMessagesForbidden{}
}

/*
NotificationsGetProjectOperationMessagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsGetProjectOperationMessagesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications get project operation messages forbidden response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications get project operation messages forbidden response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages forbidden response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications get project operation messages forbidden response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications get project operation messages forbidden response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NotificationsGetProjectOperationMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsGetProjectOperationMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsGetProjectOperationMessagesNotFound creates a NotificationsGetProjectOperationMessagesNotFound with default headers values
func NewNotificationsGetProjectOperationMessagesNotFound() *NotificationsGetProjectOperationMessagesNotFound {
	return &NotificationsGetProjectOperationMessagesNotFound{}
}

/*
NotificationsGetProjectOperationMessagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsGetProjectOperationMessagesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications get project operation messages not found response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications get project operation messages not found response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages not found response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications get project operation messages not found response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications get project operation messages not found response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NotificationsGetProjectOperationMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsGetProjectOperationMessagesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsGetProjectOperationMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsGetProjectOperationMessagesInternalServerError creates a NotificationsGetProjectOperationMessagesInternalServerError with default headers values
func NewNotificationsGetProjectOperationMessagesInternalServerError() *NotificationsGetProjectOperationMessagesInternalServerError {
	return &NotificationsGetProjectOperationMessagesInternalServerError{}
}

/*
NotificationsGetProjectOperationMessagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsGetProjectOperationMessagesInternalServerError struct {
}

// IsSuccess returns true when this notifications get project operation messages internal server error response has a 2xx status code
func (o *NotificationsGetProjectOperationMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications get project operation messages internal server error response has a 3xx status code
func (o *NotificationsGetProjectOperationMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications get project operation messages internal server error response has a 4xx status code
func (o *NotificationsGetProjectOperationMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications get project operation messages internal server error response has a 5xx status code
func (o *NotificationsGetProjectOperationMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notifications get project operation messages internal server error response a status code equal to that given
func (o *NotificationsGetProjectOperationMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NotificationsGetProjectOperationMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesInternalServerError ", 500)
}

func (o *NotificationsGetProjectOperationMessagesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesInternalServerError ", 500)
}

func (o *NotificationsGetProjectOperationMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
