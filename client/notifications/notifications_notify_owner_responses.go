// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NotificationsNotifyOwnerReader is a Reader for the NotificationsNotifyOwner structure.
type NotificationsNotifyOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsNotifyOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsNotifyOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsNotifyOwnerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsNotifyOwnerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsNotifyOwnerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsNotifyOwnerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsNotifyOwnerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsNotifyOwnerOK creates a NotificationsNotifyOwnerOK with default headers values
func NewNotificationsNotifyOwnerOK() *NotificationsNotifyOwnerOK {
	return &NotificationsNotifyOwnerOK{}
}

/*
NotificationsNotifyOwnerOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsNotifyOwnerOK struct {
}

// IsSuccess returns true when this notifications notify owner o k response has a 2xx status code
func (o *NotificationsNotifyOwnerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications notify owner o k response has a 3xx status code
func (o *NotificationsNotifyOwnerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner o k response has a 4xx status code
func (o *NotificationsNotifyOwnerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications notify owner o k response has a 5xx status code
func (o *NotificationsNotifyOwnerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications notify owner o k response a status code equal to that given
func (o *NotificationsNotifyOwnerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notifications notify owner o k response
func (o *NotificationsNotifyOwnerOK) Code() int {
	return 200
}

func (o *NotificationsNotifyOwnerOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerOK ", 200)
}

func (o *NotificationsNotifyOwnerOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerOK ", 200)
}

func (o *NotificationsNotifyOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationsNotifyOwnerBadRequest creates a NotificationsNotifyOwnerBadRequest with default headers values
func NewNotificationsNotifyOwnerBadRequest() *NotificationsNotifyOwnerBadRequest {
	return &NotificationsNotifyOwnerBadRequest{}
}

/*
NotificationsNotifyOwnerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsNotifyOwnerBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this notifications notify owner bad request response has a 2xx status code
func (o *NotificationsNotifyOwnerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications notify owner bad request response has a 3xx status code
func (o *NotificationsNotifyOwnerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner bad request response has a 4xx status code
func (o *NotificationsNotifyOwnerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications notify owner bad request response has a 5xx status code
func (o *NotificationsNotifyOwnerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications notify owner bad request response a status code equal to that given
func (o *NotificationsNotifyOwnerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the notifications notify owner bad request response
func (o *NotificationsNotifyOwnerBadRequest) Code() int {
	return 400
}

func (o *NotificationsNotifyOwnerBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsNotifyOwnerBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsNotifyOwnerBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsNotifyOwnerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsNotifyOwnerUnauthorized creates a NotificationsNotifyOwnerUnauthorized with default headers values
func NewNotificationsNotifyOwnerUnauthorized() *NotificationsNotifyOwnerUnauthorized {
	return &NotificationsNotifyOwnerUnauthorized{}
}

/*
NotificationsNotifyOwnerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsNotifyOwnerUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this notifications notify owner unauthorized response has a 2xx status code
func (o *NotificationsNotifyOwnerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications notify owner unauthorized response has a 3xx status code
func (o *NotificationsNotifyOwnerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner unauthorized response has a 4xx status code
func (o *NotificationsNotifyOwnerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications notify owner unauthorized response has a 5xx status code
func (o *NotificationsNotifyOwnerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications notify owner unauthorized response a status code equal to that given
func (o *NotificationsNotifyOwnerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the notifications notify owner unauthorized response
func (o *NotificationsNotifyOwnerUnauthorized) Code() int {
	return 401
}

func (o *NotificationsNotifyOwnerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsNotifyOwnerUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsNotifyOwnerUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsNotifyOwnerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsNotifyOwnerForbidden creates a NotificationsNotifyOwnerForbidden with default headers values
func NewNotificationsNotifyOwnerForbidden() *NotificationsNotifyOwnerForbidden {
	return &NotificationsNotifyOwnerForbidden{}
}

/*
NotificationsNotifyOwnerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsNotifyOwnerForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this notifications notify owner forbidden response has a 2xx status code
func (o *NotificationsNotifyOwnerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications notify owner forbidden response has a 3xx status code
func (o *NotificationsNotifyOwnerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner forbidden response has a 4xx status code
func (o *NotificationsNotifyOwnerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications notify owner forbidden response has a 5xx status code
func (o *NotificationsNotifyOwnerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications notify owner forbidden response a status code equal to that given
func (o *NotificationsNotifyOwnerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notifications notify owner forbidden response
func (o *NotificationsNotifyOwnerForbidden) Code() int {
	return 403
}

func (o *NotificationsNotifyOwnerForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsNotifyOwnerForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsNotifyOwnerForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsNotifyOwnerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsNotifyOwnerNotFound creates a NotificationsNotifyOwnerNotFound with default headers values
func NewNotificationsNotifyOwnerNotFound() *NotificationsNotifyOwnerNotFound {
	return &NotificationsNotifyOwnerNotFound{}
}

/*
NotificationsNotifyOwnerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsNotifyOwnerNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this notifications notify owner not found response has a 2xx status code
func (o *NotificationsNotifyOwnerNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications notify owner not found response has a 3xx status code
func (o *NotificationsNotifyOwnerNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner not found response has a 4xx status code
func (o *NotificationsNotifyOwnerNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications notify owner not found response has a 5xx status code
func (o *NotificationsNotifyOwnerNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications notify owner not found response a status code equal to that given
func (o *NotificationsNotifyOwnerNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notifications notify owner not found response
func (o *NotificationsNotifyOwnerNotFound) Code() int {
	return 404
}

func (o *NotificationsNotifyOwnerNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsNotifyOwnerNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsNotifyOwnerNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *NotificationsNotifyOwnerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsNotifyOwnerInternalServerError creates a NotificationsNotifyOwnerInternalServerError with default headers values
func NewNotificationsNotifyOwnerInternalServerError() *NotificationsNotifyOwnerInternalServerError {
	return &NotificationsNotifyOwnerInternalServerError{}
}

/*
NotificationsNotifyOwnerInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsNotifyOwnerInternalServerError struct {
}

// IsSuccess returns true when this notifications notify owner internal server error response has a 2xx status code
func (o *NotificationsNotifyOwnerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications notify owner internal server error response has a 3xx status code
func (o *NotificationsNotifyOwnerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications notify owner internal server error response has a 4xx status code
func (o *NotificationsNotifyOwnerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications notify owner internal server error response has a 5xx status code
func (o *NotificationsNotifyOwnerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notifications notify owner internal server error response a status code equal to that given
func (o *NotificationsNotifyOwnerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the notifications notify owner internal server error response
func (o *NotificationsNotifyOwnerInternalServerError) Code() int {
	return 500
}

func (o *NotificationsNotifyOwnerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerInternalServerError ", 500)
}

func (o *NotificationsNotifyOwnerInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/notifyowner][%d] notificationsNotifyOwnerInternalServerError ", 500)
}

func (o *NotificationsNotifyOwnerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
