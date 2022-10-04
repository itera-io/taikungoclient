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

// BackupDeleteRestoreReader is a Reader for the BackupDeleteRestore structure.
type BackupDeleteRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDeleteRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDeleteRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupDeleteRestoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupDeleteRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupDeleteRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupDeleteRestoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupDeleteRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupDeleteRestoreOK creates a BackupDeleteRestoreOK with default headers values
func NewBackupDeleteRestoreOK() *BackupDeleteRestoreOK {
	return &BackupDeleteRestoreOK{}
}

/*
BackupDeleteRestoreOK describes a response with status code 200, with default header values.

Success
*/
type BackupDeleteRestoreOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup delete restore o k response has a 2xx status code
func (o *BackupDeleteRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup delete restore o k response has a 3xx status code
func (o *BackupDeleteRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore o k response has a 4xx status code
func (o *BackupDeleteRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup delete restore o k response has a 5xx status code
func (o *BackupDeleteRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete restore o k response a status code equal to that given
func (o *BackupDeleteRestoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupDeleteRestoreOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreOK  %+v", 200, o.Payload)
}

func (o *BackupDeleteRestoreOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreOK  %+v", 200, o.Payload)
}

func (o *BackupDeleteRestoreOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupDeleteRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteRestoreBadRequest creates a BackupDeleteRestoreBadRequest with default headers values
func NewBackupDeleteRestoreBadRequest() *BackupDeleteRestoreBadRequest {
	return &BackupDeleteRestoreBadRequest{}
}

/*
BackupDeleteRestoreBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupDeleteRestoreBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this backup delete restore bad request response has a 2xx status code
func (o *BackupDeleteRestoreBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete restore bad request response has a 3xx status code
func (o *BackupDeleteRestoreBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore bad request response has a 4xx status code
func (o *BackupDeleteRestoreBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete restore bad request response has a 5xx status code
func (o *BackupDeleteRestoreBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete restore bad request response a status code equal to that given
func (o *BackupDeleteRestoreBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupDeleteRestoreBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDeleteRestoreBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDeleteRestoreBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupDeleteRestoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteRestoreUnauthorized creates a BackupDeleteRestoreUnauthorized with default headers values
func NewBackupDeleteRestoreUnauthorized() *BackupDeleteRestoreUnauthorized {
	return &BackupDeleteRestoreUnauthorized{}
}

/*
BackupDeleteRestoreUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupDeleteRestoreUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this backup delete restore unauthorized response has a 2xx status code
func (o *BackupDeleteRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete restore unauthorized response has a 3xx status code
func (o *BackupDeleteRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore unauthorized response has a 4xx status code
func (o *BackupDeleteRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete restore unauthorized response has a 5xx status code
func (o *BackupDeleteRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete restore unauthorized response a status code equal to that given
func (o *BackupDeleteRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupDeleteRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDeleteRestoreUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDeleteRestoreUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDeleteRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteRestoreForbidden creates a BackupDeleteRestoreForbidden with default headers values
func NewBackupDeleteRestoreForbidden() *BackupDeleteRestoreForbidden {
	return &BackupDeleteRestoreForbidden{}
}

/*
BackupDeleteRestoreForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupDeleteRestoreForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this backup delete restore forbidden response has a 2xx status code
func (o *BackupDeleteRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete restore forbidden response has a 3xx status code
func (o *BackupDeleteRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore forbidden response has a 4xx status code
func (o *BackupDeleteRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete restore forbidden response has a 5xx status code
func (o *BackupDeleteRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete restore forbidden response a status code equal to that given
func (o *BackupDeleteRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupDeleteRestoreForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreForbidden  %+v", 403, o.Payload)
}

func (o *BackupDeleteRestoreForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreForbidden  %+v", 403, o.Payload)
}

func (o *BackupDeleteRestoreForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDeleteRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteRestoreNotFound creates a BackupDeleteRestoreNotFound with default headers values
func NewBackupDeleteRestoreNotFound() *BackupDeleteRestoreNotFound {
	return &BackupDeleteRestoreNotFound{}
}

/*
BackupDeleteRestoreNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupDeleteRestoreNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this backup delete restore not found response has a 2xx status code
func (o *BackupDeleteRestoreNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete restore not found response has a 3xx status code
func (o *BackupDeleteRestoreNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore not found response has a 4xx status code
func (o *BackupDeleteRestoreNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup delete restore not found response has a 5xx status code
func (o *BackupDeleteRestoreNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup delete restore not found response a status code equal to that given
func (o *BackupDeleteRestoreNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupDeleteRestoreNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreNotFound  %+v", 404, o.Payload)
}

func (o *BackupDeleteRestoreNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreNotFound  %+v", 404, o.Payload)
}

func (o *BackupDeleteRestoreNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDeleteRestoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDeleteRestoreInternalServerError creates a BackupDeleteRestoreInternalServerError with default headers values
func NewBackupDeleteRestoreInternalServerError() *BackupDeleteRestoreInternalServerError {
	return &BackupDeleteRestoreInternalServerError{}
}

/*
BackupDeleteRestoreInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupDeleteRestoreInternalServerError struct {
}

// IsSuccess returns true when this backup delete restore internal server error response has a 2xx status code
func (o *BackupDeleteRestoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup delete restore internal server error response has a 3xx status code
func (o *BackupDeleteRestoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup delete restore internal server error response has a 4xx status code
func (o *BackupDeleteRestoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup delete restore internal server error response has a 5xx status code
func (o *BackupDeleteRestoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup delete restore internal server error response a status code equal to that given
func (o *BackupDeleteRestoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupDeleteRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreInternalServerError ", 500)
}

func (o *BackupDeleteRestoreInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/delete/restore][%d] backupDeleteRestoreInternalServerError ", 500)
}

func (o *BackupDeleteRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
