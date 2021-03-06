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

// NotificationsListReader is a Reader for the NotificationsList structure.
type NotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsListOK creates a NotificationsListOK with default headers values
func NewNotificationsListOK() *NotificationsListOK {
	return &NotificationsListOK{}
}

/* NotificationsListOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsListOK struct {
	Payload *models.NotificationHistory
}

func (o *NotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListOK  %+v", 200, o.Payload)
}
func (o *NotificationsListOK) GetPayload() *models.NotificationHistory {
	return o.Payload
}

func (o *NotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListBadRequest creates a NotificationsListBadRequest with default headers values
func NewNotificationsListBadRequest() *NotificationsListBadRequest {
	return &NotificationsListBadRequest{}
}

/* NotificationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListUnauthorized creates a NotificationsListUnauthorized with default headers values
func NewNotificationsListUnauthorized() *NotificationsListUnauthorized {
	return &NotificationsListUnauthorized{}
}

/* NotificationsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListForbidden creates a NotificationsListForbidden with default headers values
func NewNotificationsListForbidden() *NotificationsListForbidden {
	return &NotificationsListForbidden{}
}

/* NotificationsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListForbidden  %+v", 403, o.Payload)
}
func (o *NotificationsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListNotFound creates a NotificationsListNotFound with default headers values
func NewNotificationsListNotFound() *NotificationsListNotFound {
	return &NotificationsListNotFound{}
}

/* NotificationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListNotFound  %+v", 404, o.Payload)
}
func (o *NotificationsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsListInternalServerError creates a NotificationsListInternalServerError with default headers values
func NewNotificationsListInternalServerError() *NotificationsListInternalServerError {
	return &NotificationsListInternalServerError{}
}

/* NotificationsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsListInternalServerError struct {
}

func (o *NotificationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListInternalServerError ", 500)
}

func (o *NotificationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
