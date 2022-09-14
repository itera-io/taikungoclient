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

/*
NotificationsListOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsListOK struct {
	Payload *models.NotificationHistory
}

// IsSuccess returns true when this notifications list o k response has a 2xx status code
func (o *NotificationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications list o k response has a 3xx status code
func (o *NotificationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list o k response has a 4xx status code
func (o *NotificationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications list o k response has a 5xx status code
func (o *NotificationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications list o k response a status code equal to that given
func (o *NotificationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListOK  %+v", 200, o.Payload)
}

func (o *NotificationsListOK) String() string {
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

/*
NotificationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this notifications list bad request response has a 2xx status code
func (o *NotificationsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications list bad request response has a 3xx status code
func (o *NotificationsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list bad request response has a 4xx status code
func (o *NotificationsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications list bad request response has a 5xx status code
func (o *NotificationsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications list bad request response a status code equal to that given
func (o *NotificationsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NotificationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsListBadRequest) String() string {
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

/*
NotificationsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsListUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications list unauthorized response has a 2xx status code
func (o *NotificationsListUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications list unauthorized response has a 3xx status code
func (o *NotificationsListUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list unauthorized response has a 4xx status code
func (o *NotificationsListUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications list unauthorized response has a 5xx status code
func (o *NotificationsListUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications list unauthorized response a status code equal to that given
func (o *NotificationsListUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NotificationsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsListUnauthorized) String() string {
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

/*
NotificationsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsListForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications list forbidden response has a 2xx status code
func (o *NotificationsListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications list forbidden response has a 3xx status code
func (o *NotificationsListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list forbidden response has a 4xx status code
func (o *NotificationsListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications list forbidden response has a 5xx status code
func (o *NotificationsListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications list forbidden response a status code equal to that given
func (o *NotificationsListForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NotificationsListForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsListForbidden) String() string {
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

/*
NotificationsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsListNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications list not found response has a 2xx status code
func (o *NotificationsListNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications list not found response has a 3xx status code
func (o *NotificationsListNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list not found response has a 4xx status code
func (o *NotificationsListNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications list not found response has a 5xx status code
func (o *NotificationsListNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications list not found response a status code equal to that given
func (o *NotificationsListNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NotificationsListNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsListNotFound) String() string {
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

/*
NotificationsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsListInternalServerError struct {
}

// IsSuccess returns true when this notifications list internal server error response has a 2xx status code
func (o *NotificationsListInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications list internal server error response has a 3xx status code
func (o *NotificationsListInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications list internal server error response has a 4xx status code
func (o *NotificationsListInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications list internal server error response has a 5xx status code
func (o *NotificationsListInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notifications list internal server error response a status code equal to that given
func (o *NotificationsListInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NotificationsListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListInternalServerError ", 500)
}

func (o *NotificationsListInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Notifications][%d] notificationsListInternalServerError ", 500)
}

func (o *NotificationsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
