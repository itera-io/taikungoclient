// Code generated by go-swagger; DO NOT EDIT.

package project_revisions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/itera-io/taikungoclient/models"
)

// ProjectRevisionsEditReader is a Reader for the ProjectRevisionsEdit structure.
type ProjectRevisionsEditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectRevisionsEditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectRevisionsEditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectRevisionsEditBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectRevisionsEditUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectRevisionsEditForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProjectRevisionsEditNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectRevisionsEditInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProjectRevisionsEditOK creates a ProjectRevisionsEditOK with default headers values
func NewProjectRevisionsEditOK() *ProjectRevisionsEditOK {
	return &ProjectRevisionsEditOK{}
}

/* ProjectRevisionsEditOK describes a response with status code 200, with default header values.

Success
*/
type ProjectRevisionsEditOK struct {
	Payload models.Unit
}

func (o *ProjectRevisionsEditOK) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditOK  %+v", 200, o.Payload)
}
func (o *ProjectRevisionsEditOK) GetPayload() models.Unit {
	return o.Payload
}

func (o *ProjectRevisionsEditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectRevisionsEditBadRequest creates a ProjectRevisionsEditBadRequest with default headers values
func NewProjectRevisionsEditBadRequest() *ProjectRevisionsEditBadRequest {
	return &ProjectRevisionsEditBadRequest{}
}

/* ProjectRevisionsEditBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ProjectRevisionsEditBadRequest struct {
	Payload *models.ValidationProblemDetails
}

func (o *ProjectRevisionsEditBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditBadRequest  %+v", 400, o.Payload)
}
func (o *ProjectRevisionsEditBadRequest) GetPayload() *models.ValidationProblemDetails {
	return o.Payload
}

func (o *ProjectRevisionsEditBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectRevisionsEditUnauthorized creates a ProjectRevisionsEditUnauthorized with default headers values
func NewProjectRevisionsEditUnauthorized() *ProjectRevisionsEditUnauthorized {
	return &ProjectRevisionsEditUnauthorized{}
}

/* ProjectRevisionsEditUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ProjectRevisionsEditUnauthorized struct {
	Payload *models.ProblemDetails
}

func (o *ProjectRevisionsEditUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditUnauthorized  %+v", 401, o.Payload)
}
func (o *ProjectRevisionsEditUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectRevisionsEditUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectRevisionsEditForbidden creates a ProjectRevisionsEditForbidden with default headers values
func NewProjectRevisionsEditForbidden() *ProjectRevisionsEditForbidden {
	return &ProjectRevisionsEditForbidden{}
}

/* ProjectRevisionsEditForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ProjectRevisionsEditForbidden struct {
	Payload *models.ProblemDetails
}

func (o *ProjectRevisionsEditForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditForbidden  %+v", 403, o.Payload)
}
func (o *ProjectRevisionsEditForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectRevisionsEditForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectRevisionsEditNotFound creates a ProjectRevisionsEditNotFound with default headers values
func NewProjectRevisionsEditNotFound() *ProjectRevisionsEditNotFound {
	return &ProjectRevisionsEditNotFound{}
}

/* ProjectRevisionsEditNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ProjectRevisionsEditNotFound struct {
	Payload *models.ProblemDetails
}

func (o *ProjectRevisionsEditNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditNotFound  %+v", 404, o.Payload)
}
func (o *ProjectRevisionsEditNotFound) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *ProjectRevisionsEditNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectRevisionsEditInternalServerError creates a ProjectRevisionsEditInternalServerError with default headers values
func NewProjectRevisionsEditInternalServerError() *ProjectRevisionsEditInternalServerError {
	return &ProjectRevisionsEditInternalServerError{}
}

/* ProjectRevisionsEditInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ProjectRevisionsEditInternalServerError struct {
}

func (o *ProjectRevisionsEditInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v{v}/ProjectRevisions/update/{projectId}][%d] projectRevisionsEditInternalServerError ", 500)
}

func (o *ProjectRevisionsEditInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
