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

// NotificationsCreateReader is a Reader for the NotificationsCreate structure.
type NotificationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotificationsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewNotificationsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewNotificationsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotificationsCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNotificationsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNotificationsCreateOK creates a NotificationsCreateOK with default headers values
func NewNotificationsCreateOK() *NotificationsCreateOK {
	return &NotificationsCreateOK{}
}

/*
NotificationsCreateOK describes a response with status code 200, with default header values.

Success
*/
type NotificationsCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this notifications create o k response has a 2xx status code
func (o *NotificationsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notifications create o k response has a 3xx status code
func (o *NotificationsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create o k response has a 4xx status code
func (o *NotificationsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications create o k response has a 5xx status code
func (o *NotificationsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications create o k response a status code equal to that given
func (o *NotificationsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *NotificationsCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateOK  %+v", 200, o.Payload)
}

func (o *NotificationsCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateOK  %+v", 200, o.Payload)
}

func (o *NotificationsCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *NotificationsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsCreateBadRequest creates a NotificationsCreateBadRequest with default headers values
func NewNotificationsCreateBadRequest() *NotificationsCreateBadRequest {
	return &NotificationsCreateBadRequest{}
}

/*
NotificationsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type NotificationsCreateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this notifications create bad request response has a 2xx status code
func (o *NotificationsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications create bad request response has a 3xx status code
func (o *NotificationsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create bad request response has a 4xx status code
func (o *NotificationsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications create bad request response has a 5xx status code
func (o *NotificationsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications create bad request response a status code equal to that given
func (o *NotificationsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *NotificationsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationsCreateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *NotificationsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsCreateUnauthorized creates a NotificationsCreateUnauthorized with default headers values
func NewNotificationsCreateUnauthorized() *NotificationsCreateUnauthorized {
	return &NotificationsCreateUnauthorized{}
}

/*
NotificationsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type NotificationsCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications create unauthorized response has a 2xx status code
func (o *NotificationsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications create unauthorized response has a 3xx status code
func (o *NotificationsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create unauthorized response has a 4xx status code
func (o *NotificationsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications create unauthorized response has a 5xx status code
func (o *NotificationsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications create unauthorized response a status code equal to that given
func (o *NotificationsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *NotificationsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationsCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsCreateForbidden creates a NotificationsCreateForbidden with default headers values
func NewNotificationsCreateForbidden() *NotificationsCreateForbidden {
	return &NotificationsCreateForbidden{}
}

/*
NotificationsCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type NotificationsCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications create forbidden response has a 2xx status code
func (o *NotificationsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications create forbidden response has a 3xx status code
func (o *NotificationsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create forbidden response has a 4xx status code
func (o *NotificationsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications create forbidden response has a 5xx status code
func (o *NotificationsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications create forbidden response a status code equal to that given
func (o *NotificationsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *NotificationsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateForbidden  %+v", 403, o.Payload)
}

func (o *NotificationsCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsCreateNotFound creates a NotificationsCreateNotFound with default headers values
func NewNotificationsCreateNotFound() *NotificationsCreateNotFound {
	return &NotificationsCreateNotFound{}
}

/*
NotificationsCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type NotificationsCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this notifications create not found response has a 2xx status code
func (o *NotificationsCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications create not found response has a 3xx status code
func (o *NotificationsCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create not found response has a 4xx status code
func (o *NotificationsCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notifications create not found response has a 5xx status code
func (o *NotificationsCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notifications create not found response a status code equal to that given
func (o *NotificationsCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *NotificationsCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateNotFound  %+v", 404, o.Payload)
}

func (o *NotificationsCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *NotificationsCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationsCreateInternalServerError creates a NotificationsCreateInternalServerError with default headers values
func NewNotificationsCreateInternalServerError() *NotificationsCreateInternalServerError {
	return &NotificationsCreateInternalServerError{}
}

/*
NotificationsCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type NotificationsCreateInternalServerError struct {
}

// IsSuccess returns true when this notifications create internal server error response has a 2xx status code
func (o *NotificationsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notifications create internal server error response has a 3xx status code
func (o *NotificationsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notifications create internal server error response has a 4xx status code
func (o *NotificationsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this notifications create internal server error response has a 5xx status code
func (o *NotificationsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this notifications create internal server error response a status code equal to that given
func (o *NotificationsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *NotificationsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateInternalServerError ", 500)
}

func (o *NotificationsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Notifications/add][%d] notificationsCreateInternalServerError ", 500)
}

func (o *NotificationsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
