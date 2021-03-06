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

/* NotificationsGetProjectOperationMessagesOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsGetProjectOperationMessagesOK struct {
	Payload interface{}
}

func (o *NotificationsGetProjectOperationMessagesOK) Error() string {
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

/* NotificationsGetProjectOperationMessagesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsGetProjectOperationMessagesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsGetProjectOperationMessagesBadRequest) Error() string {
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

/* NotificationsGetProjectOperationMessagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsGetProjectOperationMessagesUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsGetProjectOperationMessagesUnauthorized) Error() string {
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

/* NotificationsGetProjectOperationMessagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsGetProjectOperationMessagesForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsGetProjectOperationMessagesForbidden) Error() string {
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

/* NotificationsGetProjectOperationMessagesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsGetProjectOperationMessagesNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsGetProjectOperationMessagesNotFound) Error() string {
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

/* NotificationsGetProjectOperationMessagesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsGetProjectOperationMessagesInternalServerError struct {
}

func (o *NotificationsGetProjectOperationMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/operations][%d] notificationsGetProjectOperationMessagesInternalServerError ", 500)
}

func (o *NotificationsGetProjectOperationMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
