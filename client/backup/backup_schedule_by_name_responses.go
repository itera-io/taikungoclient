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

// BackupScheduleByNameReader is a Reader for the BackupScheduleByName structure.
type BackupScheduleByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupScheduleByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupScheduleByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBackupScheduleByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBackupScheduleByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBackupScheduleByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBackupScheduleByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewBackupScheduleByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBackupScheduleByNameOK creates a BackupScheduleByNameOK with default headers values
func NewBackupScheduleByNameOK() *BackupScheduleByNameOK {
	return &BackupScheduleByNameOK{}
}

/* BackupScheduleByNameOK describes a response with status code 200, with default header values.

Success
*/
type BackupScheduleByNameOK struct {
	Payload *models.ScheduleDto
}

func (o *BackupScheduleByNameOK) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameOK  %+v", 200, o.Payload)
}
func (o *BackupScheduleByNameOK) GetPayload() *models.ScheduleDto {
	return o.Payload
}

func (o *BackupScheduleByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupScheduleByNameBadRequest creates a BackupScheduleByNameBadRequest with default headers values
func NewBackupScheduleByNameBadRequest() *BackupScheduleByNameBadRequest {
	return &BackupScheduleByNameBadRequest{}
}

/* BackupScheduleByNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type BackupScheduleByNameBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *BackupScheduleByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameBadRequest  %+v", 400, o.Payload)
}
func (o *BackupScheduleByNameBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *BackupScheduleByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupScheduleByNameUnauthorized creates a BackupScheduleByNameUnauthorized with default headers values
func NewBackupScheduleByNameUnauthorized() *BackupScheduleByNameUnauthorized {
	return &BackupScheduleByNameUnauthorized{}
}

/* BackupScheduleByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BackupScheduleByNameUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *BackupScheduleByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameUnauthorized  %+v", 401, o.Payload)
}
func (o *BackupScheduleByNameUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupScheduleByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupScheduleByNameForbidden creates a BackupScheduleByNameForbidden with default headers values
func NewBackupScheduleByNameForbidden() *BackupScheduleByNameForbidden {
	return &BackupScheduleByNameForbidden{}
}

/* BackupScheduleByNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BackupScheduleByNameForbidden struct {
	Payload *models.ProblemDetails
}

func (o *BackupScheduleByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameForbidden  %+v", 403, o.Payload)
}
func (o *BackupScheduleByNameForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupScheduleByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupScheduleByNameNotFound creates a BackupScheduleByNameNotFound with default headers values
func NewBackupScheduleByNameNotFound() *BackupScheduleByNameNotFound {
	return &BackupScheduleByNameNotFound{}
}

/* BackupScheduleByNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BackupScheduleByNameNotFound struct {
	Payload *models.ProblemDetails
}

func (o *BackupScheduleByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameNotFound  %+v", 404, o.Payload)
}
func (o *BackupScheduleByNameNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *BackupScheduleByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupScheduleByNameInternalServerError creates a BackupScheduleByNameInternalServerError with default headers values
func NewBackupScheduleByNameInternalServerError() *BackupScheduleByNameInternalServerError {
	return &BackupScheduleByNameInternalServerError{}
}

/* BackupScheduleByNameInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type BackupScheduleByNameInternalServerError struct {
}

func (o *BackupScheduleByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v{v}/Backup/schedule/{projectId}/{name}][%d] backupScheduleByNameInternalServerError ", 500)
}

func (o *BackupScheduleByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
