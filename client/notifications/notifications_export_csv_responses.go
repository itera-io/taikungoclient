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

// NotificationsExportCsvReader is a Reader for the NotificationsExportCsv structure.
type NotificationsExportCsvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsExportCsvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsExportCsvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsExportCsvBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsExportCsvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsExportCsvForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsExportCsvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsExportCsvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsExportCsvOK creates a NotificationsExportCsvOK with default headers values
func NewNotificationsExportCsvOK() *NotificationsExportCsvOK {
	return &NotificationsExportCsvOK{}
}

/* NotificationsExportCsvOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsExportCsvOK struct {
}

func (o *NotificationsExportCsvOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvOK ", 200)
}

func (o *NotificationsExportCsvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsExportCsvBadRequest creates a NotificationsExportCsvBadRequest with default headers values
func NewNotificationsExportCsvBadRequest() *NotificationsExportCsvBadRequest {
	return &NotificationsExportCsvBadRequest{}
}

/* NotificationsExportCsvBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsExportCsvBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *NotificationsExportCsvBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvBadRequest  %+v", 400, o.Payload)
}
func (o *NotificationsExportCsvBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *NotificationsExportCsvBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsExportCsvUnauthorized creates a NotificationsExportCsvUnauthorized with default headers values
func NewNotificationsExportCsvUnauthorized() *NotificationsExportCsvUnauthorized {
	return &NotificationsExportCsvUnauthorized{}
}

/* NotificationsExportCsvUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsExportCsvUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsExportCsvUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvUnauthorized  %+v", 401, o.Payload)
}
func (o *NotificationsExportCsvUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsExportCsvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsExportCsvForbidden creates a NotificationsExportCsvForbidden with default headers values
func NewNotificationsExportCsvForbidden() *NotificationsExportCsvForbidden {
	return &NotificationsExportCsvForbidden{}
}

/* NotificationsExportCsvForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsExportCsvForbidden struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsExportCsvForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvForbidden  %+v", 403, o.Payload)
}
func (o *NotificationsExportCsvForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsExportCsvForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsExportCsvNotFound creates a NotificationsExportCsvNotFound with default headers values
func NewNotificationsExportCsvNotFound() *NotificationsExportCsvNotFound {
	return &NotificationsExportCsvNotFound{}
}

/* NotificationsExportCsvNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsExportCsvNotFound struct {
	Payload *models.ProblemDetails
}

func (o *NotificationsExportCsvNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvNotFound  %+v", 404, o.Payload)
}
func (o *NotificationsExportCsvNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsExportCsvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsExportCsvInternalServerError creates a NotificationsExportCsvInternalServerError with default headers values
func NewNotificationsExportCsvInternalServerError() *NotificationsExportCsvInternalServerError {
	return &NotificationsExportCsvInternalServerError{}
}

/* NotificationsExportCsvInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsExportCsvInternalServerError struct {
}

func (o *NotificationsExportCsvInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications/download][%d] notificationsExportCsvInternalServerError ", 500)
}

func (o *NotificationsExportCsvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
