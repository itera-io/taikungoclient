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

// NotificationsUpdateReader is a Reader for the NotificationsUpdate structure.
type NotificationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsUpdateOK creates a NotificationsUpdateOK with default headers values
func NewNotificationsUpdateOK() *NotificationsUpdateOK {
	return &NotificationsUpdateOK{}
}

/* NotificationsUpdateOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsUpdateOK struct {
	Payload models.Unit
}

func (o *NotificationsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateOK  %+v", 200, o.Payload)
}
func (o *NotificationsUpdateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *NotificationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsUpdateBadRequest creates a NotificationsUpdateBadRequest with default headers values
func NewNotificationsUpdateBadRequest() *NotificationsUpdateBadRequest {
	return &NotificationsUpdateBadRequest{}
}

/* NotificationsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsUpdateBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationsUpdateBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsUpdateUnauthorized creates a NotificationsUpdateUnauthorized with default headers values
func NewNotificationsUpdateUnauthorized() *NotificationsUpdateUnauthorized {
	return &NotificationsUpdateUnauthorized{}
}

/* NotificationsUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsUpdateUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationsUpdateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsUpdateForbidden creates a NotificationsUpdateForbidden with default headers values
func NewNotificationsUpdateForbidden() *NotificationsUpdateForbidden {
	return &NotificationsUpdateForbidden{}
}

/* NotificationsUpdateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsUpdateForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateForbidden  %+v", 403, o.Payload)
}
func (o *NotificationsUpdateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsUpdateNotFound creates a NotificationsUpdateNotFound with default headers values
func NewNotificationsUpdateNotFound() *NotificationsUpdateNotFound {
	return &NotificationsUpdateNotFound{}
}

/* NotificationsUpdateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsUpdateNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateNotFound  %+v", 404, o.Payload)
}
func (o *NotificationsUpdateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsUpdateInternalServerError creates a NotificationsUpdateInternalServerError with default headers values
func NewNotificationsUpdateInternalServerError() *NotificationsUpdateInternalServerError {
	return &NotificationsUpdateInternalServerError{}
}

/* NotificationsUpdateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsUpdateInternalServerError struct {
}

func (o *NotificationsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/update][%d] notificationsUpdateInternalServerError ", 500)
}

func (o *NotificationsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}