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

// NotificationsDeleteReader is a Reader for the NotificationsDelete structure.
type NotificationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewNotificationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsDeleteOK creates a NotificationsDeleteOK with default headers values
func NewNotificationsDeleteOK() *NotificationsDeleteOK {
	return &NotificationsDeleteOK{}
}

/* NotificationsDeleteOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsDeleteOK struct {
	Payload models.Unit
}

func (o *NotificationsDeleteOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteOK  %+v", 200, o.Payload)
}
func (o *NotificationsDeleteOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *NotificationsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsDeleteNoContent creates a NotificationsDeleteNoContent with default headers values
func NewNotificationsDeleteNoContent() *NotificationsDeleteNoContent {
	return &NotificationsDeleteNoContent{}
}

/* NotificationsDeleteNoContent describes a response with status code 204, with default header values.

Success
*/
type NotificationsDeleteNoContent struct {
}

func (o *NotificationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteNoContent ", 204)
}

func (o *NotificationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsDeleteBadRequest creates a NotificationsDeleteBadRequest with default headers values
func NewNotificationsDeleteBadRequest() *NotificationsDeleteBadRequest {
	return &NotificationsDeleteBadRequest{}
}

/* NotificationsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsDeleteBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationsDeleteBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsDeleteUnauthorized creates a NotificationsDeleteUnauthorized with default headers values
func NewNotificationsDeleteUnauthorized() *NotificationsDeleteUnauthorized {
	return &NotificationsDeleteUnauthorized{}
}

/* NotificationsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsDeleteUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationsDeleteUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsDeleteForbidden creates a NotificationsDeleteForbidden with default headers values
func NewNotificationsDeleteForbidden() *NotificationsDeleteForbidden {
	return &NotificationsDeleteForbidden{}
}

/* NotificationsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsDeleteForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsDeleteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteForbidden  %+v", 403, o.Payload)
}
func (o *NotificationsDeleteForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsDeleteNotFound creates a NotificationsDeleteNotFound with default headers values
func NewNotificationsDeleteNotFound() *NotificationsDeleteNotFound {
	return &NotificationsDeleteNotFound{}
}

/* NotificationsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsDeleteNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsDeleteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteNotFound  %+v", 404, o.Payload)
}
func (o *NotificationsDeleteNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsDeleteInternalServerError creates a NotificationsDeleteInternalServerError with default headers values
func NewNotificationsDeleteInternalServerError() *NotificationsDeleteInternalServerError {
	return &NotificationsDeleteInternalServerError{}
}

/* NotificationsDeleteInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsDeleteInternalServerError struct {
}

func (o *NotificationsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/delete][%d] notificationsDeleteInternalServerError ", 500)
}

func (o *NotificationsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}