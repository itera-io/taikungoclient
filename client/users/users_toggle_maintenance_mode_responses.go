// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// UsersToggleMaintenanceModeReader is a Reader for the UsersToggleMaintenanceMode structure.
type UsersToggleMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersToggleMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersToggleMaintenanceModeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersToggleMaintenanceModeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUsersToggleMaintenanceModeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUsersToggleMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUsersToggleMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUsersToggleMaintenanceModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersToggleMaintenanceModeOK creates a UsersToggleMaintenanceModeOK with default headers values
func NewUsersToggleMaintenanceModeOK() *UsersToggleMaintenanceModeOK {
	return &UsersToggleMaintenanceModeOK{}
}

/*
UsersToggleMaintenanceModeOK describes a response with status code 200, with default header values.

Success
*/
type UsersToggleMaintenanceModeOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this users toggle maintenance mode o k response has a 2xx status code
func (o *UsersToggleMaintenanceModeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users toggle maintenance mode o k response has a 3xx status code
func (o *UsersToggleMaintenanceModeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode o k response has a 4xx status code
func (o *UsersToggleMaintenanceModeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users toggle maintenance mode o k response has a 5xx status code
func (o *UsersToggleMaintenanceModeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle maintenance mode o k response a status code equal to that given
func (o *UsersToggleMaintenanceModeOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersToggleMaintenanceModeOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeOK  %+v", 200, o.Payload)
}

func (o *UsersToggleMaintenanceModeOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeOK  %+v", 200, o.Payload)
}

func (o *UsersToggleMaintenanceModeOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *UsersToggleMaintenanceModeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleMaintenanceModeBadRequest creates a UsersToggleMaintenanceModeBadRequest with default headers values
func NewUsersToggleMaintenanceModeBadRequest() *UsersToggleMaintenanceModeBadRequest {
	return &UsersToggleMaintenanceModeBadRequest{}
}

/*
UsersToggleMaintenanceModeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersToggleMaintenanceModeBadRequest struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users toggle maintenance mode bad request response has a 2xx status code
func (o *UsersToggleMaintenanceModeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle maintenance mode bad request response has a 3xx status code
func (o *UsersToggleMaintenanceModeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode bad request response has a 4xx status code
func (o *UsersToggleMaintenanceModeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle maintenance mode bad request response has a 5xx status code
func (o *UsersToggleMaintenanceModeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle maintenance mode bad request response a status code equal to that given
func (o *UsersToggleMaintenanceModeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersToggleMaintenanceModeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeBadRequest  %+v", 400, o.Payload)
}

func (o *UsersToggleMaintenanceModeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeBadRequest  %+v", 400, o.Payload)
}

func (o *UsersToggleMaintenanceModeBadRequest) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersToggleMaintenanceModeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleMaintenanceModeUnauthorized creates a UsersToggleMaintenanceModeUnauthorized with default headers values
func NewUsersToggleMaintenanceModeUnauthorized() *UsersToggleMaintenanceModeUnauthorized {
	return &UsersToggleMaintenanceModeUnauthorized{}
}

/*
UsersToggleMaintenanceModeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UsersToggleMaintenanceModeUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users toggle maintenance mode unauthorized response has a 2xx status code
func (o *UsersToggleMaintenanceModeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle maintenance mode unauthorized response has a 3xx status code
func (o *UsersToggleMaintenanceModeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode unauthorized response has a 4xx status code
func (o *UsersToggleMaintenanceModeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle maintenance mode unauthorized response has a 5xx status code
func (o *UsersToggleMaintenanceModeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle maintenance mode unauthorized response a status code equal to that given
func (o *UsersToggleMaintenanceModeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UsersToggleMaintenanceModeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersToggleMaintenanceModeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeUnauthorized  %+v", 401, o.Payload)
}

func (o *UsersToggleMaintenanceModeUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersToggleMaintenanceModeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleMaintenanceModeForbidden creates a UsersToggleMaintenanceModeForbidden with default headers values
func NewUsersToggleMaintenanceModeForbidden() *UsersToggleMaintenanceModeForbidden {
	return &UsersToggleMaintenanceModeForbidden{}
}

/*
UsersToggleMaintenanceModeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UsersToggleMaintenanceModeForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users toggle maintenance mode forbidden response has a 2xx status code
func (o *UsersToggleMaintenanceModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle maintenance mode forbidden response has a 3xx status code
func (o *UsersToggleMaintenanceModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode forbidden response has a 4xx status code
func (o *UsersToggleMaintenanceModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle maintenance mode forbidden response has a 5xx status code
func (o *UsersToggleMaintenanceModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle maintenance mode forbidden response a status code equal to that given
func (o *UsersToggleMaintenanceModeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UsersToggleMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *UsersToggleMaintenanceModeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *UsersToggleMaintenanceModeForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersToggleMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleMaintenanceModeNotFound creates a UsersToggleMaintenanceModeNotFound with default headers values
func NewUsersToggleMaintenanceModeNotFound() *UsersToggleMaintenanceModeNotFound {
	return &UsersToggleMaintenanceModeNotFound{}
}

/*
UsersToggleMaintenanceModeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UsersToggleMaintenanceModeNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this users toggle maintenance mode not found response has a 2xx status code
func (o *UsersToggleMaintenanceModeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle maintenance mode not found response has a 3xx status code
func (o *UsersToggleMaintenanceModeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode not found response has a 4xx status code
func (o *UsersToggleMaintenanceModeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this users toggle maintenance mode not found response has a 5xx status code
func (o *UsersToggleMaintenanceModeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this users toggle maintenance mode not found response a status code equal to that given
func (o *UsersToggleMaintenanceModeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UsersToggleMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *UsersToggleMaintenanceModeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *UsersToggleMaintenanceModeNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *UsersToggleMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersToggleMaintenanceModeInternalServerError creates a UsersToggleMaintenanceModeInternalServerError with default headers values
func NewUsersToggleMaintenanceModeInternalServerError() *UsersToggleMaintenanceModeInternalServerError {
	return &UsersToggleMaintenanceModeInternalServerError{}
}

/*
UsersToggleMaintenanceModeInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UsersToggleMaintenanceModeInternalServerError struct {
}

// IsSuccess returns true when this users toggle maintenance mode internal server error response has a 2xx status code
func (o *UsersToggleMaintenanceModeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users toggle maintenance mode internal server error response has a 3xx status code
func (o *UsersToggleMaintenanceModeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users toggle maintenance mode internal server error response has a 4xx status code
func (o *UsersToggleMaintenanceModeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this users toggle maintenance mode internal server error response has a 5xx status code
func (o *UsersToggleMaintenanceModeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this users toggle maintenance mode internal server error response a status code equal to that given
func (o *UsersToggleMaintenanceModeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UsersToggleMaintenanceModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeInternalServerError ", 500)
}

func (o *UsersToggleMaintenanceModeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Users/togglemaintenancemode][%d] usersToggleMaintenanceModeInternalServerError ", 500)
}

func (o *UsersToggleMaintenanceModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
