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

// BackupDescribeBackupReader is a Reader for the BackupDescribeBackup structure.
type BackupDescribeBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDescribeBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDescribeBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupDescribeBackupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupDescribeBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupDescribeBackupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupDescribeBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupDescribeBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupDescribeBackupOK creates a BackupDescribeBackupOK with default headers values
func NewBackupDescribeBackupOK() *BackupDescribeBackupOK {
	return &BackupDescribeBackupOK{}
}

/*
BackupDescribeBackupOK describes a response with status code 200, with default header values.

Success
*/
type BackupDescribeBackupOK struct {
	Payload interface{}
}

// IsSuccess returns true when this backup describe backup o k response has a 2xx status code
func (o *BackupDescribeBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup describe backup o k response has a 3xx status code
func (o *BackupDescribeBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup o k response has a 4xx status code
func (o *BackupDescribeBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup describe backup o k response has a 5xx status code
func (o *BackupDescribeBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe backup o k response a status code equal to that given
func (o *BackupDescribeBackupOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupDescribeBackupOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupOK  %+v", 200, o.Payload)
}

func (o *BackupDescribeBackupOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupOK  %+v", 200, o.Payload)
}

func (o *BackupDescribeBackupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *BackupDescribeBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeBackupBadRequest creates a BackupDescribeBackupBadRequest with default headers values
func NewBackupDescribeBackupBadRequest() *BackupDescribeBackupBadRequest {
	return &BackupDescribeBackupBadRequest{}
}

/*
BackupDescribeBackupBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupDescribeBackupBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this backup describe backup bad request response has a 2xx status code
func (o *BackupDescribeBackupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe backup bad request response has a 3xx status code
func (o *BackupDescribeBackupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup bad request response has a 4xx status code
func (o *BackupDescribeBackupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe backup bad request response has a 5xx status code
func (o *BackupDescribeBackupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe backup bad request response a status code equal to that given
func (o *BackupDescribeBackupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupDescribeBackupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDescribeBackupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupBadRequest  %+v", 400, o.Payload)
}

func (o *BackupDescribeBackupBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupDescribeBackupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeBackupUnauthorized creates a BackupDescribeBackupUnauthorized with default headers values
func NewBackupDescribeBackupUnauthorized() *BackupDescribeBackupUnauthorized {
	return &BackupDescribeBackupUnauthorized{}
}

/*
BackupDescribeBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupDescribeBackupUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup describe backup unauthorized response has a 2xx status code
func (o *BackupDescribeBackupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe backup unauthorized response has a 3xx status code
func (o *BackupDescribeBackupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup unauthorized response has a 4xx status code
func (o *BackupDescribeBackupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe backup unauthorized response has a 5xx status code
func (o *BackupDescribeBackupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe backup unauthorized response a status code equal to that given
func (o *BackupDescribeBackupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupDescribeBackupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDescribeBackupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupDescribeBackupUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupDescribeBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeBackupForbidden creates a BackupDescribeBackupForbidden with default headers values
func NewBackupDescribeBackupForbidden() *BackupDescribeBackupForbidden {
	return &BackupDescribeBackupForbidden{}
}

/*
BackupDescribeBackupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupDescribeBackupForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup describe backup forbidden response has a 2xx status code
func (o *BackupDescribeBackupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe backup forbidden response has a 3xx status code
func (o *BackupDescribeBackupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup forbidden response has a 4xx status code
func (o *BackupDescribeBackupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe backup forbidden response has a 5xx status code
func (o *BackupDescribeBackupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe backup forbidden response a status code equal to that given
func (o *BackupDescribeBackupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupDescribeBackupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupForbidden  %+v", 403, o.Payload)
}

func (o *BackupDescribeBackupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupForbidden  %+v", 403, o.Payload)
}

func (o *BackupDescribeBackupForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupDescribeBackupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeBackupNotFound creates a BackupDescribeBackupNotFound with default headers values
func NewBackupDescribeBackupNotFound() *BackupDescribeBackupNotFound {
	return &BackupDescribeBackupNotFound{}
}

/*
BackupDescribeBackupNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupDescribeBackupNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup describe backup not found response has a 2xx status code
func (o *BackupDescribeBackupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe backup not found response has a 3xx status code
func (o *BackupDescribeBackupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup not found response has a 4xx status code
func (o *BackupDescribeBackupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup describe backup not found response has a 5xx status code
func (o *BackupDescribeBackupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup describe backup not found response a status code equal to that given
func (o *BackupDescribeBackupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupDescribeBackupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupNotFound  %+v", 404, o.Payload)
}

func (o *BackupDescribeBackupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupNotFound  %+v", 404, o.Payload)
}

func (o *BackupDescribeBackupNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupDescribeBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDescribeBackupInternalServerError creates a BackupDescribeBackupInternalServerError with default headers values
func NewBackupDescribeBackupInternalServerError() *BackupDescribeBackupInternalServerError {
	return &BackupDescribeBackupInternalServerError{}
}

/*
BackupDescribeBackupInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupDescribeBackupInternalServerError struct {
}

// IsSuccess returns true when this backup describe backup internal server error response has a 2xx status code
func (o *BackupDescribeBackupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup describe backup internal server error response has a 3xx status code
func (o *BackupDescribeBackupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup describe backup internal server error response has a 4xx status code
func (o *BackupDescribeBackupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup describe backup internal server error response has a 5xx status code
func (o *BackupDescribeBackupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup describe backup internal server error response a status code equal to that given
func (o *BackupDescribeBackupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupDescribeBackupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupInternalServerError ", 500)
}

func (o *BackupDescribeBackupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/describe/backup/{projectId}/{name}][%d] backupDescribeBackupInternalServerError ", 500)
}

func (o *BackupDescribeBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
