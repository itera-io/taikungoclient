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

// NotificationsListByUserReader is a Reader for the NotificationsListByUser structure.
type NotificationsListByUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsListByUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsListByUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsListByUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsListByUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsListByUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsListByUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsListByUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsListByUserOK creates a NotificationsListByUserOK with default headers values
func NewNotificationsListByUserOK() *NotificationsListByUserOK {
	return &NotificationsListByUserOK{}
}

/* NotificationsListByUserOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsListByUserOK struct {
	Payload *models.ListUserNotifications
}

func (o *NotificationsListByUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserOK  %+v", 200, o.Payload)
}
func (o *NotificationsListByUserOK) GetPayload() *models.ListUserNotifications {
	return o.Payload
}

func (o *NotificationsListByUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListUserNotifications)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListByUserBadRequest creates a NotificationsListByUserBadRequest with default headers values
func NewNotificationsListByUserBadRequest() *NotificationsListByUserBadRequest {
	return &NotificationsListByUserBadRequest{}
}

/* NotificationsListByUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsListByUserBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsListByUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationsListByUserBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsListByUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListByUserUnauthorized creates a NotificationsListByUserUnauthorized with default headers values
func NewNotificationsListByUserUnauthorized() *NotificationsListByUserUnauthorized {
	return &NotificationsListByUserUnauthorized{}
}

/* NotificationsListByUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsListByUserUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListByUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationsListByUserUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListByUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListByUserForbidden creates a NotificationsListByUserForbidden with default headers values
func NewNotificationsListByUserForbidden() *NotificationsListByUserForbidden {
	return &NotificationsListByUserForbidden{}
}

/* NotificationsListByUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsListByUserForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListByUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserForbidden  %+v", 403, o.Payload)
}
func (o *NotificationsListByUserForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListByUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListByUserNotFound creates a NotificationsListByUserNotFound with default headers values
func NewNotificationsListByUserNotFound() *NotificationsListByUserNotFound {
	return &NotificationsListByUserNotFound{}
}

/* NotificationsListByUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsListByUserNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListByUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserNotFound  %+v", 404, o.Payload)
}
func (o *NotificationsListByUserNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListByUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListByUserInternalServerError creates a NotificationsListByUserInternalServerError with default headers values
func NewNotificationsListByUserInternalServerError() *NotificationsListByUserInternalServerError {
	return &NotificationsListByUserInternalServerError{}
}

/* NotificationsListByUserInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsListByUserInternalServerError struct {
}

func (o *NotificationsListByUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/list][%d] notificationsListByUserInternalServerError ", 500)
}

func (o *NotificationsListByUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
