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

// BackupDescribeRestoreReader is a Reader for the BackupDescribeRestore structure.
type BackupDescribeRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDescribeRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDescribeRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupDescribeRestoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupDescribeRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupDescribeRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupDescribeRestoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupDescribeRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupDescribeRestoreOK creates a BackupDescribeRestoreOK with default headers values
func NewBackupDescribeRestoreOK() *BackupDescribeRestoreOK {
	return &BackupDescribeRestoreOK{}
}

/*
BackupDescribeRestoreOK describes a response with status code 200, with default header values.

Success
*/
type BackupDescribeRestoreOK struct {
	Payload interface{}
}

// IsSuccess returns true when this backup describe restore o k response has a 2xx status code
func (o *BackupDescribeRestoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup describe restore o k response has a 3xx status code
func (o *BackupDescribeRestoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore o k response has a 4xx status code
func (o *BackupDescribeRestoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup describe restore o k response has a 5xx status code
func (o *BackupDescribeRestoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe restore o k response a status code equal to that given
func (o *BackupDescribeRestoreOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupDescribeRestoreOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreOK  %+v", 200, o.Payload)
}

func (o *BackupDescribeRestoreOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreOK  %+v", 200, o.Payload)
}

func (o *BackupDescribeRestoreOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDescribeRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeRestoreBadRequest creates a BackupDescribeRestoreBadRequest with default headers values
func NewBackupDescribeRestoreBadRequest() *BackupDescribeRestoreBadRequest {
	return &BackupDescribeRestoreBadRequest{}
}

/*
BackupDescribeRestoreBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupDescribeRestoreBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this backup describe restore bad request response has a 2xx status code
func (o *BackupDescribeRestoreBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe restore bad request response has a 3xx status code
func (o *BackupDescribeRestoreBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore bad request response has a 4xx status code
func (o *BackupDescribeRestoreBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe restore bad request response has a 5xx status code
func (o *BackupDescribeRestoreBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe restore bad request response a status code equal to that given
func (o *BackupDescribeRestoreBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupDescribeRestoreBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDescribeRestoreBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDescribeRestoreBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupDescribeRestoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeRestoreUnauthorized creates a BackupDescribeRestoreUnauthorized with default headers values
func NewBackupDescribeRestoreUnauthorized() *BackupDescribeRestoreUnauthorized {
	return &BackupDescribeRestoreUnauthorized{}
}

/*
BackupDescribeRestoreUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupDescribeRestoreUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this backup describe restore unauthorized response has a 2xx status code
func (o *BackupDescribeRestoreUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe restore unauthorized response has a 3xx status code
func (o *BackupDescribeRestoreUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore unauthorized response has a 4xx status code
func (o *BackupDescribeRestoreUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe restore unauthorized response has a 5xx status code
func (o *BackupDescribeRestoreUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe restore unauthorized response a status code equal to that given
func (o *BackupDescribeRestoreUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupDescribeRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDescribeRestoreUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDescribeRestoreUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDescribeRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeRestoreForbidden creates a BackupDescribeRestoreForbidden with default headers values
func NewBackupDescribeRestoreForbidden() *BackupDescribeRestoreForbidden {
	return &BackupDescribeRestoreForbidden{}
}

/*
BackupDescribeRestoreForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupDescribeRestoreForbidden struct {
	Payload interface{}
}

// IsSuccess returns true when this backup describe restore forbidden response has a 2xx status code
func (o *BackupDescribeRestoreForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe restore forbidden response has a 3xx status code
func (o *BackupDescribeRestoreForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore forbidden response has a 4xx status code
func (o *BackupDescribeRestoreForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe restore forbidden response has a 5xx status code
func (o *BackupDescribeRestoreForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe restore forbidden response a status code equal to that given
func (o *BackupDescribeRestoreForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupDescribeRestoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreForbidden  %+v", 403, o.Payload)
}

func (o *BackupDescribeRestoreForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreForbidden  %+v", 403, o.Payload)
}

func (o *BackupDescribeRestoreForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDescribeRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeRestoreNotFound creates a BackupDescribeRestoreNotFound with default headers values
func NewBackupDescribeRestoreNotFound() *BackupDescribeRestoreNotFound {
	return &BackupDescribeRestoreNotFound{}
}

/*
BackupDescribeRestoreNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupDescribeRestoreNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this backup describe restore not found response has a 2xx status code
func (o *BackupDescribeRestoreNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe restore not found response has a 3xx status code
func (o *BackupDescribeRestoreNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore not found response has a 4xx status code
func (o *BackupDescribeRestoreNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe restore not found response has a 5xx status code
func (o *BackupDescribeRestoreNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe restore not found response a status code equal to that given
func (o *BackupDescribeRestoreNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupDescribeRestoreNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreNotFound  %+v", 404, o.Payload)
}

func (o *BackupDescribeRestoreNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreNotFound  %+v", 404, o.Payload)
}

func (o *BackupDescribeRestoreNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDescribeRestoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeRestoreInternalServerError creates a BackupDescribeRestoreInternalServerError with default headers values
func NewBackupDescribeRestoreInternalServerError() *BackupDescribeRestoreInternalServerError {
	return &BackupDescribeRestoreInternalServerError{}
}

/*
BackupDescribeRestoreInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupDescribeRestoreInternalServerError struct {
}

// IsSuccess returns true when this backup describe restore internal server error response has a 2xx status code
func (o *BackupDescribeRestoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe restore internal server error response has a 3xx status code
func (o *BackupDescribeRestoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe restore internal server error response has a 4xx status code
func (o *BackupDescribeRestoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup describe restore internal server error response has a 5xx status code
func (o *BackupDescribeRestoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup describe restore internal server error response a status code equal to that given
func (o *BackupDescribeRestoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupDescribeRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreInternalServerError ", 500)
}

func (o *BackupDescribeRestoreInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/restore/{projectId}/{name}][%d] backupDescribeRestoreInternalServerError ", 500)
}

func (o *BackupDescribeRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
