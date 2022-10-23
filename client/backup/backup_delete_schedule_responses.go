// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// BackupDeleteScheduleReader is a Reader for the BackupDeleteSchedule structure.
type BackupDeleteScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDeleteScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDeleteScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupDeleteScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupDeleteScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupDeleteScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupDeleteScheduleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupDeleteScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupDeleteScheduleOK creates a BackupDeleteScheduleOK with default headers values
func NewBackupDeleteScheduleOK() *BackupDeleteScheduleOK {
	return &BackupDeleteScheduleOK{}
}

/*
BackupDeleteScheduleOK describes a response with status code 200, with default header values.

Success
*/
type BackupDeleteScheduleOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup delete schedule o k response has a 2xx status code
func (o *BackupDeleteScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup delete schedule o k response has a 3xx status code
func (o *BackupDeleteScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule o k response has a 4xx status code
func (o *BackupDeleteScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup delete schedule o k response has a 5xx status code
func (o *BackupDeleteScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete schedule o k response a status code equal to that given
func (o *BackupDeleteScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupDeleteScheduleOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleOK  %+v", 200, o.Payload)
}

func (o *BackupDeleteScheduleOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleOK  %+v", 200, o.Payload)
}

func (o *BackupDeleteScheduleOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupDeleteScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteScheduleBadRequest creates a BackupDeleteScheduleBadRequest with default headers values
func NewBackupDeleteScheduleBadRequest() *BackupDeleteScheduleBadRequest {
	return &BackupDeleteScheduleBadRequest{}
}

/*
BackupDeleteScheduleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupDeleteScheduleBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this backup delete schedule bad request response has a 2xx status code
func (o *BackupDeleteScheduleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete schedule bad request response has a 3xx status code
func (o *BackupDeleteScheduleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule bad request response has a 4xx status code
func (o *BackupDeleteScheduleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete schedule bad request response has a 5xx status code
func (o *BackupDeleteScheduleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete schedule bad request response a status code equal to that given
func (o *BackupDeleteScheduleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupDeleteScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDeleteScheduleBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDeleteScheduleBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *BackupDeleteScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteScheduleUnauthorized creates a BackupDeleteScheduleUnauthorized with default headers values
func NewBackupDeleteScheduleUnauthorized() *BackupDeleteScheduleUnauthorized {
	return &BackupDeleteScheduleUnauthorized{}
}

/*
BackupDeleteScheduleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupDeleteScheduleUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup delete schedule unauthorized response has a 2xx status code
func (o *BackupDeleteScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete schedule unauthorized response has a 3xx status code
func (o *BackupDeleteScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule unauthorized response has a 4xx status code
func (o *BackupDeleteScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete schedule unauthorized response has a 5xx status code
func (o *BackupDeleteScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete schedule unauthorized response a status code equal to that given
func (o *BackupDeleteScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupDeleteScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDeleteScheduleUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDeleteScheduleUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupDeleteScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteScheduleForbidden creates a BackupDeleteScheduleForbidden with default headers values
func NewBackupDeleteScheduleForbidden() *BackupDeleteScheduleForbidden {
	return &BackupDeleteScheduleForbidden{}
}

/*
BackupDeleteScheduleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupDeleteScheduleForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup delete schedule forbidden response has a 2xx status code
func (o *BackupDeleteScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete schedule forbidden response has a 3xx status code
func (o *BackupDeleteScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule forbidden response has a 4xx status code
func (o *BackupDeleteScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete schedule forbidden response has a 5xx status code
func (o *BackupDeleteScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete schedule forbidden response a status code equal to that given
func (o *BackupDeleteScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupDeleteScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleForbidden  %+v", 403, o.Payload)
}

func (o *BackupDeleteScheduleForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleForbidden  %+v", 403, o.Payload)
}

func (o *BackupDeleteScheduleForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupDeleteScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteScheduleNotFound creates a BackupDeleteScheduleNotFound with default headers values
func NewBackupDeleteScheduleNotFound() *BackupDeleteScheduleNotFound {
	return &BackupDeleteScheduleNotFound{}
}

/*
BackupDeleteScheduleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupDeleteScheduleNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup delete schedule not found response has a 2xx status code
func (o *BackupDeleteScheduleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete schedule not found response has a 3xx status code
func (o *BackupDeleteScheduleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule not found response has a 4xx status code
func (o *BackupDeleteScheduleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete schedule not found response has a 5xx status code
func (o *BackupDeleteScheduleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete schedule not found response a status code equal to that given
func (o *BackupDeleteScheduleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupDeleteScheduleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleNotFound  %+v", 404, o.Payload)
}

func (o *BackupDeleteScheduleNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleNotFound  %+v", 404, o.Payload)
}

func (o *BackupDeleteScheduleNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupDeleteScheduleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteScheduleInternalServerError creates a BackupDeleteScheduleInternalServerError with default headers values
func NewBackupDeleteScheduleInternalServerError() *BackupDeleteScheduleInternalServerError {
	return &BackupDeleteScheduleInternalServerError{}
}

/*
BackupDeleteScheduleInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupDeleteScheduleInternalServerError struct {
}

// IsSuccess returns true when this backup delete schedule internal server error response has a 2xx status code
func (o *BackupDeleteScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete schedule internal server error response has a 3xx status code
func (o *BackupDeleteScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete schedule internal server error response has a 4xx status code
func (o *BackupDeleteScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup delete schedule internal server error response has a 5xx status code
func (o *BackupDeleteScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup delete schedule internal server error response a status code equal to that given
func (o *BackupDeleteScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupDeleteScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleInternalServerError ", 500)
}

func (o *BackupDeleteScheduleInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/schedule][%d] backupDeleteScheduleInternalServerError ", 500)
}

func (o *BackupDeleteScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
