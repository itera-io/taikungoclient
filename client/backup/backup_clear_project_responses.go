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

// BackupClearProjectReader is a Reader for the BackupClearProject structure.
type BackupClearProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupClearProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupClearProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupClearProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupClearProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupClearProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupClearProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupClearProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupClearProjectOK creates a BackupClearProjectOK with default headers values
func NewBackupClearProjectOK() *BackupClearProjectOK {
	return &BackupClearProjectOK{}
}

/*
BackupClearProjectOK describes a response with status code 200, with default header values.

Success
*/
type BackupClearProjectOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup clear project o k response has a 2xx status code
func (o *BackupClearProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup clear project o k response has a 3xx status code
func (o *BackupClearProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project o k response has a 4xx status code
func (o *BackupClearProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup clear project o k response has a 5xx status code
func (o *BackupClearProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup clear project o k response a status code equal to that given
func (o *BackupClearProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupClearProjectOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectOK  %+v", 200, o.Payload)
}

func (o *BackupClearProjectOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectOK  %+v", 200, o.Payload)
}

func (o *BackupClearProjectOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupClearProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupClearProjectBadRequest creates a BackupClearProjectBadRequest with default headers values
func NewBackupClearProjectBadRequest() *BackupClearProjectBadRequest {
	return &BackupClearProjectBadRequest{}
}

/*
BackupClearProjectBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupClearProjectBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup clear project bad request response has a 2xx status code
func (o *BackupClearProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup clear project bad request response has a 3xx status code
func (o *BackupClearProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project bad request response has a 4xx status code
func (o *BackupClearProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup clear project bad request response has a 5xx status code
func (o *BackupClearProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup clear project bad request response a status code equal to that given
func (o *BackupClearProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupClearProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectBadRequest  %+v", 400, o.Payload)
}

func (o *BackupClearProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectBadRequest  %+v", 400, o.Payload)
}

func (o *BackupClearProjectBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupClearProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupClearProjectUnauthorized creates a BackupClearProjectUnauthorized with default headers values
func NewBackupClearProjectUnauthorized() *BackupClearProjectUnauthorized {
	return &BackupClearProjectUnauthorized{}
}

/*
BackupClearProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupClearProjectUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup clear project unauthorized response has a 2xx status code
func (o *BackupClearProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup clear project unauthorized response has a 3xx status code
func (o *BackupClearProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project unauthorized response has a 4xx status code
func (o *BackupClearProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup clear project unauthorized response has a 5xx status code
func (o *BackupClearProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup clear project unauthorized response a status code equal to that given
func (o *BackupClearProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupClearProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupClearProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupClearProjectUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupClearProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupClearProjectForbidden creates a BackupClearProjectForbidden with default headers values
func NewBackupClearProjectForbidden() *BackupClearProjectForbidden {
	return &BackupClearProjectForbidden{}
}

/*
BackupClearProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupClearProjectForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup clear project forbidden response has a 2xx status code
func (o *BackupClearProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup clear project forbidden response has a 3xx status code
func (o *BackupClearProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project forbidden response has a 4xx status code
func (o *BackupClearProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup clear project forbidden response has a 5xx status code
func (o *BackupClearProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup clear project forbidden response a status code equal to that given
func (o *BackupClearProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupClearProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectForbidden  %+v", 403, o.Payload)
}

func (o *BackupClearProjectForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectForbidden  %+v", 403, o.Payload)
}

func (o *BackupClearProjectForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupClearProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupClearProjectNotFound creates a BackupClearProjectNotFound with default headers values
func NewBackupClearProjectNotFound() *BackupClearProjectNotFound {
	return &BackupClearProjectNotFound{}
}

/*
BackupClearProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupClearProjectNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup clear project not found response has a 2xx status code
func (o *BackupClearProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup clear project not found response has a 3xx status code
func (o *BackupClearProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project not found response has a 4xx status code
func (o *BackupClearProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup clear project not found response has a 5xx status code
func (o *BackupClearProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup clear project not found response a status code equal to that given
func (o *BackupClearProjectNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupClearProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectNotFound  %+v", 404, o.Payload)
}

func (o *BackupClearProjectNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectNotFound  %+v", 404, o.Payload)
}

func (o *BackupClearProjectNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupClearProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupClearProjectInternalServerError creates a BackupClearProjectInternalServerError with default headers values
func NewBackupClearProjectInternalServerError() *BackupClearProjectInternalServerError {
	return &BackupClearProjectInternalServerError{}
}

/*
BackupClearProjectInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupClearProjectInternalServerError struct {
}

// IsSuccess returns true when this backup clear project internal server error response has a 2xx status code
func (o *BackupClearProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup clear project internal server error response has a 3xx status code
func (o *BackupClearProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup clear project internal server error response has a 4xx status code
func (o *BackupClearProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup clear project internal server error response has a 5xx status code
func (o *BackupClearProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup clear project internal server error response a status code equal to that given
func (o *BackupClearProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupClearProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectInternalServerError ", 500)
}

func (o *BackupClearProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/clear/project][%d] backupClearProjectInternalServerError ", 500)
}

func (o *BackupClearProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
