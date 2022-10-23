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

// BackupCreateReader is a Reader for the BackupCreate structure.
type BackupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupCreateOK creates a BackupCreateOK with default headers values
func NewBackupCreateOK() *BackupCreateOK {
	return &BackupCreateOK{}
}

/*
BackupCreateOK describes a response with status code 200, with default header values.

Success
*/
type BackupCreateOK struct {
	Payload models.Unit
}

// IsSuccess returns true when this backup create o k response has a 2xx status code
func (o *BackupCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup create o k response has a 3xx status code
func (o *BackupCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create o k response has a 4xx status code
func (o *BackupCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup create o k response has a 5xx status code
func (o *BackupCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup create o k response a status code equal to that given
func (o *BackupCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *BackupCreateOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateOK  %+v", 200, o.Payload)
}

func (o *BackupCreateOK) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateOK  %+v", 200, o.Payload)
}

func (o *BackupCreateOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *BackupCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupCreateBadRequest creates a BackupCreateBadRequest with default headers values
func NewBackupCreateBadRequest() *BackupCreateBadRequest {
	return &BackupCreateBadRequest{}
}

/*
BackupCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupCreateBadRequest struct {
	Payload []*models.Error
}

// IsSuccess returns true when this backup create bad request response has a 2xx status code
func (o *BackupCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup create bad request response has a 3xx status code
func (o *BackupCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create bad request response has a 4xx status code
func (o *BackupCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup create bad request response has a 5xx status code
func (o *BackupCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this backup create bad request response a status code equal to that given
func (o *BackupCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *BackupCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateBadRequest  %+v", 400, o.Payload)
}

func (o *BackupCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateBadRequest  %+v", 400, o.Payload)
}

func (o *BackupCreateBadRequest) GetPayload() []*models.Error {
	return o.Payload
}

func (o *BackupCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupCreateUnauthorized creates a BackupCreateUnauthorized with default headers values
func NewBackupCreateUnauthorized() *BackupCreateUnauthorized {
	return &BackupCreateUnauthorized{}
}

/*
BackupCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupCreateUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup create unauthorized response has a 2xx status code
func (o *BackupCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup create unauthorized response has a 3xx status code
func (o *BackupCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create unauthorized response has a 4xx status code
func (o *BackupCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup create unauthorized response has a 5xx status code
func (o *BackupCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this backup create unauthorized response a status code equal to that given
func (o *BackupCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *BackupCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *BackupCreateUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupCreateForbidden creates a BackupCreateForbidden with default headers values
func NewBackupCreateForbidden() *BackupCreateForbidden {
	return &BackupCreateForbidden{}
}

/*
BackupCreateForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupCreateForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup create forbidden response has a 2xx status code
func (o *BackupCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup create forbidden response has a 3xx status code
func (o *BackupCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create forbidden response has a 4xx status code
func (o *BackupCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup create forbidden response has a 5xx status code
func (o *BackupCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this backup create forbidden response a status code equal to that given
func (o *BackupCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *BackupCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateForbidden  %+v", 403, o.Payload)
}

func (o *BackupCreateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateForbidden  %+v", 403, o.Payload)
}

func (o *BackupCreateForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupCreateNotFound creates a BackupCreateNotFound with default headers values
func NewBackupCreateNotFound() *BackupCreateNotFound {
	return &BackupCreateNotFound{}
}

/*
BackupCreateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupCreateNotFound struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this backup create not found response has a 2xx status code
func (o *BackupCreateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup create not found response has a 3xx status code
func (o *BackupCreateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create not found response has a 4xx status code
func (o *BackupCreateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this backup create not found response has a 5xx status code
func (o *BackupCreateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this backup create not found response a status code equal to that given
func (o *BackupCreateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *BackupCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateNotFound  %+v", 404, o.Payload)
}

func (o *BackupCreateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateNotFound  %+v", 404, o.Payload)
}

func (o *BackupCreateNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupCreateInternalServerError creates a BackupCreateInternalServerError with default headers values
func NewBackupCreateInternalServerError() *BackupCreateInternalServerError {
	return &BackupCreateInternalServerError{}
}

/*
BackupCreateInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupCreateInternalServerError struct {
}

// IsSuccess returns true when this backup create internal server error response has a 2xx status code
func (o *BackupCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this backup create internal server error response has a 3xx status code
func (o *BackupCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup create internal server error response has a 4xx status code
func (o *BackupCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup create internal server error response has a 5xx status code
func (o *BackupCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this backup create internal server error response a status code equal to that given
func (o *BackupCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *BackupCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateInternalServerError ", 500)
}

func (o *BackupCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v{v}/Backup/create][%d] backupCreateInternalServerError ", 500)
}

func (o *BackupCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
