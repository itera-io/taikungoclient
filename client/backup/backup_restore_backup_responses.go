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

// BackupRestoreBackupReader is a Reader for the BackupRestoreBackup structure.
type BackupRestoreBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupRestoreBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupRestoreBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupRestoreBackupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupRestoreBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupRestoreBackupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupRestoreBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupRestoreBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupRestoreBackupOK creates a BackupRestoreBackupOK with default headers values
func NewBackupRestoreBackupOK() *BackupRestoreBackupOK {
	return &BackupRestoreBackupOK{}
}

/*
BackupRestoreBackupOK describes a response with status code 200, with default header values.

Success
*/
type BackupRestoreBackupOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup restore backup o k response has a 2xx status code
func (o *BackupRestoreBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup restore backup o k response has a 3xx status code
func (o *BackupRestoreBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup o k response has a 4xx status code
func (o *BackupRestoreBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup restore backup o k response has a 5xx status code
func (o *BackupRestoreBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup restore backup o k response a status code equal to that given
func (o *BackupRestoreBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the backup restore backup o k response
func (o *BackupRestoreBackupOK) Code() int {
	return 200
}

func (o *BackupRestoreBackupOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupOK  %+v", 200, o.Payload)
}

func (o *BackupRestoreBackupOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupOK  %+v", 200, o.Payload)
}

func (o *BackupRestoreBackupOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupRestoreBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupRestoreBackupBadRequest creates a BackupRestoreBackupBadRequest with default headers values
func NewBackupRestoreBackupBadRequest() *BackupRestoreBackupBadRequest {
	return &BackupRestoreBackupBadRequest{}
}

/*
BackupRestoreBackupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupRestoreBackupBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup restore backup bad request response has a 2xx status code
func (o *BackupRestoreBackupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup restore backup bad request response has a 3xx status code
func (o *BackupRestoreBackupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup bad request response has a 4xx status code
func (o *BackupRestoreBackupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup restore backup bad request response has a 5xx status code
func (o *BackupRestoreBackupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup restore backup bad request response a status code equal to that given
func (o *BackupRestoreBackupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the backup restore backup bad request response
func (o *BackupRestoreBackupBadRequest) Code() int {
	return 400
}

func (o *BackupRestoreBackupBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupBadRequest  %+v", 400, o.Payload)
}

func (o *BackupRestoreBackupBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupBadRequest  %+v", 400, o.Payload)
}

func (o *BackupRestoreBackupBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupRestoreBackupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupRestoreBackupUnauthorized creates a BackupRestoreBackupUnauthorized with default headers values
func NewBackupRestoreBackupUnauthorized() *BackupRestoreBackupUnauthorized {
	return &BackupRestoreBackupUnauthorized{}
}

/*
BackupRestoreBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupRestoreBackupUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup restore backup unauthorized response has a 2xx status code
func (o *BackupRestoreBackupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup restore backup unauthorized response has a 3xx status code
func (o *BackupRestoreBackupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup unauthorized response has a 4xx status code
func (o *BackupRestoreBackupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup restore backup unauthorized response has a 5xx status code
func (o *BackupRestoreBackupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup restore backup unauthorized response a status code equal to that given
func (o *BackupRestoreBackupUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the backup restore backup unauthorized response
func (o *BackupRestoreBackupUnauthorized) Code() int {
	return 401
}

func (o *BackupRestoreBackupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupRestoreBackupUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupRestoreBackupUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupRestoreBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupRestoreBackupForbidden creates a BackupRestoreBackupForbidden with default headers values
func NewBackupRestoreBackupForbidden() *BackupRestoreBackupForbidden {
	return &BackupRestoreBackupForbidden{}
}

/*
BackupRestoreBackupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupRestoreBackupForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup restore backup forbidden response has a 2xx status code
func (o *BackupRestoreBackupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup restore backup forbidden response has a 3xx status code
func (o *BackupRestoreBackupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup forbidden response has a 4xx status code
func (o *BackupRestoreBackupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup restore backup forbidden response has a 5xx status code
func (o *BackupRestoreBackupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup restore backup forbidden response a status code equal to that given
func (o *BackupRestoreBackupForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the backup restore backup forbidden response
func (o *BackupRestoreBackupForbidden) Code() int {
	return 403
}

func (o *BackupRestoreBackupForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupForbidden  %+v", 403, o.Payload)
}

func (o *BackupRestoreBackupForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupForbidden  %+v", 403, o.Payload)
}

func (o *BackupRestoreBackupForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupRestoreBackupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupRestoreBackupNotFound creates a BackupRestoreBackupNotFound with default headers values
func NewBackupRestoreBackupNotFound() *BackupRestoreBackupNotFound {
	return &BackupRestoreBackupNotFound{}
}

/*
BackupRestoreBackupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupRestoreBackupNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup restore backup not found response has a 2xx status code
func (o *BackupRestoreBackupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup restore backup not found response has a 3xx status code
func (o *BackupRestoreBackupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup not found response has a 4xx status code
func (o *BackupRestoreBackupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup restore backup not found response has a 5xx status code
func (o *BackupRestoreBackupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup restore backup not found response a status code equal to that given
func (o *BackupRestoreBackupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the backup restore backup not found response
func (o *BackupRestoreBackupNotFound) Code() int {
	return 404
}

func (o *BackupRestoreBackupNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupNotFound  %+v", 404, o.Payload)
}

func (o *BackupRestoreBackupNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupNotFound  %+v", 404, o.Payload)
}

func (o *BackupRestoreBackupNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupRestoreBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupRestoreBackupInternalServerError creates a BackupRestoreBackupInternalServerError with default headers values
func NewBackupRestoreBackupInternalServerError() *BackupRestoreBackupInternalServerError {
	return &BackupRestoreBackupInternalServerError{}
}

/*
BackupRestoreBackupInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupRestoreBackupInternalServerError struct {
}

// IsSuccess returns true when this backup restore backup internal server error response has a 2xx status code
func (o *BackupRestoreBackupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup restore backup internal server error response has a 3xx status code
func (o *BackupRestoreBackupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup restore backup internal server error response has a 4xx status code
func (o *BackupRestoreBackupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup restore backup internal server error response has a 5xx status code
func (o *BackupRestoreBackupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup restore backup internal server error response a status code equal to that given
func (o *BackupRestoreBackupInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the backup restore backup internal server error response
func (o *BackupRestoreBackupInternalServerError) Code() int {
	return 500
}

func (o *BackupRestoreBackupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupInternalServerError ", 500)
}

func (o *BackupRestoreBackupInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/restore][%d] backupRestoreBackupInternalServerError ", 500)
}

func (o *BackupRestoreBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
