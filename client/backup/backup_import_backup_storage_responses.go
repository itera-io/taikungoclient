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

// BackupImportBackupStorageReader is a Reader for the BackupImportBackupStorage structure.
type BackupImportBackupStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupImportBackupStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupImportBackupStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupImportBackupStorageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupImportBackupStorageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupImportBackupStorageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupImportBackupStorageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupImportBackupStorageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupImportBackupStorageOK creates a BackupImportBackupStorageOK with default headers values
func NewBackupImportBackupStorageOK() *BackupImportBackupStorageOK {
	return &BackupImportBackupStorageOK{}
}

/*
BackupImportBackupStorageOK describes a response with status code 200, with default header values.

Success
*/
type BackupImportBackupStorageOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup import backup storage o k response has a 2xx status code
func (o *BackupImportBackupStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup import backup storage o k response has a 3xx status code
func (o *BackupImportBackupStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage o k response has a 4xx status code
func (o *BackupImportBackupStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup import backup storage o k response has a 5xx status code
func (o *BackupImportBackupStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup import backup storage o k response a status code equal to that given
func (o *BackupImportBackupStorageOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupImportBackupStorageOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageOK  %+v", 200, o.Payload)
}

func (o *BackupImportBackupStorageOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageOK  %+v", 200, o.Payload)
}

func (o *BackupImportBackupStorageOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupImportBackupStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupImportBackupStorageBadRequest creates a BackupImportBackupStorageBadRequest with default headers values
func NewBackupImportBackupStorageBadRequest() *BackupImportBackupStorageBadRequest {
	return &BackupImportBackupStorageBadRequest{}
}

/*
BackupImportBackupStorageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupImportBackupStorageBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this backup import backup storage bad request response has a 2xx status code
func (o *BackupImportBackupStorageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup import backup storage bad request response has a 3xx status code
func (o *BackupImportBackupStorageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage bad request response has a 4xx status code
func (o *BackupImportBackupStorageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup import backup storage bad request response has a 5xx status code
func (o *BackupImportBackupStorageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup import backup storage bad request response a status code equal to that given
func (o *BackupImportBackupStorageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupImportBackupStorageBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageBadRequest  %+v", 400, o.Payload)
}

func (o *BackupImportBackupStorageBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageBadRequest  %+v", 400, o.Payload)
}

func (o *BackupImportBackupStorageBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *BackupImportBackupStorageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupImportBackupStorageUnauthorized creates a BackupImportBackupStorageUnauthorized with default headers values
func NewBackupImportBackupStorageUnauthorized() *BackupImportBackupStorageUnauthorized {
	return &BackupImportBackupStorageUnauthorized{}
}

/*
BackupImportBackupStorageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupImportBackupStorageUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup import backup storage unauthorized response has a 2xx status code
func (o *BackupImportBackupStorageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup import backup storage unauthorized response has a 3xx status code
func (o *BackupImportBackupStorageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage unauthorized response has a 4xx status code
func (o *BackupImportBackupStorageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup import backup storage unauthorized response has a 5xx status code
func (o *BackupImportBackupStorageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup import backup storage unauthorized response a status code equal to that given
func (o *BackupImportBackupStorageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupImportBackupStorageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupImportBackupStorageUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupImportBackupStorageUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupImportBackupStorageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupImportBackupStorageForbidden creates a BackupImportBackupStorageForbidden with default headers values
func NewBackupImportBackupStorageForbidden() *BackupImportBackupStorageForbidden {
	return &BackupImportBackupStorageForbidden{}
}

/*
BackupImportBackupStorageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupImportBackupStorageForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup import backup storage forbidden response has a 2xx status code
func (o *BackupImportBackupStorageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup import backup storage forbidden response has a 3xx status code
func (o *BackupImportBackupStorageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage forbidden response has a 4xx status code
func (o *BackupImportBackupStorageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup import backup storage forbidden response has a 5xx status code
func (o *BackupImportBackupStorageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup import backup storage forbidden response a status code equal to that given
func (o *BackupImportBackupStorageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupImportBackupStorageForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageForbidden  %+v", 403, o.Payload)
}

func (o *BackupImportBackupStorageForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageForbidden  %+v", 403, o.Payload)
}

func (o *BackupImportBackupStorageForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupImportBackupStorageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupImportBackupStorageNotFound creates a BackupImportBackupStorageNotFound with default headers values
func NewBackupImportBackupStorageNotFound() *BackupImportBackupStorageNotFound {
	return &BackupImportBackupStorageNotFound{}
}

/*
BackupImportBackupStorageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupImportBackupStorageNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup import backup storage not found response has a 2xx status code
func (o *BackupImportBackupStorageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup import backup storage not found response has a 3xx status code
func (o *BackupImportBackupStorageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage not found response has a 4xx status code
func (o *BackupImportBackupStorageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup import backup storage not found response has a 5xx status code
func (o *BackupImportBackupStorageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup import backup storage not found response a status code equal to that given
func (o *BackupImportBackupStorageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupImportBackupStorageNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageNotFound  %+v", 404, o.Payload)
}

func (o *BackupImportBackupStorageNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageNotFound  %+v", 404, o.Payload)
}

func (o *BackupImportBackupStorageNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupImportBackupStorageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupImportBackupStorageInternalServerError creates a BackupImportBackupStorageInternalServerError with default headers values
func NewBackupImportBackupStorageInternalServerError() *BackupImportBackupStorageInternalServerError {
	return &BackupImportBackupStorageInternalServerError{}
}

/*
BackupImportBackupStorageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupImportBackupStorageInternalServerError struct {
}

// IsSuccess returns true when this backup import backup storage internal server error response has a 2xx status code
func (o *BackupImportBackupStorageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup import backup storage internal server error response has a 3xx status code
func (o *BackupImportBackupStorageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup import backup storage internal server error response has a 4xx status code
func (o *BackupImportBackupStorageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup import backup storage internal server error response has a 5xx status code
func (o *BackupImportBackupStorageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup import backup storage internal server error response a status code equal to that given
func (o *BackupImportBackupStorageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupImportBackupStorageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageInternalServerError ", 500)
}

func (o *BackupImportBackupStorageInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/location][%d] backupImportBackupStorageInternalServerError ", 500)
}

func (o *BackupImportBackupStorageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
