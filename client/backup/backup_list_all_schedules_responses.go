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

// BackupListAllSchedulesReader is a Reader for the BackupListAllSchedules structure.
type BackupListAllSchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupListAllSchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupListAllSchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupListAllSchedulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupListAllSchedulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupListAllSchedulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupListAllSchedulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupListAllSchedulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupListAllSchedulesOK creates a BackupListAllSchedulesOK with default headers values
func NewBackupListAllSchedulesOK() *BackupListAllSchedulesOK {
	return &BackupListAllSchedulesOK{}
}

/*
BackupListAllSchedulesOK describes a response with status code 200, with default header values.

Success
*/
type BackupListAllSchedulesOK struct {
	Payload *models.ListAllSchedules
}

// IsSuccess returns true when this backup list all schedules o k response has a 2xx status code
func (o *BackupListAllSchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup list all schedules o k response has a 3xx status code
func (o *BackupListAllSchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules o k response has a 4xx status code
func (o *BackupListAllSchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all schedules o k response has a 5xx status code
func (o *BackupListAllSchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all schedules o k response a status code equal to that given
func (o *BackupListAllSchedulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupListAllSchedulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesOK  %+v", 200, o.Payload)
}

func (o *BackupListAllSchedulesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesOK  %+v", 200, o.Payload)
}

func (o *BackupListAllSchedulesOK) GetPayload() *models.ListAllSchedules {
	return o.Payload
}

func (o *BackupListAllSchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAllSchedules)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllSchedulesBadRequest creates a BackupListAllSchedulesBadRequest with default headers values
func NewBackupListAllSchedulesBadRequest() *BackupListAllSchedulesBadRequest {
	return &BackupListAllSchedulesBadRequest{}
}

/*
BackupListAllSchedulesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupListAllSchedulesBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this backup list all schedules bad request response has a 2xx status code
func (o *BackupListAllSchedulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all schedules bad request response has a 3xx status code
func (o *BackupListAllSchedulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules bad request response has a 4xx status code
func (o *BackupListAllSchedulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all schedules bad request response has a 5xx status code
func (o *BackupListAllSchedulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all schedules bad request response a status code equal to that given
func (o *BackupListAllSchedulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupListAllSchedulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllSchedulesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllSchedulesBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupListAllSchedulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllSchedulesUnauthorized creates a BackupListAllSchedulesUnauthorized with default headers values
func NewBackupListAllSchedulesUnauthorized() *BackupListAllSchedulesUnauthorized {
	return &BackupListAllSchedulesUnauthorized{}
}

/*
BackupListAllSchedulesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupListAllSchedulesUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all schedules unauthorized response has a 2xx status code
func (o *BackupListAllSchedulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all schedules unauthorized response has a 3xx status code
func (o *BackupListAllSchedulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules unauthorized response has a 4xx status code
func (o *BackupListAllSchedulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all schedules unauthorized response has a 5xx status code
func (o *BackupListAllSchedulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all schedules unauthorized response a status code equal to that given
func (o *BackupListAllSchedulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupListAllSchedulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllSchedulesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllSchedulesUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllSchedulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllSchedulesForbidden creates a BackupListAllSchedulesForbidden with default headers values
func NewBackupListAllSchedulesForbidden() *BackupListAllSchedulesForbidden {
	return &BackupListAllSchedulesForbidden{}
}

/*
BackupListAllSchedulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupListAllSchedulesForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all schedules forbidden response has a 2xx status code
func (o *BackupListAllSchedulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all schedules forbidden response has a 3xx status code
func (o *BackupListAllSchedulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules forbidden response has a 4xx status code
func (o *BackupListAllSchedulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all schedules forbidden response has a 5xx status code
func (o *BackupListAllSchedulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all schedules forbidden response a status code equal to that given
func (o *BackupListAllSchedulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupListAllSchedulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllSchedulesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllSchedulesForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllSchedulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllSchedulesNotFound creates a BackupListAllSchedulesNotFound with default headers values
func NewBackupListAllSchedulesNotFound() *BackupListAllSchedulesNotFound {
	return &BackupListAllSchedulesNotFound{}
}

/*
BackupListAllSchedulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupListAllSchedulesNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all schedules not found response has a 2xx status code
func (o *BackupListAllSchedulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all schedules not found response has a 3xx status code
func (o *BackupListAllSchedulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules not found response has a 4xx status code
func (o *BackupListAllSchedulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all schedules not found response has a 5xx status code
func (o *BackupListAllSchedulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all schedules not found response a status code equal to that given
func (o *BackupListAllSchedulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupListAllSchedulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllSchedulesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllSchedulesNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllSchedulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllSchedulesInternalServerError creates a BackupListAllSchedulesInternalServerError with default headers values
func NewBackupListAllSchedulesInternalServerError() *BackupListAllSchedulesInternalServerError {
	return &BackupListAllSchedulesInternalServerError{}
}

/*
BackupListAllSchedulesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupListAllSchedulesInternalServerError struct {
}

// IsSuccess returns true when this backup list all schedules internal server error response has a 2xx status code
func (o *BackupListAllSchedulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all schedules internal server error response has a 3xx status code
func (o *BackupListAllSchedulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all schedules internal server error response has a 4xx status code
func (o *BackupListAllSchedulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all schedules internal server error response has a 5xx status code
func (o *BackupListAllSchedulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup list all schedules internal server error response a status code equal to that given
func (o *BackupListAllSchedulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupListAllSchedulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesInternalServerError ", 500)
}

func (o *BackupListAllSchedulesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedules/{projectId}][%d] backupListAllSchedulesInternalServerError ", 500)
}

func (o *BackupListAllSchedulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
