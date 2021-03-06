// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// SearchBackupCredentialsListReader is a Reader for the SearchBackupCredentialsList structure.
type SearchBackupCredentialsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchBackupCredentialsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchBackupCredentialsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchBackupCredentialsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchBackupCredentialsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchBackupCredentialsListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchBackupCredentialsListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchBackupCredentialsListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchBackupCredentialsListOK creates a SearchBackupCredentialsListOK with default headers values
func NewSearchBackupCredentialsListOK() *SearchBackupCredentialsListOK {
	return &SearchBackupCredentialsListOK{}
}

/* SearchBackupCredentialsListOK describes a response with status code 200, with default header values.

Success
*/
type SearchBackupCredentialsListOK struct {
	Payload *models.BackupCredentialsSearchList
}

func (o *SearchBackupCredentialsListOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListOK  %+v", 200, o.Payload)
}
func (o *SearchBackupCredentialsListOK) GetPayload() *models.BackupCredentialsSearchList {
	return o.Payload
}

func (o *SearchBackupCredentialsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupCredentialsSearchList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBackupCredentialsListBadRequest creates a SearchBackupCredentialsListBadRequest with default headers values
func NewSearchBackupCredentialsListBadRequest() *SearchBackupCredentialsListBadRequest {
	return &SearchBackupCredentialsListBadRequest{}
}

/* SearchBackupCredentialsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SearchBackupCredentialsListBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *SearchBackupCredentialsListBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListBadRequest  %+v", 400, o.Payload)
}
func (o *SearchBackupCredentialsListBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *SearchBackupCredentialsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBackupCredentialsListUnauthorized creates a SearchBackupCredentialsListUnauthorized with default headers values
func NewSearchBackupCredentialsListUnauthorized() *SearchBackupCredentialsListUnauthorized {
	return &SearchBackupCredentialsListUnauthorized{}
}

/* SearchBackupCredentialsListUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchBackupCredentialsListUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *SearchBackupCredentialsListUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchBackupCredentialsListUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBackupCredentialsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBackupCredentialsListForbidden creates a SearchBackupCredentialsListForbidden with default headers values
func NewSearchBackupCredentialsListForbidden() *SearchBackupCredentialsListForbidden {
	return &SearchBackupCredentialsListForbidden{}
}

/* SearchBackupCredentialsListForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchBackupCredentialsListForbidden struct {
	Payload *models.ProblemDetails
}

func (o *SearchBackupCredentialsListForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListForbidden  %+v", 403, o.Payload)
}
func (o *SearchBackupCredentialsListForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBackupCredentialsListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBackupCredentialsListNotFound creates a SearchBackupCredentialsListNotFound with default headers values
func NewSearchBackupCredentialsListNotFound() *SearchBackupCredentialsListNotFound {
	return &SearchBackupCredentialsListNotFound{}
}

/* SearchBackupCredentialsListNotFound describes a response with status code 404, with default header values.

Not Found
*/
type SearchBackupCredentialsListNotFound struct {
	Payload *models.ProblemDetails
}

func (o *SearchBackupCredentialsListNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListNotFound  %+v", 404, o.Payload)
}
func (o *SearchBackupCredentialsListNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *SearchBackupCredentialsListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchBackupCredentialsListInternalServerError creates a SearchBackupCredentialsListInternalServerError with default headers values
func NewSearchBackupCredentialsListInternalServerError() *SearchBackupCredentialsListInternalServerError {
	return &SearchBackupCredentialsListInternalServerError{}
}

/* SearchBackupCredentialsListInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchBackupCredentialsListInternalServerError struct {
}

func (o *SearchBackupCredentialsListInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/Search/backup-credentials][%d] searchBackupCredentialsListInternalServerError ", 500)
}

func (o *SearchBackupCredentialsListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
