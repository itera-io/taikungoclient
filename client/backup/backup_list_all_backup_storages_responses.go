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

// BackupListAllBackupStoragesReader is a Reader for the BackupListAllBackupStorages structure.
type BackupListAllBackupStoragesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupListAllBackupStoragesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupListAllBackupStoragesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupListAllBackupStoragesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupListAllBackupStoragesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupListAllBackupStoragesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupListAllBackupStoragesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupListAllBackupStoragesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupListAllBackupStoragesOK creates a BackupListAllBackupStoragesOK with default headers values
func NewBackupListAllBackupStoragesOK() *BackupListAllBackupStoragesOK {
	return &BackupListAllBackupStoragesOK{}
}

/*
BackupListAllBackupStoragesOK describes a response with status code 200, with default header values.

Success
*/
type BackupListAllBackupStoragesOK struct {
	Payload *models.ListAllBackupStorageLocations
}

// IsSuccess returns true when this backup list all backup storages o k response has a 2xx status code
func (o *BackupListAllBackupStoragesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup list all backup storages o k response has a 3xx status code
func (o *BackupListAllBackupStoragesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages o k response has a 4xx status code
func (o *BackupListAllBackupStoragesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all backup storages o k response has a 5xx status code
func (o *BackupListAllBackupStoragesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all backup storages o k response a status code equal to that given
func (o *BackupListAllBackupStoragesOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupListAllBackupStoragesOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesOK  %+v", 200, o.Payload)
}

func (o *BackupListAllBackupStoragesOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesOK  %+v", 200, o.Payload)
}

func (o *BackupListAllBackupStoragesOK) GetPayload() *models.ListAllBackupStorageLocations {
	return o.Payload
}

func (o *BackupListAllBackupStoragesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAllBackupStorageLocations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllBackupStoragesBadRequest creates a BackupListAllBackupStoragesBadRequest with default headers values
func NewBackupListAllBackupStoragesBadRequest() *BackupListAllBackupStoragesBadRequest {
	return &BackupListAllBackupStoragesBadRequest{}
}

/*
BackupListAllBackupStoragesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupListAllBackupStoragesBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup list all backup storages bad request response has a 2xx status code
func (o *BackupListAllBackupStoragesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all backup storages bad request response has a 3xx status code
func (o *BackupListAllBackupStoragesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages bad request response has a 4xx status code
func (o *BackupListAllBackupStoragesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all backup storages bad request response has a 5xx status code
func (o *BackupListAllBackupStoragesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all backup storages bad request response a status code equal to that given
func (o *BackupListAllBackupStoragesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupListAllBackupStoragesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllBackupStoragesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllBackupStoragesBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupListAllBackupStoragesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllBackupStoragesUnauthorized creates a BackupListAllBackupStoragesUnauthorized with default headers values
func NewBackupListAllBackupStoragesUnauthorized() *BackupListAllBackupStoragesUnauthorized {
	return &BackupListAllBackupStoragesUnauthorized{}
}

/*
BackupListAllBackupStoragesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupListAllBackupStoragesUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup list all backup storages unauthorized response has a 2xx status code
func (o *BackupListAllBackupStoragesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all backup storages unauthorized response has a 3xx status code
func (o *BackupListAllBackupStoragesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages unauthorized response has a 4xx status code
func (o *BackupListAllBackupStoragesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all backup storages unauthorized response has a 5xx status code
func (o *BackupListAllBackupStoragesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all backup storages unauthorized response a status code equal to that given
func (o *BackupListAllBackupStoragesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupListAllBackupStoragesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllBackupStoragesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllBackupStoragesUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupListAllBackupStoragesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllBackupStoragesForbidden creates a BackupListAllBackupStoragesForbidden with default headers values
func NewBackupListAllBackupStoragesForbidden() *BackupListAllBackupStoragesForbidden {
	return &BackupListAllBackupStoragesForbidden{}
}

/*
BackupListAllBackupStoragesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupListAllBackupStoragesForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup list all backup storages forbidden response has a 2xx status code
func (o *BackupListAllBackupStoragesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all backup storages forbidden response has a 3xx status code
func (o *BackupListAllBackupStoragesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages forbidden response has a 4xx status code
func (o *BackupListAllBackupStoragesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all backup storages forbidden response has a 5xx status code
func (o *BackupListAllBackupStoragesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all backup storages forbidden response a status code equal to that given
func (o *BackupListAllBackupStoragesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupListAllBackupStoragesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllBackupStoragesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllBackupStoragesForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupListAllBackupStoragesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllBackupStoragesNotFound creates a BackupListAllBackupStoragesNotFound with default headers values
func NewBackupListAllBackupStoragesNotFound() *BackupListAllBackupStoragesNotFound {
	return &BackupListAllBackupStoragesNotFound{}
}

/*
BackupListAllBackupStoragesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupListAllBackupStoragesNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup list all backup storages not found response has a 2xx status code
func (o *BackupListAllBackupStoragesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all backup storages not found response has a 3xx status code
func (o *BackupListAllBackupStoragesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages not found response has a 4xx status code
func (o *BackupListAllBackupStoragesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all backup storages not found response has a 5xx status code
func (o *BackupListAllBackupStoragesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all backup storages not found response a status code equal to that given
func (o *BackupListAllBackupStoragesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupListAllBackupStoragesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllBackupStoragesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllBackupStoragesNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupListAllBackupStoragesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllBackupStoragesInternalServerError creates a BackupListAllBackupStoragesInternalServerError with default headers values
func NewBackupListAllBackupStoragesInternalServerError() *BackupListAllBackupStoragesInternalServerError {
	return &BackupListAllBackupStoragesInternalServerError{}
}

/*
BackupListAllBackupStoragesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupListAllBackupStoragesInternalServerError struct {
}

// IsSuccess returns true when this backup list all backup storages internal server error response has a 2xx status code
func (o *BackupListAllBackupStoragesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all backup storages internal server error response has a 3xx status code
func (o *BackupListAllBackupStoragesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all backup storages internal server error response has a 4xx status code
func (o *BackupListAllBackupStoragesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all backup storages internal server error response has a 5xx status code
func (o *BackupListAllBackupStoragesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup list all backup storages internal server error response a status code equal to that given
func (o *BackupListAllBackupStoragesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupListAllBackupStoragesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesInternalServerError ", 500)
}

func (o *BackupListAllBackupStoragesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/location/{projectId}][%d] backupListAllBackupStoragesInternalServerError ", 500)
}

func (o *BackupListAllBackupStoragesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
