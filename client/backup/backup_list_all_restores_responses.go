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

// BackupListAllRestoresReader is a Reader for the BackupListAllRestores structure.
type BackupListAllRestoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupListAllRestoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupListAllRestoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupListAllRestoresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupListAllRestoresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupListAllRestoresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupListAllRestoresNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupListAllRestoresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupListAllRestoresOK creates a BackupListAllRestoresOK with default headers values
func NewBackupListAllRestoresOK() *BackupListAllRestoresOK {
	return &BackupListAllRestoresOK{}
}

/*
BackupListAllRestoresOK describes a response with status code 200, with default header values.

Success
*/
type BackupListAllRestoresOK struct {
	Payload *models.ListAllRestores
}

// IsSuccess returns true when this backup list all restores o k response has a 2xx status code
func (o *BackupListAllRestoresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup list all restores o k response has a 3xx status code
func (o *BackupListAllRestoresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores o k response has a 4xx status code
func (o *BackupListAllRestoresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all restores o k response has a 5xx status code
func (o *BackupListAllRestoresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all restores o k response a status code equal to that given
func (o *BackupListAllRestoresOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupListAllRestoresOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresOK  %+v", 200, o.Payload)
}

func (o *BackupListAllRestoresOK) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresOK  %+v", 200, o.Payload)
}

func (o *BackupListAllRestoresOK) GetPayload() *models.ListAllRestores {
	return o.Payload
}

func (o *BackupListAllRestoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAllRestores)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllRestoresBadRequest creates a BackupListAllRestoresBadRequest with default headers values
func NewBackupListAllRestoresBadRequest() *BackupListAllRestoresBadRequest {
	return &BackupListAllRestoresBadRequest{}
}

/*
BackupListAllRestoresBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupListAllRestoresBadRequest struct {
	Payload *models.ValidationProblemDetails
}

// IsSuccess returns true when this backup list all restores bad request response has a 2xx status code
func (o *BackupListAllRestoresBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all restores bad request response has a 3xx status code
func (o *BackupListAllRestoresBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores bad request response has a 4xx status code
func (o *BackupListAllRestoresBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all restores bad request response has a 5xx status code
func (o *BackupListAllRestoresBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all restores bad request response a status code equal to that given
func (o *BackupListAllRestoresBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupListAllRestoresBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllRestoresBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresBadRequest  %+v", 400, o.Payload)
}

func (o *BackupListAllRestoresBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupListAllRestoresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllRestoresUnauthorized creates a BackupListAllRestoresUnauthorized with default headers values
func NewBackupListAllRestoresUnauthorized() *BackupListAllRestoresUnauthorized {
	return &BackupListAllRestoresUnauthorized{}
}

/*
BackupListAllRestoresUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupListAllRestoresUnauthorized struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all restores unauthorized response has a 2xx status code
func (o *BackupListAllRestoresUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all restores unauthorized response has a 3xx status code
func (o *BackupListAllRestoresUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores unauthorized response has a 4xx status code
func (o *BackupListAllRestoresUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all restores unauthorized response has a 5xx status code
func (o *BackupListAllRestoresUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all restores unauthorized response a status code equal to that given
func (o *BackupListAllRestoresUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupListAllRestoresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllRestoresUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupListAllRestoresUnauthorized) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllRestoresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllRestoresForbidden creates a BackupListAllRestoresForbidden with default headers values
func NewBackupListAllRestoresForbidden() *BackupListAllRestoresForbidden {
	return &BackupListAllRestoresForbidden{}
}

/*
BackupListAllRestoresForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupListAllRestoresForbidden struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all restores forbidden response has a 2xx status code
func (o *BackupListAllRestoresForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all restores forbidden response has a 3xx status code
func (o *BackupListAllRestoresForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores forbidden response has a 4xx status code
func (o *BackupListAllRestoresForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all restores forbidden response has a 5xx status code
func (o *BackupListAllRestoresForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all restores forbidden response a status code equal to that given
func (o *BackupListAllRestoresForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupListAllRestoresForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllRestoresForbidden) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresForbidden  %+v", 403, o.Payload)
}

func (o *BackupListAllRestoresForbidden) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllRestoresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllRestoresNotFound creates a BackupListAllRestoresNotFound with default headers values
func NewBackupListAllRestoresNotFound() *BackupListAllRestoresNotFound {
	return &BackupListAllRestoresNotFound{}
}

/*
BackupListAllRestoresNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupListAllRestoresNotFound struct {
	Payload []*models.CustomProblemDetailsMg
}

// IsSuccess returns true when this backup list all restores not found response has a 2xx status code
func (o *BackupListAllRestoresNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all restores not found response has a 3xx status code
func (o *BackupListAllRestoresNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores not found response has a 4xx status code
func (o *BackupListAllRestoresNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup list all restores not found response has a 5xx status code
func (o *BackupListAllRestoresNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup list all restores not found response a status code equal to that given
func (o *BackupListAllRestoresNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupListAllRestoresNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllRestoresNotFound) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresNotFound  %+v", 404, o.Payload)
}

func (o *BackupListAllRestoresNotFound) GetPayload() []*models.CustomProblemDetailsMg {
	return o.Payload
}

func (o *BackupListAllRestoresNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupListAllRestoresInternalServerError creates a BackupListAllRestoresInternalServerError with default headers values
func NewBackupListAllRestoresInternalServerError() *BackupListAllRestoresInternalServerError {
	return &BackupListAllRestoresInternalServerError{}
}

/*
BackupListAllRestoresInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupListAllRestoresInternalServerError struct {
}

// IsSuccess returns true when this backup list all restores internal server error response has a 2xx status code
func (o *BackupListAllRestoresInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup list all restores internal server error response has a 3xx status code
func (o *BackupListAllRestoresInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup list all restores internal server error response has a 4xx status code
func (o *BackupListAllRestoresInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup list all restores internal server error response has a 5xx status code
func (o *BackupListAllRestoresInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup list all restores internal server error response a status code equal to that given
func (o *BackupListAllRestoresInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupListAllRestoresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresInternalServerError ", 500)
}

func (o *BackupListAllRestoresInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/restores/{projectId}][%d] backupListAllRestoresInternalServerError ", 500)
}

func (o *BackupListAllRestoresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
